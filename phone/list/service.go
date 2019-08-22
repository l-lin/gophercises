package list

import "github.com/l-lin/gophercises/phone/phone"

// Repository provides an access to the task storage
type Repository interface {
	GetAll() []*phone.Phone
	Count(value string) int
}

// Service provides phone listing operations
type Service interface {
	GetAll() []*phone.Phone
	Count(value string) int
}

type service struct {
	r Repository
}

// NewService creates a listing service
func NewService(r Repository) Service {
	return &service{r}
}

// GetAll whether they are completed or not
func (s *service) GetAll() []*phone.Phone {
	return s.r.GetAll()
}

func (s *service) Count(value string) int {
	return s.r.Count(value)
}
