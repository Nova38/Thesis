package state

import (
	"log/slog"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type LogedTxCtxInterface interface {
	contractapi.TransactionContextInterface
	GetLogger() *slog.Logger
	SetLogger(logger *slog.Logger)
}

// Paged Transaction Context Interface
type PagedTxCtxInterface interface {
	LogedTxCtxInterface
	GetPageSize() int32
	SetPageSize(pageSize int32)
}
