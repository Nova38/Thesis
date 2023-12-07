package common

import (
	"testing"

	v1 "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	"github.com/samber/lo"
	"google.golang.org/protobuf/proto"
)

func TestNewItemWithKey(t *testing.T) {
	type args struct {
		itemKey *v1.ItemKey
	}
	type testCase[T ItemInterface] struct {
		name string
		args args
		want T
	}
	tests := []testCase[*v1.Collection]{
		{
			name: "TestNewItemWithKey",
			args: args{
				itemKey: &v1.ItemKey{
					CollectionId: "cid",
					ItemType:     "auth.Collection",
					ItemKind:     0,
					ItemKeyParts: []string{
						"cid",
					},
				},
			},
			want: &v1.Collection{
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
			gotItem := lo.FromPtr(new(v1.Collection))

			gotItem.SetKey(tt.args.itemKey)
			if !proto.Equal(&gotItem, tt.want) {
				t.Errorf("UnmarshalNew() = %v, want %v", &gotItem, tt.want)
			}

		})
	}
}
