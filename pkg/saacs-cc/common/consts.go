package common

const (
	// ItemTypes

	HiddenItemType     = "saacs.common.v0.HiddenTxList"
	SuggestionItemType = "saacs.common.v0.Suggestion"

	CollectionItemType           = "saacs.auth.v0.Collection"
	RoleItemType                 = "saacs.auth.v0.Role"
	UserDirectMembershipItemType = "saacs.auth.v0.UserDirectMembership"
	UserCollectionRolesItemType  = "saacs.auth.v0.UserCollectionRoles"
	UserGlobalRoles              = "saacs.auth.v0.UserGlobalRoles"

	BootstrapKey = "bootstrap"
	USERCOLID    = "USERS"
	// ──────────────────────────────────────────────────────────

	// DefaultPageSize is the default page size for paginated queries
	DefaultPageSize int32 = 1000
)
