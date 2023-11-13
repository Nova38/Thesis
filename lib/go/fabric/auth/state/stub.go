package state

import (
	"encoding/json"
	"fmt"
	"github.com/mennanov/fmutils"
	"github.com/nova38/thesis/lib/go/fabric/auth"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"log/slog"
	"slices"
)

// -
// primary objects

type Stub[T Object] struct{}

// MakeCompositeKey creates a composite key from the given attributes
func MakeCompositeKey[T Object](ctx TxCtxInterface, obj T) (key string, err error) {
	namespace := obj.Namespace()
	attr, err := obj.Key()
	if err != nil {
		return "", err
	}

	ctx.GetLogger().
		Info("MakeCompositeKey",
			slog.Group("Key", "Namespace", namespace, "attr", attr))

	return ctx.GetStub().CreateCompositeKey(namespace, attr)
}

func MakeHiddenKey[T Object](ctx TxCtxInterface, obj T) (hiddenKey string, err error) {
	attr, err := obj.Key()
	if err != nil {
		return "", err
	}
	attr = append([]string{obj.Namespace()}, attr...)

	return ctx.GetStub().CreateCompositeKey(auth.HiddenNamespace, attr)
}

func (s Stub[T]) Exists(ctx TxCtxInterface, key string) bool {
	bytes, err := ctx.GetStub().GetState(key)
	if bytes == nil && err == nil {
		return false
	}

	return err == nil
}

// Create creates the object in the ledger
// returns error if the object already exists
// will panic if
//   - the key cannot be created,
//   - the object cannot be marshalled
//   - Authorization errors
func (s Stub[T]) Create(ctx TxCtxInterface, obj T) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key        = lo.Must(MakeCompositeKey(ctx, obj))
		exists     = s.Exists(ctx, key)
		bytes      = lo.Must(json.Marshal(obj))
		authorized = lo.Must(ctx.Authorize(&authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_CREATE,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		}))
	)

	if !authorized {
		return oops.Wrap(auth.UserPermissionDenied)
	}
	if exists {
		return oops.
			With("Key", key, "Namespace", obj.Namespace()).
			Wrap(auth.AlreadyExists)
	}

	return ctx.GetStub().PutState(key, bytes)
}

func (s Stub[T]) Edit(ctx TxCtxInterface, obj T, mask *fieldmaskpb.FieldMask) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key        = lo.Must(MakeCompositeKey(ctx, obj))
		exists     = s.Exists(ctx, key)
		bytes      = lo.Must(ctx.GetStub().GetState(key))
		authorized = lo.Must(ctx.Authorize(&authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_EDIT,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        mask,
		}))
	)
	if !authorized {
		return oops.Wrap(auth.UserPermissionDenied)
	}
	if !exists {
		return oops.
			With("Key", key, "Namespace", obj.Namespace()).
			Wrap(auth.KeyNotFound)
	}

	cur := new(T)
	lo.Must0(json.Unmarshal(bytes, cur))

	// TODO: change last modified

	fmutils.Filter(obj, mask.Paths)
	proto.Merge(*cur, obj)

	bytes = lo.Must(json.Marshal(cur))

	return ctx.GetStub().PutState(key, bytes)
}

func (s Stub[T]) Delete(ctx TxCtxInterface, obj T) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key        = lo.Must(MakeCompositeKey(ctx, obj))
		exists     = s.Exists(ctx, key)
		authorized = lo.Must(ctx.Authorize(&authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_DELETE,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		}))
	)

	//TODO: check if the object is referenced by other objects

	if !authorized {
		return oops.Wrap(auth.UserPermissionDenied)
	}
	if !exists {
		return oops.
			With("Key", key, "Namespace", obj.Namespace()).
			Wrap(auth.KeyNotFound)
	}

	return ctx.GetStub().DelState(key)
}

func (s Stub[T]) Get(ctx TxCtxInterface, obj T, mask *fieldmaskpb.FieldMask) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key        = lo.Must(MakeCompositeKey(ctx, obj))
		bytes      = lo.Must(ctx.GetStub().GetState(key))
		authorized = lo.Must(ctx.Authorize(&authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_VIEW,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        mask,
		}))
	)
	if !authorized {
		return oops.Wrap(auth.UserPermissionDenied)
	}

	if err = json.Unmarshal(bytes, obj); err != nil {
		return err
	}
	if mask != nil {
		if !mask.IsValid(obj) {
			return fmt.Errorf("mask is not valid")
		}
		fmutils.Filter(obj, mask.Paths)
	}

	return nil
}

func (s Stub[T]) List(
	ctx TxCtxInterface,
	obj T,
	bookmark string,
	mask *fieldmaskpb.FieldMask,
) (list []T, mk string, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	return s.ByPartialKey(ctx, obj, 0, bookmark, mask)
}

func (s Stub[T]) ByCollection(
	ctx TxCtxInterface,
	obj T,
	bookmark string,
	mask *fieldmaskpb.FieldMask,
) (list []T, mk string, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	return s.ByPartialKey(ctx, obj, 1, bookmark, mask)
}

func (s Stub[T]) ByPartialKey(
	ctx TxCtxInterface,
	obj T,
	numAttr int,
	bookmark string,
	mask *fieldmaskpb.FieldMask,
) (list []T, mk string, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	var (
		attr       = lo.Must(obj.Key())
		viewMask   = fmutils.NestedMask{}
		authorized = lo.Must(ctx.Authorize(&authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_VIEW,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        mask,
		}))
	)

	if !authorized {
		return nil, "", oops.Wrap(auth.UserPermissionDenied)
	}

	if len(attr) == 0 {
		return nil, "", fmt.Errorf("key is empty")
	} else if len(attr) < numAttr {
		return nil, "", fmt.Errorf("key has less than %d attributes", numAttr)
	}

	if mask != nil {
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

func (s Stub[T]) hiddenTxs(ctx TxCtxInterface, obj T) (list *authpb.HiddenTxList, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key    = lo.Must(MakeHiddenKey(ctx, obj))
		exists = s.Exists(ctx, key)
	)
	if !exists {
		return nil, oops.Wrap(auth.KeyNotFound)
	}

	state, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, err
	}

	list = new(authpb.HiddenTxList)
	lo.Must0(json.Unmarshal(state, list))

	return list, nil
}

func (s Stub[T]) history(
	ctx TxCtxInterface,
	obj T,
	showHidden bool,
) (history *authpb.History, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key    = lo.Must(MakeCompositeKey(ctx, obj))
		exists = s.Exists(ctx, key)
		hidden = lo.Must(s.hiddenTxs(ctx, obj))
	)

	if !exists {
		return nil, oops.Wrap(auth.KeyNotFound)
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

func (s Stub[T]) HiddenTx(ctx TxCtxInterface, obj T) (list *authpb.HiddenTxList, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		authorized = lo.Must(ctx.Authorize(&authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_VIEW_HIDDEN_TXS,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		}))
	)
	if !authorized {
		return nil, oops.Wrap(auth.UserPermissionDenied)
	}

	return s.hiddenTxs(ctx, obj)
}

func (s Stub[T]) History(ctx TxCtxInterface, obj T) (history *authpb.History, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		authorized = lo.Must(ctx.Authorize(&authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_VIEW_HISTORY,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		}))
	)
	if !authorized {
		return nil, oops.Wrap(auth.UserPermissionDenied)
	}

	return s.history(ctx, obj, false)
}

func (s Stub[T]) FullHistory(ctx TxCtxInterface, obj T) (history *authpb.History, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	var (
		authorized = lo.Must(ctx.Authorize(&authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_VIEW_HISTORY,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		}))
	)
	if !authorized {
		return nil, oops.Wrap(auth.UserPermissionDenied)
	}

	return s.history(ctx, obj, true)
}

//func (s Stub[T]) Transaction(ctx TxCtxInterface, obj T) (err error) {
//	defer func() { ctx.HandleFnError(&err, recover()) }()
//
//	//TODO implement me
//	panic("implement me")
//}

func (s Stub[T]) HideTransaction(ctx TxCtxInterface, obj T, tx *authpb.HiddenTx) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	var (
		key        = lo.Must(MakeHiddenKey(ctx, obj))
		hidden     = lo.Must(s.hiddenTxs(ctx, obj))
		authorized = lo.Must(ctx.Authorize(&authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_HIDE_TX,
			CollectionId: obj.GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		}))
	)
	if !authorized {
		return oops.Wrap(auth.UserPermissionDenied)
	}

	if hidden.Txs == nil {
		hidden.Txs = []*authpb.HiddenTx{}
	}

	if slices.ContainsFunc(hidden.Txs,
		func(tx *authpb.HiddenTx) bool { return tx.TxId == tx.TxId }) {
		return oops.Wrap(auth.AlreadyExists)
	}

	hidden.Txs = append(hidden.Txs, tx)

	bytes := lo.Must(json.Marshal(hidden))

	return ctx.GetStub().PutState(key, bytes)
}

// ════════════════════════════════════════════════════════
// Suggestion Functions
// ════════════════════════════════════════════════════════

func (s Stub[T]) Suggestion(
	ctx TxCtxInterface,
	suggestionId string,
) (suggestion *authpb.Suggestion, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	//TODO implement me
	panic("implement me")
}

func (s Stub[T]) CreateSuggestion(ctx TxCtxInterface, suggestion *authpb.Suggestion) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	//TODO implement me
	panic("implement me")
}

func (s Stub[T]) DeleteSuggestion(ctx TxCtxInterface, suggestion *authpb.Suggestion) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	//TODO implement me
	panic("implement me")
}

func (s Stub[T]) ApproveSuggestion(ctx TxCtxInterface, suggestion *authpb.Suggestion) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	//TODO implement me
	panic("implement me")
}

func (s Stub[T]) SuggestionList(
	ctx TxCtxInterface,
	obj T,
	bookmark string,
) (list []authpb.Suggestion, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	//TODO implement me
	panic("implement me")
}

func (s Stub[T]) SuggestionByCollection(
	ctx TxCtxInterface,
	obj T,
	bookmark string,
) (list []authpb.Suggestion, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	//TODO implement me
	panic("implement me")
}

func (s Stub[T]) GetSuggestionListByPartialKey(
	ctx TxCtxInterface,
	obj T,
	numAttr int,
	bookmark string,
) (list []authpb.Suggestion, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	//TODO implement me
	panic("implement me")
}

// ════════════════════════════════════════════════════════
// Reference Functions
// ════════════════════════════════════════════════════════
