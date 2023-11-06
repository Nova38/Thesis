package policy

import (
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

func (w *PathWalker) Check(
	paths []string,
	action auth_pb.Operation_Action,
) (allowed bool, err error) {
	if len(paths) == 0 {
		// Find the first entry that as a defined policy
		for _, entry := range w.entries {
			if entry == nil {
				continue
			}

			if entry.GetPolicy() != nil {
				w.entry = entry
				break
			}
		}
		if w.entry != nil {
			w.entry.GetPolicy()
		}
		return false, w.error().Errorf("no policy found")

	}
}
