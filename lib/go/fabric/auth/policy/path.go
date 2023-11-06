package policy

import (
	"log/slog"
	"os"
	"strings"

	auth_pb "github.com/nova38/thesis/lib/go/gen/auth/v1"

	_ "github.com/samber/lo"
	"github.com/samber/oops"
)

type PathWalker struct {
	entry *auth_pb.ACL_PathPermission

	entries        []*auth_pb.ACL_PathPermission
	restrict_paths []string
}

func (w *PathWalker) error() (err oops.OopsErrorBuilder) {
	return oops.
		In("ACLHandler").
		With("walker", w, "entries", w.entries, "restrict_paths", w.restrict_paths)
}

func splitPath(path string) []string {
	return strings.Split(path, ".")
}

func groupSubPaths(paths []string) (subPaths map[string][]string) {
	subPaths = make(map[string][]string)

	for _, path := range paths {
		// Split the path
		split := splitPath(path)
		base := split[0]
		sub := strings.Join(split[1:], ".")
		subPaths[base] = append(subPaths[base], sub)
	}

	return subPaths
}

func (w *PathWalker) groupEntrySubPaths() (subFields map[string][]*auth_pb.ACL_PathPermission) {
	subFields = make(map[string][]*auth_pb.ACL_PathPermission)
	for _, subPath := range w.entry.SubPaths {
		subFields[subPath.Path] = append(subFields[subPath.Path], subPath)
	}
	return subFields
}

func checkPolicy(
	policy *auth_pb.ACL_Policy_ObjectField,
	action auth_pb.Operation_Action,
) (allowed bool) {
	// Check the action
	switch action {
	case auth_pb.Operation_ACTION_EDIT:
		return policy.Edit
	case auth_pb.Operation_ACTION_VIEW:
		return policy.View
	case auth_pb.Operation_ACTION_SUGGEST_EDIT:
		return policy.SuggestEdit
	case auth_pb.Operation_ACTION_SUGGEST_APPROVE:
		return policy.SuggestApprove
	case auth_pb.Operation_ACTION_SUGGEST_REJECT:
		return policy.SuggestReject
	default:
		return false
	}
}

// // Get the index of the first entry that has a policy defined. It will be the base entry
// // entries afterwords will not be considered. Entries before it will have their
// // subPaths considered. If they are defined
// func (w *PathWalker) baseIndex() (index int, err error) {
// 	for i, v := range w.entries {
// 		if policy := v.GetPolicy(); policy != nil {
// 			return i, nil
// 		}
// 	}
// 	return -1, w.error().Errorf("no policy found")
// }

// checkLeaf - Checks the leaf node (no subPaths) for the given action.

// Should be called on a fresh walker
func (w *PathWalker) checkLeaf(action auth_pb.Operation_Action) (allowed bool, err error) {
	// Find the first entry that as a defined policy
	for _, v := range w.entries {
		if policy := v.GetPolicy(); policy != nil {
			return checkPolicy(policy, action), nil
		}
	}
	return false, w.error().Errorf("no policy found")
}

// checkNode - Checks the node for the given action. If the node has subPaths
// then it will check them as well by either calling checkLeaf or checkNode
// recursively. Should be called on a fresh walker
func (w *PathWalker) checkNode(
	path string,
	subPaths []string,
	action auth_pb.Operation_Action,
) (allowed bool, err error) {
	// Find the first entry that as a defined policy

	if len(subPaths) == 0 {
		return w.checkLeaf(action)
	}

	// Find the first entry that as a defined policy
	for _, v := range w.entries {
		// Extract the sub entries. Only ones that have the path in the subPaths
		// will be considered
		if entry, ok := v.SubPaths[path]; ok {
			// If a policy is defined here then we can stop
			if entry.GetPolicy() != nil {
				w.entry = entry
				break
			}
		}
	}

	if w.entry == nil {
		return false, w.error().Errorf("no entry found for path %s", path)
	}

	// If the entry does not allow sub paths then we just check the policy
	if !w.entry.AllowSubPaths {
		return checkPolicy(w.entry.GetPolicy(), action), nil
	}

	// If the entry allows sub paths then we need to check them
	// We need to group the sub paths by their base path
	pathMap := groupSubPaths(subPaths)

	for path, subPaths := range pathMap {
		subWalker := &PathWalker{
			entry:          &auth_pb.ACL_PathPermission{},
			entries:        []*auth_pb.ACL_PathPermission{},
			restrict_paths: []string{},
		}

		// build the subWalker by adding entries that have the path in their subPaths
		for _, entry := range w.entry.SubPaths {
			// subWalker.entries = append(subWalker.entries, entry)
		}

	}

	return false, w.error().Errorf("no policy found")
}

// Check - Checks the ACL for the given paths and action
func (w *PathWalker) Check(
	paths []string,
	action auth_pb.Operation_Action,
) (allowed bool, err error) {
	if len(paths) == 0 {
		w.checkLeaf(action)
	}

	// Path groups
	grouped := groupSubPaths(paths)

	// TODO: apply

	for path, subPaths := range grouped {
		w.checkNode(path, subPaths, action)
	}

	return false, nil
}

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)
}
