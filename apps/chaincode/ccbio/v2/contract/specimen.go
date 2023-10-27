package contract

import (
	"github.com/bufbuild/protovalidate-go"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	"github.com/nova38/thesis/lib/go/fabric/rbac"
	"github.com/nova38/thesis/lib/go/fabric/state"
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v2"
	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
	"github.com/samber/oops"
	"google.golang.org/protobuf/types/known/emptypb"
)

var validator *protovalidate.Validator

// SpecimenContract contract for handling BasicAssets
type SpecimenContractImpl struct {
	contractapi.Contract

	pb.SpecimenServiceBase
}

// AuthServiceInterface

var _ pb.SpecimenServiceInterface = (*SpecimenContractImpl)(nil)
var _ contractapi.ContractInterface = (*SpecimenContractImpl)(nil)

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

func (s *SpecimenContractImpl) BeforeTransaction(ctx rbac.AuthTxCtx) error {
	// Check if the validate is initialized

	if validator == nil {
		v, err := protovalidate.New()
		if err != nil {
			panic(err)
		}
		validator = v
	}

	// Set the operations
	ops, err := pb.SpecimenServiceGetTxOperation(ctx.GetFnName())
	if err != nil {
		return oops.
			In("BeforeTransaction").
			With("fn", ctx.GetFnName()).
			Wrap(err)
	}
	if err := ctx.SetOperation(ops); err != nil {
		return err
	}

	return nil
}

// ────────────────────────────────────────────────────────────
// Query Functions
// ────────────────────────────────────────────────────────────

func (s *SpecimenContractImpl) SpecimenGet(
	ctx rbac.AuthTxCtx,
	req *pb.SpecimenGetRequest,
) (res *pb.SpecimenGetResponse, err error) {
	if err := validator.Validate(req); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_REQUEST_INVALID.String()).
			Wrap(err)
	}

	if err := ctx.IsAuthorized(); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			Code(rbac_pb.Error_ERROR_USER_PERMISSION_DENIED.String()).
			Wrap(err)
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
	if err := state.GetState(&ctx, specimen); err != nil {
		return nil, oops.
			In(ctx.GetFnName()).
			With("id", id).
			Wrap(err)
	}

	return &pb.SpecimenGetResponse{Specimen: specimen}, nil

}

func (s *SpecimenContractImpl) SpecimenGetList(
	ctx rbac.AuthTxCtx,
	req *emptypb.Empty,
) (res *pb.SpecimenGetListResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *SpecimenContractImpl) SpecimenGetByCollection(
	ctx rbac.AuthTxCtx,
	req *pb.SpecimenGetByCollectionRequest,
) (res *pb.SpecimenGetByCollectionResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *SpecimenContractImpl) SpecimenGetHistory(
	ctx rbac.AuthTxCtx,
	req *pb.SpecimenGetHistoryRequest,
) (res *pb.SpecimenGetHistoryResponse, err error) {
	//TODO implement me
	panic("implement me")
}

// ────────────────────────────────────────────────────────────
// Invoke Functions
// ────────────────────────────────────────────────────────────

func (s *SpecimenContractImpl) SpecimenCreate(
	ctx rbac.AuthTxCtx,
	req *pb.SpecimenCreateRequest,
) (res *pb.SpecimenCreateResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *SpecimenContractImpl) SpecimenUpdate(
	ctx rbac.AuthTxCtx,
	req *pb.SpecimenUpdateRequest,
) (res *pb.SpecimenUpdateResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *SpecimenContractImpl) SpecimenDelete(
	ctx rbac.AuthTxCtx,
	req *pb.SpecimenDeleteRequest,
) (res *pb.SpecimenDeleteResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *SpecimenContractImpl) SpecimenHideTx(
	ctx rbac.AuthTxCtx,
	req *pb.SpecimenHideTxRequest,
) (res *pb.SpecimenHideTxResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *SpecimenContractImpl) SpecimenUnHideTx(
	ctx rbac.AuthTxCtx,
	req *pb.SpecimenUnHideTxRequest,
) (res *pb.SpecimenUnHideTxResponse, err error) {
	//TODO implement me
	panic("implement me")
}

// ────────────────────────────────────────────────────────────
// ────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────
// Query Suggested Functions
// ────────────────────────────────────────────────────────────

func (s *SpecimenContractImpl) GetSuggestedUpdate(
	ctx rbac.AuthTxCtx,
	req *pb.GetSuggestedUpdateRequest,
) (res *pb.GetSuggestedUpdateResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *SpecimenContractImpl) GetSuggestedUpdateBySpecimen(
	ctx rbac.AuthTxCtx,
	req *pb.GetSuggestedUpdateBySpecimenRequest,
) (res *pb.GetSuggestedUpdateBySpecimenResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *SpecimenContractImpl) GetSuggestedUpdateByCollection(
	ctx rbac.AuthTxCtx,
	req *pb.GetSuggestedUpdateByCollectionRequest,
) (res *pb.GetSuggestedUpdateByCollectionResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *SpecimenContractImpl) GetSuggestedUpdateList(
	ctx rbac.AuthTxCtx,
	req *emptypb.Empty,
) (res *pb.GetSuggestedUpdateListResponse, err error) {
	//TODO implement me
	panic("implement me")
}

// ────────────────────────────────────────────────────────────
// Invoke Suggested Functions
// ────────────────────────────────────────────────────────────

func (s *SpecimenContractImpl) SuggestedUpdateCreate(
	ctx rbac.AuthTxCtx,
	req *pb.SuggestedUpdateCreateRequest,
) (res *pb.SuggestedUpdateCreateResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *SpecimenContractImpl) SuggestedUpdateApprove(
	ctx rbac.AuthTxCtx,
	req *pb.SuggestedUpdateApproveRequest,
) (res *pb.SuggestedUpdateApproveResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *SpecimenContractImpl) SuggestedUpdateReject(
	ctx rbac.AuthTxCtx,
	req *pb.SuggestedUpdateRejectRequest,
) (res *pb.SuggestedUpdateRejectResponse, err error) {
	//TODO implement me

	panic("implement me")
}
