package policy

import (
	"log/slog"
	"slices"

	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/oops"
)

type OperationChecker struct {
	op         *authpb.Operation
	collection *authpb.Collection
	Acls       map[string]authpb.Polices
}

// ValidateOperation validates an operation against a collection,
// returns true if valid, false otherwise.
// Checks:
// - Collection is not nil
// - Operation is not nil
// - Collection item types is not nil and not empty
// - Checks per the type of action (see below)
func ValidateOperation(collection *authpb.Collection, op *authpb.Operation) (bool, error) {
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
		return false, oops.Errorf("Operation collection id is empty")
	case op.GetItemType() == "":
		return false, oops.Errorf("Operation item type is empty")
	case op.GetAction() == authpb.Action_ACTION_UNSPECIFIED:
		return false, oops.Errorf("Operation action is empty")

	}

	// Checks for validity against collection according to action type
	switch op.GetAction() {
	// Reference Actions must have both item type and secondary item type in collection reference types, and paths must be empty
	case
		authpb.Action_ACTION_REFERENCE_CREATE,
		authpb.Action_ACTION_REFERENCE_DELETE,
		authpb.Action_ACTION_REFERENCE_VIEW:

		if op.GetSecondaryItemType() == "" {
			return false, oops.Errorf("Operation secondary item type is empty")
		}

		// found := slices.Contains(collection.GetReferenceTypes(), op.GetItemType())
		// foundSecondary := slices.Contains(collection.GetReferenceTypes(), op.GetSecondaryItemType())

		// if !found || !foundSecondary {
		// 	return false, oops.Errorf(
		// 		"Operation item type or secondary item type is not in collection reference types",
		// 	)
		// }

		if op.GetPaths() != nil && len(op.GetPaths().GetPaths()) > 0 {
			return false, oops.Errorf("Operation paths is not empty")
		}

	case
		authpb.Action_ACTION_CREATE,
		authpb.Action_ACTION_DELETE,
		authpb.Action_ACTION_HIDE_TX,
		authpb.Action_ACTION_VIEW_HISTORY,
		authpb.Action_ACTION_VIEW_HIDDEN_TXS:

		found := slices.Contains(collection.GetItemTypes(), op.GetItemType())
		if !found {
			return false, oops.Errorf(
				"Operation item type %v is not in collection %v item types",
				op.GetItemType(),
				op.GetCollectionId(),
			)
		}

		if op.GetPaths() != nil && len(op.GetPaths().GetPaths()) > 0 {
			return false, oops.Errorf("Operation paths is not empty")
		}

		// Actions who might have paths
	case
		authpb.Action_ACTION_UPDATE,
		authpb.Action_ACTION_VIEW,
		authpb.Action_ACTION_SUGGEST_CREATE,
		authpb.Action_ACTION_SUGGEST_DELETE,
		authpb.Action_ACTION_SUGGEST_APPROVE:

		found := slices.Contains(collection.GetItemTypes(), op.GetItemType())
		if !found {
			return false, oops.Errorf(
				"Operation item type %v is not in collection %v 's item types",
				op.GetItemType(),
				op.GetCollectionId(),
			)
		}
	}

	return true, nil
}

// HandleOperation handles an operation
// Gets the collection of the operation and validates the operation against it
// Then sees if the default ACL of the collection allows the operation
// If so, returns true, otherwise returns false
func HandleCollectionOperation(
	ctx common.TxCtxInterface,
	op *authpb.Operation,
) (auth bool, err error) {
	ctx.GetLogger().Debug("HandleOperation", slog.Group("args", "op", op))

	// Get Operation Collection
	collection := &authpb.Collection{CollectionId: op.GetCollectionId()}

	if key, err := common.MakePrimaryKey(collection); err != nil {
		return false, oops.Wrap(err)
	} else if err = state.Get(ctx, key, collection); err != nil {
		oops.Wrap(err)
	}

	// Validate Operation is an operation of the collection
	if _, err = ValidateOperation(collection, op); err != nil {
		return false, oops.Wrap(err)
	}

	if collection.GetAuthType() == authpb.AuthType_AUTH_TYPE_NONE {
		ctx.GetLogger().Info("Auth type is no auth, short circuiting")
		return true, nil
	}

	// Get the default ACL of the collection
	acl := collection.GetDefault()
	if acl == nil {
		return false, oops.Wrap(err)
	}

	// Check if the user is authorized to perform the given action on the given reference
	if auth, err := AuthOp(ctx, op, acl); err != nil {
	} else if !auth {
		return false, oops.Wrap(common.UserPermissionDenied)
	} else if auth {
		return true, nil
	}

	policies := []*authpb.Polices{}

	// Handle the extraction of the policies of diffrent auth types
	switch collection.GetAuthType() {
	case authpb.AuthType_AUTH_TYPE_UNSPECIFIED:
		ctx.GetLogger().Error("Auth type Unspecified")
		return false, oops.Wrap(common.UserPermissionDenied)

	case authpb.AuthType_AUTH_TYPE_IDENTITY:
		policy, err := GetMembershipACL(ctx, collection.GetCollectionId())
		if err != nil {
			return false, oops.Wrap(err)
		}

		policies = append(policies, policy)
	case authpb.AuthType_AUTH_TYPE_ROLE:

	}

	return false, nil
}

func Authenticate(ctx common.TxCtxInterface, ops []*authpb.Operation) (bool, error) {
	return true, nil
}

func AuthOp(
	ctx common.TxCtxInterface,
	op *authpb.Operation,
	acl *authpb.Polices,
) (auth bool, err error) {
	// TODO implement AuthOp
	return false, nil
}

// --------------------------------------------------

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
