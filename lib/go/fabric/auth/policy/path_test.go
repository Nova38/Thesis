package policy

import (
	"testing"

	auth_pb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	_ "github.com/samber/lo"
)

var policies = []*auth_pb.ACL_PathPermission{
	{
		Path:          "a",
		AllowSubPaths: false,
		SubPaths:      map[string]*auth_pb.ACL_PathPermission{"b": {Path: "b"}},
		Policy:        &auth_pb.ACL_Policy_ObjectField{},
	},
}

func TestPathWalker_checkLeaf(t *testing.T) {
	type args struct {
		action auth_pb.Operation_Action
	}
	tests := []struct {
		name        string
		w           *PathWalker
		args        args
		wantAllowed bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAllowed, err := tt.w.checkLeaf(tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathWalker.checkLeaf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAllowed != tt.wantAllowed {
				t.Errorf("PathWalker.checkLeaf() = %v, want %v", gotAllowed, tt.wantAllowed)
			}
		})
	}
}
