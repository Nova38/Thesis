package contract

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v2"
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

// -----------------------------------------------------------------------------
// Build Contract
// -----------------------------------------------------------------------------
func NewSpecimenContract(baseName string) *SpecimenContractImpl {
	contract := new(SpecimenContractImpl)
	contract.TransactionContextHandler = &CCBioTxCtx{}

	contract.Info = metadata.InfoMetadata{
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
	}

	contract.Name = baseName + ".AuthContract"

	return contract
}
