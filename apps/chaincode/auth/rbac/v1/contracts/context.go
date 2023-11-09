package contracts

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	auth_pb "github.com/nova38/thesis/lib/go/gen/auth/v1"

	auth "github.com/nova38/thesis/lib/go/fabric/auth"
	"github.com/samber/oops"
)

var _ auth.TxCtxInterface = (*AuthTxCtx)(nil)

// TxObjects Extractors
type (
	CollectionHolder interface {
		GetCollection() *auth_pb.Collection
	}
	CollectionIdHolder interface {
		GetCollectionId() string
	}
	UserHolder interface {
		GetUser() *auth_pb.User
	}
	UserIdHolder interface {
		GetUserId() *auth_pb.User_Id
	}

	AuthTxCtxInterface interface {
		ExtractAuthTransactionItems(req interface{}) (err error)
		auth.TxCtxInterface
	}

	// AuthTxCtx is a wrapper around the contractapi.TransactionContext
	AuthTxCtx struct {
		auth.TxCtx
	}
)

func (ctx *AuthTxCtx) HandelBefore() (err error) {
	return ctx.BaseHandelBefore()
}

func (ctx *AuthTxCtx) ExtractAuthTransactionItems(req interface{}) (err error) {
	if col, ok := req.(CollectionHolder); ok {
		ctx.Collection = col.GetCollection()
		if ctx.Collection == nil {
			return oops.
				In(ctx.GetFnName()).
				Code(auth_pb.TxError_TX_ERROR_COLLECTION_INVALID.String()).
				Errorf("collection is nil")
		}
	}

	if col_id, ok := req.(CollectionIdHolder); ok {
		if _, err = ctx.SetCollection(col_id.GetCollectionId()); err != nil {
			return oops.
				In(ctx.GetFnName()).
				Code(auth_pb.TxError_TX_ERROR_COLLECTION_INVALID.String()).
				Wrap(err)
		}
	}

	if user, ok := req.(UserHolder); ok {
		ctx.User = user.GetUser()
		if ctx.User == nil {
			return oops.
				In(ctx.GetFnName()).
				Code(auth_pb.TxError_TX_ERROR_USER_INVALID.String()).
				Errorf("user is nil")
		}
	}

	if user_id, ok := req.(UserIdHolder); ok {
		user_id := user_id.GetUserId()
		if user_id == nil {
			return oops.
				In(ctx.GetFnName()).
				Code(auth_pb.TxError_TX_ERROR_USER_INVALID.String()).
				Errorf("user id is nil")
		}
		ctx.User = &auth_pb.User{
			Id: user_id,
		}
		if err := state.Get(ctx, ctx.User); err != nil {
			return oops.
				In(ctx.GetFnName()).
				Code(auth_pb.TxError_TX_ERROR_USER_INVALID.String()).
				Wrap(err)
		}

	}

	return nil
}
