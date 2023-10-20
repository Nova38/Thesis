package state

import (
	"fmt"
	"github.com/bufbuild/protovalidate-go"
	"log/slog"
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
)

func (ctx *LoggedTxCtx) GetLogger() *slog.Logger {
	return ctx.Logger
}

func (ctx *LoggedTxCtx) SetLogger(logger *slog.Logger) error {
	if logger == nil {
		return fmt.Errorf("Logger is nil")
	}

	ctx.Logger = logger
	return nil
}

func (ctx *ValidateAbleTxCtx) GetValidator() (*protovalidate.Validator, error) {

	if ctx.Validator == nil {
		v, err := protovalidate.New()
		if err != nil {
			return nil, fmt.Errorf("failed to create validator: %w", err)
		}
		ctx.Validator = v
	}

	return ctx.Validator, nil
}
