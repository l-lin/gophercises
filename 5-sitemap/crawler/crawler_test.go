package crawler

import "testing"

func TestNew(t *testing.T) {
	var tests = []struct {
		name          string
		expectedError bool
		given         string
	}{
		{"Absolute URL", false, "http://httpbin.org"},
		{"Relative URL", true, "/foobar"},
		{"Relative URL", true, "foo/bar.html"},
		{"rubbish URL", true, "htt/foo.bar"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, actual := New(tt.given)
			hasError := actual != nil
			if hasError != tt.expectedError {
				t.Errorf("(%s): expected error %v, actual %v", tt.given, tt.expectedError, hasError)
			}

		})
	}
}
