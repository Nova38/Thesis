package contracts

import (
	"log/slog"

	"github.com/nova38/thesis/lib/go/fabric/auth/state"

	"github.com/samber/lo"
	"github.com/samber/oops"

	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	cc "github.com/nova38/thesis/lib/go/gen/chaincode/auth/rbac/schema/v1"
)

func (a AuthContractImpl) CollectionGetList(
	ctx *AuthTxCtx,
) (res *cc.CollectionGetListResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	// Get the collections
	collectionList, err := state.GetFullList(ctx, &authpb.Collection{})
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(authpb.TxError_UNSPECIFIED.String()).
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
			Code(authpb.TxError_REQUEST_INVALID.String()).
			Wrap(err)
	}

	// Get the collection
	col, err := ctx.SetCollection(req.GetCollectionId())
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(authpb.TxError_COLLECTION_ALREADY_REGISTERED.String()).
			Wrap(err)
	}

	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(authpb.TxError_USER_PERMISSION_DENIED.String()).
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
		return nil, oops.Wrap(err)
	}

	panic("implement me")
}

func (a AuthContractImpl) CollectionCreate(
	ctx *AuthTxCtx,
	req *cc.CollectionCreateRequest,
) (res *cc.CollectionCreateResponse, err error) {
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
		if err = auth.ValidateCollection(req.Collection); err != nil {
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

		// 1. a role can't be in both list
		// 2. the roles in both list must be valid (above 1)
		if len(req.RolesToAdd) == 0 && len(req.RolesToRemove) == 0 {
			return nil, oops.Errorf("add and remove roles are empty")
		}

		if len(req.RolesToAdd) != len(lo.Uniq(lo.Values(req.RolesToAdd))) {
			return nil, oops.
				In(ctx.GetFnName()).
				With("roles", req.RolesToAdd).
				Errorf("add roles contains duplicates role names")
		}

		if len(req.RolesToRemove) != len(lo.Uniq(lo.Values(req.RolesToRemove))) {
			return nil, oops.
				In(ctx.GetFnName()).
				With("roles", req.RolesToRemove).
				Errorf("remove roles contains duplicates role names")
		}

		if len(lo.Intersect(lo.Values(req.RolesToAdd), lo.Values(req.RolesToRemove))) > 0 {
			return nil, oops.
				In(ctx.GetFnName()).
				With("remove", req.RolesToRemove, "add", req.RolesToAdd).
				Errorf("add and remove roles contains same role names")
		}

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
