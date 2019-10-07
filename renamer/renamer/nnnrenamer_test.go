package renamer

import "testing"

func TestNnnRenamer_transform(t *testing.T) {
	type expected struct {
		transformed string
		hasErr      bool
	}
	r := &NnnRenamer{3}
	var tests = map[string]struct {
		given    int
		expected expected
	}{
		"zero": {
			given:    0,
			expected: expected{"000", false},
		},
		"one number": {
			given:    8,
			expected: expected{"008", false},
		},
		"two numbers": {
			given:    21,
			expected: expected{"021", false},
		},
		"max numbers": {
			given:    211,
			expected: expected{"211", false},
		},
		"exceed number": {
			given:    2110,
			expected: expected{"", true},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := r.transform(tt.given)
			if actual != tt.expected.transformed {
				t.Errorf("(%d): expected %v, actual %v", tt.given, tt.expected.transformed, actual)
			}
			if tt.expected.hasErr && err == nil {
				t.Errorf("(%d): expected an error", tt.given)
			}
		})
	}
}

func TestNnnRenamer_Rename(t *testing.T) {
	type given struct {
		nb       int
		fileName string
	}
	r := &NnnRenamer{NbNumbers: 3}
	var tests = map[string]struct {
		given    given
		expected string
	}{
		"zero": {
			given:    given{0, "foobar.txt"},
			expected: "foobar_000.txt",
		},
		"one number": {
			given:    given{8, "foobar.txt"},
			expected: "foobar_008.txt",
		},
		"two numbers": {
			given:    given{21, "foobar.txt"},
			expected: "foobar_021.txt",
		},
		"max numbers": {
			given:    given{211, "foobar.txt"},
			expected: "foobar_211.txt",
		},
		"without ext": {
			given:    given{10, "foobar"},
			expected: "foobar_010",
		},
		"without base": {
			given:    given{10, ".txt"},
			expected: "010.txt",
		},
		"empty string": {
			given:    given{10, ""},
			expected: "010",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := r.Rename(tt.given.nb, tt.given.fileName)
			if actual != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}
		})
	}
}
