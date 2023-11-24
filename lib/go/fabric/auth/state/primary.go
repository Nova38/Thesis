package state

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/common"

	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Primary Items

// ──────────────────────────────────────────────────
// Query Suggested Functions
// ──────────────────────────────────────────────────
func PrimaryExists[T common.ItemInterface](ctx TxCtxInterface, obj T) bool {
	key := lo.Must(MakePrimaryKey(obj))
	return KeyExists(ctx, key)
}

func PrimaryGet[T common.ItemInterface](ctx TxCtxInterface, obj T) (err error) {
	l := &Ledger[T]{ctx: ctx}
	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_VIEW,
		CollectionId: obj.ItemKey().GetCollectionId(),
		ItemType:     obj.ItemType(),
		Paths:        nil,
	}
	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	return l.Get(obj)
}

func PrimaryGetFull[T common.ItemInterface](
	ctx TxCtxInterface,
	obj T,
) (item *FullItem[T], err error) {
	l := &Ledger[T]{ctx: ctx}
	item = &FullItem[T]{}

	ops := []*authpb.Operation{
		{
			Action:       authpb.Action_ACTION_VIEW,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		},
		{
			Action:       authpb.Action_ACTION_VIEW_HIDDEN_TXS,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		},
		{
			Action:       authpb.Action_ACTION_VIEW_HIDDEN_TXS,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		},
	}
	if auth, err := ctx.Authorize(ops); !auth || err != nil {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}
	// Get the item
	if err = l.Get(obj); err != nil {
		return nil, oops.Wrap(err)
	}
	item.Item = obj

	item.Suggestions, _, err = SuggestionListByItem(ctx, obj.ItemKey(), "")
	if err != nil {
		return nil, oops.Wrap(err)
	}

	// Get the history
	item.History, err = history(ctx, obj, true)
	if err != nil {
		return nil, oops.Wrap(err)
	}

	// state.ItemToAuthObj(obj, item.Value)

	return item, nil
}

func PrimaryByPartialKey[T common.ItemInterface](
	ctx TxCtxInterface,
	obj T,
	numAttr int,
	bookmark string,
) (list []T, mk string, err error) {
	l := &Ledger[T]{ctx: ctx}

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_VIEW,
		CollectionId: obj.ItemKey().GetCollectionId(),
		ItemType:     obj.ItemType(),
		Paths:        nil,
	}

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return nil, "", oops.Wrap(common.UserPermissionDenied)
	}

	return l.GetPartialKeyList(obj, numAttr, bookmark)
}

func PrimaryList[T common.ItemInterface](
	ctx TxCtxInterface,
	obj T,
	bookmark string,
) (list []T, mk string, err error) {
	return PrimaryByPartialKey(ctx, obj, 1, bookmark)
}

func ByCollection[T common.ItemInterface](
	ctx TxCtxInterface,
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
func PrimaryCreate[T common.ItemInterface](ctx TxCtxInterface, obj T) (err error) {
	l := &Ledger[T]{
		ctx: ctx,
	}

	// Authorize the operation
	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_CREATE,
		CollectionId: obj.ItemKey().GetCollectionId(),
		ItemType:     obj.ItemType(),
		Paths:        nil,
	}

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	// TODO: Handle Creating suggestions Domain
	// TODO: Handle Creating hiddenTx Domain

	return l.Create(obj)
}

func PrimaryUpdate[T common.ItemInterface](
	ctx TxCtxInterface,
	obj T,
	mask *fieldmaskpb.FieldMask,
) (err error) {
	l := &Ledger[T]{
		ctx: ctx,
	}
	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_UPDATE,
		CollectionId: obj.ItemKey().GetCollectionId(),
		ItemType:     obj.ItemType(),
		Paths:        mask,
	}

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	return l.Update(obj, mask)
}

func PrimaryDelete[T common.ItemInterface](ctx TxCtxInterface, obj T) (err error) {
	l := &Ledger[T]{
		ctx: ctx,
	}
	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_DELETE,
		CollectionId: obj.ItemKey().GetCollectionId(),
		ItemType:     obj.ItemType(),
		Paths:        nil,
	}

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	err = l.Delete(obj)

	// TODO: Handle deleting refs/sub items here

	// TODO: Handle deleting suggestions here
	if err := deleteSuggestionsByItem(ctx, obj.ItemKey()); err != nil {
		return oops.Wrap(err)
	}

	// TODO: Handle deleting hiddenTx items here
	if hiddenKey, err := MakeHiddenKey(obj); err != nil {
		return oops.Wrap(err)
	} else if hiddenKey != "" {
		if err := ctx.GetStub().DelState(hiddenKey); err != nil {
			return oops.Wrap(err)
		}
	}

	return err
}

func deleteSuggestionsByItem(ctx TxCtxInterface, key *authpb.ItemKey) (err error) {
	l := &Ledger[*authpb.Suggestion]{ctx: ctx}

	sList, _, err := SuggestionListByItem(ctx, key, "")
	if err != nil {
		return oops.Wrap(err)
	}
	for _, s := range sList {
		if err := l.Delete(s); err != nil {
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
