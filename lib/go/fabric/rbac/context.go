package rbac

import (
    "github.com/nova38/thesis/lib/go/fabric/state"
    "github.com/samber/oops"
    _ "strings"
    // "github.com/rs/zerolog/log"
    _ "github.com/samber/lo"

    _ "google.golang.org/protobuf/types/known/timestamppb"

    // "github.com/hyperledger-labs/cckit/identity"

    pb "github.com/nova38/thesis/lib/go/gen/rbac"
)

type TransactionObjects struct {
    User       *pb.User
    Collection *pb.Collection
    Domain     *pb.Operations_Domain
    Action     *pb.Operations_Action
}

type AuthTxCtx struct {
    state.LoggedTxCtx

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

func (ctx *AuthTxCtx) GetDomain() (*pb.Operations_Domain, error) {
    if ctx.Domain != nil {
        return ctx.Domain, nil
    }

    return nil, oops.Errorf("domain not set")
}

func (ctx *AuthTxCtx) SetDomain(domain *pb.Operations_Domain) error {
    //TODO implement me
    panic("implement me")
}

func (ctx *AuthTxCtx) GetAction() (*pb.Operations_Action, error) {
    //TODO implement me
    panic("implement me")
}

func (ctx *AuthTxCtx) SetAction(action *pb.Operations_Action) error {
    // check if the action is set
    if action == nil {
        return oops.Errorf("action is nil")
    }

    // validate the action
    
}

func (ctx *AuthTxCtx) Authorize() (bool, error) {

    if ctx.AuthChecked {
        return ctx.Authorized, nil
    }

    // Check if all the objects are set
    if ctx.Collection == nil || ctx.Domain == nil || ctx.Action == nil {
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
