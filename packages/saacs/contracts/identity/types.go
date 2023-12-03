package identity

import (
	"github.com/nova38/thesis/packages/fabric/auth/state"
	"github.com/nova38/thesis/packages/fabric/contracts"
	common "github.com/nova38/thesis/packages/saacs/auth/common"
	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	ccpb "github.com/nova38/thesis/packages/saacs/genchaincode/auth/common"
)

// ═════════════════════════════════════════════
// Transaction Context
// ═════════════════════════════════════════════

type (
	IdentityTxCtx struct {
		state.BaseTxCtx
		CollectionMemberships map[string]*authpb.UserMembership
	}
	IdentiyContract struct {
		contracts.ItemContractImpl
	}
)

var (
	_ common.TxCtxInterface                               = (*IdentityTxCtx)(nil)
	_ ccpb.GenericServiceInterface[common.TxCtxInterface] = (*IdentiyContract)(nil)
)
