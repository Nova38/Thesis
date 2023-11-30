package state

import (
	"encoding/json"
	"slices"

	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"

	"google.golang.org/protobuf/types/known/anypb"
)

// ════════════════════════════════════════════════════════
// History Functions
// ════════════════════════════════════════════════════════

func hiddenTxs[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
) (list *authpb.HiddenTxList, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		key    = lo.Must(common.MakeHiddenKey(obj))
		exists = common.KeyExists(ctx, key)
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

func history[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
	showHidden bool,
) (history *authpb.History, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	key, err := common.MakePrimaryKey(obj)
	if err != nil {
		return nil, oops.With("Object", obj).Wrap(err)
	}

	if !common.KeyExists(ctx, key) {
		return nil, oops.With("key", key).Wrap(common.KeyNotFound)
	}

	history = &authpb.History{
		Entries: []*authpb.HistoryEntry{},
	}

	// hidden = lo.Must(hiddenTxs(ctx, obj))
	if showHidden {
		history.HiddenTxs, err = hiddenTxs(ctx, obj)
		if err != nil {
			return nil, err
		}
	}

	resultIterator := lo.Must(ctx.GetStub().GetHistoryForKey(key))
	defer ctx.CloseQueryIterator(resultIterator)

	base := proto.Clone(obj).(T)
	proto.Reset(base)

	for resultIterator.HasNext() {
		queryResponse := lo.Must(resultIterator.Next())
		entry := &authpb.HistoryEntry{
			TxId:      queryResponse.GetTxId(),
			IsDelete:  queryResponse.GetIsDelete(),
			Timestamp: queryResponse.GetTimestamp(),
			Note:      "",
		}

		if history.GetHiddenTxs().GetTxs() != nil {
			for _, tx := range history.GetHiddenTxs().GetTxs() {
				if tx.GetTxId() == entry.GetTxId() {
					entry.Note = tx.GetNote()
					entry.IsHidden = true
				}
			}
		}

		if !showHidden && entry.GetIsHidden() {
			continue
		}

		tmp := proto.Clone(base)
		lo.Must0(json.Unmarshal(queryResponse.GetValue(), tmp))
		entry.Value = lo.Must(anypb.New(tmp))

		history.Entries = append(history.GetEntries(), entry)
	}

	return history, err
}

func History[T common.ItemInterface](ctx common.TxCtxInterface, obj T) (h *authpb.History, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_VIEW_HISTORY,
		CollectionId: obj.ItemKey().GetCollectionId(),
		ItemType:     obj.ItemType(),
		Paths:        nil,
	}

	authorized := lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	return history(ctx, obj, false)
}

func FullHistory[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
) (h *authpb.History, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	var (
		op = &authpb.Operation{
			Action:       authpb.Action_ACTION_VIEW_HISTORY,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		}
		authorized = lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	)

	if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	return history(ctx, obj, true)
}

func HiddenTx[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
) (l *authpb.HiddenTxList, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	var (
		op = &authpb.Operation{
			Action:       authpb.Action_ACTION_VIEW_HIDDEN_TXS,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		}
		authorized = lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	)

	if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	return hiddenTxs(ctx, obj)
}

func HideTransaction[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
	tx *authpb.HiddenTx,
) (hiddenList *authpb.HiddenTxList, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	var (
		key    = lo.Must(common.MakeHiddenKey(obj))
		hidden = lo.Must(hiddenTxs(ctx, obj))
		op     = &authpb.Operation{
			Action:       authpb.Action_ACTION_HIDE_TX,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		}
		authorized = lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	)
	if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	if hidden.GetTxs() == nil {
		hidden.Txs = []*authpb.HiddenTx{}
	}

	if slices.ContainsFunc(hidden.GetTxs(),
		func(e *authpb.HiddenTx) bool { return e.GetTxId() == tx.GetTxId() }) {
		return nil, oops.Wrap(common.AlreadyExists)
	}

	hidden.Txs = append(hidden.GetTxs(), tx)

	bytes := lo.Must(json.Marshal(hidden))

	return hidden, ctx.GetStub().PutState(key, bytes)
}

func UnHideTransaction[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
	txId string,
) (hiddenList *authpb.HiddenTxList, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	var (
		key    = lo.Must(common.MakeHiddenKey(obj))
		hidden = lo.Must(hiddenTxs(ctx, obj))
		op     = &authpb.Operation{
			Action:       authpb.Action_ACTION_HIDE_TX,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		}
		authorized = lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	)

	if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	found := false

	txs := lo.Filter(hidden.GetTxs(), func(e *authpb.HiddenTx, index int) bool {
		if e.GetTxId() == txId {
			found = true
			return false
		}
		return true
	})

	if !found {
		return nil, oops.Wrap(common.KeyNotFound)
	}

	hidden.Txs = txs

	bytes := lo.Must(json.Marshal(hidden))

	return hidden, ctx.GetStub().PutState(key, bytes)
}
