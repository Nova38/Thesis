package policy_test

import (
	"testing"

	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	"github.com/nova38/thesis/packages/saacs/policy"
	"github.com/stretchr/testify/assert"
)

func TestValidateOperation(t *testing.T) {
	collection := &authpb.Collection{
		ItemTypes: []string{"item1", "item2"},
	}

	t.Run("Valid operation", func(t *testing.T) {
		op := &authpb.Operation{
			CollectionId: "collection1",
			ItemType:     "item1",
			Action:       authpb.Action_ACTION_CREATE,
		}

		valid, err := policy.ValidateOperation(collection, op)
		assert.True(t, valid)
		assert.NoError(t, err)
	})

	t.Run("Invalid collection", func(t *testing.T) {
		op := &authpb.Operation{
			CollectionId: "collection1",
			ItemType:     "item3",
			Action:       authpb.Action_ACTION_CREATE,
		}

		valid, err := policy.ValidateOperation(nil, op)
		assert.False(t, valid)
		assert.EqualError(t, err, "Collection or Operation is nil")
	})

	t.Run("Invalid operation", func(t *testing.T) {
		op := &authpb.Operation{
			CollectionId: "collection1",
			ItemType:     "",
			Action:       authpb.Action_ACTION_UNSPECIFIED,
		}

		valid, err := policy.ValidateOperation(collection, op)
		assert.False(t, valid)
		assert.EqualError(t, err, "Operation item type is empty")
	})

	// Add more test cases for different scenarios
}
