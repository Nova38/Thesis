package models

import (
	"log/slog"

	"github.com/nova38/saacs/apps/saacs-cc/common"
	authpb "github.com/nova38/saacs/pkg/saacs-protos/auth/v1"
)

type NoAuth struct {
	Collection *authpb.Collection

	TxCtx  common.TxCtxInterface
	Logger *slog.Logger
}

func (ac *NoAuth) Authorize(op *authpb.Operation) (bool, error) {
	return true, nil
}
