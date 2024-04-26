package models

import (
	"log/slog"

	"github.com/nova38/saacs/pkg/saacs-cc/common"
	authpb "github.com/nova38/saacs/pkg/saacs-protos/saacs/auth/v0"
	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"
)

type NoAuth struct {
	Collection *authpb.Collection

	TxCtx  common.TxCtxInterface
	Logger *slog.Logger
}

func (ac *NoAuth) Authorize(op *pb.Operation) (bool, error) {
	return true, nil
}
