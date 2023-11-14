package noauth

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
)

// =============================================
// Transaction Context
// =============================================
var _ state.TxCtxInterface = (*NoAuthTxCtx)(nil)

type NoAuthTxCtx struct {
	state.BaseTxCtx
}

// =============================================
// Contract
// =============================================
