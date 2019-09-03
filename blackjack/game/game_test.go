package game

import (
	"errors"
	"testing"

	"github.com/l-lin/gophercises/deck"
)

func TestInitDealer(t *testing.T) {
	cards := deck.NewDeck(deck.Shuffle)
	nbInitCards := len(cards)
	result, dealer := initDealer(cards)

	if len(dealer.HandCard.Cards) != nbCardsOnStart {
		t.Errorf("the dealer cards should have been %d, got instead %d", nbCardsOnStart, len(dealer.HandCard.Cards))
	}
	expectedCardsLeft := nbInitCards - nbCardsOnStart
	if len(result) != expectedCardsLeft {
		t.Errorf("the number of cards left should have been %d, got instead %d", expectedCardsLeft, len(result))
	}
}

func TestInitPlayers(t *testing.T) {
	nbPlayers := 2
	cards := deck.NewDeck(deck.Shuffle)
	nbInitCards := len(cards)
	result, players := initPlayers(cards, nbPlayers)

	if len(players) != nbPlayers {
		t.Errorf("expected %d players, got %d", nbPlayers, len(players))
	}
	for i := 0; i < nbPlayers; i++ {
		if len(players[i].HandCard.Cards) != nbCardsOnStart {
			t.Errorf("expected %d cards, got %d", nbCardsOnStart, len(players[i].HandCard.Cards))
		}
	}
	expectedCardsLeft := nbInitCards - nbPlayers*nbCardsOnStart
	if len(result) != expectedCardsLeft {
		t.Errorf("the number of cards left should have been %d, got instead %d", expectedCardsLeft, len(result))
	}
}

func TestHit(t *testing.T) {
	type expected struct {
		err       error
		cardsLeft int
		card      *deck.Card
	}
	var tests = map[string]struct {
		given    []deck.Card
		expected expected
	}{
		"basic with 3 cards": {
			given: []deck.Card{
				deck.Card{Suit: deck.Spade, Rank: deck.Ten},
				deck.Card{Suit: deck.Hearth, Rank: deck.Five},
				deck.Card{Suit: deck.Club, Rank: deck.Queen},
			},
			expected: expected{
				err:       nil,
				cardsLeft: 2,
				card:      &deck.Card{Suit: deck.Spade, Rank: deck.Ten},
			},
		},
		"no cards left": {
			given: []deck.Card{},
			expected: expected{
				err:       errors.New("some error was expected"),
				cardsLeft: 0,
				card:      nil,
			},
		},
		"1 card left": {
			given: []deck.Card{
				deck.Card{Suit: deck.Hearth, Rank: deck.Five},
			},
			expected: expected{
				err:       nil,
				cardsLeft: 0,
				card:      &deck.Card{Suit: deck.Hearth, Rank: deck.Five},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			g := Game{Cards: tt.given}
			c, err := g.hit()
			if tt.expected.err != nil && err == nil {
				t.Errorf("expected error %v, actual error %v", tt.expected.err, err)
			}
			if tt.expected.card != nil && !c.Equals(*tt.expected.card) {
				t.Errorf("expected card %v, actual card %v", tt.expected.card, c)
			}
			if len(g.Cards) != tt.expected.cardsLeft {
				t.Errorf("expected %d cards left, actual %d cards left", tt.expected.cardsLeft, len(g.Cards))
			}
		})
	}
}
