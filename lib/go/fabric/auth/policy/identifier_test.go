package policy

import (
	"testing"

	auth_pb "github.com/nova38/thesis/lib/go/gen/auth/v1"
)

func TestGetACLKey(t *testing.T) {
	testCases := []struct {
		desc     string
		input    *auth_pb.Identifier
		expected string
		wantErr  bool
	}{
		{
			desc: "nil Identifier.Id",
			input: &auth_pb.Identifier{
				Id: nil,
			},
			wantErr: true,
		},
		{
			desc: "Attribute",
			input: &auth_pb.Identifier{
				Id: &auth_pb.Identifier_Attribute_{
					Attribute: &auth_pb.Identifier_Attribute{
						Name:  "name",
						Value: "value",
						MspId: "mspId",
					},
				},
			},
			expected: "Attribute[<Name:name><Value:value><MspId:mspId>]",
		},
		{
			desc: "Role",
			input: &auth_pb.Identifier{
				Id: &auth_pb.Identifier_Role_{
					Role: &auth_pb.Identifier_Role{
						Id: "roleId",
					},
				},
			},
			expected: "Role[<Id:roleId>]",
		},
		{
			desc: "Identity",
			input: &auth_pb.Identifier{
				Id: &auth_pb.Identifier_Identity_{
					Identity: &auth_pb.Identifier_Identity{
						MspId: "mspId",
						Id:    "id",
					},
				},
			},
			expected: "Identity[<MspId:mspId><Id:id>]",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual, err := GetACLKey(tC.input)
			if (err != nil) != tC.wantErr {
				t.Fatalf("GetACLKey() error = %v, wantErr %v", err, tC.wantErr)
			}
			if actual != tC.expected {
				t.Errorf("GetACLKey() = %v, want %v", actual, tC.expected)
			}
		})
	}
}
