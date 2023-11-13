package state

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"log/slog"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"google.golang.org/protobuf/proto"
)

type (
	ExtractorFunc func(ctx TxCtxInterface, msg interface{}) (interface{}, error)

	//AuthFunc func(op *authpb.Operation) (bool, error)

	Object interface {
		Key() (attr []string, err error)
		FlatKey() string
		Namespace() string
		GetCollectionId() string

		proto.Message
	}

	StubInterface[T Object] interface {
		Exists(ctx TxCtxInterface, key string) bool
		Create(ctx TxCtxInterface, obj T) (err error)
		Update(ctx TxCtxInterface, obj T) (err error)
		Delete(ctx TxCtxInterface, obj T) (err error)

		Get(ctx TxCtxInterface, obj T) (err error)
		List(ctx TxCtxInterface, obj T, bookmark string) (list []T, err error)
		ByCollection(ctx TxCtxInterface, obj T, bookmark string) (list []T, err error)
		ByPartialKey(ctx TxCtxInterface, obj T, numAttr int, bookmark string) (list []T, err error)
	}

	HistoryStubInterface[T Object] interface {
		History(ctx TxCtxInterface, obj T) (err error)
		FullHistory(ctx TxCtxInterface, obj T) (err error)
		Transaction(ctx TxCtxInterface, obj T) (err error)
		HideTransaction(ctx TxCtxInterface, obj T) (err error)
	}

	SuggestStubInterface[T Object] interface {
		Suggestion(
			ctx TxCtxInterface,
			suggestionId string,
		) (suggestion *authpb.Suggestion, err error)
		CreateSuggestion(ctx TxCtxInterface, suggestion *authpb.Suggestion) (err error)
		DeleteSuggestion(ctx TxCtxInterface, suggestion *authpb.Suggestion) (err error)
		ApproveSuggestion(ctx TxCtxInterface, suggestion *authpb.Suggestion) (err error)

		SuggestionList(
			ctx TxCtxInterface,
			obj T,
			bookmark string,
		) (list []authpb.Suggestion, err error)
		SuggestionByCollection(
			ctx TxCtxInterface,
			obj T,
			bookmark string,
		) (list []authpb.Suggestion, err error)
		SuggestionByPartialKey(
			ctx TxCtxInterface,
			obj T,
			numAttr int,
			bookmark string,
		) (list []authpb.Suggestion, err error)
	}

	GenericTxCtxInterface interface {
		TxCtxInterface
	}

	TxCtxInterface interface {
		contractapi.TransactionContextInterface
		// ════════════════════════════════════════════════════════
		// Generic injectors
		// ════════════════════════════════════════════════════════

		// AddExtractorFunc - Adds an extractor function to the transaction
		//AddExtractorFunc(name string, fn ExtractorFunc)

		// HandelBefore - Handles the before function for the transaction
		HandelBefore() (err error)
		HandleFnError(err *error, r any)
		CloseQueryIterator(resultIterator shim.CommonIteratorInterface)
		// ════════════════════════════════════════════════════════

		GetLogger() *slog.Logger

		GetPageSize() int32
		SetPageSize(pageSize int32)
		// ---------------------------------------------------------------------

		// ════════════════════════════════════════════════════════
		// Validate functions
		// ════════════════════════════════════════════════════════

		// Validate - Validates the message
		Validate(msg proto.Message) error
		// ---------------------------------------------------------------------

		// ════════════════════════════════════════════════════════
		// Info functions
		// ════════════════════════════════════════════════════════

		// GetFnName - Gets the function name for the transaction
		GetFnName() (name string)

		// ---------------------------------------------------------------------

		// ════════════════════════════════════════════════════════
		// Lifecycle functions
		// ════════════════════════════════════════════════════════

		// MakeLastModified - Makes the last modified activity
		MakeLastModified() (mod *authpb.StateActivity, err error)

		// ---------------------------------------------------------------------

		// ════════════════════════════════════════════════════════
		// Auth Functions
		// ════════════════════════════════════════════════════════

		// SetAuthenticator - Sets the authenticator for the transaction
		SetAuthenticator(fn string)
		GetAuthenticator() (fn string)

		// ════════════════════════════════════════════════════════
		//  Operations Functions
		// ════════════════════════════════════════════════════════
		//
		//SetOperation(operation *auth_pb.Operation)
		//GetOperations() (ops *auth_pb.Operation, err error)
		//SetOperationsPaths(paths *fieldmaskpb.FieldMask) (err error)

		// ---------------------------------------------------------------------

		// ════════════════════════════════════════════════════════
		//  User Functions
		// ════════════════════════════════════════════════════════

		// GetUserId Uses the ctx stub to get the user id from transaction
		// context
		GetUserId() (mspId string, userId string, err error)

		// GetUser Uses the ctx stub to get the user from the state
		//
		// # Requirements:
		//   - User to be registered
		GetUser() (user *authpb.User, err error)

		// ---------------------------------------------------------------------

		// GetCollection - Gets the collection value from the state
		//
		// # Requirements:
		//  - collection to be set
		GetCollection() (col *authpb.Collection, err error)

		// SetCollection - Sets the collection value in the state
		SetCollection(collectionId string) (col *authpb.Collection, err error)
		// ---------------------------------------------------------------------

		// ════════════════════════════════════════════════════════
		//  ACL Functions
		// ════════════════════════════════════════════════════════

		//// GetACLEntryKey - Gets the Key for the entry in the ACL
		//GetACLEntryKey() (key string, err error)

		//// GetViewMask - Gets the view mask for the collection and the role
		//GetViewMask() (mask fieldmaskpb.FieldMask, err error)

		// Authorize - Checks if the user is authorized to perform the action on
		// the collection
		//
		// # Requirements:
		//  - collection to be set
		//  - action to be set
		//  - domain to be set
		Authorize(ops *authpb.Operation) (bool, error)

		// GetMask - Request the mask for the operation from the auth service

	}
)
