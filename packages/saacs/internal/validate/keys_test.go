package validate_test

import (
	"testing"

	"github.com/nova38/thesis/packages/saacs/common"

	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	sample "github.com/nova38/thesis/packages/saacs/gen/sample/v0"
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
