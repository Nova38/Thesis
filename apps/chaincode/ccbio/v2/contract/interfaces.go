package contract

import (
	"github.com/nova38/thesis/lib/go/fabric/rbac"
)

type CCBioCtxInterface interface {
	rbac.AuthTxCtxInterface

	// Extract Transaction Objects
	ExtractFromPayload(req interface{}) (err error)
}

// Extractor interfaces
