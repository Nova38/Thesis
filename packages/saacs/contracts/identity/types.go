package identity

import (
	"github.com/nova38/thesis/packages/saacs/common"
	contracts "github.com/nova38/thesis/packages/saacs/contracts/base"
	"github.com/nova38/thesis/packages/saacs/state"

	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	ccpb "github.com/nova38/thesis/packages/saacs/gen/chaincode/common"
)

// ═════════════════════════════════════════════
// Transaction Context
// ═════════════════════════════════════════════

type (
	TxCtx struct {
		state.BaseTxCtx
		CollectionMemberships map[string]*authpb.UserMembership
	}
	IdentiyContract struct {
		contracts.ItemContractImpl
	}
)

var (
	_ common.TxCtxInterface                               = (*TxCtx)(nil)
	_ ccpb.GenericServiceInterface[common.TxCtxInterface] = (*IdentiyContract)(nil)
)
