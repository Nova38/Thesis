package rbac

import (
	"github.com/nova38/thesis/lib/go/fabric/state"

	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
)

// An interface for a transaction context that has a user and collection

type TxCtxInterface interface {
	state.LogedTxCtxInterface

	// GetUserId() (*rbac_pb.User_Id, error)

	// Uses the ctx stub to get the user from the state
	GetUser() (*rbac_pb.User, error)

	// Gets the collection value from the state
	// requires:
	//  - collection to be set
	GetCollection() (*rbac_pb.Collection, error)

	// Sets the collection value in the state
	SetCollection(collection *rbac_pb.Collection) error

	// GetRole requires a collection to be set
	// requires:
	//  - collection to be set
	//  - user to be set
	// if user is not a member of the collection it returns 0
	GetRole() (int, error)

	//
	GetRolePermission(role int, action rbac_pb.Operations_Action) (*rbac_pb.Operations, error)
}
