package contract

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/nova38/thesis/lib/go/fabric/rbac"
	"github.com/nova38/thesis/lib/go/fabric/state"
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v2"
	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
	"github.com/samber/oops"
)

var _ rbac.AuthTxCtxInterface = (*CCBioTxCtx)(nil)

type (
	TransactionObjects struct {
		Specimen  *pb.Specimen
		Suggested *pb.SuggestedUpdate
		Mask      *fieldmaskpb.FieldMask
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
	// userHolder interface {
	// 	GetUser() *rbac_pb.User
	// }
	// userIdHolder interface {
	// 	GetUserId() *rbac_pb.User_Id
	// }
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

	updatedMaskHolder interface {
		GetMask() *fieldmaskpb.FieldMask
	}

	pageHolder interface {
		GetPageSize() int32
	}
)

func (ctx *CCBioTxCtx) HandelBefore() (err error) {
	return ctx.BaseHandelBefore()
}

func (ctx *CCBioTxCtx) InitViaReq(req proto.Message) (err error) {
	if err = ctx.Validate(req); err != nil {
		return oops.Wrap(err)
	}

	// Extract the transaction items (collection id and specimen id)
	if err = ctx.ExtractTransactionItems(req); err != nil {
		return oops.Wrap(err)
	}

	if err = ctx.FetchTransactionItems(); err != nil {
		return oops.Wrap(err)
	}

	if err = ctx.IsAuthorized(); err != nil {
		return oops.Wrap(err)
	}
	return nil
}

func (ctx CCBioTxCtx) FetchTransactionItems() (err error) {
	if ctx.Suggested.Id != nil {
		if err = state.Get(ctx, ctx.Suggested); err != nil {
			return oops.With("suggested_id", ctx.Suggested.Id).Wrap(err)
		}
	}

	if ctx.Specimen.Id != nil {
		ctx.Collection.Id = &rbac_pb.Collection_Id{CollectionId: ctx.Specimen.Id.CollectionId}
		if err = state.Get(ctx, ctx.Specimen); err != nil {
			return oops.With("specimen_id", ctx.Specimen.Id).Wrap(err)
		}

	}

	if ctx.Collection.Id != nil {
		if err = state.Get(ctx, ctx.Collection); err != nil {
			return oops.With("collection_id", ctx.Collection.Id).Wrap(err)
		}
	}

	return nil
}

func (ctx *CCBioTxCtx) ExtractTransactionItems(req interface{}) (err error) {
	if ctx.Suggested == nil {
		ctx.Suggested = &pb.SuggestedUpdate{}
	}
	if ctx.Specimen == nil {
		ctx.Specimen = &pb.Specimen{}
	}
	if ctx.Collection == nil {
		ctx.Collection = &rbac_pb.Collection{}
	}

	switch v := req.(type) {

	case specimenHolder:
		if v.GetSpecimen() == nil || v.GetSpecimen().GetId() == nil {
			return oops.Errorf("Specimen is not set")
		}
		ctx.Specimen.Id = v.GetSpecimen().GetId()

	case specimenIdHolder:
		if v.GetId() == nil {
			return oops.Errorf("Collection is not set")
		}
		ctx.Specimen.Id = v.GetId()

	case suggestedUpdateHolder:
		s := v.GetSuggestedUpdate()
		if s == nil || s.GetId() == nil || s.GetId().GetSpecimenId() == nil {
			return oops.Errorf("SuggestedUpdate is not set")
		}

		ctx.Specimen.Id = s.GetId().GetSpecimenId()

	case suggestedUpdateIdHolder:
		if v.GetId() == nil || v.GetId().GetSpecimenId() == nil {
			return oops.Errorf("SuggestedUpdateId is not set")
		}

		ctx.Specimen.Id = v.GetId().GetSpecimenId()

	case collectionHolder:
		if v.GetCollection() == nil || v.GetCollection().GetId() == nil {
			return oops.Errorf("Collection is not set")
		}
		ctx.Collection.Id = v.GetCollection().GetId()

	case collectionIdHolder:
		if v.GetCollectionId() == nil {
			return oops.Errorf("Collection is not set")
		}
		ctx.Collection.Id = v.GetCollectionId()

	}

	// Extract the page size if it exists
	if page, ok := req.(pageHolder); ok {
		if page.GetPageSize() < 0 {
			ctx.PageSize = page.GetPageSize()
		}
	}

	// Extract the mask if it exists
	if mask, ok := req.(updatedMaskHolder); ok {
		ctx.Mask = mask.GetMask()
		ctx.SetOperationsPaths(ctx.Mask)
	}

	if ctx.Suggested.Id != nil {
		ctx.Specimen.Id = ctx.Suggested.Id.SpecimenId
	}
	if ctx.Specimen.Id != nil {
		ctx.Collection.Id = &rbac_pb.Collection_Id{CollectionId: ctx.Specimen.Id.CollectionId}
	}

	return nil
}

// TODO: Fix the complexity
// nolint:cyclop
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

// ExtractCollectionId extracts the collection id from the request.
// nolint:cyclop
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
