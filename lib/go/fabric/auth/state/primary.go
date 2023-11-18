package state

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/common"

	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Primary Objects

// ──────────────────────────────────────────────────
// Query Suggested Functions
// ──────────────────────────────────────────────────
func PrimaryExists[T Object](ctx TxCtxInterface, obj T) bool {
	key := lo.Must(MakeCompositeKey(obj))
	return KeyExists(ctx, key)
}

func PrimaryGet[T Object](ctx TxCtxInterface, obj T) (err error) {
	l := &Ledger[T]{ctx: ctx}
	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_OBJECT_VIEW,
		CollectionId: obj.GetCollectionId(),
		Namespace:    obj.Namespace(),
		Paths:        nil,
	}
	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	return l.Get(obj)
}

func PrimaryByPartialKey[T Object](
	ctx TxCtxInterface,
	obj T,
	numAttr int,
	bookmark string,
) (list []T, mk string, err error) {
    l := &Ledger[T]{ctx: ctx}

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_OBJECT_VIEW,
		CollectionId: obj.GetCollectionId(),
		Namespace:    obj.Namespace(),
		Paths:        nil,
	}

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return nil, "", oops.Wrap(common.UserPermissionDenied)
	}



	return l.GetPartialKeyList(obj, numAttr, bookmark)
}

func PrimaryList[T Object](
	ctx TxCtxInterface,
	obj T,
	bookmark string,
) (list []T, mk string, err error) {
	return PrimaryByPartialKey(ctx, obj, len(obj.Key()), bookmark)
}

func ByCollection[T Object](
	ctx TxCtxInterface,
	obj T,
	bookmark string,
) (list []T, mk string, err error) {
	return PrimaryByPartialKey(ctx, obj, len(obj.Key())-1, bookmark)
}

// ──────────────────────────────────────────────────
// Invoke Suggested Functions
// ──────────────────────────────────────────────────

// PrimaryCreate creates the object in the ledger
// returns error if the object already exists
// will panic if
//   - the key cannot be created,
//   - the object cannot be marshalled
//   - Authorization errors
func PrimaryCreate[T Object](ctx TxCtxInterface, obj T) (err error) {
	l := &Ledger[T]{
		ctx: ctx,
	}

	// Authorize the operation
	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_OBJECT_CREATE,
		CollectionId: obj.GetCollectionId(),
		Namespace:    obj.Namespace(),
		Paths:        nil,
	}

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	return l.Create(obj)
}

func PrimaryUpdate[T Object](ctx TxCtxInterface, obj T, mask *fieldmaskpb.FieldMask) (err error) {
	l := &Ledger[T]{
		ctx: ctx,
	}
	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_OBJECT_UPDATE,
		CollectionId: obj.GetCollectionId(),
		Namespace:    obj.Namespace(),
		Paths:        mask,
	}

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	return l.Update(obj, mask)
}

func PrimaryDelete[T Object](ctx TxCtxInterface, obj T) (err error) {
	l := &Ledger[T]{
		ctx: ctx,
	}
	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_OBJECT_DELETE,
		CollectionId: obj.GetCollectionId(),
		Namespace:    obj.Namespace(),
		Paths:        nil,
	}

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	err = l.Delete(obj)

	// TODO: Handle deleting refs/sub objects here

	return err
}
