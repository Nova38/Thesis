package identity

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	common "github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/serializer"
	"github.com/samber/oops"
)

func BeforeTransaction(ctx common.TxCtxInterface) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.HandelBefore(); err != nil {
		return oops.Wrap(err)
	}

	// Set the authenticator handler

	return nil
}

func BuildContract() *contractapi.ContractChaincode {
	contract := new(IdentiyContract)
	contract.BeforeTransaction = BeforeTransaction
	contract.TransactionContextHandler = new(IdentityTxCtx)

	sm, err := contractapi.NewChaincode(contract)
	if err != nil {
		fmt.Printf("Error creating No Auth contract: %s", err)
		panic(err)
	}
	sm.TransactionSerializer = &serializer.TxSerializer{}

	return sm
}
