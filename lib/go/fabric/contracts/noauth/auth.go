package noauth

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/policy"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/oops"
)

func (ctx *NoAuthCtx) Authorize(ops []*authpb.Operation) (bool, error) {
	ctx.GetLogger().Info("NoAuthContract.Authenticate")

	collections := map[string]*authpb.Collection{}

	for _, op := range ops {
		ctx.GetLogger().Info(op.String())

		if op.GetCollectionId() == "global" {
			ctx.Logger.Info("global collection")
			continue
		}
		if op.GetItemType() == "auth.Collection" {
			if op.GetAction() == authpb.Action_ACTION_CREATE {
				return true, nil
			}
		}

		if col, ok := collections[op.GetCollectionId()]; ok {
			valid, err := policy.ValidateOperation(col, op)
			if err != nil {
				return false, oops.Wrap(err)
			}
			if !valid {
				return false, nil
			}
		} else {
			col := &authpb.Collection{CollectionId: op.GetCollectionId()}

			key, err := common.MakePrimaryKey(col)
			if err != nil {
				return false, oops.Wrap(err)
			}
			if err := state.Get(ctx, key, col); err != nil {
				return false, oops.Wrap(err)
			}
			collections[op.GetCollectionId()] = col
			policy.ValidateOperation(col, op)
		}
	}

	return true, nil
}
