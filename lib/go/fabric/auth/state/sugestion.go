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
	obj           common.ObjectInterface
	current       common.ObjectInterface
	suggestion    *authpb.Suggestion
	bytes         []byte
}

// SuggestionToObject converts a suggestion to an object

func (s SuggestionHandler) Extract(sug *authpb.Suggestion) (err error) {
	if sug == nil {
		return oops.In("Extract").Errorf("Suggestion is nil")
	}
	s.suggestion = sug

	if s.obj, err = SuggestionToObject(s.suggestion); err != nil {
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
			Action:       authpb.Action_ACTION_OBJECT_SUGGEST_CREATE,
			CollectionId: s.GetPrimaryKey().GetCollectionId(),
			Namespace:    s.GetPrimaryKey().GetObjectType(),
			Paths:        s.GetPaths(),
		}
	)

	// Authorize the operation

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	// Extract the object from the suggestion
	if err = handler.Extract(s); err != nil {
		return oops.Wrap(err)
	}

	// Check if the object exists, if it does not, return error
	if !KeyExists(ctx, handler.objKey) {
		return oops.
			With("suggestion ID",
				s.GetSuggestionId(), "Key", handler.objKey,
				"Namespace", handler.obj.Namespace()).
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
			Action:       authpb.Action_ACTION_OBJECT_SUGGEST_DELETE,
			CollectionId: s.GetPrimaryKey().GetCollectionId(),
			Namespace:    s.GetPrimaryKey().GetObjectType(),
			Paths:        nil,
		}
	)
	// Extract the object from the suggestion
	if err = handler.Extract(s); err != nil {
		return oops.Wrap(err)
	}

	if err = l.Get(handler.suggestion); err != nil {
		return oops.
			In("DeleteSuggestion").
			With("Key", s, "Namespace", s.Namespace()).
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
) (updated *common.ObjectInterface, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	ctx.GetLogger().Debug("ApproveSuggestion", "suggestion:", s)
	var (
		l       = &Ledger[*authpb.Suggestion]{ctx: ctx}
		ol      = &Ledger[common.ObjectInterface]{ctx: ctx}
		handler = SuggestionHandler{suggestion: s}
		op      = &authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_SUGGEST_APPROVE,
			CollectionId: s.GetPrimaryKey().GetCollectionId(),
			Namespace:    s.GetPrimaryKey().GetObjectType(),
			Paths:        nil,
		}
	)

	// Extract the object from the suggestion
	if err = handler.Extract(s); err != nil {
		return nil, oops.Wrap(err)
	}

	// Get the suggestion from the ledger
	if err = l.Get(handler.suggestion); err != nil {
		return nil, oops.
			In("ApproveSuggestion").
			With("Key", s, "Namespace", s.Namespace()).
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

// Get Suggestion by its ID
func Suggestion(ctx TxCtxInterface, s *authpb.Suggestion) (err error) {
	ctx.GetLogger().Debug("GetSuggestion", "suggestion:", s)

	var (
		l = &Ledger[*authpb.Suggestion]{ctx: ctx}

		handler = SuggestionHandler{suggestion: s}
		op      = &authpb.Operation{
			Action:       authpb.Action_ACTION_OBJECT_SUGGEST_VIEW,
			CollectionId: s.GetPrimaryKey().GetCollectionId(),
			Namespace:    s.GetPrimaryKey().GetObjectType(),
			Paths:        nil,
		}
	)

	// Extract the object from the suggestion
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
		Action:       authpb.Action_ACTION_OBJECT_SUGGEST_VIEW,
		CollectionId: s.GetPrimaryKey().GetCollectionId(),
		Namespace:    s.GetPrimaryKey().GetObjectType(),
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
		return nil, "", common.ObjectInvalid
	}

	// attr = append([]string{common.SuggestionNamespace}, attr...)
	// Extract the attributes to search for
	// attr = attr[:len(attr)-numAttr]

	attr = lo.DropRight(attr, numAttr)

	ctx.GetLogger().
		Info("GetPartialSuggestedKeyList",
			slog.Group(
				"Key", "Namespace", common.SuggestionNamespace,
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
			common.SuggestionNamespace,
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

		if err := json.Unmarshal(queryResponse.Value, sug); err != nil {
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
		PrimaryKey: &authpb.ObjectKey{
			CollectionId: collection_id,
		},
	}

	return PartialSuggestionList(ctx, s, 1, "")
}

// Get Suggestion List, for the object
func SuggestionListByObject(
	ctx TxCtxInterface,
	objKey *authpb.ObjectKey,
	bookmark string,
) (list []*authpb.Suggestion, mk string, err error) {
	ctx.GetLogger().Debug("GetSuggestionList", "suggestion:", objKey)

	s := &authpb.Suggestion{
		PrimaryKey: objKey,
	}

	num := len(objKey.GetObjectIdParts()) + 2

	return PartialSuggestionList(ctx, s, num, "")
}

//func GetSuggestion[T Object](
//	ctx TxCtxInterface,
//	o T,
//	suggestionId string,
//) (suggestion *authpb.Suggestion, err error) {
//	ctx.GetLogger().Info("GetSuggestedUpdate", slog.Group("suggestionId", suggestionId))
//
//	err = ctx.SetOperation(
//		&authpb.Operation{
//			Action:       authpb.Action_ACTION_OBJECT_SUGGEST_VIEW,
//			CollectionId: o.GetCollectionId(),
//			Namespace:    o.Namespace(),
//			Paths:        nil,
//		})
//
//	if authorize, err := ctx.Authorize(); err != nil {
//		return nil, oops.Wrap(err)
//	} else if !authorize {
//		return nil, oops.Wrap(auth.UserPermissionDenied)
//	}
//
//	suggestion = &authpb.Suggestion{
//		CollectionId: o.GetCollectionId(),
//		ObjectType:   o.Namespace(),
//		ObjectId:     o.FlatKey(),
//		SuggestionId: suggestionId,
//		Paths:        nil,
//		Value:        nil,
//	}
//
//	return suggestion, nil
//}

//func PutSuggestion[T Object](
//	ctx context.TxCtxInterface,
//	o T,
//	suggestion *authpb.Suggestion,
//) {}
//
//
//func DeleteSuggestion[T Object](
//    ctx context.TxCtxInterface,
//    o T,
//    suggestion *authpb.Suggestion,
//) {}
//
//
//func GetSuggestionList[T Object](

//func GetSuggestedUpdateByObject(
//	ctx *context.TxCtxInterface,
//) (res *pb.GetSuggestedUpdateBySpecimenResponse, err error) {
//	defer func() { ctx.HandleFnError(&err, recover()) }()
//
//	if err = ctx.InitViaReq(req); err != nil {
//		return nil, oops.Wrap(err)
//	}
//
//	list, bm, err := GetPagedPartialKeyList(ctx, ctx.Suggested, 2, req.Bookmark)
//	if err != nil {
//		return nil, oops.
//			In(ctx.GetFnName()).
//			With("collection_id", ctx.Suggested.Id).
//			Wrap(err)
//	}
//
//	return &pb.GetSuggestedUpdateBySpecimenResponse{SuggestedUpdates: list, Bookmark: bm}, nil
//}
//
//func GetSuggestedUpdateByCollection(
//	ctx *context.TxCtxInterface,
//) (res *pb.GetSuggestedUpdateByCollectionResponse, err error) {
//	defer func() { ctx.HandleFnError(&err, recover()) }()
//
//	if err = ctx.InitViaReq(req); err != nil {
//		return nil, oops.Wrap(err)
//	}
//
//	list, bm, err := state.GetPagedPartialKeyList(ctx, ctx.Suggested, 1, req.Bookmark)
//	if err != nil {
//		return nil, oops.
//			In(ctx.GetFnName()).
//			With("collection_id", ctx.Suggested.Id).
//			Wrap(err)
//	}
//
//	return &pb.GetSuggestedUpdateByCollectionResponse{SuggestedUpdates: list, Bookmark: bm}, nil
//}
//
//func GetSuggestedUpdateList(
//	ctx *context.TxCtxInterface,
//	req *pb.GetSuggestedUpdateListRequest,
//) (res *pb.GetSuggestedUpdateListResponse, err error) {
//	defer func() { ctx.HandleFnError(&err, recover()) }()
//
//	if err = ctx.InitViaReq(req); err != nil {
//		return nil, oops.Wrap(err)
//	}
//
//	list, bm, err := state.GetPagedFullList(ctx, &pb.SuggestedUpdate{}, req.Bookmark)
//	if err != nil {
//		return nil, oops.
//			In(ctx.GetFnName()).
//			Wrap(err)
//	}
//
//	return &pb.GetSuggestedUpdateListResponse{SuggestedUpdates: list, Bookmark: bm}, nil
//}
//
//// ────────────────────────────────────────────────────────────
//// Invoke Suggested Functions
//// ────────────────────────────────────────────────────────────
//
//func (s SuggestHandler) SuggestedUpdateCreate(
//	ctx *context.TxCtxInterface,
//	req *pb.SuggestedUpdateCreateRequest,
//) (res *pb.SuggestedUpdateCreateResponse, err error) {
//	defer func() { ctx.HandleFnError(&err, recover()) }()
//
//	if err = ctx.InitViaReq(req); err != nil {
//		return nil, oops.Wrap(err)
//	}
//
//	if err = state.Insert(ctx, ctx.Suggested); err != nil {
//		return nil, oops.
//			In(ctx.GetFnName()).
//			With("suggested", ctx.Suggested).
//			Wrap(err)
//	}
//
//	// TODO: make sure SuggestedUpdateCreate works
//
//	return &pb.SuggestedUpdateCreateResponse{SuggestedUpdate: ctx.Suggested}, nil
//}
//
//func (s SuggestHandler) SuggestedUpdateApprove(
//	ctx *context.TxCtxInterface,
//	req *pb.SuggestedUpdateApproveRequest,
//) (res *pb.SuggestedUpdateApproveResponse, err error) {
//	// TODO implement SuggestedUpdateApprove
//	defer func() { ctx.HandleFnError(&err, recover()) }()
//
//	if err = ctx.InitViaReq(req); err != nil {
//		return nil, oops.Wrap(err)
//	}
//
//	panic("implement me")
//
//	// Remove the suggested update after it is approved
//	if err = state.Delete(ctx, ctx.Suggested); err != nil {
//		return nil, oops.
//			In(ctx.GetFnName()).
//			With("id", ctx.Suggested.Id).
//			Wrap(err)
//	}
//
//	return &pb.SuggestedUpdateApproveResponse{
//		SuggestedUpdate: ctx.Suggested,
//		Specimen:        ctx.Specimen,
//	}, nil
//}
//
//func (s SuggestHandler) SuggestedUpdateReject(
//	ctx *context.TxCtxInterface,
//	req *pb.SuggestedUpdateRejectRequest,
//) (res *pb.SuggestedUpdateRejectResponse, err error) {
//	// TODO implement me SuggestedUpdateReject
//	defer func() { ctx.HandleFnError(&err, recover()) }()
//
//	if err = ctx.InitViaReq(req); err != nil {
//		return nil, oops.Wrap(err)
//	}
//
//	// Delete the suggested update
//	if err = state.Delete(ctx, ctx.Suggested); err != nil {
//		return nil, oops.
//			In(ctx.GetFnName()).
//			With("id", ctx.Suggested.Id).
//			Wrap(err)
//	}
//
//	return &pb.SuggestedUpdateRejectResponse{SuggestedUpdate: ctx.Suggested}, nil
//}

// ──────────────────────────────────────────────────
// Query Suggested Functions
// ──────────────────────────────────────────────────

// Suggestion returns the suggestion object
// func Suggestion[T Object](
// 	ctx TxCtxInterface,
// 	obj T,
// 	suggestionId string,
// ) (suggestion *authpb.Suggestion, err error) {
// 	// defer func() { ctx.HandleFnError(&err, recover()) }()

// 	key := lo.Must(MakeSuggestionKey(obj, suggestionId))
// 	if KeyExists(ctx, key) == false {
// 		return nil, oops.Wrap(common.KeyNotFound)
// 	}

// 	op := &authpb.Operation{
// 		Action:       authpb.Action_ACTION_OBJECT_SUGGEST_CREATE,
// 		CollectionId: obj.GetCollectionId(),
// 		Namespace:    obj.Namespace(),
// 		Paths:        nil,
// 	}

// 	if authorized, err := ctx.Authorize([]*authpb.Operation{op}); err != nil {
// 		return nil, oops.Wrap(common.UserPermissionDenied)
// 	} else if !authorized {
// 		return nil, oops.Wrap(common.UserPermissionDenied)
// 	}

// 	bytes := lo.Must(ctx.GetStub().GetState(key))
// 	if err = proto.Unmarshal(bytes, suggestion); err != nil {
// 		return nil, err
// 	}

// 	return suggestion, nil
// }

// func SuggestionList[T Object](
// 	ctx TxCtxInterface,
// 	bookmark string,
// ) (list []*authpb.Suggestion, mk string, err error) {
// 	sug := &authpb.Suggestion{}

// 	if list, mk, err = List(ctx, sug, bookmark); err != nil {
// 		return nil, "", oops.Wrap(err)
// 	}

// 	return list, mk, nil
// }

// func SuggestionByCollection[T Object](
// 	ctx TxCtxInterface,
// 	obj T,
// 	bookmark string,
// ) (list []authpb.Suggestion, err error) {
// 	// defer func() { ctx.HandleFnError(&err, recover()) }()
// }

// func SuggestionListByPartialKey[T Object](
// 	ctx TxCtxInterface,
// 	obj T,
// 	numAttr int,
// 	bookmark string,
// ) (list []*authpb.Suggestion, mk string, err error) {
// 	var (
// 		attr = lo.Must(obj.Key())
// 		op   = &authpb.Operation{
// 			Action:       authpb.Action_ACTION_OBJECT_SUGGEST_VIEW,
// 			CollectionId: obj.GetCollectionId(),
// 			Namespace:    obj.Namespace(),
// 		}
// 		authorized = lo.Must(ctx.Authorize([]*authpb.Operation{op}))
// 	)

// 	if !authorized {
// 		return nil, "", oops.Wrap(common.UserPermissionDenied)
// 	}

// 	if len(attr) == 0 || len(attr) < numAttr {
// 		return nil, "", common.ObjectInvalid
// 	}

// 	attr = attr[:len(attr)-numAttr]
// 	results, metadata, err := ctx.GetStub().
// 		GetStateByPartialCompositeKeyWithPagination(obj.Namespace(), attr, ctx.GetPageSize(), bookmark)
// 	if err != nil {
// 		return nil, "", err
// 	}
// 	defer ctx.CloseQueryIterator(results)

// 	for results.HasNext() {
// 		queryResponse := lo.Must(results.Next())
// 		tmp := new(authpb.Suggestion)

// 		lo.Must0(json.Unmarshal(queryResponse.Value, tmp))

// 		list = append(list, tmp)
// 	}

// 	if metadata == nil {
// 		return nil, "", fmt.Errorf("metadata is nil")
// 	}

// 	return list, metadata.GetBookmark(), nil
// }
