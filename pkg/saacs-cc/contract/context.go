package contract

import (
	"log/slog"

	"github.com/bufbuild/protovalidate-go"
	"github.com/nova38/saacs/pkg/saacs-cc/auth/models"
	"github.com/nova38/saacs/pkg/saacs-cc/auth/policy"
	"github.com/nova38/saacs/pkg/saacs-cc/serializer"
	"github.com/nova38/saacs/pkg/saacs-cc/state"
	"github.com/samber/oops"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/nova38/saacs/pkg/saacs-cc/common"
	"github.com/nova38/saacs/pkg/saacs-cc/internal/biochain"
	cc "github.com/nova38/saacs/pkg/saacs-protos/saacs/chaincode/v0"
	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"

	authpb "github.com/nova38/saacs/pkg/saacs-protos/saacs/auth/v0"
)

var (
	// Validate that AuthCtxInterface implements the required interfaces
	_ common.TxCtxInterface = (*TxCtx)(nil)
	// Global validator
	validator *protovalidate.Validator
)

type (
	TxItems struct {
		User *pb.User
		// Collection *authpb.Collection
		Collections map[string]*authpb.Collection

		ops []*pb.Operation
	}

	TxCtx struct {
		contractapi.TransactionContext
		TxItems
		EnableSuggestion bool
		EnableHiddenTx   bool

		Logger   *slog.Logger
		PageSize int32

		Events []*pb.StateActivity
	}
)

// ═════════════════════════════════════════════

func (ctx *TxCtx) HandelBefore() (err error) {
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

	ctx.EnableHiddenTx = true
	ctx.EnableSuggestion = true

	ctx.Events = []*pb.StateActivity{}

	return nil
}

func (ctx *TxCtx) HandelAfter() (err error) {

	if ctx.ops != nil {
		event := &cc.OperationsPerformed{
			Operations: ctx.ops,
		}

		bytes, err := serializer.Marshal(event)
		if err != nil {
			ctx.LogError(err)
			return nil
		}
		ctx.GetStub().SetEvent("OperationsPerformed", bytes)
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
		*err = common.WrapError(*err)

		slog.Error((*err).Error())
	}
}

func (ctx *TxCtx) AppendEvent(event *pb.StateActivity) {
	ctx.Events = append(ctx.Events, event)
}

// Helper function to check if the bootstrap has been done
func (ctx *TxCtx) CheckBootstrap() (bool, error) {
	if v, err := ctx.GetStub().GetState(common.BootstrapKey); err != nil {
		return false, oops.Wrap(err)
	} else if v == nil {
		ctx.GetLogger().Info("Bootstrap not done")
		err := ctx.GetStub().PutState(common.BootstrapKey, []byte("true"))
		if err != nil {
			return false, oops.Wrap(err)
		}

		return false, nil
	}
	return true, nil

}

func (ctx *TxCtx) GetUser() (user *pb.User) {
	user, err := ctx.GetUserId()
	if err != nil {
		panic(err)
	}

	ctx.User = user

	return ctx.User
}

func (ctx *TxCtx) LogError(err error) {
	ctx.GetLogger().Error(err.Error(), slog.Any("error", err))
}

func (ctx *TxCtx) ErrorBase() oops.OopsErrorBuilder {
	return oops.OopsErrorBuilder{}.
		In(ctx.GetFnName()).
		User(ctx.GetUser().GetUserId(), ctx.GetUser().GetMspId())
}

func (ctx *TxCtx) CloseQueryIterator(resultIterator common.CommonIteratorInterface) {
	_ = resultIterator.Close()
}

// EnabledSuggestions returns true if the hidden tx feature is enabled,
// Enabled by default, can be disabled through build flags
// github.com/nova38/saacs/pkg/saacs-cc/common.EnabledSuggestions = ""
func (ctx *TxCtx) EnabledSuggestions() bool {
	return ctx.EnableSuggestion
}

// EnableHiddenTx returns true if the hidden tx feature is enabled,
// Enabled by default, can be disabled through build flags
// github.com/nova38/saacs/pkg/saacs-cc/common.EnableHiddenTx = ""
func (ctx *TxCtx) EnabledHidden() bool {
	return ctx.EnableHiddenTx
}

func (ctx *TxCtx) PostActionProcessing(
	item common.ItemInterface,
	ops []*pb.Operation,
) (err error) {

	if err := biochain.SpecimenPostActionProcessing(ctx, item, ops); err != nil {
		return oops.Wrap(err)
	}

	return nil
}

func (ctx *TxCtx) GetCollection(collectionId string) (col *authpb.Collection, err error) {
	if ctx.Collections == nil {
		ctx.Collections = map[string]*authpb.Collection{}
	}

	if col, ok := ctx.Collections[collectionId]; ok {
		return col, nil
	}

	col = &authpb.Collection{CollectionId: collectionId}

	if err := state.GetFromKey(ctx, col.StateKey(), col); err != nil {
		return nil, oops.Wrap(err)
	}

	ctx.Collections[collectionId] = col

	return col, nil
}

// ─────────────────────────────────────────────-
// LoggedTxCtxInterface
// ─────────────────────────────────────────────-

func (ctx *TxCtx) GetLogger() *slog.Logger {
	return ctx.Logger
}

// PagedTxCtxInterface functions

func (ctx *TxCtx) GetPageSize() int32 {
	if ctx.PageSize == 0 {
		return common.DefaultPageSize
	}
	return ctx.PageSize
}

func (ctx *TxCtx) SetPageSize(pageSize int32) {
	ctx.PageSize = pageSize
}

// ─────────────────────────────────────────────-
// ValidateAbleTxCtxInterface functions
// ─────────────────────────────────────────────-

func (ctx *TxCtx) Validate(msg proto.Message) (err error) {
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
		Code(pb.TxError_REQUEST_INVALID.String()).
		Wrap(validator.Validate(msg))
}

// ─────────────────────────────────────────────-
// AuthTxCtxInterface functions
// ─────────────────────────────────────────────-

func (ctx *TxCtx) GetFnName() (name string) {
	name, _ = ctx.GetStub().GetFunctionAndParameters()
	return name
}

func (ctx *TxCtx) MakeLastModified() (mod *pb.StateActivity, err error) {
	user, err := ctx.GetUserId()
	if err != nil {
		return nil, oops.Errorf("failed to get user: %w", err)
	}

	timestamp, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return nil, oops.Errorf("Failed to get timestamp: %w", err)
	}

	mod = &pb.StateActivity{
		UserId: user.GetUserId(),
		MspId:  user.GetMspId(),
		// Note:      fmt.Sprintf("User %v modified the state", user.GetName()),
		TxId: ctx.GetStub().GetTxID(),
		Timestamp: &timestamppb.Timestamp{
			Seconds: timestamp.GetSeconds(),
			Nanos:   timestamp.GetNanos(),
		},
	}

	ctx.Logger.Info("MakeLastModified", slog.Any("mod", mod))

	return mod, nil
}

// ═════════════════════════════════════════════
//  User Functions
// ═════════════════════════════════════════════

func (ctx *TxCtx) GetUserId() (user *pb.User, err error) {
	// Extract The info from the Client ID
	id := ctx.GetClientIdentity()
	// cert, err := id.GetX509Certificate()
	// if err != nil {
	// return nil, oops.Errorf("failed to get user certificate from CID: %s", err)
	// }
	userId, err := id.GetID()
	if err != nil {
		return nil, oops.Errorf("failed to get user ID from CID: %s", err)
	}

	mspId, err := id.GetMSPID()
	if err != nil {
		return nil, oops.Errorf("failed to get user ID from CID: %s", err)
	}

	// ctx.Logger.Info("GetUserId",
	// 	slog.Any("cert", cert),
	// 	slog.Any("userId", userId),
	// 	slog.Any("mspId", mspId))

	return &pb.User{MspId: mspId, UserId: userId}, nil
}

// ═════════════════════════════════════════════
//
//	ACL Functions
//
// ═════════════════════════════════════════════

func (ctx *TxCtx) GetSubLogger(name string) *slog.Logger {
	return ctx.Logger.WithGroup(name)
}

func (ctx *TxCtx) Authorize(operations []*pb.Operation) (auth bool, err error) {

	for _, op := range operations {
		if ctx.ops == nil {
			ctx.ops = []*pb.Operation{}
		}
		copy := proto.Clone(op).(*pb.Operation)
		ctx.ops = append(ctx.ops, copy)
	}

	opsByCol := make(map[string][]*pb.Operation)

	// Group the operations by collection
	for _, op := range operations {
		if op.GetAction() == pb.Action_ACTION_CREATE &&
			op.GetItemType() == common.CollectionItemType {
			ctx.GetLogger().Info("Operation Is to Create Collection")
			// TODO: Implement Special auth case for creating a collection
		} else {
			if _, ok := opsByCol[op.GetCollectionId()]; !ok {
				opsByCol[op.GetCollectionId()] = []*pb.Operation{}
			}
			opsByCol[op.GetCollectionId()] = append(opsByCol[op.GetCollectionId()], op)
		}

	}

	for colId, ops := range opsByCol {
		var (
			col        *authpb.Collection
			Authorizer common.Authorizer
		)

		if col, err = ctx.GetCollection(colId); err != nil {
			if colId == common.USERCOLID {
				return false, oops.
					Hint("User Collection Has not been created").
					In("Authorization").
					With("operation", ops).
					Wrap(err)
			}
			return false, oops.
				In("Authorization").
				With("operation", ops).
				Wrap(err)
		}

		// Validate Operation
		for _, op := range ops {
			if valid, err := policy.ValidateOperation(col, op); err != nil {
				return false, oops.Wrap(err)
			} else if !valid {
				return false, oops.Wrap(common.RuntimeBadOps)
			}
		}

		switch col.GetAuthType() {
		case authpb.AuthType_AUTH_TYPE_NONE:
			Authorizer = &models.NoAuth{
				Collection: col,
				Logger:     ctx.GetSubLogger("NoAuth"),
				TxCtx:      ctx,
			}

		case authpb.AuthType_AUTH_TYPE_ROLE:
			Authorizer = &models.RBAC{
				Collection: col,
				Logger:     ctx.GetSubLogger("RBAC"),
				TxCtx:      ctx,
				UserRoles:  make(map[string][]*authpb.Role),
			}
		case authpb.AuthType_AUTH_TYPE_GLOBAL_ROLE:
			Authorizer = &models.GRBAC{
				Collection: col,
				Logger:     ctx.GetSubLogger("GRBAC"),
				TxCtx:      ctx,
				UserRoles:  make(map[string][]*authpb.Role),
			}

		case authpb.AuthType_AUTH_TYPE_IDENTITY:
			Authorizer = &models.IAC{
				Collection:            col,
				CollectionMemberships: map[string]*authpb.UserDirectMembership{},
				Logger:                ctx.GetSubLogger("IAC"),
				TxCtx:                 ctx,
			}

		case authpb.AuthType_AUTH_TYPE_ATTRIBUTE:
			Authorizer = &models.ABAC{
				Collection:     col,
				UserAttributes: []*authpb.Attribute{},
				Logger:         ctx.GetSubLogger("ABAC"),
				TxCtx:          ctx,
			}

		case authpb.AuthType_AUTH_TYPE_UNSPECIFIED:
			return false, oops.Wrap(common.CollectionInvalid)
		default:
			return false, oops.Wrap(common.RuntimeBadOps)
		}

		for _, op := range ops {
			if auth, err := Authorizer.Authorize(op); err != nil {
				return false, oops.Wrap(err)
			} else if !auth {
				ctx.GetLogger().Info("User is not authorized")
				return false, nil
			}
		}

	}

	return true, nil
}
