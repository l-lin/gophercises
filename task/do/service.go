package do

// Repository provides an access to the task storage
type Repository interface {
	Do(id int)
}

// Service provides task complete operations
type Service interface {
	Do(id int)
}

type service struct {
	r Repository
}

// NewService creates an task complete service
func NewService(r Repository) Service {
	return &service{r}
}

// Do a task
func (s *service) Do(id int) {
	s.r.Do(id)
}
