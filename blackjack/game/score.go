package game

import (
	"fmt"
	"strings"
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

func (s *Score) print() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("[Dealer: %d wins](fg:yellow,mod:bold)\n", s.Dealer))
	for i := 0; i < len(s.Players); i++ {
		b.WriteString(fmt.Sprintf("[Player %d: %d wins](fg:blue,mod:bold)", i+1, s.Players[i+1]))
		if i < len(s.Players) {
			b.WriteString("\n")
		}
	}
	return b.String()
}
