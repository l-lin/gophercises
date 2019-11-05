package set

import "github.com/l-lin/gophercises/secret/secret"

// Repository to set the secret
type Repository interface {
	Set(s *secret.Secret) error
}

// Service to set the secret
type Service struct {
	r Repository
}

// New service to set the secret
func New(r Repository) Service {
	return Service{r}
}

// Set the secret
func (service *Service) Set(s *secret.Secret) error {
	return service.r.Set(s)
}
