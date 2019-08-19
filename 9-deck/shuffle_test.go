package main

import "testing"

func TestShuffle(t *testing.T) {
	given := []Card{
		Card{Suit: Hearth, Rank: Two},
		Card{Suit: Spade, Rank: Queen},
		Card{Suit: Spade, Rank: Four},
		Card{Suit: Diamond, Rank: Ace},
		Card{Suit: Club, Rank: Ten},
	}
	cards := make([]Card, len(given))
	copy(cards, given)

	cards = Shuffle(cards)

	if len(cards) != len(given) {
		t.Errorf("the result must preserve the slice length, expected %d, got %d", len(given), len(cards))
	}

	isSame := true
	for i := 0; i < len(cards); i++ {
		if cards[i] != given[i] {
			isSame = false
			break
		}
	}
	if isSame {
		t.Error("shuffle has failed miserably...")
	}
}
