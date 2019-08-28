package player

import "github.com/l-lin/gophercises/deck"

// Player of the blackjack
type Player struct {
	HandCard HandCard
}

// NewPlayer returns a new instanciated player
func NewPlayer(c1, c2 deck.Card) Player {
	return Player{
		HandCard: NewHandCard(c1, c2),
	}
}

// Hit a new card
func (p Player) Hit(c deck.Card) {
	p.HandCard.Add(c)
}
