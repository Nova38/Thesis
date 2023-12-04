package noauth

import (
	common "github.com/nova38/thesis/packages/saacs/auth/common"
	"github.com/nova38/thesis/packages/saacs/auth/state"
	"github.com/nova38/thesis/packages/saacs/contracts"
	ccpb "github.com/nova38/thesis/packages/saacs/gen/chaincode/common"
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
