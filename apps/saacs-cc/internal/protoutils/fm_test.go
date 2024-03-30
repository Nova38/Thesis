package protoutils_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/nova38/saacs/apps/saacs-cc/common"
	authpb "github.com/nova38/saacs/libs/saacs-protos-go/auth/v1"
	"github.com/samber/lo"
)

type TPointer[T common.ItemInterface] interface {
	*T
}

func OtherNew[T any]() (item T) {
	var tmpObj T
	item = tmpObj
	return item
}

func Unmarshal[T any](bytes []byte) (item T, err error) {
	value := new(T)
	err = json.Unmarshal(bytes, value)
	return lo.FromPtr(value), err
}

func makeI[T common.ItemInterface]() T {
	var i T
	item, ok := reflect.New(reflect.TypeOf(i).Elem()).Interface().(T)
	if !ok {
		panic("not ok")
	}
	fmt.Printf("t1: %T\n", i)
	fmt.Printf("t1: %T\n", item)

	return item
}

func makeIt[T common.ItemInterface](item T) T {

	return makeI[T]()
}

func TestInterface(t *testing.T) {
	c1 := &authpb.Collection{}
	// col := makeI[*authpb.Collection]()

	col := makeIt(c1)

	col.CollectionId = "test"
	text := lo.Must(json.MarshalIndent(&col, "", "  "))
	t.Fatal(string(text))

}

func NewItem[T common.ItemInterface]() T {
	var i T
	item, ok := reflect.New(reflect.TypeOf(i).Elem()).Interface().(T)
	if !ok {
		panic("not ok")
	}
	return item
}

func TestMakePrimaryKeyAttr(t *testing.T) {

	col := NewItem[*authpb.Collection]()

	col.CollectionId = "test"
	text := lo.Must(json.MarshalIndent(&col, "", "  "))
	t.Log(string(text))

	other, err := Unmarshal[authpb.Collection]([]byte(`{"collection_id":"other"}`))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	text = lo.Must(json.MarshalIndent(&other, "", "  "))
	t.Log(string(text))

	otherNew := OtherNew[authpb.Collection]()
	otherNew.CollectionId = "otherNew"
	text = lo.Must(json.MarshalIndent(&otherNew, "", "  "))
	t.Log(string(text))

	// t.Log(text)
	t.Fail()

}
