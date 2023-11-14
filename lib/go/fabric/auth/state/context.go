package state

import (
	"log/slog"

	"github.com/nova38/thesis/lib/go/fabric/auth/common"

	"github.com/hyperledger/fabric-chaincode-go/shim"

	"github.com/bufbuild/protovalidate-go"
	"github.com/samber/oops"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

var (
	// Validate that AuthCtxInterface implements the required interfaces
	_ TxCtxInterface = (*BaseTxCtx)(nil)
	// Global validator
	validator *protovalidate.Validator
)

type (
	TxObjects struct {
		User       *authpb.User
		Collection *authpb.Collection
		ops        *authpb.Operation
	}

	BaseTxCtx struct {
		contractapi.TransactionContext
		TxObjects

		Logger   *slog.Logger
		PageSize int32

		// This is the chaincode function name to call for authorization
		authFn      AuthFunc
		authorized  bool
		authChecked bool
	}
)

//func (ctx *BaseTxCtx) AddExtractorFunc(name string, fn ExtractorFunc) {
//	// TODO implement me
//	panic("implement me")
//}

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

func (ctx *BaseTxCtx) CloseQueryIterator(resultIterator shim.CommonIteratorInterface) {
	_ = resultIterator.Close()
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
		return common.DefaultPageSize
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
		Code(authpb.TxError_REQUEST_INVALID.String()).
		Wrap(validator.Validate(msg))
}

// ----------------------------------------------
// AuthTxCtxInterface functions
// ----------------------------------------------

func (ctx *BaseTxCtx) GetFnName() (name string) {
	name, _ = ctx.GetStub().GetFunctionAndParameters()
	return name
}

func (ctx *BaseTxCtx) MakeLastModified() (mod *authpb.StateActivity, err error) {
	userId, mspId, err := ctx.GetUserId()
	if err != nil {
		return nil, oops.Errorf("failed to get user: %w", err)
	}

	timestamp, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return nil, oops.Errorf("Failed to get timestamp: %w", err)
	}

	return &authpb.StateActivity{
		UserId: userId,
		MspId:  mspId,
		// Note:      fmt.Sprintf("User %v modified the state", user.GetName()),
		TxId:      ctx.GetStub().GetTxID(),
		Timestamp: timestamp,
	}, nil
}

// =============================================
//  User Functions
// =============================================

func (ctx *BaseTxCtx) GetUserId() (mspId string, userId string, err error) {
	// Extract The info from the Client ID
	id := ctx.GetClientIdentity()

	userId, err = id.GetID()
	if err != nil {
		return "", "", oops.Errorf("failed to get user certificate from CID: %s", err)
	}

	mspId, err = id.GetMSPID()
	if err != nil {
		return "", "", oops.Errorf("failed to get user ID from CID: %s", err)
	}

	return mspId, userId, nil
}

// func (ctx *BaseTxCtx) GetUser() (user *authpb.User, err error) {
// 	if ctx.User != nil {
// 		return ctx.User, nil
// 	}

// 	mspId, userId, err := ctx.GetUserId()
// 	if err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	ctx.User = &authpb.User{
// 		MspId:  mspId,
// 		UserId: userId,
// 	}

// 	// err = state.Get()

// 	return ctx.User, nil
// }

// ════════════════════════════════════════════════════════
//  Collection Functions
// ════════════════════════════════════════════════════════

// func (ctx *BaseTxCtx) GetCollection() (col *authpb.Collection, err error) {
// 	if ctx.Collection != nil {
// 		return ctx.Collection, nil
// 	}

// 	return nil, oops.Errorf("collection not set")
// }

// func (ctx *BaseTxCtx) SetCollection(
// 	collectionId string,
// ) (col *authpb.Collection, err error) {
// 	// TODO:
// 	// FIXME: Need to check in fns before calling is Authorized

// 	// See if the collection pointer has an ID and is not nil
// 	if collectionId == "" {
// 		return nil, oops.
// 			In("SetCollection").
// 			Code(authpb.TxError_COLLECTION_INVALID_ID.String()).
// 			Errorf("collection is nil or has no ID")
// 	}

// 	ctx.Collection = &authpb.Collection{
// 		CollectionId: collectionId,
// 	}

// 	if err = get(ctx, ctx.Collection); err != nil {
// 		return nil, oops.
// 			In("SetCollection").
// 			With("collectionId", collectionId).
// 			Code(authpb.TxError_COLLECTION_UNREGISTERED.String()).
// 			Wrap(err)
// 	}

// 	return ctx.Collection, nil
// }

// ════════════════════════════════════════════════════════
// Role Functions
// ════════════════════════════════════════════════════════

//func (ctx *BaseTxCtx) GetACLKey() (key string, err error) {
//	// Check the requirements
//	if ctx.User != nil || ctx.Collection != nil {
//		return key, oops.
//			With(
//				"user", ctx.User,
//				"collection", ctx.Collection,
//			).
//			Errorf("user or collection are nil")
//	}
//
//	// // Check to see if the user is a member of the collection
//	// if cRole, ok := ctx.User.Memberships[ctx.Collection.Id.CollectionId]; ok {
//	// 	return cRole.RoleId, nil
//	// }
//
//	ctx.GetLogger().Info("User is not a member of the collection, assigning to public")
//
//	return "", nil
//}

// =============================================
//  Operations Functions
// =============================================

//func (ctx *BaseTxCtx) SetOperation(op *authpb.Operation) {
//	// See if the operation pointer has an ID and is not nil
//	ctx.ops = op
//}
//
//func (ctx *BaseTxCtx) GetOperations() (ops *authpb.Operation, err error) {
//	if ctx.ops != nil {
//		return ctx.ops, nil
//	}
//	return nil, oops.Errorf("operations not set")
//}
//
//func (ctx *BaseTxCtx) SetOperationsPaths(paths *fieldmaskpb.FieldMask) (err error) {
//	if paths == nil {
//		return oops.Errorf("paths is nil")
//	}
//
//	if ctx.ops == nil {
//		return oops.Errorf("operations is nil")
//	}
//	ctx.ops.Paths = paths
//
//	return nil
//}

// =============================================
//
//	ACL Functions
//
// =============================================

//func (ctx *BaseTxCtx) ExtractAuthTransactionItems(req any) (err error) {
//	// TODO implement me
//	panic("implement me")
//}

func (ctx *BaseTxCtx) SetAuthenticator(fn AuthFunc) {
	ctx.authFn = fn
}

func (ctx *BaseTxCtx) GetAuthenticator() AuthFunc {
	return ctx.authFn
}

func (ctx *BaseTxCtx) GetViewMask() (mask *fieldmaskpb.FieldMask) {
	// TODO implement me

	if ctx.ops == nil {
		ctx.ops = &authpb.Operation{}
	}
	if ctx.ops.Paths == nil {
		ctx.ops.Paths = &fieldmaskpb.FieldMask{}
	}
	return ctx.ops.Paths
}

func (ctx *BaseTxCtx) Authorize(ops []*authpb.Operation) (auth bool, err error) {
	// if the user is already authorized, return the value
	if ctx.authChecked {
		return ctx.authorized, nil
	}
	ctx.authChecked = true

	if len(ops) == 0 {
		return false, oops.Errorf("operations is empty")
	}

	ctx.ops = ops[0]

	fn := ctx.GetAuthenticator()

	// Check if all the objects are set
	if fn == nil {
		return false, oops.Errorf("authenticator function is not set")
	}

	// Call the authenticator function
	auth, err = fn(ctx, ops)

	ctx.Logger.Info("Authorize", slog.Any("auth", auth), slog.Any("err", err))

	if auth {
		// newOps := &authpb.Operation{}

		// if err = proto.Unmarshal(res.Payload, newOps); err == nil {
		// 	ctx.ops.Paths = newOps.Paths
		// }

		return true, nil
	}
	return false, oops.With("operation", ops).Wrapf(err, "failed to authorize operation")
}
