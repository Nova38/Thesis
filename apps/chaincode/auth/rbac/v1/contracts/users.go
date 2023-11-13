package contracts

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	"github.com/samber/oops"

	"google.golang.org/protobuf/types/known/anypb"

	cc "github.com/nova38/thesis/lib/go/gen/chaincode/rbac/schema/v1"
)

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

// UserGetCurrentId Returns the current user id.
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
			Code(rbac_pb.TxError_USER_INVALID.String()).
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
			Code(rbac_pb.TxError_UNSPECIFIED.String()).
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
	if err = ctx.Validate(req); err != nil {
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
		return nil, oops.Wrap(err)
	}

	// Get the history of the users
	userHistory, err := state.GetHistory(ctx, &rbac_pb.User{Id: req.GetId()})
	if err != nil {
		return nil, oops.Wrap(err)
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
				Code(rbac_pb.TxError_COLLECTION_INVALID_ROLE_ID.String()).
				Errorf("Role %v is not valid for collection %v", req.Role, req.CollectionId)
		}

		// check if the req user is valid
		id, err := ctx.GetUserId()
		if err != nil || id == nil {
			return nil, oops.
				In(ctx.GetFnName()).
				Code(rbac_pb.TxError_UNSPECIFIED.String()).
				Errorf("User id is nil")
		}

	}

	{ // Authorize the request
		if err = ctx.IsAuthorized(); err != nil {
			return nil, oops.
				In(ctx.GetFnName()).
				Code(rbac_pb.TxError_USER_PERMISSION_DENIED.String()).
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
				Code(rbac_pb.TxError_USER_INVALID.String()).
				Wrap(err)
		}

		// Edit the user
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
