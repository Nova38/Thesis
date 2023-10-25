package rbac

import (
	_ "strings"

	"google.golang.org/protobuf/types/known/emptypb"

	// "github.com/rs/zerolog/log"
	_ "github.com/samber/lo"
	"github.com/samber/oops"
	_ "github.com/samber/oops"

	_ "google.golang.org/protobuf/types/known/timestamppb"

	// "github.com/hyperledger-labs/cckit/identity"

	"github.com/nova38/thesis/lib/go/fabric/rbac"
	"github.com/nova38/thesis/lib/go/fabric/state"

	cc "github.com/nova38/thesis/lib/go/gen/chaincode/rbac/schema/v1"
	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
)

// Check if AuthContractImpl implements AuthServiceInterface
// var _ cc.AuthServiceInterface = (*AuthContractImpl)(nil)

type AuthContractImpl struct {
	cc.AuthServiceBase
}

func HandelBefore(ctx rbac.AuthTxCtx) error {
	ops, err := cc.AuthServiceGetTxOperation(ctx.GetFnName())
	if err != nil {
		return oops.
			In("HandelBefore").
			With("fn", ctx.GetFnName()).
			Wrap(err)
	}

	err = ctx.SetOperation(ops)
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
func (a AuthContractImpl) UserGetCurrent(ctx rbac.AuthTxCtx, req *emptypb.Empty) (res *cc.UserGetCurrentResponse, err error) {
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
func (a AuthContractImpl) UserGetCurrentId(ctx rbac.AuthTxCtx, req *emptypb.Empty) (res *cc.UserGetCurrentIdResponse, err error) {
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
func (a AuthContractImpl) UserGetList(ctx rbac.AuthTxCtx, req *emptypb.Empty) (res *cc.UserGetListResponse, err error) {
	// TODO implement me
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
func (a AuthContractImpl) UserGet(ctx rbac.AuthTxCtx, req *cc.UserGetRequest) (res *cc.UserGetResponse, err error) {
	// Validate the request
	v, err := ctx.GetValidator()
	v.Validate(req)

	// Get the user
	user := &rbac_pb.User{Id: req.GetId()}

	err = state.GetState(&ctx, user)
	if err != nil {
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
func (a AuthContractImpl) UserGetHistory(ctx rbac.AuthTxCtx, req *cc.UserGetHistoryRequest) (res *cc.UserGetHistoryResponse, err error) {
	// TODO implement me
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
func (a AuthContractImpl) UserRegister(ctx rbac.AuthTxCtx, req *cc.UserRegisterRequest) (res *cc.UserRegisterResponse, err error) {
	// TODO implement me

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
	if err := state.InsertState(&ctx, user); err != nil {
		return nil, oops.
			In("UserRegister").
			Code(rbac_pb.Error_ERROR_USER_ALREADY_REGISTERED.String()).
			Wrap(err)
	}

	return &cc.UserRegisterResponse{User: user}, nil
}

func (a AuthContractImpl) UserUpdate(ctx rbac.AuthTxCtx, req *cc.UserUpdateRequest) (res *cc.UserUpdateResponse, err error) {
	// TODO implement me

	// Validate the request
	if err := ctx.Validate(req); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Wrap(err)
	}
	// TODO: Extract transaction items from the request

	// Authorize the request
	if err := ctx.IsAuthorized(); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_USER_PERMISSION_DENIED.String()).
			Wrap(err)
	}

	panic("implement me")
}

func (a AuthContractImpl) UserUpdateMembership(ctx rbac.AuthTxCtx, req *cc.UserUpdateMembershipRequest) (res *cc.UserUpdateMembershipResponse, err error) {
	// TODO implement me
	// Validate the request
	{
		err := ctx.Validate(req)
		if err != nil {
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

	if err := ctx.IsAuthorized(); err != nil {
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
	if err := state.GetState(&ctx, userToModify); err != nil {
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

func (a AuthContractImpl) CollectionGetList(ctx rbac.AuthTxCtx, req *emptypb.Empty) (res *cc.CollectionGetListResponse, err error) {
	// TODO implement me
	// panic("implement me")

	// Validate the request
	err = ctx.Validate(req)
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Wrap(err)
	}

	// Get the collections
	collectionList, err := state.GetFullStateList(&ctx, &rbac_pb.Collection{})

	return &cc.CollectionGetListResponse{
		Collections: collectionList,
	}, nil
}

func (a AuthContractImpl) CollectionGet(ctx rbac.AuthTxCtx, req *cc.CollectionGetRequest) (res *cc.CollectionGetResponse, err error) {
	// TODO implement me
	// panic("implement me")

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

	// TODO: check if user is authorized to view the collection

	if err := ctx.IsAuthorized(); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_USER_PERMISSION_DENIED.String()).
			Wrap(err)
	}

	return &cc.CollectionGetResponse{
		Collection: col,
	}, nil
}

func (a AuthContractImpl) CollectionGetHistory(ctx rbac.AuthTxCtx, req *cc.CollectionGetHistoryRequest) (res *cc.CollectionGetHistoryResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (a AuthContractImpl) CollectionCreate(ctx rbac.AuthTxCtx, req *cc.CollectionCreateRequest) (res *cc.CollectionCreateResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (a AuthContractImpl) CollectionUpdateRoles(ctx rbac.AuthTxCtx, req *cc.CollectionUpdateRolesRequest) (res *cc.CollectionUpdateRolesResponse, err error) {
	// TODO implement me
	panic("implement me")
}

func (a AuthContractImpl) CollectionUpdatePermission(ctx rbac.AuthTxCtx, req *cc.CollectionUpdatePermissionRequest) (res *cc.CollectionUpdatePermissionResponse, err error) {
	// TODO implement me
	panic("implement me")
}
