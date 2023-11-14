package common

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"

	"github.com/samber/oops"

	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	cc "github.com/nova38/thesis/lib/go/gen/chaincode/auth/common"
)

type CollectionImpl struct {
	contractapi.Contract
	cc.CollectionServiceBase
}

func (a CollectionImpl) CollectionGet(
	ctx state.TxCtxInterface,
	req *cc.CollectionGetRequest,
) (res *cc.CollectionGetResponse, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	// Validate the request
	err = ctx.Validate(req)
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(authpb.TxError_REQUEST_INVALID.String()).
			Wrap(err)
	}

	col := &authpb.Collection{CollectionId: req.GetCollectionId()}

	err = state.Get(ctx, col)

	return &cc.CollectionGetResponse{
		Collection: col,
	}, err
}

func (a CollectionImpl) CollectionGetList(
	ctx state.TxCtxInterface,
	req *cc.CollectionGetListRequest,
) (res *cc.CollectionGetListResponse, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	list, mk, err := state.List(
		ctx,
		&authpb.Collection{},
		req.GetBookmark(),
	)

	return &cc.CollectionGetListResponse{
		Collections: list,
		Bookmark:    mk,
	}, err
}

func (a CollectionImpl) CollectionGetHistory(
	ctx state.TxCtxInterface,
	req *cc.CollectionGetHistoryRequest,
) (res *cc.CollectionGetHistoryResponse, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	col := &authpb.Collection{CollectionId: req.GetCollectionId()}

	if err := state.Get(ctx, col); err != nil {
		return nil, oops.Wrap(err)
	}

	h, err := state.History(ctx, col)
	if err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.CollectionGetHistoryResponse{
		Collection: col,
		History:    h,
	}, err
}

// -------------------- Invoke ------------------------

func (a CollectionImpl) CollectionCreate(
	ctx state.TxCtxInterface,
	req *cc.CollectionCreateRequest,
) (res *cc.CollectionCreateResponse, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Create the collection
	err = state.Create(ctx, req.GetCollection())

	return &cc.CollectionCreateResponse{Collection: req.GetCollection()}, err
}

func (a CollectionImpl) CollectionUpdate(
	ctx state.TxCtxInterface,
	req *cc.CollectionUpdateRequest,
) (res *cc.CollectionUpdateResponse, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}
	col := req.GetCollection()

	err = state.Update(ctx, col, nil)

	return &cc.CollectionUpdateResponse{
		Collection: col,
	}, err
}

func (a CollectionImpl) CollectionHideTx(
	ctx state.TxCtxInterface,
	req *cc.CollectionHideTxRequest,
) (res *cc.CollectionHideTxResponse, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	hidden := req.GetHiddenTx()
	if hidden == nil {
		return nil, oops.Wrap(common.RequestInvalid)
	}

	mspId, userId, err := ctx.GetUserId()
	if err != nil {
		return nil, oops.Wrap(err)
	}

	hidden.MspId = mspId
	hidden.UserId = userId

	col := &authpb.Collection{CollectionId: req.GetCollectionId()}

	res = &cc.CollectionHideTxResponse{
		CollectionId: req.GetCollectionId(),
	}

	return res, state.HideTransaction(ctx, col, hidden)
}
