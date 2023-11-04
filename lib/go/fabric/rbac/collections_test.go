package rbac

import (
	"reflect"
	"testing"

	pb "github.com/nova38/thesis/lib/go/gen/rbac"
	"google.golang.org/protobuf/proto"
)

// Test Collections
var collection1 = &pb.Collection{
	Id: &pb.Collection_Id{
		CollectionId: "Collection1",
	},
	ObjectNamespace: "some_namespace",
	ObjectType:      "",
	Roles: map[string]string{
		"0": "public",
		"1": "admin",
		"2": "user",
		"3": "guest",
	},
	Acl: map[string]*pb.ACL{
		"0": {
			RoleDefs: &pb.ACL_Policy_Roles{
				Create: false,
				Delete: false,
			},
			RolePermissions: &pb.ACL_Policy_Roles{
				View:   false,
				Edit:   false,
				Delete: false,
			},
			Memberships: &pb.ACL_Policy_Roles{
				View:   false,
				Edit:   false,
				Delete: false,
			},
			Object: &pb.ACL_Policy_Object{
				View:        false,
				Create:      false,
				Delete:      false,
				ViewHistory: false,
				HiddenTx:    false,
			},
			ObjectPaths: &pb.ACL_PathRolePermission{
				Path:          "",
				AllowSubPaths: false,
				SubPaths:      map[string]*pb.ACL_PathRolePermission{},
				Policy: &pb.ACL_Policy_ObjectField{
					View:           false,
					Edit:           false,
					SuggestEdit:    false,
					SuggestApprove: false,
					SuggestReject:  false,
				},
			},
		},
		"1": {},
		"2": {},
		"3": {},
	},
}

func Test_splitPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitPath(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractPathPolicy(t *testing.T) {
	type args struct {
		current *pb.ACL_PathRolePermission
		path    string
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.ACL_Policy_ObjectField
		wantErr bool
	}{
		{
			name: "Empty path",
			args: args{
				current: &pb.ACL_PathRolePermission{
					Policy: &pb.ACL_Policy_ObjectField{
						View:           false,
						Edit:           false,
						SuggestEdit:    false,
						SuggestApprove: false,
						SuggestReject:  false,
					},
				},
				path: "",
			},
			want: &pb.ACL_Policy_ObjectField{
				View:           false,
				Edit:           false,
				SuggestEdit:    false,
				SuggestApprove: false,
				SuggestReject:  false,
			},
			wantErr: false,
		},
		{
			name: "Root path match",
			args: args{
				current: &pb.ACL_PathRolePermission{
					Path: "root",
				},
				path: "root",
			},
			want: &pb.ACL_Policy_ObjectField{
				// expected permission
			},
			wantErr: false,
		},
		{
			name: "Root path mismatch",
			args: args{
				current: &pb.ACL_PathRolePermission{
					Path: "root1",
				},
				path: "root2",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Nested path match",
			args: args{
				current: &pb.ACL_PathRolePermission{
					Path: "root",
					SubPaths: map[string]*pb.ACL_PathRolePermission{
						"child": {
							Path: "child",
							// child policy
						},
					},
				},
				path: "root.child",
			},
			want: &pb.ACL_Policy_ObjectField{
				// expected child policy
			},
			wantErr: false,
		},
		{
			name: "Nested path mismatch",
			args: args{
				current: &pb.ACL_PathRolePermission{
					Path: "root",
				},
				path: "root.child",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractPathPolicy(tt.args.current, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractPathPolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if proto.Equal(got, tt.want) {
				t.Errorf("ExtractPathPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPathAction(t *testing.T) {
	type args struct {
		path     string
		action   pb.ACL_Action
		policies *pb.ACL_PathRolePermission
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckPathAction(tt.args.path, tt.args.action, tt.args.policies)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPathAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckPathAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthorizeOperation(t *testing.T) {
	type args struct {
		op         *pb.ACL_Operation
		role       string
		collection *pb.Collection
	}
	tests := []struct {
		name           string
		args           args
		wantAuthorized bool
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				op:         &pb.ACL_Operation{},
				role:       "0",
				collection: collection1,
			},
			wantAuthorized: false,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAuthorized, err := AuthorizeOperation(tt.args.op, tt.args.role, tt.args.collection)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthorizeOperation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAuthorized != tt.wantAuthorized {
				t.Errorf("AuthorizeOperation() = %v, want %v", gotAuthorized, tt.wantAuthorized)
			}
		})
	}
}
