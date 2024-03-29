package noauth

import (
	"github.com/nova38/saacs/apps/chaincode/common"
	"github.com/nova38/saacs/apps/chaincode/context"
	"github.com/nova38/saacs/apps/chaincode/contracts/base"
	ccpb "github.com/nova38/saacs/libs/saacs-protos-go/chaincode/common"
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
