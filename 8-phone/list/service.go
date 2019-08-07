package list

import "github.com/l-lin/8-phone/phone"

// Repository provides an access to the task storage
type Repository interface {
	GetAll() []*phone.Phone
}

// Service provides phone listing operations
type Service interface {
	GetAll() []*phone.Phone
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
