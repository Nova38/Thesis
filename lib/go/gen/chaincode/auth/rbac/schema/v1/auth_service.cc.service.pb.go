// Code generated by proto-gen-go-auth_pb. DO NOT EDIT.
// versions:
// - protoc-gen-cckey v0.0.1
// source: chaincode/auth/rbac/schema/v1/auth_service.proto

package schemav1

import (
	fmt "fmt"
	state "github.com/nova38/thesis/lib/go/fabric/auth/state"
	v1 "github.com/nova38/thesis/lib/go/gen/auth/v1"
)

// Service AuthService
type AuthServiceInterface[T state.GenericAuthTxCtxInterface] interface {
	// UserGetCurrent: Returns the current user.
	//
	// Returns the current user.
	// # Requires:
	//   - User submitting the transaction is a registered user.
	//
	// # Operation:
	//   - Domain: ACTION_OBJECT_VIEW
	//
	// req is empty
	UserGetCurrent(ctx T) (res *UserGetCurrentResponse, err error)

	// Returns the current user id.
	//
	// # Requires:
	//   - User submitting the transaction is a registered user.
	//
	// # Operation:
	//   - Domain: ACTION_OBJECT_VIEW
	//
	// req is empty
	UserGetCurrentId(ctx T) (res *UserGetCurrentIdResponse, err error)

	// UserGetList: Returns the list of users.
	//
	// # Requires:
	//   - Non-register users can call this method.
	//
	// # Operation:
	//   - Domain: ACTION_OBJECT_VIEW
	//
	// req is empty
	UserGetList(ctx T) (res *UserGetListResponse, err error)

	// UserGet: Returns the user.
	//
	// # Requires:
	//   - Non-register users can call this method.
	//
	// # Operation:
	//   - Domain: ACTION_OBJECT_VIEW
	UserGet(ctx T, req *UserGetRequest) (res *UserGetResponse, err error)

	// UserGetHistory: Returns the user history.
	//
	// # Requires:
	//   - Non-register users can call this method.
	//
	// # Operation:
	//   - Domain: ACTION_OBJECT_VIEW_HISTORY
	UserGetHistory(ctx T, req *UserGetHistoryRequest) (res *UserGetHistoryResponse, err error)

	// UserRegister: Registers the user.
	//
	// # Requires:
	//   - The certificate for the user submitting this request must not be already registered as a user.
	//
	// # Operation:
	//   - Domain: ACTION_REGISTER_USER
	UserRegister(ctx T, req *UserRegisterRequest) (res *UserRegisterResponse, err error)

	// UserUpdateMembership: Updates the user's membership.
	//
	// # Requires:
	//   - User submitting the transaction is a registered user.
	//   - The specified user id is a registered user.
	//   - The specified collection id is a registered collection.
	//   - The user submitting the transaction is a member of the specified collection.
	//   - The user submitting the transaction the a role who has permission
	//     to update the membership of the specified collection.
	//
	// # Operation:
	//   - Domain: ACTION_OBJECT_EDIT
	UserUpdateMembership(ctx T, req *UserUpdateMembershipRequest) (res *UserUpdateMembershipResponse, err error)

	// CollectionGetList: Returns the list of collections.
	//
	// # Requires:
	//   - Non-register users can call this method.
	//
	// # Operation:
	//   - Domain: ACTION_OBJECT_VIEW
	//
	// req is empty
	CollectionGetList(ctx T) (res *CollectionGetListResponse, err error)

	// CollectionGet: Returns the collection.
	//
	// # Requires:
	//   - Non-register users can call this method.
	//
	// # Operation:
	//   - Domain: ACTION_OBJECT_VIEW
	CollectionGet(ctx T, req *CollectionGetRequest) (res *CollectionGetResponse, err error)

	// CollectionGetHistory: Returns the collection history.
	//
	// # Requires:
	//   - Non-register users can call this method.
	//
	// # Operation:
	//   - Domain: ACTION_OBJECT_VIEW_HISTORY
	CollectionGetHistory(ctx T, req *CollectionGetHistoryRequest) (res *CollectionGetHistoryResponse, err error)

	// CollectionCreate: Creates the collection.
	//
	// # Requires:
	//   - User submitting the transaction is a registered user.
	//
	// # Operation:
	//   - Domain: ACTION_REGISTER_COLLECTION
	CollectionCreate(ctx T, req *CollectionCreateRequest) (res *CollectionCreateResponse, err error)

	// CollectionUpdateRoles
	//
	// # Operation:
	//   - Domain: ACTION_OBJECT_EDIT
	CollectionUpdateRoles(ctx T, req *CollectionUpdateRolesRequest) (res *CollectionUpdateRolesResponse, err error)

	// CollectionUpdatePermission
	//
	// # Operation:
	//   - Domain: ACTION_OBJECT_EDIT
	CollectionUpdatePermission(ctx T, req *CollectionUpdatePermissionRequest) (res *CollectionUpdatePermissionResponse, err error)
}

type AuthServiceBase struct {
}

func (s *AuthServiceBase) GetEvaluateTransactions() []string {
	return []string{
		"UserGetCurrent",
		"UserGetCurrentId",
		"UserGetList",
		"UserGet",
		"UserGetHistory",
		"CollectionGetList",
		"CollectionGet",
		"CollectionGetHistory",
	}
}

func AuthServiceGetTxOperation(txName string) (op *v1.Operation, err error) {
	switch txName {
	case "UserGetCurrent":
		// action:ACTION_OBJECT_VIEW
		return &v1.Operation{
			Action: 10,
		}, nil
	case "UserGetCurrentId":
		// action:ACTION_OBJECT_VIEW
		return &v1.Operation{
			Action: 10,
		}, nil
	case "UserGetList":
		// action:ACTION_OBJECT_VIEW
		return &v1.Operation{
			Action: 10,
		}, nil
	case "UserGet":
		// action:ACTION_OBJECT_VIEW
		return &v1.Operation{
			Action: 10,
		}, nil
	case "UserGetHistory":
		// action:ACTION_OBJECT_VIEW_HISTORY
		return &v1.Operation{
			Action: 18,
		}, nil
	case "UserRegister":
		// action:ACTION_REGISTER_USER
		return &v1.Operation{
			Action: 1,
		}, nil
	case "UserUpdateMembership":
		// action:ACTION_OBJECT_EDIT
		return &v1.Operation{
			Action: 12,
		}, nil
	case "CollectionGetList":
		// action:ACTION_OBJECT_VIEW
		return &v1.Operation{
			Action: 10,
		}, nil
	case "CollectionGet":
		// action:ACTION_OBJECT_VIEW
		return &v1.Operation{
			Action: 10,
		}, nil
	case "CollectionGetHistory":
		// action:ACTION_OBJECT_VIEW_HISTORY
		return &v1.Operation{
			Action: 18,
		}, nil
	case "CollectionCreate":
		// action:ACTION_REGISTER_COLLECTION
		return &v1.Operation{
			Action: 2,
		}, nil
	case "CollectionUpdateRoles":
		// action:ACTION_OBJECT_EDIT
		return &v1.Operation{
			Action: 12,
		}, nil
	case "CollectionUpdatePermission":
		// action:ACTION_OBJECT_EDIT
		return &v1.Operation{
			Action: 12,
		}, nil
	default:
		return nil, fmt.Errorf("No operation defined for " + txName)
	}
	return nil, nil
}

func (s *AuthServiceBase) GetIgnoredFunctions() []string {
	return []string{"GetTxOperation"}
}