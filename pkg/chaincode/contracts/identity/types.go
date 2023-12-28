package identity

import (
	"github.com/nova38/saacs/pkg/chaincode/common"
	contracts "github.com/nova38/saacs/pkg/chaincode/contracts/base"
	"github.com/nova38/saacs/pkg/chaincode/state"

	authpb "github.com/nova38/saacs/pkg/chaincode/gen/auth/v1"
	ccpb "github.com/nova38/saacs/pkg/chaincode/gen/chaincode/common"
)

// ═════════════════════════════════════════════
// Transaction Context
// ═════════════════════════════════════════════

type (
	TxCtx struct {
		state.BaseTxCtx
		CollectionMemberships map[string]*authpb.UserDirectMembership
	}
	IdentiyContract struct {
		contracts.ItemContractImpl
	}
)

var (
	_ common.TxCtxInterface                               = (*TxCtx)(nil)
	_ ccpb.GenericServiceInterface[common.TxCtxInterface] = (*IdentiyContract)(nil)
)
