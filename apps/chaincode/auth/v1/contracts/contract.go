package contracts

import (
	"log/slog"
	_ "strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/samber/oops"

	"google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/nova38/thesis/lib/go/fabric/rbac"
	"github.com/nova38/thesis/lib/go/fabric/state"

	cc "github.com/nova38/thesis/lib/go/gen/chaincode/rbac/schema/v1"
	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
)

// Check if AuthContractImpl implements AuthServiceInterface
var (
	_ rbac.AuthTxCtxInterface             = (*AuthTxCtx)(nil)
	_ cc.AuthServiceInterface[*AuthTxCtx] = (*AuthContractImpl)(nil)
	_ contractapi.ContractInterface       = (*AuthContractImpl)(nil)
)

type AuthContractImpl struct {
	contractapi.Contract
	cc.AuthServiceBase
}

func (a AuthContractImpl) GetBeforeTransaction() interface{} {
	return a.BeforeTransaction
}

func (a AuthContractImpl) BeforeTransaction(ctx *AuthTxCtx) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.HandelBefore(); err != nil {
		return oops.Wrap(err)
	}

	ops, err := cc.AuthServiceGetTxOperation(ctx.GetFnName())
	if err != nil {
		return oops.
			In("BeforeTransaction").
			With("fn", ctx.GetFnName()).
			Wrap(err)
	}

	if err = ctx.SetOperation(ops); err != nil {
		return oops.
			In("BeforeTransaction").
			With("fn", ctx.GetFnName()).
			Wrap(err)
	}

	return err
}

// UserGetCurrent : Returns the current user.
//
// Returns the current user.
// # Requires:
//   - User submitting the transaction is a registered user.
//
// # Operation:
//   - Domain: DOMAIN_USER
//   - Action: ACTION_VIEW
func (a AuthContractImpl) UserGetCurrent(
	ctx *AuthTxCtx,
) (res *cc.UserGetCurrentResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	user, err := ctx.GetUser()
	if err != nil {
		return nil, err
	}

	return &cc.UserGetCurrentResponse{User: user}, nil
}

// Returns the current user id.
//
// # Requires:
//   - User submitting the transaction is a registered user.
//
// # Operation:
//   - Domain: DOMAIN_USER
//   - Action: ACTION_VIEW
func (a AuthContractImpl) UserGetCurrentId(
	ctx *AuthTxCtx,
) (res *cc.UserGetCurrentIdResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	user_id, err := ctx.GetUserId()
	if err != nil {
		return nil, oops.
			In("UserGetCurrentId").
			Code(rbac_pb.Error_ERROR_USER_INVALID.String()).
			Wrap(err)
	}
	return &cc.UserGetCurrentIdResponse{UserId: user_id}, nil
}

// UserGetList: Returns the list of users.
//
// # Requires:
//   - Non-register users can call this method.
//
// # Operation:
//   - Domain: DOMAIN_USER
//   - Action: ACTION_VIEW
func (a AuthContractImpl) UserGetList(
	ctx *AuthTxCtx,
) (res *cc.UserGetListResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	userList, err := state.GetFullList(ctx, &rbac_pb.User{})
	if err != nil {
		return nil, oops.
			In("UserGetList").
			Code(rbac_pb.Error_ERROR_UNSPECIFIED.String()).
			Wrap(err)
	}

	return &cc.UserGetListResponse{Users: userList}, nil
}

// UserGet: Returns the user.
//
// # Requires:
//   - Non-register users can call this method.
//
// # Operation:
//   - Domain: DOMAIN_USER
//   - Action: ACTION_VIEW
func (a AuthContractImpl) UserGet(
	ctx *AuthTxCtx,
	req *cc.UserGetRequest,
) (res *cc.UserGetResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	// Validate the request
	v, err := ctx.GetValidator()
	if err != nil {
		return nil, oops.Wrap(err)
	}

	if err = v.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Get the user
	user := &rbac_pb.User{Id: req.GetId()}

	if err = state.Get(ctx, user); err != nil {
		return nil, err
	}

	return &cc.UserGetResponse{User: user}, nil
}

// UserGetHistory: Returns the user history.
//
// # Requires:
//   - Non-register users can call this method.
//
// # Operation:
//   - Domain: DOMAIN_USER
//   - Action: ACTION_VIEW_HISTORY
func (a AuthContractImpl) UserGetHistory(
	ctx *AuthTxCtx,
	req *cc.UserGetHistoryRequest,
) (res *cc.UserGetHistoryResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.
			In("UserGetHistory").
			With("req", req).
			Wrap(err)
	}

	// Get the history of the users
	userHistory, err := state.GetHistory(ctx, &rbac_pb.User{Id: req.GetId()})
	if err != nil {
		return nil, oops.
			In("UserGetHistory").
			With("req", req).
			Wrap(err)
	}

	res = &cc.UserGetHistoryResponse{
		UserId:  req.Id,
		History: &rbac_pb.History{},
	}
	for _, v := range userHistory.Entries {
		e, err := anypb.New(v.State)
		if err != nil {
			return nil, oops.Wrap(err)
		}
		res.History.Entries = append(res.History.Entries, &rbac_pb.History_Entry{
			TxId:      v.TxId,
			Timestamp: v.Timestamp,
			IsDeleted: v.IsDelete,
			IsHidden:  v.IsHidden,
			State:     e,
		})
	}

	return res, nil
}

// UserRegister: Registers the user.
//
// # Requires:
//   - The certificate for the user submitting this request must not be already registered as a user.
//
// # Operation:
//   - Domain: DOMAIN_USER
//   - Action: ACTION_CREATE
func (a AuthContractImpl) UserRegister(
	ctx *AuthTxCtx,
	req *cc.UserRegisterRequest,
) (res *cc.UserRegisterResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	user := &rbac_pb.User{Name: req.Name}
	if user.Id, err = ctx.GetUserId(); err != nil {
		return nil, oops.Wrap(err)
	}

	return &cc.UserRegisterResponse{User: &rbac_pb.User{}}, state.Insert(ctx, user)
}

func (a AuthContractImpl) UserUpdateMembership(
	ctx *AuthTxCtx,
	req *cc.UserUpdateMembershipRequest,
) (res *cc.UserUpdateMembershipResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	// Validate the request
	{
		// Validate the request
		if err = ctx.Validate(req); err != nil {
			return nil, oops.Wrap(err)
		}

		// Extract AuthTransactionItems
		if err = ctx.ExtractAuthTransactionItems(req); err != nil {
			return nil, oops.Wrap(err)
		}

		col, err := ctx.GetCollection()
		if err != nil {
			return nil, oops.Wrap(err)
		}

		// Check if the role is valid
		_, ok := col.Roles[req.Role]
		if !ok {
			return nil, oops.
				In("UserUpdateMembership").
				Code(rbac_pb.Error_ERROR_COLLECTION_INVALID_ROLE_ID.String()).
				Errorf("Role %v is not valid for collection %v", req.Role, req.CollectionId)
		}

		// check if the req user is valid
		id, err := ctx.GetUserId()
		if err != nil || id == nil {
			return nil, oops.
				In(ctx.GetFnName()).
				Code(rbac_pb.Error_ERROR_UNSPECIFIED.String()).
				Errorf("User id is nil")
		}

	}

	{ // Authorize the request
		if err = ctx.IsAuthorized(); err != nil {
			return nil, oops.
				In(ctx.GetFnName()).
				Code(rbac_pb.Error_ERROR_USER_PERMISSION_DENIED.String()).
				Wrap(err)
		}
	}

	// -------------------------

	{ // Process the request
		// Get the user to modify
		userToModify := &rbac_pb.User{
			Id: req.GetId(),
		}
		if err = state.Get(ctx, userToModify); err != nil {
			return nil, oops.
				In(ctx.GetFnName()).
				Code(rbac_pb.Error_ERROR_USER_INVALID.String()).
				Wrap(err)
		}

		// Update the user
		userToModify.Roles[req.GetCollectionId().CollectionId] = &rbac_pb.User_Role{
			CollectionId: &rbac_pb.Collection_Id{
				CollectionId: req.CollectionId.CollectionId,
			},
			RoleId:    req.GetRole(),
			GrantedBy: ctx.User.GetId(),
		}

		res = &cc.UserUpdateMembershipResponse{User: userToModify}

		return res, state.Update(ctx, userToModify)

	}
}

func (a AuthContractImpl) CollectionGetList(
	ctx *AuthTxCtx,
) (res *cc.CollectionGetListResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	// Get the collections
	collectionList, err := state.GetFullList(ctx, &rbac_pb.Collection{})
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_UNSPECIFIED.String()).
			Wrap(err)
	}

	return &cc.CollectionGetListResponse{
		Collections: collectionList,
	}, nil
}

func (a AuthContractImpl) CollectionGet(
	ctx *AuthTxCtx,
	req *cc.CollectionGetRequest,
) (res *cc.CollectionGetResponse, err error) {
	defer func() {
		if err != nil {
			ctx.Logger.Error(err.Error(), slog.Any("error", err))
		}
	}()

	// Validate the request
	err = ctx.Validate(req)
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Wrap(err)
	}

	// Get the collection
	col, err := ctx.SetCollection(req.GetCollectionId())
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_COLLECTION_INVALID.String()).
			Wrap(err)
	}

	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_USER_PERMISSION_DENIED.String()).
			Wrap(err)
	}

	return &cc.CollectionGetResponse{
		Collection: col,
	}, nil
}

func (a AuthContractImpl) CollectionGetHistory(
	ctx *AuthTxCtx,
	req *cc.CollectionGetHistoryRequest,
) (res *cc.CollectionGetHistoryResponse, err error) {
	// TODO implement CollectionGetHistory
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Wrap(err)
	}

	panic("implement me")
}

func (a AuthContractImpl) CollectionCreate(
	ctx *AuthTxCtx,
	req *cc.CollectionCreateRequest,
) (res *cc.CollectionCreateResponse, err error) {
	// TODO implement CollectionCreate
	defer func() { ctx.HandleFnError(&err, recover()) }()

	{ // Validate

		// Check req
		if err = ctx.Validate(req); err != nil {
			return nil, oops.Wrap(err)
		}

		// Extract AuthTransactionItems
		if err = ctx.ExtractAuthTransactionItems(req); err != nil {
			return nil, oops.Wrap(err)
		}

		// Check if the paths are valid for the type
		if err = rbac.ValidateCollection(req.Collection); err != nil {
			return nil, oops.In(ctx.GetFnName()).With("req", req).Wrap(err)
		}

	}

	{ // Authorize
		// all users can create collections
	}

	{ // Process

		// Create the collection
		res = &cc.CollectionCreateResponse{Collection: req.Collection}

		return res, state.Insert(ctx, req.Collection)
	}
}

func (a AuthContractImpl) CollectionUpdateRoles(
	ctx *AuthTxCtx,
	req *cc.CollectionUpdateRolesRequest,
) (res *cc.CollectionUpdateRolesResponse, err error) {
	// TODO implement CollectionUpdateRoles
	defer func() { ctx.HandleFnError(&err, recover()) }()

	{ // Validate
		// Validate the request
		if err = ctx.Validate(req); err != nil {
			return nil, oops.Wrap(err)
		}

		// Extract AuthTransactionItems
		if err = ctx.ExtractAuthTransactionItems(req); err != nil {
			return nil, oops.Wrap(err)
		}
		// Check if the paths are valid for the type

		// todo other validations
	}

	{ // Authorize
		if err = ctx.IsAuthorized(); err != nil {
			return nil, oops.Wrap(err)
		}
	}

	{ // Process
		col, err := ctx.GetCollection()
		if err != nil {
			return nil, oops.Wrap(err)
		}
		if col == nil {
			return nil, oops.Errorf("collection is nil")
		}

		// todo: Change other parts???
		col.Roles = req.Roles

		// Make sure the collection is valid
		if err = rbac.ValidateCollection(col); err != nil {
			return nil, oops.In(ctx.GetFnName()).With("req", req).Wrap(err)
		}

		res = &cc.CollectionUpdateRolesResponse{Collection: col}

		return res, state.Update(ctx, col)
	}
}

func (a AuthContractImpl) CollectionUpdatePermission(
	ctx *AuthTxCtx,
	req *cc.CollectionUpdatePermissionRequest,
) (res *cc.CollectionUpdatePermissionResponse, err error) {
	// TODO implement CollectionUpdatePermission
	defer func() { ctx.HandleFnError(&err, recover()) }()

	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Extract AuthTransactionItems
	if err = ctx.ExtractAuthTransactionItems(req); err != nil {
		return nil, oops.Wrap(err)
	}
	panic("implement me")
}
