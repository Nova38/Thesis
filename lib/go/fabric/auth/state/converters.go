package state

import (
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/types/known/anypb"
)

func AuthObjToObject(obj *authpb.Object) (object Object, err error) {
	if obj == nil || obj.GetValue() == nil {
		return nil, oops.In("GetObject").Errorf("Object is nil")
	}

	m, err := obj.GetValue().UnmarshalNew()
	if err != nil {
		return nil, err
	}
	object, ok := m.(Object)

	if !ok {
		return nil, oops.In("GetObject").Errorf("Object is not a state.Object")
	}

	return object, nil
}

func ObjectToAuthObj(object Object) (obj *authpb.Object, err error) {
	if obj == nil {
		return nil, oops.In("GetObject").Errorf("Object is nil")
	}

	msg, err := anypb.New(object)
	if err != nil {
		return nil, err
	}

	obj = &authpb.Object{
		CollectionId:  object.GetCollectionId(),
		ObjectType:    string(object.ProtoReflect().Type().Descriptor().FullName()),
		ObjectIdParts: lo.Must(object.Key()),
		Value:         msg,
	}

	return obj, nil
}

func ListObjectToAuthObjs(list []Object) (objs []*authpb.Object, err error) {
	for _, o := range list {
		var msg *authpb.Object

		if msg, err = ObjectToAuthObj(o); err != nil {
			return nil, err
		}
		objs = append(objs, msg)
	}

	return objs, nil
}

// ──────────────────────────────────────────────────

func ObjectToSuggestion(obj Object) (suggestion *authpb.Suggestion, err error) {
	if suggestion == nil {
		return nil, oops.In("GetObject").Errorf("Object is nil")
	}

	msg, err := anypb.New(obj)
	if err != nil {
		return nil, err
	}

	suggestion = &authpb.Suggestion{
		CollectionId:  obj.GetCollectionId(),
		ObjectType:    string(obj.ProtoReflect().Type().Descriptor().FullName()),
		ObjectIdParts: lo.Must(obj.Key()),
		Value:         msg,
	}

	return suggestion, nil
}

func SuggestionToObject(s *authpb.Suggestion) (obj Object, err error) {
	if s == nil {
		return nil, oops.In("GetObject").Errorf("Object is nil")
	}

	m, err := s.GetValue().UnmarshalNew()
	if err != nil {
		return nil, err
	}
	object, ok := m.(Object)

	if !ok {
		return nil, oops.In("GetObject").Errorf("Object is not a state.Object")
	}

	return object, nil
}
