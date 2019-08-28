package player

import "github.com/l-lin/gophercises/deck"

// Action of a player or dealer
type Action interface {
	Hit(c deck.Card) bool
	Stand()
}
