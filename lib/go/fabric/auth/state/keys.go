package state

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/oops"
)

func KeyExists(ctx TxCtxInterface, key string) bool {
	bytes, err := ctx.GetStub().GetState(key)
	if bytes == nil && err == nil {
		return false
	}

	return err == nil
}

// ─────────────────────────────────────────────────────────────────────────────────
// ─────────────────────────────────────────────────────────────────────────────────

// MakeItemKeyAttr creates a composite key from the given attributes
func MakeItemKeyAttr(key *authpb.ItemKey) []string {
	return append(
		[]string{key.GetCollectionId()},
		key.GetItemIdParts()...,
	)
}

// MakeItemKeyPrimary
func MakeItemKeyPrimary(key *authpb.ItemKey) (itemKey string, err error) {
	if key == nil {
		return "", oops.Errorf("Invalid item key")
	}

	return shim.CreateCompositeKey(
		key.GetItemType(),
		MakeItemKeyAttr(key),
	)
}

// ─────────────────────────────────────────────────────────────────────────────────

// MakePrimaryKeyAttr creates a composite key from the given attributes
func MakePrimaryKeyAttr[T common.ItemInterface](obj T) (attr []string) {
	return append(
		[]string{obj.ItemKey().GetCollectionId()},
		obj.ItemKey().GetItemIdParts()...,
	)
}

// MakePrimaryKey creates a composite key from the given attributes
func MakePrimaryKey[T common.ItemInterface](obj T) (key string, err error) {
	return shim.CreateCompositeKey(
		obj.ItemKey().GetItemType(),
		MakePrimaryKeyAttr(obj),
	)
}

// ─────────────────────────────────────────────────────────────────────────────────

func MakeSubKeyAtter[T common.ItemInterface](obj T) (attr []string) {
	return append(
		[]string{
			obj.ItemKey().GetCollectionId(),
			obj.ItemKey().GetItemType(),
		},
		obj.ItemKey().GetItemIdParts()...,
	)
}

func MakeSubItemKeyAtter(key *authpb.ItemKey) (attr []string) {
	return append(
		[]string{key.GetCollectionId(), key.GetItemType()},
		key.GetItemIdParts()...,
	)
}

// ─────────────────────────────────────────────────────────────────────────────────

// func MakeHiddenKeyAtter[T common.ItemInterface](obj T) (attr []string) {
// 	return append([]string{common.HiddenItemType}, ...)
// }

func MakeHiddenKey[T common.ItemInterface](obj T) (hiddenKey string, err error) {
	return shim.CreateCompositeKey(
		common.HiddenItemType,
		MakeSubKeyAtter(obj),
	)
}

// ─────────────────────────────────────────────────────────────────────────────────

func MakeSuggestionKeyAtter[T common.ItemInterface](
	obj T,
	suggestionId string,
) (attr []string) {
	return append(MakeSubKeyAtter(obj), suggestionId)
}

// Key should be {<SUGGESTION>}{COLLECTION_ID}{ITEM_TYPE}{...ITEM_ID}{SuggestionId}
func MakeSuggestionKey[T common.ItemInterface](
	obj T,
	suggestionId string,
) (suggestionKey string, err error) {
	return shim.CreateCompositeKey(
		common.SuggestionItemType,
		MakeSuggestionKeyAtter(obj, suggestionId),
	)
}

func MakeItemKeySuggestion(
	objKey *authpb.ItemKey,
	suggestionId string,
) (suggestionKey string, err error) {
	return shim.CreateCompositeKey(
		common.SuggestionItemType,
		append(MakeSubItemKeyAtter(objKey), suggestionId),
	)
}

// ─────────────────────────────────────────────────────────────────────────────────
func MakeRefKeyAttrs(
	ref *authpb.ReferenceKey,
) (refKey1 []string, refKey2 []string, err error) {
	if ref == nil || (ref.GetKey1() == nil && ref.GetKey2() == nil) {
		return refKey1, refKey2, oops.Errorf("Invalid reference")
	}

	var a, b []string

	if ref.GetKey1() != nil {
		// a = append([]string{ref.Key_1.GetCollectionId(), ref.GetKey_1().GetItemType()}, ref.GetKey_1().GetItemIdParts()...)
		a = MakeSubItemKeyAtter(ref.GetKey1())
	}
	if ref.GetKey2() != nil {
		// b = append([]string{ref.GetKey_2().GetItemType()}, ref.GetKey_2().GetItemIdParts()...)
		b = MakeSubItemKeyAtter(ref.GetKey2())
	}
	refKey1 = append(a, b...)
	refKey2 = append(b, a...)

	return refKey1, refKey2, nil
}

func MakeRefKeys(
	ref *authpb.ReferenceKey,
) (refKey1 string, refKey2 string, err error) {
	// attr := obj.KeyAttr()
	// ItemKey := obj.ItemKey()

	if ref == nil || (ref.GetKey1() == nil && ref.GetKey2() == nil) {
		return "", "", oops.Errorf("Invalid reference")
	}

	var a, b, k1, k2 []string

	if ref.GetKey1() != nil {
		// a = append([]string{ref.Key_1.GetCollectionId(), ref.GetKey_1().GetItemType()}, ref.GetKey_1().GetItemIdParts()...)
		a = MakeSubItemKeyAtter(ref.GetKey1())
	}
	if ref.GetKey2() != nil {
		// b = append([]string{ref.GetKey_2().GetItemType()}, ref.GetKey_2().GetItemIdParts()...)
		b = MakeSubItemKeyAtter(ref.GetKey2())
	}

	switch {
	case ref.GetKey1() != nil && ref.GetKey2() != nil:
		{
			k1 = append(a, b...)
			k2 = append(b, a...)

			refKey1, err = shim.CreateCompositeKey(common.ReferenceItemType, k1)
			if err != nil {
				return "", "", err
			}

			refKey2, err = shim.CreateCompositeKey(common.ReferenceItemType, k2)
			if err != nil {
				return "", "", err
			}

			return refKey1, refKey2, nil
		}
	case ref.GetKey1() != nil && ref.GetKey2() == nil:
		{
			refKey1, err = shim.CreateCompositeKey(common.ReferenceItemType, a)
			if err != nil {
				return "", "", err
			}
			return refKey1, "", nil

		}
	case ref.GetKey1() == nil && ref.GetKey2() != nil:
		{
			refKey2, err = shim.CreateCompositeKey(common.ReferenceItemType, b)
			if err != nil {
				return "", "", err
			}

			return "", refKey2, nil
		}
	default:
		return "", "", oops.Errorf("Invalid reference")
	}
}
