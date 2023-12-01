package roles

import (
	common "github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	"github.com/nova38/thesis/lib/go/fabric/contracts"
	ccpb "github.com/nova38/thesis/lib/go/gen/chaincode/auth/common"
)

type (
	RolesTxCtx struct {
		state.BaseTxCtx
	}
	IdentiyContract struct {
		contracts.ItemContractImpl
	}
)

// type checking
var (
	_ common.TxCtxInterface = (*RolesTxCtx)(
		nil,
	) // _ contracts.ItemContractInterface = (*IdentiyContract)(nil)
	// see if ItemContractImpl implements the interface GenericServiceInterface
	_ ccpb.GenericServiceInterface[common.TxCtxInterface] = (*IdentiyContract)(nil)
)
