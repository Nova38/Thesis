package noauth

import (
	"encoding/json"
	"fmt"

	_ "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v0"
	_ "github.com/nova38/thesis/lib/go/gen/sample/v0"

	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/serializer"
	cc "github.com/nova38/thesis/lib/go/gen/chaincode/auth/common"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/samber/lo"
	"github.com/samber/oops"
)

func BeforeTransaction(ctx *NoAuthCtx) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.HandelBefore(); err != nil {
		return oops.Wrap(err)
	}

	return nil
}

func BuildContract() *contractapi.ContractChaincode {
	contract := new(NoAuthContract)
	contract.BeforeTransaction = BeforeTransaction
	contract.TransactionContextHandler = new(NoAuthCtx)

	sm, err := contractapi.NewChaincode(contract)
	if err != nil {
		fmt.Printf("Error creating No Auth contract: %s", err)
		panic(err)
	}
	sm.TransactionSerializer = &serializer.TxSerializer{}

	return sm
}

// ═════════════════════════════════════════════
// Additonal Functions for the NoAuthContract
// ═════════════════════════════════════════════

// NoAuthContract is the contract for the NoAuth chaincode
func (c *NoAuthContract) Bootstrap(
	ctx common.TxCtxInterface,
	req *cc.BootstrapRequest,
) (res *cc.BootstrapResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	ctx.GetLogger().Info("NoAuthContract.Boostrap")
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	// Check to see if the bootstrap has already been done
	// If so, return an error
	if err = ctx.CheckBootstrap(); err != nil {
		return nil, oops.Wrap(err)
	}

	return res, nil
}

// Test is a function that returns true
func (c *NoAuthContract) Test(ctx *NoAuthCtx) (bool, error) {
	ctx.GetLogger().Info("NoAuthContract.Test")

	return true, nil
}

// TestFail is a function that returns false and and error
func (c *NoAuthContract) TestFail(ctx *NoAuthCtx) (v bool, err error) {

	defer func() { ctx.HandleFnError(&err, recover()) }()

	ctx.GetLogger().Info("NoAuthContract.TestFail")
	e := oops.Errorf("TestFail")
	b := lo.Must(json.MarshalIndent(e, "", "  "))
	ctx.GetLogger().Info(string(b))

	return false, e
}
