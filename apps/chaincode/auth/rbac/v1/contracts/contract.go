package rbac

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/oops"

	// base "github.com/nova38/thesis/lib/go/"

	contracts "github.com/nova38/thesis/lib/go/fabric/auth/contracts"
	auth "github.com/nova38/thesis/lib/go/fabric/auth/state"
	cc "github.com/nova38/thesis/lib/go/gen/chaincode/auth/rbac/schema/v1"
)

// Check if AuthContractImpl implements AuthServiceInterface
var (
	_ auth.TxCtxInterface                 = (*AuthTxCtx)(nil)
	_ cc.RBACServiceInterface[*AuthTxCtx] = (*RbacContractImpl)(nil)
	_ contractapi.ContractInterface       = (*RbacContractImpl)(nil)
)

type RbacContractImpl struct {
	contractapi.Contract
	cc.RBACServiceBase
	contracts.CollectionImpl
}

func (a RbacContractImpl) GetBeforeTransaction() interface{} {
	return a.BeforeTransaction
}

func (a RbacContractImpl) BeforeTransaction(ctx *AuthTxCtx) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.HandelBefore(); err != nil {
		return oops.Wrap(err)
	}

	// Set the authenticator handler
	// ctx.SetAuthenticator("rbac.Authenticate")

	return nil
}

func (a RbacContractImpl) Authenticate(
	ctx *AuthTxCtx,
	ops *authpb.Operation,
) (new_ops authpb.Operation, err error) {
	// TODO implement Authenticate
	panic("implement me")
}

// NewAuthContract
// ───────────────────────────────────────────────────────────────────────────--
// Build Contract
// ───────────────────────────────────────────────────────────────────────────--
func NewAuthContract(baseName string) *RbacContractImpl {
	contract := new(RbacContractImpl)
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
