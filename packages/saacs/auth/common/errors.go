package common

import (
	"encoding/json"
	"errors"
	"fmt"

	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
)

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
	CollectionInvalid         = errors.New(authpb.TxError_COLLECTION_INVALID.String())
	CollectionInvalidItemType = errors.New(
		authpb.TxError_COLLECTION_INVALID_ITEM_TYPE.String(),
	)
	CollectionInvalidRoleId = errors.New(
		authpb.TxError_COLLECTION_INVALID_ROLE_ID.String(),
	)
	UserInvalidId         = errors.New(authpb.TxError_USER_INVALID_ID.String())
	UserUnregistered      = errors.New(authpb.TxError_USER_UNREGISTERED.String())
	UserAlreadyRegistered = errors.New(authpb.TxError_USER_ALREADY_REGISTERED.String())
	UserInvalid           = errors.New(authpb.TxError_USER_INVALID.String())
	UserNoRole            = errors.New(authpb.TxError_USER_NO_ROLE.String())
	UserPermissionDenied  = errors.New(authpb.TxError_USER_PERMISSION_DENIED.String())
	ItemInvalidId         = errors.New(authpb.TxError_ITEM_INVALID_ID.String())
	ItemUnregistered      = errors.New(authpb.TxError_ITEM_UNREGISTERED.String())
	ItemAlreadyRegistered = errors.New(authpb.TxError_ITEM_ALREADY_REGISTERED.String())
	ItemInvalid           = errors.New(authpb.TxError_ITEM_INVALID.String())
	InvalidItemFieldPath  = errors.New(authpb.TxError_INVALID_ITEM_FIELD_PATH.String())
	InvalidItemFieldValue = errors.New(
		authpb.TxError_INVALID_ITEM_FIELD_VALUE.String(),
	)
)

// var KeyNotFoundError = errors.New("Key not found")

type VerboseError struct {
	Err error
}

func WrapError(err error) error {
	b, err1 := json.MarshalIndent(err, "", "  ")
	if err1 != nil {
		fmt.Print(err1)
		return err
	}

	return errors.New(string(b))
}
