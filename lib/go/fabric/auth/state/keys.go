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
func MakeObjectKeyAttr(key *authpb.ObjectKey) []string {
	return append(
		[]string{key.GetCollectionId()},
		key.GetObjectIdParts()...,
	)
}

func MakeObjectKeyPrimary(key *authpb.ObjectKey) (objectKey string, err error) {
	return shim.CreateCompositeKey(
		key.GetObjectType(),
		MakeObjectKeyAttr(key),
	)
}

// ─────────────────────────────────────────────────────────────────────────────────

// MakePrimaryKeyAttr creates a composite key from the given attributes
func MakePrimaryKeyAttr[T common.ObjectInterface](obj T) (attr []string) {
	return append(
		[]string{obj.ObjectKey().GetCollectionId()},
		obj.ObjectKey().GetObjectIdParts()...,
	)
}

// MakePrimaryKey creates a composite key from the given attributes
func MakePrimaryKey[T common.ObjectInterface](obj T) (key string, err error) {
	return shim.CreateCompositeKey(
		obj.ObjectKey().GetObjectType(),
		MakePrimaryKeyAttr(obj),
	)
}

// ─────────────────────────────────────────────────────────────────────────────────

func MakeSubKeyAtter[T common.ObjectInterface](obj T) (attr []string) {
	return append(
		[]string{
			obj.ObjectKey().GetCollectionId(),
			obj.ObjectKey().GetObjectType(),
		},
		obj.ObjectKey().GetObjectIdParts()...,
	)
}

func MakeSubObjectKeyAtter(key *authpb.ObjectKey) (attr []string) {
	return append(
		[]string{key.GetCollectionId(), key.GetObjectType()},
		key.GetObjectIdParts()...,
	)
}

// ─────────────────────────────────────────────────────────────────────────────────

// func MakeHiddenKeyAtter[T common.ObjectInterface](obj T) (attr []string) {
// 	return append([]string{common.HiddenNamespace}, ...)
// }

func MakeHiddenKey[T common.ObjectInterface](obj T) (hiddenKey string, err error) {
	return shim.CreateCompositeKey(
		common.HiddenNamespace,
		MakeSubKeyAtter(obj),
	)
}

// ─────────────────────────────────────────────────────────────────────────────────

func MakeSuggestionKeyAtter[T common.ObjectInterface](
	obj T,
	suggestionId string,
) (attr []string) {
	return append(MakeSubKeyAtter(obj), suggestionId)
}

// Key should be {SUGGESTION}{COLLECTION_ID}{OBJECT_TYPE}{...OBJECT_ID}{SuggestionId}
func MakeSuggestionKey[T common.ObjectInterface](
	obj T,
	suggestionId string,
) (suggestionKey string, err error) {
	return shim.CreateCompositeKey(
		common.SuggestionNamespace,
		MakeSuggestionKeyAtter(obj, suggestionId),
	)
}

func MakeObjectKeySuggestion(
	objKey *authpb.ObjectKey,
	suggestionId string,
) (suggestionKey string, err error) {
	return shim.CreateCompositeKey(
		common.SuggestionNamespace,
		append(MakeSubObjectKeyAtter(objKey), suggestionId),
	)
}

// ─────────────────────────────────────────────────────────────────────────────────

func MakeRefKeys(
	ref *authpb.Reference,
) (refKey_1 string, refKey_2 string, err error) {
	// attr := obj.KeyAttr()
	// ObjectKey := obj.ObjectKey()

	if ref == nil || (ref.GetKey_1() == nil && ref.GetKey_2() == nil) {
		return "", "", oops.Errorf("Invalid reference")
	}

	var a, b, k1, k2 []string

	refBase := []string{
		common.ReferenceNamespace,
		ref.GetReferenceType(),
	}

	if ref.GetKey_1() != nil {
		// a = append([]string{ref.Key_1.GetCollectionId(), ref.GetKey_1().GetObjectType()}, ref.GetKey_1().GetObjectIdParts()...)
		a = MakeSubObjectKeyAtter(ref.GetKey_1())
	}
	if ref.GetKey_2() != nil {
		// b = append([]string{ref.GetKey_2().GetObjectType()}, ref.GetKey_2().GetObjectIdParts()...)
		b = MakeSubObjectKeyAtter(ref.GetKey_2())
	}

	if ref.GetKey_1() != nil && ref.GetKey_2() != nil {
		k1 = append(a, b...)
		k2 = append(b, a...)

		refKey_1, err = shim.CreateCompositeKey(common.ReferenceNamespace, append(refBase, k1...))
		if err != nil {
			return "", "", err
		}

		refKey_2, err = shim.CreateCompositeKey(common.ReferenceNamespace, append(refBase, k2...))
		if err != nil {
			return "", "", err
		}
	} else if ref.GetKey_1() != nil {
		refKey_1, err = shim.CreateCompositeKey(common.ReferenceNamespace, append(refBase, a...))
		if err != nil {
			return "", "", err
		}
	} else if ref.GetKey_2() != nil {
		refKey_2, err = shim.CreateCompositeKey(common.ReferenceNamespace, append(refBase, b...))
		if err != nil {
			return "", "", err
		}
	}

	return refKey_1, refKey_2, nil
}
