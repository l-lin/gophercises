package deck

import (
	"fmt"
	"testing"
)

func ExampleFilterOut() {
	cards := NewDeck(FilterOut(func(card Card) bool {
		return card.Rank != Ace
	}))
	for _, card := range cards {
		fmt.Println(card)
	}
	// Output:
	// Ace of Spade
	// Ace of Diamond
	// Ace of Club
	// Ace of Hearth
}

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
			if card.Equals(cardToFilterOut) {
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
			if c.Equals(cardToFilterOut) {
				t.Errorf("%s was not filtered out!", c)
			}
		}
	}
}
