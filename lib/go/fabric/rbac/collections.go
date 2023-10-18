package rbac

import (
	"fmt"
	"strings"

	pb "github.com/nova38/thesis/lib/gen/go/rbac"
)

// type Collection pb.Collection
// type PathRolePermission pb.Operations_PathRolePermission

//

func splitPath(path string) []string {
	return strings.Split(path, ".")
}

// walkACLPath walks though though the path and returns the permission for the path
func WalkACLPath(
	current *pb.Operations_PathRolePermission,
	path string,
) (map[int32]*pb.Operations_ObjectField, error) {
	paths := splitPath(path)

	// check to see if paths is empty
	if len(paths) == 0 {
		return nil, fmt.Errorf("empty path")
	}

	// Check to see if the first path matches the current path
	if current.Path != paths[0] {
		return nil, fmt.Errorf("path %s doesn't match current path %s", paths[0], current.Path)
	}

	// Check to see if paths only contains one element
	if len(paths) == 1 {
		// Return the permission for this path
		return current.Acl, nil
	}

	// Get the next path part
	nextPath := paths[1]
	// Check to see if there is a nested permission for the next path
	nestedPermission, ok := current.SubPaths[nextPath]
	if !ok {
		// No nested permission found for the next path
		// return current permission
		return current.Acl, nil
	}

	// Recursively call walkPath on the nested permission with the remaining paths
	remainingPaths := paths[1:]

	return WalkACLPath(nestedPermission, strings.Join(remainingPaths, "."))
}

// GetPermission returns the permission for the path
