package policy

import (
	"slices"

	"github.com/nova38/saacs/pkg/saacs-cc/common"
	authpb "github.com/nova38/saacs/pkg/saacs-protos/saacs/auth/v0"
	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"

	"github.com/samber/oops"
)

// ValidateOperation validates an operation against a collection,
// returns true if valid, false otherwise.
// Checks:
// - Collection is not nil
// - Operation is not nil
// - Collection item types is not nil and not empty
// - Checks per the type of action (see below)
func ValidateOperation(collection *authpb.Collection, op *pb.Operation) (bool, error) {

	// Sanity checks
	if collection == nil || op == nil {
		return false, oops.Errorf("Collection or Operation is nil")
	}

	// Checks for nil and empty cases
	switch {
	case collection == nil:
		return false, oops.Errorf("Collection is nil")
	case op == nil:
		return false, oops.Errorf("Operations is nil")

	case collection.GetItemTypes() == nil:
		return false, oops.Errorf("Collection item types is nil")
	case collection.GetItemTypes() != nil && len(collection.GetItemTypes()) == 0:
		return false, oops.Errorf("Collection item types is empty")

	case op.GetCollectionId() == "":
		return false, oops.With("op", op).Errorf("Operation collection id is empty")
	case op.GetItemType() == "":
		return false, oops.Errorf("Operation item type is empty")
	case op.GetAction() == pb.Action_ACTION_UNSPECIFIED:
		return false, oops.Errorf("Operation action is empty")
	}

	if op.GetItemType() == common.CollectionItemType {
		return true, nil
	}

	// Check to see if the operation item type is in the collection item types

	if !slices.Contains(collection.GetItemTypes(), op.GetItemType()) {
		return false, oops.Errorf(
			"Operation item type %v is not in collection %v item types",
			op.GetItemType(),
			op.GetCollectionId(),
		)
	}

	// Checks for validity against collection according to action type
	switch op.GetAction() {

	case
		pb.Action_ACTION_CREATE,
		pb.Action_ACTION_DELETE,
		pb.Action_ACTION_HIDE_TX,
		pb.Action_ACTION_VIEW_HISTORY,
		pb.Action_ACTION_VIEW_HIDDEN_TXS:

		// These actions should not have paths
		if op.GetPaths() != nil && len(op.GetPaths().GetPaths()) > 0 {
			return false, oops.Errorf("Operation paths is not empty")
		}

	}

	return true, nil
}

// --------------------------------------------------

func ActionOnPathPolicy(p *authpb.PathPolicy, action pb.Action) (authorized bool, found bool) {
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
