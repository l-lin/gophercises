package invoice

import "testing"

func TestAmount(t *testing.T) {
	var tests = map[string]struct {
		given    Unit
		expected float64
	}{
		"basic": {
			given:    Unit{UnitsPurchased: 2, PricePerUnit: 262},
			expected: 5.24,
		},
		"zero unit": {
			given:    Unit{UnitsPurchased: 0, PricePerUnit: 262},
			expected: 0,
		},
		"zero price": {
			given:    Unit{UnitsPurchased: 10, PricePerUnit: 0},
			expected: 0,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := tt.given.Amount()
			if actual != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}
		})
	}
}
