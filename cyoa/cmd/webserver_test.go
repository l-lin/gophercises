package cmd

import "testing"

func TestSanitizePath(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		given    string
	}{
		{"Nominal case", "foobar", "/foobar"},
		{"Nothing to change", "foobar", "foobar"},
		{"Edge case", "foobar", "/foobar/"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := sanitizePath(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %s, actual %s", tt.given, tt.expected, actual)
			}
		})
	}
}
