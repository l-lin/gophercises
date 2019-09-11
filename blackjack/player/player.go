package player

import (
	"fmt"

	"github.com/l-lin/gophercises/deck"
)

// Player of the blackjack
type Player struct {
	HandCard HandCard
	Finished bool
}

// NewPlayer returns a new instanciated player
func NewPlayer(cards ...deck.Card) *Player {
	return &Player{
		HandCard: NewHandCard(cards...),
	}
}

// Hit a new card
func (p Player) Hit(c deck.Card) {
	p.HandCard.Add(c)
}

// CompareTo another player the score
func (p *Player) CompareTo(to *Player) int {
	return p.HandCard.CompareTo(to.HandCard)
}

// Equals checks if the player is equals to the given player
func (p *Player) Equals(to *Player) bool {
	return p.HandCard.Equals(to.HandCard) && p.Finished == to.Finished
}

func (p Player) String() string {
	return fmt.Sprintf("%v", p.HandCard.Cards)
}
