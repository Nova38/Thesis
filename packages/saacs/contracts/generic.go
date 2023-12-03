package contracts

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	// "github.com/nova38/thesis/packages/saacs/auth/common"
	"github.com/nova38/thesis/packages/fabric/auth/state"
	"github.com/nova38/thesis/packages/saacs/auth/common"
	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	cc "github.com/nova38/thesis/packages/saacs/genchaincode/auth/common"
	"github.com/samber/oops"
)

type ItemContractImpl struct {
	contractapi.Contract
	cc.GenericServiceBase
}

// see if ItemContractImpl implements the interface GenericServiceInterface
var _ cc.GenericServiceInterface[common.TxCtxInterface] = (*ItemContractImpl)(nil)

// ════════════════════════════════════ Init ═══════════════════════════════════════

// ══════════════════════════════════ Helper ═════════════════════════════════════
// ────────────────────────────────── Query ──────────────────────────────────────

func (o ItemContractImpl) GetCurrentUser(
	ctx common.TxCtxInterface,
) (res *cc.GetCurrentUserResponse, err error) {

	res = &cc.GetCurrentUserResponse{}

	res.User, err = ctx.GetUserId()

	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return res, err
}

// ──────────────────────────────────── Invoke ─────────────────────────────────────

func (o ItemContractImpl) AuthorizeOperation(
	ctx common.TxCtxInterface,
	req *cc.AuthorizeOperationRequest,
) (res *cc.AuthorizeOperationResponse, err error) {

	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	authorized, err := ctx.Authorize([]*authpb.Operation{req.GetOperation()})
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	return &cc.AuthorizeOperationResponse{
		Authorized: authorized,
	}, err
}

func (o ItemContractImpl) Bootstrap(
	ctx common.TxCtxInterface,
	req *cc.BootstrapRequest,
) (res *cc.BootstrapResponse, err error) {
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	// TODO: implement bootstrap

	return &cc.BootstrapResponse{}, oops.Errorf("not implemented")
}

// ════════════════════════════════════ Item ═════════════════════════════════════

// ──────────────────────────────────── Query ──────────────────────────────────────

func (o ItemContractImpl) Get(
	ctx common.TxCtxInterface,
	req *cc.GetRequest,
) (res *cc.GetResponse, err error) {
	var (
		obj common.ItemInterface
		msg *authpb.Item
	)

	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	if obj, err = common.ItemKeyToItem(req.GetKey()); err != nil {
		return nil, oops.Wrap(err)
	}

	if err = state.PrimaryGet(ctx, obj); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	if msg, err = common.PackItem(obj); err != nil {
		ctx.LogError(err)
		return nil, err
	} else {
		return &cc.GetResponse{
			Item: msg,
		}, nil
	}
}

func (o ItemContractImpl) List(
	ctx common.TxCtxInterface,
	req *cc.ListRequest,
) (res *cc.ListResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	item, err := common.ItemKeyToItemType(req.GetKey())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	list, mk, err := state.PrimaryList(ctx, item, req.GetBookmark())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	res = &cc.ListResponse{
		Bookmark: mk,
		Items:    []*authpb.Item{},
	}

	if res.Items, err = common.ListItemToProtos(list); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return res, err
}

func (o ItemContractImpl) ListByCollection(
	ctx common.TxCtxInterface,
	req *cc.ListByCollectionRequest,
) (res *cc.ListByCollectionResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	item, err := common.ItemKeyToItemType(req.GetKey())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	item.SetKey(&authpb.ItemKey{CollectionId: req.GetKey().GetCollectionId()})

	list, mk, err := state.PrimaryList(ctx, item, req.GetBookmark())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	res = &cc.ListByCollectionResponse{
		Bookmark: mk,
		Items:    []*authpb.Item{},
	}

	if res.Items, err = common.ListItemToProtos(list); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return res, err
}

func (o ItemContractImpl) ListByAttrs(
	ctx common.TxCtxInterface,
	req *cc.ListByAttrsRequest,
) (res *cc.ListByAttrsResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	item, err := common.ItemKeyToItem(req.GetKey())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	list, mk, err := state.PrimaryByPartialKey(ctx, item, int(req.GetNumAttrs()), req.GetBookmark())
	if err != nil {
		ctx.LogError(err)

		return nil, oops.Wrap(err)
	}

	res = &cc.ListByAttrsResponse{
		Bookmark: mk,
		Items:    []*authpb.Item{},
	}

	if res.Items, err = common.ListItemToProtos(list); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return res, err
}

// ──────────────────────────────────── Invoke ─────────────────────────────────────

func (o ItemContractImpl) Create(
	ctx common.TxCtxInterface,
	req *cc.CreateRequest,
) (res *cc.CreateResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	// Get the item from the request
	obj, err := common.UnPackItem(req.GetItem())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	// Check if the item is a valid item for the collection??

	if err = state.PrimaryCreate(ctx, obj); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return &cc.CreateResponse{
		Item: req.GetItem(),
	}, err
}

func (o ItemContractImpl) Update(
	ctx common.TxCtxInterface,
	req *cc.UpdateRequest,
) (res *cc.UpdateResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	// Get the item from the request
	obj, err := common.UnPackItem(req.GetItem())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	updated, err := state.PrimaryUpdate(ctx, obj, req.GetUpdateMask())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	if item, err := common.PackItem(updated); err != nil {
		return nil, oops.Wrap(err)
	} else {
		res = &cc.UpdateResponse{
			Item: item,
		}
	}

	return res, err
}

func (o ItemContractImpl) Delete(
	ctx common.TxCtxInterface,
	req *cc.DeleteRequest,
) (res *cc.DeleteResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	// Get the item from the request
	obj, err := common.ItemKeyToItem(req.GetKey())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	if err = state.PrimaryDelete(ctx, obj); err != nil {
		return nil, oops.Wrap(err)
	}

	item, err := common.PackItem(obj)
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return &cc.DeleteResponse{
		Item: item,
	}, err
}

// ════════════════════════════════════ History ════════════════════════════════════
// ──────────────────────────────────── Query ──────────────────────────────────────

func (o ItemContractImpl) GetHistory(
	ctx common.TxCtxInterface,
	req *cc.GetHistoryRequest,
) (res *cc.GetHistoryResponse, err error) {
	var (
		obj common.ItemInterface
		h   *authpb.History
	)

	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	if obj, err = common.ItemKeyToItem(req.GetKey()); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	if h, err = state.GetHistory(ctx, obj); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return &cc.GetHistoryResponse{
		Key:     req.GetKey(),
		History: h,
	}, nil
}

func (o ItemContractImpl) GetHiddenTx(
	ctx common.TxCtxInterface,
	req *cc.GetHiddenTxRequest,
) (res *cc.GetHiddenTxResponse, err error) {
	var (
		obj  common.ItemInterface
		hTxs *authpb.HiddenTxList
	)

	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	if obj, err = common.UnPackItem(req.GetItem()); err != nil {
		return nil, oops.Wrap(err)
	}

	if hTxs, err = state.GetHiddenTx(ctx, obj); err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.GetHiddenTxResponse{
		CollectionId: obj.ItemKey().GetCollectionId(),
		HiddenTxs:    hTxs.GetTxs(),
	}, nil
}

// ──────────────────────────────────── Invoke ─────────────────────────────────────

func (o ItemContractImpl) HideTx(
	ctx common.TxCtxInterface,
	req *cc.HideTxRequest,
) (res *cc.HideTxResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	// Get the obj from the request

	obj, err := common.ItemKeyToItem(req.GetKey())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	list, err := state.HideTransaction(ctx, obj, req.GetHiddenTx())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	// Get the obj from the request from world state to return it

	return &cc.HideTxResponse{
		Key:       req.GetKey(),
		HiddenTxs: list,
	}, err
}

func (o ItemContractImpl) UnHideTx(
	ctx common.TxCtxInterface,
	req *cc.UnHideTxRequest,
) (res *cc.UnHideTxResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	// Get the item from the request
	//obj, err := common.UnPackItem(req.GetItem())
	obj, err := common.ItemKeyToItem(req.GetKey())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	list, err := state.UnHideTransaction(ctx, obj, req.GetTxId())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return &cc.UnHideTxResponse{
		Key:       req.GetKey(),
		HiddenTxs: list,
	}, err
}

// ════════════════════════════════════ Suggestions ════════════════════════════════
// ──────────────────────────────────── Query ──────────────────────────────────────

func (o ItemContractImpl) GetSuggestion(
	ctx common.TxCtxInterface,
	req *cc.GetSuggestionRequest,
) (res *cc.GetSuggestionResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	sug := &authpb.Suggestion{
		PrimaryKey:   req.GetItemKey(),
		SuggestionId: req.GetSuggestionId(),
	}

	if err = state.GetSuggestion(ctx, sug); err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.GetSuggestionResponse{
		Suggestion: sug,
	}, nil
}

func (o ItemContractImpl) SuggestionListByCollection(
	ctx common.TxCtxInterface,
	req *cc.SuggestionListByCollectionRequest,
) (res *cc.SuggestionListByCollectionResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	list, mk, err := state.SuggestionListByCollection(ctx, req.GetCollectionId(), req.GetBookmark())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	res = &cc.SuggestionListByCollectionResponse{
		Bookmark:    mk,
		Suggestions: list,
	}

	return res, nil
}

func (o ItemContractImpl) SuggestionByPartialKey(
	ctx common.TxCtxInterface,
	req *cc.SuggestionByPartialKeyRequest,
) (res *cc.SuggestionByPartialKeyResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
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
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	res = &cc.SuggestionByPartialKeyResponse{
		Bookmark:    mk,
		Suggestions: list,
	}

	return res, nil
}

// ──────────────────────────────── Invoke ─────────────────────────────────────────

func (o ItemContractImpl) SuggestionCreate(
	ctx common.TxCtxInterface,
	req *cc.SuggestionCreateRequest,
) (res *cc.SuggestionCreateResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	if err = state.SuggestionCreate(ctx, req.GetSuggestion()); err != nil {
		ctx.GetLogger().Warn("SuggestionCreate", "err", err)
		return nil, oops.Wrap(err)
	}

	return &cc.SuggestionCreateResponse{
		Suggestion: req.GetSuggestion(),
	}, nil
}

func (o ItemContractImpl) SuggestionDelete(
	ctx common.TxCtxInterface,
	req *cc.SuggestionDeleteRequest,
) (res *cc.SuggestionDeleteResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	sug := &authpb.Suggestion{
		PrimaryKey:   req.GetItemKey(),
		SuggestionId: req.GetSuggestionId(),
	}

	if err = state.SuggestionDelete(ctx, sug); err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.SuggestionDeleteResponse{
		Suggestion: sug,
	}, nil
}

func (o ItemContractImpl) SuggestionApprove(
	ctx common.TxCtxInterface,
	req *cc.SuggestionApproveRequest,
) (res *cc.SuggestionApproveResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	sug := &authpb.Suggestion{
		PrimaryKey:   req.GetItemKey(),
		SuggestionId: req.GetSuggestionId(),
	}

	u, err := state.SuggestionApprove(ctx, sug)
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	if updated, err := common.PackItem(*u); err != nil {
		return nil, oops.Wrap(err)
	} else {
		res = &cc.SuggestionApproveResponse{
			Item:       updated,
			Suggestion: sug,
		}
	}

	return res, nil
}
