package embedded

import (
	"github.com/nova38/saacs/apps/saacs-cc/common"
	"github.com/nova38/saacs/apps/saacs-cc/context"
	"github.com/nova38/saacs/apps/saacs-cc/contracts/base"
	authpb "github.com/nova38/saacs/libs/saacs-protos-go/auth/v1"
	cc "github.com/nova38/saacs/libs/saacs-protos-go/chaincode/common"
)

type (
	TxCtx struct {
		context.BaseTxCtx
		UserRoles map[string][]*authpb.Role
	}

	EmbededContract struct {
		base.ItemContractImpl
	}
)

// type checking
var (
	_ common.TxCtxInterface = (*TxCtx)(
		nil,
	) // _ contracts.ItemContractInterface = (*IdentiyContract)(nil)
	// see if ItemContractImpl implements the interface GenericServiceInterface
	_ cc.GenericServiceInterface[common.TxCtxInterface] = (*EmbededContract)(nil)
)
