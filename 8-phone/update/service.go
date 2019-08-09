package update

import "github.com/l-lin/8-phone/phone"

// Repository provides an access to the phone storage
type Repository interface {
	Update(*phone.Phone)
}

// Service provides phone updating operations
type Service interface {
	Update(*phone.Phone)
}

type service struct {
	r Repository
}

// NewService creates a updating service
func NewService(r Repository) Service {
	return &service{r}
}

// Update phone
func (s *service) Update(p *phone.Phone) {
	s.r.Update(p)
}
