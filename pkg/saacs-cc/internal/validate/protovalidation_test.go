package validate_test

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	authpb "github.com/nova38/saacs/pkg/saacs-protos/saacs/auth/v0"
)

func TestCollections(t *testing.T) {
	testCases := []struct {
		desc       string
		collection *authpb.Collection
		valid      bool
	}{
		// Role Collection
		// ───────────────────────────────────────────────────────────────────────
		{
			desc: "Valid Role Collection",
			collection: &authpb.Collection{
				CollectionId: "ColId",
				Name:         "Collection",
				AuthType:     authpb.AuthType_AUTH_TYPE_ROLE,
				ItemTypes: []string{
					"saacs.auth.v0.UserCollectionRoles",
					"saacs.auth.v0.Role",
				},
				Default:        &authpb.Polices{},
				UseAuthParents: true,
			},
			valid: true,
		},
		{
			desc: "Valid Role Collection with other ItemTypes",
			collection: &authpb.Collection{
				CollectionId: "ColId",
				Name:         "Collection",
				AuthType:     authpb.AuthType_AUTH_TYPE_ROLE,
				ItemTypes: []string{
					"saacs.auth.v0.UserCollectionRoles",
					"saacs.auth.v0.Role",
					"saacs.sample.v0.Book",
				},
				Default:        &authpb.Polices{},
				UseAuthParents: true,
			},
			valid: true,
		},
		{
			desc: "InValid Role Collection (missing Role)",
			collection: &authpb.Collection{
				CollectionId: "ColId",
				Name:         "Collection",
				AuthType:     authpb.AuthType_AUTH_TYPE_ROLE,
				ItemTypes: []string{
					"saacs.auth.v0.Role",
				},
				Default:        &authpb.Polices{},
				UseAuthParents: true,
			},
			valid: false,
		},
		{
			desc: "InValid Role Collection (missing UserCollectionRoles)",
			collection: &authpb.Collection{
				CollectionId: "ColId",
				Name:         "Collection",
				AuthType:     authpb.AuthType_AUTH_TYPE_ROLE,
				ItemTypes: []string{
					"saacs.auth.v0.Role",
				},
				Default:        &authpb.Polices{},
				UseAuthParents: true,
			},
			valid: false,
		},
		// ───────────────────────────────────────────────────────────────────────
		// Identity Collections
		// ───────────────────────────────────────────────────────────────────────
		{
			desc: "Valid Identity Collection",
			collection: &authpb.Collection{
				CollectionId: "ColId",
				Name:         "Collection",
				AuthType:     authpb.AuthType_AUTH_TYPE_IDENTITY,
				ItemTypes: []string{
					"saacs.auth.v0.UserDirectMembership",
				},
				Default:        &authpb.Polices{},
				UseAuthParents: false,
			},
			valid: true,
		},
		{
			desc: "InValid Role Collection (missing UserDirectMembership)",
			collection: &authpb.Collection{
				CollectionId:   "ColId",
				Name:           "Collection",
				AuthType:       authpb.AuthType_AUTH_TYPE_ROLE,
				ItemTypes:      []string{},
				Default:        &authpb.Polices{},
				UseAuthParents: false,
			},
			valid: false,
		},
		// Embedded Role Collection
		// ───────────────────────────────────────────────────────────────────────
		{
			desc: "Valid Embed Role Collection",
			collection: &authpb.Collection{
				CollectionId: "ColId",
				Name:         "Collection",
				AuthType:     authpb.AuthType_AUTH_TYPE_GLOBAL_ROLE,
				ItemTypes: []string{
					"saacs.auth.v0.UserGlobalRoles",
					"saacs.auth.v0.Role",
				},
				Default:        &authpb.Polices{},
				UseAuthParents: true,
			},
			valid: true,
		},
		{
			desc: "InValid Role Collection (missing Role)",
			collection: &authpb.Collection{
				CollectionId: "ColId",
				Name:         "Collection",
				AuthType:     authpb.AuthType_AUTH_TYPE_GLOBAL_ROLE,
				ItemTypes: []string{
					"saacs.auth.v0.Role",
				},
				Default:        &authpb.Polices{},
				UseAuthParents: true,
			},
			valid: false,
		},
		{
			desc: "InValid Role Collection (missing UserCollectionRoles)",
			collection: &authpb.Collection{
				CollectionId: "ColId",
				Name:         "Collection",
				AuthType:     authpb.AuthType_AUTH_TYPE_GLOBAL_ROLE,
				ItemTypes: []string{
					"saacs.auth.v0.UserGlobalRoles",
				},
				Default:        &authpb.Polices{},
				UseAuthParents: true,
			},
			valid: false,
		},
		{
			desc: "Valid NoAUth Collection",
			collection: &authpb.Collection{
				CollectionId: "ColId",
				Name:         "Collection",
				AuthType:     authpb.AuthType_AUTH_TYPE_NONE,
				ItemTypes: []string{
					"saacs.auth.v0.UserCollectionRoles",
					"saacs.auth.v0.Role",
				},
				Default:        &authpb.Polices{},
				UseAuthParents: false,
			},
			valid: true,
		},
	}
	for _, tC := range testCases {
		v, err := protovalidate.New()
		if err != nil {
			t.Fatalf("Failed to create validator: %v", err)
		}

		t.Run(tC.desc, func(t *testing.T) {
			err := v.Validate(tC.collection)
			if err != nil && tC.valid {
				t.Fatalf("Expected valid collection, got error: %v", err)
			} else if err == nil && !tC.valid {
				t.Fatalf("Expected invalid collection, got no error")
			}
		})
	}
}
