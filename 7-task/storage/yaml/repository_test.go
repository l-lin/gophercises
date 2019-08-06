package yaml

import (
	"testing"

	"github.com/l-lin/7-task/task"
)

func TestNextID(t *testing.T) {
	var tests = map[string]struct {
		expected int
		given    []*task.Task
	}{
		"basic": {
			expected: 4,
			given: []*task.Task{
				&task.Task{ID: 1},
				&task.Task{ID: 2},
				&task.Task{ID: 3},
			},
		},
		"not ordered elements": {
			expected: 5,
			given: []*task.Task{
				&task.Task{ID: 4},
				&task.Task{ID: 1},
				&task.Task{ID: 3},
			},
		},
		"no element": {
			expected: 1,
			given:    []*task.Task{},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := nextID(tt.given)
			if actual != tt.expected {
				t.Errorf("expected %d, actual %d", tt.expected, actual)
			}

		})
	}
}
