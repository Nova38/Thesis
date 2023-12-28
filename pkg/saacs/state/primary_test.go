package state

import (
	"encoding/json"
	"testing"

	authpb "github.com/nova38/thesis/pkg/saacs/gen/auth/v1"
	"google.golang.org/protobuf/proto"
)

func TestUnmarshalPrimary(t *testing.T) {
	type args struct {
		bytes []byte
		// obj   *authpb.UserDirectMembership
	}
	tests := []struct {
		name     string
		args     args
		wantItem *authpb.UserDirectMembership
		wantErr  bool
	}{{
		name: "",
		args: args{
			bytes: []byte(
				`{"collection_id":"","msp_id":"msp_id","user_id":"user_id","name":"Name"}`,
			),
		},
		wantItem: &authpb.UserDirectMembership{
			CollectionId: "",
			MspId:        "msp_id",
			UserId:       "user_id",
		},
		wantErr: false,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem := new(authpb.UserDirectMembership)

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
