package state

import (
	"encoding/json"
	"testing"

	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	"google.golang.org/protobuf/proto"
)

func TestUnmarshalNewPrimary(t *testing.T) {
	type args struct {
		bytes []byte
		base  *authpb.UserMembership
	}
	tests := []struct {
		name     string
		args     args
		wantItem *authpb.UserMembership
		wantErr  bool
	}{
		{
			name: "",
			args: args{
				bytes: []byte(
					`{"collection_id":"","msp_id":"msp_id","user_id":"user_id","name":"Name"}`,
				),
				base: &authpb.UserMembership{},
			},
			wantItem: &authpb.UserMembership{
				CollectionId: "",
				MspId:        "msp_id",
				UserId:       "user_id",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem, err := UnmarshalNewPrimary[*authpb.UserMembership](tt.args.bytes, tt.args.base)
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
		// obj   *authpb.UserMembership
	}
	tests := []struct {
		name     string
		args     args
		wantItem *authpb.UserMembership
		wantErr  bool
	}{{
		name: "",
		args: args{
			bytes: []byte(
				`{"collection_id":"","msp_id":"msp_id","user_id":"user_id","name":"Name"}`,
			),
		},
		wantItem: &authpb.UserMembership{
			CollectionId: "",
			MspId:        "msp_id",
			UserId:       "user_id",
		},
		wantErr: false,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem := new(authpb.UserMembership)

			err := json.Unmarshal(tt.args.bytes, gotItem)
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
