package context

import (
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v2"
	"github.com/samber/oops"
)

// ExtractSpecimenId extracts the specimen ID from the given request. It supports requests that implement
// the SpecimenHolder or SuggestionsIDHolder interfaces. Returns an error if the ID cannot be extracted.
func ExtractSpecimenId(req interface{}) (id *pb.Specimen_Id, err error) {

	if req == nil {
		return nil, oops.Errorf("Interface is a nil pointer")
	}

	//  Extract the

	switch v := req.(type) {
	case SpecimenHolder:
		specimen := v.GetSpecimen()

		if specimen == nil {
			return nil, oops.Errorf("Specimen is not set")
		}

		id = specimen.GetId()

	case SuggestionsIDHolder:
		id = v.GetId()

	default:
		return nil, oops.Errorf("Req doesn't contain a specimen id in its type, %T", req)
	}

	if id == nil {
		return nil, oops.Errorf("ID is nil")
	}

	// Check that the id and the collection are valid

	return id, nil

}
