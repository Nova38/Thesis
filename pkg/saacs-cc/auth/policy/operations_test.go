package policy_test

import (
	"testing"

	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"

	"github.com/nova38/saacs/pkg/saacs-cc/auth/policy"
	authpb "github.com/nova38/saacs/pkg/saacs-protos/saacs/auth/v0"
	"github.com/stretchr/testify/assert"
)

func TestValidateOperation(t *testing.T) {
	collection := &authpb.Collection{
		ItemTypes: []string{"item1", "item2"},
	}

	t.Run("Valid operation", func(t *testing.T) {
		op := &pb.Operation{
			CollectionId: "collection1",
			ItemType:     "item1",
			Action:       pb.Action_ACTION_CREATE,
		}

		valid, err := policy.ValidateOperation(collection, op)
		assert.True(t, valid)
		assert.NoError(t, err)
	})

	t.Run("Invalid collection", func(t *testing.T) {
		op := &pb.Operation{
			CollectionId: "collection1",
			ItemType:     "item3",
			Action:       pb.Action_ACTION_CREATE,
		}

		valid, err := policy.ValidateOperation(nil, op)
		assert.False(t, valid)
		assert.EqualError(t, err, "Collection or Operation is nil")
	})

	t.Run("Invalid operation", func(t *testing.T) {
		op := &pb.Operation{
			CollectionId: "collection1",
			ItemType:     "",
			Action:       pb.Action_ACTION_UNSPECIFIED,
		}

		valid, err := policy.ValidateOperation(collection, op)
		assert.False(t, valid)
		assert.EqualError(t, err, "Operation item type is empty")
	})

	// Add more test cases for different scenarios
}
