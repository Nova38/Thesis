package policy

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
)

// TODO: Implement RBAC authentification contract

// Get The roles of the user for the given collection

func GetUserRoles(ctx common.TxCtxInterface, collectionId string) ([]*authpb.Role, error) {
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

	// roles := []*authpb.Role{}
	key, err := common.MakePrimaryKey(userRoles)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if err = state.Get(ctx, key, userRoles); err != nil {
		return nil, oops.Wrap(err)
	}

	roles := []*authpb.Role{}

	for _, roleId := range userRoles.GetRoleIds() {
		role := &authpb.Role{
			RoleId: roleId,
		}
		key, err := common.MakePrimaryKey(role)
		if err != nil {
			return nil, oops.Wrap(err)
		}
		if err = state.Get(ctx, key, role); err != nil {
			return nil, oops.Wrap(err)
		}

		roles = append(roles, role)
	}

	return roles, nil
}

func GetRoleParent(ctx common.TxCtxInterface, role *authpb.Role) ([]*authpb.Role, error) {
	if role.GetParentRoleIds() == nil || len(role.GetParentRoleIds()) == 0 {
		return nil, nil
	}

	parents := []*authpb.Role{}

	for _, parrent_id := range role.GetParentRoleIds() {

		parent := &authpb.Role{
			CollectionId: role.GetCollectionId(),
			RoleId:       parrent_id,
		}

		parrentKey, err := common.MakePrimaryKey(parent)
		if err != nil {
			return nil, oops.Wrap(err)
		}

		if err := state.Get(ctx, parrentKey, parent); err != nil {
			return nil, oops.Wrap(err)
		}

		parents = append(parents, parent)
	}

	return parents, nil
}

func roleGetPolicyForItem(
	ctx common.TxCtxInterface,
	role *authpb.Role,
	itemType string,
) (itemAcl *authpb.PathPolicy, found bool, err error) {
	// Get the roles of the user for the given collection
	acl := role.GetPolices()
	if acl == nil || acl.GetItemPolicies() == nil {
		return nil, false, oops.Errorf("role acl is nil")
	}

	policies := acl.GetItemPolicies()

	itemAcl, ok := policies[itemType]
	if !ok {
		return nil, false, nil
	}
	if itemAcl == nil {
		return nil, false, oops.Errorf("item acl is nil")
	}

	return itemAcl, true, nil
}

func GetRoleACL(ctx common.TxCtxInterface, role *authpb.Role) (*authpb.Polices, error) {
	if role.GetPolices() == nil {
		return nil, oops.Errorf("role ac is nil")
	}

	return role.GetPolices(), nil
}

func RoleAuthorizeOperation(ctx common.TxCtxInterface, ops []*authpb.Operation) (bool, error) {
	ctx.GetLogger().Info("RoleAuthorizeOperations", "ops", ops)

	opsByCol := lo.GroupBy(ops, func(op *authpb.Operation) string {
		return op.GetCollectionId()
	})

	for collectionId, ops := range opsByCol {
		ctx.GetLogger().Info("RoleAuthorizeOperations", "collectionId", collectionId)

		// Get the roles of the user for the given collection
		roles, err := GetUserRoles(ctx, collectionId)
		if err != nil {
			ctx.GetLogger().Error("RoleAuthorizeOperations", "err", err)
			return false, oops.Wrap(err)
		}

		ctx.GetLogger().Info("RoleAuthorizeOperations", "roles", roles)

		// ctx.GetLogger().Info("RoleAuthorizeOperations", "ops", ops)
		for _, op := range ops {
			ctx.GetLogger().Info("RoleAuthorizeOperations", "op", op)

			// Get the roles of the user for the given collection

			// itemPolicy, found, err := roleGetPolicyForItem(ctx, role, op.GetItemType())
			if err != nil {
				return false, oops.Wrap(err)
			}
			// if !found {
			// 	// No policy found for this item type in this role
			// 	// todo: check if there is a policy for the parent role
			// 	continue
			// }

			// if policy.Allowed(itemPolicy, op.GetAction()) {
			// 	return true, nil
			// }

		}
	}

	return false, nil
}
