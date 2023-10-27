package rbac

import (
	"log/slog"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nova38/thesis/lib/go/fabric/state"
	"github.com/samber/oops"

	"google.golang.org/protobuf/types/known/fieldmaskpb"

	pb "github.com/nova38/thesis/lib/go/gen/rbac"
)

type AuthTransactionObjects struct {
	User       *pb.User
	Collection *pb.Collection
	ops        *pb.ACL_Operation
}

type AuthTxCtx struct {
	state.LoggedTxCtx
	state.ValidateAbleTxCtx

	contractapi.TransactionContext

	AuthTransactionObjects

	authorized  bool
	authChecked bool
}

func (ctx *AuthTxCtx) HandelBefore() error {
	ctx.SetLogger(slog.Default().With(
		"fn", ctx.GetFnName(),
		slog.Group(
			"tx info",
			"tx_id", ctx.GetStub().GetTxID(),
			"channel_id", ctx.GetStub().GetChannelID(),
		),
	))

	return nil
}

func (ctx *AuthTxCtx) GetLogger() *slog.Logger {
	return ctx.Logger
}

func (ctx *AuthTxCtx) SetLogger(logger *slog.Logger) error {
	ctx.Logger = logger
	return nil
}

func (ctx *AuthTxCtx) GetFnName() string {
	name, _ := ctx.GetStub().GetFunctionAndParameters()
	return name
}

func (ctx *AuthTxCtx) SetCollection(id *pb.Collection_Id) (*pb.Collection, error) {
	// See if the collection pointer has an ID and is not nil
	if id == nil || id.CollectionId == "" {
		return nil, oops.
			In("SetCollection").
			Code(pb.Error_ERROR_COLLECTION_INVALID_ID.String()).
			Errorf("collection is nil or has no ID")
	}

	ctx.Collection = &pb.Collection{
		Id: &pb.Collection_Id{
			CollectionId: id.CollectionId,
		},
	}

	// Check if the collection exists
	err := state.GetState(ctx, ctx.Collection)
	if err != nil {
		return nil, oops.
			In("SetCollection").
			With("collectionId", id.CollectionId).
			Code(pb.Error_ERROR_COLLECTION_UNREGISTERED.String()).
			Wrap(err)
	}

	return ctx.Collection, nil
}

func (ctx *AuthTxCtx) GetOperations() *pb.ACL_Operation {
	return ctx.ops
}

func (ctx *AuthTxCtx) SetOperationsPaths(paths *fieldmaskpb.FieldMask) error {
	if paths == nil {
		return oops.Errorf("paths is nil")
	}

	if ctx.ops == nil {
		return oops.Errorf("operations is nil")
	}
	ctx.ops.Paths = paths

	return nil
}

func (ctx *AuthTxCtx) SetOperation(op *pb.ACL_Operation) error {
	// See if the operation pointer has an ID and is not nil
	if op == nil {
		return oops.Errorf("operation is nil or actions is nil")
	}

	if op.Domain == pb.ACL_DOMAIN_UNSPECIFIED {
		return oops.Errorf("operation domain is unspecified")
	}
	if op.Action == pb.ACL_ACTION_UNSPECIFIED {
		return oops.Errorf("operation action type is unspecified")
	}

	ctx.ops = &pb.ACL_Operation{}

	ctx.ops.Action = op.Action
	ctx.ops.Domain = op.Domain

	if op.Paths != nil {
		ctx.ops.Paths = op.Paths
	}
	return nil
}

func (ctx *AuthTxCtx) authorize() (bool, error) {
	// if the user is already authorized, return the value
	if ctx.authChecked {
		return ctx.authorized, nil
	}
	ctx.authChecked = true

	// Check if all the objects are set
	if ctx.Collection == nil {
		return false, oops.Errorf("authorization objects not set")
	}

	user, err := ctx.GetUser()
	if err != nil {
		return false, oops.Wrap(err)
	}

	role, ok := user.Roles[ctx.Collection.Id.CollectionId]
	if !ok {
		return false, oops.
			In("AuthorizeOperation").
			Code(pb.Error_ERROR_COLLECTION_INVALID_ROLE_ID.String()).
			Errorf("Role %v is not valid for collection %v", role, ctx.Collection.Id.CollectionId)
	}

	return AuthorizeOperation(ctx.ops, role.RoleId, ctx.Collection)
}

func (ctx *AuthTxCtx) IsAuthorized() error {
	if !ctx.authChecked {

		auth, err := ctx.authorize()
		if err != nil {
			slog.Error(err.Error(), slog.Any("error", err))
			return oops.Wrap(err)
		}

		ctx.authorized = auth
	}
	return nil
}

func (ctx *AuthTxCtx) GetUser() (*pb.User, error) {
	if ctx.User != nil {
		return ctx.User, nil
	}

	id, err := ctx.GetUserId()
	if err != nil {
		return nil, oops.Wrap(err)
	}

	ctx.User = &pb.User{Id: id}

	err = state.GetState(ctx, ctx.User)

	if err != nil {
		return nil, oops.With("user_id", id).Wrap(err)
	}

	return ctx.User, nil
}

func (ctx *AuthTxCtx) GetUserId() (*pb.User_Id, error) {
	var err error

	Id := &pb.User_Id{
		MspId: "",
		Id:    "",
	}

	// Extract The info from the Client ID
	id := ctx.GetClientIdentity()

	Id.Id, err = id.GetID()
	if err != nil {
		return nil, oops.Errorf("failed to get user certificate from CID: %s", err)
	}

	Id.MspId, err = id.GetMSPID()
	if err != nil {
		return nil, oops.Errorf("failed to get user ID from CID: %s", err)
	}

	return Id, nil
}

func (ctx *AuthTxCtx) GetCollection() (*pb.Collection, error) {
	if ctx.Collection != nil {
		return ctx.Collection, nil
	}

	return nil, oops.Errorf("collection not set")
}

func (ctx *AuthTxCtx) ExtractAuthTransactionItems(req interface{}) error {
	if col, ok := req.(CollectionWrapperInterface); ok {
		ctx.Collection = col.GetCollection()
		if ctx.Collection == nil {
			return oops.
				In(ctx.GetFnName()).
				Code(pb.Error_ERROR_COLLECTION_INVALID.String()).
				Errorf("collection is nil")
		}
	}

	if col_id, ok := req.(CollectionIdWrapperInterface); ok {
		if _, err := ctx.SetCollection(col_id.GetCollectionId()); err != nil {
			return oops.
				In(ctx.GetFnName()).
				Code(pb.Error_ERROR_COLLECTION_INVALID.String()).
				Wrap(err)
		}
	}

	if user, ok := req.(UserWrapperInterface); ok {
		ctx.User = user.GetUser()
		if ctx.User == nil {
			return oops.
				In(ctx.GetFnName()).
				Code(pb.Error_ERROR_USER_INVALID.String()).
				Errorf("user is nil")
		}
	}

	if user_id, ok := req.(UserIdWrapperInterface); ok {
		user_id := user_id.GetUserId()
		if user_id == nil {
			return oops.
				In(ctx.GetFnName()).
				Code(pb.Error_ERROR_USER_INVALID.String()).
				Errorf("user id is nil")
		}
		ctx.User = &pb.User{
			Id: user_id,
		}
		if err := state.GetState(ctx, ctx.User); err != nil {
			return oops.
				In(ctx.GetFnName()).
				Code(pb.Error_ERROR_USER_INVALID.String()).
				Wrap(err)
		}

	}

	return nil
}
