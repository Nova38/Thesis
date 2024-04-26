package common

import (
	"testing"

	authpb "github.com/nova38/saacs/pkg/saacs-protos/saacs/auth/v0"
	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"
	"github.com/samber/lo"
	"google.golang.org/protobuf/proto"
)

func TestNewItemWithKey(t *testing.T) {
	type args struct {
		itemKey *pb.ItemKey
	}
	type testCase[T ItemInterface] struct {
		name string
		args args
		want T
	}
	tests := []testCase[*authpb.Collection]{
		{
			name: "TestNewItemWithKey",
			args: args{
				itemKey: &pb.ItemKey{
					CollectionId: "cid",
					ItemType:     "auth.Collection",
					ItemKind:     0,
					ItemKeyParts: []string{
						"cid",
					},
				},
			},
			want: &authpb.Collection{
				CollectionId: "cid",
				Name:         "",
				AuthType:     0,
				ItemTypes:    []string{},
				Default:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem := lo.FromPtr(new(authpb.Collection))

			gotItem.SetKey(tt.args.itemKey)
			if !proto.Equal(&gotItem, tt.want) {
				t.Errorf("UnmarshalNew() = %v, want %v", &gotItem, tt.want)
			}

		})
	}
}
