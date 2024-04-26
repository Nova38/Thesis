package common

import (
	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"

	"github.com/samber/oops"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ──────────────────────────────── Item Utils ──────────────────────────────────────────

// ItemKeyToItemType Does not populate the item's key
func ItemKeyToItemType(key *pb.ItemKey) (item ItemInterface, err error) {
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
		return nil, oops.In("KeyToItem").
			Hint("Failed to match the Item Interface").
			Wrap(ItemInvalid)
	}

	return item, nil
}

// ItemKeyToItem creates the item of the keys type and populates the item's key
func ItemKeyToItem(key *pb.ItemKey) (item ItemInterface, err error) {
	item, err = ItemKeyToItemType(key)
	if err != nil {
		return nil, oops.Wrap(err)
	}

	if key.GetItemKeyParts() == nil {
		return nil, oops.
			In("Converter").
			With("key", key).
			Code(pb.TxError_INVALID_ITEM_FIELD_VALUE.String()).
			Wrap(InvalidItemFieldValue)
	}

	item.SetKey(key)

	return item, nil
}

// ──────────────────────────────── Packing ──────────────────────────────────────────

func UnPackItem(obj *pb.Item) (item ItemInterface, err error) {
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

func PackItem(item ItemInterface) (obj *pb.Item, err error) {
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

	obj = &pb.Item{
		Value: msg,
	}

	return obj, nil
}

func ListItemToProtos(list []ItemInterface) (objs []*pb.Item, err error) {
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

func ItemToSuggestion(obj ItemInterface) (suggestion *pb.Suggestion, err error) {
	if obj == nil {
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

	suggestion = &pb.Suggestion{
		PrimaryKey: primaryKey,
		Paths:      &fieldmaskpb.FieldMask{},
		Value:      msg,
	}

	return suggestion, nil
}

func SuggestionToItem(s *pb.Suggestion) (obj ItemInterface, err error) {
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

func MaybePagation(p *pb.Pagination) (pagination *pb.Pagination) {
	if p != nil {
		return p
	}

	return &pb.Pagination{
		PageSize: DefaultPageSize,
		Bookmark: "",
	}
}
