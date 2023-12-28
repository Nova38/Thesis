package roles

import (
	"github.com/nova38/thesis/pkg/saacs/common"
	"github.com/nova38/thesis/pkg/saacs/contracts/base"
	"github.com/nova38/thesis/pkg/saacs/state"

	authpb "github.com/nova38/thesis/pkg/saacs/gen/auth/v1"
	cc "github.com/nova38/thesis/pkg/saacs/gen/chaincode/common"
)

type (
	TxCtx struct {
		state.BaseTxCtx
		UserRoles map[string][]*authpb.Role
	}

	RoleContract struct {
		base.ItemContractImpl
	}
)

// type checking
var (
	_ common.TxCtxInterface = (*TxCtx)(
		nil,
	) // _ contracts.ItemContractInterface = (*IdentiyContract)(nil)
	// see if ItemContractImpl implements the interface GenericServiceInterface
	_ cc.GenericServiceInterface[common.TxCtxInterface] = (*RoleContract)(nil)
)
