package list

import "github.com/l-lin/task/task"

// Repository provides an access to the task storage
type Repository interface {
	GetAll() []*task.Task
	GetIncompletes() []*task.Task
}

// Service provides task listing operations
type Service interface {
	GetAll() []*task.Task
	GetIncompletes() []*task.Task
}

type service struct {
	r Repository
}

// NewService creates a listing service
func NewService(r Repository) Service {
	return &service{r}
}

// GetAll whether they are completed or not
func (s *service) GetAll() []*task.Task {
	return s.r.GetAll()
}

// GetIncompletes tasks
func (s *service) GetIncompletes() []*task.Task {
	return s.r.GetIncompletes()
}
