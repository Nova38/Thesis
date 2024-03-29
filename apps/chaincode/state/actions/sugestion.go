package actions

import (
	"github.com/nova38/saacs/apps/chaincode/common"
	"github.com/nova38/saacs/apps/chaincode/state"
	authpb "github.com/nova38/saacs/libs/saacs-protos-go/auth/v1"
	"github.com/samber/oops"
)

// ════════════════════════════════════════════════════════
// Suggestion Functions
// ════════════════════════════════════════════════════════

// ──────────────────────────────────────────────────
// Invoke Suggested Functions
// ──────────────────────────────────────────────────

func SuggestionCreate(ctx common.TxCtxInterface, s *authpb.Suggestion) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	ctx.GetLogger().Debug("CreateSuggestion", "suggestion:", s)

	if ctx.Validate(s) != nil {
		return ctx.ErrorBase().Wrap(err)
	}

	// Authorize the operation
	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_SUGGEST_CREATE,
		CollectionId: s.GetPrimaryKey().GetCollectionId(),
		ItemType:     s.GetPrimaryKey().GetItemType(),
		Paths:        s.GetPaths(),
	}
	if auth, err := ctx.Authorize([]*authpb.Operation{op}); err != nil {
		return oops.Wrap(err)
	} else if !auth {
		return oops.Wrap(common.UserPermissionDenied)
	}

	// Extract the item from the suggestion

	if exists := state.KeyExists(ctx, common.MakeStateKey(s.GetPrimaryKey())); !exists {
		return oops.
			With(
				"suggestion ID", s.GetSuggestionId(),
				"Primary Key", s.GetPrimaryKey(),
			).
			Wrap(common.KeyNotFound)
	}

	if exists := state.Exists(ctx, s); exists {
		return oops.
			With(
				"suggestion ID", s.GetSuggestionId(),
				"Primary Key", s.GetPrimaryKey(),
			).
			Wrap(common.AlreadyExists)
	}

	return state.Put(ctx, s)

}

func SuggestionDelete(ctx common.TxCtxInterface, s *authpb.Suggestion) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	ctx.GetLogger().Debug("DeleteSuggestion", "suggestion:", s)

	if ctx.Validate(s) != nil {
		return ctx.ErrorBase().Wrap(err)
	}

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_SUGGEST_DELETE,
		CollectionId: s.GetPrimaryKey().GetCollectionId(),
		ItemType:     s.GetPrimaryKey().GetItemType(),
		Paths:        s.GetPaths(),
	}

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); err != nil {
		return oops.Wrap(err)
	} else if !auth {
		return oops.Wrap(common.UserPermissionDenied)
	}

	return state.Delete(ctx, s)
}

func SuggestionApprove(
	ctx common.TxCtxInterface,
	s *authpb.Suggestion,
) (updated common.ItemInterface, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	ctx.GetLogger().Debug("ApproveSuggestion", "suggestion:", s)

	// Authorize Viewing and Retrieving the suggestion
	if ctx.Validate(s) != nil {
		return nil, ctx.ErrorBase().Wrap(err)
	}

	if err = GetSuggestion(ctx, s); err != nil {
		return nil, oops.Wrap(err)
	}

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_SUGGEST_APPROVE,
		CollectionId: s.GetPrimaryKey().GetCollectionId(),
		ItemType:     s.GetPrimaryKey().GetItemType(),
		Paths:        s.GetPaths(),
	}

	// Authorize the operation
	if auth, err := ctx.Authorize([]*authpb.Operation{op}); err != nil {
		return nil, oops.
			With("Operation", op).
			Wrap(err)
	} else if !auth {
		return nil, oops.
			With("Operation", op).
			Wrap(common.UserPermissionDenied)
	}

	// Process the Update
	src, err := common.ItemKeyToItem(s.GetPrimaryKey())
	if err != nil {
		return nil, oops.Wrap(err)
	} else if err = PrimaryGet(ctx, src); err != nil {
		return nil, oops.Hint("Shouldn't be possible").Wrap(common.KeyNotFound)
	}

	v, err := s.GetValue().UnmarshalNew()
	if err != nil {
		return nil, oops.Wrap(err)
	}

	update, ok := v.(common.ItemInterface)
	if !ok {
		return nil, oops.Wrap(common.ItemInvalid)
	}

	// Update the primary object
	updated, err = common.UpdateItem(s.GetPaths(), src, update)

	// update the primary object
	if err = state.Put(ctx, updated); err != nil {
		return nil, oops.
			With("Primary", src, "Suggestion", s).
			Wrap(err)
	}

	// Delete the suggestion
	if err = state.Delete(ctx, s); err != nil {
		return nil, oops.
			With("Suggestion", s).
			Wrap(err)
	}

	return updated, nil
}

// ──────────────────────────────────────────────────
// Query Suggested Functions
// ──────────────────────────────────────────────────

// GetSuggestion by its ID
func GetSuggestion(ctx common.TxCtxInterface, s *authpb.Suggestion) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	ctx.GetLogger().Debug("GetSuggestion", "suggestion:", s)

	// Authorize the operation
	auth, err := ctx.Authorize([]*authpb.Operation{
		{
			Action:       authpb.Action_ACTION_SUGGEST_VIEW,
			CollectionId: s.GetPrimaryKey().GetCollectionId(),
			ItemType:     s.GetPrimaryKey().GetItemType(),
			Paths:        nil,
		},
	})
	if err != nil {
		return oops.Wrap(err)
	} else if !auth {
		return oops.Wrap(common.UserPermissionDenied)
	}

	if err := state.GetFromKey(ctx, s.StateKey(), s); err != nil {
		return oops.
			With("Suggestion", s).
			Wrap(err)
	}

	return nil
}

// PartialSuggestionList returns a list of suggestions based on the partial key
func PartialSuggestionList(
	ctx common.TxCtxInterface,
	s *authpb.Suggestion,
	numAttr int,
	bookmark string,
) (list []*authpb.Suggestion, mk string, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	ctx.GetLogger().Debug("GetPartialSuggestionList", "suggestion:", s)

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_SUGGEST_VIEW,
		CollectionId: s.GetPrimaryKey().GetCollectionId(),
		ItemType:     s.GetPrimaryKey().GetItemType(),
		Paths:        nil,
	}
	// Authorize the operation
	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return nil, "", oops.Wrap(common.UserPermissionDenied)
	}

	return state.GetPartialKeyList(ctx, s, numAttr, bookmark)
}

// Get Full Suggestion List, for the collection
func SuggestionListByCollection(ctx common.TxCtxInterface, collectionId string, bookmark string) (
	list []*authpb.Suggestion, mk string, err error,
) {
	ctx.GetLogger().Debug("GetSuggestionListByCollection", "collection_id:", collectionId)

	s := &authpb.Suggestion{
		PrimaryKey: &authpb.ItemKey{
			CollectionId: collectionId,
		},
	}

	return PartialSuggestionList(ctx, s, 1, bookmark)
}

// Get Suggestion List, for the item
func SuggestionListByItem(
	ctx common.TxCtxInterface,
	objKey *authpb.ItemKey,
	bookmark string,
) (list []*authpb.Suggestion, mk string, err error) {
	ctx.GetLogger().Debug("GetSuggestionList", "suggestion:", objKey)

	s := &authpb.Suggestion{
		PrimaryKey: objKey,
	}

	num := len(objKey.GetItemKeyParts()) + 2

	return PartialSuggestionList(ctx, s, num, bookmark)
}
