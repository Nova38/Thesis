package contract

import (
	"github.com/nova38/thesis/lib/go/fabric/state"
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v2"
	"github.com/samber/oops"
)

// ────────────────────────────────────────────────────────────
// Query Suggested Functions
// ────────────────────────────────────────────────────────────

func (s *SpecimenContractImpl) GetSuggestedUpdate(
	ctx *CCBioTxCtx,
	req *pb.GetSuggestedUpdateRequest,
) (res *pb.GetSuggestedUpdateResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.InitViaReq(req); err != nil {
		return nil, oops.Wrap(err)
	}

	return &pb.GetSuggestedUpdateResponse{SuggestedUpdate: ctx.Suggested}, nil
}

func (s *SpecimenContractImpl) GetSuggestedUpdateBySpecimen(
	ctx *CCBioTxCtx,
	req *pb.GetSuggestedUpdateBySpecimenRequest,
) (res *pb.GetSuggestedUpdateBySpecimenResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.InitViaReq(req); err != nil {
		return nil, oops.Wrap(err)
	}

	list, bm, err := state.GetPagedPartialKeyList(ctx, ctx.Suggested, 2, req.Bookmark)
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("collection_id", ctx.Suggested.Id).
			Wrap(err)
	}

	return &pb.GetSuggestedUpdateBySpecimenResponse{SuggestedUpdates: list, Bookmark: bm}, nil
}

func (s *SpecimenContractImpl) GetSuggestedUpdateByCollection(
	ctx *CCBioTxCtx,
	req *pb.GetSuggestedUpdateByCollectionRequest,
) (res *pb.GetSuggestedUpdateByCollectionResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.InitViaReq(req); err != nil {
		return nil, oops.Wrap(err)
	}

	list, bm, err := state.GetPagedPartialKeyList(ctx, ctx.Suggested, 1, req.Bookmark)
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("collection_id", ctx.Suggested.Id).
			Wrap(err)
	}

	return &pb.GetSuggestedUpdateByCollectionResponse{SuggestedUpdates: list, Bookmark: bm}, nil
}

func (s *SpecimenContractImpl) GetSuggestedUpdateList(
	ctx *CCBioTxCtx,
	req *pb.GetSuggestedUpdateListRequest,
) (res *pb.GetSuggestedUpdateListResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.InitViaReq(req); err != nil {
		return nil, oops.Wrap(err)
	}

	list, bm, err := state.GetPagedFullList(ctx, &pb.SuggestedUpdate{}, req.Bookmark)
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Wrap(err)
	}

	return &pb.GetSuggestedUpdateListResponse{SuggestedUpdates: list, Bookmark: bm}, nil
}

// ────────────────────────────────────────────────────────────
// Invoke Suggested Functions
// ────────────────────────────────────────────────────────────

func (s *SpecimenContractImpl) SuggestedUpdateCreate(
	ctx *CCBioTxCtx,
	req *pb.SuggestedUpdateCreateRequest,
) (res *pb.SuggestedUpdateCreateResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.InitViaReq(req); err != nil {
		return nil, oops.Wrap(err)
	}

	if err = state.Insert(ctx, ctx.Suggested); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("suggested", ctx.Suggested).
			Wrap(err)
	}

	// TODO: make sure SuggestedUpdateCreate works

	return &pb.SuggestedUpdateCreateResponse{SuggestedUpdate: ctx.Suggested}, nil
}

func (s *SpecimenContractImpl) SuggestedUpdateApprove(
	ctx *CCBioTxCtx,
	req *pb.SuggestedUpdateApproveRequest,
) (res *pb.SuggestedUpdateApproveResponse, err error) {
	// TODO implement SuggestedUpdateApprove
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.InitViaReq(req); err != nil {
		return nil, oops.Wrap(err)
	}

	panic("implement me")

	// Remove the suggested update after it is approved
	if err = state.Delete(ctx, ctx.Suggested); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("id", ctx.Suggested.Id).
			Wrap(err)
	}

	return &pb.SuggestedUpdateApproveResponse{
		SuggestedUpdate: ctx.Suggested,
		Specimen:        ctx.Specimen,
	}, nil
}

func (s *SpecimenContractImpl) SuggestedUpdateReject(
	ctx *CCBioTxCtx,
	req *pb.SuggestedUpdateRejectRequest,
) (res *pb.SuggestedUpdateRejectResponse, err error) {
	// TODO implement me SuggestedUpdateReject
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.InitViaReq(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Delete the suggested update
	if err = state.Delete(ctx, ctx.Suggested); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("id", ctx.Suggested.Id).
			Wrap(err)
	}

	return &pb.SuggestedUpdateRejectResponse{SuggestedUpdate: ctx.Suggested}, nil
}
