package rbac

import (
	"strings"

	pb "github.com/nova38/thesis/lib/go/gen/rbac"
	"github.com/samber/oops"
)

// type Collection pb.Collection
// type PathRolePermission pb.Operations_PathRolePermission

//

func splitPath(path string) []string {
	return strings.Split(path, ".")
}

// WalkACLPath walks through though the path and returns the permission for the path

func ExtractPathPolicy(
	current *pb.ACL_PathRolePermission,
	path string,
) (*pb.ACL_Policy_ObjectField, error) {
	paths := splitPath(path)

	// check to see if paths is empty
	if len(paths) == 0 {
		return nil, oops.
			In("WalkPATH").
			Code(pb.Error_ERROR_INVALID_OBJECT_FIELD_PATH.String()).
			Errorf("path is empty")
	}

	// Check to see if the first path matches the current path
	if current.Path != paths[0] {
		return nil, oops.
			In("WalkPATH").
			Code(pb.Error_ERROR_INVALID_OBJECT_FIELD_PATH.String()).
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
			Code(pb.Error_ERROR_INVALID_OBJECT_FIELD_PATH.String()).
			With("path", path, "current path policy", current).
			Errorf("invalid path, this should not be thrown")
	}
}

func CheckPathAction(
	path string,
	action pb.ACL_Action,
	policies *pb.ACL_PathRolePermission,
) (bool, error) {
	policy, err := ExtractPathPolicy(policies, path)
	if err != nil {
		return false, err
	}

	switch action {
	case pb.ACL_ACTION_VIEW:
		return policy.View, nil
	case pb.ACL_ACTION_EDIT:
		return policy.Edit, nil
	case pb.ACL_ACTION_SUGGEST_EDIT:
		return policy.SuggestEdit, nil
	case pb.ACL_ACTION_SUGGEST_APPROVE:
		return policy.SuggestApprove, nil
	case pb.ACL_ACTION_SUGGEST_REJECT:
		return policy.SuggestReject, nil
	default:
		return false, oops.
			In("CheckPathAction").
			Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
			Hint("Action not one of the valid options").
			With("action", action).
			Errorf("invalid action for object")
	}
}

// GetPermission returns the permission for the path
// TODO: reduce complexity
// nolint:cyclop
func AuthorizeOperation(
	op *pb.ACL_Operation,
	role int32,
	collection *pb.Collection,
) (authorized bool, err error) {
	// Pointer Safety check
	if op == nil || collection == nil {
		return false, oops.
			In("AuthorizeOperation").
			Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
			Hint("Operation or collection is nil").
			Errorf("invalid operation or collection")
	}

	// Check if the role is valid
	acl, ok := collection.Acl[role]
	if !ok {
		return false, oops.
			In("AuthorizeOperation").
			Code(pb.Error_ERROR_COLLECTION_INVALID_ROLE_ID.String()).
			Errorf("Role %v is not valid for collection %v", role, collection.Id.CollectionId)
	}

	switch op.Domain {

	case pb.ACL_DOMAIN_UNSPECIFIED:
		return false, oops.
			In("AuthorizeOperation").
			Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
			Hint("Domain not specified").
			Errorf("invalid domain")

	case pb.ACL_DOMAIN_COLLECTION:
		switch op.Action {
		case pb.ACL_ACTION_CREATE:
			// Everyone can create a collection
			return true, nil
		default:
			return false, oops.
				In("AuthorizeOperation").
				Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
				Hint("Only Create is defined").
				With("domain", op.Domain.String(), "action", op.Action).
				Errorf("invalid action for collection")

		}

		// Collection User Membership
	case pb.ACL_DOMAIN_COLLECTION_MEMBERSHIP:
		if acl.Memberships == nil {
			return false, oops.
				In("AuthorizeOperation").
				Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
				Hint("Memberships is nil").
				With("domain", op.Domain.String(), "action", op.Action).
				Errorf("ACL is invalid")
		}

		switch op.Action {
		case pb.ACL_ACTION_VIEW:
			// View the current memberships of the collection
			return acl.Memberships.View, nil
		case pb.ACL_ACTION_EDIT:
			// Edit the an users membership in the collection
			return acl.Memberships.Edit, nil
		case pb.ACL_ACTION_DELETE:
			// Delete an users membership in the collection
			return acl.Memberships.Delete, nil
		default:
			return false, oops.
				In("AuthorizeOperation").
				Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
				Hint("Action not one of the valid options").
				With("domain", op.Domain.String(), "action", op.Action).
				Errorf("invalid action for membership")
		}

	case pb.ACL_DOMAIN_COLLECTION_PERMISSION:
		if acl.RolePermissions == nil {
			return false, oops.
				In("AuthorizeOperation").
				Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
				Hint("RolePermissions is nil").
				With("domain", op.Domain.String(), "action", op.Action).
				Errorf("ACL is invalid")
		}

		switch op.Action {
		case pb.ACL_ACTION_VIEW:
			// View the current permissions of roles the collection
			return acl.RolePermissions.View, nil
		case pb.ACL_ACTION_EDIT:
			// Edit the permissions of a role in the collection
			return acl.RolePermissions.Edit, nil
		case pb.ACL_ACTION_DELETE:
			// Delete the permissions of a role in the collection
			return acl.RolePermissions.Delete, nil
		default:
			return false, oops.
				In("AuthorizeOperation").
				Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
				Hint("Action not one of the valid options").
				With("domain", op.Domain.String(), "action", op.Action).
				Errorf("invalid action for permissions")

		}

	case pb.ACL_DOMAIN_COLLECTION_ROLES:
		if acl.RoleDefs == nil {
			return false, oops.
				In("AuthorizeOperation").
				Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
				Hint("RoleDefs is nil").
				With("domain", op.Domain.String(), "action", op.Action).
				Errorf("ACL is invalid")
		}

		switch op.Action {
		case pb.ACL_ACTION_CREATE:
			// Create a new role in the collection
			return acl.RoleDefs.Create, nil
		case pb.ACL_ACTION_DELETE:
			// Delete a role in the collection
			return acl.RoleDefs.Delete, nil
		default:
			return false, oops.
				In("AuthorizeOperation").
				Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
				Hint("Action not one of the valid options").
				With("domain", op.Domain.String(), "action", op.Action).
				Errorf("invalid action for roles")

		}

	case pb.ACL_DOMAIN_USER:
		switch op.Action {
		case pb.ACL_ACTION_VIEW:
			// View the current user
			return true, nil
		case pb.ACL_ACTION_CREATE:
			// Create a new user
			return true, nil
		default:
			return false, oops.
				In("AuthorizeOperation").
				Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
				Hint("Action not one of the valid options").
				With("domain", op.Domain.String(), "action", op.Action).
				Errorf("invalid action for user")

		}

	case pb.ACL_DOMAIN_OBJECT:
		{

			if acl.Object == nil {
				return false, oops.
					In("AuthorizeOperation").
					Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
					Hint("Object is nil").
					With("domain", op.Domain.String(), "action", op.Action).
					Errorf("ACL is invalid")
			}

			switch op.Action {
			case pb.ACL_ACTION_VIEW:
				return acl.Object.View, nil
			case pb.ACL_ACTION_CREATE:
				return acl.Object.Create, nil
			case pb.ACL_ACTION_DELETE:
				return acl.Object.Delete, nil
			case pb.ACL_ACTION_VIEW_HISTORY:
				return acl.Object.ViewHistory, nil
			case pb.ACL_ACTION_HIDDEN_TX:
				return acl.Object.HiddenTx, nil
			default:
				return false, oops.
					In("AuthorizeOperation").
					Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
					Hint("Action not one of the valid options").
					With("domain", op.Domain.String(), "action", op.Action).
					Errorf("invalid action for object")
			}
		}

	case pb.ACL_DOMAIN_OBJECT_FIELD:
		{
			if op.Paths == nil {
				return false, oops.
					In("AuthorizeOperation").
					Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
					Hint("Paths is nil").
					With("domain", op.Domain.String(), "action", op.Action).
					Errorf("Ops is invalid")
			}
			if acl.ObjectPaths == nil {
				return false, oops.
					In("AuthorizeOperation").
					Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
					Hint("ObjectPaths is nil").
					With("domain", op.Domain.String(), "action", op.Action).
					Errorf("ACL is invalid")
			}

			authorized := true
			defer func() {
				if p := recover(); p != nil {
					authorized = false
					err = oops.
						In("AuthorizeOperation").
						Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
						Hint("Panic").
						With("domain", op.Domain.String(), "action", op.Action).
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
		return false, oops.
			In("AuthorizeOperation").
			Code(pb.Error_ERROR_RUNTIME_BAD_OPS.String()).
			Hint("Domain not one of the valid options").
			With("domain", op.Domain).
			Errorf("invalid domain")

	}
}
