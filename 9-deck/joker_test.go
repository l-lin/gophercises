package main

import "testing"

func TestAddJoker(t *testing.T) {
	given := []Card{
		Card{Suit: Hearth, Rank: Two},
		Card{Suit: Spade, Rank: Queen},
		Card{Suit: Spade, Rank: Four},
		Card{Suit: Diamond, Rank: Ace},
		Card{Suit: Club, Rank: Ten},
	}
	cards := make([]Card, len(given))
	copy(cards, given)

	cards = AddJokers(cards)

	if len(cards) == len(given) {
		t.Errorf("the result must have additional jokers, expected %d cards, got %d cards", len(given), len(cards))
	}
}
