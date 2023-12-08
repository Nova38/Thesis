package noauth

import (
	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	"github.com/nova38/thesis/packages/saacs/policy"
	"github.com/nova38/thesis/packages/saacs/state"
	"github.com/samber/oops"
)

func (ctx *Ctx) Authorize(ops []*authpb.Operation) (auth bool, err error) {

	// ────────────────────────────────── profile ──────────────────────────────────────

	ctx.GetLogger().Info("NoAuthContract.Authenticate")

	collections := map[string]*authpb.Collection{}

	for _, op := range ops {
		ctx.GetLogger().Info(op.String())

		if op.GetItemType() == "auth.Collection" {
			if op.GetAction() == authpb.Action_ACTION_CREATE {
				auth = true
				err = nil
				return
			}
		}

		if col, ok := collections[op.GetCollectionId()]; ok {
			return policy.ValidateOperation(col, op)

		} else {
			col := &authpb.Collection{CollectionId: op.GetCollectionId()}

			if err = state.Get(ctx, col); err != nil {
				auth = false
				err = oops.Wrap(err)
				return
			}

			collections[op.GetCollectionId()] = col

			return policy.ValidateOperation(col, op)

		}
	}

	return true, nil
}
