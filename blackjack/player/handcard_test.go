package player

import (
	"testing"

	"github.com/l-lin/gophercises/deck"
)

func TestCompute(t *testing.T) {
	var tests = map[string]struct {
		given    HandCard
		expected int
	}{
		"one card of 2": {
			given: NewHandCard(
				deck.Card{Rank: deck.Two},
			),
			expected: 2,
		},
		"two cards": {
			given: NewHandCard(
				deck.Card{Rank: deck.Two},
				deck.Card{Rank: deck.Five},
			),
			expected: 7,
		},
		"three cards exceed 21": {
			given: NewHandCard(
				deck.Card{Rank: deck.Two},
				deck.Card{Rank: deck.Five},
				deck.Card{Rank: deck.Jack},
			),
			expected: 17,
		},
		"three cards with ace": {
			given: NewHandCard(
				deck.Card{Rank: deck.Ace},
				deck.Card{Rank: deck.Five},
				deck.Card{Rank: deck.Jack},
			),
			expected: 16,
		},
		"multiple aces exceed 21": {
			given: NewHandCard(
				deck.Card{Rank: deck.Ace},
				deck.Card{Rank: deck.Queen},
				deck.Card{Rank: deck.Ace},
				deck.Card{Rank: deck.King},
				deck.Card{Rank: deck.Ace},
			),
			expected: 23,
		},
		"multiple aces with different values": {
			given: NewHandCard(
				deck.Card{Rank: deck.Ace},
				deck.Card{Rank: deck.Queen},
				deck.Card{Rank: deck.Ace},
				deck.Card{Rank: deck.Ace},
			),
			expected: 23,
		},
		"blackjack": {
			given: NewHandCard(
				deck.Card{Rank: deck.Ace},
				deck.Card{Rank: deck.King},
			),
			expected: 21,
		},
		"no cards": {
			given:    NewHandCard(),
			expected: 0,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := tt.given.Compute()
			if actual != tt.expected {
				t.Errorf("expected %d, actual %d", tt.expected, actual)
			}
		})
	}
}

func TestIsBlackJack(t *testing.T) {
	var tests = map[string]struct {
		given    HandCard
		expected bool
	}{
		"blackjack with jack": {
			given: NewHandCard(
				deck.Card{Rank: deck.Ace},
				deck.Card{Rank: deck.Jack},
			),
			expected: true,
		},
		"blackjack with queen": {
			given: NewHandCard(
				deck.Card{Rank: deck.Ace},
				deck.Card{Rank: deck.Queen},
			),
			expected: true,
		},
		"blackjack with king": {
			given: NewHandCard(
				deck.Card{Rank: deck.Ace},
				deck.Card{Rank: deck.King},
			),
			expected: true,
		},
		"score 21 but with 3 cards": {
			given: NewHandCard(
				deck.Card{Rank: deck.Ace},
				deck.Card{Rank: deck.Ace},
				deck.Card{Rank: deck.Nine},
			),
			expected: false,
		},
		"score 21 but with 3 different cards": {
			given: NewHandCard(
				deck.Card{Rank: deck.Five},
				deck.Card{Rank: deck.Seven},
				deck.Card{Rank: deck.Nine},
			),
			expected: false,
		},
		"not 21 score": {
			given: NewHandCard(
				deck.Card{Rank: deck.Five},
				deck.Card{Rank: deck.Seven},
			),
			expected: false,
		},
		"no cards": {
			given:    NewHandCard(),
			expected: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := tt.given.IsBlackJack()
			if actual != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}

		})
	}
}

func TestCompareTo(t *testing.T) {
	var tests = map[string]struct {
		given    []HandCard
		expected int
	}{
		"less": {
			given: []HandCard{
				NewHandCard(
					deck.Card{Rank: deck.Five},
					deck.Card{Rank: deck.Seven},
				),
				NewHandCard(
					deck.Card{Rank: deck.Jack},
					deck.Card{Rank: deck.Seven},
				),
			},
			expected: -1,
		},
		"greater": {
			given: []HandCard{
				NewHandCard(
					deck.Card{Rank: deck.Jack},
					deck.Card{Rank: deck.Queen},
				),
				NewHandCard(
					deck.Card{Rank: deck.Jack},
					deck.Card{Rank: deck.Seven},
				),
			},
			expected: 1,
		},
		"equals": {
			given: []HandCard{
				NewHandCard(
					deck.Card{Rank: deck.Ace},
					deck.Card{Rank: deck.Nine},
				),
				NewHandCard(
					deck.Card{Rank: deck.Ten},
					deck.Card{Rank: deck.Ten},
				),
			},
			expected: 0,
		},
		"equals but more cards": {
			given: []HandCard{
				NewHandCard(
					deck.Card{Rank: deck.Jack},
					deck.Card{Rank: deck.Nine},
					deck.Card{Rank: deck.Ace},
				),
				NewHandCard(
					deck.Card{Rank: deck.Jack},
					deck.Card{Rank: deck.King},
				),
			},
			expected: -1,
		},
		"equals but less cards": {
			given: []HandCard{
				NewHandCard(
					deck.Card{Rank: deck.Jack},
					deck.Card{Rank: deck.Nine},
					deck.Card{Rank: deck.Ace},
				),
				NewHandCard(
					deck.Card{Rank: deck.Jack},
					deck.Card{Rank: deck.Two},
					deck.Card{Rank: deck.Five},
					deck.Card{Rank: deck.Three},
				),
			},
			expected: 1,
		},
		"from is busted": {
			given: []HandCard{
				NewHandCard(
					deck.Card{Rank: deck.Jack},
					deck.Card{Rank: deck.Nine},
					deck.Card{Rank: deck.Three},
				),
				NewHandCard(
					deck.Card{Rank: deck.Jack},
					deck.Card{Rank: deck.Two},
					deck.Card{Rank: deck.Five},
					deck.Card{Rank: deck.Three},
				),
			},
			expected: -1,
		},
		"to is busted": {
			given: []HandCard{
				NewHandCard(
					deck.Card{Rank: deck.Jack},
					deck.Card{Rank: deck.Nine},
					deck.Card{Rank: deck.Ace},
				),
				NewHandCard(
					deck.Card{Rank: deck.Jack},
					deck.Card{Rank: deck.Two},
					deck.Card{Rank: deck.Five},
					deck.Card{Rank: deck.Five},
				),
			},
			expected: 1,
		},
		"both are busted": {
			given: []HandCard{
				NewHandCard(
					deck.Card{Rank: deck.Jack},
					deck.Card{Rank: deck.Nine},
					deck.Card{Rank: deck.Ace},
				),
				NewHandCard(
					deck.Card{Rank: deck.Jack},
					deck.Card{Rank: deck.Two},
					deck.Card{Rank: deck.Five},
					deck.Card{Rank: deck.Five},
				),
			},
			expected: 1,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := tt.given[0].CompareTo(tt.given[1])
			if actual != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}

		})
	}
}
