package noauth

import (
	common "github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	"github.com/nova38/thesis/lib/go/fabric/contracts"
	ccpb "github.com/nova38/thesis/lib/go/gen/chaincode/auth/common"
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
