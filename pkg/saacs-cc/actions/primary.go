package actions

import (
	"log/slog"

	"github.com/nova38/saacs/pkg/saacs-cc/common"
	"github.com/nova38/saacs/pkg/saacs-cc/state"
	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"

	"github.com/samber/oops"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Primary Items

// ──────────────────────────────────────────────────
// Query Suggested Functions
// ──────────────────────────────────────────────────

func PrimaryGet[T common.ItemInterface](ctx common.TxCtxInterface, obj T) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	op := &pb.Operation{
		Action:       pb.Action_ACTION_VIEW,
		CollectionId: obj.ItemKey().GetCollectionId(),
		ItemType:     obj.ItemType(),
		Paths:        nil,
	}
	if auth, err := ctx.Authorize([]*pb.Operation{op}); !auth || err != nil {
		return ctx.ErrorBase().Wrap(common.UserPermissionDenied)
	}

	if err := state.Get(ctx, obj); err != nil {
		return oops.Wrap(err)
	}

	return nil
}

func PrimaryGetFull[T common.ItemInterface](
	ctx common.TxCtxInterface,
	obj T,
	historyOptions *pb.HistoryOptions,
) (fullItem *pb.FullItem, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	ops := []*pb.Operation{
		{
			Action:       pb.Action_ACTION_VIEW,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		},
		{
			Action:       pb.Action_ACTION_SUGGEST_VIEW,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		},
		{
			Action:       pb.Action_ACTION_VIEW_HISTORY,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		},
	}

	if historyOptions.GetHidden().GetInclude() {
		ops = append(ops, &pb.Operation{
			Action:       pb.Action_ACTION_VIEW_HIDDEN_TXS,
			CollectionId: obj.ItemKey().GetCollectionId(),
			ItemType:     obj.ItemType(),
			Paths:        nil,
		})

		if historyOptions.GetHidden().GetMspIds() != nil &&
			len(historyOptions.GetHidden().GetMspIds()) > 0 {
			ops = append(ops, &pb.Operation{
				Action:       pb.Action_ACTION_VIEW_MSP_HIDDEN_TX,
				CollectionId: obj.ItemKey().GetCollectionId(),
				ItemType:     obj.ItemType(),
				Paths:        nil,
			})
		}
	}

	if auth, err := ctx.Authorize(ops); !auth || err != nil {
		return nil, ctx.ErrorBase().Wrap(common.UserPermissionDenied)
	}

	// l := &Ledger[T]{ctx: ctx}
	fullItem = &pb.FullItem{}

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
	fullItem.History, err = getHistory(ctx, obj, historyOptions.GetHidden())
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

	op := &pb.Operation{
		Action:       pb.Action_ACTION_VIEW,
		CollectionId: obj.ItemKey().GetCollectionId(),
		ItemType:     obj.ItemType(),
		Paths:        nil,
	}

	if auth, err := ctx.Authorize([]*pb.Operation{op}); !auth || err != nil {
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
	return PrimaryByPartialKey(ctx, obj, 0, bookmark)
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
	ops := []*pb.Operation{{
		Action:       pb.Action_ACTION_CREATE,
		CollectionId: obj.ItemKey().GetCollectionId(),
		ItemType:     obj.ItemType(),
		Paths:        nil,
	}}

	if auth, err := ctx.Authorize(ops); err != nil {
		return oops.Wrap(err)
	} else if !auth {
		return oops.Wrap(common.UserPermissionDenied)
	}

	if err = ctx.PostActionProcessing(obj, ops); err != nil {
		return oops.Wrap(err)
	}

	// if err := state.Insert(ctx, obj); err != nil {
	// 	return oops.Wrap(err)
	// }

	if err = (state.Ledger[T]{}.PrimaryCreate(ctx, obj)); err != nil {
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

	ops := []*pb.Operation{{
		Action:       pb.Action_ACTION_UPDATE,
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

	ctx.GetLogger().Info("obj key: %v")

	// current, ok := proto.Clone(obj).(T)

	current, ok := obj.ProtoReflect().New().Interface().(T)
	if !ok {
		return updated, ctx.ErrorBase().
			With("operation", ops).
			Wrap(common.ItemInvalid)
	}
	current.SetKey(obj.ItemKey())

	// ctx.GetLogger().Info("current: %v", current)

	if err = state.Get(ctx, current); err != nil {
		return obj, ctx.ErrorBase().Wrap(err)
	}

	ctx.GetLogger().Info("Update Spec => ",
		slog.Any("obj", obj),
		slog.Any("key", obj.ItemType()),
		slog.Any("mask", mask),
		slog.Any("current", current),
	)

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

	auth, err := ctx.Authorize([]*pb.Operation{{
		Action:       pb.Action_ACTION_DELETE,
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

func deleteSuggestionsByItem(ctx common.TxCtxInterface, key *pb.ItemKey) (err error) {
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
