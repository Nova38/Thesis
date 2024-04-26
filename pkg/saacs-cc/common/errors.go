package common

import (
	"encoding/json"
	"errors"
	"fmt"

	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"
)

var (
	Unspecified            = errors.New(pb.TxError_UNSPECIFIED.String())
	RequestInvalid         = errors.New(pb.TxError_REQUEST_INVALID.String())
	Runtime                = errors.New(pb.TxError_RUNTIME.String())
	RuntimeBadOps          = errors.New(pb.TxError_RUNTIME_BAD_OPS.String())
	KeyNotFound            = errors.New(pb.TxError_KEY_NOT_FOUND.String())
	AlreadyExists          = errors.New(pb.TxError_KEY_ALREADY_EXISTS.String())
	CollectionInvalidId    = errors.New(pb.TxError_COLLECTION_INVALID_ID.String())
	CollectionUnregistered = errors.New(pb.TxError_COLLECTION_UNREGISTERED.String())
	AlreadyRegistered      = errors.New(
		pb.TxError_COLLECTION_ALREADY_REGISTERED.String(),
	)
	CollectionInvalid         = errors.New(pb.TxError_COLLECTION_INVALID.String())
	CollectionInvalidItemType = errors.New(
		pb.TxError_COLLECTION_INVALID_ITEM_TYPE.String(),
	)
	CollectionInvalidRoleId = errors.New(
		pb.TxError_COLLECTION_INVALID_ROLE_ID.String(),
	)
	UserInvalidId         = errors.New(pb.TxError_USER_INVALID_ID.String())
	UserUnregistered      = errors.New(pb.TxError_USER_UNREGISTERED.String())
	UserAlreadyRegistered = errors.New(pb.TxError_USER_ALREADY_REGISTERED.String())
	UserInvalid           = errors.New(pb.TxError_USER_INVALID.String())
	UserNoRole            = errors.New(pb.TxError_USER_NO_ROLE.String())
	UserPermissionDenied  = errors.New(pb.TxError_USER_PERMISSION_DENIED.String())
	ItemInvalidId         = errors.New(pb.TxError_ITEM_INVALID_ID.String())
	ItemUnregistered      = errors.New(pb.TxError_ITEM_UNREGISTERED.String())
	ItemAlreadyRegistered = errors.New(pb.TxError_ITEM_ALREADY_REGISTERED.String())
	ItemInvalid           = errors.New(pb.TxError_ITEM_INVALID.String())
	InvalidItemFieldPath  = errors.New(pb.TxError_INVALID_ITEM_FIELD_PATH.String())
	InvalidItemFieldValue = errors.New(
		pb.TxError_INVALID_ITEM_FIELD_VALUE.String(),
	)
)

// var KeyNotFoundError = errors.New("Key not found")

type VerboseError struct {
	Err error
}

func WrapError(err error) error {

	fmt.Printf("%+v", err)

	// b, err1 := json.MarshalIndent(err, "", "  ")
	b, err := json.Marshal(err)
	if err != nil {
		fmt.Print(err)
		return err
	}

	return errors.New(string(b))
}
