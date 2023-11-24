package state

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ════════════════════════════════════════════════════════
// References Functions
// ════════════════════════════════════════════════════════

// ──────────────────────────────── Query ────────────────────────────────────────

func GetReference(
	ctx TxCtxInterface,
	ref *authpb.ReferenceKey,
) (reference *authpb.Reference, err error) {
	if err = AuthRef(ctx, ref, authpb.Action_ACTION_REFERENCE_VIEW); err != nil {
		return nil, oops.Wrap(err)
	}

	return ReferenceGetPacked(ctx, ref)
}

func PartialReferenceList(
	ctx TxCtxInterface,
	ref *authpb.ReferenceKey,
	numAttr int,
	bookmark string,
) (list []*authpb.ReferenceKey, mk string, err error) {
	// TODO: Implement PartialReferenceList function

	ctx.GetLogger().Debug("PartialReferenceList", slog.Group("args", "ref", ref, "numAttr", numAttr, "bookmark", bookmark))

	if err = AuthRef(ctx, ref, authpb.Action_ACTION_REFERENCE_VIEW); err != nil {
		return nil, mk, oops.Wrap(err)
	}

	// we only need the first key if the state is actually valid
	attr, _, err := MakeRefKeyAttrs(ref)
	if err != nil {
		return nil, mk, oops.Wrap(err)
	}

	if len(attr) == 0 || len(attr) < numAttr {
		return nil, mk, oops.Wrap(common.RequestInvalid)
	}

	attr = lo.DropRight(attr, numAttr)
	ctx.GetLogger().
		Info("PartialReferenceList",
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

func ReferenceListByItem(
	ctx TxCtxInterface,
	key *authpb.ItemKey,
	bookmark string,
) (list []*authpb.ReferenceKey, mk string, err error) {
	refKey := &authpb.ReferenceKey{
		Key1: key,
	}

	if err = AuthRef(ctx, refKey, authpb.Action_ACTION_REFERENCE_VIEW); err != nil {
		return nil, mk, oops.Wrap(err)
	}

	return PartialReferenceList(ctx, refKey, int(ctx.GetPageSize()), bookmark)
}

func ReferenceByCollection(
	ctx TxCtxInterface,
	collectionId string,
	bookmark string,
) (list []*authpb.ReferenceKey, mk string, err error) {
	refKey := &authpb.ReferenceKey{
		Key1: &authpb.ItemKey{
			CollectionId: collectionId,
		},
	}

	if err = AuthRef(ctx, refKey, authpb.Action_ACTION_REFERENCE_VIEW); err != nil {
		return nil, mk, oops.Wrap(err)
	}

	return PartialReferenceList(ctx, refKey, int(ctx.GetPageSize()), bookmark)
}

// ──────────────────────────────── Invoke ───────────────────────────────────────

func ReferenceCreate(ctx TxCtxInterface, reference *authpb.ReferenceKey) (err error) {
	ctx.GetLogger().Debug("ReferenceCreate", slog.Group("ref", reference))
	if err = AuthRef(ctx, reference, authpb.Action_ACTION_DELETE); err != nil {
		return oops.Wrap(err)
	}

	// See if the reference already exists

	if objExist, err := ReferencedObjectsExist(ctx, reference); err != nil {
		return oops.Wrap(err)
	} else if !objExist {
		return oops.Wrap(common.KeyNotFound)
	}

	if refExist, err := ReferencesExist(ctx, reference); err != nil {
		return oops.Wrap(err)
	} else if refExist {
		return oops.Wrap(common.AlreadyExists)
	}

	k1, k2, err := MakeRefKeys(reference)
	if err != nil {
		// ctx.GetLogger().Error("Error making reference keys", "ref", reference, slog.Group("k1", k1, "k2", k2))
		return oops.Wrap(err)
	}

	// Put the reference into the ledger
	// Needed because an empty value is not allowed
	value := []byte{0x00}

	if err = ctx.GetStub().PutState(k1, value); err != nil {
		return oops.Wrap(err)
	}
	if err = ctx.GetStub().PutState(k2, value); err != nil {
		return oops.Wrap(err)
	}

	return nil
}

func ReferenceDelete(ctx TxCtxInterface, reference *authpb.ReferenceKey) (err error) {
	ctx.GetLogger().Debug("ReferenceDelete", slog.Group("args", "ref", reference.String()))
	if err = AuthRef(ctx, reference, authpb.Action_ACTION_DELETE); err != nil {
		return oops.Wrap(err)
	}

	if objExist, err := ReferencedObjectsExist(ctx, reference); err != nil {
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
	k1, k2, err := MakeRefKeys(reference)
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

// AuthRef checks if the user is authorized to perform the given action on the given reference
func AuthRef(ctx TxCtxInterface, ref *authpb.ReferenceKey, action authpb.Action) (err error) {
	// ctx.GetLogger().Debug("AuthRef", slog.Group("ref", ref, "action", action))

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
			Action:            action,
			CollectionId:      ref.GetKey1().GetCollectionId(),
			ItemType:          ref.GetKey1().GetItemType(),
			SecondaryItemType: ref.GetKey2().GetItemType(),
			Paths:             &fieldmaskpb.FieldMask{},
		},
		{
			Action:            action,
			CollectionId:      ref.GetKey2().GetCollectionId(),
			ItemType:          ref.GetKey2().GetItemType(),
			SecondaryItemType: ref.GetKey1().GetItemType(),
			Paths:             &fieldmaskpb.FieldMask{},
		},
	}
}

// ReferencedObjectsExist checks if the objects referenced by the given reference exist
func ReferencedObjectsExist(
	ctx TxCtxInterface,
	reference *authpb.ReferenceKey,
) (exists bool, err error) {
	// See if the objects exist
	k1, err := MakeItemKeyPrimary(reference.GetKey1())
	if err != nil {
		return false, oops.Wrap(err)
	}

	k2, err := MakeItemKeyPrimary(reference.GetKey2())
	if err != nil {
		return false, oops.Wrap(err)
	}

	if !Exists(ctx, k1) || !Exists(ctx, k2) {
		ctx.GetLogger().
			Info("Items with given keys does not exist", slog.Group("keys", "Key1", k1, "Key2", k2))
		return false, oops.Wrap(common.KeyNotFound)
	}

	return true, nil
}

// ReferencesExist checks if the references for the given reference exist
func ReferencesExist(ctx TxCtxInterface, reference *authpb.ReferenceKey) (exists bool, err error) {
	// See if the reference already exists
	k1, k2, err := MakeRefKeys(reference)
	if err != nil {
		return false, oops.Wrap(err)
	}

	// See if key1 exists
	if !Exists(ctx, k1) || !Exists(ctx, k2) {
		ctx.GetLogger().
			Info("References for objects already exists with given keys does not exist", slog.Group("keys", "Key1", k1, "Key2", k2))
		return false, oops.Wrap(common.KeyNotFound)
	}

	return true, nil
}

func ReferenceGetPacked(ctx TxCtxInterface, key *authpb.ReferenceKey) (ref *authpb.Reference, err error) {
	item1, item2, err := ReferenceKeyToItems(key)
	if err != nil {
		return nil, oops.Wrap(err)
	}

	if err = Get(ctx, item1); err != nil {
		ctx.GetLogger().Error("Error getting item1", "ref", ref, slog.Group("item1", item1))
		return nil, oops.With("ref", ref).Wrap(err)
	}

	if err = Get(ctx, item2); err != nil {
		ctx.GetLogger().Error("Error getting item2", "ref", ref, slog.Group("item2", item1))
		return nil, oops.Wrap(err)
	}

	return PackReference(key, item1, item2)
}

// ReferenceDeleteByItem deletes all the references for the given item
func ReferenceDeleteByItem(ctx TxCtxInterface, key *authpb.ItemKey) (err error) {
	// Get all the references for the given item

	refs, _, err := ReferenceListByItem(ctx, key, "")
	if err != nil {
		return oops.Wrap(err)
	}
	ctx.GetLogger().Debug("ReferenceDeleteByItem", slog.Group("To Delete", "refs", refs))

	for _, ref := range refs {
		if err = ReferenceDelete(ctx, ref); err != nil {
			return oops.Wrap(err)
		}
	}

	return nil
}
