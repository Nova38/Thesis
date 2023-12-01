package identity

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
)

func Authorize(ctx common.TxCtxInterface, ops []*authpb.Operation) (bool, error) {
	ctx.GetLogger().Info("NoAuthContract.Authenticate")

	for _, op := range ops {
		ctx.GetLogger().Info(op.String())

		// policy.ValidateOperation()
	}

	return true, nil
}
