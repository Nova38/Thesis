package context

import (
	"github.com/nova38/thesis/lib/go/fabric/rbac"
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v2"
)

type CCBioCtxInterface interface {
	rbac.AuthTxCtxInterface

	// Extract Transaction Objects
	ExtractFromPayload(req interface{}) (err error)
}

// Extractor interfaces

type SpecimenHolder interface {
	GetSpecimen() *pb.Specimen
}
type SuggestionsIDHolder interface {
	GetId() *pb.Specimen_Id
}

type SuggestedUpdateHolder interface {
	GetSuggestedUpdate() *pb.SuggestedUpdate
}
type SuggestedUpdateIDHolder interface {
	GetId() *pb.SuggestedUpdate_Id
}
