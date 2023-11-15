package state

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
)

// MakeCompositeKey creates a composite key from the given attributes
func MakeCompositeKey[T Object](obj T) (key string, err error) {
	namespace := obj.Namespace()
	attr, err := obj.Key()
	if err != nil {
		return "", err
	}
	return shim.CreateCompositeKey(namespace, attr)
}

func MakeHiddenKey[T Object](obj T) (hiddenKey string, err error) {
	attr, err := obj.Key()
	if err != nil {
		return "", err
	}
	attr = append([]string{obj.Namespace()}, attr...)

	return shim.CreateCompositeKey(common.HiddenNamespace, attr)
}

func MakeSuggestionKey[T Object](
	obj T,
	suggestionId string,
) (suggestionKey string, err error) {
	attr, err := obj.Key()
	if err != nil {
		return "", err
	}
	attr = append([]string{obj.Namespace()}, attr...)
	if suggestionId != "" {
		attr = append(attr, suggestionId)
	}

	return shim.CreateCompositeKey(common.SuggestionNamespace, attr)
}
