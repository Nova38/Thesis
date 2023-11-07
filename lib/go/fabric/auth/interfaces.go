package auth

import (
	"log/slog"

	auth_pb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

type (
	GenericAuthTxCtxInterface interface {
		AuthCtxInterface
	}
	AuthCtxInterface interface {
		// =============================================
		// Generic injectors
		// =============================================

		// SetAuthenticator - Sets the authenticator for the transaction
		// SetAuthenticator(auth Authenticater) (err error)
		// GetAuthenticator() (auth Authenticater)

		// =============================================
		// Logging functions
		// =============================================

		// GetLogger - Gets the logger for the transaction
		GetLogger() *slog.Logger

		// -------------------------------------------------------------------------

		// ==================================================
		// Paging functions
		// ==================================================

		// GetPageSize - Gets the page size for the transaction
		GetPageSize() int32
		SetPageSize(pageSize int32)
		// -------------------------------------------------------------------------

		// ==================================================
		// Validate functions
		// ==================================================

		// Validate - Validates the message
		Validate(msg proto.Message) error
		// -------------------------------------------------------------------------

		// ==================================================
		// Info functions
		// ==================================================

		// GetFnName - Gets the function name for the transaction
		GetFnName() (name string)

		// -------------------------------------------------------------------------

		// ==================================================
		// Lifecycle functions
		// ==================================================

		// HandelBefore - Handles the before function for the transaction
		HandelBefore() (err error)

		// MakeLastModified - Makes the last modified activity
		MakeLastModified() (mod *auth_pb.StateActivity, err error)

		// -------------------------------------------------------------------------

		// =============================================
		//  Operations Functions
		// =============================================

		SetOperation(operation *auth_pb.Operation) (err error)
		GetOperations() (ops *auth_pb.Operation, err error)
		SetOperationsPaths(paths *fieldmaskpb.FieldMask) (err error)

		// -------------------------------------------------------------------------

		// ==================================================
		//  User Functions
		// ==================================================

		// GetUserId Uses the ctx stub to get the user id from transaction context
		GetUserId() (*auth_pb.User_Id, error)

		// GetUser Uses the ctx stub to get the user from the state
		//
		// # Requirements:
		//   - User to be registered
		GetUser() (user *auth_pb.User, err error)

		// -------------------------------------------------------------------------
		// =============================================
		//  Collection Functions
		// =============================================

		// GetCollection - Gets the collection value from the state
		//
		// # Requirements:
		//  - collection to be set
		GetCollection() (col *auth_pb.Collection, err error)

		// SetCollection - Sets the collection value in the state
		SetCollection(id *auth_pb.Collection_Id) (col *auth_pb.Collection, err error)
		// -------------------------------------------------------------------------

		// ==================================================
		//  ACL Functions
		// ==================================================

		// GetACLKey - Gets the Key for the entry in the ACL
		GetACLEntryKey() (key string, err error)

		// GetViewMask - Gets the view mask for the collection and the role
		GetViewMask() (mask fieldmaskpb.FieldMask, err error)

		// Authorize - Checks if the user is authorized to perform the action on the collection
		//
		// # Requirements:
		//  - collection to be set
		//  - action to be set
		//  - domain to be set
		Authorize() (bool, error)
	}
)
