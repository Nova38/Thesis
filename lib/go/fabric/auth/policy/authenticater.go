package policy

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
)

func Authenticate(ctx state.TxCtxInterface, ops []*authpb.Operation) (auth bool, err error) {
	return true, nil
}

func BuildRadix()
