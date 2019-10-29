package user

import "testing"

func TestUniquUsers(t *testing.T) {
	var tests = map[string]struct {
		given    []User
		expected []User
	}{
		"no double": {
			given: []User{
				User{ID: "123"},
				User{ID: "321"},
			},
			expected: []User{
				User{ID: "123"},
				User{ID: "321"},
			},
		},
		"one double": {
			given: []User{
				User{ID: "123"},
				User{ID: "321"},
				User{ID: "123"},
			},
			expected: []User{
				User{ID: "123"},
				User{ID: "321"},
			},
		},
		"two doubles": {
			given: []User{
				User{ID: "123"},
				User{ID: "321"},
				User{ID: "123"},
				User{ID: "321"},
			},
			expected: []User{
				User{ID: "123"},
				User{ID: "321"},
			},
		},
		"one triple": {
			given: []User{
				User{ID: "123"},
				User{ID: "321"},
				User{ID: "123"},
				User{ID: "123"},
			},
			expected: []User{
				User{ID: "123"},
				User{ID: "321"},
			},
		},
		"empty slice": {
			given:    []User{},
			expected: []User{},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := UniqueUsers(tt.given)
			if len(tt.expected) != len(actual) {
				t.Errorf("expected %v slice size, actual %v slice size", len(tt.expected), len(actual))
			}
			for i := 0; i < len(tt.expected); i++ {
				if actual[i] != tt.expected[i] {
					t.Errorf("expected %v, actual %v", tt.expected[i], actual[i])
				}
			}
		})
	}
}
