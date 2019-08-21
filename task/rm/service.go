package rm

// Repository provides an access to the task storage
type Repository interface {
	Remove(id int)
}

// Service provides task removing operations
type Service interface {
	Remove(id int)
}

type service struct {
	r Repository
}

// NewService creates an removing service
func NewService(r Repository) Service {
	return &service{r}
}

// Remove a task
func (s *service) Remove(id int) {
	s.r.Remove(id)
}
