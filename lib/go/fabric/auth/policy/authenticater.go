package policy

import (
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
)

func BuildPolicyChecker(policies []*authpb.PathPolicy) {
}

func ActionOnPathPolicy(p *authpb.PathPolicy, action authpb.Action) (authorized bool, found bool) {
	if p == nil && p.GetActions() == nil {
		return false, false
	}

	for _, a := range p.GetActions() {
		if a == action {
			return true, true
		}
	}

	return false, false
}
