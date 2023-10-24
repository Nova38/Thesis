package context

import (
	"github.com/nova38/thesis/lib/go/fabric/rbac"
	"github.com/nova38/thesis/lib/go/fabric/state"
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v2"
	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
	"github.com/samber/oops"
)

type TransactionObjects struct {
	Specimen *pb.Specimen
}

type CCBioCtx struct {
	rbac.AuthTxCtx

	TransactionObjects
}

func (ctx *CCBioCtx) InitFromPayload(req interface{}) (err error) {

	if req == nil {
		return oops.Errorf("Interface is a nil pointer")
	}

	ctx.Specimen = &pb.Specimen{}
	ctx.Collection = &rbac_pb.Collection{}

	// Todo: should we require that the specimen and the roles are not already set

	//  Extract the
	id := &pb.Specimen_Id{}

	switch v := req.(type) {
	case SpecimenHolder:
		specimen := v.GetSpecimen()

		if specimen == nil {
			return oops.Errorf("Specimen is not set")
		}

		id = specimen.GetId()

	case SuggestionsIDHolder:
		id = v.GetId()
	}

	if id == nil {
		return oops.Errorf("ID is nil")
	}

	// Check that the id and the collection are valid

	return nil

}
