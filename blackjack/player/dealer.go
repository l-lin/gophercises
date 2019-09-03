package player

import "github.com/l-lin/gophercises/deck"

// Dealer of the blackjac
type Dealer struct {
	Player
}

// NewDealer returns a new instanciated dealer
func NewDealer(cards ...deck.Card) *Dealer {
	d := &Dealer{}
	d.HandCard = NewHandCard(cards...)
	return d
}
