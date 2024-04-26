package contract

import (
	"github.com/nova38/saacs/pkg/saacs-cc/actions"
	"github.com/nova38/saacs/pkg/saacs-cc/common"
	cc "github.com/nova38/saacs/pkg/saacs-protos/saacs/chaincode/v0"
	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"
	"github.com/samber/oops"
)

// ════════════════════════════════════ Item ═════════════════════════════════════

// ──────────────────────────────────── Query ──────────────────────────────────────

func (o ContractImpl) Get(
	ctx common.TxCtxInterface,
	req *cc.GetRequest,
) (res *cc.GetResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	var (
		obj common.ItemInterface
		msg *pb.Item
	)

	// Validate the request

	if obj, err = common.ItemKeyToItem(req.GetKey()); err != nil {
		ctx.LogError(err)
		return nil, oops.Hint("Failed Convert To an Object").
			With("ItemKind", req.GetKey().GetItemKind()).
			Wrap(err)
	}

	if err = actions.PrimaryGet(ctx, obj); err != nil {
		ctx.LogError(err)
		return nil, oops.In("Contract").Wrap(err)
	}

	if msg, err = common.PackItem(obj); err != nil {
		ctx.LogError(err)
		return nil, err
	} else {
		ctx.GetLogger().Debug("Get Item")
		return &cc.GetResponse{
			Item: msg,
		}, nil
	}

}

func (o ContractImpl) GetFull(
	ctx common.TxCtxInterface,
	req *cc.GetFullRequest,
) (res *cc.GetFullResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	var (
		obj common.ItemInterface
	)

	// Validate the request

	if obj, err = common.ItemKeyToItem(req.GetKey()); err != nil {
		return nil, oops.Wrap(err)
	}

	if req.GetHistoryOptions() == nil {
		req.HistoryOptions = &pb.HistoryOptions{
			Hidden: &pb.HiddenOptions{
				Include: false,
				MspIds:  []string{},
			},
		}
	} else if req.GetHistoryOptions().GetHidden() == nil {
		req.HistoryOptions.Hidden = &pb.HiddenOptions{
			Include: false,
			MspIds:  []string{},
		}

	}

	full, err := actions.PrimaryGetFull(ctx, obj, req.GetHistoryOptions().Include)
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return &cc.GetFullResponse{
		Item:        full.GetItem(),
		History:     full.GetHistory(),
		Suggestions: full.GetSuggestions(),
	}, nil

}

func (o ContractImpl) List(
	ctx common.TxCtxInterface,
	req *cc.ListRequest,
) (res *cc.ListResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	pagination := common.MaybePagation(req.GetPagination())

	if pagination.GetPageSize() != 0 {
		ctx.GetLogger().Info("Limit is not 0", "limit", pagination.GetPageSize())
		ctx.SetPageSize(pagination.GetPageSize())
	}

	item, err := common.ItemKeyToItemType(req.GetKey())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	list, mk, err := actions.PrimaryList(
		ctx,
		item,
		pagination.GetBookmark(),
	)
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	res = &cc.ListResponse{
		Pagination: &pb.Pagination{
			Bookmark: mk,
			PageSize: pagination.GetPageSize(),
		},
		Items: []*pb.Item{},
	}

	if res.Items, err = common.ListItemToProtos(list); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return res, err
}

func (o ContractImpl) ListByAttrs(
	ctx common.TxCtxInterface,
	req *cc.ListByAttrsRequest,
) (res *cc.ListByAttrsResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	// Validate the request
	pagination := common.MaybePagation(req.GetPagination())

	if pagination.GetPageSize() != 0 {
		ctx.GetLogger().Info("Limit is not 0", "limit", pagination.GetPageSize())
		ctx.SetPageSize(pagination.GetPageSize())
	}

	item, err := common.ItemKeyToItem(req.GetKey())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	list, mk, err := actions.PrimaryByPartialKey(
		ctx,
		item,
		int(req.GetNumAttrs()),
		pagination.GetBookmark(),
	)
	if err != nil {
		ctx.LogError(err)

		return nil, oops.Wrap(err)
	}

	res = &cc.ListByAttrsResponse{
		Pagination: &pb.Pagination{
			Bookmark: mk,
			PageSize: pagination.GetPageSize(),
		},
		Items: []*pb.Item{},
	}

	if res.Items, err = common.ListItemToProtos(list); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return res, err
}

// ──────────────────────────────────── Invoke ─────────────────────────────────────

func (o ContractImpl) Create(
	ctx common.TxCtxInterface,
	req *cc.CreateRequest,
) (res *cc.CreateResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	// Validate the request

	// Get the item from the request
	obj, err := common.UnPackItem(req.GetItem())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	// Check if the item is a valid item for the collection??

	if err = actions.PrimaryCreate(ctx, obj); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	item, err := common.PackItem(obj)
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return &cc.CreateResponse{
		Item: item,
	}, err
}

func (o ContractImpl) Update(
	ctx common.TxCtxInterface,
	req *cc.UpdateRequest,
) (res *cc.UpdateResponse, err error) {
	// Validate the request

	// Get the item from the request
	obj, err := common.UnPackItem(req.GetItem())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	updated, err := actions.PrimaryUpdate(ctx, obj, req.GetUpdateMask())
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

func (o ContractImpl) Delete(
	ctx common.TxCtxInterface,
	req *cc.DeleteRequest,
) (res *cc.DeleteResponse, err error) {
	// Validate the request

	// Get the item from the request
	obj, err := common.ItemKeyToItem(req.GetKey())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	if err = actions.PrimaryDelete(ctx, obj); err != nil {
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

func (o ContractImpl) GetHistory(
	ctx common.TxCtxInterface,
	req *cc.GetHistoryRequest,
) (res *cc.GetHistoryResponse, err error) {
	var (
		obj     common.ItemInterface
		h       *pb.History
		options *pb.HistoryOptions
	)

	if obj, err = common.ItemKeyToItem(req.GetKey()); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	if req.GetHistoryOptions() == nil {
		req.HistoryOptions = &pb.HiddenOptions{}
	}

	if h, err = actions.GetHistory(ctx, obj, options); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return &cc.GetHistoryResponse{
		Key:     req.GetKey(),
		History: h,
	}, nil
}

func (o ContractImpl) GetHiddenTx(
	ctx common.TxCtxInterface,
	req *cc.GetHiddenTxRequest,
) (res *cc.GetHiddenTxResponse, err error) {
	var (
		obj  common.ItemInterface
		hTxs *pb.HiddenTxList
	)

	if obj, err = common.UnPackItem(req.GetItem()); err != nil {
		return nil, oops.Wrap(err)
	}

	if hTxs, err = actions.GetHiddenTx(ctx, obj); err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.GetHiddenTxResponse{
		CollectionId: obj.ItemKey().GetCollectionId(),
		HiddenTxs:    hTxs.GetTxs(),
	}, nil
}

// ──────────────────────────────────── Invoke ─────────────────────────────────────

func (o ContractImpl) HideTx(
	ctx common.TxCtxInterface,
	req *cc.HideTxRequest,
) (res *cc.HideTxResponse, err error) {

	// Get the obj from the request

	obj, err := common.ItemKeyToItem(req.GetKey())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	list, err := actions.HideTransaction(ctx, obj, req.GetHiddenTx())
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

func (o ContractImpl) UnHideTx(
	ctx common.TxCtxInterface,
	req *cc.UnHideTxRequest,
) (res *cc.UnHideTxResponse, err error) {

	// Get the item from the request
	// obj, err := common.UnPackItem(req.GetItem())
	obj, err := common.ItemKeyToItem(req.GetKey())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	list, err := actions.UnHideTransaction(ctx, obj, req.GetTxId())
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

func (o ContractImpl) GetSuggestion(
	ctx common.TxCtxInterface,
	req *cc.GetSuggestionRequest,
) (res *cc.GetSuggestionResponse, err error) {
	// Validate the request

	sug := &pb.Suggestion{
		PrimaryKey:   req.GetItemKey(),
		SuggestionId: req.GetSuggestionId(),
	}

	if err = actions.GetSuggestion(ctx, sug); err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.GetSuggestionResponse{
		Suggestion: sug,
	}, nil
}

func (o ContractImpl) SuggestionListByCollection(
	ctx common.TxCtxInterface,
	req *cc.SuggestionListByCollectionRequest,
) (res *cc.SuggestionListByCollectionResponse, err error) {
	// Validate the request
	pagination := common.MaybePagation(req.GetPagination())

	list, mk, err := actions.SuggestionListByCollection(
		ctx,
		req.GetCollectionId(),
		pagination.GetBookmark(),
	)
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	res = &cc.SuggestionListByCollectionResponse{
		Pagination: &pb.Pagination{
			PageSize: pagination.GetPageSize(),
			Bookmark: mk,
		},
		Suggestions: list,
	}

	return res, nil
}

func (o ContractImpl) SuggestionByPartialKey(
	ctx common.TxCtxInterface,
	req *cc.SuggestionByPartialKeyRequest,
) (res *cc.SuggestionByPartialKeyResponse, err error) {
	// Validate the request
	pagination := common.MaybePagation(req.GetPagination())

	sug := &pb.Suggestion{
		PrimaryKey: req.GetItemKey(),
	}

	list, mk, err := actions.PartialSuggestionList(
		ctx,
		sug,
		int(req.GetNumAttrs()),
		pagination.GetBookmark(),
	)
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	res = &cc.SuggestionByPartialKeyResponse{
		Pagination: &pb.Pagination{
			PageSize: pagination.GetPageSize(),
			Bookmark: mk,
		},
		Suggestions: list,
	}

	return res, nil
}

func (o ContractImpl) SuggestionListByItem(
	ctx common.TxCtxInterface,
	req *cc.SuggestionListByItemRequest,
) (res *cc.SuggestionListByItemResponse, err error) {
	// Validate the request
	pagination := common.MaybePagation(req.GetPagination())

	list, mk, err := actions.SuggestionListByItem(
		ctx,
		req.GetItemKey(),
		pagination.GetBookmark(),
	)
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	res = &cc.SuggestionListByItemResponse{
		Pagination: &pb.Pagination{
			PageSize: pagination.GetPageSize(),
			Bookmark: mk,
		},
		Suggestions: list,
	}

	return res, nil
}

// ──────────────────────────────── Invoke ─────────────────────────────────────────

func (o ContractImpl) SuggestionCreate(
	ctx common.TxCtxInterface,
	req *cc.SuggestionCreateRequest,
) (res *cc.SuggestionCreateResponse, err error) {
	// Validate the request

	if err = actions.SuggestionCreate(ctx, req.GetSuggestion()); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return &cc.SuggestionCreateResponse{
		Suggestion: req.GetSuggestion(),
	}, nil
}

func (o ContractImpl) SuggestionDelete(
	ctx common.TxCtxInterface,
	req *cc.SuggestionDeleteRequest,
) (res *cc.SuggestionDeleteResponse, err error) {
	// Validate the request

	sug := &pb.Suggestion{
		PrimaryKey:   req.GetItemKey(),
		SuggestionId: req.GetSuggestionId(),
	}

	if err = actions.SuggestionDelete(ctx, sug); err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.SuggestionDeleteResponse{
		Suggestion: sug,
	}, nil
}

func (o ContractImpl) SuggestionApprove(
	ctx common.TxCtxInterface,
	req *cc.SuggestionApproveRequest,
) (res *cc.SuggestionApproveResponse, err error) {
	// Validate the request

	sug := &pb.Suggestion{
		PrimaryKey:   req.GetItemKey(),
		SuggestionId: req.GetSuggestionId(),
	}

	u, err := actions.SuggestionApprove(ctx, sug)
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
