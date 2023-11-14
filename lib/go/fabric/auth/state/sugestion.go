package state

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
)

// ════════════════════════════════════════════════════════
// Suggestion Functions
// ════════════════════════════════════════════════════════

// ──────────────────────────────────────────────────
// Query Suggested Functions
// ──────────────────────────────────────────────────

// Suggestion returns the suggestion object
func Suggestion[T Object](
	ctx TxCtxInterface,
	obj T,
	suggestionId string,
) (suggestion *authpb.Suggestion, err error) {
	// defer func() { ctx.HandleFnError(&err, recover()) }()

	key := lo.Must(MakeSuggestionKey(ctx, obj, suggestionId))
	if KeyExists(ctx, key) == false {
		return nil, oops.Wrap(common.KeyNotFound)
	}

	op := &authpb.Operation{
		Action:       authpb.Action_ACTION_OBJECT_SUGGEST_CREATE,
		CollectionId: obj.GetCollectionId(),
		Namespace:    obj.Namespace(),
		Paths:        nil,
	}

	if authorized, err := ctx.Authorize([]*authpb.Operation{op}); err != nil {
		return nil, oops.Wrap(common.UserPermissionDenied)
	} else if !authorized {
		return nil, oops.Wrap(common.UserPermissionDenied)
	}

	bytes := lo.Must(ctx.GetStub().GetState(key))
	if err = proto.Unmarshal(bytes, suggestion); err != nil {
		return nil, err
	}

	return suggestion, nil
}

func SuggestionList[T Object](
	ctx TxCtxInterface,
	obj T,
	bookmark string,
) (list []authpb.Suggestion, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	// TODO implement me
	panic("implement me")
}

func SuggestionByCollection[T Object](
	ctx TxCtxInterface,
	obj T,
	bookmark string,
) (list []authpb.Suggestion, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	// TODO implement me
	panic("implement me")
}

func GetSuggestionListByPartialKey[T Object](
	ctx TxCtxInterface,
	obj T,
	numAttr int,
	bookmark string,
) (list []authpb.Suggestion, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	// TODO implement me
	panic("implement me")
}

// ──────────────────────────────────────────────────
// Invoke Suggested Functions
// ──────────────────────────────────────────────────

func CreateSuggestion[T Object](ctx TxCtxInterface, suggestion *authpb.Suggestion) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	// TODO implement me
	panic("implement me")
}

func DeleteSuggestion[T Object](ctx TxCtxInterface, suggestion *authpb.Suggestion) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	// TODO implement me
	panic("implement me")
}

func ApproveSuggestion[T Object](ctx TxCtxInterface, suggestion *authpb.Suggestion) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	// TODO implement me
	panic("implement me")
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
