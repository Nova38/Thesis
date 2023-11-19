package policy

// import (
// 	auth_pb "github.com/nova38/thesis/lib/go/gen/auth/v1"
// 	"github.com/samber/oops"
// )

// type (
// 	ACLHandler struct {
// 		op      *auth_pb.Operation
// 		entries []*auth_pb.State
// 	}
// )

// func BuildACLHandler(
// 	col *auth_pb.Collection,
// 	op *auth_pb.Operation,
// 	id *auth_pb.Identifier,
// ) (handler *ACLHandler, err error) {
// 	// Insure that all pointers are not nil
// 	if col == nil || op == nil || id == nil ||
// 		col.GetSettings() != nil ||
// 		col.GetAcl() == nil {
// 		return nil, oops.Errorf("nil pointer")
// 	}

// 	handler = &ACLHandler{
// 		op:      op,
// 		entries: []*auth_pb.ACL_Entry{},
// 	}

// 	// Get the Identifier key
// 	idKey, err := GetIdentifierKey(id)
// 	if err != nil {
// 		return nil, oops.Wrapf(err, "failed to get identifier key")
// 	}

// 	// Retrieve the Identifier information from the collection
// 	idEntry, ok := col.Identifiers[idKey]
// 	if !ok {
// 		return nil, oops.Errorf("identifier not found in collection")
// 	}

// 	// Retrieve the identifiers ACL entry
// 	baseEntry, ok := col.Acl.Entries[idKey]
// 	if !ok {
// 		return nil, oops.Errorf("identifier acl entry not found in collection")
// 	}

// 	// Add the identifier entry to the handler
// 	handler.entries = append(handler.entries, baseEntry)

// 	// See if the collection has inheritance enabled
// 	if col.Settings.AclInheritanceEnabled {
// 		for _, parent := range idEntry.Parents {
// 			parentEntry, ok := col.Acl.Entries[parent]
// 			if !ok {
// 				return nil, oops.With("").
// 					Wrapf(err, "failed to get identifier key")
// 			}
// 			handler.entries = append(handler.entries, parentEntry)
// 		}
// 	}

// 	// add the default entry
// 	if col.Acl.Default != nil {
// 		handler.entries = append(handler.entries, col.Acl.Default)
// 	}

// 	// add the public entry
// 	if col.Acl.Public != nil {
// 		handler.entries = append(handler.entries, col.Acl.Public)
// 	}

// 	return handler, nil
// }

// func (h *ACLHandler) valid() (err error) {
// 	valid := h.op != nil && h.entries != nil && len(h.entries) > 0
// 	if !valid {
// 		return h.error().Errorf("invalid handler")
// 	}

// 	return nil
// }

// func (h *ACLHandler) error() (builder oops.OopsErrorBuilder) {
// 	return oops.
// 		In("ACLHandler").
// 		With("operation", h.op, "entries", h.entries)
// }

// func (h *ACLHandler) IsAllowed() (allowed bool, err error) {
// 	if err = h.valid(); err != nil {
// 		return false, err
// 	}

// 	switch h.op.Domain {
// 	case auth_pb.Operation_DOMAIN_UNSPECIFIED:
// 		return false, h.error().Errorf("invalid domain")
// 	case auth_pb.Operation_DOMAIN_COLLECTION:
// 		return h.collection()
// 	case auth_pb.Operation_DOMAIN_COLLECTION_MEMBERSHIP:
// 		return h.collectionMembership()
// 	case auth_pb.Operation_DOMAIN_COLLECTION_PERMISSION:
// 		return h.collectionPermission()
// 	case auth_pb.Operation_DOMAIN_COLLECTION_ROLES:
// 		return h.collectionRoles()
// 	case auth_pb.Operation_DOMAIN_USER:
// 		return h.user()
// 	case auth_pb.Operation_DOMAIN_OBJECT:
// 		return h.object()
// 	case auth_pb.Operation_DOMAIN_OBJECT_FIELD:
// 		return h.objectField()
// 	default:
// 		return false, h.error().Errorf("invalid domain")

// 	}
// }

// // ════════════════════════════════════════════════==
// // Domain Handlers
// // ════════════════════════════════════════════════==

// // collection - Handles the operations in the
// // collection domain
// func (h *ACLHandler) collection() (allowed bool, err error) {
// 	if err = h.valid(); err != nil {
// 		return false, err
// 	}
// 	// Everyone can create a collection
// 	if h.op.Action == auth_pb.Operation_ACTION_CREATE {
// 		return true, nil
// 	}
// 	return false, nil
// }

// // user - Handles the operations in the
// // user domain
// func (h *ACLHandler) user() (allowed bool, err error) {
// 	if err = h.valid(); err != nil {
// 		return false, err
// 	}

// 	// Everyone can create a user
// 	switch h.op.Action {
// 	case auth_pb.Operation_ACTION_CREATE:
// 		return true, nil
// 	case auth_pb.Operation_ACTION_VIEW:
// 		// View the user
// 		return true, nil
// 	default:
// 		return false, h.error().Errorf("invalid action for user")

// 	}
// }

// // collectionMembership - Handles the operations in the
// // collection membership domain
// func (h *ACLHandler) collectionMembership() (allowed bool, err error) {
// 	if err = h.valid(); err != nil {
// 		return false, err
// 	}

// 	var policy *auth_pb.ACL_Policy_Permission

// 	// Extract the policy from the entries
// 	for _, entry := range h.entries {
// 		if entry.GetMemberships() == nil {
// 			continue
// 		}
// 		policy = entry.GetMemberships()
// 		break
// 	}
// 	if policy == nil {
// 		return false, h.error().Errorf("no policy found")
// 	}

// 	switch h.op.Action {
// 	case auth_pb.Operation_ACTION_VIEW:
// 		// View the current memberships of the collection
// 		return policy.View, nil
// 	case auth_pb.Operation_ACTION_EDIT:
// 		// Edit the an users membership in the collection
// 		return policy.Edit, nil
// 	case auth_pb.Operation_ACTION_DELETE:
// 		// Delete an users membership in the collection
// 		return policy.Delete, nil
// 	default:
// 		return false, h.error().
// 			Hint("Action not one of the valid options").
// 			Errorf("invalid action for membership")
// 	}
// }

// // collectionPermission - Handles the operations in the
// // collection permission domain
// func (h *ACLHandler) collectionPermission() (allowed bool, err error) {
// 	if err = h.valid(); err != nil {
// 		return false, err
// 	}

// 	var policy *auth_pb.ACL_Policy_Permission

// 	// Extract the policy from the entries
// 	for _, entry := range h.entries {
// 		if entry.GetRolePermissions() == nil {
// 			continue
// 		}
// 		policy = entry.GetRolePermissions()
// 		break
// 	}
// 	if policy == nil {
// 		return false, h.error().Errorf("no policy found")
// 	}

// 	switch h.op.Action {
// 	case auth_pb.Operation_ACTION_VIEW:
// 		// View the permissions of the collection
// 		return policy.View, nil
// 	case auth_pb.Operation_ACTION_EDIT:
// 		// Edit the permissions of the collection
// 		return policy.Edit, nil
// 	case auth_pb.Operation_ACTION_DELETE:
// 		// Delete the permissions of the collection
// 		return policy.Delete, nil
// 	default:
// 		return false, h.error().
// 			Errorf("invalid action for permission")
// 	}
// }

// // collectionRoles - Handles the operations in the
// // collection roles domain
// func (h *ACLHandler) collectionRoles() (allowed bool, err error) {
// 	if err = h.valid(); err != nil {
// 		return false, err
// 	}

// 	var policy *auth_pb.ACL_Policy_Permission

// 	for _, entry := range h.entries {
// 		if entry.GetRoleDefs() == nil {
// 			continue
// 		}
// 		policy = entry.GetRoleDefs()
// 		break
// 	}

// 	if policy == nil {
// 		return false, h.error().Errorf("no policy found")
// 	}

// 	switch h.op.Action {
// 	case auth_pb.Operation_ACTION_EDIT:
// 		// Create a new role in the collection
// 		return policy.Create, nil
// 	case auth_pb.Operation_ACTION_DELETE:
// 		// Delete a role in the collection
// 		return policy.Delete, nil
// 	default:
// 		return false, h.error().
// 			Errorf("invalid action for roles")
// 	}
// }

// // object - Handles the operations in the
// // object domain
// func (h *ACLHandler) object() (allowed bool, err error) {
// 	if err = h.valid(); err != nil {
// 		return false, err
// 	}

// 	var policy *auth_pb.ACL_Policy_Object

// 	for _, entry := range h.entries {
// 		if entry.GetObject() == nil {
// 			continue
// 		}
// 		policy = entry.GetObject()
// 		break
// 	}

// 	if policy == nil {
// 		return false, h.error().Errorf("no policy found")
// 	}

// 	switch h.op.Action {
// 	case auth_pb.Operation_ACTION_VIEW:
// 		return policy.View, nil
// 	case auth_pb.Operation_ACTION_CREATE:
// 		return policy.Create, nil
// 	case auth_pb.Operation_ACTION_DELETE:
// 		return policy.Delete, nil
// 	case auth_pb.Operation_ACTION_VIEW_HISTORY:
// 		return policy.ViewHistory, nil
// 	case auth_pb.Operation_ACTION_HIDDEN_TX:
// 		return policy.HiddenTx, nil
// 	default:
// 		return false, h.error().
// 			Errorf("invalid action for object")
// 	}
// }

// func (h *ACLHandler) objectField() (allowed bool, err error) {
// 	// TODO: implement
// 	baseWalker := &PathWalker{
// 		entries:        []*auth_pb.ACL_PathPermission{},
// 		restrict_paths: []string{},
// 	}
// 	for _, entry := range h.entries {
// 		if entry.ObjectPaths == nil {
// 			continue
// 		}
// 		baseWalker.entries = append(baseWalker.entries, entry.ObjectPaths)
// 	}
// 	if len(baseWalker.entries) == 0 {
// 		return false, h.error().Errorf("no policy found")
// 	}

// 	return allowed, nil
// }

// // ════════════════════════════════════════════════==
// // Object Field Helpers
// // ════════════════════════════════════════════════==
