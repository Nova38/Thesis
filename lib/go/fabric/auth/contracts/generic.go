package contracts

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	// "github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	cc "github.com/nova38/thesis/lib/go/gen/chaincode/auth/common"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
)

type ObjectContractImpl struct {
	contractapi.Contract
	cc.GenericServiceBase
}

// ================================== Object =================================
// ---------------------------------- Query ----------------------------------

func (o ObjectContractImpl) Get(
	ctx state.TxCtxInterface,
	req *cc.GetRequest,
) (res *cc.GetResponse, err error) {
	var (
		obj state.Object
		msg *authpb.Object
	)

	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	if obj, err = state.AuthObjToObject(req.GetObject()); err != nil {
		return nil, oops.Wrap(err)
	}

	if err = state.Get(ctx, obj); err != nil {
		return nil, err
	}

	if msg, err = state.ObjectToAuthObj(obj); err != nil {
		return nil, err
	} else {
		return &cc.GetResponse{
			Object: msg,
		}, nil
	}
}

func (o ObjectContractImpl) List(
	ctx state.TxCtxInterface,
	req *cc.ListRequest,
) (res *cc.ListResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	obj, err := state.AuthObjToObject(req.GetObject())
	if err != nil {
		return nil, oops.Wrap(err)
	}
	proto.Reset(obj) // reset the object to its default state, so that we can get the full list

	list, mk, err := state.List(ctx, obj, req.GetBookmark())

	res = &cc.ListResponse{
		Bookmark: mk,
		Objects:  []*authpb.Object{},
	}

	//for _, o := range list {
	//	var msg *authpb.Object
	//	if msg, err = common.ObjectToAuthObj(o); err != nil {
	//		return nil, err
	//	}
	//	res.Objects = append(res.Objects, msg)
	//}
	if res.Objects, err = state.ListObjectToAuthObjs(list); err != nil {
		return nil, oops.Wrap(err)
	}

	return res, err
}

func (o ObjectContractImpl) ListByCollection(
	ctx state.TxCtxInterface,
	req *cc.ListByCollectionRequest,
) (res *cc.ListByCollectionResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	obj, err := state.AuthObjToObject(req.GetObject())
	if err != nil {
		return nil, oops.Wrap(err)
	}
	// proto.Reset(obj) // reset the object to its default state, so that we can get the full list

	list, mk, err := state.List(ctx, obj, req.GetBookmark())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	res = &cc.ListByCollectionResponse{
		Bookmark: mk,
		Objects:  []*authpb.Object{},
	}

	if res.Objects, err = state.ListObjectToAuthObjs(list); err != nil {
		return nil, oops.Wrap(err)
	}

	return res, err
}

func (o ObjectContractImpl) ListByAttrs(
	ctx state.TxCtxInterface,
	req *cc.ListByAttrsRequest,
) (res *cc.ListByAttrsResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	obj, err := state.AuthObjToObject(req.GetObject())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	list, mk, err := state.List(ctx, obj, req.GetBookmark())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	res = &cc.ListByAttrsResponse{
		Bookmark: mk,
		Objects:  []*authpb.Object{},
	}

	if res.Objects, err = state.ListObjectToAuthObjs(list); err != nil {
		return nil, oops.Wrap(err)
	}

	return res, err
}

func (o ObjectContractImpl) History(
	ctx state.TxCtxInterface,
	req *cc.HistoryRequest,
) (res *cc.HistoryResponse, err error) {
	var (
		obj state.Object
		h   *authpb.History
	)

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	if obj, err = state.AuthObjToObject(req.GetObject()); err != nil {
		return nil, oops.Wrap(err)
	}

	if h, err = state.History(ctx, obj); err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.HistoryResponse{
		History: h,
	}, nil
}

func (o ObjectContractImpl) HiddenTx(
	ctx state.TxCtxInterface,
	req *cc.HiddenTxRequest,
) (res *cc.HiddenTxResponse, err error) {
	var (
		obj  state.Object
		hTxs *authpb.HiddenTxList
	)

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	if obj, err = state.AuthObjToObject(req.GetObject()); err != nil {
		return nil, oops.Wrap(err)
	}

	if hTxs, err = state.HiddenTx(ctx, obj); err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.HiddenTxResponse{
		CollectionId: obj.GetCollectionId(),
		HiddenTxs:    hTxs.Txs,
	}, nil
}

// ------------------------------- Invoke ------------------------------------

func (o ObjectContractImpl) Create(
	ctx state.TxCtxInterface,
	req *cc.CreateRequest,
) (res *cc.CreateResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Get the object from the request
	obj, err := state.AuthObjToObject(req.GetObject())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	err = state.Create(ctx, obj)

	return &cc.CreateResponse{
		Object: req.GetObject(),
	}, err
}

func (o ObjectContractImpl) Update(
	ctx state.TxCtxInterface,
	req *cc.UpdateRequest,
) (res *cc.UpdateResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Get the object from the request
	obj, err := state.AuthObjToObject(req.GetObject())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	err = state.Update(ctx, obj, req.GetUpdateMask())

	return &cc.UpdateResponse{
		Object: req.GetObject(),
	}, err
}

func (o ObjectContractImpl) Delete(
	ctx state.TxCtxInterface,
	req *cc.DeleteRequest,
) (res *cc.DeleteResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Get the object from the request
	obj, err := state.AuthObjToObject(req.GetObject())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	err = state.Delete(ctx, obj)

	return &cc.DeleteResponse{
		Object: req.GetObject(),
	}, err
}

func (o ObjectContractImpl) HideTx(
	ctx state.TxCtxInterface,
	req *cc.HideTxRequest,
) (res *cc.HideTxResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Get the object from the request
	obj, err := state.AuthObjToObject(req.GetObject())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	list, err := state.HideTransaction(ctx, obj, req.GetHiddenTx())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.HideTxResponse{
		Object:    req.GetObject(),
		HiddenTxs: list,
	}, err
}

func (o ObjectContractImpl) UnHideTx(
	ctx state.TxCtxInterface,
	req *cc.UnHideTxRequest,
) (res *cc.UnHideTxResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Get the object from the request
	obj, err := state.AuthObjToObject(req.GetObject())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	list, err := state.UnHideTransaction(ctx, obj, req.GetTxId())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.UnHideTxResponse{
		Object:    req.GetObject(),
		HiddenTxs: list,
	}, err
}

// ================================ Suggestions ================================
// -------------------------------- Query --------------------------------------

func (o ObjectContractImpl) Suggestion(
	ctx state.TxCtxInterface,
	req *cc.SuggestionRequest,
) (res *cc.SuggestionResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (o ObjectContractImpl) SuggestionList(
	ctx state.TxCtxInterface,
	req *cc.SuggestionListRequest,
) (res *cc.SuggestionListResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (o ObjectContractImpl) SuggestionListByCollection(
	ctx state.TxCtxInterface,
	req *cc.SuggestionListByCollectionRequest,
) (res *cc.SuggestionListByCollectionResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (o ObjectContractImpl) SuggestionByPartialKey(
	ctx state.TxCtxInterface,
	req *cc.SuggestionByPartialKeyRequest,
) (res *cc.SuggestionByPartialKeyResponse, err error) {
	// TODO implement me
	panic("implement me")
}

// ---------------------------------- Invoke -----------------------------------

func (o ObjectContractImpl) SuggestionCreate(
	ctx state.TxCtxInterface,
	req *cc.SuggestionCreateRequest,
) (res *cc.SuggestionCreateResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (o ObjectContractImpl) SuggestionDelete(
	ctx state.TxCtxInterface,
	req *cc.SuggestionDeleteRequest,
) (res *cc.SuggestionDeleteResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (o ObjectContractImpl) SuggestionApprove(
	ctx state.TxCtxInterface,
	req *cc.SuggestionApproveRequest,
) (res *cc.SuggestionApproveResponse, err error) {
	// TODO implement me
	panic("implement me")
}
