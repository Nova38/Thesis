package rbac

import (
	"github.com/nova38/thesis/lib/go/fabric/state"
	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

/*
   Transaction Context interfaces
*/

// An interface for a transaction context that has a user and collection

type GenericAuthTxCtxInterface interface {
	AuthTxCtxInterface
}

type AuthTxCtxInterface interface {
	state.LoggedTxCtxInterface
	state.ValidateAbleTxCtxInterface
	state.PagedTxCtxInterface

	// HandleBefore: Handles the before function for the transaction
	HandelBefore() (err error)

	// =============================================
	// Util functions
	// =============================================

	GetFnName() (name string)
	MakeLastModified() (mod *rbac_pb.StateActivity, err error)
	// ----------------------------------------------

	// =============================================
	//  User Functions
	// =============================================

	// GetUserId Uses the ctx stub to get the user id from transaction context
	//
	GetUserId() (*rbac_pb.User_Id, error)

	// GetUser Uses the ctx stub to get the user from the state
	// # Requirements:
	//   - User to be registered
	GetUser() (user *rbac_pb.User, err error)
	// ----------------------------------------------

	// =============================================
	//  Collection Functions
	// =============================================

	// GetCollection Gets the collection value from the state
	// requires:
	//  - collection to be set
	GetCollection() (col *rbac_pb.Collection, err error)

	// SetCollection Sets the collection value in the state
	SetCollection(id *rbac_pb.Collection_Id) (col *rbac_pb.Collection, err error)
	// ----------------------------------------------

	// =============================================
	//  Role Functions
	// =============================================

	// GetRole: Gets the role of the user in the collection.
	//
	//
	// # Requirements:
	//   - collection to be set
	//   - user to be set
	// if user is not a member of the collection it returns 0 for public user
	GetRole() (role int, err error)

	// ----------------------------------------------

	// =============================================
	//  Operations Functions
	// =============================================

	SetOperation(operation *rbac_pb.ACL_Operation) (err error)
	GetOperations() (ops *rbac_pb.ACL_Operation, err error)
	SetOperationsPaths(paths *fieldmaskpb.FieldMask) (err error)

	// =============================================
	//  ACL Functions
	// =============================================

	// Authorize: Checks if the user is authorized to perform the action on the collection
	//
	// # Requirements:
	//  - collection to be set
	//  - action to be set
	//  - domain to be set
	Authorize() (bool, error)
}

// AuthTransactionObjectInterfaces

// type RoleWrapperInterface interface {
//     GetRole() (int, error)
// }
// type RolePermissionWrapperInterface interface {
//     GetRolePermission(role int) (*rbac_pb.ACL, error)
// }
