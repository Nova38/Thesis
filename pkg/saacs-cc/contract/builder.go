package contract

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nova38/saacs/pkg/saacs-cc/common"
	"github.com/nova38/saacs/pkg/saacs-cc/serializer"
	cc "github.com/nova38/saacs/pkg/saacs-protos/saacs/chaincode/v0"
	"github.com/samber/oops"
)

type ContractImpl struct {
	contractapi.Contract
	cc.ItemServiceBase
	cc.UtilsServiceBase
}

// see if ItemContractImpl implements the interface GenericServiceInterface
var _ cc.ItemServiceInterface[common.TxCtxInterface] = (*ContractImpl)(nil)
var _ cc.UtilsServiceInterface[common.TxCtxInterface] = (*ContractImpl)(nil)

// ════════════════════════════════════ Build ═══════════════════════════════════════
func BeforeTransaction(ctx common.TxCtxInterface) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.HandelBefore(); err != nil {
		return oops.Wrap(err)
	}

	return nil
}

func BuildContract() *contractapi.ContractChaincode {
	contract := new(ContractImpl)
	contract.BeforeTransaction = BeforeTransaction
	contract.TransactionContextHandler = new(TxCtx)
	oops.StackTraceMaxDepth = 5

	sm, err := contractapi.NewChaincode(contract)
	if err != nil {
		fmt.Printf("Error creating contract: %s", err)
		panic(err)
	}

	// todo check if we want to use a binary serializer
	sm.TransactionSerializer = &serializer.PbTxSerializer{}

	return sm
}
