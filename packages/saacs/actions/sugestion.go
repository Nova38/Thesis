package actions

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/mennanov/fmutils"
	"github.com/nova38/thesis/packages/saacs/common"
	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
)

// ════════════════════════════════════════════════════════
// Suggestion Functions
// ════════════════════════════════════════════════════════

// ──────────────────────────────────────────────────
// Helpers
// ──────────────────────────────────────────────────

// type SuggestionHandler struct {
//	suggestionKey string
//	objKey        string
//	obj           common.ItemInterface
//	suggestion    *authpb.Suggestion
//	bytes         []byte
//}

// SuggestionToItem converts a suggestion to an item

// func (s SuggestionHandler) Extract(sug *authpb.Suggestion) (err error) {
//	if sug == nil {
//		return oops.In("Extract").Errorf("Suggestion is nil")
//	}
//	s.suggestion = sug
//
//	if s.obj, err = common.SuggestionToItem(s.suggestion); err != nil {
//		return oops.Wrap(err)
//	}
//
//	if s.objKey, err = common.MakePrimaryKey(s.obj); err != nil {
//		return oops.Wrap(err)
//	}
//
//	if s.suggestionKey, err = common.MakeSuggestionPrimaryKey(s.obj, s.suggestion.GetSuggestionId()); err != nil {
//		return oops.Wrap(err)
//	}
//
//	if s.bytes, err = json.Marshal(s.obj); err != nil {
//		return oops.Wrap(err)
//	}
//
//	return nil
//}

// ──────────────────────────────────────────────────
// Invoke Suggested Functions
// ──────────────────────────────────────────────────

func SuggestionCreate(ctx common.TxCtxInterface, s *authpb.Suggestion) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	ctx.GetLogger().Debug("CreateSuggestion", "suggestion:", s)

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

	if pKey, err := common.MakeItemKeyPrimary(s.GetPrimaryKey()); err != nil {
		return oops.Wrap(err)
	} else if !common.KeyExists(ctx, pKey) {
		return oops.
			With("suggestion ID",
				s.GetSuggestionId(), "Key", pKey,
				"ItemType").
			Wrap(common.KeyNotFound)

	}

	// Make the suggestion key
	sKey, err := common.MakeItemKeySuggestion(s.GetPrimaryKey(), s.GetSuggestionId())
	if err != nil {
		return oops.
			Hint("Make Item Suggestion Key Failed").
			With("Primary Key", s.GetPrimaryKey()).
			Wrap(err)
	}

	bytes, err := json.Marshal(s)
	if err != nil {
		return oops.Hint("Json Marshal Faild").Wrap(err)
	}

	return ctx.GetStub().PutState(sKey, bytes)
	// Check if the item exists, if it does not, return error

}

func SuggestionDelete(ctx common.TxCtxInterface, s *authpb.Suggestion) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	ctx.GetLogger().Debug("DeleteSuggestion", "suggestion:", s)

	// Authorize Viewing and Retrieving the suggestion
	if err = GetSuggestion(ctx, s); err != nil {
		return oops.
			Hint("Failed to be able to read suggestions").
			Wrap(err)
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

	sKey, err := common.MakeItemKeySuggestion(s.GetPrimaryKey(), s.GetSuggestionId())
	if err != nil {
		return oops.Wrap(err)
	}

	return ctx.GetStub().DelState(sKey)
}

func SuggestionApprove(
	ctx common.TxCtxInterface,
	s *authpb.Suggestion,
) (updated *common.ItemInterface, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	ctx.GetLogger().Debug("ApproveSuggestion", "suggestion:", s)

	// Authorize Viewing and Retrieving the suggestion
	if err = GetSuggestion(ctx, s); err != nil {
		return nil, oops.Wrap(err)
	}
	sKey, err := common.MakeItemKeySuggestion(s.GetPrimaryKey(), s.GetSuggestionId())
	if err != nil {
		return nil, oops.Wrap(err)
	}

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_SUGGEST_APPROVE,
		CollectionId: s.GetPrimaryKey().GetCollectionId(),
		ItemType:     s.GetPrimaryKey().GetItemType(),
		Paths:        s.GetPaths(),
	}

	// Authorize the operation
	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	// Process the Update
	primary, err := common.ItemKeyToItem(s.GetPrimaryKey())
	if err != nil {
		return nil, oops.Wrap(err)
	} else if err = PrimaryGet(ctx, primary); err != nil {
		return nil, oops.Hint("Shouldn't be possible").Wrap(common.KeyNotFound)
	}

	update, err := s.GetValue().UnmarshalNew()
	if err != nil {
		return nil, oops.Wrap(err)
	}

	// Apply the mask to the Updating item
	// fmutils.Overwrite(primary, update, s.GetPaths().GetPaths())
	fmutils.Filter(primary, s.GetPaths().GetPaths())
	proto.Merge(primary, update)
	pKey, err := common.MakePrimaryKey(primary)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if bytes, err := json.Marshal(primary); err != nil {
		return nil, oops.Wrap(err)
	} else if err = ctx.GetStub().PutState(pKey, bytes); err != nil {
		return nil, oops.Wrap(err)
	}

	// Delete the suggestion

	if err = ctx.GetStub().DelState(sKey); err != nil {
		return nil, oops.Wrap(err)
	}

	return &primary, nil
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

	sKey, err := common.MakeItemKeySuggestion(s.GetPrimaryKey(), s.GetSuggestionId())
	if err != nil {
		return oops.Wrap(err)
	}

	if bytes, err := ctx.GetStub().GetState(sKey); err != nil {
		return oops.Wrap(err)
	} else if bytes == nil && err == nil {
		return oops.
			With("suggestion ID",
				s.GetSuggestionId(), "Key", sKey,
				"ItemType").
			Wrap(common.KeyNotFound)
	} else if err := json.Unmarshal(bytes, s); err != nil {
		return oops.Hint("Json Unmarshal Failed").Wrap(err)
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

	if s.GetPrimaryKey() == nil {
		return nil, "", oops.Wrap(common.ItemInvalid)
	}

	attr := common.MakeItemKeySuggestionKeyAttr(s.GetPrimaryKey(), "")
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
	defer func(results shim.StateQueryIteratorInterface) {
		err := results.Close()
		if err != nil {
			ctx.GetLogger().Error("GetPartialSuggestedKeyList", "Error", err)
		}
	}(results)

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

	num := len(objKey.GetItemIdParts()) + 2

	return PartialSuggestionList(ctx, s, num, bookmark)
}
