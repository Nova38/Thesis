package contracts

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/oops"
)

// TODO: Implement NoAuth authentification contract

// Get The user membership for the given collection
func GetUserMembership(ctx common.TxCtxInterface, collectionId string) (*authpb.UserMembership, error) {
	user, err := ctx.GetUserId()
	if err != nil {
		return nil, err
	}

	membership := &authpb.UserMembership{
		CollectionId: collectionId,
		MspId:        user.GetUserId(),
		UserId:       user.GetMspId(),
	}

	if err = state.Get(ctx, membership); err != nil {
		return nil, err
	}

	return membership, nil
}

func GetMembershipACL(ctx common.TxCtxInterface, collectionId string) (*authpb.Polices, error) {
	membership, err := GetUserMembership(ctx, collectionId)
	if err != nil {
		return nil, oops.Wrap(err)
	}

	if membership.GetPolices() == nil {
		return nil, oops.Errorf("membership ac is nil")
	}

	return membership.GetPolices(), nil
}
