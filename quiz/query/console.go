package query

import (
	"github.com/l-lin/gophercises/quiz/problem"
	"github.com/manifoldco/promptui"
)

// ConsoleQuerier displays the questions in the console
type ConsoleQuerier struct{}

// AskReady ask the user if he/she is ready to take the quiz
func (cq *ConsoleQuerier) AskReady() (bool, error) {
	prompt := promptui.Select{
		Label: "Are you ready?",
		Items: []string{"yes", "no"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		return false, err
	}
	return "yes" == result, nil
}

// Query the problem to the user
func (cq *ConsoleQuerier) Query(pb *problem.Problem) (bool, error) {
	prompt := promptui.Prompt{
		Label: pb.Question + " ",
	}
	result, err := prompt.Run()
	if err != nil {
		return false, err
	}
	if pb.IsCorrect(result) {
		return true, nil
	}
	return false, nil
}
