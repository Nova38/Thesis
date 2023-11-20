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
func PrimaryExists[T common.ObjectInterface](ctx TxCtxInterface, obj T) bool {
	key := lo.Must(MakePrimaryKey(obj))
	return KeyExists(ctx, key)
}

func PrimaryGet[T common.ObjectInterface](ctx TxCtxInterface, obj T) (err error) {
	l := &Ledger[T]{ctx: ctx}
	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_VIEW,
		CollectionId: obj.ObjectKey().GetCollectionId(),
		Namespace:    obj.Namespace(),
		Paths:        nil,
	}
	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	return l.Get(obj)
}

func PrimaryGetFull[T common.ObjectInterface](
	ctx TxCtxInterface,
	obj T,
) (object *FullObject[T], err error) {
	l := &Ledger[T]{ctx: ctx}
	object = &FullObject[T]{}

	ops := []*authpb.Operation{
		{
			Action:       authpb.Action_ACTION_VIEW,
			CollectionId: obj.ObjectKey().GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		},
		{
			Action:       authpb.Action_ACTION_VIEW_HIDDEN_TXS,
			CollectionId: obj.ObjectKey().GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		},
		{
			Action:       authpb.Action_ACTION_VIEW_HIDDEN_TXS,
			CollectionId: obj.ObjectKey().GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		},
	}
	if auth, err := ctx.Authorize(ops); !auth || err != nil {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}
	// Get the object
	if err = l.Get(obj); err != nil {
		return nil, oops.Wrap(err)
	}
	object.Object = obj

	object.Suggestions, _, err = SuggestionListByObject(ctx, obj.ObjectKey(), "")
	if err != nil {
		return nil, oops.Wrap(err)
	}

	// Get the history
	object.History, err = history(ctx, obj, true)
	if err != nil {
		return nil, oops.Wrap(err)
	}

	// state.ObjectToAuthObj(obj, object.Value)

	return object, nil
}

func PrimaryByPartialKey[T common.ObjectInterface](
	ctx TxCtxInterface,
	obj T,
	numAttr int,
	bookmark string,
) (list []T, mk string, err error) {
	l := &Ledger[T]{ctx: ctx}

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_VIEW,
		CollectionId: obj.ObjectKey().GetCollectionId(),
		Namespace:    obj.Namespace(),
		Paths:        nil,
	}

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return nil, "", oops.Wrap(common.UserPermissionDenied)
	}

	return l.GetPartialKeyList(obj, numAttr, bookmark)
}

func PrimaryList[T common.ObjectInterface](
	ctx TxCtxInterface,
	obj T,
	bookmark string,
) (list []T, mk string, err error) {
	return PrimaryByPartialKey(ctx, obj, 1, bookmark)
}

func ByCollection[T common.ObjectInterface](
	ctx TxCtxInterface,
	obj T,
	bookmark string,
) (list []T, mk string, err error) {
	return PrimaryByPartialKey(ctx, obj, 1, bookmark)
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
func PrimaryCreate[T common.ObjectInterface](ctx TxCtxInterface, obj T) (err error) {
	l := &Ledger[T]{
		ctx: ctx,
	}

	// Authorize the operation
	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_CREATE,
		CollectionId: obj.ObjectKey().GetCollectionId(),
		Namespace:    obj.Namespace(),
		Paths:        nil,
	}

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	return l.Create(obj)
}

func PrimaryUpdate[T common.ObjectInterface](
	ctx TxCtxInterface,
	obj T,
	mask *fieldmaskpb.FieldMask,
) (err error) {
	l := &Ledger[T]{
		ctx: ctx,
	}
	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_UPDATE,
		CollectionId: obj.ObjectKey().GetCollectionId(),
		Namespace:    obj.Namespace(),
		Paths:        mask,
	}

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	return l.Update(obj, mask)
}

func PrimaryDelete[T common.ObjectInterface](ctx TxCtxInterface, obj T) (err error) {
	l := &Ledger[T]{
		ctx: ctx,
	}
	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_DELETE,
		CollectionId: obj.ObjectKey().GetCollectionId(),
		Namespace:    obj.Namespace(),
		Paths:        nil,
	}

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	err = l.Delete(obj)

	// TODO: Handle deleting refs/sub objects here

	// TODO: Handle deleting suggestions here
	// TODO: Handle deleting hiddenTx objects here

	return err
}

// func PrimaryDeleteFromKey(ctx TxCtxInterface, key *authpb.ObjectKey) (obj *authpb.Object, err error) {
// 	k, err := MakeObjectKeyPrimary(key)
// 	if err != nil {
// 		return obj, oops.Wrap(err)
// 	}

// 	bytes, err := ctx.GetStub().GetState(k)

// 	if bytes == nil && err == nil {
// 		return nil, oops.Wrap(common.KeyNotFound)
// 	}

// 	if err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	obj = &authpb.Object{}
// 	if err = json.Unmarshal(bytes, obj); err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	// obj, err = l.GetFromKey(k)
// 	// if err != nil {
// 	// 	return obj, oops.Wrap(err)
// 	// }

// 	// err = PrimaryDelete(ctx, obj)
// 	// if err != nil {
// 	// 	return obj, oops.Wrap(err)
// 	// }

// 	return obj, nil
// }

// func PrimaryGetFromKey[T common.ObjectInterface](ctx TxCtxInterface, key *authpb.ObjectKey) (obj T, err error) {
// 	l := &Ledger[T]{ctx: ctx}
// 	op := &authpb.Operation{
// 		Action:       authpb.Action_ACTION_VIEW,
// 		CollectionId: key.GetCollectionId(),
// 		Namespace:    key.GetObjectType(),
// 		Paths:        nil,
// 	}
// 	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
// 		return obj, oops.Wrap(common.UserPermissionDenied)
// 	}

// 	k, err := MakeObjectKeyPrimary(key)
// 	if err != nil {
// 		return obj, oops.Wrap(err)
// 	}

// 	obj, err = l.GetFromKey(k)
// 	if err != nil {
// 		return obj, oops.Wrap(err)
// 	}

// 	return obj, l.Get(obj)
// }
