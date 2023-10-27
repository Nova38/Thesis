package rbac

import (
	"fmt"
	"log/slog"

	"github.com/bufbuild/protovalidate-go"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nova38/thesis/lib/go/fabric/state"
	"github.com/samber/oops"

	// "github.com/rs/zerolog/log"
	_ "github.com/samber/lo"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
)

const (
	// DefaultPageSize default page size
	DefaultPageSize int32 = 10000
)

var validator *protovalidate.Validator

type AuthTransactionObjects struct {
	User       *rbac_pb.User
	Collection *rbac_pb.Collection
	ops        *rbac_pb.ACL_Operation
}

type ContextHelpers struct {
	Logger    *slog.Logger
	PageSize  int32
	Validator *protovalidate.Validator
}

type AuthTxCtx struct {
	contractapi.TransactionContext

	ContextHelpers
	AuthTransactionObjects

	// auth values
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

	if validator == nil {
		v, err := protovalidate.New()
		if err != nil {
			panic(err)
		}
		validator = v
	}

	return nil
}

func (ctx *AuthTxCtx) GetLogger() *slog.Logger {
	return ctx.Logger
}

func (ctx *AuthTxCtx) SetLogger(logger *slog.Logger) error {
	if logger == nil {
		return oops.Errorf("logger is nil")
	}
	ctx.Logger = logger
	return nil
}

// PagedTxCtxInterface functions
func (ctx *AuthTxCtx) GetPageSize() int32 {
	if ctx.PageSize == 0 {
		return DefaultPageSize
	}
	return ctx.PageSize
}

func (ctx *AuthTxCtx) SetPageSize(pageSize int32) {
	ctx.PageSize = pageSize
}

// ----------------------------------------------
// ValidateAbleTxCtxInterface functions
// ----------------------------------------------
func (ctx *AuthTxCtx) GetValidator() (*protovalidate.Validator, error) {
	if ctx.Validator == nil {
		v, err := protovalidate.New()
		if err != nil {
			return nil, oops.Errorf("failed to create validator: %w", err)
		}
		ctx.Validator = v
	}

	return ctx.Validator, nil
}

func (ctx *AuthTxCtx) Validate(msg proto.Message) error {
	v, err := ctx.GetValidator()
	if err != nil {
		return oops.Errorf("failed to get validator: %w", err)
	}
	if v == nil {
		return oops.Errorf("validator is nil")
	}
	return oops.
		In(ctx.GetFnName()).
		Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
		Wrap(v.Validate(msg))
}

func (ctx *AuthTxCtx) GetFnName() string {
	name, _ := ctx.GetStub().GetFunctionAndParameters()
	return name
}

func (ctx *AuthTxCtx) SetCollection(id *rbac_pb.Collection_Id) error {
	// TODO:
	// FIXME: Need to check in fns before calling is Authorized

	// See if the collection pointer has an ID and is not nil
	if id == nil || id.CollectionId == "" {
		return oops.
			In("SetCollection").
			Code(rbac_pb.Error_ERROR_COLLECTION_INVALID_ID.String()).
			Errorf("collection is nil or has no ID")
	}

	ctx.Collection = &rbac_pb.Collection{
		Id: &rbac_pb.Collection_Id{
			CollectionId: id.CollectionId,
		},
	}

	// Check if the collection exists
	err := state.GetState(ctx, ctx.Collection)
	if err != nil {
		return oops.
			In("SetCollection").
			With("collectionId", id.CollectionId).
			Code(rbac_pb.Error_ERROR_COLLECTION_UNREGISTERED.String()).
			Wrap(err)
	}

	return nil
}

func (ctx *AuthTxCtx) GetOperations() *rbac_pb.ACL_Operation {
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

func (ctx *AuthTxCtx) SetOperation(op *rbac_pb.ACL_Operation) error {
	// See if the operation pointer has an ID and is not nil
	if op == nil {
		return oops.Errorf("operation is nil or actions is nil")
	}

	if op.Domain == rbac_pb.ACL_DOMAIN_UNSPECIFIED {
		return oops.Errorf("operation domain is unspecified")
	}
	if op.Action == rbac_pb.ACL_ACTION_UNSPECIFIED {
		return oops.Errorf("operation action type is unspecified")
	}

	ctx.ops = &rbac_pb.ACL_Operation{}

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
			Code(rbac_pb.Error_ERROR_COLLECTION_INVALID_ROLE_ID.String()).
			Errorf("Role %v is not valid for collection %v", role, ctx.Collection.Id.CollectionId)
	}

	return AuthorizeOperation(ctx.ops, role.RoleId, ctx.Collection)
}

func (ctx *AuthTxCtx) IsAuthorized() error {
	if !ctx.authChecked {

		auth, err := ctx.authorize()
		if err != nil {
			slog.Error(err.Error(), slog.Any("error", err))
			return oops.
				In(ctx.GetFnName()).
				Code(rbac_pb.Error_ERROR_USER_PERMISSION_DENIED.String()).
				Wrap(err)
		}

		ctx.authorized = auth
	}
	return nil
}

func (ctx *AuthTxCtx) GetUser() (*rbac_pb.User, error) {
	if ctx.User != nil {
		return ctx.User, nil
	}

	id, err := ctx.GetUserId()
	if err != nil {
		return nil, oops.Wrap(err)
	}

	ctx.User = &rbac_pb.User{Id: id}

	err = state.GetState(ctx, ctx.User)

	if err != nil {
		return nil, oops.With("user_id", id).Wrap(err)
	}

	return ctx.User, nil
}

func (ctx *AuthTxCtx) GetUserId() (*rbac_pb.User_Id, error) {
	var err error

	Id := &rbac_pb.User_Id{
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

func (ctx *AuthTxCtx) GetCollection() (*rbac_pb.Collection, error) {
	if ctx.Collection != nil {
		return ctx.Collection, nil
	}

	return nil, oops.Errorf("collection not set")
}

func (ctx *AuthTxCtx) ExtractAuthTransactionItems(req interface{}) error {
	if col, ok := req.(CollectionHolder); ok {
		ctx.Collection = col.GetCollection()
		if ctx.Collection == nil {
			return oops.
				In(ctx.GetFnName()).
				Code(rbac_pb.Error_ERROR_COLLECTION_INVALID.String()).
				Errorf("collection is nil")
		}
	}

	if col_id, ok := req.(CollectionIdHolder); ok {
		if _, err := ctx.SetCollection(col_id.GetCollectionId()); err != nil {
			return oops.
				In(ctx.GetFnName()).
				Code(rbac_pb.Error_ERROR_COLLECTION_INVALID.String()).
				Wrap(err)
		}
	}

	if user, ok := req.(UserHolder); ok {
		ctx.User = user.GetUser()
		if ctx.User == nil {
			return oops.
				In(ctx.GetFnName()).
				Code(rbac_pb.Error_ERROR_USER_INVALID.String()).
				Errorf("user is nil")
		}
	}

	if user_id, ok := req.(UserIdHolder); ok {
		user_id := user_id.GetUserId()
		if user_id == nil {
			return oops.
				In(ctx.GetFnName()).
				Code(rbac_pb.Error_ERROR_USER_INVALID.String()).
				Errorf("user id is nil")
		}
		ctx.User = &rbac_pb.User{
			Id: user_id,
		}
		if err := state.GetState(ctx, ctx.User); err != nil {
			return oops.
				In(ctx.GetFnName()).
				Code(rbac_pb.Error_ERROR_USER_INVALID.String()).
				Wrap(err)
		}

	}

	return nil
}

func (ctx *AuthTxCtx) MakeLastModified() (*rbac_pb.StateActivity, error) {
	user, err := ctx.GetUser()
	if err != nil {
		return nil, oops.Errorf("failed to get user: %w", err)
	}

	timestamp, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return nil, oops.Errorf("Failed to get timestamp: %w", err)
	}

	return &rbac_pb.StateActivity{
		UserId:    user.GetId(),
		Note:      fmt.Sprintf("User %v modified the state", user.GetName()),
		TxId:      ctx.GetStub().GetTxID(),
		Timestamp: timestamp,
	}, nil
}
