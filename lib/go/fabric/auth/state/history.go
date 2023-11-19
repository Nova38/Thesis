package state

import (
	"encoding/json"
	"slices"

	"github.com/nova38/thesis/lib/go/fabric/auth/common"

	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/types/known/anypb"
)

// ════════════════════════════════════════════════════════
// History Functions
// ════════════════════════════════════════════════════════

func hiddenTxs[T common.ObjectInterface](
	ctx TxCtxInterface,
	obj T,
) (list *authpb.HiddenTxList, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key    = lo.Must(MakeHiddenKey(obj))
		exists = KeyExists(ctx, key)
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

func history[T common.ObjectInterface](
	ctx TxCtxInterface,
	obj T,
	showHidden bool,
) (history *authpb.History, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key    = lo.Must(MakePrimaryKey(obj))
		exists = KeyExists(ctx, key)
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
			TxId:      queryResponse.GetTxId(),
			IsDelete:  queryResponse.GetIsDelete(),
			Timestamp: queryResponse.GetTimestamp(),
			Note:      "",
		}

		if hidden.Txs != nil {
			for _, tx := range hidden.GetTxs() {
				if tx.GetTxId() == entry.GetTxId() {
					entry.Note = tx.GetNote()
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

func History[T common.ObjectInterface](ctx TxCtxInterface, obj T) (h *authpb.History, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_OBJECT_VIEW_HISTORY,
		CollectionId: obj.ObjectKey().GetCollectionId(),
		Namespace:    obj.Namespace(),
		Paths:        nil,
	}

	authorized := lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	return history(ctx, obj, false)
}

func FullHistory[T common.ObjectInterface](
	ctx TxCtxInterface,
	obj T,
) (h *authpb.History, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	var (
		op = &authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_VIEW_HISTORY,
			CollectionId: obj.ObjectKey().GetCollectionId(),
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

func HiddenTx[T common.ObjectInterface](
	ctx TxCtxInterface,
	obj T,
) (l *authpb.HiddenTxList, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		op = &authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_VIEW_HIDDEN_TXS,
			CollectionId: obj.ObjectKey().GetCollectionId(),
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

func HideTransaction[T common.ObjectInterface](
	ctx TxCtxInterface,
	obj T,
	tx *authpb.HiddenTx,
) (hiddenList *authpb.HiddenTxList, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	var (
		key    = lo.Must(MakeHiddenKey(obj))
		hidden = lo.Must(hiddenTxs(ctx, obj))
		op     = &authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_HIDE_TX,
			CollectionId: obj.ObjectKey().GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		}
		authorized = lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	)
	if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	if hidden.Txs == nil {
		hidden.Txs = []*authpb.HiddenTx{}
	}

	if slices.ContainsFunc(hidden.Txs,
		func(e *authpb.HiddenTx) bool { return e.TxId == tx.TxId }) {
		return nil, oops.Wrap(common.AlreadyExists)
	}

	hidden.Txs = append(hidden.Txs, tx)

	bytes := lo.Must(json.Marshal(hidden))

	return hidden, ctx.GetStub().PutState(key, bytes)
}

func UnHideTransaction[T common.ObjectInterface](
	ctx TxCtxInterface,
	obj T,
	txId string,
) (hiddenList *authpb.HiddenTxList, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	var (
		key    = lo.Must(MakeHiddenKey(obj))
		hidden = lo.Must(hiddenTxs(ctx, obj))
		op     = &authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_HIDE_TX,
			CollectionId: obj.ObjectKey().GetCollectionId(),
			Namespace:    obj.Namespace(),
			Paths:        nil,
		}
		authorized = lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	)

	if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	found := false
	for i, tx := range hidden.Txs {
		if tx.TxId == txId {
			hidden.Txs = append(hidden.Txs[:i], hidden.Txs[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return nil, oops.Wrap(common.KeyNotFound)
	}

	bytes := lo.Must(json.Marshal(hidden))

	return hidden, ctx.GetStub().PutState(key, bytes)
}
