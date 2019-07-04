package problem

import (
	"fmt"
	"io"
	"strings"
)

// Problem is the data representation of a problem from the input
type Problem struct {
	Question string
	Answer   string
}

// IsCorrect checks if the given answer is correct for this problem
func (p *Problem) IsCorrect(answer string) bool {
	return strings.TrimSpace(answer) == strings.TrimSpace(p.Answer)
}

func (p *Problem) String() string {
	return fmt.Sprintf("question: %s, answer: %s", p.Question, p.Answer)
}

// ProblemsParser extracts inputs to a slice of Problem
type ProblemsParser interface {
	Parse(in io.Reader) ([]*Problem, error)
}
