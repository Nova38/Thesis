package models

import (
	"log/slog"

	"github.com/nova38/saacs/apps/saacs-cc/auth/policy"
	"github.com/nova38/saacs/apps/saacs-cc/common"
	authpb "github.com/nova38/saacs/pkg/saacs-protos/auth/v1"
	"github.com/samber/oops"
)

type ABAC struct {
	Collection *authpb.Collection

	UserAttributes []*authpb.Attribute

	TxCtx  common.TxCtxInterface
	Logger *slog.Logger
}

func (ac *ABAC) Authorize(op *authpb.Operation) (bool, error) {
	// ═════════════════════════════════════════════
	// Default Policy
	// ═════════════════════════════════════════════
	if auth, err := policy.AuthorizedPolicy(ac.Collection.GetDefault(), op); err != nil {
		return false, oops.Wrap(err)
	} else if auth {
		ac.Logger.Info("User is authorized by default")
		return true, nil
	}

	return true, nil
}
