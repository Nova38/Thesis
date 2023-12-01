package state

import (
	"log/slog"

	"github.com/bufbuild/protovalidate-go"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/samber/oops"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/nova38/thesis/lib/go/fabric/auth/common"

	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
)

var (
	// Validate that AuthCtxInterface implements the required interfaces
	_ common.TxCtxInterface = (*BaseTxCtx)(nil)
	// Global validator
	validator *protovalidate.Validator
)

type (
	RawLedger[T common.ItemInterface] struct {
	}
	TxItems struct {
		User *authpb.User
		// Collection *authpb.Collection
		Collections map[string]*authpb.Collection

		ops *authpb.Operation
	}

	BaseTxCtx struct {
		contractapi.TransactionContext
		TxItems

		Logger   *slog.Logger
		PageSize int32
	}
)

// ═════════════════════════════════════════════

func (ctx *BaseTxCtx) HandelBefore() (err error) {
	ctx.Logger = slog.Default().With(
		"fn", ctx.GetFnName(),
		slog.Group(
			"tx info",
			"tx_id", ctx.GetStub().GetTxID(),
			"channel_id", ctx.GetStub().GetChannelID(),
		),
	)

	ctx.User, err = ctx.GetUserId()
	if err != nil {
		return oops.Errorf("failed to get user: %w", err)
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
		*err = common.WrapError(*err)
		slog.Error((*err).Error())
	}
}

// Helper function to check if the bootstrap has been done
func (ctx *BaseTxCtx) CheckBootstrap() (bool, error) {
	if v, err := ctx.GetStub().GetState(common.BootstrapKey); err != nil {
		return false, oops.Wrap(err)
	} else if v == nil && err == nil {
		ctx.GetLogger().Info("Bootstrap not done")
		ctx.GetStub().PutState(common.BootstrapKey, []byte("true"))

		return false, nil
	}
	return true, nil

}

func (ctx *BaseTxCtx) LogError(err error) {
	ctx.LogError(err)
}

func (ctx *BaseTxCtx) ErrorBase() oops.OopsErrorBuilder {
	return oops.OopsErrorBuilder{}.
		In(ctx.GetFnName()).
		User(ctx.User.GetUserId(), ctx.User.GetMspId())
}

func (ctx *BaseTxCtx) CloseQueryIterator(resultIterator shim.CommonIteratorInterface) {
	_ = resultIterator.Close()
}

// EnabledSuggestions returns true if the hidden tx feature is enabled,
// Enabled by default, can be disabled through build flags
// github.com/nova38/thesis/lib/go/fabric/auth/common.EnabledSuggestions = ""
func (ctx *BaseTxCtx) EnabledSuggestions() bool {
	return common.EnableSuggestion == common.EnableHiddenTxValue
}

// EnableHiddenTx returns true if the hidden tx feature is enabled,
// Enabled by default, can be disabled through build flags
// github.com/nova38/thesis/lib/go/fabric/auth/common.EnableHiddenTx = ""
func (ctx *BaseTxCtx) EnabledHidden() bool {
	return common.EnableHiddenTx == common.EnableHiddenTxValue
}

func (ctx *BaseTxCtx) PostActionProcessing(
	item common.ItemInterface,
	ops []*authpb.Operation,
) (err error) {
	return nil
}

func (ctx *BaseTxCtx) GetCollection(collectionId string) (col *authpb.Collection, err error) {
	if ctx.Collections == nil {
		ctx.Collections = map[string]*authpb.Collection{}
	}

	if col, ok := ctx.Collections[collectionId]; ok {
		return col, nil
	}

	col = &authpb.Collection{CollectionId: collectionId}
	if key, err := common.MakePrimaryKey(col); err != nil {
		return nil, oops.Wrap(err)
	} else if err := Get(ctx, key, col); err != nil {
		return nil, oops.Wrap(err)
	}

	ctx.Collections[collectionId] = col

	return col, nil
}

// ─────────────────────────────────────────────-
// LoggedTxCtxInterface
// ─────────────────────────────────────────────-

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

// ─────────────────────────────────────────────-
// ValidateAbleTxCtxInterface functions
// ─────────────────────────────────────────────-

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

// ─────────────────────────────────────────────-
// AuthTxCtxInterface functions
// ─────────────────────────────────────────────-

func (ctx *BaseTxCtx) GetFnName() (name string) {
	name, _ = ctx.GetStub().GetFunctionAndParameters()
	return name
}

func (ctx *BaseTxCtx) MakeLastModified() (mod *authpb.StateActivity, err error) {
	user, err := ctx.GetUserId()
	if err != nil {
		return nil, oops.Errorf("failed to get user: %w", err)
	}

	timestamp, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return nil, oops.Errorf("Failed to get timestamp: %w", err)
	}

	return &authpb.StateActivity{
		UserId: user.GetUserId(),
		MspId:  user.GetMspId(),
		// Note:      fmt.Sprintf("User %v modified the state", user.GetName()),
		TxId:      ctx.GetStub().GetTxID(),
		Timestamp: timestamp,
	}, nil
}

// ═════════════════════════════════════════════
//  User Functions
// ═════════════════════════════════════════════

func (ctx *BaseTxCtx) GetUserId() (user *authpb.User, err error) {
	// Extract The info from the Client ID
	id := ctx.GetClientIdentity()

	userId, err := id.GetID()
	if err != nil {
		return nil, oops.Errorf("failed to get user certificate from CID: %s", err)
	}

	mspId, err := id.GetMSPID()
	if err != nil {
		return nil, oops.Errorf("failed to get user ID from CID: %s", err)
	}

	return &authpb.User{MspId: mspId, UserId: userId}, nil
}

// func (ctx *BaseTxCtx) GetUser() (*authpb.User, error) {
// 	if ctx.User != nil {
// 		return ctx.User, nil
// 	}
// 	user, err := ctx.GetUserId()
// 	if err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	ctx.User = user

// 	Get(ctx, ctx.User)

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

// 	if key, err := common.MakePrimaryKey(ctx.Collection); err != nil {
// 		return nil, oops.Wrap(err)
// 	} else if err = Get(ctx, key, ctx.Collection); err != nil {
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

// func (ctx *BaseTxCtx) GetACLKey() (key string, err error) {
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

// ═════════════════════════════════════════════
//  Operations Functions
// ═════════════════════════════════════════════

// func (ctx *BaseTxCtx) SetOperation(op *authpb.Operation) {
//	// See if the operation pointer has an ID and is not nil
//	ctx.ops = op
//}
//
// func (ctx *BaseTxCtx) GetOperations() (ops *authpb.Operation, err error) {
//	if ctx.ops != nil {
//		return ctx.ops, nil
//	}
//	return nil, oops.Errorf("operations not set")
//}
//
// func (ctx *BaseTxCtx) SetOperationsPaths(paths *fieldmaskpb.FieldMask) (err error) {
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

// ═════════════════════════════════════════════
//
//	ACL Functions
//
// ═════════════════════════════════════════════

// func (ctx *BaseTxCtx) ExtractAuthTransactionItems(req any) (err error) {
//	// TODO implement me
//	panic("implement me")
//}

func (ctx *BaseTxCtx) GetViewMask() (mask *fieldmaskpb.FieldMask) {
	// TODO implement me

	if ctx.ops == nil {
		ctx.ops = &authpb.Operation{}
	}
	if ctx.ops.GetPaths() == nil {
		ctx.ops.Paths = &fieldmaskpb.FieldMask{}
	}

	return ctx.ops.GetPaths()
}

func (ctx *BaseTxCtx) Authorize(ops []*authpb.Operation) (auth bool, err error) {
	panic("Should Never Be Called, Implement in child class")
}
