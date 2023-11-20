package state

import authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"

// ════════════════════════════════════════════════════════
// References Functions
// ════════════════════════════════════════════════════════

// ──────────────────────────────── Utils ────────────────────────────────────────

// // GetObjectByReference returns the objects referenced by the given reference
// func GetObjectByReference(ctx TxCtxInterface, reference *authpb.Reference) (object *authpb.Object, err error) {
// 	return nil, nil
// }

// ──────────────────────────────── Query ────────────────────────────────────────

func GetReference(ctx TxCtxInterface, ref *authpb.Reference) (reference *authpb.Reference, err error) {
	// op := &authpb.Operation{
	// 	Action:       authpb.Action_ACTION_SUGGEST_VIEW,
	// 	CollectionId: s.GetPrimaryKey().GetCollectionId(),
	// 	Namespace:    s.GetPrimaryKey().GetObjectType(),
	// 	Paths:        nil,
	// }

	return nil, nil
}

func PartialReferenceList(ctx TxCtxInterface, key string, numAttr int, bookmark string) (list []*authpb.Reference, mk string, err error) {
	return nil, "", nil
}

func ReferenceListByObject(ctx TxCtxInterface, key *authpb.ObjectKey, bookmark string) (list []*authpb.Reference, mk string, err error) {
	return nil, "", nil
}

func ReferenceByCollection(ctx TxCtxInterface, collectionId string, bookmark string) (list []*authpb.Reference, mk string, err error) {
	return nil, "", nil
}

// ──────────────────────────────── Invoke ───────────────────────────────────────

func ReferenceCreate(ctx TxCtxInterface, reference *authpb.Reference) (err error) {
	return nil
}

func ReferenceDelete(ctx TxCtxInterface, reference *authpb.Reference) (err error) {
	return nil
}

func ReferenceDeleteByObject(ctx TxCtxInterface, key *authpb.ObjectKey) {
}
