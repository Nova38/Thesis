package state

import (
	"log/slog"

	"github.com/nova38/thesis/lib/go/fabric/auth/common"
)

// -
// Keys

// MakeCompositeKey creates a composite key from the given attributes
func MakeCompositeKey[T Object](ctx TxCtxInterface, obj T) (key string, err error) {
	namespace := obj.Namespace()
	attr, err := obj.Key()
	if err != nil {
		return "", err
	}

	ctx.GetLogger().
		Info("MakeCompositeKey",
			slog.Group("Key", "Namespace", namespace, "attr", attr))

	return ctx.GetStub().CreateCompositeKey(namespace, attr)
}

func MakeHiddenKey[T Object](ctx TxCtxInterface, obj T) (hiddenKey string, err error) {
	attr, err := obj.Key()
	if err != nil {
		return "", err
	}
	attr = append([]string{obj.Namespace()}, attr...)

	return ctx.GetStub().CreateCompositeKey(common.HiddenNamespace, attr)
}

func MakeSuggestionKey[T Object](
	ctx TxCtxInterface,
	obj T,
	suggestionId string,
) (suggestionKey string, err error) {
	attr, err := obj.Key()
	if err != nil {
		return "", err
	}
	attr = append([]string{obj.Namespace()}, attr...)
	attr = append(attr, suggestionId)

	return ctx.GetStub().CreateCompositeKey(common.SuggestionNamespace, attr)
}
