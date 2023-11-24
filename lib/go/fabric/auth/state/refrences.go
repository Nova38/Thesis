package state

import (
	"log/slog"

	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
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
) (reference *authpb.ReferenceKey, err error) {
	if err = AuthRef(ctx, ref, authpb.Action_ACTION_REFERENCE_VIEW); err != nil {
		return nil, oops.Wrap(err)
	}

	return nil, nil
}

func PartialReferenceList(
	ctx TxCtxInterface,
	key string,
	numAttr int,
	bookmark string,
) (list []*authpb.ReferenceKey, mk string, err error) {
	// TODO: Implement PartialReferenceList function

	return nil, "", nil
}

func ReferenceListByItem(
	ctx TxCtxInterface,
	key *authpb.ItemKey,
	bookmark string,
) (list []*authpb.ReferenceKey, mk string, err error) {
	// todo: Implement ReferenceListByItem function
	panic("not implemented")
	return nil, "", nil
}

func ReferenceByCollection(
	ctx TxCtxInterface,
	collectionId string,
	bookmark string,
) (list []*authpb.ReferenceKey, mk string, err error) {
	// todo: Implement ReferenceByCollection function
	panic("not implemented")

	return nil, "", nil
}

// ──────────────────────────────── Invoke ───────────────────────────────────────

func ReferenceCreate(ctx TxCtxInterface, reference *authpb.ReferenceKey) (err error) {
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
	op := RefToOp(ref, action)

	if auth, err := ctx.Authorize([]*authpb.Operation{op}); !auth || err != nil {
		return oops.Wrap(common.UserPermissionDenied)
	}

	return nil
}

// RefToOp converts a reference to an operation
func RefToOp(ref *authpb.ReferenceKey, action authpb.Action) (op *authpb.Operation) {
	return &authpb.Operation{
		Action:            action,
		CollectionId:      ref.GetCollectionId(),
		ItemType:          ref.GetKey1().GetItemType(),
		SecondaryItemType: ref.GetKey2().GetItemType(),
		Paths:             &fieldmaskpb.FieldMask{},
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

// ReferenceDeleteByItem deletes all the references for the given item
func ReferenceDeleteByItem(ctx TxCtxInterface, key *authpb.ItemKey) (err error) {
	// Get all the references for the given item

	refs, _, err := ReferenceListByItem(ctx, key, "")
	if err != nil {
		return oops.Wrap(err)
	}

	for _, ref := range refs {
		if err = ReferenceDelete(ctx, ref); err != nil {
			return oops.Wrap(err)
		}
	}

	// Delete all the references for the given item, in both directions
	// todo: Implement ReferenceDeleteByItem function
	panic("not implemented")
}

// ReferenceKeyToItems converts a reference key to the item object

func ReferenceKeyToItems(ref *authpb.ReferenceKey) (item1 common.ItemInterface, item2 common.ItemInterface, err error) {
	// TODO: Implement ReferenceKeyToItems function

	panic("not implemented")
	return nil, nil, nil
}

// PackReference packs the items into anypb and returns the reference
func PackReference(ref *authpb.ReferenceKey, item1 common.ItemInterface, item2 common.ItemInterface) (reference authpb.Reference, err error) {
	panic("not implemented")
}
