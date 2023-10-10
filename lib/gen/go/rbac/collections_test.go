package rbac

import (
	"reflect"
	"testing"
)

func TestOperations_PathRolePermission_walkPath(t *testing.T) {
	type fields struct {
		Path     string
		SubPaths map[string]*Operations_PathRolePermission
		Acl      map[int32]*Operations_ObjectField
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[int32]*Operations_ObjectField
		wantErr bool
	}{
		{
			name: "Nested Path Exists",
			fields: fields{
				Path: "path1",
				SubPaths: map[string]*Operations_PathRolePermission{
					"path2": {
						Path: "path2",
						Acl: map[int32]*Operations_ObjectField{
							1: {},
						},
					},
				},
			},
			args: args{
				path: "path1.path2",
			},
			want: map[int32]*Operations_ObjectField{
				1: {},
			},
			wantErr: false,
		},
		{
			name: "path does not exist",
			fields: fields{
				Path: "path1",
				SubPaths: map[string]*Operations_PathRolePermission{
					"path2": {
						Path: "path2",
						Acl: map[int32]*Operations_ObjectField{
							1: {},
						},
					},
				},
			},
			args: args{
				path: "path3",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "path is empty",
			fields: fields{
				Path: "path1",
				SubPaths: map[string]*Operations_PathRolePermission{
					"path2": {
						Path: "path2",
						Acl: map[int32]*Operations_ObjectField{
							1: {},
						},
					},
				},
			},
			args: args{
				path: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			current := &Operations_PathRolePermission{
				Path:     tt.fields.Path,
				SubPaths: tt.fields.SubPaths,
				Acl:      tt.fields.Acl,
			}
			got, err := current.walkPath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("walkPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("walkPath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
