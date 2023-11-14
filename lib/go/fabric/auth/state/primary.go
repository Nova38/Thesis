package state

import (
	"encoding/json"
	"fmt"

	"github.com/nova38/thesis/lib/go/fabric/auth/common"

	"github.com/mennanov/fmutils"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func KeyExists(ctx TxCtxInterface, key string) bool {
	bytes, err := ctx.GetStub().GetState(key)
	if bytes == nil && err == nil {
		return false
	}

	return err == nil
}

// Primary Objects

// ──────────────────────────────────────────────────
// Query Suggested Functions
// ──────────────────────────────────────────────────
func PrimaryExists[T Object](ctx TxCtxInterface, obj T) bool {
	key := lo.Must(MakeCompositeKey(ctx, obj))
	return KeyExists(ctx, key)
}

func Get[T Object](ctx TxCtxInterface, obj T) (err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		mask  = ctx.GetViewMask()
		key   = lo.Must(MakeCompositeKey(ctx, obj))
		bytes = lo.Must(ctx.GetStub().GetState(key))
		op    = &authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_VIEW,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        mask,
		}
		authorized = lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	)

	if !authorized {
		return oops.Wrap(common.UserPermissionDenied)
	}

	if err = json.Unmarshal(bytes, obj); err != nil {
		return err
	}
	if mask != nil && len(mask.Paths) > 0 {
		if !mask.IsValid(obj) {
			return fmt.Errorf("mask is not valid")
		}
		fmutils.Filter(obj, mask.Paths)
	}

	return nil
}

func List[T Object](
	ctx TxCtxInterface,
	obj T,
	bookmark string,
) (list []T, mk string, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	return ByPartialKey(ctx, obj, 0, bookmark)
}

func ByCollection[T Object](
	ctx TxCtxInterface,
	obj T,
	bookmark string,
) (list []T, mk string, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	return ByPartialKey(ctx, obj, 1, bookmark)
}

func ByPartialKey[T Object](
	ctx TxCtxInterface,
	obj T,
	numAttr int,
	bookmark string,
) (list []T, mk string, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	var (
		mask     = ctx.GetViewMask()
		attr     = lo.Must(obj.Key())
		viewMask = fmutils.NestedMask{}
		op       = &authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_VIEW,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        mask,
		}
		authorized = lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	)

	if !authorized {
		return nil, "", oops.Wrap(common.UserPermissionDenied)
	}

	if len(attr) == 0 || len(attr) < numAttr {
		return nil, "", common.ObjectInvalid
	}

	if mask != nil && len(mask.Paths) > 0 {
		if !mask.IsValid(obj) {
			return nil, "", fmt.Errorf("mask is not valid")
		}
		viewMask = fmutils.NestedMaskFromPaths(mask.Paths)
	}

	attr = attr[:len(attr)-numAttr]
	results, metadata, err := ctx.GetStub().
		GetStateByPartialCompositeKeyWithPagination(obj.Namespace(), attr, ctx.GetPageSize(), bookmark)
	if err != nil {
		return nil, "", err
	}
	defer ctx.CloseQueryIterator(results)

	for results.HasNext() {
		queryResponse := lo.Must(results.Next())
		tmp := new(T)

		lo.Must0(json.Unmarshal(queryResponse.Value, tmp))

		if mask != nil {
			viewMask.Filter(*tmp)
		}

		list = append(list, *tmp)
	}

	if metadata == nil {
		return nil, "", fmt.Errorf("metadata is nil")
	}

	return list, metadata.GetBookmark(), nil
}

// ──────────────────────────────────────────────────
// Invoke Suggested Functions
// ──────────────────────────────────────────────────

// Create creates the object in the ledger
// returns error if the object already exists
// will panic if
//   - the key cannot be created,
//   - the object cannot be marshalled
//   - Authorization errors
func Create[T Object](ctx TxCtxInterface, obj T) (err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key    = lo.Must(MakeCompositeKey(ctx, obj))
		exists = KeyExists(ctx, key)
		bytes  = lo.Must(json.Marshal(obj))
		op     = &authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_CREATE,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		}
		authorized = lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	)

	if !authorized {
		return oops.Wrap(common.UserPermissionDenied)
	}
	if exists {
		return oops.
			With("Key", key, "Namespace", obj.Namespace()).
			Wrap(common.AlreadyExists)
	}

	return ctx.GetStub().PutState(key, bytes)
}

func Update[T Object](ctx TxCtxInterface, obj T, mask *fieldmaskpb.FieldMask) (err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key   = lo.Must(MakeCompositeKey(ctx, obj))
		bytes = lo.Must(ctx.GetStub().GetState(key))
		op    = &authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_UPDATE,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        mask,
		}
		authorized = lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	)
	if !authorized {
		return oops.Wrap(common.UserPermissionDenied)
	}

	cur := new(T)
	lo.Must0(json.Unmarshal(bytes, cur))
	bytes = []byte{}

	// TODO: change last modified

	fmutils.Filter(obj, mask.Paths)
	proto.Merge(*cur, obj)

	bytes = lo.Must(json.Marshal(cur))
	lo.Must0(json.Unmarshal(bytes, obj))

	return ctx.GetStub().PutState(key, bytes)
}

func Delete[T Object](ctx TxCtxInterface, obj T) (err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key    = lo.Must(MakeCompositeKey(ctx, obj))
		exists = KeyExists(ctx, key)
		op     = &authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_DELETE,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		}
		authorized = lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	)

	// TODO: check if the object is referenced by other objects

	if !authorized {
		return oops.Wrap(common.UserPermissionDenied)
	}
	if !exists {
		return oops.
			With("Key", key, "Namespace", obj.Namespace()).
			Wrap(common.KeyNotFound)
	}

	return ctx.GetStub().DelState(key)
}
