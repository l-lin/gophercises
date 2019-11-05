package get

import "github.com/l-lin/gophercises/secret/secret"

// Repository to get the secret
type Repository interface {
	Get(key, encodingKey string) (*secret.Secret, error)
}
