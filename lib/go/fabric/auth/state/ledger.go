package state

import (
	"encoding/json"

	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/samber/oops"
)

// UTIL Functions

// Returns true if the key exists in the ledger
func Exists(ctx common.TxCtxInterface, key string) bool {
	bytes, err := ctx.GetStub().GetState(key)
	if bytes == nil && err == nil {
		return false
	}

	return err == nil
}

// Get returns the item from the ledger
func Get[T common.ItemInterface](ctx common.TxCtxInterface, key string, obj T) (err error) {
	bytes, err := ctx.GetStub().GetState(key)
	if bytes == nil && err == nil {
		return oops.
			With("Key", key, "ItemType", obj.ItemType()).
			Wrap(common.KeyNotFound)
	}
	if err != nil {
		return oops.
			With("Key", key, "ItemType", obj.ItemType()).
			Wrap(err)
	}

	if err = json.Unmarshal(bytes, obj); err != nil {
		return oops.Wrap(err)
	}

	return nil
}

func Put[T common.ItemInterface](ctx common.TxCtxInterface, key string, obj T) (err error) {
	var (
		bytes []byte
	)

	if bytes, err = json.Marshal(obj); err != nil {
		return err
	}

	return ctx.GetStub().PutState(key, bytes)
}

// ════════════════════════════════════════════════════════
// Invoke Functions
// ════════════════════════════════════════════════════════

// Insert inserts the item into the ledger
// returns error if the item already exists

// ════════════════════════════════════════════════════════
// Query Functions
// ════════════════════════════════════════════════════════

// func (l *Ledger[T]) GetFromKey(key string) (obj T, err error) {
// 	return GetFromKey[T](l.ctx, key)
// }

// GetPartialKeyList returns a list of items of type T
// T must implement StateItem interface
// numAttr is the number of attributes in the key to search for
// func (l *Ledger[T]) GetPartialKeyList(
// 	obj T,
// 	numAttr int,
// 	bookmark string,
// ) (list []*T, mk string, err error) {
// 	// obj = []*T{}
// 	l.ctx.GetLogger().Info("GetPartialKeyList")

// 	var (
// 		itemtype = obj.ItemType()
// 		attr     = obj.KeyAttr()
// 	)

// 	if len(attr) == 0 || len(attr) < numAttr {
// 		return nil, "", common.ItemInvalid
// 	}

// 	// Extract the attributes to search for
// 	// attr = attr[:len(attr)-numAttr]
// 	attr = lo.DropRight(attr, numAttr)

// 	l.ctx.GetLogger().
// 		Info("GetPartialKeyList",
// 			slog.Group(
// 				"Key", "ItemType", itemtype,
// 				slog.Int("numAttr", numAttr),
// 				slog.Any("attr", attr),
// 				slog.Group(
// 					"Paged",
// 					"Bookmark", bookmark,
// 					"PageSize", strconv.Itoa(int(l.ctx.GetPageSize())),
// 				),
// 			),
// 		)

// 	results, meta, err := l.ctx.GetStub().
// 		GetStateByPartialCompositeKeyWithPagination(
// 			obj.ItemKey().GetItemType(),
// 			attr,
// 			l.ctx.GetPageSize(),
// 			bookmark,
// 		)
// 	if err != nil {
// 		return nil, "", err
// 	}
// 	defer func(results shim.StateQueryIteratorInterface) {
// 		err := results.Close()
// 		if err != nil {
// 			l.ctx.GetLogger().Error("GetPartialKeyList", "Error", err)
// 		}
// 	}(results)

// 	for results.HasNext() {
// 		queryResponse, err := results.Next()
// 		if err != nil {
// 			return nil, "", oops.Wrapf(err, "Error getting next item")
// 		}

// 		if queryResponse == nil {
// 			return nil, "", oops.Errorf("queryResponse is nil")
// 		}

// 		var tmpObj T

// 		// if err := json.Unmarshal(queryResponse.GetValue(), obj); err != nil {
// 		// 	return nil, "", oops.Wrap(err)
// 		// }

// 		if err = protojson.Unmarshal(queryResponse.GetValue(), tmpObj); err != nil {
// 			return nil, "", oops.Wrap(err)
// 		}

// 		list = append(list, &obj)
// 	}

// 	return list, meta.GetBookmark(), nil
// }

// ════════════════════════════════════════════════════════
// Raw Functions
// ════════════════════════════════════════════════════════

// func GetFromKey[T common.ItemInterface](ctx common.TxCtxInterface, key string) (obj T, err error) {
// 	bytes, err := ctx.GetStub().GetState(key)
// 	if bytes == nil && err == nil {
// 		return obj, oops.
// 			With("Key", key, "ItemType", obj.ItemType()).
// 			Wrap(common.KeyNotFound)
// 	} else if err != nil {
// 		return obj, oops.Wrap(err)
// 	}

// 	if err = json.Unmarshal(bytes, obj); err != nil {
// 		return obj, oops.Wrap(err)
// 	}

// 	return obj, nil
// }

// bytes, err := l.ctx.GetStub().GetState(key)
// 	if bytes == nil && err == nil {
// 		return obj, oops.
// 			With("Key", key, "ItemType", obj.ItemType()).
// 			Wrap(common.KeyNotFound)
// 	} else if err != nil {
// 		return obj, oops.Wrap(err)
// 	}

// 	if err = json.Unmarshal(bytes, obj); err != nil {
// 		return obj, oops.Wrap(err)
// 	}

// 	return obj, nil

// ════════════════════════════════════════════════════════
// Invoke Functions
// ════════════════════════════════════════════════════════

// Insert inserts the item into the ledger
// returns error if the item already exists
//func Insert[T common.ItemInterface](ctx common.TxCtxInterface, obj T) (err error) {
//	var (
//		key   string
//		bytes []byte
//	)
//
//	if key, err = common.MakePrimaryKey(obj); err != nil {
//		return err
//	}
//
//	if Exists(ctx, key) {
//		return oops.
//			With("Key", key, "ItemType", obj.ItemType()).
//			Wrap(common.AlreadyExists)
//	}
//
//	if bytes, err = json.Marshal(obj); err != nil {
//		return err
//	}
//
//	return ctx.GetStub().PutState(key, bytes)
//}

// Edit updates the item in the ledger
// returns error if the item does not exist
//func Update[T common.ItemInterface](
//	ctx common.TxCtxInterface,
//	update T,
//	mask *fieldmaskpb.FieldMask,
//) (obj T, err error) {
//	var (
//		key   string
//		bytes []byte
//	)
//
//}
//
//// Delete deletes the item from the ledger
//func Delete[T common.ItemInterface](ctx common.TxCtxInterface, in T) (err error) {
//	if err != nil {
//		return err
//	}
//
//	if err = ctx.GetStub().DelState(key); err != nil {
//		return oops.Wrap(err)
//	}
//
//	return nil
//}
