package complete

import "github.com/l-lin/7-task/task"

// Repository provides an access to the task storage
type Repository interface {
	GetCompleted() []*task.Task
}

// Service provides task adding operations
type Service interface {
	GetCompleted() []*task.Task
}

type service struct {
	r Repository
}

// NewService creates an completed listing service
func NewService(r Repository) Service {
	return &service{r}
}

// GetCompleted tasks
func (s *service) GetCompleted() []*task.Task {
	return s.r.GetCompleted()
}
