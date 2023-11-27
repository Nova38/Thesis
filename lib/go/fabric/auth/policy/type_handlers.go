package policy

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
)

func handleCollection(ctx common.TxCtxInterface, op *authpb.Operation) (err error) {
	ctx.GetLogger().Info("handleCollection")
	return nil
}
