package auth

import (
	"html/template"
	"strings"

	auth_pb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// type Collection pb.Collection
// type PathRolePermission pb.Operations_PathRolePermission
func ValidateCollection(c *auth_pb.Collection) (err error) {
	// Check to see if the type is valid
	if _, err = GetType(c); err != nil {
		return oops.
			With("type", c.GetObjectType()).
			Wrap(err)
	}

	// Check if the paths are valid for the type
	if err = ValidatePaths(c); err != nil {
		return oops.Errorf("invalid paths for collection %v", c.Id.CollectionId)
	}

	// // Check if the roles are valid
	// if err = validateRoles(c); err != nil {
	// 	return oops.Errorf("invalid roles for collection %v", c.Id.CollectionId)
	// }

	return nil
}

//

func GetIdentifierKey(id *auth_pb.Identifier) (key string, err error) {
	(template.New("identifier").Parse("{{.Id}}:{{.Type}}:{{.Value}}"))

	return
}

func splitPath(path string) []string {
	return strings.Split(path, ".")
}

func getSubPaths(
	current *auth_pb.ACL_PathPermission,
) []string {
	paths := []string{current.Path}

	if current.SubPaths == nil || len(current.SubPaths) == 0 {
		return paths
	}
	for path := range current.SubPaths {
		subPaths := getSubPaths(current.SubPaths[path])
		for _, subPath := range subPaths {
			paths = append(paths, path+"."+subPath)
		}
	}

	return paths
}

func GetAllPaths(c *auth_pb.Collection) []string {
	acl := c.GetAcl()

	paths := lo.Uniq(lo.Flatten(lo.MapToSlice(
		acl.GetEntries(),
		func(key string, v *auth_pb.ACL_Entry) []string {
			return getSubPaths(v.ObjectPaths)
		},
	)))

	return paths
}

// Collection Type helpers

func GetType(c *auth_pb.Collection) (protoreflect.MessageType, error) {
	name := protoreflect.FullName(c.ObjectType)

	t, err := protoregistry.GlobalTypes.FindMessageByName(name)
	if err != nil {
		return nil, oops.
			With("type name", name).
			Wrap(err)
	}

	// Check if the type is valid
	if t == nil {
		return nil, oops.
			With("type name", name).
			Wrap(err)
	}

	return t, nil
}

func checkPaths(c *auth_pb.Collection) (err error) {
	t, err := GetType(c)
	if err != nil {
		return err
	}

	defPaths := GetAllPaths(c)

	if len(defPaths) == 0 {
		// No paths defined
		return oops.Errorf("no paths defined")
	}

	mask, err := fieldmaskpb.New(t.New().Interface(), defPaths...)
	if err != nil {
		return oops.
			With("paths", defPaths).
			Wrap(err)
	}
	if !mask.IsValid(t.New().Interface()) {
		return oops.
			With("paths", defPaths).
			Wrap(err)
	}
	// Check if the paths are valid
	//for _, path := range defPaths {
	//	if err := t.Fields().ByJSONName(path); err != nil {
	//		return oops.
	//			With("path", path).
	//			Wrap(err)
	//	}
	//}

	return nil
}

func ValidatePaths(c *auth_pb.Collection) (err error) {
	return checkPaths(c)
}

// WalkACLPath walks through though the path and returns the permission for the path

func ExtractPathPolicy(
	current *auth_pb.ACL_PathPermission,
	path string,
) (*auth_pb.ACL_Policy_ObjectField, error) {
	paths := splitPath(path)

	// check to see if paths is empty
	if len(paths) == 0 {
		return nil, oops.
			In("WalkPATH").
			Code(auth_pb.Error_ERROR_INVALID_OBJECT_FIELD_PATH.String()).
			Errorf("path is empty")
	}

	// Check to see if the first path matches the current path
	if current.Path != paths[0] {
		return nil, oops.
			In("WalkPATH").
			Code(auth_pb.Error_ERROR_INVALID_OBJECT_FIELD_PATH.String()).
			With("path", path, "current path policy", current).
			Errorf("path %s doesn't match current path %s", paths[0], current.Path)
	}

	switch {
	case len(paths) == 1:
		// At the lowest level of the path
		// Return the permission for this path
		fallthrough
	case !current.AllowSubPaths:
		// If the current path doesn't allow sub paths
		// return the current permission
		fallthrough
	case current.SubPaths == nil || len(current.SubPaths) == 0:
		// If the current path doesn't have any sub paths
		// return the current permission
		return current.GetPolicy(), nil

	case len(paths) > 1:
		// Get the next path part
		nextPath := paths[1]
		// Check to see if there is a nested permission for the next path
		nestedPermission, ok := current.SubPaths[nextPath]

		if !ok {
			// No nested permission found for the next path
			// return current permission
			return current.GetPolicy(), nil
		}

		// Recursively call walkPath on the nested permission with the remaining paths
		remainingPaths := paths[1:]

		return ExtractPathPolicy(nestedPermission, strings.Join(remainingPaths, "."))

	default:
		return nil, oops.
			In("WalkPATH").
			Code(auth_pb.Error_ERROR_INVALID_OBJECT_FIELD_PATH.String()).
			With("path", path, "current path policy", current).
			Errorf("invalid path, this should not be thrown")
	}
}

func CheckPathAction(
	path string,
	action auth_pb.Operation_Action,
	policies *auth_pb.ACL_PathPermission,
) (bool, error) {
	policy, err := ExtractPathPolicy(policies, path)
	if err != nil {
		return false, err
	}

	switch action {
	case auth_pb.Operation_ACTION_DELETE:
		return policy.View, nil
	case auth_pb.Operation_ACTION_EDIT:
		return policy.Edit, nil
	case auth_pb.Operation_ACTION_SUGGEST_EDIT:
		return policy.SuggestEdit, nil
	case auth_pb.Operation_ACTION_SUGGEST_APPROVE:
		return policy.SuggestApprove, nil
	case auth_pb.Operation_ACTION_SUGGEST_REJECT:
		return policy.SuggestReject, nil
	default:
		return false, oops.
			In("CheckPathAction").
			Code(auth_pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
			Hint("Action not one of the valid options").
			With("action", action).
			Errorf("invalid action for object")
	}
}

// GetPermission returns the permission for the path
// TODO: reduce complexity
// nolint:cyclop
func _AuthorizeOperation(
	op *auth_pb.Operation,
	role string,
	collection *auth_pb.Collection,
) (authorized bool, err error) {
	// Pointer Safety check
	if op == nil || collection == nil {
		return false, oops.
			In("AuthorizeOperation").
			Code(auth_pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
			Hint("Operation or collection is nil").
			Errorf("invalid operation or collection")
	}

	// Check if the role is valid
	acl, ok := collection.Acl.Entries[role]
	if !ok {
		return false, oops.
			In("AuthorizeOperation").
			Code(auth_pb.Error_ERROR_COLLECTION_INVALID_ROLE_ID.String()).
			Errorf("Role %v is not valid for collection %v", role, collection.Id.CollectionId)
	}

	errorBuilder := oops.
		In("AuthorizeOperation").
		Code(auth_pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
		With("domain", op.Domain.String(), "action", op.Action)

	switch op.Domain {

	case auth_pb.Operation_DOMAIN_UNSPECIFIED:
		return false, errorBuilder.
			Hint("Domain not specified").
			Errorf("invalid domain")

	case auth_pb.Operation_DOMAIN_COLLECTION:
		switch op.Action {
		case auth_pb.Operation_ACTION_CREATE:
			// Everyone can create a collection
			return true, nil
		default:
			return false, errorBuilder.
				Hint("Only Create is a defined option").
				Errorf("invalid action for collection")

		}

		// Collection User Membership
	case auth_pb.Operation_DOMAIN_COLLECTION_MEMBERSHIP:
		if acl.Memberships == nil {
			return false, errorBuilder.
				Hint("Memberships is nil").
				Errorf("ACL is invalid")
		}

		switch op.Action {
		case auth_pb.Operation_ACTION_VIEW:
			// View the current memberships of the collection
			return acl.Memberships.View, nil
		case auth_pb.Operation_ACTION_EDIT:
			// Edit the an users membership in the collection
			return acl.Memberships.Edit, nil
		case auth_pb.Operation_ACTION_DELETE:
			// Delete an users membership in the collection
			return acl.Memberships.Delete, nil
		default:
			return false, errorBuilder.
				Hint("Action not one of the valid options").
				Errorf("invalid action for membership")
		}

	case auth_pb.Operation_DOMAIN_COLLECTION_PERMISSION:
		if acl.RolePermissions == nil {
			return false, errorBuilder.
				Hint("RolePermissions is nil").
				Errorf("ACL is invalid")
		}

		switch op.Action {
		case auth_pb.Operation_ACTION_VIEW:
			// View the current permissions of roles the collection
			return acl.RolePermissions.View, nil
		case auth_pb.Operation_ACTION_EDIT:
			// Edit the permissions of a role in the collection
			return acl.RolePermissions.Edit, nil
		case auth_pb.Operation_ACTION_DELETE:
			// Delete the permissions of a role in the collection
			return acl.RolePermissions.Delete, nil
		default:
			return false, errorBuilder.
				Hint("Action not one of the valid options").
				Errorf("invalid action for permissions")

		}

	case auth_pb.Operation_DOMAIN_COLLECTION_ROLES:
		if acl.RoleDefs == nil {
			return false, errorBuilder.
				Hint("RoleDefs is nil").
				Errorf("ACL is invalid")
		}

		switch op.Action {
		case auth_pb.Operation_ACTION_EDIT:
			// Create a new role in the collection
			return acl.RoleDefs.Create, nil
		case auth_pb.Operation_ACTION_DELETE:
			// Delete a role in the collection
			return acl.RoleDefs.Delete, nil
		default:
			return false, errorBuilder.
				Hint("Action not one of the valid options").
				Errorf("invalid action for roles")

		}

	case auth_pb.Operation_DOMAIN_USER:
		switch op.Action {
		case auth_pb.Operation_ACTION_VIEW:
			// View the current user
			return true, nil
		case auth_pb.Operation_ACTION_CREATE:
			// Create a new user
			return true, nil
		default:
			return false, errorBuilder.
				Hint("Action not one of the valid options").
				Errorf("invalid action for user")

		}

	case auth_pb.Operation_DOMAIN_OBJECT:
		{

			if acl.Object == nil {
				return false, errorBuilder.
					Hint("Object is nil").
					Errorf("ACL is invalid")
			}

			switch op.Action {
			case auth_pb.Operation_ACTION_VIEW:
				return acl.Object.View, nil
			case auth_pb.Operation_ACTION_CREATE:
				return acl.Object.Create, nil
			case auth_pb.Operation_ACTION_DELETE:
				return acl.Object.Delete, nil
			case auth_pb.Operation_ACTION_VIEW_HISTORY:
				return acl.Object.ViewHistory, nil
			case auth_pb.Operation_ACTION_HIDDEN_TX:
				return acl.Object.HiddenTx, nil
			default:
				return false, errorBuilder.
					Hint("Action not one of the valid options").
					Errorf("invalid action for object")
			}
		}

	case auth_pb.Operation_DOMAIN_OBJECT_FIELD:
		{
			if op.Paths == nil {
				return false, errorBuilder.
					Hint("Paths is nil").
					Errorf("Ops is invalid")
			}
			if acl.ObjectPaths == nil {
				return false, errorBuilder.
					Hint("ObjectPaths is nil").
					Errorf("ACL is invalid")
			}

			authorized := true
			defer func() {
				if p := recover(); p != nil {
					authorized = false
					err = errorBuilder.
						Hint("Panic").
						Errorf("invalid path for object, %v", p)
				}
			}()

			for _, path := range op.Paths.GetPaths() {
				path_valid, err := CheckPathAction(path, op.Action, acl.ObjectPaths)
				if err != nil {
					panic(err)
				}
				authorized = authorized && path_valid
			}

			return authorized, nil

		}
	default:
		return false, errorBuilder.
			Hint("Domain not one of the valid options").
			Errorf("invalid domain")

	}
}
