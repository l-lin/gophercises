package renamer

import (
	"testing"
)

func TestOfRenamer_Rename(t *testing.T) {
	type given struct {
		nb       int
		fileName string
	}
	renamer := &OfRenamer{Max: 100}
	var tests = map[string]struct {
		given    given
		expected string
	}{
		"just a file name": {
			given:    given{2, "foobar.txt"},
			expected: "foobar (2 of 100).txt",
		},
		"without extension": {
			given:    given{2, "foobar"},
			expected: "foobar (2 of 100)",
		},
		"empty fileName": {
			given:    given{2, ""},
			expected: "(2 of 100)",
		},
		"space": {
			given:    given{2, " "},
			expected: "  (2 of 100)",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := renamer.Rename(tt.given.nb, tt.given.fileName)
			if actual != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}
		})
	}
}
