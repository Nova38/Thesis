package rbac

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
)

var _ common.TxCtxInterface = (*AuthTxCtx)(nil)

// TxItems Extractors
type (
	CollectionHolder interface {
		GetCollection() *authpb.Collection
	}
	CollectionIdHolder interface {
		GetCollectionId() string
	}
	UserHolder interface {
		GetUser() *authpb.User
	}
	UserIdHolder interface {
		GetUserId() string
		GetMspId() string
	}

	AuthTxCtxInterface interface {
		ExtractAuthTransactionItems(req interface{}) (err error)
		common.TxCtxInterface
	}

	// AuthTxCtx is a wrapper around the contractapi.TransactionContext
	AuthTxCtx struct {
		state.BaseTxCtx
	}
)

//func (ctx *AuthTxCtx) ExtractAuthTransactionItems(req interface{}) (err error) {
//	if col, ok := req.(CollectionHolder); ok {
//		ctx.Collection = col.GetCollection()
//		if ctx.Collection == nil {
//			return oops.
//				In(ctx.GetFnName()).
//				Code(authpb.TxError_COLLECTION_INVALID.String()).
//				Errorf("collection is nil")
//		}
//	}
//
//	if col_id, ok := req.(CollectionIdHolder); ok {
//		if _, err = ctx.SetCollection(col_id.GetCollectionId()); err != nil {
//			return oops.
//				In(ctx.GetFnName()).
//				Code(authpb.TxError_COLLECTION_INVALID.String()).
//				Wrap(err)
//		}
//	}
//
//	if user, ok := req.(UserHolder); ok {
//		ctx.User = user.GetUser()
//		if ctx.User == nil {
//			return oops.
//				In(ctx.GetFnName()).
//				Code(authpb.TxError_USER_INVALID.String()).
//				Errorf("user is nil")
//		}
//	}
//
//	if userId, ok := req.(UserIdHolder); ok {
//		userId := userId.GetUserId()
//		if userId == nil {
//			return oops.
//				In(ctx.GetFnName()).
//				Code(authpb.TxError_USER_INVALID.String()).
//				Errorf("user id is nil")
//		}
//		ctx.User = &authpb.User{
//			Id: userId,
//		}
//		if err := state.Get(ctx, ctx.User); err != nil {
//			return oops.
//				In(ctx.GetFnName()).
//				Code(authpb.TxError_USER_INVALID.String()).
//				Wrap(err)
//		}
//
//	}
//
//	return nil
//}
