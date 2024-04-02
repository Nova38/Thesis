package state

import (
	"github.com/nova38/saacs/apps/saacs-cc/common"
	authpb "github.com/nova38/saacs/libs/saacs-protos-go/auth/v1"
	"github.com/samber/oops"
)

func (l Ledger[T]) PrimaryCreate(ctx common.TxCtxInterface, obj T) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if obj.ItemKind() != authpb.ItemKind_ITEM_KIND_PRIMARY_ITEM {
		return oops.With(
			"Key", obj.StateKey(),
			"ItemType", obj.ItemType(),
		).Wrap(common.Runtime)
	}

	if err := Put[T](ctx, obj); err != nil {
		return oops.
			With("Key", obj.StateKey(),
				"ItemType", obj.ItemType(),
				"obj", obj).
			Hint("Failed to create object").Wrap(err)
	}

	if ctx.EnabledHidden() {
		hiddenTxList := &authpb.HiddenTxList{
			PrimaryKey: obj.ItemKey(),
			Txs:        []*authpb.HiddenTx{},
		}

		if err := Put[*authpb.HiddenTxList](ctx, hiddenTxList); err != nil {
			return oops.
				With("Object Key", obj.StateKey(),
					"HiddenTxList Key", hiddenTxList.StateKey(),
				).
				Hint("Failed to create hiddenTxList").Wrap(err)
		}
	}

	return nil
}

func (l Ledger[T]) PrimaryDelete(ctx common.TxCtxInterface, obj T) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if obj.ItemKind() != authpb.ItemKind_ITEM_KIND_PRIMARY_ITEM {
		return oops.With(
			"Key", obj.StateKey(),
			"ItemType", obj.ItemType(),
		).Wrap(common.Runtime)
	}

	if err := Delete(ctx, obj); err != nil {
		return oops.With(
			"Key", obj.StateKey(),
			"ItemType", obj.ItemType(),
		).Wrap(err)
	}

	if ctx.EnabledHidden() {
		hiddenTxList := &authpb.HiddenTxList{
			PrimaryKey: obj.ItemKey(),
		}

		if err := Delete(ctx, hiddenTxList); err != nil {
			return oops.
				With("Object Key", obj.StateKey(),
					"HiddenTxList Key", hiddenTxList.StateKey(),
				).
				Hint("Failed to delete hiddenTxList").Wrap(err)
		}
	}

	return nil
}
