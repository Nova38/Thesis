package roles

import (
	common "github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	"github.com/nova38/thesis/lib/go/fabric/contracts"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	cc "github.com/nova38/thesis/lib/go/gen/chaincode/auth/common"
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
