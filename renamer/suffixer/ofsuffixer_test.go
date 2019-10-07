package suffixer

import "testing"

func TestOfSuffixer_Extract(t *testing.T) {
	type expected struct {
		base, ext string
		nb        int
	}
	s := &OfSuffixer{Max: 100}
	var tests = map[string]struct {
		given    string
		expected expected
	}{
		"basic": {
			given:    "foobar (2 of 100).txt",
			expected: expected{"foobar", ".txt", 2},
		},
		"without suffix": {
			given:    "foobar.txt",
			expected: expected{"foobar", ".txt", 0},
		},
		"without ext": {
			given:    "foobar (2 of 100)",
			expected: expected{"foobar", "", 2},
		},
		"without suffix & ext": {
			given:    "foobar",
			expected: expected{"foobar", "", 0},
		},
		"without base": {
			given:    " (1 of 100).txt",
			expected: expected{"", ".txt", 1},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actualBase, actualExt, actualNb := s.Extract(tt.given)
			if actualBase != tt.expected.base {
				t.Errorf("expected base %v, actual %v", tt.expected.base, actualBase)
			}
			if actualExt != tt.expected.ext {
				t.Errorf("expected ext %v, actual %v", tt.expected.ext, actualExt)
			}
			if actualNb != tt.expected.nb {
				t.Errorf("expected nb %v, actual %v", tt.expected.nb, actualNb)
			}
		})
	}
}
