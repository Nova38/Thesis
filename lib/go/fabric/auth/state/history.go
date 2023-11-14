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

func hiddenTxs[T Object](ctx TxCtxInterface, obj T) (list *authpb.HiddenTxList, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key    = lo.Must(MakeHiddenKey(ctx, obj))
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

func history[T Object](
	ctx TxCtxInterface,
	obj T,
	showHidden bool,
) (history *authpb.History, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key    = lo.Must(MakeCompositeKey(ctx, obj))
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
