package auth

import (
	"log/slog"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type (
	Authenticater interface {
		any
		// SetCollection(*auth_pb.Collection)
		// SetIdentifier(*auth_pb.Identifier)
		// SetOperation(*auth_pb.Operation)

		// GetCollection() *auth_pb.Collection
		// GetIdentifier() *auth_pb.Identifier
		// GetOperation() *auth_pb.Operation
	}
	GBaseTxCtx[T Authenticater] struct {
		contractapi.TransactionContext

		AuthTransactionObjects

		Logger   *slog.Logger
		PageSize int32

		// auth values
		authorized  bool
		authChecked bool
	}
)

// AuthorizeOperation
