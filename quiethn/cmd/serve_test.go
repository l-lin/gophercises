package cmd

import "testing"

func TestComputeNbStoriesToFetch(t *testing.T) {
	var tests = map[string]struct {
		given    int
		expected int
	}{
		"30": {
			given:    30,
			expected: 37,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := computeNbStoriesToFetch(tt.given)
			if actual != tt.expected {
				t.Errorf("(%d): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}
