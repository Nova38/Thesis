package policy_test

import (
	"testing"

	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	"github.com/nova38/thesis/packages/saacs/policy"
	"github.com/stretchr/testify/assert"
)

func TestAuthorizedPolicy(t *testing.T) {

	policies := &authpb.Polices{
		ItemPolicies: map[string]*authpb.PathPolicy{"item1": {
			Path:          "",
			FullPath:      "",
			AllowSubPaths: false,
			SubPaths:      map[string]*authpb.PathPolicy{},
			Actions: []authpb.Action{
				authpb.Action_ACTION_CREATE,
			},
		}},
		DefaultPolicy: &authpb.PathPolicy{
			Path:          "",
			FullPath:      "",
			AllowSubPaths: false,
			SubPaths:      map[string]*authpb.PathPolicy{},
			Actions: []authpb.Action{
				authpb.Action_ACTION_VIEW,
			},
		},
	}

	t.Run("Operation allowed on specific policy", func(t *testing.T) {
		op := &authpb.Operation{
			CollectionId: "collection1",
			ItemType:     "item1",
			Action:       authpb.Action_ACTION_CREATE,
		}

		allowed, err := policy.AuthorizedPolicy(policies, op)
		assert.True(t, allowed)
		assert.NoError(t, err)
	})

	t.Run("Operation allowed on default policy", func(t *testing.T) {
		op := &authpb.Operation{
			CollectionId: "collection2",
			ItemType:     "item2",
			Action:       authpb.Action_ACTION_VIEW,
		}

		allowed, err := policy.AuthorizedPolicy(policies, op)
		assert.True(t, allowed)
		assert.NoError(t, err)
	})

	t.Run("Operation not authorized", func(t *testing.T) {
		op := &authpb.Operation{
			CollectionId: "collection3",
			ItemType:     "item3",
			Action:       authpb.Action_ACTION_UPDATE,
		}

		allowed, err := policy.AuthorizedPolicy(policies, op)
		assert.False(t, allowed)
		assert.NoError(t, err)
	})
}
