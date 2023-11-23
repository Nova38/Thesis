package state

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
)

// ════════════════════════════════════════════════════════
// Suggestion Functions
// ════════════════════════════════════════════════════════

// ──────────────────────────────────────────────────
// Helpers
// ──────────────────────────────────────────────────
type SuggestionHandler struct {
	suggestionKey string
	objKey        string
	obj           common.ItemInterface
	current       common.ItemInterface
	suggestion    *authpb.Suggestion
	bytes         []byte
}

// SuggestionToItem converts a suggestion to an item

func (s SuggestionHandler) Extract(sug *authpb.Suggestion) (err error) {
	if sug == nil {
		return oops.In("Extract").Errorf("Suggestion is nil")
	}
	s.suggestion = sug

	if s.obj, err = SuggestionToItem(s.suggestion); err != nil {
		return oops.Wrap(err)
	}

	if s.objKey, err = MakePrimaryKey(s.obj); err != nil {
		return oops.Wrap(err)
	}

	if s.suggestionKey, err = MakeSuggestionKey(s.obj, s.suggestion.GetSuggestionId()); err != nil {
		return oops.Wrap(err)
	}

	if s.bytes, err = json.Marshal(s.obj); err != nil {
		return oops.Wrap(err)
	}

	return nil
}

// ──────────────────────────────────────────────────
// Invoke Suggested Functions
// ──────────────────────────────────────────────────

func SuggestionCreate(ctx TxCtxInterface, s *authpb.Suggestion) (err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	ctx.GetLogger().Debug("CreateSuggestion", "suggestion:", s)

	var (
		l       = &Ledger[*authpb.Suggestion]{ctx: ctx}
		handler = SuggestionHandler{suggestion: s}
		op      = &authpb.Operation{
			Action:       authpb.Action_ACTION_SUGGEST_CREATE,
			CollectionId: s.GetPrimaryKey().GetCollectionId(),
			ItemType:     s.GetPrimaryKey().GetItemType(),
			Paths:        s.GetPaths(),
		}
	)

	// Authorize the operation

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	// Extract the item from the suggestion
	if err = handler.Extract(s); err != nil {
		return oops.Wrap(err)
	}

	// Check if the item exists, if it does not, return error
	if !KeyExists(ctx, handler.objKey) {
		return oops.
			With("suggestion ID",
				s.GetSuggestionId(), "Key", handler.objKey,
				"ItemType", handler.obj.ItemType()).
			Wrap(common.KeyNotFound)
	}

	return l.Create(handler.suggestion)
}

func SuggestionDelete(ctx TxCtxInterface, s *authpb.Suggestion) (err error) {
	ctx.GetLogger().Debug("DeleteSuggestion", "suggestion:", s)

	var (
		l       = &Ledger[*authpb.Suggestion]{ctx: ctx}
		handler = SuggestionHandler{suggestion: s}
		op      = &authpb.Operation{
			Action:       authpb.Action_ACTION_SUGGEST_DELETE,
			CollectionId: s.GetPrimaryKey().GetCollectionId(),
			ItemType:     s.GetPrimaryKey().GetItemType(),
			Paths:        nil,
		}
	)
	// Extract the item from the suggestion
	if err = handler.Extract(s); err != nil {
		return oops.Wrap(err)
	}

	if err = l.Get(handler.suggestion); err != nil {
		return oops.
			In("DeleteSuggestion").
			With("Key", s, "ItemType", s.ItemType()).
			Wrap(common.KeyNotFound)
	}

	// Set The paths
	op.Paths = handler.suggestion.GetPaths()

	// Authorize the operation
	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	return l.Delete(handler.suggestion)
}

func SuggestionApprove(
	ctx TxCtxInterface,
	s *authpb.Suggestion,
) (updated *common.ItemInterface, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	ctx.GetLogger().Debug("ApproveSuggestion", "suggestion:", s)
	var (
		l       = &Ledger[*authpb.Suggestion]{ctx: ctx}
		ol      = &Ledger[common.ItemInterface]{ctx: ctx}
		handler = SuggestionHandler{suggestion: s}
		op      = &authpb.Operation{
			Action:       authpb.Action_ACTION_SUGGEST_APPROVE,
			CollectionId: s.GetPrimaryKey().GetCollectionId(),
			ItemType:     s.GetPrimaryKey().GetItemType(),
			Paths:        nil,
		}
	)

	// Extract the item from the suggestion
	if err = handler.Extract(s); err != nil {
		return nil, oops.Wrap(err)
	}

	// Get the suggestion from the ledger
	if err = l.Get(handler.suggestion); err != nil {
		return nil, oops.
			In("ApproveSuggestion").
			With("Key", s, "ItemType", s.ItemType()).
			Wrap(common.KeyNotFound)
	}

	// Set The paths
	op.Paths = handler.suggestion.GetPaths()

	// Authorize the operation
	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	// Process the Update
	if err = ol.Update(handler.obj, s.GetPaths()); err != nil {
		return nil, oops.Wrap(err)
	}

	if err = l.Delete(handler.suggestion); err != nil {
		return nil, oops.Wrap(err)
	}

	return &handler.obj, nil
}

// ──────────────────────────────────────────────────
// Query Suggested Functions
// ──────────────────────────────────────────────────

// Get GetSuggestion by its ID
func GetSuggestion(ctx TxCtxInterface, s *authpb.Suggestion) (err error) {
	ctx.GetLogger().Debug("GetSuggestion", "suggestion:", s)

	var (
		l = &Ledger[*authpb.Suggestion]{ctx: ctx}

		handler = SuggestionHandler{suggestion: s}
		op      = &authpb.Operation{
			Action:       authpb.Action_ACTION_SUGGEST_VIEW,
			CollectionId: s.GetPrimaryKey().GetCollectionId(),
			ItemType:     s.GetPrimaryKey().GetItemType(),
			Paths:        nil,
		}
	)

	// Extract the item from the suggestion
	if err = handler.Extract(s); err != nil {
		return oops.Wrap(err)
	}

	// Authorize the operation
	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	return l.Get(handler.suggestion)
}

// partialSuggestedKeyList returns a list of suggestions based on the partial key
func PartialSuggestionList(
	ctx TxCtxInterface,
	s *authpb.Suggestion,
	numAttr int,
	bookmark string,
) (list []*authpb.Suggestion, mk string, err error) {
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

	handler := SuggestionHandler{suggestion: s}
	if err = handler.Extract(s); err != nil {
		return nil, "", oops.Wrap(err)
	}

	attr := MakeSuggestionKeyAtter(handler.obj, "")
	if len(attr) == 0 || len(attr) < numAttr {
		return nil, "", common.ItemInvalid
	}

	// attr = append([]string{common.SuggestionItemType}, attr...)
	// Extract the attributes to search for
	// attr = attr[:len(attr)-numAttr]

	attr = lo.DropRight(attr, numAttr)

	ctx.GetLogger().
		Info("GetPartialSuggestedKeyList",
			slog.Group(
				"Key", "ItemType", common.SuggestionItemType,
				slog.Int("numAttr", numAttr),
				slog.Any("attr", attr),
				slog.Group(
					"Paged",
					"Bookmark", bookmark,
					"PageSize", strconv.Itoa(int(ctx.GetPageSize())),
				),
			),
		)

	results, metadata, err := ctx.GetStub().
		GetStateByPartialCompositeKeyWithPagination(
			common.SuggestionItemType,
			attr,
			ctx.GetPageSize(),
			bookmark,
		)
	if err != nil {
		return nil, "", err
	}
	defer results.Close()

	for results.HasNext() {
		queryResponse, err := results.Next()
		if err != nil {
			return nil, "", err
		}
		sug := new(authpb.Suggestion)

		if err := json.Unmarshal(queryResponse.GetValue(), sug); err != nil {
			return nil, "", err
		}

		list = append(list, sug)
	}

	if metadata == nil {
		return nil, "", fmt.Errorf("metadata is nil")
	}

	return list, metadata.GetBookmark(), nil
}

// Get Full Suggestion List, for the collection
func SuggestionListByCollection(ctx TxCtxInterface, collection_id string, bookmark string) (
	list []*authpb.Suggestion, mk string, err error,
) {
	ctx.GetLogger().Debug("GetSuggestionListByCollection", "collection_id:", collection_id)

	s := &authpb.Suggestion{
		PrimaryKey: &authpb.ItemKey{
			CollectionId: collection_id,
		},
	}

	return PartialSuggestionList(ctx, s, 1, "")
}

// Get Suggestion List, for the item
func SuggestionListByItem(
	ctx TxCtxInterface,
	objKey *authpb.ItemKey,
	bookmark string,
) (list []*authpb.Suggestion, mk string, err error) {
	ctx.GetLogger().Debug("GetSuggestionList", "suggestion:", objKey)

	s := &authpb.Suggestion{
		PrimaryKey: objKey,
	}

	num := len(objKey.GetItemIdParts()) + 2

	return PartialSuggestionList(ctx, s, num, "")
}
