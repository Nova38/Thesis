package state

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"

	"github.com/mennanov/fmutils"
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
)

// ════════════════════════════════════════════════════════
// Suggestion Functions
// ════════════════════════════════════════════════════════

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

// ──────────────────────────────────────────────────
// Invoke Suggested Functions
// ──────────────────────────────────────────────────

func CreateSuggestion[T Object](ctx TxCtxInterface, s *authpb.Suggestion) (err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	sKey := lo.Must(MakeCompositeKey(s))
	if KeyExists(ctx, sKey) {
		return oops.
			With("Key", sKey, "Namespace", s.Namespace()).
			Wrap(common.AlreadyExists)
	}

	obj, err := SuggestionToObject(s)
	if err != nil {
		return oops.Wrap(err)
	}

	objKey := lo.Must(MakeCompositeKey(obj))
	if !KeyExists(ctx, objKey) {
		return oops.
			With("suggestion Key", sKey, "Key", objKey, "Namespace", obj.Namespace()).
			Wrap(common.KeyNotFound)
	}

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_OBJECT_SUGGEST_CREATE,
		CollectionId: s.GetCollectionId(),
		Namespace:    s.GetObjectType(),
		Paths:        nil,
	}
	authorized := lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	if !authorized {
		return oops.Wrap(common.UserPermissionDenied)
	}

	bytes := lo.Must(json.Marshal(s))

	return ctx.GetStub().PutState(sKey, bytes)
}

func DeleteSuggestion[T Object](ctx TxCtxInterface, s *authpb.Suggestion) (err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()
	sKey := lo.Must(MakeCompositeKey(s))
	if !KeyExists(ctx, sKey) {
		return oops.
			In("DeleteSuggestion").
			With("Key", sKey, "Namespace", s.Namespace()).
			Wrap(common.KeyNotFound)
	}

	obj, err := SuggestionToObject(s)
	if err != nil {
		return oops.Wrap(err)
	}

	objKey := lo.Must(MakeCompositeKey(obj))
	if !KeyExists(ctx, objKey) {
		return oops.
			With("suggestion Key", sKey, "Key", objKey, "Namespace", obj.Namespace()).
			Wrap(common.KeyNotFound)
	}

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_OBJECT_SUGGEST_DELETE,
		CollectionId: s.GetCollectionId(),
		Namespace:    s.GetObjectType(),
		Paths:        nil,
	}

	authorized := lo.Must(ctx.Authorize([]*authpb.Operation{op}))
	if !authorized {
		return oops.Wrap(common.UserPermissionDenied)
	}

	return ctx.GetStub().DelState(sKey)
}

func ApproveSuggestion[T Object](ctx TxCtxInterface, s *authpb.Suggestion) (err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = get(ctx, s); err != nil {
		// ctx.GetLogger().Error("ApproveSuggestion", err.Error())

		return oops.
			In("DeleteSuggestion").
			With("Key", s, "Namespace", s.Namespace()).
			Wrap(common.KeyNotFound)
	}

	update, err := SuggestionToObject(s)
	if err != nil {
		return oops.Wrap(err)
	}

	c := proto.Clone(update)

	cur, ok := c.(Object)
	if !ok {
		return oops.In("ApproveSuggestion").Errorf("Object is not an Object")
	}

	if err = get(ctx, cur); err != nil {
		// ctx.GetLogger().Error("ApproveSuggestion", err.Error())
		return oops.
			In("DeleteSuggestion").
			With("Key", s, "Namespace", s.Namespace()).
			Wrap(err)
	}

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_OBJECT_SUGGEST_APPROVE,
		CollectionId: s.GetCollectionId(),
		Namespace:    s.GetObjectType(),
		Paths:        s.GetPaths(),
	}

	if !lo.Must(ctx.Authorize([]*authpb.Operation{op})) {
		return oops.Wrap(common.UserPermissionDenied)
	}

	// Fetch the suggestion object

	fmutils.Overwrite(update, cur, s.GetPaths().GetPaths())
	ctx.GetStub().PutState(lo.Must(MakeCompositeKey(update)), lo.Must(json.Marshal(update)))

	// todo Add LastUpdate to the object

	sKey := lo.Must(MakeCompositeKey(s))

	return ctx.GetStub().DelState(sKey)
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
