package contracts

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	"github.com/samber/oops"

	"github.com/nova38/thesis/lib/go/fabric/rbac"
	cc "github.com/nova38/thesis/lib/go/gen/chaincode/rbac/schema/v1"
)

// Check if AuthContractImpl implements AuthServiceInterface
var (
	_ rbac.AuthTxCtxInterface             = (*AuthTxCtx)(nil)
	_ cc.AuthServiceInterface[*AuthTxCtx] = (*AuthContractImpl)(nil)
	_ contractapi.ContractInterface       = (*AuthContractImpl)(nil)
)

type AuthContractImpl struct {
	contractapi.Contract
	cc.AuthServiceBase
}

func (a AuthContractImpl) GetBeforeTransaction() interface{} {
	return a.BeforeTransaction
}

func (a AuthContractImpl) BeforeTransaction(ctx *AuthTxCtx) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.HandelBefore(); err != nil {
		return oops.Wrap(err)
	}

	ops, err := cc.AuthServiceGetTxOperation(ctx.GetFnName())
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

	return err
}

// -----------------------------------------------------------------------------
// Build Contract
// -----------------------------------------------------------------------------
func NewAuthContract(baseName string) *AuthContractImpl {
	contract := new(AuthContractImpl)
	contract.TransactionContextHandler = &AuthTxCtx{}

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
