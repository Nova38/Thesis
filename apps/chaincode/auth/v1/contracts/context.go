package contract

import (
	"github.com/nova38/thesis/lib/go/fabric/rbac"
	"github.com/nova38/thesis/lib/go/fabric/state"
	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
	"github.com/samber/oops"
)

var _ rbac.AuthTxCtxInterface = (*AuthTxCtx)(nil)

// AuthTransactionObjects Extractors
type (
	CollectionHolder interface {
		GetCollection() *rbac_pb.Collection
	}
	CollectionIdHolder interface {
		GetCollectionId() *rbac_pb.Collection_Id
	}
	UserHolder interface {
		GetUser() *rbac_pb.User
	}
	UserIdHolder interface {
		GetUserId() *rbac_pb.User_Id
	}

	AuthTxCtxInterface interface {
		ExtractAuthTransactionItems(req interface{}) (err error)
		rbac.AuthTxCtxInterface
	}

	// AuthTxCtx is a wrapper around the contractapi.TransactionContext
	AuthTxCtx struct {
		rbac.TxCtx
	}
)

func (ctx *AuthTxCtx) ExtractAuthTransactionItems(req interface{}) (err error) {
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
		if _, err = ctx.SetCollection(col_id.GetCollectionId()); err != nil {
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
