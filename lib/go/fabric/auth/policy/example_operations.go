package policy

import (
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

var Policies = map[string]authpb.Operation{
	"Create Collection": {
		Action:       authpb.Action_ACTION_CREATE,
		CollectionId: "collection_id",
		ItemType:     "auth.Collection",
		Paths:        &fieldmaskpb.FieldMask{},
	},
	"Register User": {
		Action:       authpb.Action_ACTION_CREATE,
		CollectionId: "GLOBAL",
		ItemType:     "auth.User",
		Paths:        &fieldmaskpb.FieldMask{},
	},
	"Register Role": {
		Action:       authpb.Action_ACTION_CREATE,
		CollectionId: "collection_id",
		ItemType:     "auth.Role",
		Paths:        &fieldmaskpb.FieldMask{},
	},
	"Edit Role Permissions": {
		Action:       authpb.Action_ACTION_UPDATE,
		CollectionId: "collection_id",
		ItemType:     "auth.Role",
		Paths: &fieldmaskpb.FieldMask{
			Paths: []string{"ac"},
		},
	},
	"Edit User Membership": {
		Action:       authpb.Action_ACTION_REFERENCE_CREATE,
		CollectionId: "collection_id",
		ItemType:     "auth.UserMembership",
		Paths:        &fieldmaskpb.FieldMask{},
	},
}
