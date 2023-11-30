package common

import (
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/oops"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ItemKeyToItemType Does not populate the item's key
func ItemKeyToItemType(key *authpb.ItemKey) (item ItemInterface, err error) {
	if key == nil {
		return nil, oops.In("GetItem").Errorf("ItemKey is nil")
	}

	name := protoreflect.FullName(key.GetItemType())

	t, err := protoregistry.GlobalTypes.FindMessageByName(name)
	if err != nil {
		return nil, err
	}

	item, ok := t.New().Interface().(ItemInterface)

	if !ok {
		return nil, oops.In("GetItem").Errorf("Item is not a state.Item")
	}

	return item, nil
}

func ItemKeyToItem(key *authpb.ItemKey) (item ItemInterface, err error) {
	item, err = ItemKeyToItemType(key)
	if err != nil {
		return nil, oops.Wrap(err)
	}

	if key.GetItemIdParts() == nil {
		return nil, oops.Errorf("ItemKey is nil")
	}

	item.SetKey(key)

	return item, nil
}

// ──────────────────────────────── Packing ──────────────────────────────────────────

func UnPackItem(obj *authpb.Item) (item ItemInterface, err error) {
	if obj == nil || obj.GetValue() == nil {
		return nil, oops.In("GetItem").Errorf("Item is nil")
	}

	m, err := obj.GetValue().UnmarshalNew()
	if err != nil {
		return nil, err
	}
	item, ok := m.(ItemInterface)

	if !ok {
		return nil, oops.In("GetItem").Errorf("Item is not a state.Item")
	}

	return item, nil
}

func PackItem(item ItemInterface) (obj *authpb.Item, err error) {
	if item == nil {
		return nil, oops.In("GetItem").Errorf("Item is nil")
	}

	msg, err := anypb.New(item)
	if err != nil {
		return nil, err
	}

	key := item.ItemKey()
	if key == nil {
		return nil, oops.In("GetItem").Errorf("ItemKey is nil")
	}

	obj = &authpb.Item{
		Key:   key,
		Value: msg,
	}

	return obj, nil
}

func ListItemToProtos(list []ItemInterface) (objs []*authpb.Item, err error) {
	for _, item := range list {
		msg, err := PackItem(item)
		if err != nil {
			return nil, oops.Wrap(err)
		}
		objs = append(objs, msg)
	}
	return objs, nil
}

// ──────────────────────────────── Suggestions ──────────────────────────────────────────

func ItemToSuggestion(obj ItemInterface) (suggestion *authpb.Suggestion, err error) {
	if suggestion == nil {
		return nil, oops.In("GetItem").Errorf("Item is nil")
	}

	msg, err := anypb.New(obj)
	if err != nil {
		return nil, err
	}

	primaryKey := obj.ItemKey()
	if primaryKey == nil {
		return nil, oops.In("GetItem").Errorf("ItemKey is nil")
	}

	suggestion = &authpb.Suggestion{
		PrimaryKey: primaryKey,
		Paths:      &fieldmaskpb.FieldMask{},
		Value:      msg,
	}

	return suggestion, nil
}

func SuggestionToItem(s *authpb.Suggestion) (obj ItemInterface, err error) {
	if s == nil {
		return nil, oops.In("GetItem").Errorf("Item is nil")
	}

	m, err := s.GetValue().UnmarshalNew()
	if err != nil {
		return nil, err
	}
	item, ok := m.(ItemInterface)

	if !ok {
		return nil, oops.In("GetItem").Errorf("Item is not a state.Item")
	}

	return item, nil
}

// ──────────────────────────────── References ──────────────────────────────────────────

// ReferenceKeyToItems converts a reference key to the item object

func ReferenceKeyToItems(
	ref *authpb.ReferenceKey,
) (item1 ItemInterface, item2 ItemInterface, err error) {
	if ref == nil {
		return nil, nil, oops.In("GetItem").Errorf("Item is nil")
	}

	if ref.GetKey1() == nil && ref.GetKey2() == nil {
		return nil, nil, oops.In("GetItem").Errorf("ItemKey is nil")
	}

	if ref.GetKey1() != nil {
		// Get the items from the ledger
		item1, err = ItemKeyToItem(ref.GetKey1())
		if err != nil {
			return nil, nil, oops.Wrap(err)
		}
	}

	if ref.GetKey2() != nil {
		item2, err = ItemKeyToItem(ref.GetKey2())
		if err != nil {
			return item1, nil, oops.Wrap(err)
		}
	}

	return item1, item2, nil
}

// PackReference packs the items into anypb and returns the reference
func PackReference(
	ref *authpb.ReferenceKey,
	item1 ItemInterface,
	item2 ItemInterface,
) (reference *authpb.Reference, err error) {
	reference = &authpb.Reference{
		Reference: ref,
		Item1:     &authpb.Item{},
		Item2:     &authpb.Item{},
	}

	reference.Item1, err = PackItem(item1)
	if err != nil {
		return nil, oops.Hint("Error packing item1").Wrap(err)
	}

	reference.Item2, err = PackItem(item2)
	if err != nil {
		return nil, oops.Hint("Error packing item2").Wrap(err)
	}

	return reference, nil
}
