package common

import (
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
)

// ──────────────────────────────── Item Utils ──────────────────────────────────────────
func CloneItemWithKey[T ItemInterface](item T) (clone T, err error) {
	clone, ok := proto.Clone(item).(T)
	if !ok {
		return clone, oops.
			With("Base Object", item).
			Errorf("Error cloning object")
	}

	proto.Reset(clone)

	clone.SetKey(item.ItemKey())

	return clone, nil
}
