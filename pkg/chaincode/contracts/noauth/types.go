package noauth

import (
	"github.com/nova38/saacs/pkg/chaincode/common"
	"github.com/nova38/saacs/pkg/chaincode/context"
	"github.com/nova38/saacs/pkg/chaincode/contracts/base"
	ccpb "github.com/nova38/saacs/pkg/chaincode/gen/chaincode/common"
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
