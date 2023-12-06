package state

import (
	"encoding/json"

	"github.com/nova38/thesis/packages/saacs/common"
	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
)

func (l Ledger[T]) PrimaryGet(ctx common.TxCtxInterface, obj T) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err := Get(ctx, obj.StateKey(), obj); err != nil {
		return oops.With(
			"Key", obj.StateKey(),
			"ItemType", obj.ItemType(),
		).Wrap(err)
	}

	return nil
}

func (l Ledger[T]) PrimaryPut(ctx common.TxCtxInterface, obj T) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if obj.ItemKind() != authpb.ItemKind_ITEM_KIND_PRIMARY_ITEM {
		return oops.With(
			"Key", obj.StateKey(),
			"ItemType", obj.ItemType(),
		).Wrap(common.Runtime)
	}

	if err := Put(ctx, obj); err != nil {
		return oops.With(
			"Key", obj.StateKey(),
			"ItemType", obj.ItemType(),
		).Wrap(err)
	}

	return nil
}

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

func UnmarshalNewPrimary[T common.ItemInterface](bytes []byte, base T) (item T, err error) {

	if base.ItemKind() != authpb.ItemKind_ITEM_KIND_PRIMARY_ITEM {
		return base, oops.With(
			"Key", base.StateKey(),
			"ItemType", base.ItemType(),
		).Wrap(common.Runtime)
	}

	item, ok := proto.Clone(base).(T)
	if !ok {
		return item, oops.Errorf("Failed to clone")
	}
	proto.Reset(item)

	if err = json.Unmarshal(bytes, item); err != nil {
		return item, oops.Wrap(err)
	}
	return item, nil
}
