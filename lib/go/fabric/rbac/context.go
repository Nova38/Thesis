package rbac

import (
	"fmt"
	"log/slog"

	"github.com/bufbuild/protovalidate-go"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nova38/thesis/lib/go/fabric/state"
	"github.com/samber/oops"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
)

const (
	// DefaultPageSize default page size
	DefaultPageSize int32 = 10000
)

var (
	// Validate that AuthTxCtx implements the required interfaces
	_ AuthTxCtxInterface               = (*TxCtx)(nil)
	_ state.LoggedTxCtxInterface       = (*TxCtx)(nil)
	_ state.ValidateAbleTxCtxInterface = (*TxCtx)(nil)
	_ state.PagedTxCtxInterface        = (*TxCtx)(nil)

	// Global validator
	validator *protovalidate.Validator
)

type (
	AuthTransactionObjects struct {
		User       *rbac_pb.User
		Collection *rbac_pb.Collection
		ops        *rbac_pb.ACL_Operation
	}

	ContextHelpers struct {
		Logger    *slog.Logger
		PageSize  int32
		Validator *protovalidate.Validator
	}

	TxCtx struct {
		contractapi.TransactionContext

		ContextHelpers
		AuthTransactionObjects

		// auth values
		authorized  bool
		authChecked bool
	}
)

func (ctx *TxCtx) HandelBefore() (err error) {
	return ctx.BaseHandelBefore()
}

func (ctx *TxCtx) BaseHandelBefore() (err error) {
	err = ctx.SetLogger(slog.Default().With(
		"fn", ctx.GetFnName(),
		slog.Group(
			"tx info",
			"tx_id", ctx.GetStub().GetTxID(),
			"channel_id", ctx.GetStub().GetChannelID(),
		),
	))
	if err != nil {
		return err
	}

	if validator == nil {
		v, err := protovalidate.New()
		if err != nil {
			panic(err)
		}
		validator = v
	}

	return nil
}

func (ctx *TxCtx) HandleFnError(err *error, r any) {
	if ctx.Logger == nil {
		ctx.Logger = slog.Default()
	}

	if r != nil {
		ctx.Logger.Error("Panic", slog.Any("panic", r))
		e := oops.Errorf("Panic: %v", r)
		err = &e
	}

	if *err != nil {
		slog.Error((*err).Error())
	}
}

// ----------------------------------------------
// LoggedTxCtxInterface
// ----------------------------------------------
func (ctx *TxCtx) GetLogger() *slog.Logger {
	return ctx.Logger
}

func (ctx *TxCtx) SetLogger(logger *slog.Logger) error {
	if logger == nil {
		return oops.Errorf("logger is nil")
	}
	ctx.Logger = logger
	return nil
}

// PagedTxCtxInterface functions
func (ctx *TxCtx) GetPageSize() int32 {
	if ctx.PageSize == 0 {
		return DefaultPageSize
	}
	return ctx.PageSize
}

func (ctx *TxCtx) SetPageSize(pageSize int32) {
	ctx.PageSize = pageSize
}

// ----------------------------------------------
// ValidateAbleTxCtxInterface functions
// ----------------------------------------------
func (ctx *TxCtx) GetValidator() (*protovalidate.Validator, error) {
	if ctx.Validator == nil {
		v, err := protovalidate.New()
		if err != nil {
			return nil, oops.Errorf("failed to create validator: %w", err)
		}
		ctx.Validator = v
	}

	return ctx.Validator, nil
}

func (ctx *TxCtx) Validate(msg proto.Message) error {
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

// ----------------------------------------------
// AuthTxCtxInterface functions
// ----------------------------------------------

func (ctx *TxCtx) GetFnName() (name string) {
	name, _ = ctx.GetStub().GetFunctionAndParameters()
	return name
}

func (ctx *TxCtx) MakeLastModified() (mod *rbac_pb.StateActivity, err error) {
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

// =============================================
//  User Functions
// =============================================

func (ctx *TxCtx) GetUserId() (userId *rbac_pb.User_Id, err error) {
	userId = &rbac_pb.User_Id{}

	// Extract The info from the Client ID
	id := ctx.GetClientIdentity()

	userId.Id, err = id.GetID()
	if err != nil {
		return nil, oops.Errorf("failed to get user certificate from CID: %s", err)
	}

	userId.MspId, err = id.GetMSPID()
	if err != nil {
		return nil, oops.Errorf("failed to get user ID from CID: %s", err)
	}

	return userId, nil
}

func (ctx *TxCtx) GetUser() (user *rbac_pb.User, err error) {
	if ctx.User != nil {
		return ctx.User, nil
	}

	id, err := ctx.GetUserId()
	if err != nil {
		return nil, oops.Wrap(err)
	}

	ctx.User = &rbac_pb.User{Id: id}

	err = state.Get(ctx, ctx.User)

	if err != nil {
		return nil, oops.With("user_id", id).Wrap(err)
	}

	return ctx.User, nil
}

// =============================================
//  Collection Functions
// =============================================

func (ctx *TxCtx) GetCollection() (col *rbac_pb.Collection, err error) {
	if ctx.Collection != nil {
		return ctx.Collection, nil
	}

	return nil, oops.Errorf("collection not set")
}

func (ctx *TxCtx) SetCollection(id *rbac_pb.Collection_Id) (col *rbac_pb.Collection, err error) {
	// TODO:
	// FIXME: Need to check in fns before calling is Authorized

	// See if the collection pointer has an ID and is not nil
	if id == nil || id.CollectionId == "" {
		return nil, oops.
			In("SetCollection").
			Code(rbac_pb.Error_ERROR_COLLECTION_INVALID_ID.String()).
			Errorf("collection is nil or has no ID")
	}

	ctx.Collection = &rbac_pb.Collection{
		Id: &rbac_pb.Collection_Id{
			CollectionId: id.CollectionId,
		},
	}

	if err = state.Get(ctx, ctx.Collection); err != nil {
		return nil, oops.
			In("SetCollection").
			With("collectionId", id.CollectionId).
			Code(rbac_pb.Error_ERROR_COLLECTION_UNREGISTERED.String()).
			Wrap(err)
	}

	return ctx.Collection, nil
}

// =============================================
//  Role Functions
// =============================================

func (ctx *TxCtx) GetRole() (role string, err error) {
	// Check the requirements
	if ctx.User != nil || ctx.Collection != nil {
		return role, oops.
			With(
				"user", ctx.User,
				"collection", ctx.Collection,
			).
			Errorf("user or collection are nil")
	}

	// Check to see if the user is a member of the collection
	if cRole, ok := ctx.User.Roles[ctx.Collection.Id.CollectionId]; ok {
		return cRole.RoleId, nil
	}

	ctx.GetLogger().Info("User is not a member of the collection, assigning to public")

	return "0", nil
}

// =============================================
//  Operations Functions
// =============================================

func (ctx *TxCtx) SetOperation(op *rbac_pb.ACL_Operation) (err error) {
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

func (ctx *TxCtx) GetOperations() (ops *rbac_pb.ACL_Operation, err error) {
	if ctx.ops != nil {
		return ctx.ops, nil
	}
	return nil, oops.Errorf("operations not set")
}

func (ctx *TxCtx) SetOperationsPaths(paths *fieldmaskpb.FieldMask) (err error) {
	if paths == nil {
		return oops.Errorf("paths is nil")
	}

	if ctx.ops == nil {
		return oops.Errorf("operations is nil")
	}
	ctx.ops.Paths = paths

	return nil
}

// =============================================
//
//	ACL Functions
//
// =============================================
func (ctx *TxCtx) Authorize() (auth bool, err error) {
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

func (ctx *TxCtx) IsAuthorized() (err error) {
	if !ctx.authChecked {

		auth, err := ctx.Authorize()
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
