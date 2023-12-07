package common

import (
	"encoding/json"
	"reflect"

	"github.com/samber/lo"
)

// ──────────────────────────────── Item Utils ──────────────────────────────────────────

func NewItem[T any]() T {
	var i T
	item, ok := reflect.New(reflect.TypeOf(i).Elem()).Interface().(T)
	if !ok {
		panic("not ok")
	}
	return item
}

func UnmarshalNew[T any](bytes []byte) (item T, err error) {
	value := new(T)
	err = json.Unmarshal(bytes, value)
	return lo.FromPtr(value), err
}
