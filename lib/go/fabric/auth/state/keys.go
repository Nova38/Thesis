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
func MakeItemKeyAttr(key *authpb.ItemKey) []string {
	return append(
		[]string{key.GetCollectionId()},
		key.GetItemIdParts()...,
	)
}

func MakeItemKeyPrimary(key *authpb.ItemKey) (itemKey string, err error) {
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

// Key should be {SUGGESTION}{COLLECTION_ID}{ITEM_TYPE}{...ITEM_ID}{SuggestionId}
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

func MakeRefKeys(
	ref *authpb.Reference,
) (refKey_1 string, refKey_2 string, err error) {
	// attr := obj.KeyAttr()
	// ItemKey := obj.ItemKey()

	if ref == nil || (ref.GetKey_1() == nil && ref.GetKey_2() == nil) {
		return "", "", oops.Errorf("Invalid reference")
	}

	var a, b, k1, k2 []string

	refBase := []string{
		common.ReferenceItemType,
		ref.GetReferenceType(),
	}

	if ref.GetKey_1() != nil {
		// a = append([]string{ref.Key_1.GetCollectionId(), ref.GetKey_1().GetItemType()}, ref.GetKey_1().GetItemIdParts()...)
		a = MakeSubItemKeyAtter(ref.GetKey_1())
	}
	if ref.GetKey_2() != nil {
		// b = append([]string{ref.GetKey_2().GetItemType()}, ref.GetKey_2().GetItemIdParts()...)
		b = MakeSubItemKeyAtter(ref.GetKey_2())
	}

	if ref.GetKey_1() != nil && ref.GetKey_2() != nil {
		k1 = append(a, b...)
		k2 = append(b, a...)

		refKey_1, err = shim.CreateCompositeKey(common.ReferenceItemType, append(refBase, k1...))
		if err != nil {
			return "", "", err
		}

		refKey_2, err = shim.CreateCompositeKey(common.ReferenceItemType, append(refBase, k2...))
		if err != nil {
			return "", "", err
		}
	} else if ref.GetKey_1() != nil {
		refKey_1, err = shim.CreateCompositeKey(common.ReferenceItemType, append(refBase, a...))
		if err != nil {
			return "", "", err
		}
	} else if ref.GetKey_2() != nil {
		refKey_2, err = shim.CreateCompositeKey(common.ReferenceItemType, append(refBase, b...))
		if err != nil {
			return "", "", err
		}
	}

	return refKey_1, refKey_2, nil
}
