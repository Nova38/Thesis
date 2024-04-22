package models

import (
	"log/slog"

	"github.com/nova38/saacs/apps/saacs-cc/auth/policy"
	"github.com/nova38/saacs/apps/saacs-cc/common"
	"github.com/nova38/saacs/apps/saacs-cc/state"
	authpb "github.com/nova38/saacs/pkg/saacs-protos/auth/v1"
	"github.com/samber/oops"
)

type IAC struct {
	Collection            *authpb.Collection
	CollectionMemberships map[string]*authpb.UserDirectMembership
	TxCtx                 common.TxCtxInterface
	Logger                *slog.Logger
}

func (ac *IAC) Authorize(op *authpb.Operation) (bool, error) {

	ac.Logger.Info("Authenticate")

	// ═════════════════════════════════════════════
	// Default Policy
	// ═════════════════════════════════════════════

	if auth, err := policy.AuthorizedPolicy(ac.Collection.GetDefault(), op); err != nil {
		return false, oops.Wrap(err)
	} else if auth {
		ac.Logger.Info("User is authorized by default")
		return true, nil
	}
	// ═════════════════════════════════════════════
	// Check Membership
	// ═════════════════════════════════════════════

	// Get the user membership
	membership, err := ac.getUserDirectMembership(op.GetCollectionId())
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

func (ac *IAC) getUserDirectMembership(
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
