package identity

import (
	"github.com/nova38/saacs/apps/saacs-cc/common"
	"github.com/nova38/saacs/apps/saacs-cc/context"
	contracts "github.com/nova38/saacs/apps/saacs-cc/contracts/base"
	authpb "github.com/nova38/saacs/pkg/saacs-protos/auth/v1"
	ccpb "github.com/nova38/saacs/pkg/saacs-protos/chaincode/common"
)

// ═════════════════════════════════════════════
// Transaction Context
// ═════════════════════════════════════════════

type (
	TxCtx struct {
		context.BaseTxCtx
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
