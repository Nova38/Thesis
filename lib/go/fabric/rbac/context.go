package rbac

import (
	_ "strings"

	"github.com/nova38/thesis/lib/go/fabric/state"
	"github.com/samber/oops"

	// "github.com/rs/zerolog/log"
	_ "github.com/samber/lo"

	_ "google.golang.org/protobuf/types/known/timestamppb"

	// "github.com/hyperledger-labs/cckit/identity"

	pb "github.com/nova38/thesis/lib/go/gen/rbac"
)

type TransactionObjects struct {
	User       *pb.User
	Collection *pb.Collection
	ops        *pb.Operations
}

type AuthTxCtx struct {
	state.LoggedTxCtx
	state.RegistryTxCtx

	TransactionObjects

	Authorized  bool
	AuthChecked bool
}

func (ctx *AuthTxCtx) SetCollection(collection *pb.Collection) error {

	// See if the collection pointer has an ID and is not nil
	if collection == nil || collection.Id == nil || collection.Id.CollectionId == "" {
		return oops.Errorf("collection is nil or has no ID")
	}

	ctx.Collection = &pb.Collection{
		Id: &pb.Collection_Id{
			CollectionId: collection.Id.CollectionId,
		},
	}

	// Check if the collection exists
	err := state.GetState(ctx, ctx.Collection)

	if err != nil {
		return oops.
			In("SetCollection").
			With("collectionId", collection.GetId().CollectionId).
			Wrap(err)
	}

	return nil
}

func (ctx *AuthTxCtx) GetOperation() (*pb.Operations, error) {
	if ctx.Domain == pb.Operations_DOMAIN_UNSPECIFIED || ctx.Action == nil {
		return nil, oops.Errorf("operation not set")
	}

	return &pb.Operations{
		Domain: ctx.Domain,
		Action: ctx.Action,
	}, nil
}

func (ctx *AuthTxCtx) SetOperation(op *pb.Operations) error {

	// See if the operation pointer has an ID and is not nil
	if op == nil || op.Action == nil {
		return oops.Errorf("operation is nil or actions is nil")
	}

	// Set Domain

	if op.Domain == pb.Operations_DOMAIN_UNSPECIFIED {
		return oops.Errorf("operation domain is unspecified")
	}

	ctx.Domain = op.Domain

	// Set Action

	// switch op.Action.Type {
	// case pb.Operations_Action_TYPE_UNSPECIFIED:
	// 	return oops.Errorf("operation action type is unspecified")
	// case pb.Operations_Action_TYPE_CREATE, pb.Operations_Action_TYPE_DELETE:
	// }

	if op.Action.Type == pb.Operations_Action_TYPE_UNSPECIFIED {
		return oops.Errorf("operation action type is unspecified")
	}

	ctx.Action = op.Action

	// TODO: Should we validate the action here?

	return nil
}

func (ctx *AuthTxCtx) Authorize() (bool, error) {

	if ctx.AuthChecked {
		return ctx.Authorized, nil
	}

	// Check if all the objects are set
	if ctx.Collection == nil || ctx.Action == nil {
		return false, oops.Errorf("authorization objects not set")
	}

	return true, nil
}

func (ctx *AuthTxCtx) IsAuthorized() (bool, error) {
	if !ctx.AuthChecked {
		return false, oops.Errorf("authorization not checked")
	}
	return ctx.Authorized, nil

}

func (ctx *AuthTxCtx) GetUser() (*pb.User, error) {

	if ctx.User != nil {
		return ctx.User, nil
	}

	id, err := ctx.GetUserId()

	if err != nil {
		return nil, oops.Wrap(err)
	}

	ctx.User = &pb.User{Id: id}

	err = state.GetState(ctx, ctx.User)

	if err != nil {
		return nil, oops.With("user_id", id).Wrap(err)
	}

	return ctx.User, nil
}

func (ctx *AuthTxCtx) GetUserId() (*pb.User_Id, error) {
	var err error

	Id := &pb.User_Id{
		MspId: "",
		Id:    "",
	}

	// Extract The info from the Client ID
	id := ctx.GetClientIdentity()

	Id.Id, err = id.GetID()
	if err != nil {
		return nil, oops.Errorf("failed to get user certificate from CID: %s", err)
	}

	Id.MspId, err = id.GetMSPID()
	if err != nil {
		return nil, oops.Errorf("failed to get user ID from CID: %s", err)
	}

	return Id, nil
}

func (ctx *AuthTxCtx) GetCollection() (*pb.Collection, error) {
	if ctx.Collection != nil {
		return ctx.Collection, nil
	}

	return nil, oops.Errorf("collection not set")
}
