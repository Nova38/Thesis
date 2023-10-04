package schema

import "errors"

// Namespaces
// -------------------------

//goland:noinspection GoSnakeCaseUsage,GoSnakeCaseUsage,GoSnakeCaseUsage,GoSnakeCaseUsage
const (
	NS_USER            = "User"
	NS_COL             = "Collection"
	NS_SPECIMEN        = "Specimen"
	NS_SUGGESTEDUPDATE = "SuggestedUpdate"
)

func (u *User) Namespace() string {
	return NS_USER
}
func (c *Collection) Namespace() string {
	return NS_COL
}
func (s *Specimen) Namespace() string {
	return NS_SPECIMEN
}

func (s *SuggestedUpdate) Namespace() string {
	return NS_SUGGESTEDUPDATE
}

func (u *User) Key() (attr []string, err error) {
	if id := u.GetId(); id != nil {
		return []string{id.GetMspId(), id.GetId()}, nil
	}
	return nil, errors.New("User.Id is nil")
}

func (c *Collection) Key() (attr []string, err error) {
	if id := c.GetId(); id != nil {
		return []string{id.GetCollectionId()}, nil
	}
	return nil, errors.New("Collection.Id is nil")
}

func (s *Specimen) Key() (attr []string, err error) {
	if id := s.GetId(); id != nil {
		return []string{id.GetCollectionId(), id.GetId()}, nil
	}
	return nil, errors.New("Specimen.Id is nil")
}

func (s *SuggestedUpdate) Key() (attr []string, err error) {

	if id := s.GetId(); id != nil {
		if id.GetSpecimenId() == nil {
			return nil, errors.New("SuggestedUpdate.Id.SpecimenId is nil")
		}
		return []string{id.GetSpecimenId().GetCollectionId(), id.GetSpecimenId().GetId(), id.GetId()}, nil
	}
	return nil, errors.New("SuggestedUpdate.Id is nil")
}

// Request Types

type Requests interface {
	CollectionCreateRequest | CollectionUpdateRequest | GetCollectionRequest
}
