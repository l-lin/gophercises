//go:generate stringer -type=Suit
package main

// Suit represents the type of card
type Suit int

const (
	// Spade card type
	Spade Suit = iota
	// Diamond card type
	Diamond
	// Club card type
	Club
	// Hearth card type
	Hearth
)

var suits = [...]Suit{Spade, Diamond, Club, Hearth}
