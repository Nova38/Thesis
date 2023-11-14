package common

import (
	"errors"

	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
)

//type InvalidStateObjectError struct {
//	MSG string
//}

var (
	Unspecified            = errors.New(authpb.TxError_UNSPECIFIED.String())
	RequestInvalid         = errors.New(authpb.TxError_REQUEST_INVALID.String())
	Runtime                = errors.New(authpb.TxError_RUNTIME.String())
	RuntimeBadOps          = errors.New(authpb.TxError_RUNTIME_BAD_OPS.String())
	KeyNotFound            = errors.New(authpb.TxError_KEY_NOT_FOUND.String())
	AlreadyExists          = errors.New(authpb.TxError_KEY_ALREADY_EXISTS.String())
	CollectionInvalidId    = errors.New(authpb.TxError_COLLECTION_INVALID_ID.String())
	CollectionUnregistered = errors.New(authpb.TxError_COLLECTION_UNREGISTERED.String())
	AlreadyRegistered      = errors.New(
		authpb.TxError_COLLECTION_ALREADY_REGISTERED.String(),
	)
	CollectionInvalid           = errors.New(authpb.TxError_COLLECTION_INVALID.String())
	CollectionInvalidObjectType = errors.New(
		authpb.TxError_COLLECTION_INVALID_OBJECT_TYPE.String(),
	)
	CollectionInvalidRoleId = errors.New(
		authpb.TxError_COLLECTION_INVALID_ROLE_ID.String(),
	)
	UserInvalidId           = errors.New(authpb.TxError_USER_INVALID_ID.String())
	UserUnregistered        = errors.New(authpb.TxError_USER_UNREGISTERED.String())
	UserAlreadyRegistered   = errors.New(authpb.TxError_USER_ALREADY_REGISTERED.String())
	UserInvalid             = errors.New(authpb.TxError_USER_INVALID.String())
	UserNoRole              = errors.New(authpb.TxError_USER_NO_ROLE.String())
	UserDeletedRole         = errors.New(authpb.TxError_USER_DELETED_ROLE.String())
	UserPermissionDenied    = errors.New(authpb.TxError_USER_PERMISSION_DENIED.String())
	ObjectInvalidId         = errors.New(authpb.TxError_OBJECT_INVALID_ID.String())
	ObjectUnregistered      = errors.New(authpb.TxError_OBJECT_UNREGISTERED.String())
	ObjectAlreadyRegistered = errors.New(authpb.TxError_OBJECT_ALREADY_REGISTERED.String())
	ObjectInvalid           = errors.New(authpb.TxError_OBJECT_INVALID.String())
	InvalidObjectFieldPath  = errors.New(authpb.TxError_INVALID_OBJECT_FIELD_PATH.String())
	InvalidObjectFieldValue = errors.New(
		authpb.TxError_INVALID_OBJECT_FIELD_VALUE.String(),
	)
)

// var KeyNotFoundError = errors.New("Key not found")
