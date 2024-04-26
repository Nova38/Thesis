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

type Identity struct{}

func (c *Identity) CreateCollection(
	ctx common.TxCtxInterface,
	col *authpb.Collection,
) (res *authpb.Collection, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	ctx.GetLogger().Info("RoleContract.CreateCollection")
	user, err := ctx.GetUserId()
	if err != nil {
		return nil, oops.Wrap(err)
	}

	if state.Exists(ctx, col) {
		return nil, oops.Errorf("Collection already exists")
	}

	col.AuthType = authpb.AuthType_AUTH_TYPE_IDENTITY

	// Make Sure the auth types are in the collection
	authTypes := []string{
		common.CollectionItemType,
		common.UserDirectMembershipItemType,
	}
	col.ItemTypes = append(col.GetItemTypes(), authTypes...)
	col.ItemTypes = lo.Uniq(col.GetItemTypes()) // Deduplicate the item types

	// Exclude the auth types from the default policy
	// col.Default.DefaultExcludedTypes = append(
	// 	col.GetDefault().GetDefaultExcludedTypes(),
	// 	authTypes...)
	col.Default.DefaultExcludedTypes = lo.Uniq(col.GetDefault().GetDefaultExcludedTypes())

	// Add the auth types to the collection
	membership := &authpb.UserDirectMembership{
		CollectionId: col.GetCollectionId(),
		MspId:        user.GetMspId(),
		UserId:       user.GetUserId(),
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
		},
	}

	col.ItemTypes = append(col.GetItemTypes(), []string{membership.ItemType()}...)
	col.ItemTypes = lo.Uniq(col.GetItemTypes())

	// state.Put(ctx, col)
	if err = (state.Ledger[*authpb.Collection]{}.PrimaryCreate(ctx, col)); err != nil {
		return nil, oops.Wrap(err)
	}
	if err = (state.Ledger[*authpb.UserDirectMembership]{}.PrimaryCreate(ctx, membership)); err != nil {
		return nil, oops.Wrap(err)
	}

	ctx.GetLogger().Info("Bootstrapping", slog.Any("membership", membership))

	return col, nil

}
