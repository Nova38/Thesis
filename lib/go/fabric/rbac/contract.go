package rbac

import (
	_ "strings"
	// "github.com/rs/zerolog/log"
	_ "github.com/samber/lo"
	_ "github.com/samber/oops"

	_ "google.golang.org/protobuf/types/known/timestamppb"

	// "github.com/hyperledger-labs/cckit/identity"

	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
)

type AuthContractImpl struct{}
