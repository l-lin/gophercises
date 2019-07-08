package problem

import "testing"

func TestIsCorrect(t *testing.T) {
	pb := &Problem{Answer: "foobar"}
	var tests = []struct {
		name     string
		expected bool
		given    string
	}{
		{"Correct answer", true, "foobar"},
		{"Incorrect answer", false, "barfoo"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := pb.IsCorrect(tt.given)
			if actual != tt.expected {
				t.Errorf("pb.IsCorrect(%s): expected %v, actual %v", tt.given, tt.expected, actual)
			}

		})
	}
}
