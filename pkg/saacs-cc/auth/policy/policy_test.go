package policy_test

import (
	"testing"

	"github.com/nova38/saacs/pkg/saacs-cc/auth/policy"
	authpb "github.com/nova38/saacs/pkg/saacs-protos/saacs/auth/v0"
	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestAuthorizedPolicyLight(t *testing.T) {

	policies := &authpb.Polices{
		ItemPolicies: map[string]*authpb.PathPolicy{"item1": {
			Path:          "",
			FullPath:      "",
			AllowSubPaths: false,
			SubPaths:      map[string]*authpb.PathPolicy{},
			Actions: []pb.Action{
				pb.Action_ACTION_CREATE,
			},
		}},
		DefaultPolicy: &authpb.PathPolicy{
			Path:          "",
			FullPath:      "",
			AllowSubPaths: false,
			SubPaths:      map[string]*authpb.PathPolicy{},
			Actions: []pb.Action{
				pb.Action_ACTION_VIEW,
			},
		},
	}

	t.Run("Operation allowed on specific policy", func(t *testing.T) {
		op := &pb.Operation{
			CollectionId: "collection1",
			ItemType:     "item1",
			Action:       pb.Action_ACTION_CREATE,
		}

		allowed, err := policy.AuthorizedPolicy(policies, op)
		assert.True(t, allowed)
		assert.NoError(t, err)
	})

	t.Run("Operation allowed on default policy", func(t *testing.T) {
		op := &pb.Operation{
			CollectionId: "collection2",
			ItemType:     "item2",
			Action:       pb.Action_ACTION_VIEW,
		}

		allowed, err := policy.AuthorizedPolicy(policies, op)
		assert.True(t, allowed)
		assert.NoError(t, err)
	})

	t.Run("Operation not authorized", func(t *testing.T) {
		op := &pb.Operation{
			CollectionId: "collection3",
			ItemType:     "item3",
			Action:       pb.Action_ACTION_UPDATE,
		}

		allowed, err := policy.AuthorizedPolicy(policies, op)
		assert.False(t, allowed)
		assert.NoError(t, err)
	})
}

func TestAuthorizePathPolicy(t *testing.T) {

	// authPolicies := &authpb.Polices{
	// 	ItemPolicies: map[string]*authpb.PathPolicy{
	// 		common.UserGlobalRoles: {
	// 			Path:          "",
	// 			FullPath:      "",
	// 			AllowSubPaths: false,
	// 			SubPaths:      map[string]*authpb.PathPolicy{},
	// 			Actions: []pb.Action{
	// 				pb.Action_ACTION_CREATE,
	// 			},
	// 		},
	// 	},

	// 	DefaultPolicy: &authpb.PathPolicy{
	// 		Path:          "",
	// 		FullPath:      "",
	// 		AllowSubPaths: false,
	// 		SubPaths:      map[string]*authpb.PathPolicy{},
	// 		Actions: []pb.Action{
	// 			pb.Action_ACTION_VIEW,
	// 			pb.Action_ACTION_SUGGEST_VIEW,
	// 			pb.Action_ACTION_VIEW_HISTORY,
	// 		},
	// 	},
	// }

	type args struct {
		op         *pb.Operation
		pathPolicy *authpb.PathPolicy
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Operation allowed on specific policy",
			args: args{
				op: &pb.Operation{
					Action:       pb.Action_ACTION_VIEW,
					CollectionId: "",
					ItemType:     "",
					Paths:        &fieldmaskpb.FieldMask{},
				},
				pathPolicy: &authpb.PathPolicy{
					Path:          "",
					FullPath:      "",
					AllowSubPaths: false,
					SubPaths:      map[string]*authpb.PathPolicy{},
					Actions: []pb.Action{
						pb.Action_ACTION_VIEW,
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := policy.AuthorizePathPolicy(tt.args.pathPolicy, tt.args.op); got != tt.want {
				t.Errorf("AuthorizePathPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}
