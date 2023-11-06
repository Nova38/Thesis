package state

import (
	"log/slog"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"google.golang.org/protobuf/proto"
)

type LoggedTxCtxInterface interface {
	contractapi.TransactionContextInterface
	GetLogger() *slog.Logger
	// SetLogger(logger *slog.Logger) error

	// HandleFnError Handles the error from a function
	// also handles panics
	HandleFnError(err *error, r any)
}

// PagedTxCtxInterface Paged Transaction Context Interface
type PagedTxCtxInterface interface {
	LoggedTxCtxInterface
	GetPageSize() int32
	SetPageSize(pageSize int32)
}

type ValidateAbleTxCtxInterface interface {
	Validate(msg proto.Message) error
}

type InfoAbleTxCtxInterface interface {
	GetFnName() (name string)
}
