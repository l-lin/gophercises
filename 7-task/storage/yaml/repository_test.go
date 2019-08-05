package yaml

import (
	"testing"
	"time"

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

func TestSameDay(t *testing.T) {
	type arg struct {
		t1, t2 time.Time
	}
	var tests = map[string]struct {
		expected bool
		given    arg
	}{
		"same day": {
			expected: true,
			given: arg{
				t1: time.Date(2019, time.August, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2019, time.August, 5, 10, 0, 0, 0, time.UTC),
			},
		},
		"different day": {
			expected: false,
			given: arg{
				t1: time.Date(2019, time.August, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2019, time.August, 6, 10, 0, 0, 0, time.UTC),
			},
		},
		"different location": {
			expected: true,
			given: arg{
				t1: time.Date(2019, time.August, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2019, time.August, 6, 0, 0, 0, 0, time.FixedZone("GMT", 2)),
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := sameDay(tt.given.t1, tt.given.t2)
			if actual != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}

		})
	}
}
