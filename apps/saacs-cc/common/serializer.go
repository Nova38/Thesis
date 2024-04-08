package common

//import (
//	"encoding/json"
//
//    "github.com/samber/oops"
//    "google.golang.org/protobuf/proto"
//)
//
//// Protobuf or JSON
//
//const (
//	JsonMode = iota
//	ProtoMode
//)
//
//var SerializerMode = JsonMode
//
//type Marshaller func(interface{}) ([]byte, error)
//
//func MarshalProto[T proto.Message](a T) ([]byte, error) {
//	switch SerializerMode {
//	case JsonMode:
//		return json.Marshal(a)
//	case ProtoMode:
//		return proto.MarshalOptions{Deterministic: true}.Marshal(a)
//	return nil, oops.Errorf("Unsupported marshalling serializer")
//}

//func UnMarshalProto[T proto.Message]([]byte) (T, error) {
//
//    switch SerializerMode {
//    case JsonMode:
//        return
//    case ProtoMode:
//
//    return nil, nil
//}

//func Marshal(a any) ([]byte, error) {
//	switch SerializerMode {
//	case JsonMode:
//		return json.Marshal(a)
//	case ProtoMode:
//
//		wrapper := &structpb.Struct{
//			Fields: map[string]*structpb.Value{"data": a},
//		}
//		return proto.Marshal(wrapper)
//	}
//	return nil, nil
//}
