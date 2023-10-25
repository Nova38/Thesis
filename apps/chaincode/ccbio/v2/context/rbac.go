package context

import (
	"github.com/nova38/thesis/lib/go/fabric/rbac"
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v2"
)

type TransactionObjects struct {
	Specimen *pb.Specimen
}

type CCBioCtx struct {
	rbac.AuthTxCtx

	TransactionObjects
}
