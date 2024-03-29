package roles

import (
	"github.com/nova38/saacs/apps/chaincode/common"
	"github.com/nova38/saacs/apps/chaincode/context"
	"github.com/nova38/saacs/apps/chaincode/contracts/base"
	authpb "github.com/nova38/saacs/libs/saacs-protos-go/auth/v1"
	cc "github.com/nova38/saacs/libs/saacs-protos-go/chaincode/common"
)

type (
	TxCtx struct {
		context.BaseTxCtx
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
	) // _ contracts.ItemContractInterface = (*IdentityContract)(nil)
	// see if ItemContractImpl implements the interface GenericServiceInterface
	_ cc.GenericServiceInterface[common.TxCtxInterface] = (*RoleContract)(nil)
)
