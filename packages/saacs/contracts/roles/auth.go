package roles

import (
	"log/slog"

	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	"github.com/nova38/thesis/packages/saacs/policy"
	"github.com/nova38/thesis/packages/saacs/state"
	"github.com/samber/lo"
	"github.com/samber/oops"
)

func (ctx *RolesTxCtx) Authorize(ops []*authpb.Operation) (bool, error) {
	ctx.GetLogger().Info("NoAuthContract.Authenticate")

	for _, op := range ops {
		if auth, err := ctx.authorized(op); err != nil {
			return false, oops.Wrap(err)
		} else if !auth {
			ctx.Logger.Info("User is not authorized")
			return false, nil
		}
	}

	return true, nil
}

func (ctx *RolesTxCtx) authorized(op *authpb.Operation) (bool, error) {
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

func (ctx *RolesTxCtx) checkParents(
	roles []*authpb.Role,
	checked []string,
	op *authpb.Operation,
) (bool, []string, error) {

	parents := []string{}

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

func (ctx *RolesTxCtx) getUserRoles(collectionId string) ([]*authpb.Role, error) {

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
		UserId:       user.GetMspId(),
		RoleIds:      []string{},
	}

	if err := (state.Ledger[*authpb.UserCollectionRoles]{}.PrimaryGet(ctx, userRoles)); err != nil {
		return nil, oops.Wrap(err)
	}

	roles, err := ctx.getRoles(userRoles.GetRoleIds())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	ctx.UserRoles[collectionId] = roles

	return roles, nil
}

func (ctx *RolesTxCtx) getRoles(roleIds []string) ([]*authpb.Role, error) {
	roles := []*authpb.Role{}

	for _, roleId := range roleIds {
		role := &authpb.Role{
			RoleId: roleId,
		}

		err := state.Ledger[*authpb.Role]{}.PrimaryGet(ctx, role)
		if err != nil {
			return nil, oops.Wrap(err)
		}

		roles = append(roles, role)
	}

	return roles, nil
}

func (ctx *RolesTxCtx) getNewParents(checked []string, parents []string) ([]*authpb.Role, error) {

	parents = lo.Uniq(parents)

	unckecked, _ := lo.Difference(parents, checked)

	if len(unckecked) == 0 {
		return nil, nil
	}

	roles := []*authpb.Role{}

	for _, roleId := range unckecked {
		role := &authpb.Role{
			RoleId: roleId,
		}

		if err := (state.Ledger[*authpb.Role]{}.PrimaryGet(ctx, role)); err != nil {
			return nil, oops.Wrap(err)
		}

		roles = append(roles, role)
	}

	return roles, nil
}

// func (ctx *RolesTxCtx) getRoleParent(checked []string, role *authpb.Role) ([]*authpb.Role, []string, error) {
// 	if role.GetParentRoleIds() == nil || len(role.GetParentRoleIds()) == 0 {
// 		return nil, nil, nil
// 	}

// 	parents := []*authpb.Role{}

// 	newIds, already := lo.Difference(role.GetParentRoleIds(), checked)

// 	ctx.Logger.Debug("New Parrent Roles", "newIds", newIds)
// 	ctx.Logger.Debug("Already Checked", "already", already)

// 	for _, parrentId := range role.GetParentRoleIds() {

// 		parent := &authpb.Role{
// 			CollectionId: role.GetCollectionId(),
// 			RoleId:       parrentId,
// 		}

// 		parrentKey, err := common.MakePrimaryKey(parent)
// 		if err != nil {
// 			return nil, nil, oops.Wrap(err)
// 		}

// 		ctx.Logger.Debug("Parrent Key", "parrentKey", parrentKey)
// 		if err := state.Get(ctx, parrentKey, parent); err != nil {
// 			return nil, nil, oops.Wrap(err)
// 		}

// 		parents = append(parents, parent)
// 	}

// 	ids := lo.Union(checked, role.GetParentRoleIds())

// 	return parents, ids, nil
// }

// // checkRoles recursively checks the roles and their parents,
// // it keeps track of the roles that have already been checked to avoid inheritance loops
// func (ctx *RolesTxCtx) checkRoles(unchecked []*authpb.Role, checked []string, op *authpb.Operation) (bool, []string, error) {
// 	if len(unchecked) == 0 {
// 		return false, nil, nil
// 	}

// 	// Check last batch of roles
// 	for _, role := range unchecked {
// 		ctx.Logger.Info("Role", "role", role)
// 		if auth, err := policy.AuthorizedPolicy(role.GetPolices(), op); err != nil {
// 			return false, nil, oops.Wrap(err)
// 		} else if auth {
// 			ctx.Logger.Info("User is authorized by role")
// 			return true, nil, nil
// 		}

// 		// Add the role to the checked list
// 		checked = append(checked, role.GetRoleId())
// 	}

// 	// Get the roles parents

// 	parentRoles := []*authpb.Role{}

// 	fetched := slices.Clone(checked)

//     lo.ForEach(unchecked, func(role *authpb.Role, i int) {
//         parents, newIds, err := ctx.getRoleParent(fetched, role)
//         if err != nil {

// 	for _, role := range unchecked {
// 		var parents []*authpb.Role
// 		parents, roleIds, err = ctx.getRoleParent(roleIds, role)
// 		if err != nil {
// 			return false, oops.Wrap(err)
// 		}
// 		parentRoles = append(parentRoles, parents...)
// 	}

// 	// Add the role to the checked list
// 	checked = append(checked, role.GetRoleId())

// 	// Check the role's parents
// 	for _, parrentId := range role.GetParentRoleIds() {

// 		parent := &authpb.Role{
// 			CollectionId: role.GetCollectionId(),
// 			RoleId:       parrentId,
// 		}

// 		parrentKey, err := common.MakePrimaryKey(parent)
// 		if err != nil {
// 			return false, oops.Wrap(err)
// 		}

// 		ctx.Logger.Debug("Parrent Key", "parrentKey", parrentKey)
// 		if err := state.Get(ctx, parrentKey, parent); err != nil {
// 			return false, oops.Wrap(err)
// 		}

// 		if auth, err := policy.AuthorizedPolicy(parent.GetPolices(), op); err != nil {
// 			return false, oops.Wrap(err)
// 		} else if auth {
// 			ctx.Logger.Info("User is authorized by role")
// 			return true, nil
// 		}

// 		// Add the role to the checked list
// 		checked = append(checked, parent.GetRoleId())
// 	}

// 	return false, nil
// }
