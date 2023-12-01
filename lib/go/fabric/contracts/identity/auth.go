package identity

import (
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
)

func (ctx *IdentityTxCtx) Authorize(ops []*authpb.Operation) (bool, error) {

	ctx.GetLogger().Info("NoAuthContract.Authenticate")

	for _, op := range ops {
		ctx.GetLogger().Info(op.String())

		// policy.ValidateOperation()
	}

	return true, nil
}
