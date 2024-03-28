package noauth

import (
	"log/slog"

	authpb "github.com/nova38/saacs/libs/saacs-protos-go/auth/v1"
	"github.com/nova38/saacs/pkg/chaincode/policy"
	"github.com/nova38/saacs/pkg/chaincode/state"
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
			ctx.Logger.Debug("Collection found in cache")
			return policy.ValidateOperation(col, op)

		} else {
			ctx.Logger.Debug("Collection not found in cache")

			col := &authpb.Collection{CollectionId: op.GetCollectionId()}

			if err = state.Get(ctx, col); err != nil {
				auth = false
				err = oops.Wrap(err)
				return
			}

			collections[op.GetCollectionId()] = col

			ctx.Logger.Debug("Collection found", slog.Any("col", col))

			return policy.ValidateOperation(col, op)

		}
	}

	return true, nil
}
