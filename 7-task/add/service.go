package add

import "github.com/l-lin/7-task/task"

// Repository provides an access to the task storage
type Repository interface {
	Add(t *task.Task)
}

// Service provides task adding operations
type Service interface {
	Add(t *task.Task)
}

type service struct {
	r Repository
}

// NewService creates an adding service
func NewService(r Repository) Service {
	return &service{r}
}

// Add a task
func (s *service) Add(t *task.Task) {
	s.r.Add(t)
}
