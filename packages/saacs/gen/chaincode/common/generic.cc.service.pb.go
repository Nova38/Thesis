// Code generated by proto-gen-go-auth_pb. DO NOT EDIT.
// versions:
// - protoc-gen-cckey v0.0.1
// source: chaincode/common/generic.proto

package common

import (
	fmt "fmt"
	common "github.com/nova38/thesis/packages/saacs/common"
	v1 "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
)

// Service GenericService
type GenericServiceInterface[T common.GenericTxCtxInterface] interface {
	// ══════════════════════════════════ Helper ═════════════════════════════════════
	// ────────────────────────────────── Query ──────────────────────────────────────
	//
	//	rpc GetAllTypes(google.protobuf.Empty) returns (GetAllTypesResponse) {
	//	  option (auth.transaction_type) = TRANSACTION_TYPE_QUERY;
	//	  option (auth.operation) = {action: ACTION_UTILITY};
	//	}
	//
	// # Operation:
	//   - Domain: ACTION_VIEW
	//
	// req is empty
	GetCurrentUser(ctx T) (res *GetCurrentUserResponse, err error)

	// ──────────────────────────────── Invoke ───────────────────────────────────────
	// # Operation:
	//   - Domain: ACTION_UTILITY
	Bootstrap(ctx T, req *BootstrapRequest) (res *BootstrapResponse, err error)

	// AuthorizeOperation
	//
	// # Operation:
	//   - Domain: ACTION_UTILITY
	AuthorizeOperation(ctx T, req *AuthorizeOperationRequest) (res *AuthorizeOperationResponse, err error)

	// Get
	//
	// # Operation:
	//   - Domain: ACTION_VIEW
	Get(ctx T, req *GetRequest) (res *GetResponse, err error)

	// List
	//
	// # Operation:
	//   - Domain: ACTION_VIEW
	List(ctx T, req *ListRequest) (res *ListResponse, err error)

	// ListByCollection
	//
	// # Operation:
	//   - Domain: ACTION_VIEW
	ListByCollection(ctx T, req *ListByCollectionRequest) (res *ListByCollectionResponse, err error)

	// ListByAttrs
	//
	// # Operation:
	//   - Domain: ACTION_VIEW
	ListByAttrs(ctx T, req *ListByAttrsRequest) (res *ListByAttrsResponse, err error)

	// Create
	//
	// # Operation:
	//   - Domain: ACTION_CREATE
	Create(ctx T, req *CreateRequest) (res *CreateResponse, err error)

	// Update
	//
	// # Operation:
	//   - Domain: ACTION_UPDATE
	Update(ctx T, req *UpdateRequest) (res *UpdateResponse, err error)

	// Delete
	//
	// # Operation:
	//   - Domain: ACTION_DELETE
	Delete(ctx T, req *DeleteRequest) (res *DeleteResponse, err error)

	// GetHistory
	//
	// # Operation:
	//   - Domain: ACTION_VIEW_HISTORY
	GetHistory(ctx T, req *GetHistoryRequest) (res *GetHistoryResponse, err error)

	// GetHiddenTx
	//
	// # Operation:
	//   - Domain: ACTION_VIEW_HIDDEN_TXS
	GetHiddenTx(ctx T, req *GetHiddenTxRequest) (res *GetHiddenTxResponse, err error)

	// HideTx
	//
	// # Operation:
	//   - Domain: ACTION_HIDE_TX
	HideTx(ctx T, req *HideTxRequest) (res *HideTxResponse, err error)

	// UnHideTx
	//
	// # Operation:
	//   - Domain: ACTION_UNHIDE_TX
	UnHideTx(ctx T, req *UnHideTxRequest) (res *UnHideTxResponse, err error)

	// GetSuggestion
	//
	// # Operation:
	//   - Domain: ACTION_SUGGEST_VIEW
	GetSuggestion(ctx T, req *GetSuggestionRequest) (res *GetSuggestionResponse, err error)

	// SuggestionListByCollection
	//
	// # Operation:
	//   - Domain: ACTION_SUGGEST_VIEW
	SuggestionListByCollection(ctx T, req *SuggestionListByCollectionRequest) (res *SuggestionListByCollectionResponse, err error)

	// SuggestionByPartialKey
	//
	// # Operation:
	//   - Domain: ACTION_SUGGEST_VIEW
	SuggestionByPartialKey(ctx T, req *SuggestionByPartialKeyRequest) (res *SuggestionByPartialKeyResponse, err error)

	// ──────────────────────────────── Invoke ───────────────────────────────────────
	// # Operation:
	//   - Domain: ACTION_SUGGEST_CREATE
	SuggestionCreate(ctx T, req *SuggestionCreateRequest) (res *SuggestionCreateResponse, err error)

	// SuggestionDelete
	//
	// # Operation:
	//   - Domain: ACTION_SUGGEST_DELETE
	SuggestionDelete(ctx T, req *SuggestionDeleteRequest) (res *SuggestionDeleteResponse, err error)

	// SuggestionApprove
	//
	// # Operation:
	//   - Domain: ACTION_SUGGEST_CREATE
	SuggestionApprove(ctx T, req *SuggestionApproveRequest) (res *SuggestionApproveResponse, err error)
}

type GenericServiceBase struct {
}

func (s *GenericServiceBase) GetEvaluateTransactions() []string {
	return []string{
		"GetCurrentUser",
		"Get",
		"List",
		"ListByCollection",
		"ListByAttrs",
		"GetHistory",
		"GetHiddenTx",
		"GetSuggestion",
		"SuggestionListByCollection",
		"SuggestionByPartialKey",
	}
}

func GenericServiceGetTxOperation(txName string) (op *v1.Operation, err error) {
	switch txName {
	case "GetCurrentUser":
		// action:ACTION_VIEW
		return &v1.Operation{
			Action: 10,
		}, nil
	case "Bootstrap":
		// action:ACTION_UTILITY
		return &v1.Operation{
			Action: 1,
		}, nil
	case "AuthorizeOperation":
		// action:ACTION_UTILITY
		return &v1.Operation{
			Action: 1,
		}, nil
	case "Get":
		// action:ACTION_VIEW
		return &v1.Operation{
			Action: 10,
		}, nil
	case "List":
		// action:ACTION_VIEW
		return &v1.Operation{
			Action: 10,
		}, nil
	case "ListByCollection":
		// action:ACTION_VIEW
		return &v1.Operation{
			Action: 10,
		}, nil
	case "ListByAttrs":
		// action:ACTION_VIEW
		return &v1.Operation{
			Action: 10,
		}, nil
	case "Create":
		// action:ACTION_CREATE
		return &v1.Operation{
			Action: 11,
		}, nil
	case "Update":
		// action:ACTION_UPDATE
		return &v1.Operation{
			Action: 12,
		}, nil
	case "Delete":
		// action:ACTION_DELETE
		return &v1.Operation{
			Action: 13,
		}, nil
	case "GetHistory":
		// action:ACTION_VIEW_HISTORY
		return &v1.Operation{
			Action: 18,
		}, nil
	case "GetHiddenTx":
		// action:ACTION_VIEW_HIDDEN_TXS
		return &v1.Operation{
			Action: 19,
		}, nil
	case "HideTx":
		// action:ACTION_HIDE_TX
		return &v1.Operation{
			Action: 20,
		}, nil
	case "UnHideTx":
		// action:ACTION_UNHIDE_TX
		return &v1.Operation{
			Action: 21,
		}, nil
	case "GetSuggestion":
		// action:ACTION_SUGGEST_VIEW
		return &v1.Operation{
			Action: 14,
		}, nil
	case "SuggestionListByCollection":
		// action:ACTION_SUGGEST_VIEW
		return &v1.Operation{
			Action: 14,
		}, nil
	case "SuggestionByPartialKey":
		// action:ACTION_SUGGEST_VIEW
		return &v1.Operation{
			Action: 14,
		}, nil
	case "SuggestionCreate":
		// action:ACTION_SUGGEST_CREATE
		return &v1.Operation{
			Action: 15,
		}, nil
	case "SuggestionDelete":
		// action:ACTION_SUGGEST_DELETE
		return &v1.Operation{
			Action: 16,
		}, nil
	case "SuggestionApprove":
		// action:ACTION_SUGGEST_CREATE
		return &v1.Operation{
			Action: 15,
		}, nil
	default:
		return nil, fmt.Errorf("No operation defined for " + txName)
	}
	return nil, nil
}

func (s *GenericServiceBase) GetIgnoredFunctions() []string {
	return []string{"GetTxOperation"}
}
