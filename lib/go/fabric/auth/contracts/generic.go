package contracts

import (
	"log/slog"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	// "github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	ccpb "github.com/nova38/thesis/lib/go/gen/chaincode/auth/common"
	"github.com/samber/oops"
)

type ItemContractImpl struct {
	contractapi.Contract
	ccpb.GenericServiceBase
}

// see if ItemContractImpl implements the interface GenericServiceInterface
var _ ccpb.GenericServiceInterface[common.TxCtxInterface] = (*ItemContractImpl)(nil)

// ════════════════════════════════════ Init ═══════════════════════════════════════

// ══════════════════════════════════ Helper ═════════════════════════════════════
// ────────────────────────────────── Query ──────────────────────────────────────

func (o ItemContractImpl) GetCurrentUser(
	ctx common.TxCtxInterface,
) (res *ccpb.GetCurrentUserResponse, err error) {

	res = &ccpb.GetCurrentUserResponse{}

	res.User, err = ctx.GetUserId()

	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	return res, err
}

// ──────────────────────────────────── Invoke ─────────────────────────────────────

func (o ItemContractImpl) AuthorizeOperation(
	ctx common.TxCtxInterface,
	req *ccpb.AuthorizeOperationRequest,
) (res *ccpb.AuthorizeOperationResponse, err error) {

	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	authorized, err := ctx.Authorize([]*authpb.Operation{req.GetOperation()})
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	return &ccpb.AuthorizeOperationResponse{
		Authorized: authorized,
	}, err
}

// ════════════════════════════════════ Item ═════════════════════════════════════

// ──────────────────────────────────── Query ──────────────────────────────────────

func (o ItemContractImpl) Get(
	ctx common.TxCtxInterface,
	req *ccpb.GetRequest,
) (res *ccpb.GetResponse, err error) {
	var (
		obj common.ItemInterface
		msg *authpb.Item
	)

	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	if obj, err = common.ItemKeyToItem(req.GetKey()); err != nil {
		return nil, oops.Wrap(err)
	}

	if err = state.PrimaryGet(ctx, obj); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", (err)))
		return nil, oops.Wrap(err)
	}

	if msg, err = common.PackItem(obj); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, err
	} else {
		return &ccpb.GetResponse{
			Item: msg,
		}, nil
	}
}

func (o ItemContractImpl) List(
	ctx common.TxCtxInterface,
	req *ccpb.ListRequest,
) (res *ccpb.ListResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	item, err := common.ItemKeyToItemType(req.GetKey())
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	list, mk, err := state.PrimaryList(ctx, item, req.GetBookmark())
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	res = &ccpb.ListResponse{
		Bookmark: mk,
		Items:    []*authpb.Item{},
	}

	if res.Items, err = common.ListItemToProtos(list); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	return res, err
}

func (o ItemContractImpl) ListByCollection(
	ctx common.TxCtxInterface,
	req *ccpb.ListByCollectionRequest,
) (res *ccpb.ListByCollectionResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	item, err := common.ItemKeyToItemType(req.GetKey())
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	item.SetKey(&authpb.ItemKey{CollectionId: req.GetKey().GetCollectionId()})

	list, mk, err := state.PrimaryList(ctx, item, req.GetBookmark())
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	res = &ccpb.ListByCollectionResponse{
		Bookmark: mk,
		Items:    []*authpb.Item{},
	}

	if res.Items, err = common.ListItemToProtos(list); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	return res, err
}

func (o ItemContractImpl) ListByAttrs(
	ctx common.TxCtxInterface,
	req *ccpb.ListByAttrsRequest,
) (res *ccpb.ListByAttrsResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	item, err := common.ItemKeyToItem(req.GetKey())
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	list, mk, err := state.PrimaryByPartialKey(ctx, item, int(req.GetNumAttrs()), req.GetBookmark())
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))

		return nil, oops.Wrap(err)
	}

	res = &ccpb.ListByAttrsResponse{
		Bookmark: mk,
		Items:    []*authpb.Item{},
	}

	if res.Items, err = common.ListItemToProtos(list); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	return res, err
}

// ──────────────────────────────────── Invoke ─────────────────────────────────────

func (o ItemContractImpl) Create(
	ctx common.TxCtxInterface,
	req *ccpb.CreateRequest,
) (res *ccpb.CreateResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	// Get the item from the request
	obj, err := common.UnPackItem(req.GetItem())
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	// Check if the item is a valid item for the collection??

	if err = state.PrimaryCreate(ctx, obj); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	return &ccpb.CreateResponse{
		Item: req.GetItem(),
	}, err
}

func (o ItemContractImpl) Update(
	ctx common.TxCtxInterface,
	req *ccpb.UpdateRequest,
) (res *ccpb.UpdateResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	// Get the item from the request
	obj, err := common.UnPackItem(req.GetItem())
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	updated, err := state.PrimaryUpdate(ctx, obj, req.GetUpdateMask())
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	if item, err := common.PackItem(updated); err != nil {
		return nil, oops.Wrap(err)
	} else {
		res = &ccpb.UpdateResponse{
			Item: item,
		}
	}

	return res, err
}

func (o ItemContractImpl) Delete(
	ctx common.TxCtxInterface,
	req *ccpb.DeleteRequest,
) (res *ccpb.DeleteResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	// Get the item from the request
	obj, err := common.UnPackItem(req.GetItem())
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	if err = state.PrimaryDelete(ctx, obj); err != nil {
		return nil, oops.Wrap(err)
	}

	item, err := common.PackItem(obj)
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	return &ccpb.DeleteResponse{
		Item: item,
	}, err
}

// ════════════════════════════════════ History ════════════════════════════════════
// ──────────────────────────────────── Query ──────────────────────────────────────

func (o ItemContractImpl) History(
	ctx common.TxCtxInterface,
	req *ccpb.HistoryRequest,
) (res *ccpb.HistoryResponse, err error) {
	var (
		obj common.ItemInterface
		h   *authpb.History
	)

	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	if obj, err = common.ItemKeyToItem(req.GetKey()); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	if h, err = state.GetHistory(ctx, obj); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	return &ccpb.HistoryResponse{
		Key:     req.GetKey(),
		History: h,
	}, nil
}

func (o ItemContractImpl) HiddenTx(
	ctx common.TxCtxInterface,
	req *ccpb.HiddenTxRequest,
) (res *ccpb.HiddenTxResponse, err error) {
	var (
		obj  common.ItemInterface
		hTxs *authpb.HiddenTxList
	)

	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	if obj, err = common.UnPackItem(req.GetItem()); err != nil {
		return nil, oops.Wrap(err)
	}

	if hTxs, err = state.GetHiddenTx(ctx, obj); err != nil {
		return nil, oops.Wrap(err)
	}

	return &ccpb.HiddenTxResponse{
		CollectionId: obj.ItemKey().GetCollectionId(),
		HiddenTxs:    hTxs.GetTxs(),
	}, nil
}

// ──────────────────────────────────── Invoke ─────────────────────────────────────

func (o ItemContractImpl) HideTx(
	ctx common.TxCtxInterface,
	req *ccpb.HideTxRequest,
) (res *ccpb.HideTxResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	// Get the obj from the request

	obj, err := common.ItemKeyToItem(req.GetKey())
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	list, err := state.HideTransaction(ctx, obj, req.GetHiddenTx())
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	// Get the obj from the request from world state to return it

	return &ccpb.HideTxResponse{
		Key:       req.GetKey(),
		HiddenTxs: list,
	}, err
}

func (o ItemContractImpl) UnHideTx(
	ctx common.TxCtxInterface,
	req *ccpb.UnHideTxRequest,
) (res *ccpb.UnHideTxResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	// Get the item from the request
	//obj, err := common.UnPackItem(req.GetItem())
	obj, err := common.ItemKeyToItem(req.GetKey())
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	list, err := state.UnHideTransaction(ctx, obj, req.GetTxId())
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	return &ccpb.UnHideTxResponse{
		Key:       req.GetKey(),
		HiddenTxs: list,
	}, err
}

// ════════════════════════════════════ Suggestions ════════════════════════════════
// ──────────────────────────────────── Query ──────────────────────────────────────

func (o ItemContractImpl) Suggestion(
	ctx common.TxCtxInterface,
	req *ccpb.SuggestionRequest,
) (res *ccpb.SuggestionResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	sug := &authpb.Suggestion{
		PrimaryKey:   req.GetItemKey(),
		SuggestionId: req.GetSuggestionId(),
	}

	if err = state.GetSuggestion(ctx, sug); err != nil {
		return nil, oops.Wrap(err)
	}

	return &ccpb.SuggestionResponse{
		Suggestion: sug,
	}, nil
}

func (o ItemContractImpl) SuggestionListByCollection(
	ctx common.TxCtxInterface,
	req *ccpb.SuggestionListByCollectionRequest,
) (res *ccpb.SuggestionListByCollectionResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	list, mk, err := state.SuggestionListByCollection(ctx, req.GetCollectionId(), req.GetBookmark())
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	res = &ccpb.SuggestionListByCollectionResponse{
		Bookmark:    mk,
		Suggestions: list,
	}

	return res, nil
}

func (o ItemContractImpl) SuggestionByPartialKey(
	ctx common.TxCtxInterface,
	req *ccpb.SuggestionByPartialKeyRequest,
) (res *ccpb.SuggestionByPartialKeyResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	sug := &authpb.Suggestion{
		PrimaryKey: req.GetItemKey(),
	}

	list, mk, err := state.PartialSuggestionList(
		ctx,
		sug,
		int(req.GetNumAttrs()),
		req.GetBookmark(),
	)
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	res = &ccpb.SuggestionByPartialKeyResponse{
		Bookmark:    mk,
		Suggestions: list,
	}

	return res, nil
}

// ──────────────────────────────── Invoke ─────────────────────────────────────────

func (o ItemContractImpl) SuggestionCreate(
	ctx common.TxCtxInterface,
	req *ccpb.SuggestionCreateRequest,
) (res *ccpb.SuggestionCreateResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	if err = state.SuggestionCreate(ctx, req.GetSuggestion()); err != nil {
		ctx.GetLogger().Warn("SuggestionCreate", "err", err)
		return nil, oops.Wrap(err)
	}

	return &ccpb.SuggestionCreateResponse{
		Suggestion: req.GetSuggestion(),
	}, nil
}

func (o ItemContractImpl) SuggestionDelete(
	ctx common.TxCtxInterface,
	req *ccpb.SuggestionDeleteRequest,
) (res *ccpb.SuggestionDeleteResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	sug := &authpb.Suggestion{
		PrimaryKey:   req.GetItemKey(),
		SuggestionId: req.GetSuggestionId(),
	}

	if err = state.SuggestionDelete(ctx, sug); err != nil {
		return nil, oops.Wrap(err)
	}

	return &ccpb.SuggestionDeleteResponse{
		Suggestion: sug,
	}, nil
}

func (o ItemContractImpl) SuggestionApprove(
	ctx common.TxCtxInterface,
	req *ccpb.SuggestionApproveRequest,
) (res *ccpb.SuggestionApproveResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}
	sug := &authpb.Suggestion{
		PrimaryKey:   req.GetItemKey(),
		SuggestionId: req.GetSuggestionId(),
	}

	u, err := state.SuggestionApprove(ctx, sug)
	if err != nil {
		ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
		return nil, oops.Wrap(err)
	}

	if updated, err := common.PackItem(*u); err != nil {
		return nil, oops.Wrap(err)
	} else {
		res = &ccpb.SuggestionApproveResponse{
			Item:       updated,
			Suggestion: sug,
		}
	}

	return res, nil
}
