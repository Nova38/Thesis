package rbac

// import (
// 	"github.com/nova38/thesis/packages/saacs/auth/common"
// 	"github.com/nova38/thesis/packages/fabric/auth/state"
// 	"google.golang.org/protobuf/types/known/fieldmaskpb"

// 	"github.com/samber/oops"

// 	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
// 	cc "github.com/nova38/thesis/packages/saacs/genchaincode/auth/rbac/schema/v1"
// )

// func (a RbacContractImpl) GetMembershipByUser(
// 	ctx *AuthTxCtx,
// 	req *cc.MembershipsGetByUserRequest,
// ) (res *cc.GetMembershipsByUserResponse, err error) {
// 	defer func() { ctx.HandleFnError(&err, recover()) }()

// 	// Validate the request
// 	if err = ctx.Validate(req); err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	user := authpb.User{
// 		MspId:  req.MspId,
// 		UserId: req.UserId,
// 	}

// 	if exists := state.PrimaryExists(ctx, &user); !exists {
// 		return nil, oops.Wrap(common.UserUnregistered)
// 	}

// 	mem := authpb.Membership{
// 		CollectionId: req.CollectionId,
// 		MspId:        req.MspId,
// 		UserId:       req.UserId,
// 	}

// 	// res = &cc.UserGetMembershipResponse{Membership: &mem}

// 	list, _, err := state.PrimaryByPartialKey(ctx, &mem, 3, "")
// 	if err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	res = &cc.GetMembershipsByUserResponse{
// 		User:        &user,
// 		Memberships: list,
// 	}

// 	return res, nil
// }

// func (a RbacContractImpl) GetMembershipByCollection(
// 	ctx *AuthTxCtx,
// 	req *cc.GetMembershipsByCollectionRequest,
// ) (res *cc.GetMembershipsByCollectionResponse, err error) {
// 	defer func() { ctx.HandleFnError(&err, recover()) }()

// 	// Validate the request
// 	if err = ctx.Validate(req); err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	mem := authpb.Membership{
// 		CollectionId: req.CollectionId,
// 	}

// 	list, _, err := state.PrimaryByPartialKey(ctx, &mem, 3, "")
// 	if err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	res = &cc.GetMembershipsByCollectionResponse{Memberships: list}

// 	return res, nil
// }

// // ══════════════════════════════════════════════════════════════════════════════==
// // Invoke Functions
// // ══════════════════════════════════════════════════════════════════════════════==
// func (a RbacContractImpl) CreateMembership(
// 	ctx *AuthTxCtx,
// 	req *cc.UpdateMembershipRequest,
// ) (res *cc.UpdateMembershipResponse, err error) {
// 	defer func() { ctx.HandleFnError(&err, recover()) }()

// 	// Validate the request
// 	if err = ctx.Validate(req); err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	mem := authpb.Membership{
// 		CollectionId: req.CollectionId,
// 		RoleId:       req.RoleId,
// 		MspId:        req.MspId,
// 		UserId:       req.UserId,
// 	}
// 	res = &cc.UpdateMembershipResponse{Membership: &mem}

// 	return res, state.PrimaryCreate(ctx, &mem)
// }

// func (a RbacContractImpl) UpdateMembership(
// 	ctx *AuthTxCtx,
// 	req *cc.UpdateMembershipRequest,
// ) (res *cc.UpdateMembershipResponse, err error) {
// 	defer func() { ctx.HandleFnError(&err, recover()) }()

// 	// Validate the request
// 	if err = ctx.Validate(req); err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	mem := authpb.Membership{
// 		CollectionId: req.CollectionId,
// 		RoleId:       req.RoleId,
// 		MspId:        req.MspId,
// 		UserId:       req.UserId,
// 	}
// 	res = &cc.UpdateMembershipResponse{Membership: &mem}

// 	return res, state.PrimaryUpdate(ctx, &mem, &fieldmaskpb.FieldMask{})
// }

// func (a RbacContractImpl) DeleteMembership(
// 	ctx *AuthTxCtx,
// 	req *cc.DeleteMembershipRequest,
// ) (res *cc.DeleteMembershipResponse, err error) {
// 	defer func() { ctx.HandleFnError(&err, recover()) }()
// 	// Validate the request
// 	if err = ctx.Validate(req); err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	mem := authpb.Membership{
// 		CollectionId: req.CollectionId,
// 		RoleId:       req.RoleId,
// 		MspId:        req.MspId,
// 		UserId:       req.UserId,
// 	}

// 	res = &cc.DeleteMembershipResponse{Membership: &mem}

// 	return res, state.PrimaryDelete(ctx, &mem)
// }
