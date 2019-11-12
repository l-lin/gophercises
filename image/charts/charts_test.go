package charts

import "testing"

func Test_computeMaxHeight(t *testing.T) {
	var tests = map[string]struct {
		given    []int
		expected int
	}{
		"basic": {
			given:    []int{10, 20, 5, 2},
			expected: 20 * BarHeightCoeff,
		},
		"single value": {
			given:    []int{1},
			expected: 1 * BarHeightCoeff,
		},
		"empty slice": {
			given:    []int{},
			expected: 0,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := computeMaxHeight(tt.given)
			if actual != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}
		})
	}
}

func Test_computeMaxWidth(t *testing.T) {
	var tests = map[string]struct {
		given    []int
		expected int
	}{
		"basic": {
			given:    []int{10, 20, 5, 2},
			expected: 4*BarWidth + 3*SepWidth,
		},
		"one value": {
			given:    []int{10},
			expected: BarWidth,
		},
		"two values": {
			given:    []int{10, 20},
			expected: 2*BarWidth + SepWidth,
		},
		"empty slice": {
			given:    []int{},
			expected: 0,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := computeMaxWidth(tt.given)
			if actual != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}
		})
	}
}
