package collections

import (
	"log/slog"

	authpb "github.com/nova38/saacs/pkg/saacs-protos/saacs/auth/v0"
	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"

	"github.com/nova38/saacs/pkg/saacs-cc/common"
	"github.com/nova38/saacs/pkg/saacs-cc/state"
	"github.com/samber/lo"
	"github.com/samber/oops"
)

type GlobalRBAC struct{}

func (c *GlobalRBAC) BootstrapGlobal(
	ctx common.TxCtxInterface,
	collectionId string,
) (bootstraped bool, err error) {
	defer func() { ctx.HandleFnError(nil, recover()) }()
	ctx.GetLogger().Info("EmbedRoleContract.BootstrapGlobal")

	col := &authpb.Collection{
		CollectionId: "global",
	}

	// See if the global role collection already exists
	if err := state.Get(ctx, col); err != nil {
		// Collection does not exist, create it
		col.AuthType = authpb.AuthType_AUTH_TYPE_GLOBAL_ROLE
		col.CollectionId = "global"
		col.ItemTypes = []string{
			common.UserGlobalRoles,
			common.RoleItemType,
		}
		col.UseAuthParents = false
		col.Default = &authpb.Polices{
			ItemPolicies: map[string]*authpb.PathPolicy{},
			DefaultPolicy: &authpb.PathPolicy{
				Path:          "",
				FullPath:      "",
				AllowSubPaths: false,
				SubPaths:      map[string]*authpb.PathPolicy{},
				Actions: []pb.Action{
					pb.Action_ACTION_VIEW,
					pb.Action_ACTION_SUGGEST_VIEW,
					pb.Action_ACTION_VIEW_HISTORY,
				},
			},
			DefaultExcludedTypes: []string{},
		}

		// Create the collection
		if err = (state.Ledger[*authpb.Collection]{}.PrimaryCreate(ctx, col)); err != nil {
			return false, oops.
				Hint("Creating New Collection failed").
				With("collectionId", col.GetCollectionId(),
					"AuthType", authpb.AuthType_AUTH_TYPE_GLOBAL_ROLE,
				).
				Wrap(err)
		}

		// Create the manager role object

		role := &authpb.Role{
			CollectionId: col.GetCollectionId(),
			RoleId:       "manager",
			Polices: &authpb.Polices{
				ItemPolicies: map[string]*authpb.PathPolicy{},
				DefaultPolicy: &authpb.PathPolicy{
					Path:          "",
					FullPath:      "",
					AllowSubPaths: false,
					Actions: []pb.Action{
						pb.Action_ACTION_UTILITY,
						pb.Action_ACTION_VIEW,
						pb.Action_ACTION_CREATE,
						pb.Action_ACTION_UPDATE,
						pb.Action_ACTION_DELETE,
						pb.Action_ACTION_SUGGEST_VIEW,
						pb.Action_ACTION_SUGGEST_CREATE,
						pb.Action_ACTION_SUGGEST_DELETE,
						pb.Action_ACTION_SUGGEST_APPROVE,
						pb.Action_ACTION_VIEW_HISTORY,
						pb.Action_ACTION_VIEW_HIDDEN_TXS,
						pb.Action_ACTION_HIDE_TX,
					},
				},
				// Exclude the auth types from the default policy
			},
			Note:          "Default Admin Role",
			ParentRoleIds: []string{},
		}

		if err = (state.Ledger[*authpb.Role]{}.PrimaryCreate(ctx, role)); err != nil {
			return false, oops.Wrap(err)
		}

		user, err := ctx.GetUserId()
		if err != nil {
			return false, oops.With("Failed", "at making global manger role").Wrap(err)
		}

		userState := &authpb.UserGlobalRoles{
			CollectionId: "global",
			MspId:        user.GetMspId(),
			UserId:       user.GetUserId(),
			Roles: map[string]*authpb.RoleIDList{
				col.GetCollectionId(): {RoleId: []string{"manager"}},
				"global":              {RoleId: []string{"manager"}},
			},
		}

		if err = (state.Ledger[*authpb.UserGlobalRoles]{}.PrimaryCreate(ctx, userState)); err != nil {
			return false, oops.With("Failed", "to make global user role").Wrap(err)
		}

		return true, nil

	}

	// Authorize the updating of the current user's global roles
	op := &pb.Operation{
		Action:       pb.Action_ACTION_UPDATE,
		CollectionId: "global",
		ItemType:     common.UserGlobalRoles,
		Paths:        nil,
	}
	if auth, err := ctx.Authorize([]*pb.Operation{op}); err != nil {
		return false, oops.Wrap(err)
	} else if !auth {
		return false, oops.Wrap(common.UserPermissionDenied)
	}

	return false, nil

}

func (c *GlobalRBAC) CreateCollection(
	ctx common.TxCtxInterface,
	col *authpb.Collection,
) (res *authpb.Collection, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	ctx.GetLogger().Info("EmbedRoleContract.CreateCollection")

	user, err := ctx.GetUserId()
	if err != nil {
		return nil, oops.Wrap(err)
	}

	if state.Exists(ctx, col) {
		return nil, oops.Wrap(common.AlreadyExists)
	}

	// See if the global role collection already exists

	bootstrapedGlobal, err := c.BootstrapGlobal(ctx, col.GetCollectionId())
	if err != nil {
		return nil, oops.Hint("Failed to bootstraped global when needed to").Wrap(err)
	}

	///
	col.AuthType = authpb.AuthType_AUTH_TYPE_GLOBAL_ROLE

	// Make Sure the auth types are in the collection
	authTypes := []string{
		common.CollectionItemType,
		common.RoleItemType,
		common.UserGlobalRoles,
	}
	col.ItemTypes = append(col.GetItemTypes(), authTypes...)
	col.ItemTypes = lo.Uniq(col.GetItemTypes()) // Deduplicate the item types

	// col.Default.DefaultExcludedTypes = append(
	// 	col.GetDefault().GetDefaultExcludedTypes(),
	// 	authTypes...)
	// col.Default.DefaultExcludedTypes = lo.Uniq(col.GetDefault().GetDefaultExcludedTypes())

	// Add the auth types to the collection
	role := &authpb.Role{
		CollectionId: col.GetCollectionId(),
		RoleId:       "manager",
		Polices: &authpb.Polices{
			ItemPolicies: map[string]*authpb.PathPolicy{},
			DefaultPolicy: &authpb.PathPolicy{
				Path:          "",
				FullPath:      "",
				AllowSubPaths: false,
				Actions: []pb.Action{
					pb.Action_ACTION_UTILITY,
					pb.Action_ACTION_VIEW,
					pb.Action_ACTION_CREATE,
					pb.Action_ACTION_UPDATE,
					pb.Action_ACTION_DELETE,
					pb.Action_ACTION_SUGGEST_VIEW,
					pb.Action_ACTION_SUGGEST_CREATE,
					pb.Action_ACTION_SUGGEST_DELETE,
					pb.Action_ACTION_SUGGEST_APPROVE,
					pb.Action_ACTION_VIEW_HISTORY,
					pb.Action_ACTION_VIEW_HIDDEN_TXS,
					pb.Action_ACTION_HIDE_TX,
				},
			},
			// Exclude the auth types from the default policy
		},
		Note:          "Default Admin Role",
		ParentRoleIds: []string{},
	}

	userState := &authpb.UserGlobalRoles{
		CollectionId: "global",
		MspId:        user.GetMspId(),
		UserId:       user.GetUserId(),
	}

	if state.Exists(ctx, userState) {
		if err = state.Get(ctx, userState); err != nil {
			return nil, oops.Hint("User State Exists but failed to retrieve").Wrap(err)
		}
	}

	userState.Roles[col.GetCollectionId()] = &authpb.RoleIDList{RoleId: []string{"manager"}}

	col.ItemTypes = append(col.GetItemTypes(), []string{role.ItemType(), userState.ItemType()}...)
	col.ItemTypes = lo.Uniq(col.GetItemTypes())

	// Create the collection
	if err = (state.Ledger[*authpb.Collection]{}.PrimaryCreate(ctx, col)); err != nil {
		return nil, oops.
			Hint("Creating New Collection failed").
			With("collectionId", col.GetCollectionId(),
				"AuthType", authpb.AuthType_AUTH_TYPE_GLOBAL_ROLE,
			).
			Wrap(err)
	}

	// Create the role object
	if err = (state.Ledger[*authpb.Role]{}.PrimaryCreate(ctx, role)); err != nil {
		return nil, oops.Wrap(err)
	}

	// if we bootstraped we dont need to create the user object as it already exists
	if bootstrapedGlobal {
		return col, nil
	}

	// Create the user object
	if err = (state.Ledger[*authpb.UserGlobalRoles]{}.PrimaryCreate(ctx, userState)); err != nil {
		return nil, oops.Wrap(err)
	}

	ctx.GetLogger().Info("Bootstrapping",
		slog.Any("role", role),
		slog.Any("userRole", userState),
	)

	return col, nil
}
