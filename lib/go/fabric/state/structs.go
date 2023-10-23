package state

import (
	"fmt"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/reflect/protoregistry"
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

	RegistryTxCtx struct {
		LoggedTxCtx
		Registry *protoregistry.Types
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

func (ctx *RegistryTxCtx) SetRegistry(registry *protoregistry.Types) error {
	if registry == nil {
		return fmt.Errorf("registry is nil")
	}

	ctx.Registry = registry
	return nil
}

func (ctx *RegistryTxCtx) GetRegistry() (*protoregistry.Types, error) {
	if ctx.Registry == nil {
		return nil, fmt.Errorf("registry is nil")
	}

	return ctx.Registry, nil
}
