package context

// nolint

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/nova38/thesis/lib/go/fabric/state"

	//"github.com/charmbracelet/log"
	//"github.com/pkg/errors"

	// "github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"

	// "github.com/hyperledger-labs/cckit/identity"

	schema "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v1"
)

//	type AuthTransport struct {
//		User       schema.User
//		Collection schema.Collection
//		Role       schema.Role
//		// Permissions *schema.Permissions
//	}

func HandelBefore(ctx TxContext) error {
	fn, params := ctx.GetStub().GetFunctionAndParameters()

	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		TimeFormat:      time.Kitchen,
		ReportTimestamp: true,
		Prefix:          "Chaincode",
		Level:           log.DebugLevel,
	}).With("fn", fn)
	// ctx.SetLogger(logger)

	logger.Info("Handling Before", "params", params)

	return nil
}

type AuthTxContext struct {
	contractapi.TransactionContext
	Logger      *slog.Logger
	User        *schema.User
	Collection  *schema.Collection
	Role        schema.Role
	Action      *schema.Action
	Authorized  bool
	AuthChecked bool
}

func (ctx *AuthTxContext) GetLogger() *slog.Logger {
	if ctx.Logger != nil {
		return ctx.Logger
	}

	ctx.Logger = slog.With("fn", "GetLogger")
	// ctx.Logger = log.NewWithOptions(os.Stderr, log.Options{})
	return ctx.Logger
}

func (ctx *AuthTxContext) SetLogger(logger *slog.Logger) error {
	if logger == nil {
		return fmt.Errorf("logger is nil")
	}
	ctx.Logger = logger
	return nil
}

func (ctx *AuthTxContext) IsAuthorized() (bool, error) {
	if !ctx.AuthChecked {
		return false, fmt.Errorf("authorization was not checked")
	}
	return ctx.Authorized, nil
}

// InitViaRequest initializes the context based on the request
// It sets the collection and action
// If the action is a preAuth action, it will also check authorization
// If it fails to authorize, it will return an error
// If it succeeds, it will set ctx.Authorized to true
func (ctx *AuthTxContext) InitViaRequest(req interface{}) (err error) {
	// Set preAuth to true if we should immediately check authorization
	preAuth := false

	ctx.Collection = &schema.Collection{
		Id: &schema.Collection_Id{},
	}

	ctx.Action = &schema.Action{
		Level:   schema.Action_LEVEL_UNSPECIFIED,
		Domains: []schema.Action_Domain{},
	}

	// Handle the Collection ID
	switch v := req.(type) {
	case CollectionIDTypeHolder:
		id := v.GetId()
		if id == nil {
			// slog.Warn("failed to extract collection id from the collection payload")
			return oops.
				In("InitViaRequest").
				With("req", req).
				Errorf("failed to extract collection id from the collection payload")
		}

		ctx.Collection.Id.CollectionId = id.GetCollectionId()
	case CollectionHolder:
		collection := v.GetCollection()
		if err != nil {
			return oops.
				In("InitViaRequest").
				With("req", req).
				Errorf("failed to extract collection from the payload")
		}

		id := collection.GetId()
		if id == nil {
			slog.Warn("failed to extract collection id from the collection payload")
			return oops.
				In("InitViaRequest").
				With("req", req).
				Errorf("failed to extract collection id from the collection payload")
		}

		ctx.Collection.Id.CollectionId = id.GetCollectionId()
	case SpecimenIDHolder:
		id := v.GetId()
		if id == nil {
			slog.Warn("failed to extract collection id from the specimen payload")
			return oops.
				In("InitViaRequest").
				With("req", req).
				Errorf("failed to extract collection id from the specimen payload")
		}

		ctx.Collection.Id.CollectionId = id.GetCollectionId()
	case SpecimenHolder:
		specimen := v.GetSpecimen()
		if specimen == nil {
			slog.Warn("failed to extract collection id from the specimen payload")
			return oops.
				In("InitViaRequest").
				With("req", req).
				Errorf("failed to extract collection id from the specimen payload")
		}
		id := specimen.GetId()
		if id == nil {
			slog.Warn("failed to extract collection id from the specimen payload")
			return oops.
				In("InitViaRequest").
				With("req", req).
				Errorf("failed to extract collection id from the specimen payload")
		}

		ctx.Collection.Id.CollectionId = id.GetCollectionId()
	case SuggestionsIDHolder:
		id := v.GetId()
		if id == nil {
			slog.Warn("SuggestionID is nil")
			return fmt.Errorf("failed to extract collection id from the suggestions payload")
		}

		speciemenId := id.GetSpecimenId()
		if speciemenId == nil {
			slog.Warn("SuggestionID's SpecimenID is nil")
			return fmt.Errorf("failed to extract collection id from the suggestions payload")
		}

		ctx.Collection.Id.CollectionId = speciemenId.GetCollectionId()

	default:
		slog.Warn(fmt.Sprintf("Failed to extract collection id from request. Type:%T", v))

		return oops.
			In("InitViaRequest-CollectionExtract").
			With("req", req).
			Errorf("failed to extract collection id from request: %T", v)
	}

	// --------------------------------------------
	// Get Collection from WorldState

	// Set the collection in the context
	if err := ctx.SetCollection(ctx.Collection); err != nil {

		// TODO: Handle when it is okay to not have a collection

		fn, _ := ctx.GetStub().GetFunctionAndParameters()

		if !strings.Contains(fn, "CollectionCreate") {
			return oops.
				In("InitViaRequest-CollectionState").
				With("fn", fn).
				With("collection", ctx.Collection).
				With("req", req).
				Wrapf(err, "Failed to set collection from: %s", err)
		}

	}
	slog.Info("Set collection", "collection", ctx.Collection)

	// --------------------------------------------

	// Helper functions to set the action level
	levelView := func() {
		ctx.Action.Level = schema.Action_LEVEL_VIEW
	}

	setLevel := func(level schema.Action_Level) {
		ctx.Action.Level = level
	}
	addDomain := func(domain schema.Action_Domain) {
		ctx.AddActionDomain(domain)
	}

	// Handle the Action Type
	switch v := req.(type) {

	// Query
	// --------------------------------------------
	// User
	case *schema.GetUserRequest:
		levelView()

	// Collection
	case *schema.GetCollectionRequest:
		levelView()

	// Specimen
	case *schema.GetSpecimenRequest:
		levelView()
	case *schema.GetSpecimenByCollectionRequest:
		levelView()
	case *schema.GetSpecimenHistoryRequest:
		// levelView()
		// This is a special case where they might ask for the hidden txs
		setLevel(schema.Action_LEVEL_HIDE_TX)
	case *schema.GetSuggestedUpdateRequest:
		levelView()
	case *schema.GetSuggestedUpdateBySpecimenRequest:
		levelView()
	case *schema.GetSuggestedUpdateByCollectionRequest:
		levelView()

	// --------------------------------------------
	// Invoke
	// --------------------------------------------

	// User
	case *schema.UserRegisterRequest:
		setLevel(schema.Action_LEVEL_INIT)
	case *schema.UpdateMembershipRequest:
		setLevel(schema.Action_LEVEL_EDIT)
		addDomain(schema.Action_DOMAIN_USERS)
		preAuth = true

	// Collection
	case *schema.CollectionCreateRequest:
		setLevel(schema.Action_LEVEL_INIT)
		addDomain(schema.Action_DOMAIN_COLLECTION)
	case *schema.CollectionUpdateRequest:
		setLevel(schema.Action_LEVEL_EDIT)
		addDomain(schema.Action_DOMAIN_ROLES)
		preAuth = true

	// Specimen
	case *schema.SpecimenCreateRequest:
		setLevel(schema.Action_LEVEL_CREATE)
		addDomain(schema.Action_DOMAIN_SPECIMEN)
		preAuth = true
	case *schema.SpecimenDeleteRequest:
		setLevel(schema.Action_LEVEL_DELETE)
		addDomain(schema.Action_DOMAIN_SPECIMEN)
		preAuth = true
	case *schema.SpecimenHideTxRequest:
		setLevel(schema.Action_LEVEL_HIDE_TX)
		addDomain(schema.Action_DOMAIN_SPECIMEN)
		preAuth = true

	case *schema.SpecimenUnHideTxRequest:
		setLevel(schema.Action_LEVEL_HIDE_TX)
		addDomain(schema.Action_DOMAIN_SPECIMEN)
		preAuth = true

	case *schema.SpecimenUpdateRequest:
		setLevel(schema.Action_LEVEL_EDIT)

	default:
		return oops.
			In("InitViaRequest-ActionType").
			With("req", req).
			Errorf("unknown request type: %T", v)

	}

	if preAuth {
		if valid, err := ctx.Authorize(); err != nil {
			ctx.AuthChecked = true
			ctx.Authorized = false
			slog.Error("Failed to Authorize", "err", err)
			return fmt.Errorf("Failed to Authorize: %w", err)

		} else {
			ctx.AuthChecked = true
			ctx.Authorized = valid

			if !valid {

				slog.Error("PreAuth: Not authorized", "actions", ctx.Action)
				return errors.New("Not authorized")
			}
			slog.Debug("Authorized", "valid", valid)
		}
	}

	return nil
}

func (ctx *AuthTxContext) MakeLastModified() (*schema.LastModified, error) {
	user, err := ctx.GetUser()
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	timestamp, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return nil, fmt.Errorf("Failed to get timestamp: %w", err)
	}

	return &schema.LastModified{
		UserId:    user.GetId(),
		UserName:  user.GetName(),
		TxId:      ctx.GetStub().GetTxID(),
		UpdatedAt: timestamp,
	}, nil
}

func (ctx *AuthTxContext) SetAction(action *schema.Action) error {
	if action == nil {
		return fmt.Errorf("action is nil")
	}

	ctx.Action = action

	return nil
}

func (ctx *AuthTxContext) GetActionDomains() []schema.Action_Domain {
	if ctx.Action == nil {
		return []schema.Action_Domain{}
	}

	return ctx.Action.Domains
}

func (ctx *AuthTxContext) AddActionDomain(domain schema.Action_Domain) {
	ctx.Action.Domains = append(ctx.Action.Domains, domain)
}

func (ctx *AuthTxContext) GetUser() (*schema.User, error) {
	if ctx.User != nil {
		return ctx.User, nil
	}

	id, err := ctx.GetUserId()
	if err != nil {
		return nil, err
	}

	ctx.User = &schema.User{
		Id: &schema.User_Id{
			MspId: id.GetMspId(),
			Id:    id.GetId(),
		},
	}

	if err := state.GetState(ctx, ctx.User); err != nil {
		return nil, fmt.Errorf("failed to get user from state: %w", err)
	}

	return ctx.User, nil
}

func (ctx *AuthTxContext) GetUserId() (*schema.User_Id, error) {
	var err error

	//if ctx.currentUser != nil {
	//	return ctx.currentUser, nil
	//}

	Id := &schema.User_Id{
		MspId: "",
		Id:    "",
	}

	// Extract The info from the Client ID
	id := ctx.GetClientIdentity()

	Id.Id, err = id.GetID()
	if err != nil {
		return nil, fmt.Errorf("failed to get user certificate from CID: %s", err)
	}

	Id.MspId, err = id.GetMSPID()
	if err != nil {
		return nil, fmt.Errorf("failed to get user ID from CID: %s", err)
	}

	return Id, nil
}

func (ctx *AuthTxContext) SetCollection(collection *schema.Collection) error {
	if collection == nil {
		return fmt.Errorf("collection is nil")
	}

	ctx.Collection = collection

	if err := state.GetState(ctx, ctx.Collection); err != nil {
		return fmt.Errorf("Failed to get collection from state:%w", err)
	}

	return nil
}

func (ctx *AuthTxContext) GetCollection() (*schema.Collection, error) {
	if ctx.Collection != nil {
		return ctx.Collection, nil
	}

	return nil, fmt.Errorf("collection is nil")
}

func (ctx *AuthTxContext) GetRole() (schema.Role, error) {
	user, err := ctx.GetUser()
	if err != nil {
		return schema.Role_ROLE_PUBLIC_UNSPECIFIED, fmt.Errorf("failed to get user: %w", err)
	}

	collection, err := ctx.GetCollection()
	if err != nil {
		return schema.Role_ROLE_PUBLIC_UNSPECIFIED, fmt.Errorf("failed to get collection: %w", err)
	}

	if collection.GetId() == nil {
		return schema.Role_ROLE_PUBLIC_UNSPECIFIED, fmt.Errorf("collection.GetId() is nil")
	}
	collectionId := collection.GetId().GetCollectionId()

	// At this point we have a user and a collection, so we shouldn't return an error

	if user.Memberships == nil {
		slog.Warn("user.Memberships is nil")
		return schema.Role_ROLE_PUBLIC_UNSPECIFIED, nil
		// return nil, fmt.Errorf("user.Memberships is nil")
	}

	if role, exists := user.Memberships[collectionId]; exists {
		return role, nil
	}

	return schema.Role_ROLE_PUBLIC_UNSPECIFIED, nil
}

func (ctx *AuthTxContext) Authorize() (valid bool, err error) {
	action := ctx.Action

	if action == nil {
		return false, fmt.Errorf("action is nil")
	}

	if err := action.ValidateAll(); err != nil {
		return false, fmt.Errorf("failed to validate action: %w", err)
	}

	ctx.Action = action

	// Get The collection and the role

	collection, err := ctx.GetCollection()
	if err != nil {
		return false, fmt.Errorf("failed to get collection: %w", err)
	}
	ac := collection.GetAccessControl()
	if ac == nil {
		return false, fmt.Errorf("ac is nil")
	}
	if err := ac.ValidateAll(); err != nil {
		return false, fmt.Errorf("Failed to validate ac: %w", err)
	}

	role, err := ctx.GetRole()
	if err != nil {
		return false, fmt.Errorf("Failed to get role: %w", err)
	}

	// AccessControlActions Sections:
	//	roles

	// AccessControlActions
	//	users

	// SpecimenActions sections:
	//	specimen

	// SectionActions sections:
	//	primary
	//	secondary
	//	taxon
	//	georeference
	//	images
	//	loans
	//	grants
	//	hidden

	// SuggestedActions sections:
	// suggested

	invalidAction := false

	// Check if the action is valid for the role
	switch action.GetLevel() {

	case schema.Action_LEVEL_UNSPECIFIED:
		return false, fmt.Errorf("action is unspecified")
	case schema.Action_LEVEL_VIEW:
		// Can be any of the following type of roles
		// 		- AccessControlActions
		//		- SpecimenActions
		//		- SectionActions
		return true, nil
	case schema.Action_LEVEL_INIT:
		// Can be any of the following type of roles
		return true, nil
	case schema.Action_LEVEL_EDIT:
		// Can be any of the following type of roles
		// 		- AccessControlActions
		//		- SectionActions
		valid = lo.EveryBy(action.GetDomains(), func(domain schema.Action_Domain) bool {
			switch domain {
			// AccessControlActions
			case schema.Action_DOMAIN_ROLES:
				return lo.Contains(ac.GetRoles().GetEdit(), role)
			case schema.Action_DOMAIN_USERS:
				return lo.Contains(ac.GetUsers().GetEdit(), role)
			// SectionActions
			case schema.Action_DOMAIN_PRIMARY:
				return lo.Contains(ac.GetPrimary().GetEdit(), role)
			case schema.Action_DOMAIN_SECONDARY:
				return lo.Contains(ac.GetSecondary().GetEdit(), role)
			case schema.Action_DOMAIN_TAXON:
				return lo.Contains(ac.GetTaxon().GetEdit(), role)
			case schema.Action_DOMAIN_GEOREFERENCE:
				return lo.Contains(ac.GetGeoreference().GetEdit(), role)
			case schema.Action_DOMAIN_IMAGES:
				return lo.Contains(ac.GetImages().GetEdit(), role)
			case schema.Action_DOMAIN_LOANS:
				return lo.Contains(ac.GetLoans().GetEdit(), role)
			case schema.Action_DOMAIN_GRANTS:
				return lo.Contains(ac.GetGrants().GetEdit(), role)

			default:
				slog.Warn("Invalid domain", "level", action.GetLevel(), "domain", domain)
				invalidAction = true
				return false
			}
		})
		if invalidAction {
			return false, fmt.Errorf("invalid action")
		}
		return valid, nil
	case schema.Action_LEVEL_SUGGEST_EDIT:
		// Can be any of the following type of roles
		// 		- SectionActions
		// 		- AccessControlActions
		valid = lo.EveryBy(action.GetDomains(), func(domain schema.Action_Domain) bool {
			switch domain {
			// SectionActions
			case schema.Action_DOMAIN_PRIMARY:
				return lo.Contains(ac.GetPrimary().GetSuggestEdit(), role)
			case schema.Action_DOMAIN_SECONDARY:
				return lo.Contains(ac.GetSecondary().GetSuggestEdit(), role)
			case schema.Action_DOMAIN_TAXON:
				return lo.Contains(ac.GetTaxon().GetSuggestEdit(), role)
			case schema.Action_DOMAIN_GEOREFERENCE:
				return lo.Contains(ac.GetGeoreference().GetSuggestEdit(), role)
			case schema.Action_DOMAIN_IMAGES:
				return lo.Contains(ac.GetImages().GetSuggestEdit(), role)
			case schema.Action_DOMAIN_LOANS:
				return lo.Contains(ac.GetLoans().GetSuggestEdit(), role)
			case schema.Action_DOMAIN_GRANTS:
				return lo.Contains(ac.GetGrants().GetSuggestEdit(), role)
			default:
				slog.Warn("invalid domain", "level", action.GetLevel(), "domain", domain)
				invalidAction = true
				return false
			}
		})
		if invalidAction {
			return false, fmt.Errorf("invalid action")
		}
		return valid, nil
	case schema.Action_LEVEL_SUGGEST_APPROVE:
		// Can be any of the following type of roles
		// 		- SuggestedActions
		valid = lo.EveryBy(action.GetDomains(), func(domain schema.Action_Domain) bool {
			switch domain {
			// SectionActions
			case schema.Action_DOMAIN_PRIMARY:
				return lo.Contains(ac.GetPrimary().GetSuggestApprove(), role)
			case schema.Action_DOMAIN_SECONDARY:
				return lo.Contains(ac.GetSecondary().GetSuggestApprove(), role)
			case schema.Action_DOMAIN_TAXON:
				return lo.Contains(ac.GetTaxon().GetSuggestApprove(), role)
			case schema.Action_DOMAIN_GEOREFERENCE:
				return lo.Contains(ac.GetGeoreference().GetSuggestApprove(), role)
			case schema.Action_DOMAIN_IMAGES:
				return lo.Contains(ac.GetImages().GetSuggestApprove(), role)
			case schema.Action_DOMAIN_LOANS:
				return lo.Contains(ac.GetLoans().GetSuggestApprove(), role)
			case schema.Action_DOMAIN_GRANTS:
				return lo.Contains(ac.GetGrants().GetSuggestApprove(), role)
			default:
				slog.Warn("invalid domain", "level", action.GetLevel(), "domain", domain)
				invalidAction = true
				return false
			}
		})
		if invalidAction {
			return false, fmt.Errorf("invalid action")
		}
		return valid, nil
	case schema.Action_LEVEL_SUGGEST_REJECT:
		// Can be any of the following type of roles
		// 		- SectionActions
		valid = lo.EveryBy(action.GetDomains(), func(domain schema.Action_Domain) bool {
			switch domain {
			// SectionActions
			case schema.Action_DOMAIN_PRIMARY:
				return lo.Contains(ac.GetPrimary().GetSuggestReject(), role)
			case schema.Action_DOMAIN_SECONDARY:
				return lo.Contains(ac.GetSecondary().GetSuggestReject(), role)
			case schema.Action_DOMAIN_TAXON:
				return lo.Contains(ac.GetTaxon().GetSuggestReject(), role)
			case schema.Action_DOMAIN_GEOREFERENCE:
				return lo.Contains(ac.GetGeoreference().GetSuggestReject(), role)
			case schema.Action_DOMAIN_IMAGES:
				return lo.Contains(ac.GetImages().GetSuggestReject(), role)
			case schema.Action_DOMAIN_LOANS:
				return lo.Contains(ac.GetLoans().GetSuggestReject(), role)
			case schema.Action_DOMAIN_GRANTS:
				return lo.Contains(ac.GetGrants().GetSuggestReject(), role)
			default:
				slog.Warn("invalid domain", "level", action.GetLevel(), "domain", domain)
				invalidAction = true
				return false
			}
		})

		if invalidAction {
			return false, fmt.Errorf("invalid action")
		}
		return valid, nil
	case schema.Action_LEVEL_CREATE:
		// Can be any of the following type of roles
		// 		- SpecimenActions
		return lo.Contains(ac.GetSpecimen().GetCreate(), role), nil
	case schema.Action_LEVEL_DELETE:
		// Can be any of the following type of roles
		// 		- SpecimenActions
		return lo.Contains(ac.GetSpecimen().GetDelete(), role), nil
	case schema.Action_LEVEL_HIDE_TX:
		// Can be any of the following type of roles
		// 		- SpecimenActions
		return lo.Contains(ac.GetSpecimen().GetHideTx(), role), nil
	default:
		slog.Warn("invalid level", "level", action.GetLevel())
		invalidAction = true
		return false, fmt.Errorf("invalid level")
	}
}

func (ctx *AuthTxContext) GetAction() (*schema.Action, error) {
	if ctx.Action != nil {
		return ctx.Action, nil
	}

	return nil, fmt.Errorf("action is nil")
}

func (ctx *AuthTxContext) MakeClearedLastModified() *schema.LastModified {
	return &schema.LastModified{
		UserId: &schema.User_Id{
			MspId: "",
			Id:    "",
		},
		UserName: "",
		TxId:     "",
		UpdatedAt: &timestamppb.Timestamp{
			Seconds: 0,
			Nanos:   0,
		},
	}
}

// func (ctx *AuthTxContext) SetCollectionViaRequest(req interface{}) (err error) {
// 	ctx.Collection = &schema.Collection{
// 		Id: &schema.Collection_Id{},
// 	}

// 	switch v := req.(type) {
// 	case CollectionIDTypeHolder:
// 		id := v.GetId()
// 		if id == nil {
// 			slog.Warn("failed to extract collection id from the collection payload")
// 			return fmt.Errorf("failed to extract collection id from the collection payload")
// 		}

// 		ctx.Collection.Id.CollectionId = id.GetCollectionId()

// 	case SpecimenIDHolder:
// 		id := v.GetId()
// 		if id == nil {
// 			slog.Warn("failed to extract collection id from the specimen payload")
// 			return fmt.Errorf("failed to extract collection id from the specimen payload")
// 		}

// 		ctx.Collection.Id.CollectionId = id.GetCollectionId()

// 	default:
// 		return fmt.Errorf("unknown request type: %T", v)
// 	}

// 	return nil
// }
