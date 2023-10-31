package contract

import (
	"github.com/mennanov/fmutils"
	"github.com/nova38/thesis/lib/go/fabric/state"
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v2"
	"github.com/nova38/thesis/lib/go/gen/rbac"
	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// ────────────────────────────────────────────────────────────
// Query Functions
// ────────────────────────────────────────────────────────────

func (s *SpecimenContractImpl) SpecimenGet(
	ctx *CCBioTxCtx,
	req *pb.SpecimenGetRequest,
) (res *pb.SpecimenGetResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.InitViaReq(req); err != nil {
		return nil, oops.Wrap(err)
	}

	return &pb.SpecimenGetResponse{Specimen: ctx.Specimen}, nil
}

func (s *SpecimenContractImpl) SpecimenGetList(
	ctx *CCBioTxCtx,
	req *pb.SpecimenGetListRequest,
) (res *pb.SpecimenGetListResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}

	req.GetPageSize()
	list, bm, err := state.GetPagedFullList(ctx, &pb.Specimen{}, req.Bookmark)
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Wrap(err)
	}

	return &pb.SpecimenGetListResponse{Specimens: list, Bookmark: bm}, nil
}

func (s *SpecimenContractImpl) SpecimenGetByCollection(
	ctx *CCBioTxCtx,
	req *pb.SpecimenGetByCollectionRequest,
) (res *pb.SpecimenGetByCollectionResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.InitViaReq(req); err != nil {
		return nil, oops.Wrap(err)
	}

	list, err := state.GetPartialKeyList(ctx, ctx.Specimen, 1)
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("collection_id", req.Id).
			Wrap(err)
	}

	return &pb.SpecimenGetByCollectionResponse{Specimens: list}, nil
}

func (s *SpecimenContractImpl) SpecimenGetHistory(
	ctx *CCBioTxCtx,
	req *pb.SpecimenGetHistoryRequest,
) (res *pb.SpecimenGetHistoryResponse, err error) {
	// TODO implement SpecimenGetHistory

	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.InitViaReq(req); err != nil {
		return nil, oops.Wrap(err)
	}

	history, err := state.GetHistory(ctx, ctx.Specimen)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	res = &pb.SpecimenGetHistoryResponse{
		History: &pb.Specimen_History{
			Id:      ctx.Specimen.Id,
			Entries: []*rbac.History_Entry{},
		},
	}

	for _, h := range history.Entries {
		a, err := anypb.New(h.State)
		if err != nil {
			return nil, oops.Wrap(err)
		}

		res.History.Entries = append(res.History.Entries, &rbac.History_Entry{
			TxId:      h.TxId,
			Timestamp: h.Timestamp,
			IsDeleted: h.IsDelete,
			IsHidden:  h.IsHidden,
			State:     a,
		})
	}

	return res, nil
}

// ────────────────────────────────────────────────────────────
// Invoke Functions
// ────────────────────────────────────────────────────────────

func (s *SpecimenContractImpl) SpecimenCreate(
	ctx *CCBioTxCtx,
	req *pb.SpecimenCreateRequest,
) (res *pb.SpecimenCreateResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	{ // init
		if err = ctx.Validate(req); err != nil {
			return nil, oops.Wrap(err)
		}

		// Extract the transaction items (collection id and specimen id)
		if err = ctx.ExtractTransactionItems(req); err != nil {
			return nil, oops.Wrap(err)
		}

		if err = state.Get(ctx, ctx.Collection); err != nil {
			return nil, oops.
				In(ctx.GetFnName()).
				With("collection_id", ctx.Collection.Id).
				Wrap(err)
		}

		if err = ctx.IsAuthorized(); err != nil {
			return nil, oops.Wrap(err)
		}
	}

	madeAt, err := ctx.MakeLastModified()

	spec := req.GetSpecimen()
	if spec == nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Errorf("specimen is required")
	}

	// Create the new specimen
	specimen := &pb.Specimen{
		Id:           spec.GetId(),
		Primary:      spec.GetPrimary(),
		Secondary:    spec.GetSecondary(),
		Taxon:        spec.GetTaxon(),
		Georeference: spec.GetGeoreference(),
		Images:       spec.GetImages(),
		Loans:        spec.GetLoans(),
		Grants:       spec.GetGrants(),
		HiddenTxs:    map[string]*pb.Specimen_HiddenTx{},
		LastModified: madeAt,
	}
	// Update all the last modified fields
	specimen.Primary.LastModified = madeAt
	specimen.Secondary.LastModified = madeAt
	specimen.Taxon.LastModified = madeAt
	specimen.Georeference.LastModified = madeAt

	for _, image := range specimen.Images {
		image.LastModified = madeAt
	}
	for _, loan := range specimen.Loans {
		loan.LastModified = madeAt
	}
	for _, grant := range specimen.Grants {
		grant.LastModified = madeAt
	}

	// Insert the new specimen
	if err = state.Insert(ctx, specimen); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("specimen", specimen).
			Wrap(err)
	}

	return &pb.SpecimenCreateResponse{Specimen: specimen}, nil
}

func (s *SpecimenContractImpl) SpecimenUpdate(
	ctx *CCBioTxCtx,
	req *pb.SpecimenUpdateRequest,
) (res *pb.SpecimenUpdateResponse, err error) {
	// TODO implement SpecimenUpdate

	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.InitViaReq(req); err != nil {
		return nil, oops.Wrap(err)
	}

	{
		// panic("TODO: implement SpecimenUpdate")
		// Apply the mask
		// req.UpdateMask.Append(req.Specimen, "id")
		fmutils.Filter(req.Specimen, req.UpdateMask.Paths)
		proto.Merge(ctx.Specimen, req.Specimen)
		mod, err := ctx.MakeLastModified()
		if err != nil {
			return nil, oops.Wrap(err)
		}
		SetLastModByMask(ctx.Specimen, req.UpdateMask, mod)
	}

	// Update the specimen
	if err = state.Update(ctx, ctx.Specimen); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("specimen", ctx.Specimen).
			Wrap(err)
	}

	return &pb.SpecimenUpdateResponse{Specimen: ctx.Specimen}, nil
}

func (s *SpecimenContractImpl) SpecimenDelete(
	ctx *CCBioTxCtx,
	req *pb.SpecimenDeleteRequest,
) (res *pb.SpecimenDeleteResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.InitViaReq(req); err != nil {
		return nil, oops.Wrap(err)
	}

	if err = state.Delete(ctx, ctx.Specimen); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("specimen", ctx.Specimen).
			Wrap(err)
	}

	return &pb.SpecimenDeleteResponse{Specimen: ctx.Specimen}, nil
}

func (s *SpecimenContractImpl) SpecimenHideTx(
	ctx *CCBioTxCtx,
	req *pb.SpecimenHideTxRequest,
) (res *pb.SpecimenHideTxResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.InitViaReq(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Ensure it is not the last tx
	if ctx.Specimen.LastModified.TxId == req.Tx.TxId {
		return nil, oops.
			In(ctx.GetFnName()).
			With("tx_id", req.Tx.TxId).
			Errorf("cannot hide last tx")
	}

	in, err := state.TxIdInHistory(ctx, ctx.Specimen, req.Tx.TxId)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if !in {
		return nil, oops.
			In(ctx.GetFnName()).
			With("tx_id", req.Tx.TxId).
			Errorf("tx_id not in history")
	}

	// Add the tx to the hidden txs
	ctx.Specimen.HiddenTxs[req.Tx.TxId] = req.Tx

	// Update the specimen
	if err = state.Update(ctx, ctx.Specimen); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("specimen", ctx.Specimen).
			Wrap(err)
	}

	return &pb.SpecimenHideTxResponse{Specimen: ctx.Specimen}, nil
}

func (s *SpecimenContractImpl) SpecimenUnHideTx(
	ctx *CCBioTxCtx,
	req *pb.SpecimenUnHideTxRequest,
) (res *pb.SpecimenUnHideTxResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.InitViaReq(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// check if it is in the hidden txs
	if _, ok := ctx.Specimen.HiddenTxs[req.Tx.TxId]; !ok {
		return nil, oops.
			In(ctx.GetFnName()).
			With("tx_id", req.Tx.TxId).
			Errorf("tx_id not in hidden txs")
	}

	// remove the tx from the hidden txs
	delete(ctx.Specimen.HiddenTxs, req.Tx.TxId)

	// Update the specimen
	if err = state.Update(ctx, ctx.Specimen); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("specimen", ctx.Specimen).
			Wrap(err)
	}

	return &pb.SpecimenUnHideTxResponse{Specimen: ctx.Specimen}, nil
}

// ────────────────────────────────────────────────────────────
// ────────────────────────────────────────────────────────────
