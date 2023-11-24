package main

import (
	"fmt"
	"testing"

	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	// "github.com/bufbuild/protovalidate-go"
	// pb "github.com/nova38/thesis/lib/go/gen/rbac"
	_ "github.com/nova38/thesis/lib/go/gen/auth/v1"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func TestV(t *testing.T) {
}

func main() {
	name := protoreflect.FullName("auth.Suggestion")

	t, err := protoregistry.GlobalTypes.FindMessageByName(name)
	if err != nil {
		fmt.Println(err)
	}

	// item := t.New().Interface()

	item, ok := t.New().Interface().(common.ItemInterface)
	if !ok {
		fmt.Println("not ok")
	}

	item.ItemKey()

	fd := item.ProtoReflect().Descriptor().Fields().ByName("suggestion_id")
	fmt.Println(fd.Name())
	fmt.Println(fd.TextName())
	fmt.Println(fd.FullName())
	// fmt.Println(fd.())
	item.ProtoReflect().Set(fd, protoreflect.ValueOf("test"))

	// Modifiy suggestion_id field

	fmt.Println(item)
}
