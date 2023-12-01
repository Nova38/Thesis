package common

const (
	// ItemTypes

	HiddenItemType     = "auth.HiddenTxList"
	SuggestionItemType = "auth.Suggestion"
	ReferenceItemType  = "auth.Reference"

	BootstrapKey = "bootstrap"

	// ──────────────────────────────────────────────────────────
	EnableHiddenTxValue   = "enable_hidden_tx"
	EnableSuggestionValue = "enable_suggestion"
	// ──────────────────────────────────────────────────────────

	// DefaultPageSize is the default page size for paginated queries
	DefaultPageSize int32 = 10000
)

var (
	EnableHiddenTx   = "enable_hidden_tx"
	EnableSuggestion = "enable_suggestion"
)
