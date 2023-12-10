package embedded

import (
	"github.com/nova38/thesis/packages/saacs/common"
	"github.com/nova38/thesis/packages/saacs/contracts/base"
	"github.com/nova38/thesis/packages/saacs/state"

	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	cc "github.com/nova38/thesis/packages/saacs/gen/chaincode/common"
)

type (
	TxCtx struct {
		state.BaseTxCtx
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
