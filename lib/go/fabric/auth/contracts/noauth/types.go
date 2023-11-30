package noauth

import (
	common "github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/contracts"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
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
