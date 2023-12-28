package noauth

import (
	"github.com/nova38/thesis/pkg/saacs/common"
	"github.com/nova38/thesis/pkg/saacs/state"

	"github.com/nova38/thesis/pkg/saacs/contracts/base"
	ccpb "github.com/nova38/thesis/pkg/saacs/gen/chaincode/common"
)

type (
	Contract struct {
		base.ItemContractImpl
	}

	Ctx struct {
		state.BaseTxCtx
	}
)

var (
	_ common.TxCtxInterface                               = (*Ctx)(nil)
	_ ccpb.GenericServiceInterface[common.TxCtxInterface] = (*Contract)(nil)
)
