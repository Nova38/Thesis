package context

import (
	"github.com/nova38/thesis/lib/go/fabric/rbac"
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v2"
	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
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
		if v.GetSpecimen() == nil || v.GetSpecimen().GetId() == nil {
			return nil, oops.Errorf("Specimen is not set")
		}
		id = v.GetSpecimen().GetId()

	case SuggestionsIdHolder:
		if v.GetId() == nil {
			return nil, oops.Errorf("Collection is not set")
		}
		id = v.GetId()

	case SuggestedUpdateHolder:
		s := v.GetSuggestedUpdate()
		if s == nil || s.GetId() == nil || s.GetId().GetSpecimenId() == nil {
			return nil, oops.Errorf("SuggestedUpdate is not set")
		}

		id = s.GetId().GetSpecimenId()

	case SuggestedUpdateIDHolder:
		if v.GetId() == nil || v.GetId().GetSpecimenId() == nil {
			return nil, oops.Errorf("SuggestedUpdateId is not set")
		}

		id = v.GetId().GetSpecimenId()

	default:
		return nil, oops.Errorf("Req doesn't contain a collection id in its type, %T", req)
	}

	if id == nil {
		return nil, oops.Errorf("ID is nil")
	}

	if id.CollectionId == "" || id.Id == "" {
		return nil, oops.Errorf("ID is empty")
	}
	// Check that the id and the collection are valid

	return id, nil
}

func ExtractCollectionId(req interface{}) (id *rbac_pb.Collection_Id, err error) {
	if req == nil {
		return nil, oops.Errorf("Interface is a nil pointer")
	}

	//  Extract the

	switch v := req.(type) {
	case rbac.CollectionHolder:
		if v.GetCollection() == nil {
			return nil, oops.Errorf("Collection is not set")
		}
		id = v.GetCollection().GetId()
	case rbac.CollectionIdHolder:
		id = v.GetCollectionId()
	case SpecimenHolder:
		if v.GetSpecimen() == nil || v.GetSpecimen().GetId() == nil {
			return nil, oops.Errorf("Specimen is not set")
		}
		id = &rbac_pb.Collection_Id{CollectionId: v.GetSpecimen().GetId().CollectionId}

	case SuggestionsIdHolder:
		if v.GetId() == nil {
			return nil, oops.Errorf("Collection is not set")
		}
		id = &rbac_pb.Collection_Id{CollectionId: v.GetId().CollectionId}

	case SuggestedUpdateHolder:
		s := v.GetSuggestedUpdate()
		if s == nil || s.GetId() == nil || s.GetId().GetSpecimenId() == nil {
			return nil, oops.Errorf("SuggestedUpdate is not set")
		}

		id = &rbac_pb.Collection_Id{CollectionId: s.GetId().GetSpecimenId().CollectionId}

	case SuggestedUpdateIDHolder:
		if v.GetId() == nil || v.GetId().GetSpecimenId() == nil {
			return nil, oops.Errorf("SuggestedUpdateId is not set")
		}

		id = &rbac_pb.Collection_Id{CollectionId: v.GetId().GetSpecimenId().CollectionId}

	default:
		return nil, oops.Errorf("Req doesn't contain a collection id in its type, %T", req)
	}

	if id == nil {
		return nil, oops.Errorf("ID is nil")
	}

	if id.CollectionId == "" {
		return nil, oops.Errorf("ID is empty")
	}
	// Check that the id and the collection are valid

	return id, nil
}
