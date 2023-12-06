package actions

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/nova38/thesis/packages/saacs/common"
	"github.com/nova38/thesis/packages/saacs/state"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ════════════════════════════════════════════════════════
// References Functions
// ════════════════════════════════════════════════════════

// ──────────────────────────────── Query ────────────────────────────────────────

func GetReference(
	ctx common.TxCtxInterface,
	ref *authpb.ReferenceKey,
) (reference *authpb.Reference, err error) {
	if err = authRef(ctx, ref, authpb.Action_ACTION_REFERENCE_VIEW); err != nil {
		return nil, oops.Wrap(err)
	}

	return referenceGetPacked(ctx, ref)
}

// PartialReferenceKeysList gets the reference keys for the given reference key
//
// Checks Auth Directly
func PartialReferenceKeysList(
	ctx common.TxCtxInterface,
	ref *authpb.ReferenceKey,
	numAttr int,
	bookmark string,
) (list []*authpb.ReferenceKey, mk string, err error) {
	// TODO: Implement PartialReferenceKeysList function

	ctx.GetLogger().
		Debug("PartialReferenceKeysList", slog.Group("args", "ref", ref, "numAttr", numAttr, "bookmark", bookmark))

	if err = authRef(ctx, ref, authpb.Action_ACTION_REFERENCE_VIEW); err != nil {
		return nil, mk, oops.Wrap(err)
	}

	// we only need the first key if the state is actually valid
	attr, _, err := common.MakeRefKeyAttrs(ref)
	if err != nil {
		return nil, mk, oops.Wrap(err)
	}

	if len(attr) == 0 || len(attr) < numAttr {
		return nil, mk, oops.Wrap(common.RequestInvalid)
	}

	attr = lo.DropRight(attr, numAttr)
	ctx.GetLogger().
		Info("PartialReferenceKeysList",
			slog.Group(
				"Key", "ItemType", common.ReferenceItemType,
				slog.Int("numAttr", numAttr),
				slog.Any("attr", attr),
				slog.Group(
					"Paged",
					"Bookmark", bookmark,
					"PageSize", strconv.Itoa(int(ctx.GetPageSize())),
				),
			),
		)
	results, metadata, err := ctx.GetStub().
		GetStateByPartialCompositeKeyWithPagination(
			common.ReferenceItemType,
			attr,
			ctx.GetPageSize(),
			bookmark,
		)
	if err != nil {
		return nil, "", err
	}
	defer func(results shim.StateQueryIteratorInterface) {
		err := results.Close()
		if err != nil {
			ctx.GetLogger().Error("GetPartialSuggestedKeyList", "Error", err)
		}
	}(results)

	for results.HasNext() {
		queryResponse, err := results.Next()
		if err != nil {
			return nil, "", err
		}
		refKey := new(authpb.ReferenceKey)

		if err := json.Unmarshal(queryResponse.GetValue(), refKey); err != nil {
			return nil, "", err
		}

		list = append(list, refKey)
	}

	if metadata == nil {
		return nil, "", fmt.Errorf("metadata is nil")
	}

	return list, metadata.GetBookmark(), nil
}

func ReferenceKeysByItem(
	ctx common.TxCtxInterface,
	key *authpb.ItemKey,
	bookmark string,
) (list []*authpb.ReferenceKey, mk string, err error) {
	refKey := &authpb.ReferenceKey{
		Key1: key,
	}

	if err = authRef(ctx, refKey, authpb.Action_ACTION_REFERENCE_VIEW); err != nil {
		return nil, mk, oops.Wrap(err)
	}

	return PartialReferenceKeysList(ctx, refKey, int(ctx.GetPageSize()), bookmark)
}

// ReferencesByItem gets the references for the given item
// Indirectly checks Auth via ReferenceKeysByItem
func ReferencesByItem(
	ctx common.TxCtxInterface,
	key *authpb.ItemKey,
	bookmark string,
) (list []*authpb.Reference, mk string, err error) {
	refKeys, mk, err := ReferenceKeysByItem(ctx, key, bookmark)
	if err != nil {
		return nil, mk, oops.Wrap(err)
	}

	for _, refKey := range refKeys {
		ref, err := GetReference(ctx, refKey)
		if err != nil {
			return nil, mk, oops.Wrap(err)
		}
		list = append(list, ref)
	}

	return list, mk, nil
}

// ReferenceKeysByCollection gets the references for the given collection
// Indirectly checks Auth via ReferenceKeysByItem
func ReferenceKeysByCollection(
	ctx common.TxCtxInterface,
	key1 *authpb.ItemKey,
	collectionId string,
	bookmark string,
) (list []*authpb.ReferenceKey, mk string, err error) {
	if key1 == nil || collectionId == "" {
		return nil, mk, oops.Wrap(common.RequestInvalid)
	}

	refKey := &authpb.ReferenceKey{
		Key1: key1,
		Key2: &authpb.ItemKey{
			CollectionId: collectionId,
		},
	}

	if err = authRef(ctx, refKey, authpb.Action_ACTION_REFERENCE_VIEW); err != nil {
		return nil, mk, oops.Wrap(err)
	}

	return PartialReferenceKeysList(ctx, refKey, int(ctx.GetPageSize()), bookmark)
}

// ──────────────────────────────── Invoke ───────────────────────────────────────

func ReferenceCreate(
	ctx common.TxCtxInterface,
	refKey *authpb.ReferenceKey,
) (reference *authpb.Reference, err error) {
	ctx.GetLogger().Debug("ReferenceCreate", slog.Group("args", "refKey", refKey))
	if err = authRef(ctx, refKey, authpb.Action_ACTION_DELETE); err != nil {
		return nil, oops.Wrap(err)
	}

	// See if the reference already exists

	if objExist, err := referencedObjectsExist(ctx, refKey); err != nil {
		return nil, oops.Wrap(err)
	} else if !objExist {
		return nil, oops.Wrap(common.KeyNotFound)
	}

	if refExist, err := ReferencesExist(ctx, refKey); err != nil {
		return nil, oops.Wrap(err)
	} else if refExist {
		return nil, oops.Wrap(common.AlreadyExists)
	}

	k1, k2, err := common.MakeRefKeys(refKey)
	if err != nil {
		// ctx.GetLogger().Error("Error making reference keys", "ref", reference, slog.Group("k1", k1, "k2", k2))
		return nil, oops.Wrap(err)
	}

	// Put the reference into the ledger
	// Needed because an empty value is not allowed
	value := []byte{0x00}

	if err = ctx.GetStub().PutState(k1, value); err != nil {
		return nil, oops.Wrap(err)
	}
	if err = ctx.GetStub().PutState(k2, value); err != nil {
		return nil, oops.Wrap(err)
	}

	// TODO: Get the items from the ledger to pack the references

	return nil, nil
}

// ReferenceDelete deletes the reference from the ledger
//
// Checks Auth Directly
func ReferenceDelete(ctx common.TxCtxInterface, reference *authpb.ReferenceKey) error {
	ctx.GetLogger().Debug("ReferenceDelete", slog.Group("args", "ref", reference.String()))

	if err := authRef(ctx, reference, authpb.Action_ACTION_DELETE); err != nil {
		return oops.Wrap(err)
	}

	if objExist, err := referencedObjectsExist(ctx, reference); err != nil {
		return oops.Wrap(err)
	} else if !objExist {
		return oops.Wrap(common.KeyNotFound)
	}

	if refExist, err := ReferencesExist(ctx, reference); err != nil {
		return oops.Wrap(err)
	} else if !refExist {
		return oops.Wrap(common.KeyNotFound)
	}

	// Delete the reference
	k1, k2, err := common.MakeRefKeys(reference)
	ctx.GetLogger().Debug("ReferenceDelete", slog.Group("keys", "Key1", k1, "Key2", k2))
	if err != nil {
		return oops.Wrap(err)
	}

	if err = ctx.GetStub().DelState(k1); err != nil {
		return oops.Wrap(err)
	}
	if err = ctx.GetStub().DelState(k2); err != nil {
		return oops.Wrap(err)
	}

	return nil
}

// ──────────────────────────────── Utils ──────────────────────────────────────────

// ReferencesExist checks if the references for the given reference to exist
//
// Does not check Auth
func ReferencesExist(
	ctx common.TxCtxInterface,
	reference *authpb.ReferenceKey,
) (exists bool, err error) {
	// See if the reference already exists
	k1, k2, err := common.MakeRefKeys(reference)
	if err != nil {
		return false, oops.Wrap(err)
	}

	// See if key1 exists
	if !state.KeyExists(ctx, k1) || !state.KeyExists(ctx, k2) {
		ctx.GetLogger().
			Info("References for objects already exists with given keys does not exist", slog.Group("keys", "Key1", k1, "Key2", k2))
		return false, oops.Wrap(common.KeyNotFound)
	}

	return true, nil
}

// ──────────────────────────────── Internal Utils ──────────────────────────────────────────

// authRef checks if the user is authorized to perform the given action on the given reference
func authRef(
	ctx common.TxCtxInterface,
	ref *authpb.ReferenceKey,
	action authpb.Action,
) error {
	// ctx.GetLogger().Debug("authRef", slog.Group("ref", ref, "action", action))

	if auth, err := ctx.Authorize(RefToOp(ref, action)); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	return nil
}

// RefToOp converts a reference to an operation
func RefToOp(ref *authpb.ReferenceKey, action authpb.Action) (ops []*authpb.Operation) {
	// ctx.GetLogger().Debug("RefToOp", slog.Group("ref", ref, "action", action))
	return []*authpb.Operation{
		{
			Action:       action,
			CollectionId: ref.GetKey1().GetCollectionId(),
			ItemType:     ref.GetKey1().GetItemType(),
			Paths:        &fieldmaskpb.FieldMask{},
		},
		{
			Action:       action,
			CollectionId: ref.GetKey2().GetCollectionId(),
			ItemType:     ref.GetKey2().GetItemType(),
			Paths:        &fieldmaskpb.FieldMask{},
		},
	}
}

// referencedObjectsExist checks if the objects referenced by the given reference exist
// Does not check Auth
func referencedObjectsExist(
	ctx common.TxCtxInterface,
	reference *authpb.ReferenceKey,
) (exists bool, err error) {
	// See if the objects exist

	err1, ok := lo.TryWithErrorValue(func() error {

		k1 := common.MakeStateKey(reference.GetKey1())
		k2 := common.MakeStateKey(reference.GetKey1())

		if !state.KeyExists(ctx, k1) || !state.KeyExists(ctx, k2) {
			ctx.GetLogger().
				Info("Items with given keys does not exist", slog.Group("keys", "Key1", k1, "Key2", k2))
			return oops.Wrap(common.KeyNotFound)
		}

		return nil
	})

	if err, errOk := err1.(oops.OopsError); !errOk {
		return false, oops.Wrap(err)
	}

	return ok, oops.Wrap(err)
}

// referenceGetPacked gets the referenced objects for the given reference key
// Assumes that the reference exists
func referenceGetPacked(
	ctx common.TxCtxInterface,
	key *authpb.ReferenceKey,
) (ref *authpb.Reference, err error) {
	item1, item2, err := common.ReferenceKeyToItems(key)
	if err != nil {
		return nil, oops.Wrap(err)
	}

	if err := state.Get(ctx, item1.StateKey(), item1); err != nil {
		ctx.GetLogger().Error("Error getting item1",
			slog.Any("ref", ref),
			slog.Group("item1", item1))
		return nil, oops.With("ref", ref).Wrap(err)
	}

	if err := state.Get(ctx, item2.StateKey(), item2); err != nil {
		ctx.GetLogger().Error("Error getting item1",
			slog.Any("ref", ref),
			slog.Group("item1", item2))
		return nil, oops.With("ref", ref).Wrap(err)
	}

	return common.PackReference(key, item1, item2)
}
