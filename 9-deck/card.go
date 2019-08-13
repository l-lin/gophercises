package main

import "fmt"

// Card from a deck
type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}

func (c Card) absRank() int {
	// 100 is the next decimal of maxRank
	// TODO: compute next decimal
	return int(c.Suit)*100 + int(c.Rank)
}

// NewDeck with sorted cards
func NewDeck(opt func([]Card)) []Card {
	cards := []Card{}
	for _, s := range suits {
		for i := minRank; i <= maxRank; i++ {
			cards = append(cards, Card{Suit: s, Rank: i})
		}
	}

	opt(cards)
	return cards
}
