package identity

import (
	"github.com/nova38/saacs/apps/chaincode/common"
	"github.com/nova38/saacs/apps/chaincode/context"
	contracts "github.com/nova38/saacs/apps/chaincode/contracts/base"
	authpb "github.com/nova38/saacs/libs/saacs-protos-go/auth/v1"
	ccpb "github.com/nova38/saacs/libs/saacs-protos-go/chaincode/common"
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
