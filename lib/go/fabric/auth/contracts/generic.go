package contracts

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	// "github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
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

// ════════════════════════════════════ Object ═════════════════════════════════════
// ──────────────────────────────────── Query ──────────────────────────────────────

func (o ObjectContractImpl) Get(
	ctx state.TxCtxInterface,
	req *cc.GetRequest,
) (res *cc.GetResponse, err error) {
	var (
		obj common.ObjectInterface
		msg *authpb.Object
	)

	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	if obj, err = state.ProtoToObject(req.GetObject()); err != nil {
		return nil, oops.Wrap(err)
	}

	if err = state.PrimaryGet(ctx, obj); err != nil {
		return nil, err
	}

	if msg, err = state.ObjectToProto(obj); err != nil {
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

	obj, err := state.ProtoToObject(req.GetObject())
	if err != nil {
		return nil, oops.Wrap(err)
	}
	proto.Reset(obj)
	// reset the object to its default state, so that we can get the full list

	list, mk, err := state.PrimaryList(ctx, obj, req.GetBookmark())

	res = &cc.ListResponse{
		Bookmark: mk,
		Objects:  []*authpb.Object{},
	}

	if res.Objects, err = state.ListObjectToProtos(list); err != nil {
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

	obj, err := state.ProtoToObject(req.GetObject())
	if err != nil {
		return nil, oops.Wrap(err)
	}
	// proto.Reset(obj) // reset the object to its default state, so that we can get the full list

	list, mk, err := state.PrimaryList(ctx, obj, req.GetBookmark())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	res = &cc.ListByCollectionResponse{
		Bookmark: mk,
		Objects:  []*authpb.Object{},
	}

	if res.Objects, err = state.ListObjectToProtos(list); err != nil {
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

	obj, err := state.ProtoToObject(req.GetObject())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	list, mk, err := state.PrimaryList(ctx, obj, req.GetBookmark())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	res = &cc.ListByAttrsResponse{
		Bookmark: mk,
		Objects:  []*authpb.Object{},
	}

	if res.Objects, err = state.ListObjectToProtos(list); err != nil {
		return nil, oops.Wrap(err)
	}

	return res, err
}

func (o ObjectContractImpl) History(
	ctx state.TxCtxInterface,
	req *cc.HistoryRequest,
) (res *cc.HistoryResponse, err error) {
	var (
		obj common.ObjectInterface
		h   *authpb.History
	)

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	if obj, err = state.ProtoToObject(req.GetObject()); err != nil {
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
		obj  common.ObjectInterface
		hTxs *authpb.HiddenTxList
	)

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	if obj, err = state.ProtoToObject(req.GetObject()); err != nil {
		return nil, oops.Wrap(err)
	}

	if hTxs, err = state.HiddenTx(ctx, obj); err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.HiddenTxResponse{
		CollectionId: obj.ObjectKey().CollectionId,
		HiddenTxs:    hTxs.Txs,
	}, nil
}

// ──────────────────────────────────── Invoke ─────────────────────────────────────

func (o ObjectContractImpl) Create(
	ctx state.TxCtxInterface,
	req *cc.CreateRequest,
) (res *cc.CreateResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Get the object from the request
	obj, err := state.ProtoToObject(req.GetObject())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	// Check if the object is a valid object for the collection??

	err = state.PrimaryCreate(ctx, obj)

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
	obj, err := state.ProtoToObject(req.GetObject())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	err = state.PrimaryUpdate(ctx, obj, req.GetUpdateMask())

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
	obj, err := state.ProtoToObject(req.GetObject())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	if err = state.PrimaryDelete(ctx, obj); err != nil {
		return nil, oops.Wrap(err)
	}

	object, err := state.ObjectToProto(obj)
	if err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.DeleteResponse{
		Object: object,
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
	obj, err := state.ProtoToObject(req.GetObject())
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
	obj, err := state.ProtoToObject(req.GetObject())
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

// ════════════════════════════════════ Suggestions ════════════════════════════════
// ──────────────────────────────────── Query ──────────────────────────────────────

func (o ObjectContractImpl) Suggestion(
	ctx state.TxCtxInterface,
	req *cc.SuggestionRequest,
) (res *cc.SuggestionResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	sug := &authpb.Suggestion{
		PrimaryKey:   req.GetObjectKey(),
		SuggestionId: req.GetSuggestionId(),
	}

	if err = state.Suggestion(ctx, sug); err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.SuggestionResponse{
		Suggestion: sug,
	}, nil
}

func (o ObjectContractImpl) SuggestionListByCollection(
	ctx state.TxCtxInterface,
	req *cc.SuggestionListByCollectionRequest,
) (res *cc.SuggestionListByCollectionResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	list, mk, err := state.SuggestionListByCollection(ctx, req.GetCollectionId(), req.GetBookmark())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	res = &cc.SuggestionListByCollectionResponse{
		Bookmark:    mk,
		Suggestions: list,
	}

	return res, nil
}

func (o ObjectContractImpl) SuggestionByPartialKey(
	ctx state.TxCtxInterface,
	req *cc.SuggestionByPartialKeyRequest,
) (res *cc.SuggestionByPartialKeyResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}
	sug := &authpb.Suggestion{
		PrimaryKey: req.GetObjectKey(),
	}

	list, mk, err := state.PartialSuggestionList(
		ctx,
		sug,
		int(req.GetNumAttrs()),
		req.GetBookmark(),
	)
	if err != nil {
		return nil, oops.Wrap(err)
	}

	res = &cc.SuggestionByPartialKeyResponse{
		Bookmark:    mk,
		Suggestions: list,
	}

	return res, nil
}

// ---------------------------------- Invoke -----------------------------------

func (o ObjectContractImpl) SuggestionCreate(
	ctx state.TxCtxInterface,
	req *cc.SuggestionCreateRequest,
) (res *cc.SuggestionCreateResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
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

func (o ObjectContractImpl) SuggestionDelete(
	ctx state.TxCtxInterface,
	req *cc.SuggestionDeleteRequest,
) (res *cc.SuggestionDeleteResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	sug := &authpb.Suggestion{
		PrimaryKey:   req.GetObjectKey(),
		SuggestionId: req.GetSuggestionId(),
	}

	if err = state.SuggestionDelete(ctx, sug); err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.SuggestionDeleteResponse{
		Suggestion: sug,
	}, nil
}

func (o ObjectContractImpl) SuggestionApprove(
	ctx state.TxCtxInterface,
	req *cc.SuggestionApproveRequest,
) (res *cc.SuggestionApproveResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	sug := &authpb.Suggestion{
		PrimaryKey:   req.GetObjectKey(),
		SuggestionId: req.GetSuggestionId(),
	}

	u, err := state.SuggestionApprove(ctx, sug)
	if err != nil {
		return nil, oops.Wrap(err)
	}

	updated := &authpb.Object{}

	if updated, err = state.ObjectToProto(*u); err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.SuggestionApproveResponse{
		Object:     updated,
		Suggestion: sug,
	}, nil
}
