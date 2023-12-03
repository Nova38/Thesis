package identity

import (
	"log/slog"

	"github.com/nova38/thesis/packages/fabric/auth/policy"
	"github.com/nova38/thesis/packages/fabric/auth/state"
	"github.com/nova38/thesis/packages/saacs/auth/common"
	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	"github.com/samber/oops"
)

func (ctx *IdentityTxCtx) GetUserMembership(
	collectionId string,
) (*authpb.UserMembership, error) {

	if ctx.CollectionMemberships == nil {
		ctx.CollectionMemberships = make(map[string]*authpb.UserMembership)
	}

	if membership, ok := ctx.CollectionMemberships[collectionId]; ok {
		return membership, nil
	}

	user, err := ctx.GetUserId()
	if err != nil {
		return nil, err
	}

	membership := &authpb.UserMembership{
		CollectionId: collectionId,
		MspId:        user.GetUserId(),
		UserId:       user.GetMspId(),
	}

	if key, err := common.MakePrimaryKey(membership); err != nil {
		return nil, oops.Wrap(err)
	} else if err = state.Get(ctx, key, membership); err != nil {
		return nil, err
	}

	ctx.CollectionMemberships[collectionId] = membership

	return membership, nil
}

// ──────────────────────────────── Query ────────────────────────────────────────

func (ctx *IdentityTxCtx) Authorize(ops []*authpb.Operation) (bool, error) {

	ctx.GetLogger().Info("NoAuthContract.Authenticate")

	for _, op := range ops {
		if auth, err := ctx.authorized(op); err != nil {
			return false, oops.Wrap(err)
		} else if !auth {
			ctx.Logger.Info("User is not authorized")
			return false, nil
		}
	}

	return true, nil
}

func (ctx *IdentityTxCtx) authorized(op *authpb.Operation) (bool, error) {
	ctx.GetLogger().Info(op.String())

	// Handle special case of creating a collection
	if op.GetItemType() == "auth.Collection" {
		if op.GetAction() == authpb.Action_ACTION_CREATE {
			ctx.Logger.Info(
				"User is authorized to create a collection",
				slog.Group("auth", "collection", op.GetCollectionId()),
			)
			return true, nil
		}
	}

	// Get the collection
	col, err := ctx.GetCollection(op.GetCollectionId())
	if err != nil {
		return false, oops.Wrap(err)
	}

	// Validate the operation
	if valid, err := policy.ValidateOperation(col, op); err != nil {
		return false, oops.Hint("Invalid Operation").Wrap(err)
	} else if !valid {
		return false, nil
	}

	// TODO: Check if the action is allowed for any user in the collection

	// ═════════════════════════════════════════════
	// Default Policy
	// ═════════════════════════════════════════════

	if auth, err := policy.AuthorizedPolicy(col.GetDefault(), op); err != nil {
		return false, oops.Wrap(err)
	} else if auth {
		ctx.Logger.Info("User is authorized by default")
		return true, nil
	}
	// ═════════════════════════════════════════════
	// Check Membership
	// ═════════════════════════════════════════════

	// Get the user membership
	membership, err := ctx.GetUserMembership(op.GetCollectionId())
	if err != nil {
		return false, oops.Wrap(err)
	}
	if membership == nil {
		ctx.Logger.Info(
			"User is not a member of the collection",
			slog.Group("auth", "collection", op.GetCollectionId()),
		)
		return false, nil
	}

	if auth, err := policy.AuthorizedPolicy(membership.GetPolices(), op); err != nil {
		return false, oops.Wrap(err)
	} else if auth {
		ctx.Logger.Info("User is authorized by membership")
		return true, nil
	}

	return false, nil
}
