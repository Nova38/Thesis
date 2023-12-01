package roles

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/serializer"

	"github.com/samber/oops"
)

// BuildContract builds the roles contract
func BuildContract() *contractapi.ContractChaincode {
	contract := new(IdentiyContract)
	contract.BeforeTransaction = BeforeTransaction
	contract.TransactionContextHandler = new(RolesTxCtx)

	sm, err := contractapi.NewChaincode(contract)
	if err != nil {
		fmt.Printf("Error creating No Auth contract: %s", err)
		panic(err)
	}
	sm.TransactionSerializer = &serializer.TxSerializer{}

	return sm
}

// BeforeTransaction is the before transaction handler
func BeforeTransaction(ctx common.TxCtxInterface) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.HandelBefore(); err != nil {
		return oops.Wrap(err)
	}

	return nil
}

// Boostrap is the boostrap function
func Boostrap(ctx common.TxCtxInterface) error {
	ctx.GetLogger().Info("NoAuthContract.Boostrap")

	return nil
}
