package phone

import "testing"

func TestNormalize(t *testing.T) {
	var tests = map[string]struct {
		expected string
		given    string
	}{
		"same pattern": {
			expected: "1234567890",
			given:    "1234567890",
		},
		"with spaces": {
			expected: "1234567891",
			given:    "123 456 7891",
		},
		"with spaces & parenthesises": {
			expected: "1234567892",
			given:    "(123) 456 7892",
		},
		"with spaces & parenthesises & hyphens": {
			expected: "1234567893",
			given:    "(123) 456-7893",
		},
		"with hyphens": {
			expected: "1234567894",
			given:    "123-456-7894",
		},
		"with parenthesises & hyphens": {
			expected: "1234567892",
			given:    "(123)456-7892",
		},
		"empty string": {
			expected: "",
			given:    "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			p := &Phone{Value: tt.given}
			actual := p.Normalize()
			if actual != tt.expected {
				t.Errorf("(%s): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}

}
