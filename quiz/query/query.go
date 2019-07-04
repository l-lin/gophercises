package query

import (
	"fmt"

	"github.com/l-lin/quiz/problem"
)

// Querier asks the user the questions
type Querier interface {
	Query(pbs *problem.Problem) (bool, error)
}

// Query the questions
func Query(pbs []*problem.Problem) (*Result, error) {
	nbCorrectAnswers := 0
	querier := ConsoleQuerier{}
	for _, pb := range pbs {
		result, err := querier.query(pb)
		if err != nil {
			return nil, err
		}
		if result {
			nbCorrectAnswers++
		}
	}
	return &Result{NbCorrectAnswers: nbCorrectAnswers, NbTotalAnswers: len(pbs)}, nil
}

// Result of the quizz
type Result struct {
	NbCorrectAnswers, NbTotalAnswers int
}

func (r *Result) String() string {
	return fmt.Sprintf("Result: %d / %d", r.NbCorrectAnswers, r.NbTotalAnswers)
}
