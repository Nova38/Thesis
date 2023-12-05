package noauth

import (
	"github.com/nova38/thesis/packages/saacs/common"
	"github.com/nova38/thesis/packages/saacs/state"

	"github.com/nova38/thesis/packages/saacs/contracts/base"
	ccpb "github.com/nova38/thesis/packages/saacs/gen/chaincode/common"
)

type (
	NoAuthContract struct {
		base.ItemContractImpl
	}

	NoAuthCtx struct {
		state.BaseTxCtx
	}
)

var (
	_ common.TxCtxInterface                               = (*NoAuthCtx)(nil)
	_ ccpb.GenericServiceInterface[common.TxCtxInterface] = (*NoAuthContract)(nil)
)
