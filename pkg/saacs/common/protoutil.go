package common

import (
	"encoding/json"
	"reflect"

	"github.com/mennanov/fmutils"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
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

// ──────────────────────────────── FieldMask Utils ──────────────────────────────────────────

// UpdateItem updates the object with the update object using the field mask
// The objects that are passed in should not be used after this function is called
// The returned object is a new object that is a copy of the original object with the update applied
func UpdateItem[T ItemInterface](mask *fieldmaskpb.FieldMask, src, update T) (T, error) {

	key := src.ItemKey()

	// If the mask is nil or empty, return the update applies to the whole object
	if mask == nil || len(mask.GetPaths()) == 0 {
		// Set the key to make sure the update is applied to the correct object
		// the key must not change
		update.SetKey(key)
		dst, ok := proto.Clone(update).(T)
		if !ok {
			return src, oops.Errorf("could not clone object")
		}
		return dst, nil
	}

	if !mask.IsValid(src) {
		return src, oops.
			With(
				"mask", mask,
				"object", src,
				"update", update,
			).
			Errorf("invalid mask")
	}

	// Filter the update object to only the fields specified in the mask
	fmutils.Filter(update, mask.GetPaths())

	// Clone the object so we don't mutate the original
	dst, ok := proto.Clone(src).(T)
	if !ok {
		return src, oops.Errorf("could not clone object")
	}

	// Remove any fields that are in the mask from the object
	fmutils.Prune(dst, mask.GetPaths())

	// Merge the update into the object
	proto.Merge(dst, update)

	// Set the key to make sure the update is applied to the correct object
	// the key must not change
	dst.SetKey(key)

	return dst, nil
}
