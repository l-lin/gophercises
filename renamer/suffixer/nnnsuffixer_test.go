package suffixer

import "testing"

func TestNnnSuffixer_Extract(t *testing.T) {
	type expected struct {
		base, ext string
		nb        int
	}
	s := &NnnSuffixer{3}
	var tests = map[string]struct {
		given    string
		expected expected
	}{
		"basic": {
			given:    "n_008.txt",
			expected: expected{"n", ".txt", 8},
		},
		"without suffix": {
			given:    "n.txt",
			expected: expected{"n", ".txt", 0},
		},
		"without ext": {
			given:    "n_002",
			expected: expected{"n", "", 2},
		},
		"without suffix & ext": {
			given:    "foobar",
			expected: expected{"foobar", "", 0},
		},
		"without base": {
			given:    "_002.txt",
			expected: expected{"", ".txt", 2},
		},
		"not same numbers of numbers": {
			given:    "n_0002.txt",
			expected: expected{"n_0002", ".txt", 0},
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
