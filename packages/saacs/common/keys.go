package common

import (
	"slices"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
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
		key.GetItemKeyParts()...,
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
func MakePrimaryKeyAttr[T ItemInterface](obj T) (attr []string) {
	return append(
		[]string{obj.ItemKey().GetCollectionId()},
		obj.ItemKey().GetItemKeyParts()...,
	)
}

// MakePrimaryKey creates a composite key from the given attributes
func MakePrimaryKey[T ItemInterface](obj T) (key string, err error) {
	return shim.CreateCompositeKey(
		obj.ItemKey().GetItemType(),
		MakePrimaryKeyAttr(obj),
	)
}

// ─────────────────────────────────────────────────────────────────────────────────

func MakeSubKeyAtter[T ItemInterface](obj T) (attr []string) {
	return append(
		[]string{
			obj.ItemKey().GetCollectionId(),
			obj.ItemKey().GetItemType(),
		},
		obj.ItemKey().GetItemKeyParts()...,
	)
}

func MakeSubItemKeyAtter(key *authpb.ItemKey) (attr []string) {
	return append(
		[]string{key.GetCollectionId(), key.GetItemType()},
		key.GetItemKeyParts()...,
	)
}

// ─────────────────────────────────────────────────────────────────────────────────

// func MakeHiddenKeyAtter[T common.ItemInterface](obj T) (attr []string) {
// 	return append([]string{common.HiddenItemType}, ...)
// }

func MakeHiddenKey[T ItemInterface](obj T) (hiddenKey string, err error) {
	return shim.CreateCompositeKey(
		HiddenItemType,
		MakeSubKeyAtter(obj),
	)
}

// ─────────────────────────────────────────────────────────────────────────────────

func MakeSuggestionKeyAtter[T ItemInterface](
	obj T,
	suggestionId string,
) (attr []string) {
	return append(MakeSubKeyAtter(obj), suggestionId)
}

// Key should be {<SUGGESTION>}{COLLECTION_ID}{ITEM_TYPE}{...ITEM_ID}{SuggestionId}
func MakeSuggestionPrimaryKey[T ItemInterface](
	obj T,
	suggestionId string,
) (suggestionKey string, err error) {
	return shim.CreateCompositeKey(
		SuggestionItemType,
		MakeSuggestionKeyAtter(obj, suggestionId),
	)
}

func MakeSuggestionKey[T ItemInterface](
	suggestion *authpb.Suggestion,
) (suggestionKey string, err error) {
	return shim.CreateCompositeKey(
		SuggestionItemType,
		append(MakeSubItemKeyAtter(suggestion.GetPrimaryKey()), suggestion.GetSuggestionId()),
	)
}

func MakeItemKeySuggestion(
	objKey *authpb.ItemKey,
	suggestionId string,
) (suggestionKey string, err error) {
	return shim.CreateCompositeKey(
		SuggestionItemType,
		append(MakeSubItemKeyAtter(objKey), suggestionId),
	)
}

func MakeItemKeySuggestionKeyAttr(
	objKey *authpb.ItemKey,
	suggestionId string,
) (attr []string) {
	return append(MakeSubItemKeyAtter(objKey), suggestionId)
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
		// a = append([]string{ref.Key_1.GetCollectionId(), ref.GetKey_1().GetItemType()}, ref.GetKey_1().GetItemKeyParts()...)
		a = MakeSubItemKeyAtter(ref.GetKey1())
	}
	if ref.GetKey2() != nil {
		// b = append([]string{ref.GetKey_2().GetItemType()}, ref.GetKey_2().GetItemKeyParts()...)
		b = MakeSubItemKeyAtter(ref.GetKey2())
	}
	refKey1 = slices.Clone(a)
	refKey1 = append(refKey1, b...)

	refKey2 = slices.Clone(b)
	refKey2 = append(refKey2, a...)

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
		// a = append([]string{ref.Key_1.GetCollectionId(), ref.GetKey_1().GetItemType()}, ref.GetKey_1().GetItemKeyParts()...)
		a = MakeSubItemKeyAtter(ref.GetKey1())
	}
	if ref.GetKey2() != nil {
		// b = append([]string{ref.GetKey_2().GetItemType()}, ref.GetKey_2().GetItemKeyParts()...)
		b = MakeSubItemKeyAtter(ref.GetKey2())
	}

	switch {
	case ref.GetKey1() != nil && ref.GetKey2() != nil:
		{
			k1 = slices.Clone(a)
			k1 = append(k1, b...)
			k2 = slices.Clone(b)
			k2 = append(k2, a...)

			refKey1, err = shim.CreateCompositeKey(ReferenceItemType, k1)
			if err != nil {
				return "", "", err
			}

			refKey2, err = shim.CreateCompositeKey(ReferenceItemType, k2)
			if err != nil {
				return "", "", err
			}

			return refKey1, refKey2, nil
		}
	case ref.GetKey1() != nil && ref.GetKey2() == nil:
		{
			refKey1, err = shim.CreateCompositeKey(ReferenceItemType, a)
			if err != nil {
				return "", "", err
			}
			return refKey1, "", nil

		}
	case ref.GetKey1() == nil && ref.GetKey2() != nil:
		{
			refKey2, err = shim.CreateCompositeKey(ReferenceItemType, b)
			if err != nil {
				return "", "", err
			}

			return "", refKey2, nil
		}
	default:
		return "", "", oops.Errorf("Invalid reference")
	}
}
