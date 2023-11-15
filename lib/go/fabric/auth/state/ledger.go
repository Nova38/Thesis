package state

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/nova38/thesis/lib/go/fabric/auth/common"

	"github.com/samber/oops"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

// UTIL Functions

// Exists returns true if the object exists in the ledger
func exists(ctx TxCtxInterface, key string) bool {
	bytes, err := ctx.GetStub().GetState(key)
	if bytes == nil && err == nil {
		return false
	}

	return err == nil
}

// ════════════════════════════════════════════════════════
// Primary Functions
// ════════════════════════════════════════════════════════

// Insert inserts the object into the ledger
// returns error if the object already exists
//func insert[T Object](ctx TxCtxInterface, obj T) (err error) {
//	key, err := MakeCompositeKey( obj)
//	if err != nil {
//		return err
//	}
//
//	if exists(ctx, key) {
//		return oops.
//			With("Key", key, "Namespace", obj.Namespace()).
//			Wrap(auth.AlreadyExists)
//	}
//
//	bytes, err := json.Marshal(obj)
//	if err != nil {
//		return err
//	}
//
//	return ctx.GetStub().PutState(key, bytes)
//}

//// Edit updates the object in the ledger
//// returns error if the object does not exist
//func update[T Object](ctx TxCtxInterface, obj T) (err error) {
//	key, err := MakeCompositeKey( obj)
//	if err != nil {
//		return err
//	}
//
//	if !exists(ctx, key) {
//		return oops.
//			With("Key", key, "Namespace", obj.Namespace()).
//			Wrap(auth.KeyNotFound)
//	}
//
//	bytes, err := json.Marshal(obj)
//	if err != nil {
//		return err
//	}
//
//	return ctx.GetStub().PutState(key, bytes)
//}

// Get returns the object from the ledger
func get[T Object](ctx TxCtxInterface, in T) (err error) {
	namespace := in.Namespace()
	ctx.GetLogger().Info("fn: GetState", "Namespace", namespace)

	key, err := MakeCompositeKey(in)
	if err != nil {
		return err
	}

	bytes, err := ctx.GetStub().GetState(key)
	if bytes == nil && err == nil {
		return oops.
			With("Key", key, "Namespace", in.Namespace()).
			Wrap(common.AlreadyExists)
	}

	if err = json.Unmarshal(bytes, in); err != nil {
		return err
	}
	return nil
}

// Delete deletes the object from the ledger
func delete[T Object](ctx TxCtxInterface, in T) (err error) {
	key, err := MakeCompositeKey(in)
	if err != nil {
		return err
	}

	return ctx.GetStub().DelState(key)
}

// GetHistory returns the history of the object from the ledger
func getHistory[T Object](
	ctx TxCtxInterface,
	in T,
) (list HistoryList[T], err error) {
	key, err := MakeCompositeKey(in)
	if err != nil {
		return HistoryList[T]{}, err
	}

	resultIterator, err := ctx.GetStub().GetHistoryForKey(key)
	if err != nil {
		return HistoryList[T]{}, err
	}
	defer func(resultIterator shim.HistoryQueryIteratorInterface) {
		_ = resultIterator.Close()
	}(resultIterator)

	for resultIterator.HasNext() {
		queryResponse, err := resultIterator.Next()
		if err != nil {
			return HistoryList[T]{}, err
		}
		obj := new(T)

		if err := json.Unmarshal(queryResponse.Value, &obj); err != nil {
			return HistoryList[T]{}, err
		}

		entry := &HistoryEntry[T]{
			TxId:      queryResponse.TxId,
			Timestamp: queryResponse.Timestamp,
			IsDelete:  queryResponse.IsDelete,
			State:     *obj,
		}

		list.Entries = append(list.Entries, entry)
	}

	return list, nil
}

func txIdInHistory[T Object](ctx TxCtxInterface, in T, txId string) (bool, error) {
	key, err := MakeCompositeKey(in)
	if err != nil {
		return false, err
	}

	resultIterator, err := ctx.GetStub().GetHistoryForKey(key)
	if err != nil {
		return false, err
	}
	defer func(resultIterator shim.HistoryQueryIteratorInterface) {
		_ = resultIterator.Close()
	}(resultIterator)

	for resultIterator.HasNext() {
		queryResponse, err := resultIterator.Next()
		if err != nil {
			return false, err
		}
		if queryResponse.TxId == txId {
			return true, nil
		}
	}

	return false, nil
}

// ------------------------------------------------------------

// ------------------------------------------------------------

// GetPartialKeyList returns a list of objects of type T
// T must implement StateObject interface
// numAttr is the number of attributes in the key to search for
func getPartialKeyList[T Object](
	ctx TxCtxInterface,
	in T,
	numAttr int,
) (list []T, err error) {
	// obj = []*T{}
	ctx.GetLogger().Info("GetPartialKeyList")
	namespace := in.Namespace()

	attr, err := in.Key()
	if err != nil {
		return nil, err
	}
	if len(attr) == 0 {
		return nil, fmt.Errorf("key is empty")
	}
	if len(attr) < numAttr {
		return nil, fmt.Errorf("key has less than %d attributes", numAttr)
	}

	attr = attr[:len(attr)-numAttr]

	ctx.GetLogger().
		Info("GetPartialKeyList", slog.Group("Key", "Namespace", namespace, slog.Int("numAttr", numAttr), slog.Any("attr", attr)))

	resultIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(namespace, attr)
	if err != nil {
		return nil, err
	}
	defer func(resultIterator shim.StateQueryIteratorInterface) {
		_ = resultIterator.Close()
	}(resultIterator)

	for resultIterator.HasNext() {
		queryResponse, err := resultIterator.Next()
		if err != nil {
			return nil, err
		}
		obj := new(T)

		if err := json.Unmarshal(queryResponse.Value, &obj); err != nil {
			return nil, err
		}

		list = append(list, *obj)
	}

	return list, nil
}

func getFullList[T Object](ctx TxCtxInterface, in T) (list []T, err error) {
	// obj = []*T{}

	namespace := in.Namespace()

	// key, err := ctx.GetStub().CreateCompositeKey()

	resultIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(namespace, []string{})
	if err != nil {
		return nil, err
	}
	defer func(resultIterator shim.StateQueryIteratorInterface) {
		_ = resultIterator.Close()
	}(resultIterator)

	for resultIterator.HasNext() {
		queryResponse, err := resultIterator.Next()
		if err != nil {
			return nil, err
		}
		obj := new(T)

		if err := json.Unmarshal(queryResponse.Value, &obj); err != nil {
			return nil, err
		}

		list = append(list, *obj)
	}

	return list, nil
}

// ------------------------------------------------------------
// Pagination
// ------------------------------------------------------------

// getPartialKeyList returns a list of objects of type T
// T must implement StateObject interface
// numAttr is the number of attributes in the key to search for
func getPagedPartialKeyList[T Object](
	ctx TxCtxInterface,
	in T,
	numAttr int,
	bookmark string,
) (list []T, nextBookmark string, err error) {
	// obj = []*T{}
	// ctx.GetLogger().Info("GetPagedPartialKeyList")
	namespace := in.Namespace()

	attr, err := in.Key()
	if err != nil {
		return nil, "", err
	}
	if len(attr) == 0 {
		return nil, "", fmt.Errorf("key is empty")
	}
	if len(attr) < numAttr {
		return nil, "", fmt.Errorf("key has less than %d attributes", numAttr)
	}

	attr = attr[:len(attr)-numAttr]

	ctx.GetLogger().
		Info("GetPartialKeyList",
			slog.Group(
				"Key", "Namespace", namespace,
				slog.Int("numAttr", numAttr),
				slog.Any("attr", attr),
				slog.Group(
					"Paged",
					"Bookmark", bookmark,
					"PageSize", strconv.Itoa(int(ctx.GetPageSize())),
				),
			),
		)

	resultIterator, resMeta, err := ctx.GetStub().
		GetStateByPartialCompositeKeyWithPagination(namespace, attr, ctx.GetPageSize(), bookmark)
	if err != nil {
		return nil, "", err
	}
	defer func(resultIterator shim.StateQueryIteratorInterface) {
		_ = resultIterator.Close()
	}(resultIterator)

	for resultIterator.HasNext() {
		queryResponse, err := resultIterator.Next()
		if err != nil {
			return nil, resMeta.Bookmark, err
		}
		obj := new(T)

		if err := json.Unmarshal(queryResponse.Value, &obj); err != nil {
			return nil, resMeta.Bookmark, err
		}

		list = append(list, *obj)
	}

	return list, resMeta.GetBookmark(), nil
}

// ------------------------------------------------------------
// Pagination
// ------------------------------------------------------------

// GetPartialKeyList returns a list of objects of type T
func getPagedFullList[T Object](
	ctx TxCtxInterface,
	in T,
	bookmark string,
) (list []T, nextBookmark string, err error) {
	// obj = []*T{}

	namespace := in.Namespace()

	// key, err := ctx.GetStub().CreateCompositeKey()

	ctx.GetLogger().
		Info("GetPartialKeyList",
			slog.Group("Key", "Namespace", namespace,
				slog.Group("Paged",
					"Bookmark", bookmark, "PageSize",
					strconv.Itoa(int(ctx.GetPageSize()))),
			),
		)

	resultIterator, resMeta, err := ctx.GetStub().
		GetStateByPartialCompositeKeyWithPagination(namespace, []string{}, ctx.GetPageSize(), bookmark)
	if err != nil {
		return nil, "", err
	}
	defer func(resultIterator shim.StateQueryIteratorInterface) {
		_ = resultIterator.Close()
	}(resultIterator)

	for resultIterator.HasNext() {
		queryResponse, err := resultIterator.Next()
		if err != nil {
			return nil, "", err
		}
		obj := new(T)

		if err := json.Unmarshal(queryResponse.Value, &obj); err != nil {
			return nil, "", err
		}

		list = append(list, *obj)
	}

	return list, resMeta.GetBookmark(), nil
}
