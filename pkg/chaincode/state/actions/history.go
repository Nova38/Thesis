package actions

import (
	"encoding/json"
	"slices"

	authpb "github.com/nova38/saacs/lib/saacs-protos-go/auth/v1"
	"github.com/nova38/saacs/pkg/chaincode/common"
	"github.com/nova38/saacs/pkg/chaincode/state"
	"github.com/samber/lo"
	"github.com/samber/oops"

	"google.golang.org/protobuf/types/known/anypb"
)

// ════════════════════════════════════════════════════════
// History Functions
// ════════════════════════════════════════════════════════

func GetHiddenTxs[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
) (list *authpb.HiddenTxList, key string, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if !ctx.EnabledHidden() {
		return nil, "", nil
	}

	list = &authpb.HiddenTxList{
		PrimaryKey: &authpb.ItemKey{
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     common.HiddenItemType,
			ItemKind:     authpb.ItemKind_ITEM_KIND_SUB_ITEM,
			ItemKeyParts: []string{
				obj.ItemKey().GetItemType(),
			},
		},
		Txs: []*authpb.HiddenTx{},
	}
	list.PrimaryKey.ItemKeyParts = append(
		list.GetPrimaryKey().GetItemKeyParts(),
		obj.ItemKey().GetItemKeyParts()...)

	if err := state.GetFromKey(ctx, list.StateKey(), list); err != nil {
		return nil, key, oops.Wrap(err)
	}

	return list, list.StateKey(), nil
}

func GetHistory[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
) (h *authpb.History, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	authorized, err := ctx.Authorize([]*authpb.Operation{
		{
			Action:       authpb.Action_ACTION_VIEW_HISTORY,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		},
	})
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	return getHistory(ctx, obj, false)
}

func FullHistory[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
) (h *authpb.History, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	authorized, err := ctx.Authorize([]*authpb.Operation{
		{
			Action:       authpb.Action_ACTION_VIEW_HISTORY,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		},
		{
			Action:       authpb.Action_ACTION_VIEW_HIDDEN_TXS,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		},
	})

	if err != nil {
		return nil, oops.Wrap(err)
	}
	if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	return getHistory(ctx, obj, true)
}

func GetHiddenTx[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
) (l *authpb.HiddenTxList, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	authorized, err := ctx.Authorize([]*authpb.Operation{
		{
			Action:       authpb.Action_ACTION_VIEW_HIDDEN_TXS,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		}})
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	l, _, err = GetHiddenTxs(ctx, obj)

	return l, err
}

func HideTransaction[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
	tx *authpb.HiddenTx,
) (hiddenList *authpb.HiddenTxList, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	authorized, err := ctx.Authorize([]*authpb.Operation{
		{
			Action:       authpb.Action_ACTION_HIDE_TX,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		}})

	switch {
	case err != nil:
		return nil, oops.Wrap(err)
	case !authorized:
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	hidden, key, err := GetHiddenTxs(ctx, obj)

	switch {
	case err != nil:
		return nil, oops.Wrap(err)
	case hidden == nil:
		hidden = &authpb.HiddenTxList{
			Txs: []*authpb.HiddenTx{},
		}
	case hidden.GetTxs() == nil:
		hidden.Txs = []*authpb.HiddenTx{}

	case slices.ContainsFunc(hidden.GetTxs(),
		func(e *authpb.HiddenTx) bool { return e.GetTxId() == tx.GetTxId() }):
		return nil, oops.Wrap(common.AlreadyExists)
	}

	hidden.Txs = append(hidden.GetTxs(), tx)

	bytes, err := json.Marshal(hidden)
	if err != nil {
		return nil, oops.Wrap(err)
	}

	return hidden, ctx.GetStub().PutState(key, bytes)
}

func UnHideTransaction[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
	txId string,
) (hiddenList *authpb.HiddenTxList, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	authorized, err := ctx.Authorize([]*authpb.Operation{
		{
			Action:       authpb.Action_ACTION_HIDE_TX,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		},
	})

	if err != nil {
		return nil, oops.Wrap(err)
	}
	if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	hidden, key, err := GetHiddenTxs(ctx, obj)
	if err != nil {
		return nil, oops.Wrap(err)
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

	bytes, err := json.Marshal(hidden)

	if err != nil {
		return nil, oops.Wrap(err)
	}

	return hidden, ctx.GetStub().PutState(key, bytes)
}

func getHistory[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
	showHidden bool,
) (history *authpb.History, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	key := obj.StateKey()

	history = &authpb.History{
		Entries: []*authpb.HistoryEntry{},
	}

	if ctx.EnabledHidden() && showHidden {
		history.HiddenTxs, _, err = GetHiddenTxs(ctx, obj)
		if err != nil {
			return nil, err
		}
	}

	resultIterator, err := ctx.GetStub().GetHistoryForKey(key)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	defer ctx.CloseQueryIterator(resultIterator)

	for resultIterator.HasNext() {
		queryResponse, err := resultIterator.Next()
		if err != nil {
			return nil, oops.Wrap(err)
		}
		entry := &authpb.HistoryEntry{
			TxId:      queryResponse.GetTxId(),
			IsDelete:  queryResponse.GetIsDelete(),
			Timestamp: queryResponse.GetTimestamp(),
			Note:      "",
		}

		if ctx.EnabledHidden() && history.GetHiddenTxs() != nil &&
			history.GetHiddenTxs().GetTxs() != nil {

			ctx.GetLogger().Info("Hidden History Enabled, checking for hidden txs")
			if history.GetHiddenTxs().GetTxs() != nil {
				for _, tx := range history.GetHiddenTxs().GetTxs() {
					if tx.GetTxId() == entry.GetTxId() {
						entry.Note = tx.GetNote()
						entry.IsHidden = true
					}
				}
			}

			if !showHidden && entry.GetIsHidden() {
				ctx.GetLogger().Info("Hidden History Enabled but not shown, skipping hidden tx")
				continue
			}

		}
		tmp, ok := obj.ProtoReflect().New().Interface().(T)
		if !ok {
			return nil, oops.Errorf("Error cloning object")
		}

		if queryResponse.GetValue() == nil {
			break
		}

		if err := json.Unmarshal(queryResponse.GetValue(), tmp); err != nil {
			return nil, oops.Wrap(err)
		}

		entry.Value, err = anypb.New(tmp)
		if err != nil {
			return nil, oops.Wrap(err)
		}

		history.Entries = append(history.GetEntries(), entry)
	}

	return history, err
}
