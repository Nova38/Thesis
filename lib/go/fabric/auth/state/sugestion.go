package state

type SuggestHandler struct{}

//────────────────────────────────────────────────────────────
//Query Suggested Functions
//────────────────────────────────────────────────────────────

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
