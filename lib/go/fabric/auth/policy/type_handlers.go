package policy

import authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"

func handleCollection(ctx TxCtxInterface, op *authpb.Operation) (err error) {
	ctx.GetLogger().Info("handleCollection")
	return nil
}
