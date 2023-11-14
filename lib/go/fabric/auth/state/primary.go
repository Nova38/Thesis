package state

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/nova38/thesis/lib/go/fabric/auth/common"

	"github.com/mennanov/fmutils"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func Exists(ctx TxCtxInterface, key string) bool {
	bytes, err := ctx.GetStub().GetState(key)
	if bytes == nil && err == nil {
		return false
	}

	return err == nil
}

// Primary Objects

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
		exists = Exists(ctx, key)
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

	// TODO: change last modified

	fmutils.Filter(obj, mask.Paths)
	proto.Merge(*cur, obj)

	bytes = lo.Must(json.Marshal(cur))

	return ctx.GetStub().PutState(key, bytes)
}

func Delete[T Object](ctx TxCtxInterface, obj T) (err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key    = lo.Must(MakeCompositeKey(ctx, obj))
		exists = Exists(ctx, key)
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

// ════════════════════════════════════════════════════════
// History Functions
// ════════════════════════════════════════════════════════

func hiddenTxs[T Object](ctx TxCtxInterface, obj T) (list *authpb.HiddenTxList, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key    = lo.Must(MakeHiddenKey(ctx, obj))
		exists = Exists(ctx, key)
	)
	if !exists {
		return nil, oops.Wrap(common.KeyNotFound)
	}

	state, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, err
	}

	list = new(authpb.HiddenTxList)
	lo.Must0(json.Unmarshal(state, list))

	return list, nil
}

func history[T Object](
	ctx TxCtxInterface,
	obj T,
	showHidden bool,
) (history *authpb.History, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key    = lo.Must(MakeCompositeKey(ctx, obj))
		exists = Exists(ctx, key)
		hidden = lo.Must(hiddenTxs(ctx, obj))
	)

	if !exists {
		return nil, oops.Wrap(common.KeyNotFound)
	}

	history = &authpb.History{
		Entries:   []*authpb.HistoryEntry{},
		HiddenTxs: lo.Ternary(showHidden, hidden, nil),
	}

	resultIterator := lo.Must(ctx.GetStub().GetHistoryForKey(key))
	defer ctx.CloseQueryIterator(resultIterator)

	for resultIterator.HasNext() {
		queryResponse := lo.Must(resultIterator.Next())
		entry := &authpb.HistoryEntry{
			TxId:      queryResponse.TxId,
			IsDelete:  queryResponse.IsDelete,
			Timestamp: queryResponse.Timestamp,
			Note:      "",
		}

		if hidden.Txs != nil {
			for _, tx := range hidden.Txs {
				if tx.TxId == entry.TxId {
					entry.Note = tx.Note
					entry.IsHidden = true
				}
			}
		}

		if !showHidden && entry.IsHidden {
			continue
		}

		tmp := new(T)
		lo.Must0(json.Unmarshal(queryResponse.Value, tmp))
		entry.Value = lo.Must(anypb.New(*tmp))

		history.Entries = append(history.Entries, entry)
	}

	return history, err
}

func HiddenTx[T Object](ctx TxCtxInterface, obj T) (l *authpb.HiddenTxList, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		op = &authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_VIEW_HIDDEN_TXS,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		}
		authorized = lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	)

	if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	return hiddenTxs(ctx, obj)
}

func History[T Object](ctx TxCtxInterface, obj T) (h *authpb.History, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_OBJECT_VIEW_HISTORY,
		CollectionId: obj.GetCollectionId(),
		Namespace:    obj.Namespace(),
		Paths:        nil,
	}

	authorized := lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	return history(ctx, obj, false)
}

func FullHistory[T Object](ctx TxCtxInterface, obj T) (h *authpb.History, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	var (
		op = &authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_VIEW_HISTORY,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		}
		authorized = lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	)

	if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	return history(ctx, obj, true)
}

func HideTransaction[T Object](ctx TxCtxInterface, obj T, tx *authpb.HiddenTx) (err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	var (
		key    = lo.Must(MakeHiddenKey(ctx, obj))
		hidden = lo.Must(hiddenTxs(ctx, obj))
		op     = &authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_HIDE_TX,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		}
		authorized = lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	)
	if !authorized {
		return oops.Wrap(common.UserPermissionDenied)
	}

	if hidden.Txs == nil {
		hidden.Txs = []*authpb.HiddenTx{}
	}

	if slices.ContainsFunc(hidden.Txs,
		func(e *authpb.HiddenTx) bool { return e.TxId == tx.TxId }) {
		return oops.Wrap(common.AlreadyExists)
	}

	hidden.Txs = append(hidden.Txs, tx)

	bytes := lo.Must(json.Marshal(hidden))

	return ctx.GetStub().PutState(key, bytes)
}
