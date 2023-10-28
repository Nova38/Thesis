package contract

import (
	"log/slog"
	_ "strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/samber/oops"

	_ "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/nova38/thesis/lib/go/fabric/rbac"
	"github.com/nova38/thesis/lib/go/fabric/state"

	cc "github.com/nova38/thesis/lib/go/gen/chaincode/rbac/schema/v1"
	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
)

// Check if AuthContractImpl implements AuthServiceInterface
var (
	_ rbac.AuthTxCtxInterface       = (*AuthTxCtx)(nil)
	_ cc.AuthServiceInterface       = (*AuthContractImpl)(nil)
	_ contractapi.ContractInterface = (*AuthContractImpl)(nil)
)

type AuthContractImpl struct {
	contractapi.Contract
	cc.AuthServiceBase
}

func (a AuthContractImpl) GetBeforeTransaction() interface{} {
	return a.BeforeTransaction
}

func (a AuthContractImpl) BeforeTransaction(ctx AuthTxCtx) (err error) {
	defer func() {
		if err != nil && ctx.Logger != nil {
			ctx.Logger.Error(err.Error(), slog.Any("error", err))
		}
	}()

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

// UserGetCurrent: Returns the current user.
//
// Returns the current user.
// # Requires:
//   - User submitting the transaction is a registered user.
//
// # Operation:
//   - Domain: DOMAIN_USER
//   - Action: ACTION_VIEW
func (a AuthContractImpl) UserGetCurrent(
	ctx AuthTxCtx,
) (res *cc.UserGetCurrentResponse, err error) {
	defer func() {
		if err != nil && ctx.Logger != nil {
			ctx.Logger.Error(err.Error(), slog.Any("error", err))
		}
	}()

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
	ctx AuthTxCtx,
) (res *cc.UserGetCurrentIdResponse, err error) {
	defer func() {
		if err != nil && ctx.Logger != nil {
			ctx.Logger.Error(err.Error(), slog.Any("error", err))
		}
	}()
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
	ctx AuthTxCtx,
) (res *cc.UserGetListResponse, err error) {
	defer func() {
		if err != nil && ctx.Logger != nil {
			ctx.Logger.Error(err.Error(), slog.Any("error", err))
		}
	}()
	userList, err := state.GetFullStateList(&ctx, &rbac_pb.User{})
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
	ctx AuthTxCtx,
	req *cc.UserGetRequest,
) (res *cc.UserGetResponse, err error) {
	defer func() {
		if err != nil && ctx.Logger != nil {
			ctx.Logger.Error(err.Error(), slog.Any("error", err))
		}
	}()

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

	if err = state.GetState(&ctx, user); err != nil {
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
	ctx AuthTxCtx,
	req *cc.UserGetHistoryRequest,
) (res *cc.UserGetHistoryResponse, err error) {
	defer func() {
		if err != nil && ctx.Logger != nil {
			ctx.Logger.Error(err.Error(), slog.Any("error", err))
		}
	}()
	// TODO implement UserGetHistory
	panic("implement me")
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
	ctx AuthTxCtx,
	req *cc.UserRegisterRequest,
) (res *cc.UserRegisterResponse, err error) {
	defer func() {
		if err != nil && ctx.Logger != nil {
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

	id, err := ctx.GetUserId()
	if err != nil {
		return nil, err
	}

	// Create the user
	user := &rbac_pb.User{
		Id: &rbac_pb.User_Id{
			MspId: id.GetMspId(),
			Id:    id.GetId(),
		},
		Name: req.GetName(),
	}

	// Check if the user already exists
	if err = state.InsertState(&ctx, user); err != nil {
		return nil, oops.
			In("UserRegister").
			Code(rbac_pb.Error_ERROR_USER_ALREADY_REGISTERED.String()).
			Wrap(err)
	}

	return &cc.UserRegisterResponse{User: user}, nil
}

func (a AuthContractImpl) UserUpdate(
	ctx AuthTxCtx,
	req *cc.UserUpdateRequest,
) (res *cc.UserUpdateResponse, err error) {
	// TODO implement UserUpdate
	defer func() {
		if err != nil && ctx.Logger != nil {
			ctx.Logger.Error(err.Error(), slog.Any("error", err))
		}
	}()
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Wrap(err)
	}
	// TODO: Extract transaction items from the request

	// Authorize the request
	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_USER_PERMISSION_DENIED.String()).
			Wrap(err)
	}

	panic("implement me")
}

func (a AuthContractImpl) UserUpdateMembership(
	ctx AuthTxCtx,
	req *cc.UserUpdateMembershipRequest,
) (res *cc.UserUpdateMembershipResponse, err error) {
	defer func() {
		if err != nil && ctx.Logger != nil {
			ctx.Logger.Error(err.Error(), slog.Any("error", err))
		}
	}()
	// Validate the request
	{

		if err = ctx.Validate(req); err != nil {
			return nil, oops.
				In(ctx.GetFnName()).
				Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
				Wrap(err)
		}
		col, err := ctx.SetCollection(req.CollectionId)
		if err != nil {
			return nil, oops.
				In(ctx.GetFnName()).
				Code(rbac_pb.Error_ERROR_COLLECTION_INVALID.String()).
				Wrap(err)
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

	// Authorize the request

	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_USER_PERMISSION_DENIED.String()).
			Wrap(err)
	}

	// -------------------------
	// Process the request
	// -------------------------

	// Get the user to modify
	userToModify := &rbac_pb.User{
		Id: req.GetId(),
	}
	if err = state.GetState(&ctx, userToModify); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_USER_INVALID.String()).
			Wrap(err)
	}

	// Update the user
	userToModify.Roles[req.GetCollectionId().CollectionId] = &rbac_pb.UserRole{
		CollectionId: &rbac_pb.Collection_Id{
			CollectionId: req.CollectionId.CollectionId,
		},
		RoleId:    req.GetRole(),
		GrantedBy: ctx.User.GetId(),
	}

	return &cc.UserUpdateMembershipResponse{User: userToModify},
		state.UpdateState(&ctx, userToModify)
}

func (a AuthContractImpl) CollectionGetList(
	ctx AuthTxCtx,
) (res *cc.CollectionGetListResponse, err error) {
	defer func() {
		if err != nil && ctx.Logger != nil {
			ctx.Logger.Error(err.Error(), slog.Any("error", err))
		}
	}()

	// Get the collections
	collectionList, err := state.GetFullStateList(&ctx, &rbac_pb.Collection{})
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
	ctx AuthTxCtx,
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
	ctx AuthTxCtx,
	req *cc.CollectionGetHistoryRequest,
) (res *cc.CollectionGetHistoryResponse, err error) {
	// TODO implement CollectionGetHistory
	defer func() {
		if err != nil && ctx.Logger != nil {
			ctx.Logger.Error(err.Error(), slog.Any("error", err))
		}
	}()
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
	defer func() {
		if err != nil && ctx.Logger != nil {
			ctx.Logger.Error(err.Error(), slog.Any("error", err))
		}
	}()
	if err = ctx.Validate(req); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Wrap(err)
	}

	panic("implement me")
}

func (a AuthContractImpl) CollectionUpdateRoles(
	ctx AuthTxCtx,
	req *cc.CollectionUpdateRolesRequest,
) (res *cc.CollectionUpdateRolesResponse, err error) {
	// TODO implement CollectionUpdateRoles
	defer func() {
		if err != nil && ctx.Logger != nil {
			ctx.Logger.Error(err.Error(), slog.Any("error", err))
		}
	}()
	if err = ctx.Validate(req); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Wrap(err)
	}

	panic("implement me")
}

func (a AuthContractImpl) CollectionUpdatePermission(
	ctx AuthTxCtx,
	req *cc.CollectionUpdatePermissionRequest,
) (res *cc.CollectionUpdatePermissionResponse, err error) {
	// TODO implement CollectionUpdatePermission
	defer func() {
		if err != nil && ctx.Logger != nil {
			ctx.Logger.Error(err.Error(), slog.Any("error", err))
		}
	}()
	if err = ctx.Validate(req); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Wrap(err)
	}

	panic("implement me")
}
