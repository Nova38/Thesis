package contract

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/nova38/saacs/pkg/saacs-cc/auth/collections"
	"github.com/nova38/saacs/pkg/saacs-cc/common"
	"github.com/nova38/saacs/pkg/saacs-cc/serializer"
	authpb "github.com/nova38/saacs/pkg/saacs-protos/saacs/auth/v0"
	cc "github.com/nova38/saacs/pkg/saacs-protos/saacs/chaincode/v0"
	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"
	"github.com/samber/oops"
)

// ══════════════════════════════════ Helper ═════════════════════════════════════
// ────────────────────────────────── Query ──────────────────────────────────────

func (o ContractImpl) GetCurrentUser(
	ctx common.TxCtxInterface,
	req *cc.GetCurrentUserRequest,
) (res *cc.GetCurrentUserResponse, err error) {

	res = &cc.GetCurrentUserResponse{}

	res.User, err = ctx.GetUserId()

	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	return res, err
}

func (o ContractImpl) GetCollectionsList(
	ctx common.TxCtxInterface,
	req *cc.GetCollectionsListRequest,
) (res *cc.GetCollectionsListResponse, err error) {

	col := &authpb.Collection{}

	results, err := ctx.GetStub().
		GetStateByPartialCompositeKey(
			col.ItemType(),
			[]string{},
		)
	if err != nil {
		ctx.GetLogger().Info("error")

		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	ctx.GetLogger().Info("No error")
	defer func(results shim.StateQueryIteratorInterface) {
		err := results.Close()
		if err != nil {
			ctx.GetLogger().Error("GetCollectionsList", "Error", err)
		}
	}(results)

	res = &cc.GetCollectionsListResponse{
		Collections: []*authpb.Collection{},
	}

	for results.HasNext() {

		tmp := &authpb.Collection{}

		queryResponse, err := results.Next()
		if err != nil || queryResponse == nil {
			return nil, oops.Wrapf(err, "Error getting next item")
		}

		if err = serializer.Unmarshal(queryResponse.GetValue(), tmp); err != nil {
			return nil, oops.Wrap(err)
		}

		res.Collections = append(res.GetCollections(), tmp)
	}

	// list, _, err := state.GetPartialKeyList(ctx, col, 0, "")
	// if err != nil {
	// 	ctx.LogError(err)
	// 	return nil, oops.Wrap(err)
	// }

	return res, nil
}

// ──────────────────────────────────── Invoke ─────────────────────────────────────

func (o ContractImpl) AuthorizeOperation(
	ctx common.TxCtxInterface,
	req *cc.AuthorizeOperationRequest,
) (res *cc.AuthorizeOperationResponse, err error) {

	authorized, err := ctx.Authorize([]*pb.Operation{req.GetOperation()})
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	return &cc.AuthorizeOperationResponse{
		Authorized: authorized,
	}, err
}

func (o ContractImpl) Bo(
	ctx common.TxCtxInterface,
) (err error) {
	return ctx.GetStub().PutState("bootstrap", []byte("true"))

}
func (o ContractImpl) Bootstrap(
	ctx common.TxCtxInterface,
	req *cc.BootstrapRequest,
) (res *cc.BootstrapResponse, err error) {

	// ctx.GetStub().PutState("bootstrap", []byte("true"))

	switch req.GetCollection().GetAuthType() {
	case authpb.AuthType_AUTH_TYPE_GLOBAL_ROLE:
		ctx.GetLogger().Info("Global Role Auth BootStrap")

		builder := &collections.GlobalRBAC{}
		if _, err = builder.CreateCollection(ctx, req.GetCollection()); err != nil {
			ctx.LogError(err)
			return nil, oops.In("Contract").Wrap(err)
		}
		return &cc.BootstrapResponse{Collection: req.GetCollection(), Success: true}, nil
	case authpb.AuthType_AUTH_TYPE_ROLE:
		builder := &collections.RBAC{}
		if _, err = builder.CreateCollection(ctx, req.GetCollection()); err != nil {
			ctx.LogError(err)
			return nil, oops.In("Contract").Wrap(err)
		}
		return &cc.BootstrapResponse{Collection: req.GetCollection(), Success: true}, nil
	case authpb.AuthType_AUTH_TYPE_IDENTITY:
		builder := &collections.Identity{}
		if _, err = builder.CreateCollection(ctx, req.GetCollection()); err != nil {
			ctx.LogError(err)
			return nil, oops.In("Contract").Wrap(err)
		}
		return &cc.BootstrapResponse{Collection: req.GetCollection(), Success: true}, nil
	case authpb.AuthType_AUTH_TYPE_NONE:
		ctx.GetLogger().Info("No Auth")
		builder := &collections.NoAuth{}
		if _, err = builder.CreateCollection(ctx, req.GetCollection()); err != nil {
			ctx.LogError(err)
			return nil, oops.In("Contract").Wrap(err)
		}
		return &cc.BootstrapResponse{Collection: req.GetCollection(), Success: true}, nil
	default:
		return nil, oops.In("Contract").Errorf("Invalid AuthType")
	}

}
