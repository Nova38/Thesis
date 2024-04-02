package noauth

import (
	"encoding/json"
	"fmt"

	"github.com/nova38/saacs/apps/saacs-cc/serializer"
	authpb "github.com/nova38/saacs/libs/saacs-protos-go/auth/v1"
	_ "github.com/nova38/saacs/libs/saacs-protos-go/biochain/v1"
	_ "github.com/nova38/saacs/libs/saacs-protos-go/sample/v0"

	"github.com/nova38/saacs/apps/saacs-cc/common"
	"github.com/nova38/saacs/apps/saacs-cc/state"
	cc "github.com/nova38/saacs/libs/saacs-protos-go/chaincode/common"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/samber/lo"
	"github.com/samber/oops"
)

func BeforeTransaction(ctx *Ctx) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.HandelBefore(); err != nil {
		return oops.Wrap(err)
	}

	return nil
}

func BuildContract() *contractapi.ContractChaincode {
	contract := new(Contract)
	contract.BeforeTransaction = BeforeTransaction
	contract.TransactionContextHandler = new(Ctx)

	sm, err := contractapi.NewChaincode(contract)
	if err != nil {
		fmt.Printf("Error creating No Auth contract: %s", err)
		panic(err)
	}
	sm.TransactionSerializer = &serializer.TxSerializer{}

	return sm
}

func NoSubBeforeTransaction(ctx *Ctx) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.HandelBefore(); err != nil {
		return oops.Wrap(err)
	}

	ctx.BaseTxCtx.EnableHiddenTx = false
	ctx.BaseTxCtx.EnableSuggestion = false

	return nil
}
func NoSubBuildContract() *contractapi.ContractChaincode {
	contract := new(Contract)
	contract.BeforeTransaction = NoSubBeforeTransaction
	contract.TransactionContextHandler = new(Ctx)

	sm, err := contractapi.NewChaincode(contract)
	if err != nil {
		fmt.Printf("Error creating No Auth contract: %s", err)
		panic(err)
	}
	sm.TransactionSerializer = &serializer.TxSerializer{}

	return sm
}

// ═════════════════════════════════════════════
// Additional Functions for the NoAuthContract
// ═════════════════════════════════════════════

// NoAuthContract is the contract for the NoAuth saacs-cc
func (c *Contract) Bootstrap(
	ctx common.TxCtxInterface,
	req *cc.BootstrapRequest,
) (res *cc.BootstrapResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	ctx.GetLogger().Info("NoAuthContract.Bootstrap")
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

		err := state.Ledger[*authpb.Collection]{}.PrimaryCreate(ctx, col)
		if err != nil {
			return nil, ctx.ErrorBase().Wrap(err)
		}

	}

	return &cc.BootstrapResponse{Success: true}, nil
}

// Test is a function that returns true
func (c *Contract) Test(ctx *Ctx) (bool, error) {
	ctx.GetLogger().Info("NoAuthContract.Test")

	return true, nil
}

// TestFail is a function that returns false and and error
func (c *Contract) TestFail(ctx *Ctx) (v bool, err error) {

	defer func() { ctx.HandleFnError(&err, recover()) }()

	ctx.GetLogger().Info("NoAuthContract.TestFail")
	e := oops.Errorf("TestFail")
	b := lo.Must(json.MarshalIndent(e, "", "  "))
	ctx.GetLogger().Info(string(b))

	return false, e
}

func (c *Contract) CreateCollection(
	ctx common.TxCtxInterface,
	req *cc.CreateCollectionRequest,
) (res *cc.CreateCollectionResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	ctx.GetLogger().Info("NoAuthContract.CreateCollection")

	col := req.GetCollection()

	if ctx.Validate(col) != nil {
		return nil, ctx.ErrorBase().Wrap(err)
	}

	// Make Sure the auth types are in the collection
	authTypes := []string{
		common.CollectionItemType,
		common.RoleItemType,
		common.UserCollectionRolesItemType,
	}
	col.ItemTypes = append(col.GetItemTypes(), authTypes...)
	col.ItemTypes = lo.Uniq(col.GetItemTypes()) // Deduplicate the item types

	err = state.Ledger[*authpb.Collection]{}.PrimaryCreate(ctx, col)

	if err != nil {
		return nil, ctx.ErrorBase().Wrap(err)
	}

	return &cc.CreateCollectionResponse{Collection: col}, nil
}
