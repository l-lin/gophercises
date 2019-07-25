package link

import "testing"

func TestNew(t *testing.T) {
	var tests = []struct {
		name          string
		expectedError bool
		given         string
	}{
		{"Absolute URL", false, "http://httpbin.org"},
		{"Relative URL starting with /", false, "/foobar"},
		{"Relative URL not starting with /", false, "foo/bar.html"},
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

func TestIsSameLink(t *testing.T) {
	l, _ := New("http://httpbin.org")
	sameLink, _ := New("http://httpbin.org")
	differentLink, _ := New("http://httpbin.org/foobar")
	differentLinkRelative, _ := New("/get")
	var tests = []struct {
		name     string
		expected bool
		given    *Link
	}{
		{"Same link", true, sameLink},
		{"Different link absolute", false, differentLink},
		{"Different link relative", false, differentLinkRelative},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := l.IsSameLink(*tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %v, actual %v", tt.given, tt.expected, actual)
			}

		})
	}

}
