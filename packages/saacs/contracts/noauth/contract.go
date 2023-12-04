package noauth

import (
	"encoding/json"
	"fmt"

	_ "github.com/nova38/thesis/packages/saacs/gen/biochain/v1"
	_ "github.com/nova38/thesis/packages/saacs/gen/sample/v0"
	"github.com/nova38/thesis/packages/saacs/serializer"

	"github.com/nova38/thesis/packages/saacs/auth/common"
	"github.com/nova38/thesis/packages/saacs/auth/state"
	cc "github.com/nova38/thesis/packages/saacs/gen/chaincode/common"

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

func NoSubBeforeTransaction(ctx *NoAuthCtx) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.HandelBefore(); err != nil {
		return oops.Wrap(err)
	}

	ctx.BaseTxCtx.EnableHiddenTx = false
	ctx.BaseTxCtx.EnableSuggestion = false

	return nil
}
func NoSubBuildContract() *contractapi.ContractChaincode {
	contract := new(NoAuthContract)
	contract.BeforeTransaction = NoSubBeforeTransaction
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
	if bootstraped, err := ctx.CheckBootstrap(); err != nil {
		return nil, oops.Wrap(err)
	} else if bootstraped {
		ctx.GetLogger().Warn("Bootstrap already done")
		return nil, oops.Errorf("Bootstrap already done")
	}

	for _, col := range req.GetCollections() {
		colKey := lo.Must(common.MakePrimaryKey(col))

		colBytes := lo.Must(json.Marshal(col))

		lo.Must0(ctx.GetStub().PutState(colKey, colBytes), "PutCollection")

	}

	return &cc.BootstrapResponse{Success: true}, nil
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

func (c *NoAuthContract) CreateCollection(
	ctx common.TxCtxInterface,
	req *cc.CreateCollectionRequest,
) (res *cc.CreateCollectionResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	ctx.GetLogger().Info("NoAuthContract.CreateCollection")

	col := req.GetCollection()

	colKey := lo.Must(common.MakePrimaryKey(col))
	if state.Exists(ctx, colKey) {
		return nil, oops.Errorf("Collection already exists")
	}

	colBytes := lo.Must(json.Marshal(col))

	lo.Must0(ctx.GetStub().PutState(colKey, colBytes), "PutCollection")

	return &cc.CreateCollectionResponse{Collection: col}, nil
}
