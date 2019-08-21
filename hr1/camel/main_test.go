package main

import "testing"

func TestCamelCase(t *testing.T) {
	var tests = []struct {
		name     string
		expected int32
		given    string
	}{
		{"nominal case", 2, "fooBar"},
		{"start with a upper case", 2, "FooBar"},
		{"empty string", 0, ""},
		{"long string", 10, "somethingReallyLongIHopeItCanComputeItRight"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := camelcase(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %v, actual %v", tt.given, tt.expected, actual)
			}

		})
	}

}
