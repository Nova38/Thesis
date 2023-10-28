package contract

import (
	"github.com/nova38/thesis/lib/go/fabric/rbac"
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v2"
	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
	"github.com/samber/oops"
)

var _ rbac.AuthTxCtxInterface = (*CCBioTxCtx)(nil)

type (
	TransactionObjects struct {
		Specimen *pb.Specimen
	}

	CCBioTxCtx struct {
		*rbac.TxCtx

		TransactionObjects
	}
)

type (
	collectionHolder interface {
		GetCollection() *rbac_pb.Collection
	}
	collectionIdHolder interface {
		GetCollectionId() *rbac_pb.Collection_Id
	}
	userHolder interface {
		GetUser() *rbac_pb.User
	}
	userIdHolder interface {
		GetUserId() *rbac_pb.User_Id
	}
	specimenHolder interface {
		GetSpecimen() *pb.Specimen
	}
	specimenIdHolder interface {
		GetId() *pb.Specimen_Id
	}
	suggestedUpdateHolder interface {
		GetSuggestedUpdate() *pb.SuggestedUpdate
	}
	suggestedUpdateIdHolder interface {
		GetId() *pb.SuggestedUpdate_Id
	}
)

func ExtractSpecimenId(req interface{}) (id *pb.Specimen_Id, err error) {
	if req == nil {
		return nil, oops.Errorf("Interface is a nil pointer")
	}

	//  Extract the

	switch v := req.(type) {

	case specimenHolder:
		if v.GetSpecimen() == nil || v.GetSpecimen().GetId() == nil {
			return nil, oops.Errorf("Specimen is not set")
		}
		id = v.GetSpecimen().GetId()

	case specimenIdHolder:
		if v.GetId() == nil {
			return nil, oops.Errorf("Collection is not set")
		}
		id = v.GetId()

	case suggestedUpdateHolder:
		s := v.GetSuggestedUpdate()
		if s == nil || s.GetId() == nil || s.GetId().GetSpecimenId() == nil {
			return nil, oops.Errorf("SuggestedUpdate is not set")
		}

		id = s.GetId().GetSpecimenId()

	case suggestedUpdateIdHolder:
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

func extractCollectionId(req interface{}) (id *rbac_pb.Collection_Id, err error) {
	if req == nil {
		return nil, oops.Errorf("Interface is a nil pointer")
	}

	//  Extract the

	switch v := req.(type) {
	case collectionHolder:
		if v.GetCollection() == nil {
			return nil, oops.Errorf("Collection is not set")
		}
		id = v.GetCollection().GetId()
	case collectionIdHolder:
		id = v.GetCollectionId()
	case specimenHolder:
		if v.GetSpecimen() == nil || v.GetSpecimen().GetId() == nil {
			return nil, oops.Errorf("Specimen is not set")
		}
		id = &rbac_pb.Collection_Id{CollectionId: v.GetSpecimen().GetId().CollectionId}

	case specimenIdHolder:
		if v.GetId() == nil {
			return nil, oops.Errorf("Collection is not set")
		}
		id = &rbac_pb.Collection_Id{CollectionId: v.GetId().CollectionId}

	case suggestedUpdateHolder:
		s := v.GetSuggestedUpdate()
		if s == nil || s.GetId() == nil || s.GetId().GetSpecimenId() == nil {
			return nil, oops.Errorf("SuggestedUpdate is not set")
		}

		id = &rbac_pb.Collection_Id{CollectionId: s.GetId().GetSpecimenId().CollectionId}

	case suggestedUpdateIdHolder:
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
