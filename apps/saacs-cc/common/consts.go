package common

const (
	// ItemTypes

	HiddenItemType     = "auth.HiddenTxList"
	SuggestionItemType = "auth.Suggestion"
	ReferenceItemType  = "auth.Reference"
	CollectionItemType = "auth.Collection"

	RoleItemType                 = "auth.Role"
	UserDirectMembershipItemType = "auth.UserDirectMembership"
	UserCollectionRolesItemType  = "auth.UserCollectionRoles"
	UserEmbeddedRoles            = "auth.UserEmbeddedRoles"

	BootstrapKey = "bootstrap"

	// ──────────────────────────────────────────────────────────

	// DefaultPageSize is the default page size for paginated queries
	DefaultPageSize int32 = 1000
)
