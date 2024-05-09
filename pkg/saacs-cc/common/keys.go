package common

import (
	"strings"

	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"
)

const sep = string(rune(0))

// MakeComposeKey creates a composite key from the given attributes
//
// Based on the "github.com/hyperledger/fabric-saacs-cc-go/shim" package's CreateCompositeKey
// so that we don't have to import the shim package in this package
func MakeComposeKey(namespace string, attrs []string) (key string, err error) {
	// return shim.CreateCompositeKey(namespace, []string{attrs})
	key = namespace + sep

	if len(attrs) == 0 {
		return key, nil
	}
	return key + strings.Join(attrs, sep) + sep, nil

	// for _, attr := range attrs {
	// 	// TODO Validate the attribute

	// 	key = key + attr + sep
	// }
	// return key, nil

}

// ─────────────────────────────────────────────────────────────────────────────────
// ─────────────────────────────────────────────────────────────────────────────────

// MakeStateKey creates a composite key from the given attributes
// Key should be {<ITEM_TYPE>}{COLLECTION_ID}{...ITEM_ID}
// Panics if ItemType or CollectionId is nil or an empty string
func MakeStateKey(objKey *pb.ItemKey) (key string) {

	attrs := objKey.GetItemKeyParts()
	if attrs == nil {
		panic("ItemKeyParts is nil")
	}

	collectionId := objKey.GetCollectionId()
	if collectionId == "" {
		panic("CollectionId is nil")
	}

	key = sep + objKey.GetItemType() + sep + collectionId + sep

	if len(attrs) == 0 {
		return key
	}
	key = key + strings.Join(attrs, sep) + sep

	return key
}

func KeyToSubKey(objKey *pb.ItemKey, subType string) (subKey *pb.ItemKey) {
	subKey = &pb.ItemKey{
		ItemType:     subType,
		CollectionId: objKey.GetCollectionId(),
		ItemKeyParts: objKey.GetItemKeyParts(),
	}

	return subKey
}

// ─────────────────────────────────────────────────────────────────────────────────

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

func MakeSubItemKeyAtter(key *pb.ItemKey) (attr []string) {
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
	return MakeComposeKey(
		HiddenItemType,
		MakeSubKeyAtter(obj),
	)
}

func MakeItemHiddenKey[T ItemInterface](objKey *pb.ItemKey) (hiddenKey string) {
	subKey := KeyToSubKey(objKey, HiddenItemType)
	return MakeStateKey(subKey)
}

// ─────────────────────────────────────────────────────────────────────────────────
func MakeItemSuggestionDomain[T ItemInterface](objKey *pb.ItemKey) (hiddenKey string) {
	subKey := KeyToSubKey(objKey, SuggestionItemType)
	return MakeStateKey(subKey)
}

func MakeSuggestionDomainKey(s *pb.Suggestion) (suggestionKey string, err error) {
	base, err := MakeComposeKey(
		SuggestionItemType,
		MakeSubItemKeyAtter(s.GetPrimaryKey()),
	)
	if err != nil {
		return "", err
	}
	return base, nil
}

func MakeSuggestionKey[T ItemInterface](obj T) (suggestionKey string, err error) {
	return MakeComposeKey(
		SuggestionItemType,
		MakeSubKeyAtter(obj),
	)
}

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
	return MakeComposeKey(
		SuggestionItemType,
		MakeSuggestionKeyAtter(obj, suggestionId),
	)
}

func MakeItemKeySuggestion(
	objKey *pb.ItemKey,
	suggestionId string,
) (suggestionKey string) {
	subKey := KeyToSubKey(objKey, SuggestionItemType)
	subKey.ItemKeyParts = append(subKey.GetItemKeyParts(), suggestionId)

	return MakeStateKey(subKey)
}

func MakeItemKeySuggestionKeyAttr(
	objKey *pb.ItemKey,
	suggestionId string,
) (attr []string) {
	return append(MakeSubItemKeyAtter(objKey), suggestionId)
}
