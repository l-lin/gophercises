package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Suit: Spade, Rank: King})
	fmt.Println(Card{Suit: Diamond, Rank: Queen})
	fmt.Println(Card{Suit: Hearth, Rank: Eight})
	fmt.Println(Card{Suit: RedJoker})
	// Output:
	// King of Spade
	// Queen of Diamond
	// Eight of Hearth
	// RedJoker
}

func ExampleCard_ToASCII() {
	fmt.Println(Card{Suit: Spade, Rank: Ten}.ToASCII())
	fmt.Println(Card{Suit: Diamond, Rank: Eight}.ToASCII())
	fmt.Println(Card{Suit: Club, Rank: Jack}.ToASCII())
	fmt.Println(Card{Suit: Hearth, Rank: Ace}.ToASCII())
	fmt.Println(Card{Suit: BlackJoker}.ToASCII())
	fmt.Println(Card{Suit: RedJoker}.ToASCII())
	// Output:
	// ┌────────┐
	// │10 .    │
	// │  / \   │
	// │ (_,_)  │
	// │   I  10│
	// └────────┘
	// ┌────────┐
	// │8  /\   │
	// │  /  \  │
	// │  \  /  │
	// │   \/  8│
	// └────────┘
	// ┌────────┐
	// │J  _    │
	// │  ( )   │
	// │ (_x_)  │
	// │   Y   J│
	// └────────┘
	// ┌────────┐
	// │A _  _  │
	// │ ( \/ ) │
	// │  \  /  │
	// │   \/  A│
	// └────────┘
	// ┌────────┐
	// │* \||/ K│
	// │J /~~\ O│
	// │O( o o)J│
	// │K \ v/ *│
	// └────────┘
	// ┌────────┐
	// │+ \||/ K│
	// │J /~~\ O│
	// │O( o o)J│
	// │K \ v/ +│
	// └────────┘
}

func TestNewDeck(t *testing.T) {
	cards := NewDeck()
	expectedCardsNb := 52
	if len(cards) != expectedCardsNb {
		t.Errorf("not enough cards, expected %d, got %d", expectedCardsNb, len(cards))
	}
}

func TestFromDecks(t *testing.T) {
	deck1 := NewDeck()
	deck2 := NewDeck()

	cards := FromDecks(deck1, deck2)

	expectedCardsNb := len(deck1) + len(deck2)

	if len(cards) != expectedCardsNb {
		t.Errorf("not enough cards, expected %d cards, got %d cards", expectedCardsNb, len(cards))
	}
}

func TestComputeCoeff(t *testing.T) {
	var tests = map[string]struct {
		given    int
		expected int
	}{
		"13":    {13, 100},
		"99":    {99, 100},
		"1":     {1, 10},
		"81234": {81234, 100000},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := computeCoeff(tt.given)
			if actual != tt.expected {
				t.Errorf("(%d): expected %d, actual %d", tt.given, tt.expected, actual)
			}
		})
	}
}
