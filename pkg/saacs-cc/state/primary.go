package state

import (
	"github.com/nova38/saacs/pkg/saacs-cc/common"
	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"

	"github.com/samber/oops"
)

func (l Ledger[T]) PrimaryCreate(ctx common.TxCtxInterface, obj T) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if obj.ItemKind() != pb.ItemKind_ITEM_KIND_PRIMARY_ITEM {
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

	// Make Empty Suggestion Key
	if err := Put(ctx, &pb.Suggestion{PrimaryKey: obj.ItemKey()}); err != nil {
		return oops.
			With("Object Key", obj.StateKey(),
				"Suggestion Key", obj.ItemKey(),
			).
			In("Primary Object Creation").
			Hint("Failed to create Suggestion Key").Wrap(err)
	}

	// Make HiddenTxList for the primary item
	if ctx.EnabledHidden() {
		hiddenTxList := &pb.HiddenTxList{
			PrimaryKey: obj.ItemKey(),
			Txs:        []*pb.HiddenTx{},
		}

		if err := Put(ctx, hiddenTxList); err != nil {
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

	if obj.ItemKind() != pb.ItemKind_ITEM_KIND_PRIMARY_ITEM {
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

	if err := Delete(ctx, &pb.Suggestion{PrimaryKey: obj.ItemKey()}); err != nil {
		return oops.
			With("Object Key", obj.StateKey(),
				"Suggestion Key", obj.ItemKey(),
			).
			Hint("Failed to delete Suggestion Domain Key").Wrap(err)
	}

	if ctx.EnabledHidden() {
		hiddenTxList := &pb.HiddenTxList{
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
