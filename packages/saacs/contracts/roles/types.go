package roles

import (
	"github.com/nova38/thesis/packages/fabric/auth/state"
	"github.com/nova38/thesis/packages/fabric/contracts"
	common "github.com/nova38/thesis/packages/saacs/auth/common"
	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	cc "github.com/nova38/thesis/packages/saacs/genchaincode/auth/common"
)

type (
	RolesTxCtx struct {
		state.BaseTxCtx
		UserRoles map[string][]*authpb.Role
	}

	RoleContract struct {
		contracts.ItemContractImpl
	}
)

// type checking
var (
	_ common.TxCtxInterface = (*RolesTxCtx)(
		nil,
	) // _ contracts.ItemContractInterface = (*IdentiyContract)(nil)
	// see if ItemContractImpl implements the interface GenericServiceInterface
	_ cc.GenericServiceInterface[common.TxCtxInterface] = (*RoleContract)(nil)
)
