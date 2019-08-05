package task

import (
	"fmt"
	"time"
)

// Task represents a TODO
type Task struct {
	ID            int
	Content       string
	Created       time.Time
	Completed     bool
	CompletedTime *time.Time
}

func (t *Task) String() string {
	return fmt.Sprintf("%d. %s", t.ID, t.Content)
}
