package identity

import (
	common "github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/contracts"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	ccpb "github.com/nova38/thesis/lib/go/gen/chaincode/auth/common"
)

// ═════════════════════════════════════════════
// Transaction Context
// ═════════════════════════════════════════════

type (
	IdentityTxCtx struct {
		state.BaseTxCtx
	}
	IdentiyContract struct {
		contracts.ItemContractImpl
	}
)

var (
	_ common.TxCtxInterface                               = (*IdentityTxCtx)(nil)
	_ ccpb.GenericServiceInterface[common.TxCtxInterface] = (*IdentiyContract)(nil)
)
