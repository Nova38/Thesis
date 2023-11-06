package auth

import (
	"fmt"
	"log/slog"

	"github.com/bufbuild/protovalidate-go"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nova38/thesis/lib/go/fabric/state"
	"github.com/samber/oops"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	auth_pb "github.com/nova38/thesis/lib/go/gen/auth/v1"
)

const (
	// DefaultPageSize default page size
	DefaultPageSize int32 = 10000
)

var (
	// Validate that AuthTxCtx implements the required interfaces
	_ IAuthTxCtx                       = (*BaseTxCtx)(nil)
	_ state.LoggedTxCtxInterface       = (*BaseTxCtx)(nil)
	_ state.ValidateAbleTxCtxInterface = (*BaseTxCtx)(nil)
	_ state.PagedTxCtxInterface        = (*BaseTxCtx)(nil)

	// Global validator
	validator *protovalidate.Validator
)

type (
	AuthTransactionObjects struct {
		User       *auth_pb.User
		Identifier *auth_pb.Identifier
		Collection *auth_pb.Collection
		ops        *auth_pb.Operation
	}
	BaseTxCtx struct {
		contractapi.TransactionContext
		AuthTransactionObjects

		Logger   *slog.Logger
		PageSize int32

		// auth values
		authorized  bool
		authChecked bool
	}
)

// Unimplemented
func (ctx *BaseTxCtx) GetACLEntryKey() (key string, err error) {
	panic("Not Implemented In BaseTxCtx")
}

func (ctx *BaseTxCtx) GetViewMask() (mask fieldmaskpb.FieldMask, err error) {
	panic("Not Implemented In BaseTxCtx")
}

// =============================================

func (ctx *BaseTxCtx) HandelBefore() (err error) {
	ctx.Logger = slog.Default().With(
		"fn", ctx.GetFnName(),
		slog.Group(
			"tx info",
			"tx_id", ctx.GetStub().GetTxID(),
			"channel_id", ctx.GetStub().GetChannelID(),
		),
	)

	if validator == nil {
		v, err := protovalidate.New()
		if err != nil {
			panic(err)
		}
		validator = v
	}

	return nil
}

func (ctx *BaseTxCtx) HandleFnError(err *error, r any) {
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
func (ctx *BaseTxCtx) GetLogger() *slog.Logger {
	return ctx.Logger
}

// PagedTxCtxInterface functions
func (ctx *BaseTxCtx) GetPageSize() int32 {
	if ctx.PageSize == 0 {
		return DefaultPageSize
	}
	return ctx.PageSize
}

func (ctx *BaseTxCtx) SetPageSize(pageSize int32) {
	ctx.PageSize = pageSize
}

// ----------------------------------------------
// ValidateAbleTxCtxInterface functions
// ----------------------------------------------

func (ctx *BaseTxCtx) Validate(msg proto.Message) (err error) {
	if validator == nil {
		validator, err = protovalidate.New()
		if err != nil {
			return oops.Errorf("failed to create validator: %w", err)
		}
	}

	if msg == nil {
		return oops.Errorf("message is nil")
	}

	return oops.
		In(ctx.GetFnName()).
		Code(auth_pb.Error_ERROR_REQUEST_INVALID.String()).
		Wrap(validator.Validate(msg))
}

// ----------------------------------------------
// AuthTxCtxInterface functions
// ----------------------------------------------

func (ctx *BaseTxCtx) GetFnName() (name string) {
	name, _ = ctx.GetStub().GetFunctionAndParameters()
	return name
}

func (ctx *BaseTxCtx) MakeLastModified() (mod *auth_pb.StateActivity, err error) {
	user, err := ctx.GetUser()
	if err != nil {
		return nil, oops.Errorf("failed to get user: %w", err)
	}

	timestamp, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return nil, oops.Errorf("Failed to get timestamp: %w", err)
	}

	return &auth_pb.StateActivity{
		UserId:    user.GetId(),
		Note:      fmt.Sprintf("User %v modified the state", user.GetName()),
		TxId:      ctx.GetStub().GetTxID(),
		Timestamp: timestamp,
	}, nil
}

// =============================================
//  User Functions
// =============================================

func (ctx *BaseTxCtx) GetUserId() (userId *auth_pb.User_Id, err error) {
	userId = &auth_pb.User_Id{}

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

func (ctx *BaseTxCtx) GetUser() (user *auth_pb.User, err error) {
	if ctx.User != nil {
		return ctx.User, nil
	}

	id, err := ctx.GetUserId()
	if err != nil {
		return nil, oops.Wrap(err)
	}

	ctx.User = &auth_pb.User{Id: id}

	err = state.Get(ctx, ctx.User)

	if err != nil {
		return nil, oops.With("user_id", id).Wrap(err)
	}

	return ctx.User, nil
}

// =============================================
//  Collection Functions
// =============================================

func (ctx *BaseTxCtx) GetCollection() (col *auth_pb.Collection, err error) {
	if ctx.Collection != nil {
		return ctx.Collection, nil
	}

	return nil, oops.Errorf("collection not set")
}

func (ctx *BaseTxCtx) SetCollection(
	id *auth_pb.Collection_Id,
) (col *auth_pb.Collection, err error) {
	// TODO:
	// FIXME: Need to check in fns before calling is Authorized

	// See if the collection pointer has an ID and is not nil
	if id == nil || id.CollectionId == "" {
		return nil, oops.
			In("SetCollection").
			Code(auth_pb.Error_ERROR_COLLECTION_INVALID_ID.String()).
			Errorf("collection is nil or has no ID")
	}

	ctx.Collection = &auth_pb.Collection{
		Id: &auth_pb.Collection_Id{
			CollectionId: id.CollectionId,
		},
	}

	if err = state.Get(ctx, ctx.Collection); err != nil {
		return nil, oops.
			In("SetCollection").
			With("collectionId", id.CollectionId).
			Code(auth_pb.Error_ERROR_COLLECTION_UNREGISTERED.String()).
			Wrap(err)
	}

	return ctx.Collection, nil
}

// =============================================
//  Role Functions
// =============================================

func (ctx *BaseTxCtx) GetACLKey() (key string, err error) {
	// Check the requirements
	if ctx.User != nil || ctx.Collection != nil {
		return key, oops.
			With(
				"user", ctx.User,
				"collection", ctx.Collection,
			).
			Errorf("user or collection are nil")
	}

	// // Check to see if the user is a member of the collection
	// if cRole, ok := ctx.User.Memberships[ctx.Collection.Id.CollectionId]; ok {
	// 	return cRole.RoleId, nil
	// }

	ctx.GetLogger().Info("User is not a member of the collection, assigning to public")

	return "", nil
}

// =============================================
//  Operations Functions
// =============================================

func (ctx *BaseTxCtx) SetOperation(op *auth_pb.Operation) (err error) {
	// See if the operation pointer has an ID and is not nil
	if op == nil {
		return oops.Errorf("operation is nil or actions is nil")
	}

	if op.Domain == auth_pb.Operation_DOMAIN_UNSPECIFIED {
		return oops.Errorf("operation domain is unspecified")
	}
	if op.Action == auth_pb.Operation_ACTION_UNSPECIFIED {
		return oops.Errorf("operation action type is unspecified")
	}

	ctx.ops = &auth_pb.Operation{}

	ctx.ops.Action = op.Action
	ctx.ops.Domain = op.Domain

	if op.Paths != nil {
		ctx.ops.Paths = op.Paths
	}
	return nil
}

func (ctx *BaseTxCtx) GetOperations() (ops *auth_pb.Operation, err error) {
	if ctx.ops != nil {
		return ctx.ops, nil
	}
	return nil, oops.Errorf("operations not set")
}

func (ctx *BaseTxCtx) SetOperationsPaths(paths *fieldmaskpb.FieldMask) (err error) {
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
func (ctx *BaseTxCtx) Authorize() (auth bool, err error) {
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

	role, ok := user.Memberships[ctx.Collection.Id.CollectionId]
	if !ok {
		return false, oops.
			In("AuthorizeOperation").
			Code(auth_pb.Error_ERROR_COLLECTION_INVALID_ROLE_ID.String()).
			Errorf("Role %v is not valid for collection %v", role, ctx.Collection.Id.CollectionId)
	}

	// return AuthorizeOperation(ctx.ops, role.RoleId, ctx.Collection)
	return false, nil
}

func (ctx *BaseTxCtx) IsAuthorized() (err error) {
	if !ctx.authChecked {

		auth, err := ctx.Authorize()
		if err != nil {
			slog.Error(err.Error(), slog.Any("error", err))
			return oops.
				In(ctx.GetFnName()).
				Code(auth_pb.Error_ERROR_USER_PERMISSION_DENIED.String()).
				Wrap(err)
		}

		ctx.authorized = auth
	}
	return nil
}

func t() {
}
