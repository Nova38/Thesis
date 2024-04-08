package identity

import (
	"log/slog"

	"github.com/nova38/saacs/apps/saacs-cc/auth/policy"
	"github.com/nova38/saacs/apps/saacs-cc/common"
	"github.com/nova38/saacs/apps/saacs-cc/state"
	authpb "github.com/nova38/saacs/libs/saacs-protos-go/auth/v1"
	"github.com/samber/oops"
)

type IAC struct {
	Collections           map[string]*authpb.Collection
	CollectionMemberships map[string]*authpb.UserDirectMembership
	TxCtx                 common.TxCtxInterface
	Logger                *slog.Logger
}

func (ac *IAC) Authorize(ops []*authpb.Operation) (bool, error) {

	ac.Logger.Info("NoAuthContract.Authenticate")

	for _, op := range ops {
		if auth, err := ac.authorized(op); err != nil {
			return false, oops.Wrap(err)
		} else if !auth {
			ac.Logger.Info("User is not authorized")
			return false, nil
		}
	}

	return true, nil
}

func (ac *IAC) GetUserDirectMembership(
	collectionId string,
) (*authpb.UserDirectMembership, error) {

	if ac.CollectionMemberships == nil {
		ac.CollectionMemberships = make(map[string]*authpb.UserDirectMembership)
	}

	// if membership, ok := actions.CollectionMemberships[collectionId]; ok {
	// 	return membership, nil
	// }

	user, err := ac.TxCtx.GetUserId()
	if err != nil {
		return nil, err
	}

	membership := &authpb.UserDirectMembership{
		CollectionId: collectionId,
		MspId:        user.GetMspId(),
		UserId:       user.GetUserId(),
	}

	if err = state.Get(ac.TxCtx, membership); err != nil {
		return nil, oops.Wrap(err)
	}

	ac.CollectionMemberships[collectionId] = membership

	return membership, nil
}

// ──────────────────────────────── Query ────────────────────────────────────────

func (ac *IAC) authorized(op *authpb.Operation) (bool, error) {
	ac.Logger.Info(op.String())

	// Handle special case of creating a collection
	if op.GetItemType() == "auth.Collection" {
		if op.GetAction() == authpb.Action_ACTION_CREATE {
			ac.Logger.Info(
				"User is authorized to create a collection",
				slog.Group("auth", "collection", op.GetCollectionId()),
			)
			return true, nil
		}
	}

	// Get the collection
	col, err := ac.TxCtx.GetCollection(op.GetCollectionId())
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
		ac.Logger.Info("User is authorized by default")
		return true, nil
	}
	// ═════════════════════════════════════════════
	// Check Membership
	// ═════════════════════════════════════════════

	// Get the user membership
	membership, err := ac.GetUserDirectMembership(op.GetCollectionId())
	if err != nil {
		return false, oops.Wrap(err)
	}
	if membership == nil {
		ac.Logger.Info(
			"User is not a member of the collection",
			slog.Group("auth", "collection", op.GetCollectionId()),
		)
		return false, nil
	}

	if auth, err := policy.AuthorizedPolicy(membership.GetPolices(), op); err != nil {
		return false, oops.Wrap(err)
	} else if auth {
		ac.Logger.Info("User is authorized by membership")
		return true, nil
	}

	return false, nil
}
