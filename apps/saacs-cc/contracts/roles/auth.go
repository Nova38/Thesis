package roles

import (
	"log/slog"

	"github.com/nova38/saacs/apps/saacs-cc/auth/policy"
	"github.com/nova38/saacs/apps/saacs-cc/state"
	authpb "github.com/nova38/saacs/libs/saacs-protos-go/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
)

func (ctx *TxCtx) Authorize(ops []*authpb.Operation) (bool, error) {
	ctx.GetLogger().Info("Roles.Authenticate")
	ctx.GetLogger().Info("Roles.Authenticate", slog.Group("auth", "ops", ops))
	for _, op := range ops {
		if auth, err := ctx.authorized(op); err != nil {
			return false, oops.With("op", op).Wrap(err)

		} else if !auth {
			ctx.Logger.Info("User is not authorized")
			return false, nil
		}
	}

	return true, nil
}

func (ctx *TxCtx) authorized(op *authpb.Operation) (bool, error) {
	ctx.GetLogger().Info(op.String())

	// Handle special case of creating a collection
	if op.GetItemType() == "auth.Collection" {
		if op.GetAction() == authpb.Action_ACTION_CREATE {
			ctx.Logger.Info(
				"User is authorized to create a collection",
				slog.Group("auth", "collection", op.GetCollectionId()),
			)
			return true, nil
		}
	}

	// Get the collection
	col, err := ctx.GetCollection(op.GetCollectionId())
	if err != nil {
		return false, oops.Wrap(err)
	}

	// Validate the operation
	if valid, err := policy.ValidateOperation(col, op); err != nil {
		return false, oops.Hint("Invalid Operation").Wrap(err)
	} else if !valid {
		return false, nil
	}

	// TODO: Check if the action is allowed for any user in the collection

	// ═════════════════════════════════════════════
	// Default Policy
	// ═════════════════════════════════════════════

	switch auth, err := policy.AuthorizedPolicy(col.GetDefault(), op); {
	case err != nil:
		return false, oops.Wrap(err)
	case auth:
		ctx.Logger.Info("User is authorized by default")
		return true, nil
	}

	// ═════════════════════════════════════════════
	// Check Role Membership
	// ═════════════════════════════════════════════
	roles, err := ctx.getUserRoles(op.GetCollectionId())
	if err != nil {
		return false, oops.Wrap(err)
	}

	for _, role := range roles {
		ctx.Logger.Info("Role", "role", role)
		if auth, err := policy.AuthorizedPolicy(role.GetPolices(), op); err != nil {
			return false, oops.Wrap(err)
		} else if auth {
			ctx.Logger.Info("User is authorized by role")
			return true, nil
		}
	}

	if auth, checked, err := ctx.checkParents(roles, []string{}, op); err != nil {
		return false, oops.Wrap(err)
	} else if auth {
		ctx.Logger.Info("User is authorized by parrent role")
		return true, nil
	} else {
		ctx.Logger.Info("User is not authorized by parrent role", "checked", checked)
	}

	return false, nil
}

func (ctx *TxCtx) checkParents(
	roles []*authpb.Role,
	checked []string,
	op *authpb.Operation,
) (bool, []string, error) {

	var parents []string

	for _, role := range roles {
		parents = append(parents, role.GetParentRoleIds()...)
	}

	parents = lo.Uniq(parents)
	ctx.Logger.Debug("New Parent RolesIds To check", "parents", parents)

	parentRoles, err := ctx.getNewParents(checked, parents)
	if err != nil {
		return false, nil, oops.Hint("Failed to get new parrents").Wrap(err)
	}

	for _, role := range parentRoles {
		if auth, err := policy.AuthorizedPolicy(role.GetPolices(), op); err != nil {
			return false, nil, oops.Wrap(err)
		} else if auth {
			ctx.Logger.Info("User is authorized by parrent role", slog.Group("auth", slog.Any("role", role), slog.Any("op", op)))
			return true, checked, nil
		}
		checked = append(checked, role.GetRoleId())

		var auth bool

		auth, checked, err = ctx.checkParents([]*authpb.Role{role}, checked, op)
		if err != nil {
			return false, checked, oops.Wrap(err)
		}

		if auth {
			return true, checked, nil
		}
	}

	return false, checked, nil
}

func (ctx *TxCtx) getUserRoles(collectionId string) ([]*authpb.Role, error) {

	if collectionId == "" {
		return nil, oops.Errorf("collectionId is empty")
	}

	// If we chached the user roles, return them
	if ctx.UserRoles == nil {
		ctx.UserRoles = map[string][]*authpb.Role{}
	} else if roles, ok := ctx.UserRoles[collectionId]; ok {
		return roles, nil
	}

	user, err := ctx.GetUserId()
	if err != nil {
		return nil, oops.Wrap(err)
	}

	userRoles := &authpb.UserCollectionRoles{
		CollectionId: collectionId,
		MspId:        user.GetMspId(),
		UserId:       user.GetUserId(),
		RoleIds:      []string{},
	}

	if err := state.Get(ctx, userRoles); err != nil {
		return nil, oops.Wrap(err)
	}

	roles, err := ctx.getRoles(collectionId, userRoles.GetRoleIds())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	ctx.UserRoles[collectionId] = roles

	return roles, nil
}

func (ctx *TxCtx) getRoles(collectionId string, roleIds []string) ([]*authpb.Role, error) {
	var roles []*authpb.Role

	for _, roleId := range roleIds {
		role := &authpb.Role{
			CollectionId: collectionId,
			RoleId:       roleId,
		}

		if err := state.Get(ctx, role); err != nil {
			return nil, oops.Wrap(err)
		}

		roles = append(roles, role)
	}

	return roles, nil
}

func (ctx *TxCtx) getNewParents(checked []string, parents []string) ([]*authpb.Role, error) {

	parents = lo.Uniq(parents)

	unckecked, _ := lo.Difference(parents, checked)

	if len(unckecked) == 0 {
		return nil, nil
	}

	var roles []*authpb.Role

	for _, roleId := range unckecked {
		role := &authpb.Role{
			RoleId: roleId,
		}

		if err := state.Get(ctx, role); err != nil {
			return nil, oops.Wrap(err)
		}

		roles = append(roles, role)
	}

	return roles, nil
}
