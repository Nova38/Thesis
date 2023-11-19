package contracts

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"

	"github.com/samber/oops"

	cc "github.com/nova38/thesis/lib/go/gen/chaincode/auth/common"
)

type UserImpl struct {
	contractapi.Contract
	cc.CollectionServiceBase
}

// ─────────────────────────────────-- Query ─────────────────────────────────--

func (a UserImpl) UserGetCurrentId(
	ctx state.TxCtxInterface,
) (res *cc.UserGetCurrentIdResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	user, err := ctx.GetUserId()
	if err != nil {
		return nil, oops.
			In("UserGetCurrentId").
			Wrap(err)
	}

	return &cc.UserGetCurrentIdResponse{
		MspId:  user.GetMspId(),
		UserId: user.GetUserId(),
	}, nil
}

func (a UserImpl) UserGetCurrent(
	ctx state.TxCtxInterface,
) (res *cc.UserGetCurrentResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	user, err := ctx.GetUserId()
	if err != nil {
		return nil, err
	}

	err = state.PrimaryGet(ctx, user)
	if err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.UserGetCurrentResponse{User: user}, nil
}

// func (a UserImpl) UserGet(
// 	ctx state.TxCtxInterface,
// 	req *cc.UserGetRequest,
// ) (res *cc.UserGetResponse, err error) {
// 	defer func() { ctx.HandleFnError(&err, recover()) }()

// 	// Validate the request
// 	if err = ctx.Validate(req); err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	// Get the user
// 	user := &authpb.User{
// 		MspId:  req.GetMspId(),
// 		UserId: req.GetUserId(),
// 	}

// 	if err = state.PrimaryGet(ctx, user); err != nil {
// 		return nil, err
// 	}

// 	return &cc.UserGetResponse{User: user}, nil
// }

// func (a UserImpl) UserGetList(
// 	ctx state.TxCtxInterface,
// 	req *cc.UserGetListRequest,
// ) (res *cc.UserGetListResponse, err error) {
// 	defer func() { ctx.HandleFnError(&err, recover()) }()

// 	// Validate the request
// 	if err = ctx.Validate(req); err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	userList, _, err := state.PrimaryList(ctx, &authpb.User{}, "")
// 	if err != nil {
// 		return nil, oops.
// 			In("UserGetList").
// 			Wrap(err)
// 	}

// 	return &cc.UserGetListResponse{Users: userList}, nil
// }

// func (a UserImpl) UserGetHistory(
// 	ctx state.TxCtxInterface,
// 	req *cc.UserGetHistoryRequest,
// ) (res *cc.UserGetHistoryResponse, err error) {
// 	defer func() { ctx.HandleFnError(&err, recover()) }()

// 	if err = ctx.Validate(req); err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	// Get the history of the users
// 	userHistory, err := state.History(ctx, &authpb.User{
// 		MspId:  req.GetMspId(),
// 		UserId: req.GetUserId(),
// 	})
// 	if err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	return &cc.UserGetHistoryResponse{History: userHistory}, nil
// }

// func (a UserImpl) UserGetHiddenTx(
// 	ctx state.TxCtxInterface,
// 	req *cc.UserGetHiddenTxRequest,
// ) (res *cc.UserGetHiddenTxResponse, err error) {
// 	defer func() { ctx.HandleFnError(&err, recover()) }()

// 	if err = ctx.Validate(req); err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	// Get the history of the users
// 	list, err := state.HiddenTx(ctx, &authpb.User{
// 		MspId:  req.GetMspId(),
// 		UserId: req.GetUserId(),
// 	})
// 	if err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	return &cc.UserGetHiddenTxResponse{
// 		MspId:     req.GetMspId(),
// 		UserId:    req.GetUserId(),
// 		HiddenTxs: list,
// 	}, nil
// }

// ─────────────────────────────────- Invoke ─────────────────────────────────--

// func (a UserImpl) UserCreate(
// 	ctx state.TxCtxInterface,
// 	req *cc.UserCreateRequest,
// ) (res *cc.UserCreateResponse, err error) {
// 	defer func() { ctx.HandleFnError(&err, recover()) }()

// 	// Validate the request
// 	if err = ctx.Validate(req); err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	user, err := ctx.GetUserId()
// 	if err != nil {
// 		return nil, oops.
// 			In("UserRegister").
// 			Wrap(err)
// 	}

// 	user.Name = req.GetName()

// 	res = &cc.UserCreateResponse{User: user}

// 	return res, state.PrimaryCreate(ctx, user)
// }

// func (a UserImpl) UserUpdate(
// 	ctx state.TxCtxInterface,
// 	req *cc.UserUpdateRequest,
// ) (res *cc.UserUpdateResponse, err error) {
// 	defer func() { ctx.HandleFnError(&err, recover()) }()

// 	// Validate the request
// 	if err = ctx.Validate(req); err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	user := &authpb.User{
// 		UserId: req.GetUserId(),
// 		MspId:  req.GetMspId(),
// 	}

// 	err = state.PrimaryGet(ctx, user)
// 	if err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	// Update the user
// 	user.Name = req.GetName()

// 	return &cc.UserUpdateResponse{User: user}, state.PrimaryUpdate(ctx, user, nil)
// }

// func (a UserImpl) UserDelete(
// 	ctx state.TxCtxInterface,
// 	req *cc.UserDeleteRequest,
// ) (res *cc.UserDeleteResponse, err error) {
// 	defer func() { ctx.HandleFnError(&err, recover()) }()

// 	// Validate the request
// 	if err = ctx.Validate(req); err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	user := &authpb.User{
// 		UserId: req.GetUserId(),
// 		MspId:  req.GetMspId(),
// 	}

// 	err = state.PrimaryDelete(ctx, user)

// 	return &cc.UserDeleteResponse{User: user}, err
// }

// func (a UserImpl) UserHideTx(
// 	ctx state.TxCtxInterface,
// 	req *cc.UserHideTxRequest,
// ) (res *cc.UserHideTxResponse, err error) {
// 	defer func() { ctx.HandleFnError(&err, recover()) }()

// 	// Validate the request
// 	if err = ctx.Validate(req); err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	user := &authpb.User{
// 		UserId: req.GetUserId(),
// 		MspId:  req.GetMspId(),
// 	}

// 	_, err = state.UnHideTransaction(ctx, user, req.GetHiddenTx().GetTxId())

// 	return &cc.UserHideTxResponse{User: user}, err
// }
