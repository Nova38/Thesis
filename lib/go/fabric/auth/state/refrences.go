package state

import authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"

// ════════════════════════════════════════════════════════
// References Functions
// ════════════════════════════════════════════════════════

// ──────────────────────────────── Utils ────────────────────────────────────────

// // GetItemByReference returns the items referenced by the given reference
// func GetItemByReference(ctx TxCtxInterface, reference *authpb.Reference) (item *authpb.Item, err error) {
// 	return nil, nil
// }

// ──────────────────────────────── Query ────────────────────────────────────────

func GetReference(
	ctx TxCtxInterface,
	ref *authpb.Reference,
) (reference *authpb.Reference, err error) {
	// op := &authpb.Operation{
	// 	Action:       authpb.Action_ACTION_SUGGEST_VIEW,
	// 	CollectionId: s.GetPrimaryKey().GetCollectionId(),
	// 	ItemType:    s.GetPrimaryKey().GetItemType(),
	// 	Paths:        nil,
	// }

	return nil, nil
}

func PartialReferenceList(
	ctx TxCtxInterface,
	key string,
	numAttr int,
	bookmark string,
) (list []*authpb.Reference, mk string, err error) {
	return nil, "", nil
}

func ReferenceByType(
	ctx TxCtxInterface,
	itemType string,
	bookmark string,
) (list []*authpb.Reference, mk string, err error) {
	return nil, "", nil
}

func ReferenceListByItem(
	ctx TxCtxInterface,
	key *authpb.ItemKey,
	bookmark string,
) (list []*authpb.Reference, mk string, err error) {
	return nil, "", nil
}

func ReferenceByCollection(
	ctx TxCtxInterface,
	collectionId string,
	bookmark string,
) (list []*authpb.Reference, mk string, err error) {
	return nil, "", nil
}

// ──────────────────────────────── Invoke ───────────────────────────────────────

func ReferenceCreate(ctx TxCtxInterface, reference *authpb.Reference) (err error) {
	return nil
}

func ReferenceDelete(ctx TxCtxInterface, reference *authpb.Reference) (err error) {
	return nil
}

func ReferenceDeleteByItem(ctx TxCtxInterface, key *authpb.ItemKey) {
}
