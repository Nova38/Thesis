package noauth

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/serializer"
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
func (c *NoAuthContract) Boostrap(ctx common.TxCtxInterface) error {
	ctx.GetLogger().Info("NoAuthContract.Boostrap")

	return nil
}

// Test is a function that returns true
func (c *NoAuthContract) Test(ctx *NoAuthCtx) (bool, error) {
	ctx.GetLogger().Info("NoAuthContract.Test")

	return true, nil
}

// TestFail is a function that returns false and and error
func (c *NoAuthContract) TestFail(ctx *NoAuthCtx) (bool, error) {
	ctx.GetLogger().Info("NoAuthContract.TestFail")
	e := oops.Errorf("TestFail")
	b := lo.Must(json.MarshalIndent(e, "", "  "))
	ctx.GetLogger().Info(string(b))

	return false, e
}
