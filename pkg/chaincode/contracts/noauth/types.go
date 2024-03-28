package noauth

import (
	ccpb "github.com/nova38/saacs/libs/saacs-protos-go/chaincode/common"
	"github.com/nova38/saacs/pkg/chaincode/common"
	"github.com/nova38/saacs/pkg/chaincode/context"
	"github.com/nova38/saacs/pkg/chaincode/contracts/base"
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
