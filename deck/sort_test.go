package deck

import (
	"fmt"
	"testing"
)

func ExampleDefaultSort() {
	for _, card := range DefaultSort([]Card{
		Card{Suit: Hearth, Rank: Two},
		Card{Suit: Spade, Rank: Queen},
		Card{Suit: Spade, Rank: Four},
		Card{Suit: Diamond, Rank: Ace},
		Card{Suit: Club, Rank: Ten},
	}) {
		fmt.Println(card)
	}
	// Output:
	// Four of Spade
	// Queen of Spade
	// Ace of Diamond
	// Ten of Club
	// Two of Hearth
}

func TestDefaultSort(t *testing.T) {
	cards := []Card{
		Card{Suit: Hearth, Rank: Two},
		Card{Suit: Spade, Rank: Queen},
		Card{Suit: Spade, Rank: Four},
		Card{Suit: Diamond, Rank: Ace},
		Card{Suit: Club, Rank: Ten},
	}
	// cannot sort copied slices for some unknown reason...
	// go does not perform the Less function to all elements
	given := make([]Card, len(cards))
	copy(given, cards)

	cards = DefaultSort(cards)

	expectedRanges := [...]int{2, 1, 3, 4, 0}
	for gotRange, expectedRange := range expectedRanges {
		if !cards[gotRange].Equals(given[expectedRange]) {
			t.Errorf("card number %d must be %s, got %s", gotRange, given[expectedRange].String(), cards[gotRange].String())
		}
	}
}
