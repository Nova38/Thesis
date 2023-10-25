package rbac

import (
	_ "strings"
	// "github.com/rs/zerolog/log"
	_ "github.com/samber/lo"
	_ "github.com/samber/oops"

	_ "google.golang.org/protobuf/types/known/timestamppb"

	// "github.com/hyperledger-labs/cckit/identity"

	cc "github.com/nova38/thesis/lib/go/gen/chaincode/rbac/schema/v1"
)

// Check if AuthContractImpl implements AuthServiceInterface
var _ cc.AuthServiceInterface = (*AuthContractImpl)(nil)

type AuthContractImpl struct {
	cc.AuthServiceBase
}
