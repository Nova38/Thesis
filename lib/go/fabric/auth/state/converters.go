package state

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/oops"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func ProtoToItem(obj *authpb.Item) (item common.ItemInterface, err error) {
	if obj == nil || obj.GetValue() == nil {
		return nil, oops.In("GetItem").Errorf("Item is nil")
	}

	m, err := obj.GetValue().UnmarshalNew()
	if err != nil {
		return nil, err
	}
	item, ok := m.(common.ItemInterface)

	if !ok {
		return nil, oops.In("GetItem").Errorf("Item is not a state.Item")
	}

	return item, nil
}

// ItemKeyToItemType Does not populate the item's key
func ItemKeyToItemType(key *authpb.ItemKey) (item common.ItemInterface, err error) {
	if key == nil {
		return nil, oops.In("GetItem").Errorf("ItemKey is nil")
	}

	name := protoreflect.FullName(key.GetItemType())

	t, err := protoregistry.GlobalTypes.FindMessageByName(name)
	if err != nil {
		return nil, err
	}

	item, ok := t.New().Interface().(common.ItemInterface)

	if !ok {
		return nil, oops.In("GetItem").Errorf("Item is not a state.Item")
	}

	return item, nil
}

func ItemToProto(item common.ItemInterface) (obj *authpb.Item, err error) {
	if obj == nil {
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

func ListItemToProtos(list []common.ItemInterface) (objs []*authpb.Item, err error) {
	for _, o := range list {
		var msg *authpb.Item

		if msg, err = ItemToProto(o); err != nil {
			return nil, err
		}
		objs = append(objs, msg)
	}

	return objs, nil
}

// ──────────────────────────────────────────────────

func ItemToSuggestion(obj common.ItemInterface) (suggestion *authpb.Suggestion, err error) {
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

func SuggestionToItem(s *authpb.Suggestion) (obj common.ItemInterface, err error) {
	if s == nil {
		return nil, oops.In("GetItem").Errorf("Item is nil")
	}

	m, err := s.GetValue().UnmarshalNew()
	if err != nil {
		return nil, err
	}
	item, ok := m.(common.ItemInterface)

	if !ok {
		return nil, oops.In("GetItem").Errorf("Item is not a state.Item")
	}

	return item, nil
}
