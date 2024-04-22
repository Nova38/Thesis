package base

import (
	"encoding/json"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nova38/saacs/pkg/saacs-cc/common"
	actions2 "github.com/nova38/saacs/pkg/saacs-cc/state/actions"
	authpb "github.com/nova38/saacs/pkg/saacs-protos/auth/v1"
	cc "github.com/nova38/saacs/pkg/saacs-protos/chaincode/common"
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

func (o ItemContractImpl) GetCollectionsList(
	ctx common.TxCtxInterface,
) (res *cc.GetCollectionsListResponse, err error) {

	ctx.GetLogger().Info("GetCollectionsList")

	col := &authpb.Collection{}

	results, err := ctx.GetStub().
		GetStateByPartialCompositeKey(
			col.ItemType(),
			[]string{},
		)
	if err != nil {
		ctx.GetLogger().Info("error")

		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	ctx.GetLogger().Info("No error")
	defer func(results shim.StateQueryIteratorInterface) {
		err := results.Close()
		if err != nil {
			ctx.GetLogger().Error("GetCollectionsList", "Error", err)
		}
	}(results)

	res = &cc.GetCollectionsListResponse{
		Collections: []*authpb.Collection{},
	}

	for results.HasNext() {

		tmp := &authpb.Collection{}

		queryResponse, err := results.Next()
		if err != nil || queryResponse == nil {
			return nil, oops.Wrapf(err, "Error getting next item")
		}

		if err = json.Unmarshal(queryResponse.GetValue(), tmp); err != nil {
			return nil, oops.Wrap(err)
		}

		res.Collections = append(res.GetCollections(), tmp)
	}

	// list, _, err := state.GetPartialKeyList(ctx, col, 0, "")
	// if err != nil {
	// 	ctx.LogError(err)
	// 	return nil, oops.Wrap(err)
	// }

	return res, nil
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

	if err = actions2.PrimaryGet(ctx, obj); err != nil {
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

func (o ItemContractImpl) GetFull(
	ctx common.TxCtxInterface,
	req *cc.GetFullRequest,
) (res *cc.GetFullResponse, err error) {

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

	full, err := actions2.PrimaryGetFull(ctx, item, req.GetShowHidden())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return &cc.GetFullResponse{FullItem: full}, nil

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

	if req.GetLimit() != 0 {
		ctx.GetLogger().Info("Limit is not 0", "limit", req.GetLimit())
		ctx.SetPageSize(int32(req.GetLimit()))
	}

	item, err := common.ItemKeyToItemType(req.GetKey())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	list, mk, err := actions2.PrimaryList(ctx, item, req.GetBookmark())
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

	if req.GetLimit() != 0 {
		ctx.GetLogger().Info("Limit is not 0", "limit", req.GetLimit())
		ctx.SetPageSize(int32(req.GetLimit()))
	}

	item, err := common.ItemKeyToItemType(req.GetKey())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	item.SetKey(&authpb.ItemKey{CollectionId: req.GetKey().GetCollectionId()})

	// list, mk, err := actions.PrimaryList(ctx, item, req.GetBookmark())
	list, mk, err := actions2.PrimaryByPartialKey(
		ctx,
		item,
		int(0),
		req.GetBookmark(),
	)

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

	if req.GetLimit() != 0 {
		ctx.GetLogger().Info("Limit is not 0", "limit", req.GetLimit())
		ctx.SetPageSize(int32(req.GetLimit()))
	}

	item, err := common.ItemKeyToItem(req.GetKey())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	list, mk, err := actions2.PrimaryByPartialKey(
		ctx,
		item,
		int(req.GetNumAttrs()),
		req.GetBookmark(),
	)
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

	if err = actions2.PrimaryCreate(ctx, obj); err != nil {
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

	updated, err := actions2.PrimaryUpdate(ctx, obj, req.GetUpdateMask())
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

	if err = actions2.PrimaryDelete(ctx, obj); err != nil {
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

	if h, err = actions2.GetHistory(ctx, obj); err != nil {
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

	if hTxs, err = actions2.GetHiddenTx(ctx, obj); err != nil {
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

	list, err := actions2.HideTransaction(ctx, obj, req.GetHiddenTx())
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
	// obj, err := common.UnPackItem(req.GetItem())
	obj, err := common.ItemKeyToItem(req.GetKey())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	list, err := actions2.UnHideTransaction(ctx, obj, req.GetTxId())
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

	if err = actions2.GetSuggestion(ctx, sug); err != nil {
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
	list, mk, err := actions2.SuggestionListByCollection(
		ctx,
		req.GetCollectionId(),
		req.GetBookmark(),
	)
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

	list, mk, err := actions2.PartialSuggestionList(
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
	if err = actions2.SuggestionCreate(ctx, req.GetSuggestion()); err != nil {
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

	if err = actions2.SuggestionDelete(ctx, sug); err != nil {
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

	u, err := actions2.SuggestionApprove(ctx, sug)
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	if updated, err := common.PackItem(u); err != nil {
		return nil, oops.Wrap(err)
	} else {
		res = &cc.SuggestionApproveResponse{
			Item:       updated,
			Suggestion: sug,
		}
	}

	return res, nil
}
