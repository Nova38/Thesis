package rbac

import (
	"fmt"
	"strings"
)

// type Collection pb.Collection
// type PathRolePermission pb.Operations_PathRolePermission

//

func splitPath(path string) []string {
	return strings.Split(path, ".")
}

// Search PathRolePermission for a given paths and return the permission for the given role

// GetPathToRoleObjectField returns the permission for the given path and role
// func (c *Collection) t(path string, role int) (*pb.Operations_ObjectField, error) {

// 	// Get the next path part

// 	// Check to see if

// 	return nil, fmt.Errorf("no permission found for path: %v, role: %v, action: %v", paths, role)
// }

func (current *Operations_PathRolePermission) walkPath(
	path string,
) (map[int32]*Operations_ObjectField, error) {
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

	return nestedPermission.walkPath(strings.Join(remainingPaths, "."))

}
