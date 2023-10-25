package contracts

import (
	"fmt"
	"log/slog"

	"github.com/nova38/thesis/lib/go/fabric/state"

	"github.com/nova38/thesis/apps/chaincode/ccbio/v1/context"

	"github.com/pkg/errors"
	"github.com/samber/oops"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	schema "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v1"
)

// AuthContract contract for handling BasicAssets
type AuthContract struct {
	contractapi.Contract
}

func BuildAuthContract() *AuthContract {
	return &AuthContract{
		Contract: contractapi.Contract{
			Name:              "ccbio.Auth",
			BeforeTransaction: context.HandelBefore,
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
// Contract interface special functions
// ────────────────────────────────────────────────────────────

func (s *AuthContract) GetEvaluateTransactions() []string {
	return []string{
		"GetUser",
		"GetCurrentUserId",
		"GetUser",
		"GetUserList",
		"GetCollection",
		"GetCollectionList",
	}
}

func (s *AuthContract) GetIgnoredFunctions() []string {
	return []string{}
}

// ────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────
// Query Functions
// ────────────────────────────────────────────────────────────

// -------------------------
// User
// -------------------------
func (s *AuthContract) GetCurrentUser(ctx context.TxContext) (*schema.User, error) {
	// return ctx.CurrentUser()

	id, err := ctx.GetUserId()
	if err != nil {
		return nil, err
	}

	user := &schema.User{
		Id: &schema.User_Id{
			MspId: id.GetMspId(),
			Id:    id.GetId(),
		},
	}

	if err := state.GetState(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthContract) GetCurrentUserId(ctx context.TxContext) (*schema.User_Id, error) {
	return ctx.GetUserId()
}

// GetUser  returns the user with the given ID
func (s *AuthContract) GetUser(ctx context.TxContext, id *schema.GetUserRequest) (user *schema.User, er error) {
	if err := id.ValidateAll(); err != nil {
		return nil, err
	}

	idReq := id.GetId()
	if idReq == nil {
		return nil, fmt.Errorf("id is nil")
	}

	user = &schema.User{
		Id: &schema.User_Id{
			MspId: id.GetId().GetMspId(),
			Id:    id.GetId().GetId(),
		},
	}

	slog.Info("GetUser", "user", user)

	if err := state.GetState(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserList returns all users
func (s *AuthContract) GetUserList(ctx context.TxContext) (users *schema.User_List, err error) {
	list, err := state.GetFullStateList(ctx, &schema.User{})
	if err != nil {
		return nil, err
	}

	users = &schema.User_List{
		Items: list,
	}

	return users, nil
}

//func GetUserByCollection(ctx TxContext, req schema.GetUserByCollectionRequest) ([]*schema.User, error) {
//	if err := req.ValidateAll(); err != nil {
//		return nil, err
//	}
//	id := req.GetId()
//
//	// Check if collection exists
//	colExists := state.Exists(ctx, &schema.Collection{
//		Id: &schema.Collection_Id{
//			CollectionId: id.GetCollectionId(),
//		},
//	})
//
//	if !colExists {
//		return nil, &state.KeyNotFoundError{
//			Key:       id.GetCollectionId(),
//			Namespace: schema.NS_COL,
//			MSG:       "Collection not found",
//		}
//	}
//
//	users, err := state.GetFullStateList(ctx, &schema.User{})
//	if err != nil {
//		return nil, err
//	}
//
//	lo.Filter(users, func(user *schema.User) bool {
//		return user.Memberships[id.GetCollectionId()] != schema.Role_ROLE_NONE
//	})
//
//
//}

// -------------------------
// Collection
// -------------------------

func (s *AuthContract) GetCollection(ctx context.TxContext, req *schema.GetCollectionRequest) (*schema.Collection, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}
	id := req.GetId()

	collection := &schema.Collection{
		Id: &schema.Collection_Id{
			CollectionId: id.GetCollectionId(),
		},
	}

	if err := state.GetState(ctx, collection); err != nil {
		return nil, err
	}

	return collection, nil
}

func (s *AuthContract) GetCollectionList(ctx context.TxContext) (*schema.Collection_List, error) {
	list, err := state.GetFullStateList(ctx, &schema.Collection{})
	if err != nil {
		return nil, err
	}

	return &schema.Collection_List{Items: list}, nil
}

// ────────────────────────────────────────────────────────────
// Invoke Functions
// ────────────────────────────────────────────────────────────

// -------------------------
// User
// -------------------------

func (s *AuthContract) UserRegister(ctx context.TxContext, req *schema.UserRegisterRequest) (*schema.User, error) {
	id, err := ctx.GetUserId()
	if err != nil {
		return nil, err
	}

	user := &schema.User{
		Id: &schema.User_Id{
			MspId: id.GetMspId(),
			Id:    id.GetId(),
		},
		Name:        req.Name,
		Affiliation: req.Affiliation,
		Email:       req.Email,
		Memberships: map[string]schema.Role{},
	}

	if err := state.InsertState(ctx, user); err != nil {
		return nil, errors.Wrap(err, "UserRegister failed")
	}

	slog.Info(fmt.Sprintf("User registered: %+v", user))

	return user, nil
}

// AddTestUsers returns the user who is invoking the transaction
func (s *AuthContract) AddTestUsers(ctx context.TxContext) (*schema.User, error) {
	fmt.Println("UserTesting")
	user := schema.User{
		Id: &schema.User_Id{
			MspId: "MSPID",
			Id:    "ID",
		},
		Name:        "Fake User",
		Email:       "item@tes.co",
		Affiliation: "I am a affiliation",
		Memberships: map[string]schema.Role{"base": schema.Role_ROLE_MANAGER},
	}

	slog.Info("AddTestUsers", "users", user.String())
	err := state.InsertState(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthContract) UserUpdateMembership(ctx context.TxContext, req *schema.UpdateMembershipRequest) (*schema.User, error) {
	// Validate request
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	// -------------------------
	// Authorize the request
	// -------------------------

	// Set the collection and do authorization, Should return error if not authorized
	if err := ctx.InitViaRequest(req); err != nil {
		return nil, errors.Wrap(err, "Collection Update via request failed")
	}

	// -------------------------
	// Process the request
	// -------------------------

	// Get the user to modify
	userToModify := &schema.User{
		Id: &schema.User_Id{
			MspId: req.GetUserId().GetMspId(),
			Id:    req.GetUserId().GetId(),
		},
	}
	if err := state.GetState(ctx, userToModify); err != nil {
		return nil, errors.Wrap(err, "GetState failed")
	}

	// Update the user
	userToModify.Memberships[req.GetCollectionId().CollectionId] = req.NewRole

	return userToModify, state.UpdateState(ctx, userToModify)
}

// --------------------------------------------------
// Collection
// --------------------------------------------------

// CollectionCreate creates a new collection and adds the user as a manager
// Requires the user to be registered
func (s *AuthContract) CollectionCreate(ctx context.TxContext, req *schema.CollectionCreateRequest) (*schema.Collection, error) {
	logger := ctx.GetLogger()
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if err := req.ValidateAll(); err != nil {
		// logger.Error("CollectionCreate failed", "err", err)
		logger.Error(err.Error(), "error", err)

		return nil, oops.
			In("CollectionCreate").
			Hint("Invalid request").
			Wrapf(err, "CollectionCreate failed")
	}

	// Set the collection and do authorization
	if err := ctx.InitViaRequest(req); err != nil {
		logger.Error("CollectionCreate failed", "err", err)
		return nil, oops.
			In("CollectionCreate-Init").
			Wrapf(err, "Failed Extracting collection from request")
	}

	// Get the user to modify
	user, err := ctx.GetUser()
	if err != nil {
		return nil, oops.
			In("CollectionCreate").
			Wrapf(err, "user not found")
	}

	// Create the new collection
	if err := state.InsertState(ctx, req.GetCollection()); err != nil {
		// slog.Error("CollectionCreate failed", "err", err)
		logger.Error("Failed to insert new state", "error", err)

		return nil, oops.
			In("CollectionCreate-InsertState").
			User(user.String()).
			Wrapf(err, "Failed to insert new state")
	}

	// Update the User's memberships
	collection_id := req.GetCollection().GetId().GetCollectionId()
	new_memberships := map[string]schema.Role{collection_id: schema.Role_ROLE_MANAGER}

	if user.Memberships == nil {
		user.Memberships = new_memberships
	} else {
		user.Memberships[collection_id] = schema.Role_ROLE_MANAGER
	}

	if err := state.UpdateState(ctx, user); err != nil {
		slog.Error("CollectionCreate failed", "err", err)
		return nil, oops.
			In("CollectionCreate-UpdateMembership").
			User(user.String()).
			Wrapf(err, "Failed to update user memberships")
	}

	return req.Collection, nil
}

// CollectionUpdate updates the collection role permissions
func (s *AuthContract) CollectionUpdate(ctx context.TxContext, req *schema.CollectionUpdateRequest) (*schema.Collection, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	// -------------------------
	// Authorize the request
	// -------------------------

	// Set the collection and do authorization
	if err := ctx.InitViaRequest(req); err != nil {
		return nil, errors.Wrap(err, "Collection Update via request failed")
	}

	// -------------------------
	// Process the request
	// -------------------------
	if err := state.UpdateState(ctx, req.GetCollection()); err != nil {
		return nil, errors.Wrap(err, "CollectionUpdate failed")
	}

	return req.GetCollection(), nil
}
