package game

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

// Score of the players and dealer
type Score struct {
	Players map[int]int
	Dealer  int
}

func newScore(nbPlayers int) *Score {
	scores := make(map[int]int, nbPlayers)
	for i := 1; i <= nbPlayers; i++ {
		scores[i] = 0
	}
	return &Score{Players: scores, Dealer: 0}
}

func (s *Score) diplay() {
	fmt.Println("")
	for i, v := range s.Players {
		fmt.Println(aurora.BrightBlack(fmt.Sprintf("Player %d: %d wins", i, v)).BgBrightBlue().Bold())
	}
	fmt.Println(aurora.BrightBlack(fmt.Sprintf("Dealer: %d wins", s.Dealer)).BgBrightYellow().Bold())
	fmt.Println("")
}
