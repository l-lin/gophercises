package query

import (
	"github.com/l-lin/quiz/problem"
	"github.com/manifoldco/promptui"
)

// ConsoleQuerier displays the questions in the console
type ConsoleQuerier struct{}

func (cq *ConsoleQuerier) query(pb *problem.Problem) (bool, error) {
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
