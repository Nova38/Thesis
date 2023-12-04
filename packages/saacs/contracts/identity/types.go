package identity

import (
	common "github.com/nova38/thesis/packages/saacs/auth/common"
	"github.com/nova38/thesis/packages/saacs/auth/state"
	"github.com/nova38/thesis/packages/saacs/contracts"
	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	ccpb "github.com/nova38/thesis/packages/saacs/gen/chaincode/common"
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
