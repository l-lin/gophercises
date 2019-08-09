package rm

// Repository provides an access to the phone storage
type Repository interface {
	Delete(id int)
}

// Service provides phone deleting operations
type Service interface {
	Delete(id int)
}

type service struct {
	r Repository
}

// NewService creates a listing service
func NewService(r Repository) Service {
	return &service{r}
}

// Delete a phone by its id
func (s *service) Delete(id int) {
	s.r.Delete(id)
}
