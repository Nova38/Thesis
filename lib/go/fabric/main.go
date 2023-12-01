package main

import (
	"testing"
	//	"google.golang.org/protobuf/reflect/protoreflect"
	//	"google.golang.org/protobuf/reflect/protoregistry"
	//	"google.golang.org/protobuf/types/known/anypb"
	//
	// )
)

func TestV(t *testing.T) {
}

func main() {
	// name := protoreflect.FullName("auth.Suggestion")

	// t, err := protoregistry.GlobalTypes.FindMessageByName(name)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // item := t.New().Interface()

	// item, ok := t.New().Interface().(common.ItemInterface)
	// if !ok {
	// 	fmt.Println("not ok")
	// }

	// item.ItemKey()

	// fd := item.ProtoReflect().Descriptor().Fields().ByName("suggestion_id")
	// fmt.Println(fd.Name())
	// fmt.Println(fd.TextName())
	// fmt.Println(fd.FullName())
	// // fmt.Println(fd.())
	// item.ProtoReflect().Set(fd, protoreflect.ValueOf("test"))

	// // Modifiy suggestion_id field
	// any, err := anypb.New(&authpb.User{})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(any)

	// i := &authpb.User{
	// 	CollectionId: "",
	// 	MspId:        "msp_id",
	// 	UserId:       "user_id",
	// 	Name:         "Name",
	// }

	// p, e := common.PackItem(i)

	// b, e := p.MarshalJSON()
	// // protojson.Marshal(any)
	// if e != nil {
	// 	fmt.Println(e)
	// }

	// fmt.Println(string(b))

	// fmt.Println(item)

	str := "{\"collectionId\":\"\",\"mspId\":\"Org1MSP\",\"name\":\"\",\"userId\":\"eDUwOTo6Q049b3JnMWFkbWluLE9VPWFkbWluLE89SHlwZXJsZWRnZXIsU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUzo6Q049ZmFicmljLWNhLXNlcnZlcixPVT1GYWJyaWMsTz1IeXBlcmxlZGdlcixTVD1Ob3J0aCBDYXJvbGluYSxDPVVT\"}"

	// bytes := []byte(str)

	// user := &authpb.User{}
	// _ = json.Unmarshal(bytes, user)

	// fmt.Println(user)

}
