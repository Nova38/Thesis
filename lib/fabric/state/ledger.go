package state

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Generic structures

//goland:noinspection GoNameStartsWithPackageName
type StateObject interface {
	Key() (attr []string, err error)
	Namespace() (namespace string)
}

//goland:noinspection GoNameStartsWithPackageName
type StateObjectList[T StateObject] struct {
	Entries []T `json:"entries"`
}

type HistoryEntry[T StateObject] struct {
	TxId      string                 `json:"txId"`
	Timestamp *timestamppb.Timestamp `json:"timestamp"`
	IsDelete  bool                   `json:"isDelete"`
	T         T                      `json:"object"`
}

type HistoryList[T StateObject] struct {
	Entries []HistoryEntry[T] `json:"entries"`
}

// UTIL Functions

// MakeCompositeKey creates a composite key from the given attributes
func MakeCompositeKey[T StateObject](ctx contractapi.TransactionContextInterface, obj T) (key string, err error) {
	namespace := obj.Namespace()
	attr, err := obj.Key()
	if err != nil {
		return "", err
	}

	slog.Info("MakeCompositeKey", slog.Group("Key", "Namespace", namespace, "attr", attr))
	key, err = ctx.GetStub().CreateCompositeKey(namespace, attr)

	if err != nil {
		return "", err
	}
	return key, nil
}

func ObjExists[T StateObject](ctx contractapi.TransactionContextInterface, obj T) (bool, error) {
	key, err := MakeCompositeKey(ctx, obj)
	if err != nil {
		return false, err
	}

	return Exists(ctx, key), nil
}

// Exists returns true if the object exists in the ledger
func Exists(ctx contractapi.TransactionContextInterface, key string) bool {

	bytes, err := ctx.GetStub().GetState(key)
	if bytes == nil && err == nil {
		return false
	}

	return err == nil
}

// PutState puts the object into the ledger
func PutState[T StateObject](ctx contractapi.TransactionContextInterface, obj T) (err error) {

	key, err := MakeCompositeKey(ctx, obj)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(key, bytes)
}

// InsertState inserts the object into the ledger
// returns error if the object already exists
func InsertState[T StateObject](ctx contractapi.TransactionContextInterface, obj T) (err error) {

	key, err := MakeCompositeKey(ctx, obj)
	if err != nil {
		return err
	}

	if Exists(ctx, key) {
		return &KeyAlreadyExistsError{
			Key:       key,
			Namespace: obj.Namespace(),
			MSG:       "Key already exists",
		}
	}

	bytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(key, bytes)
}

// UpdateState updates the object in the ledger
// returns error if the object does not exist
func UpdateState[T StateObject](ctx contractapi.TransactionContextInterface, obj T) (err error) {

	key, err := MakeCompositeKey(ctx, obj)
	if err != nil {
		return err
	}

	if !Exists(ctx, key) {
		return &KeyNotFoundError{
			Key:       key,
			Namespace: obj.Namespace(),
			MSG:       "Key not found",
		}
	}

	bytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(key, bytes)
}

// GetState returns the object from the ledger
func GetState[T StateObject](ctx contractapi.TransactionContextInterface, in T) (err error) {
	namespace := in.Namespace()
	slog.Info("fn: GetState", "Namespace", namespace)

	key, err := MakeCompositeKey(ctx, in)
	if err != nil {
		return err
	}

	bytes, err := ctx.GetStub().GetState(key)
	if bytes == nil && err == nil {
		return &KeyNotFoundError{
			Key:       key,
			Namespace: namespace,
			MSG:       "Key Not Found",
		}
	}

	if err = json.Unmarshal(bytes, in); err != nil {
		return err
	}
	return nil
}

// DeleteState deletes the object from the ledger
func DeleteState[T StateObject](ctx contractapi.TransactionContextInterface, in T) (err error) {
	key, err := MakeCompositeKey(ctx, in)
	if err != nil {
		return err
	}

	return ctx.GetStub().DelState(key)
}

// GetStateHistory returns the history of the object from the ledger
func GetStateHistory[T StateObject](ctx contractapi.TransactionContextInterface, in T) (list HistoryList[T], err error) {
	key, err := MakeCompositeKey(ctx, in)
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

		entry := HistoryEntry[T]{
			TxId:      queryResponse.TxId,
			Timestamp: queryResponse.Timestamp,
			IsDelete:  queryResponse.IsDelete,
			T:         *obj,
		}

		list.Entries = append(list.Entries, entry)
	}

	return list, nil
}

func TxIdInHistory[T StateObject](ctx contractapi.TransactionContextInterface, in T, txId string) (bool, error) {
	key, err := MakeCompositeKey(ctx, in)
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

// GetPartialKeyList returns a list of objects of type T
// T must implement StateObject interface
// numAttr is the number of attributes in the key to search for
func GetPartialKeyList[T StateObject](ctx contractapi.TransactionContextInterface, in T, numAttr int) (list []T, err error) {
	// obj = []*T{}
	slog.Info("GetPartialKeyList")
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

	slog.Info("GetPartialKeyList", slog.Group("Key", "Namespace", namespace, slog.Int("numAttr", numAttr), slog.Any("attr", attr)))

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

func GetFullStateList[T StateObject](ctx contractapi.TransactionContextInterface, in T) (list []T, err error) {
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
