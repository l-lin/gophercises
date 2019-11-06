package get

import "github.com/l-lin/gophercises/secret/secret"

// Repository to get the secret
type Repository interface {
	Get(key string) (*secret.Secret, error)
}

// Service to get the secret
type Service struct {
	r Repository
}

// New service to get the secret
func New(r Repository) Service {
	return Service{r}
}

// Get the secret
func (service *Service) Get(key string) (*secret.Secret, error) {
	return service.r.Get(key)
}
