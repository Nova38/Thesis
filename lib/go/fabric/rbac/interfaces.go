package rbac

import (
	"github.com/nova38/thesis/lib/go/fabric/state"
	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
)

/*
   Transaction Context interfaces
*/

// An interface for a transaction context that has a user and collection

type AuthTxCtxInterface interface {
	state.LoggedTxCtxInterface
	state.ValidateAbleTxCtxInterface

	// GetUserId Uses the ctx stub to get the user id from transaction context
	GetUserId() (*rbac_pb.User_Id, error)

	// GetUser Uses the ctx stub to get the user from the state
	GetUser() (*rbac_pb.User, error)

	// GetCollection Gets the collection value from the state
	// requires:
	//  - collection to be set
	GetCollection() (*rbac_pb.Collection, error)

	// SetCollection Sets the collection value in the state
	SetCollection(collection *rbac_pb.Collection) error

	// GetRole: Gets the role of the user in the collection.
	//
	//
	// Requirements:
	//  - collection to be set
	//  - user to be set
	// if user is not a member of the collection it returns 0 for public user
	GetRole() (int, error)

	// GetRolePermission: Gets the permission of the user in the collection.
	//
	// # Requirements:
	//  - collection to be set
	//  - user to be set
	// if user is not a member of the collection it returns that public user permissions
	GetRolePermission(role int) (*rbac_pb.ACL, error)

	// GetDomain() (*rbac_pb.Operations_Domain, error)
	// SetDomain(domain *rbac_pb.Operations_Domain) error

	// GetAction() (*rbac_pb.Operations_Action, error)
	// SetAction(action *rbac_pb.Operations_Action) error

	SetOperation(operation *rbac_pb.ACL_Operation) error
	GetOperation() (*rbac_pb.ACL_Operation, error)

	// Authorize: Checks if the user is authorized to perform the action on the collection
	//
	// # Requirements:
	//  - collection to be set
	//  - action to be set
	//  - domain to be set
	Authorize() (bool, error)

	MakeLastModified() (*rbac_pb.StateActivity, error)
}

// AuthTransactionObjectInterfaces

type CollectionHolder interface {
	GetCollection() *rbac_pb.Collection
}
type CollectionIdHolder interface {
	GetCollectionId() *rbac_pb.Collection_Id
}
type UserHolder interface {
	GetUser() *rbac_pb.User
}
type UserIdHolder interface {
	GetUserId() *rbac_pb.User_Id
}

// type RoleWrapperInterface interface {
//     GetRole() (int, error)
// }
// type RolePermissionWrapperInterface interface {
//     GetRolePermission(role int) (*rbac_pb.ACL, error)
// }
