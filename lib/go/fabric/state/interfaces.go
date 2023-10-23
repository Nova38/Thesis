package state

import (
	"log/slog"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/reflect/protoregistry"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type LoggedTxCtxInterface interface {
	contractapi.TransactionContextInterface
	GetLogger() *slog.Logger
	SetLogger(logger *slog.Logger) error
}

// PagedTxCtxInterface Paged Transaction Context Interface
type PagedTxCtxInterface interface {
	LoggedTxCtxInterface
	GetPageSize() int32
	SetPageSize(pageSize int32)
}

type ValidateAbleTxCtxInterface interface {
	GetValidator() (*protovalidate.Validator, error)
}

type RegistryTxCtxInterface interface {
	LoggedTxCtxInterface
	GetRegistry() (*protoregistry.Types, error)
	SetRegistry(registry *protoregistry.Types) error
}
