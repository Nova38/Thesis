package collections

import (
	authpb "github.com/nova38/saacs/pkg/saacs-protos/saacs/auth/v0"

	"github.com/nova38/saacs/pkg/saacs-cc/common"
	"github.com/nova38/saacs/pkg/saacs-cc/state"
	"github.com/samber/lo"
)

type NoAuth struct{}

func (c *NoAuth) CreateCollection(
	ctx common.TxCtxInterface,
	col *authpb.Collection,
) (res *authpb.Collection, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	ctx.GetLogger().Info("NoAuthContract.CreateCollection")

	if ctx.Validate(col) != nil {
		return nil, ctx.ErrorBase().Wrap(err)
	}

	// Check if the collection already exists
	if exist := state.Exists(ctx, col); exist {
		return nil, ctx.ErrorBase().
			With("CollectionId", col.CollectionId).
			Wrap(common.AlreadyExists)
	}

	// Make Sure the auth types are in the collection
	authTypes := []string{
		common.CollectionItemType,
		common.RoleItemType,
		common.UserCollectionRolesItemType,
	}
	col.ItemTypes = append(col.GetItemTypes(), authTypes...)
	col.ItemTypes = lo.Uniq(col.GetItemTypes()) // Deduplicate the item types

	err = state.Ledger[*authpb.Collection]{}.PrimaryCreate(ctx, col)

	if err != nil {
		return nil, ctx.ErrorBase().Wrap(err)
	}

	return col, nil
}
