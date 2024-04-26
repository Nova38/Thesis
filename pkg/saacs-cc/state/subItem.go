package state

import (
	"github.com/nova38/saacs/pkg/saacs-cc/common"
	"github.com/nova38/saacs/pkg/saacs-cc/serializer"
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

	if bytes, err := serializer.Marshal(obj); err != nil {
		return oops.Hint("Failed To Marshal").Wrap(err)
	} else {
		if err := ctx.GetStub().PutState(key, bytes); err != nil {
			return oops.With("key", key).Wrap(err)
		}
	}

	return nil
}

func UnmarshalSubItem[T common.ItemInterface](bytes []byte, obj T) (err error) {
	return serializer.Unmarshal(bytes, obj)
}
func UnmarshalNewSubItem[T common.ItemInterface](bytes []byte, base T) (item T, err error) {
	item, ok := proto.Clone(base).(T)
	if !ok {
		return item, oops.Errorf("Failed to clone")
	}
	proto.Reset(item)

	if err = serializer.Unmarshal(bytes, item); err != nil {
		return item, oops.Wrap(err)
	}
	return item, nil
}
