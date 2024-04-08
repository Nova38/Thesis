package embedded

import (
	"log/slog"

	"github.com/nova38/saacs/apps/saacs-cc/common"
	authpb "github.com/nova38/saacs/libs/saacs-protos-go/auth/v1"
)

type ERBAC struct {
	Collections map[string]*authpb.Collection

	CollectionMemberships map[string]*authpb.UserDirectMembership

	TxCtx  common.TxCtxInterface
	Logger *slog.Logger
}

func (ac *ERBAC) Authorize(ops []*authpb.Operation) (bool, error) {

	return true, nil
}
