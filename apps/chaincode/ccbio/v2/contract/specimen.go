package contracts

import (
	"fmt"

	state "github.com/nova38/thesis/lib/go/fabric/state"
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v2"

	rbac "github.com/nova38/thesis/lib/go/fabric/rbac"

	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"

	"log/slog"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

// SpecimenContract contract for handling BasicAssets
type SpecimenContract struct {
	contractapi.Contract
}

func BuildSpecimenContract() *SpecimenContract {
	return &SpecimenContract{
		Contract: contractapi.Contract{
			Name:              "ccbio.Specimen",
			BeforeTransaction: rbac.AuthTxCtx.HandelBefore,
			Info: metadata.InfoMetadata{
				Description: "",
				Title:       "Biochain Chaincode",
				Contact: &metadata.ContactMetadata{
					Name:  "Thomas Atkins",
					URL:   "https://biochain.iitc.ku.edu",
					Email: "tom@ku.edu",
				},
				License: &metadata.LicenseMetadata{
					Name: "MIT",
					URL:  "https://example.com",
				},
				Version: "latest",
			},
		},
	}
}

// ────────────────────────────────────────────────────────────

func (s *SpecimenContract) GetEvaluateTransactions() []string {
	return []string{
		"GetSpecimen",
		"GetSpecimenList",
		"GetSpecimenByCollection",
		"GetSpecimenHistory",
	}
}

//func (s *BioContract) GetBeforeTransaction() interface{} {
//	return HandelBefore
//}

// ────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────
// Query Functions
// ────────────────────────────────────────────────────────────

// -------------------------
// Specimen
// -------------------------

func (s *SpecimenContract) GetSpecimen(ctx rbac.AuthTxCtx, req *pb.GetSpecimenRequest) (*pb.Specimen, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}
	id := req.GetId()

	specimen := &pb.Specimen{
		Id: &pb.Specimen_Id{
			CollectionId: id.GetCollectionId(),
			Id:           id.GetId(),
		},
	}

	if err := state.GetState(&ctx, specimen); err != nil {
		return nil, err
	}

	return specimen, nil
}

func (s *SpecimenContract) GetSpecimenList(ctx rbac.AuthTxCtx) ([]*pb.Specimen, error) {
	return state.GetFullStateList(&ctx, &pb.Specimen{})
}

func (s *SpecimenContract) GetSpecimenByCollection(ctx rbac.AuthTxCtx, req *pb.GetSpecimenByCollectionRequest) ([]*pb.Specimen, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}
	id := req.GetId()

	// Check if collection exists
	colExists, err := state.ObjExists(&ctx, &rbac_pb.Collection{
		Id: &rbac_pb.Collection_Id{
			CollectionId: id.GetCollectionId(),
		},
	})
	if err != nil {
		slog.Error("GetSpecimenByCollection failed to check if collection exists", "err", err)
		return nil, err
	}

	if !colExists {
		return nil, &state.KeyNotFoundError{
			Key:       id.GetCollectionId(),
			Namespace: "Collection",
			MSG:       "Collection not found",
		}
	}

	specimen := &pb.Specimen{
		Id: &pb.Specimen_Id{
			CollectionId: id.GetCollectionId(),
			Id:           "",
		},
	}
	slog.Info("SpecimenId", slog.Any("specimen", specimen))
	return state.GetPartialKeyList(&ctx, specimen, 1)
}

// GetSpecimenHistory returns the history of the specimen
//   - If the request includes hidden txs:
//   - The history will include hidden txs
//   - Requires the user to be registered
//   - Requires the user to be a have a role authorized to see hidden txs
//   - If the request does not include hidden txs, then the history will not include hidden txs
func (s *SpecimenContract) GetSpecimenHistory(ctx rbac.AuthTxCtx, req *pb.GetSpecimenHistoryRequest) (*pb.Specimen_History, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}
	id := req.GetId()

	specimen := &pb.Specimen{
		Id: &pb.Specimen_Id{
			CollectionId: id.GetCollectionId(),
			Id:           id.GetId(),
		},
	}
	err := state.GetState(&ctx, specimen)
	if err != nil {
		return nil, errors.Wrap(err, "GetSpecimenHistory failed to get specimen")
	}

	slog.Info("GetSpecimenHistory", "specimen", specimen)

	history := &pb.Specimen_History{
		Id:      id,
		Entries: []*pb.Specimen_History_Entry{},
	}

	rawHistory, err := state.GetStateHistory(&ctx, specimen)
	if err != nil {
		return nil, err
	}

	// See if they have requested to see the hidden txs
	if req.GetIncludeHidden() {
		// Specify the actions to authorize
		err := ctx.SetAction(&rbac_pb.Action{
			Level:   pb.Action_LEVEL_HIDE_TX,
			Domains: []pb.Action_Domain{},
		})
		if err != nil {
			slog.Error("SetAction failed", "err", err)
			return nil, errors.Wrap(err, "SetAction failed")
		}

		// Authorize the request and return error if not authorized
		valid, err := ctx.Authorize()
		if err != nil {
			slog.Error("Authorize user to create a specimen failed", "err", err)
			return nil, errors.Wrap(err, "Authorize user update membership failed")
		}
		if !valid {
			action, err := ctx.GetAction()
			if err != nil {
				slog.Error("GetAction failed", "err", err)
				return nil, errors.Wrap(err, "GetAction failed")
			}
			slog.Error("Not authorized", "actions", action)
			return nil, errors.New("Not authorized")
		}

	}

	for _, h := range rawHistory.Entries {
		entry := &pb.Specimen_History_Entry{}

		// check to see if the entry is hidden by seeing if it is a key in the hidden txs map
		if _, ok := specimen.HiddenTxs[h.TxId]; ok {
			entry.IsHidden = true

			if !req.GetIncludeHidden() {
				entry.State = nil
			}
		}

		history.Entries = append(history.Entries, entry)

	}
	return history, nil
}

// -------------------------
// Specimen Suggestions
// -------------------------

// GetSuggestedUpdate returns the suggested update with the given ID
func (s *SpecimenContract) GetSuggestedUpdate(ctx rbac.AuthTxCtx, req *pb.GetSuggestedUpdateRequest) (*pb.SuggestedUpdate, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}
	id := req.GetId()

	suggestion := &pb.SuggestedUpdate{
		Id: &pb.SuggestedUpdate_Id{
			Id: id.GetId(),
			SpecimenId: &pb.Specimen_Id{
				CollectionId: id.GetSpecimenId().GetCollectionId(),
				Id:           id.GetSpecimenId().GetId(),
			},
		},
	}

	if err := state.GetState(&ctx, suggestion); err != nil {
		return nil, err
	}

	return suggestion, nil
}

func (s *SpecimenContract) GetSuggestedUpdateList(ctx rbac.AuthTxCtx) ([]*pb.SuggestedUpdate, error) {
	return state.GetFullStateList(ctx, &pb.SuggestedUpdate{})
}

func (s *SpecimenContract) GetSuggestedUpdateBySpecimen(ctx rbac.AuthTxCtx, req *pb.GetSuggestedUpdateBySpecimenRequest) ([]*pb.SuggestedUpdate, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}
	id := req.GetId()

	suggestion := &pb.SuggestedUpdate{
		Id: &pb.SuggestedUpdate_Id{
			SpecimenId: &pb.Specimen_Id{
				CollectionId: id.GetCollectionId(),
				Id:           id.GetId(),
			},
		},
	}

	list, err := state.GetPartialKeyList(ctx, suggestion, 2)
	if err != nil {
		return nil, errors.Wrap(err, "GetSuggestedUpdateBySpecimen failed")
	}

	return list, nil
}

func (s *SpecimenContract) GetSuggestedUpdateByCollectionRequest(ctx rbac.AuthTxCtx, req *pb.GetSuggestedUpdateByCollectionRequest) ([]*pb.SuggestedUpdate, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}
	id := req.GetId()

	suggestion := &pb.SuggestedUpdate{
		Id: &pb.SuggestedUpdate_Id{
			SpecimenId: &pb.Specimen_Id{
				CollectionId: id.GetCollectionId(),
			},
		},
	}

	list, err := state.GetPartialKeyList(ctx, suggestion, 1)
	if err != nil {
		return nil, errors.Wrap(err, "GetSuggestedUpdateByCollection failed")
	}

	return list, nil
}

// ────────────────────────────────────────────────────────────
// Invoke Functions
// ────────────────────────────────────────────────────────────

// --------------------------------------------------
// Specimen
// --------------------------------------------------

func (s *SpecimenContract) Stest(ctx rbac.AuthTxCtx, req *pb.STest) {
	fmt.Printf("STest: %+v\n", req)
}

// SpecimenCreate creates a new specimen in the collection
//   - Requires the user to be registered
//   - Requires the user to be a have a role authorized to create specimens
func (s *SpecimenContract) SpecimenCreate(ctx rbac.AuthTxCtx, req *pb.SpecimenCreateRequest) (*pb.Specimen, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	// -------------------------
	// Authorize the request
	// -------------------------

}

// SpecimenUpdate updates the specimen
// If any of the fields contain an empty last modified field that field will be ignored
func (s *SpecimenContract) SpecimenUpdate(ctx rbac.AuthTxCtx, req *pb.SpecimenUpdateRequest) (*pb.Specimen, error) {
	// Validate request
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	// Initialize the auth context via the request
	// This function should not try to authorize the request
	// as it needs to still add the domains to the actions list
	if err := ctx.InitViaRequest(req); err != nil {
		return nil, errors.Wrap(err, "SpecimenUpdate init via request failed")
	}

	// -------------------------
	// Get the current specimen
	// -------------------------

	currentSpecimen := &pb.Specimen{
		Id: &pb.Specimen_Id{
			CollectionId: req.GetSpecimen().GetId().GetCollectionId(),
			Id:           req.GetSpecimen().GetId().GetId(),
		},
	}

	// Get the specimen
	if err := state.GetState(&ctx, currentSpecimen); err != nil {
		return nil, errors.Wrap(err, "SpecimenUnHideTx failed to get specimen")
	}

	// -------------------------
	// Process the request
	// -------------------------
	// This value will be reused for all the fields and should not be modified
	updatedAt, err := ctx.MakeLastModified()
	if err != nil {
		return nil, errors.Wrap(err, "SpecimenUpdate failed to get current modified struct")
	}

	// Primary
	{
		modified, err := UpdateMessageField("primary", currentSpecimen.GetPrimary(), req.GetSpecimen().GetPrimary())
		if err != nil {
			return nil, errors.Wrap(err, "SpecimenUpdate failed to update primary")
		}
		if modified {
			ctx.AddActionDomain(pb.Action_DOMAIN_PRIMARY)
			currentSpecimen.Primary.LastModified = updatedAt
		}
	}

	// Secondary
	{
		modified, err := UpdateMessageField("secondary", currentSpecimen.GetSecondary(), req.GetSpecimen().GetSecondary())
		if err != nil {
			return nil, errors.Wrap(err, "SpecimenUpdate failed to update secondary")
		}
		if modified {
			ctx.AddActionDomain(pb.Action_DOMAIN_SECONDARY)
			currentSpecimen.Secondary.LastModified = updatedAt
		}
	}

	// Taxon
	{
		modified, err := UpdateMessageField("taxon", currentSpecimen.GetTaxon(), req.GetSpecimen().GetTaxon())
		if err != nil {
			return nil, errors.Wrap(err, "SpecimenUpdate failed to update taxon")
		}
		if modified {
			ctx.AddActionDomain(pb.Action_DOMAIN_TAXON)
			currentSpecimen.Taxon.LastModified = updatedAt
		}
	}

	// Georeference
	{
		modified, err := UpdateMessageField("georeference", currentSpecimen.GetGeoreference(), req.GetSpecimen().GetGeoreference())
		if err != nil {
			return nil, errors.Wrap(err, "SpecimenUpdate failed to update georeference")
		}
		if modified {
			ctx.AddActionDomain(pb.Action_DOMAIN_GEOREFERENCE)
			currentSpecimen.Georeference.LastModified = updatedAt
		}

	}

	// Images

	// String Properties
	// -------------------------

	// Check if the loans have changed
	if currentSpecimen.GetLoans() != req.GetSpecimen().GetLoans() {
		slog.Info("SpecimenUpdate", "loans", req.GetSpecimen().GetLoans())
		ctx.AddActionDomain(pb.Action_DOMAIN_LOANS)
		currentSpecimen.Loans = req.GetSpecimen().GetLoans()
	}

	// Check if the grants have changed
	if currentSpecimen.GetGrants() != req.GetSpecimen().GetGrants() {
		slog.Info("SpecimenUpdate", "grants", req.GetSpecimen().GetGrants())
		// actions.Domains = append(actions.Domains, pb.Action_DOMAIN_GRANTS)
		ctx.AddActionDomain(pb.Action_DOMAIN_GRANTS)
		currentSpecimen.Grants = req.GetSpecimen().GetGrants()
	}

	// Map Properties
	// -------------------------

	// Check if the images have changed

	// Get the current images
	currentImages := currentSpecimen.GetImages()
	// Get the new images
	newImages := req.GetSpecimen().GetImages()

	cKeys := lo.Keys(currentImages)
	nKeys := lo.Keys(newImages)

	if len(cKeys) == 0 && len(nKeys) == 0 {

		removed, added := lo.Difference(cKeys, nKeys)
		shared := lo.Union(cKeys, nKeys)

		updatedImages := lo.PickByKeys(newImages, added)
		// any updated?
		modified := (len(removed) > 0) || (len(added) > 0)

		for _, key := range shared {
			newItem := newImages[key]
			oldItem := currentImages[key]
			if !proto.Equal(ctx.MakeClearedLastModified(), newItem.GetLastModified()) {
				if !proto.Equal(oldItem, newItem) {
					updatedImages[key] = newItem
					modified = true
					// Update the last modified
					modified, ok := proto.Clone(updatedAt).(*pb.LastModified)
					if !ok {
						return nil, errors.New("Failed to clone updatedAt")
					}
					updatedImages[key].LastModified = modified
				}
				updatedImages[key] = oldItem
			} else {
				updatedImages[key] = oldItem
			}
		}
		currentSpecimen.Images = updatedImages

		if modified {
			ctx.AddActionDomain(pb.Action_DOMAIN_IMAGES)
		}

	} else {
		slog.Info("SpecimenUpdate", "images", "skipping => Both maps are empty")
	}

	// -------------------------
	// Authorize the request
	// -------------------------

	// Authorize the request and return error if not authorized
	valid, err := ctx.Authorize()
	if err != nil {
		slog.Error("Authorize user to update specimen failed", "err", err)
		return nil, errors.Wrap(err, "Authorize user update membership failed")
	}
	if !valid {
		action, err := ctx.GetAction()
		if err != nil {
			return nil, errors.Wrap(err, "SpecimenUpdate failed to get actions")
		}
		slog.Error("Not authorized", "actions", action)
		return nil, errors.New("Not authorized")
	}

	currentSpecimen.LastModified = updatedAt

	return currentSpecimen, state.UpdateState(ctx, currentSpecimen)
}

// SpecimenDelete removes the specimen from the world state
func (s *SpecimenContract) SpecimenDelete(ctx rbac.AuthTxCtx, req *pb.SpecimenDeleteRequest) (*pb.Specimen, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	// Extract the id

	id := req.GetId()
	specimen := &pb.Specimen{
		Id: &pb.Specimen_Id{
			CollectionId: id.GetCollectionId(),
			Id:           id.GetId(),
		},
	}
	err := state.GetState(&ctx, specimen)

	if err != nil {
		return nil, oops.In("SpecimenDelete").Code("StateNotFound").Wrap(err)
	}

	// -------------------------
	// Authorize the request
	// -------------------------
	// Set the collection
	ctx.SetCollection(&rbac_pb.Collection{
		Id: &rbac_pb.Collection_Id{CollectionId: specimen.Id.GetCollectionId()}},
	)
	ctx.SetOperation(&rbac_pb.Operations{
		Domain: rbac_pb.Operations_DOMAIN_OBJECT,
		Action: &rbac_pb.Operations_Action{Type: rbac_pb.Operations_Action_TYPE_DELETE},
	})

	// Authorize the request and return error if not authorized
	authorized, err := ctx.Authorize()
	if err != nil {
		slog.Error("Authorize user to delete specimen failed", "err", err)
		return nil, oops.In("SpecimenDelete").Wrap(err)
	}
	if !authorized {
		slog.Error("Not authorized To Delete", "User", &ctx.User)
		return nil, oops.In("SpecimenDelete").With("User", &ctx.User).Wrap(errors.New("Not authorized"))
	}

	// -------------------------
	// Process the request
	// -------------------------

	err = state.DeleteState(&ctx, specimen)

	if err != nil {
		return nil, oops.In("SpecimenDelete").Code("DeleteStateFailed").Wrap(err)
	}

	return specimen, nil
}

// SpecimenHideTx hides the tx from the specimen
//   - Adds the tx to the hidden txs
//   - Sets the last modified field
//   - Requires the user to be registered
//   - Requires the user to be a have a role authorized to hide txs
//   - Requires the tx to be in the specimen history
//   - Requires the tx to not already be hidden
//   - Requires the tx to not be the last modified tx
func (s *SpecimenContract) SpecimenHideTx(ctx rbac.AuthTxCtx, req *pb.SpecimenHideTxRequest) (*pb.Specimen, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	// -------------------------
	// Authorize the request
	// -------------------------

	// Initialize the auth context via the request, Should return error if not authorized
	if err := ctx.InitViaRequest(req); err != nil {
		return nil, errors.Wrap(err, "SpecimenCreate init via request failed")
	}

	// -------------------------
	// Process the request
	// -------------------------
	specimen := &pb.Specimen{
		Id: &pb.Specimen_Id{
			CollectionId: req.GetId().GetCollectionId(),
			Id:           req.GetId().GetId(),
		},
	}
	// Get the specimen
	if err := state.GetState(&ctx, specimen); err != nil {
		return nil, errors.Wrap(err, "SpecimenHideTx failed to get specimen")
	}

	// Check if the tx is already hidden
	if _, ok := specimen.HiddenTxs[req.GetTx().GetTxId()]; ok {
		return nil, errors.New("Tx already hidden")
	}

	// Check if the tx is the last modified tx
	if specimen.LastModified.GetTxId() == req.GetTx().GetTxId() {
		slog.Error("Tx is the last modified tx", "tx", req.GetTx())
		return nil, errors.New("Tx is the last modified tx")
	}

	// Check if the tx is in history
	found, err := state.TxIdInHistory(ctx, specimen, req.GetTx().GetTxId())
	if err != nil {
		return nil, errors.Wrap(err, "SpecimenHideTx failed to check if tx is in history")
	}
	if !found {
		return nil, errors.New("Failed to find tx in history")
	}

	// Add the hidden tx
	specimen.HiddenTxs[req.GetTx().GetTxId()] = req.GetTx()

	// Register the last modified tx id
	updatedAt, err := ctx.MakeLastModified()
	if err != nil {
		return nil, errors.Wrap(err, "SpecimenHideTx failed to get current modified struct")
	}
	specimen.LastModified = updatedAt

	return specimen, state.UpdateState(ctx, specimen)
}

// SpecimenUnHideTx removes the tx from the specimen's hidden txs
//   - Sets the last modified field
//   - Requires the user to be registered
//   - Requires the user to be a have a role authorized to unHide txs
//   - Requires the tx to be hidden
func (s *SpecimenContract) SpecimenUnHideTx(ctx rbac.AuthTxCtx, req *pb.SpecimenUnHideTxRequest) (*pb.Specimen, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	// -------------------------
	// Authorize the request
	// -------------------------

	// Initialize the auth context via the request, Should return error if not authorized
	if err := ctx.InitViaRequest(req); err != nil {
		return nil, errors.Wrap(err, "SpecimenUnHideTx init via request failed")
	}

	// -------------------------
	// Process the request
	// -------------------------
	specimen := &pb.Specimen{
		Id: &pb.Specimen_Id{
			CollectionId: req.GetId().GetCollectionId(),
			Id:           req.GetId().GetId(),
		},
	}
	// Get the specimen
	if err := state.GetState(&ctx, specimen); err != nil {
		return nil, errors.Wrap(err, "SpecimenUnHideTx failed to get specimen")
	}

	// Check if the tx is already hidden
	if _, ok := specimen.HiddenTxs[req.GetTx().GetTxId()]; !ok {
		return nil, errors.New("Tx Not hidden")
	}

	// Remove the hidden tx from the map
	delete(specimen.HiddenTxs, req.GetTx().GetTxId())

	// Register the last modified tx id
	updatedAt, err := ctx.MakeLastModified()
	if err != nil {
		return nil, errors.Wrap(err, "SpecimenUnHideTx failed to get current modified struct")
	}
	specimen.LastModified = updatedAt

	return specimen, state.UpdateState(ctx, specimen)
}

// -------------------------
// suggest Specimens
// -------------------------

func (s *SpecimenContract) SuggestedUpdateCreate(ctx rbac.AuthTxCtx, req *pb.SuggestedUpdateCreateRequest) (*pb.SuggestedUpdate, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	// -------------------------
	// Authorize the request
	// -------------------------

	// Initialize the auth context via the request, Should return error if not authorized
	if err := ctx.InitViaRequest(req); err != nil {
		return nil, errors.Wrap(err, "SpecimenUnHideTx init via request failed")
	}
	return nil, nil
}

func (s *SpecimenContract) SuggestedUpdateReject(ctx rbac.AuthTxCtx, req *pb.SuggestedUpdateRejectRequest) (*pb.SuggestedUpdate, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	// -------------------------
	// Authorize the request
	// -------------------------

	// Initialize the auth context via the request, Should return error if not authorized
	if err := ctx.InitViaRequest(req); err != nil {
		return nil, errors.Wrap(err, "SpecimenUnHideTx init via request failed")
	}
	return nil, nil
}

func (s *SpecimenContract) SuggestedUpdateApprove(ctx rbac.AuthTxCtx, req *pb.SuggestedUpdateApproveRequest) (*pb.SuggestedUpdate, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	// -------------------------
	// Authorize the request
	// -------------------------

	// Initialize the auth context via the request, Should return error if not authorized
	if err := ctx.InitViaRequest(req); err != nil {
		return nil, errors.Wrap(err, "SpecimenUnHideTx init via request failed")
	}
	return nil, nil
}
