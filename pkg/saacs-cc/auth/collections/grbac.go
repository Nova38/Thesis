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
