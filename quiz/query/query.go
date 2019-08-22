package query

import (
	"fmt"
	"log"
	"time"

	"github.com/l-lin/gophercises/quiz/problem"
)

// Querier asks the user the questions
type Querier interface {
	AskReady() (bool, error)
	Query(pbs *problem.Problem) (bool, error)
}

// Query the questions
func Query(querier Querier, pbs []*problem.Problem, timer time.Duration) (*Result, error) {
	finished := make(chan bool, 1)
	errCh := make(chan error, 1)
	r := &Result{NbTotalAnswers: len(pbs)}

	go performQuery(finished, errCh, r, pbs, querier)

	select {
	case <-finished:
	case err := <-errCh:
		log.Fatalln(err)
	case <-time.After(timer):
		fmt.Println("Timed out")
	}
	return r, nil
}

func performQuery(finished chan<- bool, errCh chan<- error, result *Result, pbs []*problem.Problem, querier Querier) {
	for _, pb := range pbs {
		r, err := querier.Query(pb)
		if err != nil {
			errCh <- err
		}
		if r {
			result.NbCorrectAnswers++
		}
	}
	finished <- true
}

// Result of the quiz
type Result struct {
	NbCorrectAnswers, NbTotalAnswers int
	Err                              error
}

func (r *Result) String() string {
	return fmt.Sprintf("Result: %d / %d", r.NbCorrectAnswers, r.NbTotalAnswers)
}
