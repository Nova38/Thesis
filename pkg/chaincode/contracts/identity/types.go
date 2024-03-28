package identity

import (
	authpb "github.com/nova38/saacs/libs/saacs-protos-go/auth/v1"
	ccpb "github.com/nova38/saacs/libs/saacs-protos-go/chaincode/common"
	"github.com/nova38/saacs/pkg/chaincode/common"
	"github.com/nova38/saacs/pkg/chaincode/context"
	contracts "github.com/nova38/saacs/pkg/chaincode/contracts/base"
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
