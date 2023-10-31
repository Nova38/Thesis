package contract

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	"github.com/nova38/thesis/lib/go/fabric/state"
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v2"
	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
	"github.com/samber/oops"
)

// SpecimenContract contract for handling BasicAssets
type SpecimenContractImpl struct {
	contractapi.Contract

	pb.SpecimenServiceBase
}

// AuthServiceInterface

var (
	_ pb.SpecimenServiceInterface[*CCBioTxCtx] = (*SpecimenContractImpl)(nil)
	_ contractapi.ContractInterface            = (*SpecimenContractImpl)(nil)
)

func BuildSpecimenContract() *SpecimenContractImpl {
	return &SpecimenContractImpl{
		Contract: contractapi.Contract{
			Name: "ccbio.Specimen",
			// BeforeTransaction: ,
			Info: metadata.InfoMetadata{
				Description: "",
				Title:       "Biochain Chaincode",
				Contact: &metadata.ContactMetadata{
					Name:  "Thomas Atkins",
					URL:   "https://biochain.iitc.ku.edu",
					Email: "tom@ku.edu",
				},
				License: &metadata.LicenseMetadata{
					Name: "MIT",
					URL:  "https://example.com",
				},
				Version: "latest",
			},
		},
	}
}

// ────────────────────────────────────────────────────────────

func (s *SpecimenContractImpl) GetBeforeTransaction() interface{} {
	return s.BeforeTransaction
}

func (s *SpecimenContractImpl) BeforeTransaction(ctx *CCBioTxCtx) (err error) {
	// Check if the validate is initialized
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.HandelBefore(); err != nil {
		return oops.Wrap(err)
	}
	// Set the operations
	ops, err := pb.SpecimenServiceGetTxOperation(ctx.GetFnName())
	if err != nil {
		return oops.
			In("BeforeTransaction").
			With("fn", ctx.GetFnName()).
			Wrap(err)
	}
	if err = ctx.SetOperation(ops); err != nil {
		return oops.
			In("BeforeTransaction").
			With("fn", ctx.GetFnName()).
			Wrap(err)
	}

	return nil
}

// ────────────────────────────────────────────────────────────
// Query Functions
// ────────────────────────────────────────────────────────────

func (s *SpecimenContractImpl) SpecimenGet(
	ctx *CCBioTxCtx,
	req *pb.SpecimenGetRequest,
) (res *pb.SpecimenGetResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Extract the collection id
	colId, err := extractCollectionId(req)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if _, err = ctx.SetCollection(colId); err != nil {
		return nil, oops.Wrap(err)
	}

	// Authorize the operation
	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}

	id := req.GetId()
	if id == nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Errorf("id is required")
	}

	specimen := &pb.Specimen{
		Id: &pb.Specimen_Id{
			CollectionId: id.CollectionId,
			Id:           id.Id,
		},
	}
	if err = state.Get(&ctx, specimen); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("id", id).
			Wrap(err)
	}

	return &pb.SpecimenGetResponse{Specimen: specimen}, nil
}

func (s *SpecimenContractImpl) SpecimenGetList(
	ctx *CCBioTxCtx,
) (res *pb.SpecimenGetListResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}

	list, err := state.GetFullList(&ctx, &pb.Specimen{})
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Wrap(err)
	}

	return &pb.SpecimenGetListResponse{Specimens: list}, nil
}

func (s *SpecimenContractImpl) SpecimenGetByCollection(
	ctx *CCBioTxCtx,
	req *pb.SpecimenGetByCollectionRequest,
) (res *pb.SpecimenGetByCollectionResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Extract the collection id
	colId, err := extractCollectionId(req)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if _, err = ctx.SetCollection(colId); err != nil {
		return nil, oops.Wrap(err)
	}

	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}

	collectionId := req.Id
	if collectionId == nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Errorf("collection id is required")
	}

	spec := &pb.Specimen{Id: &pb.Specimen_Id{CollectionId: collectionId.CollectionId}}

	list, err := state.GetPartialKeyList(&ctx, spec, 1)
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("collection_id", collectionId).
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

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Extract the collection id
	colId, err := extractCollectionId(req)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if _, err = ctx.SetCollection(colId); err != nil {
		return nil, oops.Wrap(err)
	}

	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}
	panic("implement me")
}

// ────────────────────────────────────────────────────────────
// Invoke Functions
// ────────────────────────────────────────────────────────────

func (s *SpecimenContractImpl) SpecimenCreate(
	ctx *CCBioTxCtx,
	req *pb.SpecimenCreateRequest,
) (res *pb.SpecimenCreateResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Extract the collection id
	colId, err := extractCollectionId(req)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if _, err = ctx.SetCollection(colId); err != nil {
		return nil, oops.Wrap(err)
	}

	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
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
	if err = state.Insert(&ctx, specimen); err != nil {
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

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}
	panic("implement me")
}

func (s *SpecimenContractImpl) SpecimenDelete(
	ctx *CCBioTxCtx,
	req *pb.SpecimenDeleteRequest,
) (res *pb.SpecimenDeleteResponse, err error) {
	// TODO implement SpecimenDelete

	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Extract the collection id
	colId, err := extractCollectionId(req)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if _, err = ctx.SetCollection(colId); err != nil {
		return nil, oops.Wrap(err)
	}
	// Authorize the operation
	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}

	panic("implement me")
}

func (s *SpecimenContractImpl) SpecimenHideTx(
	ctx *CCBioTxCtx,
	req *pb.SpecimenHideTxRequest,
) (res *pb.SpecimenHideTxResponse, err error) {
	// TODO implement SpecimenHideTx

	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Extract the collection id
	colId, err := extractCollectionId(req)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if _, err = ctx.SetCollection(colId); err != nil {
		return nil, oops.Wrap(err)
	}
	// Authorize the operation
	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}
	panic("implement me")
}

func (s *SpecimenContractImpl) SpecimenUnHideTx(
	ctx *CCBioTxCtx,
	req *pb.SpecimenUnHideTxRequest,
) (res *pb.SpecimenUnHideTxResponse, err error) {
	// TODO implement SpecimenUnHideTx
	recover()

	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}

	// Extract the collection id
	colId, err := extractCollectionId(req)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if _, err = ctx.SetCollection(colId); err != nil {
		return nil, oops.Wrap(err)
	}
	// Authorize the operation
	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}
	panic("implement me")
}

// ────────────────────────────────────────────────────────────
// ────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────
// Query Suggested Functions
// ────────────────────────────────────────────────────────────

func (s *SpecimenContractImpl) GetSuggestedUpdate(
	ctx *CCBioTxCtx,
	req *pb.GetSuggestedUpdateRequest,
) (res *pb.GetSuggestedUpdateResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}
	uId := req.GetId()
	if uId == nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Errorf("id is required")
	}

	colId, err := extractCollectionId(req)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if _, err = ctx.SetCollection(colId); err != nil {
		return nil, oops.Wrap(err)
	}

	// Authorize the operation
	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}

	// Business logic
	update := &pb.SuggestedUpdate{Id: uId}

	if err = state.Get(&ctx, update); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("id", uId).
			Wrap(err)
	}

	return &pb.GetSuggestedUpdateResponse{SuggestedUpdate: update}, nil
}

func (s *SpecimenContractImpl) GetSuggestedUpdateBySpecimen(
	ctx *CCBioTxCtx,
	req *pb.GetSuggestedUpdateBySpecimenRequest,
) (res *pb.GetSuggestedUpdateBySpecimenResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}
	// Extract the collection id
	colId, err := extractCollectionId(req)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if _, err = ctx.SetCollection(colId); err != nil {
		return nil, oops.Wrap(err)
	}
	// Authorize the operation
	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}

	id := req.GetId()
	if id == nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Errorf("id is required")
	}

	spec := &pb.SuggestedUpdate{Id: &pb.SuggestedUpdate_Id{
		SpecimenId: &pb.Specimen_Id{
			CollectionId: id.CollectionId,
			Id:           id.Id,
		},
	}}

	list, err := state.GetPartialKeyList(&ctx, spec, 2)
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("collection_id", id).
			Wrap(err)
	}

	return &pb.GetSuggestedUpdateBySpecimenResponse{SuggestedUpdates: list}, nil
}

func (s *SpecimenContractImpl) GetSuggestedUpdateByCollection(
	ctx *CCBioTxCtx,
	req *pb.GetSuggestedUpdateByCollectionRequest,
) (res *pb.GetSuggestedUpdateByCollectionResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}
	// Extract the collection id
	colId, err := extractCollectionId(req)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if _, err = ctx.SetCollection(colId); err != nil {
		return nil, oops.Wrap(err)
	}
	// Authorize the operation

	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}

	id := req.GetId()
	if id == nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Errorf("id is required")
	}

	spec := &pb.SuggestedUpdate{Id: &pb.SuggestedUpdate_Id{
		SpecimenId: &pb.Specimen_Id{
			CollectionId: id.CollectionId,
		},
	}}

	list, err := state.GetPartialKeyList(&ctx, spec, 1)
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("collection_id", id).
			Wrap(err)
	}

	return &pb.GetSuggestedUpdateByCollectionResponse{SuggestedUpdates: list}, nil
}

func (s *SpecimenContractImpl) GetSuggestedUpdateList(
	ctx *CCBioTxCtx,
) (res *pb.GetSuggestedUpdateListResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}

	list, err := state.GetFullList(&ctx, &pb.SuggestedUpdate{})
	if err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Wrap(err)
	}

	return &pb.GetSuggestedUpdateListResponse{SuggestedUpdates: list}, nil
}

// ────────────────────────────────────────────────────────────
// Invoke Suggested Functions
// ────────────────────────────────────────────────────────────

func (s *SpecimenContractImpl) SuggestedUpdateCreate(
	ctx *CCBioTxCtx,
	req *pb.SuggestedUpdateCreateRequest,
) (res *pb.SuggestedUpdateCreateResponse, err error) {
	// TODO implement SuggestedUpdateCreate

	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}
	// Extract the collection id
	colId, err := extractCollectionId(req)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if _, err = ctx.SetCollection(colId); err != nil {
		return nil, oops.Wrap(err)
	}
	// Authorize the operation

	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}
	panic("implement me")
}

func (s *SpecimenContractImpl) SuggestedUpdateApprove(
	ctx *CCBioTxCtx,
	req *pb.SuggestedUpdateApproveRequest,
) (res *pb.SuggestedUpdateApproveResponse, err error) {
	// TODO implement SuggestedUpdateApprove
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}
	// Extract the collection id
	colId, err := extractCollectionId(req)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if _, err = ctx.SetCollection(colId); err != nil {
		return nil, oops.Wrap(err)
	}
	// Authorize the operation
	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}
	panic("implement me")
}

func (s *SpecimenContractImpl) SuggestedUpdateReject(
	ctx *CCBioTxCtx,
	req *pb.SuggestedUpdateRejectRequest,
) (res *pb.SuggestedUpdateRejectResponse, err error) {
	// TODO implement me SuggestedUpdateReject
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.Validate(req); err != nil {
		return nil, oops.Wrap(err)
	}
	// Extract the collection id
	colId, err := extractCollectionId(req)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	if _, err = ctx.SetCollection(colId); err != nil {
		return nil, oops.Wrap(err)
	}
	// Authorize the operation
	if err = ctx.IsAuthorized(); err != nil {
		return nil, oops.Wrap(err)
	}

	// Get the suggested update
	id := req.GetId()
	if id == nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Errorf("id is required")
	}

	suggestedUpdate := &pb.SuggestedUpdate{Id: id}
	if err = state.Get(&ctx, suggestedUpdate); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("id", id).
			Wrap(err)
	}

	// Delete the suggested update
	if err = state.Delete(&ctx, suggestedUpdate); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("id", id).
			Wrap(err)
	}

	return &pb.SuggestedUpdateRejectResponse{SuggestedUpdate: suggestedUpdate}, nil
}
