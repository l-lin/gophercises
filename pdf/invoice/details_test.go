package invoice

import "testing"

func TestInvoiceTotal(t *testing.T) {
	var tests = map[string]struct {
		given    []Unit
		expected float64
	}{
		"2 units": {
			given: []Unit{
				Unit{PricePerUnit: 100, UnitsPurchased: 1},
				Unit{PricePerUnit: 150, UnitsPurchased: 3},
			},
			expected: 5.50,
		},
		"price 0": {
			given: []Unit{
				Unit{PricePerUnit: 0, UnitsPurchased: 1},
				Unit{PricePerUnit: 0, UnitsPurchased: 3},
			},
			expected: 0,
		},
		"1 unit": {
			given: []Unit{
				Unit{PricePerUnit: 10, UnitsPurchased: 3},
			},
			expected: 0.3,
		},
		"no unit": {
			given:    []Unit{},
			expected: 0,
		},
		"gophercises example": {
			given: []Unit{
				Unit{
					PricePerUnit:   375, // in cents
					UnitsPurchased: 220,
				}, {
					PricePerUnit:   822, // in cents
					UnitsPurchased: 50,
				}, {
					PricePerUnit:   1455, // in cents
					UnitsPurchased: 3,
				},
			},
			expected: 1279.65,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			b := Bill{Units: tt.given}
			actual := b.InvoiceTotal()
			if actual != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}
		})
	}
}
