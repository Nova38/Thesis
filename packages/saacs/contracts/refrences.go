package contracts

import (
	"github.com/nova38/thesis/packages/fabric/auth/state"
	"github.com/nova38/thesis/packages/saacs/auth/common"
	ccpb "github.com/nova38/thesis/packages/saacs/genchaincode/auth/common"
	"github.com/samber/oops"
)

// ════════════════════════════════════ References ═════════════════════════════════
// ──────────────────────────────────── Query ──────────────────────────────────────

// Reference returns the reference if it exists
func (o ItemContractImpl) Reference(
	ctx common.TxCtxInterface,
	req *ccpb.ReferenceRequest,
) (res *ccpb.ReferenceResponse, err error) {
	// todo: Reference

	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	if v, err := state.GetReference(ctx, req.GetReference()); err != nil {
		return nil, oops.Wrap(err)
	} else if v == nil && err == nil {
		return &ccpb.ReferenceResponse{
			Exists: false,
		}, nil
	}

	return &ccpb.ReferenceResponse{
		Exists: true,
	}, nil
}

// ReferenceByPartialKey
func (o ItemContractImpl) ReferenceByPartialKey(
	ctx common.TxCtxInterface,
	req *ccpb.ReferenceByPartialKeyRequest,
) (res *ccpb.ReferenceByPartialKeyResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	list, mk, err := state.PartialReferenceKeysList(
		ctx,
		req.GetReference(),
		int(ctx.GetPageSize()),
		req.GetBookmark(),
	)
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	res = &ccpb.ReferenceByPartialKeyResponse{
		Bookmark:   mk,
		References: list,
	}
	return res, nil
}

// // todo: ReferenceListByType
// func (o ItemContractImpl) ReferenceByType(
// 	ctx state.TxCtxInterface,
// 	req *cc.ReferenceListByTypeRequest,
// ) (res *cc.ReferenceListByTypeResponse, err error) {
// 	// Validate the request
// 	if err = ctx.Validate(req); err != nil {
// 		return nil, oops.Wrap(err)
// 	}
// 	list, mk, err := state.ReferenceByType(ctx, req.GetReferenceType(), req.GetBookmark())
// 	if err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	res = &cc.ReferenceListByTypeResponse{
// 		Bookmark:   mk,
// 		References: list,
// 	}

// 	return res, nil
// }

// todo: ReferenceByCollection
// func (o ItemContractImpl) ReferenceByCollection(
// 	ctx common.TxCtxInterface,
// 	req *cc.ReferenceByCollectionRequest,
// ) (res *cc.ReferenceByCollectionResponse, err error) {
// 	// Validate the request
// 	if err = ctx.Validate(req); err != nil {
// 		return nil, oops.Wrap(err)
// 	}
// 	list, mk, err := state.ReferenceKeysByCollection(ctx, req.GetCollectionId(), req.GetBookmark(), "")
// 	if err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	res = &cc.ReferenceByCollectionResponse{
// 		Bookmark:   mk,
// 		References: list,
// 	}

// 	return res, nil
// }

// todo: ReferenceByItem
func (o ItemContractImpl) ReferenceByItem(
	ctx common.TxCtxInterface,
	req *ccpb.ReferenceByItemRequest,
) (res *ccpb.ReferenceByItemResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	list, mk, err := state.ReferenceKeysByItem(ctx, req.GetItemKey(), req.GetBookmark())
	if err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	res = &ccpb.ReferenceByItemResponse{
		Bookmark:   mk,
		References: list,
	}

	return res, nil
}

// ──────────────────────────────────── Invoke ─────────────────────────────────────

// ReferenceCreate creates a reference
func (o ItemContractImpl) ReferenceCreate(
	ctx common.TxCtxInterface,
	req *ccpb.ReferenceCreateRequest,
) (res *ccpb.ReferenceCreateResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	if _, err = state.ReferenceCreate(ctx, req.GetRefKey()); err != nil {
		return nil, oops.Wrap(err)
	}

	return &ccpb.ReferenceCreateResponse{
		RefKey: req.GetRefKey(),
	}, nil
}

// ReferenceDelete deletes a reference
func (o ItemContractImpl) ReferenceDelete(
	ctx common.TxCtxInterface,
	req *ccpb.ReferenceDeleteRequest,
) (res *ccpb.ReferenceDeleteResponse, err error) {
	// Validate the request
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}
	if err = state.ReferenceDelete(ctx, req.GetRefKey()); err != nil {
		return nil, oops.Wrap(err)
	}

	return &ccpb.ReferenceDeleteResponse{
		RefKey: req.GetRefKey(),
	}, nil
}
