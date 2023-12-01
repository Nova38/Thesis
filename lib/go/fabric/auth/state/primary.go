package state

import (
	"encoding/json"
	"log/slog"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/mennanov/fmutils"
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Primary Items

func UnmarshalPrimary[T common.ItemInterface](bytes []byte, obj T) (err error) {
	return json.Unmarshal(bytes, obj)
}
func UnmarshalNewPrimary[T common.ItemInterface](bytes []byte, base T) (item T, err error) {
	// item = new(T)
	item, ok := proto.Clone(base).(T)
	if !ok {
		return item, oops.Errorf("Failed to clone")
	}
	proto.Reset(item)

	if err = json.Unmarshal(bytes, item); err != nil {
		return item, oops.Wrap(err)
	}
	return item, nil
}

// ──────────────────────────────────────────────────
// Query Suggested Functions
// ──────────────────────────────────────────────────

// PrimaryExists returns true if the item exists in the ledger
func PrimaryExists[T common.ItemInterface](ctx common.TxCtxInterface, obj T) bool {

	key := lo.Must(common.MakePrimaryKey(obj))
	return common.KeyExists(ctx, key)
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
func PrimaryGet[T common.ItemInterface](ctx common.TxCtxInterface, obj T) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_VIEW,
		CollectionId: obj.ItemKey().GetCollectionId(),
		ItemType:     obj.ItemType(),
		Paths:        nil,
	}
	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	key, err := common.MakePrimaryKey(obj)
	if err != nil {
		return oops.Wrap(err)
	}

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

func PrimaryGetFull[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
) (fullItem *authpb.FullItem, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	// l := &Ledger[T]{ctx: ctx}
	fullItem = &authpb.FullItem{}

	// Get the item
	if err = PrimaryGet(ctx, obj); err != nil {
		return nil, oops.Wrap(err)
	}

	if fullItem.Item, err = common.PackItem(obj); err != nil {
		return nil, oops.Wrap(err)
	}

	fullItem.Suggestions, _, err = SuggestionListByItem(ctx, obj.ItemKey(), "")
	if err != nil {
		return nil, oops.Wrap(err)
	}

	// Get the history
	fullItem.History, err = getHistory(ctx, obj, true)
	if err != nil {
		return nil, oops.Wrap(err)
	}

	// // Get the References
	// refKeys, _, err := ReferencesByItem(ctx, obj.ItemKey(), "")
	// if err != nil {
	// 	return nil, oops.Wrap(err)
	// }
	// fullItem.References = refKeys

	return fullItem, nil
}

func PrimaryByPartialKey[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
	numAttr int,
	bookmark string,
) (list []T, mk string, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_VIEW,
		CollectionId: obj.ItemKey().GetCollectionId(),
		ItemType:     obj.ItemType(),
		Paths:        nil,
	}

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return nil, "", oops.Wrap(common.UserPermissionDenied)
	}

	var (
		itemType = obj.ItemType()
		attr     = obj.KeyAttr()
	)

	if len(attr) == 0 || len(attr) < numAttr {
		return nil, "", common.ItemInvalid
	}

	// Extract the attributes to search for
	// attr = attr[:len(attr)-numAttr]
	attr = lo.DropRight(attr, numAttr)

	ctx.GetLogger().
		Info("GetPartialKeyList",
			slog.Group(
				"Key", "ItemType", itemType,
				slog.Int("numAttr", numAttr),
				slog.Any("attr", attr),
				slog.Group(
					"Paged",
					"Bookmark", bookmark,
					"PageSize", strconv.Itoa(int(ctx.GetPageSize())),
				),
			),
		)

	results, meta, err := ctx.GetStub().
		GetStateByPartialCompositeKeyWithPagination(
			obj.ItemKey().GetItemType(),
			attr,
			ctx.GetPageSize(),
			bookmark,
		)
	if err != nil {
		return nil, "", err
	}
	defer func(results shim.StateQueryIteratorInterface) {
		err := results.Close()
		if err != nil {
			ctx.GetLogger().Error("GetPartialKeyList", "Error", err)
		}
	}(results)

	base := proto.Clone(obj).(T)
	proto.Reset(base)

	for results.HasNext() {
		queryResponse, err := results.Next()
		if err != nil {
			return nil, "", oops.Wrapf(err, "Error getting next item")
		}
		if queryResponse == nil {
			return nil, "", oops.Errorf("queryResponse is nil")
		}

		// if err := json.Unmarshal(queryResponse.GetValue(), obj); err != nil {
		// 	return nil, "", oops.Wrap(err)
		// }
		// var tmpObj
		item := proto.Clone(base).(T)

		if err = protojson.Unmarshal(queryResponse.GetValue(), item); err != nil {
			return nil, "", oops.Wrap(err)
		}

		list = append(list, item)
	}

	return list, meta.GetBookmark(), nil
	// l.GetPartialKeyList(obj, numAttr, bookmark)
}

func PrimaryList[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
	bookmark string,
) (list []T, mk string, err error) {
	return PrimaryByPartialKey(ctx, obj, len(obj.KeyAttr()), bookmark)
}

func ByCollection[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
	bookmark string,
) (list []T, mk string, err error) {
	return PrimaryByPartialKey(ctx, obj, 1, bookmark)
}

// ──────────────────────────────────────────────────
// Invoke Suggested Functions
// ──────────────────────────────────────────────────

// PrimaryCreate creates the item in the ledger
// returns error if the item already exists
// will panic if
//   - the key cannot be created,
//   - the item cannot be marshalled
//   - Authorization errors
func PrimaryCreate[T common.ItemInterface](ctx common.TxCtxInterface, obj T) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	// Authorize the operation
	ops := []*authpb.Operation{{
		Action:       authpb.Action_ACTION_CREATE,
		CollectionId: obj.ItemKey().GetCollectionId(),
		ItemType:     obj.ItemType(),
		Paths:        nil,
	}}

	auth, err := ctx.Authorize(ops)

	if err != nil {
		return oops.Wrap(err)
	} else if !auth {
		return oops.Wrap(common.UserPermissionDenied)
	}

	key, err := common.MakePrimaryKey(obj)

	if err != nil {
		return err
	}

	if Exists(ctx, key) {
		return oops.
			With("Key", key, "ItemType", obj.ItemType()).
			Wrap(common.AlreadyExists)
	}

	ctx.PostActionProcessing(obj, ops)

	if bytes, err := json.Marshal(obj); err != nil {
		return oops.Hint("Failed To Marshal").Wrap(err)
	} else {
		if err := ctx.GetStub().PutState(key, bytes); err != nil {
			return oops.With("key", key).Wrap(err)
		}
	}

	if ctx.EnabledHidden() {
		hiddenKey, err := common.MakeHiddenKey(obj)
		if err != nil {
			return oops.Wrap(err)
		}
		hidden := authpb.HiddenTxList{
			PrimaryKey: obj.ItemKey(),
			Txs:        []*authpb.HiddenTx{},
		}
		hiddenBytes, err := json.Marshal(&hidden)

		if err != nil {
			return oops.Wrap(err)
		}

		err = ctx.GetStub().PutState(hiddenKey, hiddenBytes)
		if err != nil {
			return oops.With(
				"Key", hiddenKey,
				"ItemType", obj.ItemType(),
			).Wrap(err)
		}
	}

	return nil
}

func PrimaryUpdate[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
	mask *fieldmaskpb.FieldMask,
) (updated T, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	auth, err := ctx.Authorize([]*authpb.Operation{{
		Action:       authpb.Action_ACTION_UPDATE,
		CollectionId: obj.ItemKey().GetCollectionId(),
		ItemType:     obj.ItemType(),
		Paths:        mask,
	}})
	if err != nil {
		return updated, oops.Wrap(err)
	} else if !auth {
		return updated, oops.Wrap(common.UserPermissionDenied)
	}

	// Get the current item from the ledger
	key, err := common.MakePrimaryKey(obj)
	if err != nil {
		return obj, err
	}

	current := proto.Clone(obj).(T)
	proto.Reset(current)

	current.SetKey(obj.ItemKey())

	if bytes, err := ctx.GetStub().GetState(key); err != nil {
		return obj, oops.Wrap(err)
	} else if err = json.Unmarshal(bytes, current); err != nil {
		return obj, oops.Wrap(err)
	}

	// Update the item
	fmutils.Overwrite(current, obj, mask.GetPaths())

	// Put the item back into the ledger
	if bytes, err := json.Marshal(current); err != nil {
		return current, oops.Wrap(err)
	} else if err := ctx.GetStub().PutState(key, bytes); err != nil {
		return current, oops.Wrap(err)
	}

	return current, nil

}

func PrimaryDelete[T common.ItemInterface](ctx common.TxCtxInterface, obj T) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	auth, err := ctx.Authorize([]*authpb.Operation{{
		Action:       authpb.Action_ACTION_DELETE,
		CollectionId: obj.ItemKey().GetCollectionId(),
		ItemType:     obj.ItemType(),
		Paths:        nil,
	}})

	if err != nil {
		return oops.Wrap(err)
	} else if !auth {
		return oops.Wrap(common.UserPermissionDenied)
	}

	// if err := referenceDeleteByItem(ctx, obj.ItemKey()); err != nil {
	// 	return oops.Wrap(err)
	// }
	// Should we delete the object refs in other collections? (there shouldn't be any except for users)

	if ctx.EnabledSuggestions() {
		if err := deleteSuggestionsByItem(ctx, obj.ItemKey()); err != nil {
			return oops.Wrap(err)
		}
	}

	if ctx.EnabledHidden() {
		if hiddenKey, err := common.MakeHiddenKey(obj); err != nil {
			return oops.Wrap(err)
		} else if hiddenKey != "" {
			if err := ctx.GetStub().DelState(hiddenKey); err != nil {
				return oops.Wrap(err)
			}
		}

	}

	if key, err := common.MakePrimaryKey(obj); err != nil {
		return oops.Wrap(err)
	} else if err := ctx.GetStub().DelState(key); err != nil {
		return oops.In("Primary").With("key", key).Wrap(err)
	}

	return nil
}

func deleteSuggestionsByItem(ctx common.TxCtxInterface, key *authpb.ItemKey) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	sList, _, err := SuggestionListByItem(ctx, key, "")
	if err != nil {
		return oops.Wrap(err)
	}
	for _, s := range sList {

		if sKey, err := common.MakeItemKeySuggestion(key, s.GetSuggestionId()); err != nil {
			return oops.Wrap(err)
		} else if err := ctx.GetStub().DelState(sKey); err != nil {
			return oops.Wrap(err)
		}
	}
	return nil
}

// func PrimaryDeleteFromKey(ctx TxCtxInterface, key *authpb.ItemKey) (obj *authpb.Item, err error) {
// 	k, err := MakeItemKeyPrimary(key)
// 	if err != nil {
// 		return obj, oops.Wrap(err)
// 	}

// 	bytes, err := ctx.GetStub().GetState(k)

// 	if bytes == nil && err == nil {
// 		return nil, oops.Wrap(common.KeyNotFound)
// 	}

// 	if err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	obj = &authpb.Item{}
// 	if err = json.Unmarshal(bytes, obj); err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	// obj, err = l.GetFromKey(k)
// 	// if err != nil {
// 	// 	return obj, oops.Wrap(err)
// 	// }

// 	// err = PrimaryDelete(ctx, obj)
// 	// if err != nil {
// 	// 	return obj, oops.Wrap(err)
// 	// }

// 	return obj, nil
// }
//
// func PrimaryGetFromKey[T common.ItemInterface](ctx TxCtxInterface, key *authpb.ItemKey) (obj T, err error) {
//	l := &Ledger[T]{ctx: ctx}
//	op := &authpb.Operation{
//		Action:       authpb.Action_ACTION_VIEW,
//		CollectionId: key.GetCollectionId(),
//		ItemType:   key.GetItemType(),
//		Paths:        nil,
//	}
//	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
//		return obj, oops.Wrap(common.UserPermissionDenied)
//	}
//
//	k, err := MakeItemKeyPrimary(key)
//	if err != nil {
//		return obj, oops.Wrap(err)
//	}
//
//	obj, err = l.GetFromKey(k)
//	if err != nil {
//		return obj, oops.Wrap(err)
//	}
//
//	return obj, l.Get(obj)
//}
