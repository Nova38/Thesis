package context

import (
	"github.com/charmbracelet/log"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nova38/thesis/lib/go/fabric/state"
	schema "github.com/nova38/thesis/lib/go/gen/ccbio/schema/v1"
)

type CollectionIDTypeHolder interface {
	GetId() *schema.Collection_Id
}
type CollectionHolder interface {
	GetCollection() *schema.Collection
}
type SpecimenIDHolder interface {
	GetId() *schema.Specimen_Id
}
type SpecimenHolder interface {
	GetSpecimen() *schema.Specimen
}

type SuggestionsIDHolder interface {
	GetId() *schema.SuggestedUpdate_Id
}
type SuggestionsHolder interface {
	GetId() *schema.SuggestedUpdate_Id
}

type LoggingTxContext interface {
	contractapi.TransactionContextInterface
	contractapi.SettableTransactionContextInterface
	GetLogger() *log.Logger
	SetLogger(logger *log.Logger) error
}

type TxContext interface {
	state.LogedTxCtxInterface

	GetUserId() (*schema.User_Id, error)
	GetUser() (*schema.User, error)

	// MakeLastModified Makes a LastModified object with the current user and timestamp
	MakeLastModified() (*schema.LastModified, error)

	// MakeClearedLastModified Makes a LastModified that is empty
	// This is used when we want to see if we should update a field.
	MakeClearedLastModified() *schema.LastModified
	//SetUser(user *schema.User) error

	GetCollection() (*schema.Collection, error)
	SetCollection(collection *schema.Collection) error
	InitViaRequest(req interface{}) error

	// GetRole requires a collection to be set
	GetRole() (schema.Role, error)
	//SetRole(role *schema.Role) error

	Authorize() (valid bool, err error)
	IsAuthorized() (valid bool, err error)

	GetAction() (*schema.Action, error)
	SetAction(action *schema.Action) error

	AddActionDomain(domain schema.Action_Domain)
}

type LastModifiedTxContext interface {
	MakeLastModified() (*schema.LastModified, error)
	MakeClearedLastModified() *schema.LastModified
}
