package policy

import (
	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

var _ = map[string]pb.Operation{
	"Create Collection": {
		Action:       pb.Action_ACTION_CREATE,
		CollectionId: "collection_id",
		ItemType:     "auth.Collection",
		Paths:        &fieldmaskpb.FieldMask{},
	},
	"Register User": {
		Action:       pb.Action_ACTION_CREATE,
		CollectionId: "GLOBAL",
		ItemType:     "auth.User",
		Paths:        &fieldmaskpb.FieldMask{},
	},
	"Register Role": {
		Action:       pb.Action_ACTION_CREATE,
		CollectionId: "collection_id",
		ItemType:     "auth.Role",
		Paths:        &fieldmaskpb.FieldMask{},
	},
	"Edit Role Permissions": {
		Action:       pb.Action_ACTION_UPDATE,
		CollectionId: "collection_id",
		ItemType:     "auth.Role",
		Paths: &fieldmaskpb.FieldMask{
			Paths: []string{"ac"},
		},
	},
}
