package actions

import (
	"github.com/nova38/thesis/packages/saacs/common"
	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	"github.com/nova38/thesis/packages/saacs/state"

	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Primary Items

// ──────────────────────────────────────────────────
// Query Suggested Functions
// ──────────────────────────────────────────────────

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

	if err := state.Get(ctx, obj); err != nil {
		return oops.Wrap(err)
	}

	return nil

}

func PrimaryGetFull[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
	showHidden bool,
) (fullItem *authpb.FullItem, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	ops := []*authpb.Operation{
		{
			Action:       authpb.Action_ACTION_VIEW,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		},
		{
			Action:       authpb.Action_ACTION_SUGGEST_VIEW,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		},
		{
			Action:       authpb.Action_ACTION_VIEW_HISTORY,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		},
	}

	if showHidden {
		ops = append(ops, &authpb.Operation{
			Action:       authpb.Action_ACTION_VIEW_HIDDEN_TXS,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		})
	}

	if auth, err := ctx.Authorize(ops); !auth || err != nil {
		return nil, ctx.ErrorBase().Wrap(common.UserPermissionDenied)
	}

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

	return state.GetPartialKeyList(ctx, obj, numAttr, bookmark)
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

	if auth, err := ctx.Authorize(ops); err != nil {
		return oops.Wrap(err)
	} else if !auth {
		return oops.Wrap(common.UserPermissionDenied)
	}

	if err := ctx.PostActionProcessing(obj, ops); err != nil {
		return oops.Wrap(err)
	}

	if err := state.Insert(ctx, obj); err != nil {
		return oops.Wrap(err)
	}

	return nil
}

func PrimaryUpdate[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
	mask *fieldmaskpb.FieldMask,
) (updated T, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	ops := []*authpb.Operation{{
		Action:       authpb.Action_ACTION_UPDATE,
		CollectionId: obj.ItemKey().GetCollectionId(),
		ItemType:     obj.ItemType(),
		Paths:        mask,
	}}

	if auth, err := ctx.Authorize(ops); err != nil {
		return updated, ctx.ErrorBase().
			With("operation", ops).
			Wrap(err)
	} else if !auth {
		return updated, ctx.ErrorBase().
			With("operation", ops).
			Wrap(common.UserPermissionDenied)
	}

	current, ok := obj.ProtoReflect().New().Interface().(T)
	if !ok {
		return updated, ctx.ErrorBase().
			With("operation", ops).
			Wrap(common.ItemInvalid)
	}
	proto.Reset(current)

	if err = state.Get(ctx, obj); err != nil {
		return obj, ctx.ErrorBase().Wrap(err)
	}

	val, err := common.UpdateItem(mask, current, obj)
	if err != nil {
		return obj, ctx.ErrorBase().Wrap(err)
	}

	if err = ctx.PostActionProcessing(val, ops); err != nil {
		return obj, ctx.ErrorBase().Wrap(err)
	}
	if err = state.Put(ctx, val); err != nil {
		return obj, ctx.ErrorBase().Wrap(err)
	}

	return val, nil

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
		return ctx.ErrorBase().Wrap(common.UserPermissionDenied)
	}

	if ctx.EnabledSuggestions() {
		if err := deleteSuggestionsByItem(ctx, obj.ItemKey()); err != nil {
			return ctx.ErrorBase().Wrap(err)
		}
	}

	if ctx.EnabledHidden() {
		if hiddenKey, err := common.MakeHiddenKey(obj); err != nil {
			return ctx.ErrorBase().Wrap(err)
		} else if hiddenKey != "" {
			if err := ctx.GetStub().DelState(hiddenKey); err != nil {
				return ctx.ErrorBase().Wrap(err)
			}
		}
	}

	if err := state.Delete(ctx, obj); err != nil {
		return ctx.ErrorBase().
			In("Primary").
			With("key", obj.ItemKey(), "state key", obj.StateKey()).
			Wrap(err)
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
		if err := state.Delete(ctx, s); err != nil {
			return oops.Wrap(err)
		}
	}
	return nil
}
