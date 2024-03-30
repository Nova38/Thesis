package validate_test

import (
	"testing"

	"github.com/nova38/saacs/apps/saacs-cc/common"

	authpb "github.com/nova38/saacs/libs/saacs-protos-go/auth/v1"
	sample "github.com/nova38/saacs/libs/saacs-protos-go/sample/v0"
)

func TestMakePrimaryKeyAttr(t *testing.T) {
	type args struct {
		obj common.ItemInterface
	}
	tests := []struct {
		name    string
		args    args
		wantKey string
	}{
		{
			name: "Simple Item",
			args: args{
				obj: &sample.SimpleItem{
					CollectionId: "SimpleColId",
					Id:           "SimpleId",
					Name:         "",
					Quantity:     0,
				},
			},
			wantKey: "\u0000sample.SimpleItem\u0000SimpleColId\u0000SimpleId\u0000",
		},
		{
			name: "Role Item",
			args: args{
				obj: &authpb.Role{
					CollectionId: "ColId",
					RoleId:       "RoleId",
				},
			},
			wantKey: "\u0000auth.Role\u0000ColId\u0000RoleId\u0000",
		},
		{
			name: "Suggestion Item",
			args: args{
				obj: &authpb.Suggestion{
					PrimaryKey: &authpb.ItemKey{
						CollectionId: "ColId",
						ItemType:     "auth.Role",
						ItemKind:     authpb.ItemKind_ITEM_KIND_PRIMARY_ITEM,
						ItemKeyParts: []string{"RoleId"},
					},
					SuggestionId: "SuggestionId",
				},
			},
			wantKey: "\u0000auth.Suggestion\u0000ColId\u0000auth.Role\u0000RoleId\u0000SuggestionId\u0000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey := tt.args.obj.StateKey()
			if gotKey != tt.wantKey {
				t.Errorf("MakePrimaryKey() = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}
