package state

import (
	"log/slog"

	"github.com/bufbuild/protovalidate-go"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
	// "google.golang.org/protobuf/reflect/protoregistry"
)

type (
	LoggedTxCtx struct {
		LoggedTxCtxInterface
		Logger *slog.Logger
	}
	PagedTxCtx struct {
		LoggedTxCtxInterface
		PageSize int32
	}
	ValidateAbleTxCtx struct {
		ValidateAbleTxCtxInterface
		Validator *protovalidate.Validator
	}

	// RegistryTxCtx struct {
	// 	LoggedTxCtx
	// 	Registry *protoregistry.Types
	// }
)

func (ctx *LoggedTxCtx) GetLogger() *slog.Logger {
	return ctx.Logger
}

func (ctx *LoggedTxCtx) NewLogger() (*slog.Logger, error) {
	ctx.Logger = &slog.Logger{}
	return ctx.Logger, nil
}

func (ctx *ValidateAbleTxCtx) GetValidator() (*protovalidate.Validator, error) {
	if ctx.Validator == nil {
		v, err := protovalidate.New()
		if err != nil {
			return nil, oops.Errorf("failed to create validator: %w", err)
		}
		ctx.Validator = v
	}

	return ctx.Validator, nil
}

func (ctx *ValidateAbleTxCtx) Validate(msg proto.Message) error {
	v, err := ctx.GetValidator()
	if err != nil {
		return oops.Errorf("failed to get validator: %w", err)
	}
	if v == nil {
		return oops.Errorf("validator is nil")
	}
	return v.Validate(msg)
}
