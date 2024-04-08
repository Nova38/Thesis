package rabc

import (
	"log/slog"

	"github.com/nova38/saacs/apps/saacs-cc/common"
	authpb "github.com/nova38/saacs/libs/saacs-protos-go/auth/v1"
)

type RBAC struct {
	Collections map[string]*authpb.Collection

	CollectionMemberships map[string]*authpb.UserDirectMembership

	TxCtx  common.TxCtxInterface
	Logger *slog.Logger
}

func (ac *RBAC) Authorize(ops []*authpb.Operation) (bool, error) {

	return true, nil
}
