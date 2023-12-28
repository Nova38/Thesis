package policy

import (
	"log/slog"
	"slices"
	"strings"

	authpb "github.com/nova38/saacs/pkg/chaincode/gen/auth/v1"
	"github.com/samber/lo"
)

func AuthorizedPolicy(policy *authpb.Polices, op *authpb.Operation) (bool, error) {

	// Check to see if there is a policy that matches the operation object type
	// We assume that the more specific policy would be more permisive
	if policy.GetItemPolicies() != nil {
		if itemPolicy, ok := policy.GetItemPolicies()[op.GetItemType()]; ok {
			if AuthorizePathPolicy(itemPolicy, op) {
				slog.Debug("Operation is allowed on specific policy")
				return true, nil
			}
		}

	}

	// Fallback to the default policy if their is one.
	//  Check the default policy to see if their is a general perm that satisfies the request
	if policy.GetDefaultPolicy() != nil {
		if policy.GetDefaultExcludedTypes() != nil {
			if lo.Contains(policy.GetDefaultExcludedTypes(), op.GetItemType()) {
				slog.Debug("Operation is not allowed on default policy")
				return false, nil
			}

		}
		if AuthorizePathPolicy(policy.GetDefaultPolicy(), op) {
			slog.Debug("Operation is allowed on default policy")
			return true, nil
		}
	}

	slog.Default().
		Debug("Operation is not authorized",
			slog.Group("operation",
				slog.Any("Operation", op), "policy",
				slog.Any("Policy", policy)))
	return false, nil
}

func AuthorizePathPolicy(pathPolicy *authpb.PathPolicy, op *authpb.Operation) bool {

	if pathPolicy.GetActions() != nil {

		// If the action is in the list of allowed actions,
		// we don't need to search any further
		if slices.Contains(pathPolicy.GetActions(), op.GetAction()) {
			return true
		}
	}

	{ // SubPath sanity check
		// If the operation has no paths and the policy doesn't let them
		// on the root path, then they are not authorized
		if op.GetPaths() == nil || len(op.GetPaths().GetPaths()) == 0 {
			slog.Default().Debug("Operation has no paths",
				slog.Group("operation", slog.Any("Operation", op)),
			)
			return false
		}

		if !pathPolicy.GetAllowSubPaths() {
			slog.Default().Debug("PathPolicy does not allow sub paths",
				slog.Group("operation", "action", op.GetAction().String(), "paths", op.GetPaths().GetPaths()),
			)
			return false
		}

		if pathPolicy.GetSubPaths() == nil || len(pathPolicy.GetSubPaths()) == 0 {
			slog.Default().Debug("PathPolicy has no sub paths",
				slog.Group("operation", "action", op.GetAction().String(), "paths", op.GetPaths().GetPaths()),
			)
			return false
		}
	}

	return walkSubPaths(pathPolicy, op.GetAction(), op.GetPaths().GetPaths())

}

func walkSubPaths(pathPolicy *authpb.PathPolicy, action authpb.Action, paths []string) bool {

	subPaths := make(map[string][]string)

	for _, path := range paths {
		split := strings.Split(path, ".")
		if len(split) == 1 {
			subPaths[split[0]] = []string{""}
		} else if len(split) > 1 {
			subPaths[split[0]] = split[1:]
		}
	}

	for path, children := range subPaths {
		if subPathPolicy, ok := pathPolicy.GetSubPaths()[path]; ok {
			if !authorizeSubPaths(subPathPolicy, action, children) {
				slog.Default().Debug("Operation is not authorized",
					slog.Group("operation", "action", action.String(), "path", path, "children", children),
				)
				return false
			}
		}

		slog.Default().Debug("SubPathPolicy doesn't have a matching sub path")
		return false
	}

	return true
}

func authorizeSubPaths(
	subPathPolicy *authpb.PathPolicy,
	action authpb.Action,
	paths []string,
) bool {

	if subPathPolicy.GetActions() != nil {
		if slices.Contains(subPathPolicy.GetActions(), action) {
			return true
		}
		slog.Default().Debug("SubPathPolicy does not allow action, checking sub paths")

	}

	if !subPathPolicy.GetAllowSubPaths() {
		slog.Default().Debug("SubPathPolicy does not allow sub paths")
		return false
	}
	if subPathPolicy.GetSubPaths() == nil || len(subPathPolicy.GetSubPaths()) == 0 {
		slog.Default().Debug("SubPathPolicy has no sub paths")
		return false
	}

	return walkSubPaths(subPathPolicy, action, paths)

}
