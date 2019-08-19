package main

import (
	"testing"
)

func TestFilterOut(t *testing.T) {
	given := []Card{
		Card{Suit: Hearth, Rank: Two},
		Card{Suit: Spade, Rank: Queen},
		Card{Suit: Spade, Rank: Four},
		Card{Suit: Diamond, Rank: Ace},
		Card{Suit: Club, Rank: Ten},
	}
	cards := make([]Card, len(given))
	copy(cards, given)

	cardsToFilterOut := []Card{
		Card{Suit: Diamond, Rank: Ace},
		Card{Suit: Spade, Rank: Queen},
		Card{Suit: Club, Rank: Two},
	}

	cards = FilterOut(func(card Card) bool {
		for _, cardToFilterOut := range cardsToFilterOut {
			if card.Rank == cardToFilterOut.Rank && card.Suit == cardToFilterOut.Suit {
				return true
			}
		}
		return false
	})(cards)

	if len(cards) != len(given)-2 {
		t.Errorf("2 cards was filtered out from the deck, got %d cards, expected %d cards", len(cards), len(given)-2)
	}
	for _, c := range cards {
		for _, cardToFilterOut := range cardsToFilterOut {
			if c.Rank == cardToFilterOut.Rank && c.Suit == cardToFilterOut.Suit {
				t.Errorf("%s was not filtered out!", c)
			}
		}
	}
}
