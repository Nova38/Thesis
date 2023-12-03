package noauth

import (
	"github.com/nova38/thesis/packages/fabric/auth/state"
	"github.com/nova38/thesis/packages/fabric/contracts"
	common "github.com/nova38/thesis/packages/saacs/auth/common"
	ccpb "github.com/nova38/thesis/packages/saacs/genchaincode/auth/common"
)

type (
	NoAuthContract struct {
		contracts.ItemContractImpl
	}

	NoAuthCtx struct {
		state.BaseTxCtx
	}
)

var (
	_ common.TxCtxInterface                               = (*NoAuthCtx)(nil)
	_ ccpb.GenericServiceInterface[common.TxCtxInterface] = (*NoAuthContract)(nil)
)
