package contract

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	"github.com/nova38/thesis/lib/go/fabric/rbac"
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

// SpecimenContract contract for handling BasicAssets
type SpecimenContractImpl struct {
	contractapi.Contract
	pb.SpecimenServiceBase
}

// AuthServiceInterface

var _ pb.SpecimenServiceInterface = (*SpecimenContractImpl)(nil)

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

// ────────────────────────────────────────────────────────────
// Query Functions
// ────────────────────────────────────────────────────────────

func (s *SpecimenContractImpl) SpecimenGet(
	ctx rbac.AuthTxCtx,
	req *pb.SpecimenGetRequest,
) (res *pb.SpecimenGetResponse, err error) {
	//TODO implement me
	panic("implement me")
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
