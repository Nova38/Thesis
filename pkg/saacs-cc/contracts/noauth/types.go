package noauth

import (
	"github.com/nova38/saacs/apps/saacs-cc/common"
	"github.com/nova38/saacs/apps/saacs-cc/context"
	"github.com/nova38/saacs/apps/saacs-cc/contracts/base"
	ccpb "github.com/nova38/saacs/pkg/saacs-protos/chaincode/common"
)

type (
	Contract struct {
		base.ItemContractImpl
	}

	Ctx struct {
		context.BaseTxCtx
	}
)

var (
	_ common.TxCtxInterface                               = (*Ctx)(nil)
	_ ccpb.GenericServiceInterface[common.TxCtxInterface] = (*Contract)(nil)
)
