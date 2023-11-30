package state

import (
	"testing"

	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"google.golang.org/protobuf/proto"
)

func TestUnmarshalNewPrimary(t *testing.T) {
	type args struct {
		bytes []byte
		base  *authpb.User
	}
	tests := []struct {
		name     string
		args     args
		wantItem *authpb.User
		wantErr  bool
	}{
		{
			name: "",
			args: args{
				bytes: []byte(
					`{"collection_id":"","msp_id":"msp_id","user_id":"user_id","name":"Name"}`,
				),
				base: &authpb.User{},
			},
			wantItem: &authpb.User{
				CollectionId: "",
				MspId:        "msp_id",
				UserId:       "user_id",
				Name:         "Name",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem, err := UnmarshalNewPrimary[*authpb.User](tt.args.bytes, tt.args.base)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalNewPrimary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !proto.Equal(gotItem, tt.wantItem) {
				t.Errorf("UnmarshalNewPrimary() = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}

func TestUnmarshalPrimary(t *testing.T) {
	type args struct {
		bytes []byte
		obj   *authpb.User
	}
	tests := []struct {
		name     string
		args     args
		wantItem *authpb.User
		wantErr  bool
	}{{
		name: "",
		args: args{
			bytes: []byte(
				`{"collection_id":"","msp_id":"msp_id","user_id":"user_id","name":"Name"}`,
			),
		},
		wantItem: &authpb.User{
			CollectionId: "",
			MspId:        "msp_id",
			UserId:       "user_id",
			Name:         "Name",
		},
		wantErr: false,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem := new(authpb.User)

			err := UnmarshalPrimary(tt.args.bytes, gotItem)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalNewPrimary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !proto.Equal(gotItem, tt.wantItem) {
				t.Errorf("UnmarshalNewPrimary() = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}
