package main

import "testing"

func TestCaesarCipher(t *testing.T) {
	type args struct {
		s string
		k int
	}

	var tests = map[string]struct {
		given    args
		expected string
	}{
		"simple example": {
			given:    args{s: "foobar", k: 1},
			expected: "gppcbs",
		},
		"example provided by hacker rank": {
			given:    args{s: "middle-Outz", k: 2},
			expected: "okffng-Qwvb",
		},
		"rotation": {
			given:    args{s: "RsTZ Pop", k: 10},
			expected: "BcDJ Zyz",
		},
		"empty string": {
			given:    args{s: "", k: 10},
			expected: "",
		},
		"all special characters": {
			given:    args{s: "%*+ ----", k: 8},
			expected: "%*+ ----",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := caesarCipher(tt.given.s, tt.given.k)
			if actual != tt.expected {
				t.Errorf("(%s, %d): expected %s, actual %s", tt.given.s, tt.given.k, tt.expected, actual)
			}
		})
	}
}
