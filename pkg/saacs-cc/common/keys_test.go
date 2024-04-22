package common

import (
	"reflect"
	"testing"

	authpb "github.com/nova38/saacs/pkg/saacs-protos/auth/v1"
)

// func TestKeyExists(t *testing.T) {
// 	type args struct {
// 		ctx TxCtxInterface
// 		key string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := KeyExists(tt.args.ctx, tt.args.key); got != tt.want {
// 				t.Errorf("KeyExists() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestMakeItemKeyAttr(t *testing.T) {
	type args struct {
		key *authpb.ItemKey
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1 part",
			args: args{
				key: &authpb.ItemKey{
					CollectionId: "collection",
					ItemKeyParts: []string{"id"},
				},
			},
			want: []string{
				"id",
			},
		},
		{
			name: "2 parts",
			args: args{
				key: &authpb.ItemKey{
					CollectionId: "collection",
					ItemKeyParts: []string{"id1", "id2"},
				},
			},
			want: []string{
				"id1",
				"id2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.key.GetItemKeyParts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeItemKeyAttr() = %v, want %v", got, tt.want)
			}
		})
	}
}
