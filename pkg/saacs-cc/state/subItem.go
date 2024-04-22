package state

import (
	"encoding/json"

	"github.com/nova38/saacs/apps/saacs-cc/common"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
)

func (l Ledger[T]) SubItemGet(ctx common.TxCtxInterface, obj T) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err := GetFromKey(ctx, obj.StateKey(), obj); err != nil {
		return oops.With(
			"Key", obj.StateKey(),
			"ItemType", obj.ItemType(),
		).Wrap(err)
	}

	return nil
}

func (l Ledger[T]) SubItemPut(ctx common.TxCtxInterface, obj T) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err := Put(ctx, obj); err != nil {
		return oops.With(
			"Key", obj.StateKey(),
			"ItemType", obj.ItemType(),
		).Wrap(err)
	}

	return nil
}

func (l Ledger[T]) SubItemCreate(ctx common.TxCtxInterface, obj T) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	key := obj.StateKey()

	if KeyExists(ctx, key) {
		return oops.
			With("Key", key, "ItemType", obj.ItemType()).
			Wrap(common.AlreadyExists)
	}

	if bytes, err := json.Marshal(obj); err != nil {
		return oops.Hint("Failed To Marshal").Wrap(err)
	} else {
		if err := ctx.GetStub().PutState(key, bytes); err != nil {
			return oops.With("key", key).Wrap(err)
		}
	}

	return nil
}

func UnmarshalSubItem[T common.ItemInterface](bytes []byte, obj T) (err error) {
	return json.Unmarshal(bytes, obj)
}
func UnmarshalNewSubItem[T common.ItemInterface](bytes []byte, base T) (item T, err error) {
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
