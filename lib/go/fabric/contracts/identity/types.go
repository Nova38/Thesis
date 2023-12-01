package identity

import (
	common "github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	"github.com/nova38/thesis/lib/go/fabric/contracts"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	ccpb "github.com/nova38/thesis/lib/go/gen/chaincode/auth/common"
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
