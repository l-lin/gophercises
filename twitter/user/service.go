package user

import (
	"math/rand"
	"time"
)

// Service to manage users
type Service interface {
	FindAll() []User
	SaveAll([]User)
	PickWinner([]User) *User
}

// Repository to manage users
type Repository interface {
	FindAll() []User
	SaveAll([]User)
}

type service struct {
	r Repository
}

// NewService returns a new service to manage users
func NewService(r Repository) Service {
	return &service{r}
}

// FindAll users
func (s *service) FindAll() []User {
	return s.r.FindAll()
}

// SaveAll users
func (s *service) SaveAll(users []User) {
	s.r.SaveAll(users)
}

// PickWinner from a given slice of users
func (s *service) PickWinner(users []User) *User {
	if len(users) == 0 {
		return nil
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return &users[r.Intn(len(users))]
}
