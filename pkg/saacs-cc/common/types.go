package common

import (
	"log/slog"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	authpb "github.com/nova38/saacs/pkg/saacs-protos/auth/v1"
	v1 "github.com/nova38/saacs/pkg/saacs-protos/auth/v1"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
)

type (
	ServerConfig struct {
		CCID    string
		Address string
	}

	ExtractorFunc func(ctx TxCtxInterface, msg interface{}) (interface{}, error)

	ItemInterface interface {

		// KeyAttr - Returns the key attributes for the item
		KeyAttr() (attr []string)
		SetKeyAttr(attr []string) int

		SetKey(key *authpb.ItemKey)
		ItemKey() *authpb.ItemKey

		ItemType() string

		// ItemKind - Returns the item kind
		//
		// The valid types are:
		//  - Primary
		//  - SubItem (Suggestions and HiddenTxLists)
		ItemKind() authpb.ItemKind

		// KeySchema - Returns the key schema for the item
		KeySchema() *authpb.KeySchema

		// StateKey - Returns the state key for the item
		// This is equivalent to calling MakeStateKey on the ItemKey.
		// It Implements the construction of a composite key just like
		// the CreateCompositeKey function from the shim package
		StateKey() string

		proto.Message
	}

	GenericTxCtxInterface interface {
		TxCtxInterface
	}

	Authorizer interface {

		// Authorize - Checks if the user is authorized to perform the action on
		// the collection
		//
		// # Requirements:
		//  - collection to be set
		//  - action to be set
		//  - domain to be set
		Authorize(op *v1.Operation) (bool, error)
	}

	TxCtxInterface interface {
		contractapi.TransactionContextInterface
		// ════════════════════════════════════════════════════════
		// Generic injectors
		// ════════════════════════════════════════════════════════

		// AddExtractorFunc - Adds an extractor function to the transaction
		// AddExtractorFunc(name string, fn ExtractorFunc)

		// HandelBefore - Handles the before function for the transaction
		HandelBefore() (err error)
		HandleFnError(err *error, r any)
		CloseQueryIterator(resultIterator CommonIteratorInterface)
		// ════════════════════════════════════════════════════════

		GetLogger() *slog.Logger
		LogError(err error)

		GetPageSize() int32
		SetPageSize(pageSize int32)
		// ─────────────────────────────────────────────────────────────────────

		// ════════════════════════════════════════════════════════
		// Validate functions
		// ════════════════════════════════════════════════════════

		// Validate - Validates the message
		Validate(msg proto.Message) error
		ErrorBase() oops.OopsErrorBuilder
		// ─────────────────────────────────────────────────────────────────────

		// ════════════════════════════════════════════════════════
		// Info functions
		// ════════════════════════════════════════════════════════

		// GetFnName - Gets the function name for the transaction
		GetFnName() (name string)

		// ─────────────────────────────────────────────────────────────────────

		// ════════════════════════════════════════════════════════
		// Lifecycle functions
		// ════════════════════════════════════════════════════════

		// MakeLastModified - Makes the last modified activity
		MakeLastModified() (mod *v1.StateActivity, err error)

		// ─────────────────────────────────────────────────────────────────────

		// ════════════════════════════════════════════════════════
		// Auth Functions
		// ════════════════════════════════════════════════════════
		EnabledSuggestions() bool
		EnabledHidden() bool
		// ════════════════════════════════════════════════════════
		//  Operations Functions
		// ════════════════════════════════════════════════════════
		//
		// SetOperation(operation *authpb.Operation)
		// GetOperations() (ops *authpb.Operation, err error)
		// SetOperationsPaths(paths *fieldmaskpb.FieldMask) (err error)

		// ─────────────────────────────────────────────────────────────────────

		// ════════════════════════════════════════════════════════
		//  User Functions
		// ════════════════════════════════════════════════════════

		// GetUserId Uses the ctx stub to get the user id from transaction
		// context
		GetUserId() (user *v1.User, err error)

		// GetUser Retrieves the user from the transaction context or panics
		GetUser() (user *v1.User)

		// GetUser Uses the ctx stub to get the user from the state
		//
		// # Requirements:
		//   - User to be registered
		// GetUser() (user *v1.User, err error)
		// ─────────────────────────────────────────────────────────────────────

		// GetCollection - Gets the collection value from the state
		//
		GetCollection(collectionId string) (col *authpb.Collection, err error)
		//
		//// SetCollection - Sets the collection value in the state
		// SetCollection(collectionId string) (col *authpb.Collection, err error)
		// ─────────────────────────────────────────────────────────────────────

		// ════════════════════════════════════════════════════════
		//  ACL Functions
		// ════════════════════════════════════════════════════════

		// GetViewMask() (mask *fieldmaskpb.FieldMask)

		// Authorize - Checks if the user is authorized to perform the action on
		// the collection
		//
		// # Requirements:
		//  - collection to be set
		//  - action to be set
		//  - domain to be set
		Authorize(ops []*v1.Operation) (bool, error)

		// GetMask - Request the mask for the operation from the auth service

		// PostActionProcessing - Used to modify the item after the action has been performed
		// Before the item is saved to the state, this is used to add the
		// last modified activity to the item
		PostActionProcessing(item ItemInterface, ops []*v1.Operation) (err error)
		CheckBootstrap() (bootstraped bool, err error)
	}

	// ─────────────────────────────────────────────────────────────────────

	// CommonIteratorInterface allows a saacs-cc to check whether any more result
	// to be fetched from an iterator and close it when done.
	//
	// From on the "github.com/hyperledger/fabric-saacs-cc-go/shim" package's CommonIteratorInterface
	CommonIteratorInterface interface {
		// HasNext returns true if the range query iterator contains additional keys
		// and values.
		HasNext() bool

		// Close closes the iterator. This should be called when done
		// reading from the iterator to free up resources.
		Close() error
	}
)
