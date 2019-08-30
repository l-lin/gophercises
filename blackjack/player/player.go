package player

import "github.com/l-lin/gophercises/deck"

// Player of the blackjack
type Player struct {
	HandCard HandCard
}

// NewPlayer returns a new instanciated player
func NewPlayer(cards ...deck.Card) Player {
	return Player{
		HandCard: NewHandCard(cards...),
	}
}

// Hit a new card
func (p Player) Hit(c deck.Card) {
	p.HandCard.Add(c)
}
