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

func ProtoToObject(obj *authpb.Object) (object common.ObjectInterface, err error) {
	if obj == nil || obj.GetValue() == nil {
		return nil, oops.In("GetObject").Errorf("Object is nil")
	}

	m, err := obj.GetValue().UnmarshalNew()
	if err != nil {
		return nil, err
	}
	object, ok := m.(common.ObjectInterface)

	if !ok {
		return nil, oops.In("GetObject").Errorf("Object is not a state.Object")
	}

	return object, nil
}

// Does not populate the object's key
func ObjectKeyToObjectType(key *authpb.ObjectKey) (object common.ObjectInterface, err error) {
	if key == nil {
		return nil, oops.In("GetObject").Errorf("ObjectKey is nil")
	}

	name := protoreflect.FullName(key.GetObjectType())

	t, err := protoregistry.GlobalTypes.FindMessageByName(name)
	if err != nil {
		return nil, err
	}

	object, ok := t.New().Interface().(common.ObjectInterface)

	if !ok {
		return nil, oops.In("GetObject").Errorf("Object is not a state.Object")
	}

	return object, nil
}

func ObjectToProto(object common.ObjectInterface) (obj *authpb.Object, err error) {
	if obj == nil {
		return nil, oops.In("GetObject").Errorf("Object is nil")
	}

	msg, err := anypb.New(object)
	if err != nil {
		return nil, err
	}

	key := object.ObjectKey()
	if key == nil {
		return nil, oops.In("GetObject").Errorf("ObjectKey is nil")
	}

	obj = &authpb.Object{
		Key:   key,
		Value: msg,
	}

	return obj, nil
}

func ListObjectToProtos(list []common.ObjectInterface) (objs []*authpb.Object, err error) {
	for _, o := range list {
		var msg *authpb.Object

		if msg, err = ObjectToProto(o); err != nil {
			return nil, err
		}
		objs = append(objs, msg)
	}

	return objs, nil
}

// ──────────────────────────────────────────────────

func ObjectToSuggestion(obj common.ObjectInterface) (suggestion *authpb.Suggestion, err error) {
	if suggestion == nil {
		return nil, oops.In("GetObject").Errorf("Object is nil")
	}

	msg, err := anypb.New(obj)
	if err != nil {
		return nil, err
	}

	primaryKey := obj.ObjectKey()
	if primaryKey == nil {
		return nil, oops.In("GetObject").Errorf("ObjectKey is nil")
	}

	suggestion = &authpb.Suggestion{
		PrimaryKey: primaryKey,
		Paths:      &fieldmaskpb.FieldMask{},
		Value:      msg,
	}

	return suggestion, nil
}

func SuggestionToObject(s *authpb.Suggestion) (obj common.ObjectInterface, err error) {
	if s == nil {
		return nil, oops.In("GetObject").Errorf("Object is nil")
	}

	m, err := s.GetValue().UnmarshalNew()
	if err != nil {
		return nil, err
	}
	object, ok := m.(common.ObjectInterface)

	if !ok {
		return nil, oops.In("GetObject").Errorf("Object is not a state.Object")
	}

	return object, nil
}
