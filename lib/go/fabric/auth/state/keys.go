package state

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
)

// MakeCompositeKey creates a composite key from the given attributes
func MakeCompositeKey[T Object](obj T) (key string, err error) {
	col := obj.GetCollectionId()
	attr := obj.Key()

	// key = append([]string{obj.Namespace()}, attr...)
	attr = append([]string{}, attr...)

	return shim.CreateCompositeKey(col, attr)
}

func MakeHiddenKey[T Object](obj T) (hiddenKey string, err error) {
	attr := obj.Key()

	attr = append([]string{common.HiddenNamespace}, attr...)

	return shim.CreateCompositeKey(obj.GetCollectionId(), attr)
}

func MakeSuggestionKey[T Object](
	obj T,
	suggestionId string,
) (suggestionKey string, err error) {
	attr := obj.Key()

	attr = append([]string{common.SuggestionNamespace}, attr...)
	if suggestionId != "" {
		attr = append(attr, suggestionId)
	}

	return shim.CreateCompositeKey(obj.GetCollectionId(), attr)
}

func KeyExists(ctx TxCtxInterface, key string) bool {
	bytes, err := ctx.GetStub().GetState(key)
	if bytes == nil && err == nil {
		return false
	}

	return err == nil
}
