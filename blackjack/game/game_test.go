package game

import (
	"errors"
	"testing"

	"github.com/l-lin/gophercises/blackjack/player"
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

func TestFinish(t *testing.T) {
	type expected struct {
		nbWPlayer int
		wPlayer   *player.Player
		wDealer   *player.Dealer
	}
	var tests = map[string]struct {
		given Game
		expected
	}{
		"no one is over, player 1 wins": {
			given: Game{
				Players: []*player.Player{
					player.NewPlayer(newCard(1), newCard(9)),
					player.NewPlayer(newCard(9), newCard(9)),
				},
				Dealer: player.NewDealer(newCard(1), newCard(8)),
			},
			expected: expected{
				nbWPlayer: 1,
				wPlayer:   player.NewPlayer(newCard(1), newCard(9)),
				wDealer:   nil,
			},
		},
		"no one is over, player 2 wins": {
			given: Game{
				Players: []*player.Player{
					player.NewPlayer(newCard(9), newCard(9)),
					player.NewPlayer(newCard(1), newCard(9)),
				},
				Dealer: player.NewDealer(newCard(1), newCard(8)),
			},
			expected: expected{
				nbWPlayer: 2,
				wPlayer:   player.NewPlayer(newCard(1), newCard(9)),
				wDealer:   nil,
			},
		},
		"no one is over, dealer wins": {
			given: Game{
				Players: []*player.Player{
					player.NewPlayer(newCard(1), newCard(8)),
					player.NewPlayer(newCard(9), newCard(9)),
				},
				Dealer: player.NewDealer(newCard(1), newCard(9)),
			},
			expected: expected{
				nbWPlayer: 0,
				wPlayer:   nil,
				wDealer:   player.NewDealer(newCard(1), newCard(9)),
			},
		},
		"player 2 is over, player 1 wins": {
			given: Game{
				Players: []*player.Player{
					player.NewPlayer(newCard(1), newCard(9)),
					player.NewPlayer(newCard(9), newCard(9), newCard(10)),
				},
				Dealer: player.NewDealer(newCard(1), newCard(7)),
			},
			expected: expected{
				nbWPlayer: 1,
				wPlayer:   player.NewPlayer(newCard(1), newCard(9)),
				wDealer:   nil,
			},
		},
		"player 1 is over, player 2 wins": {
			given: Game{
				Players: []*player.Player{
					player.NewPlayer(newCard(9), newCard(9), newCard(10)),
					player.NewPlayer(newCard(1), newCard(9)),
				},
				Dealer: player.NewDealer(newCard(1), newCard(7)),
			},
			expected: expected{
				nbWPlayer: 2,
				wPlayer:   player.NewPlayer(newCard(1), newCard(9)),
				wDealer:   nil,
			},
		},
		"both players are over, dealer wins": {
			given: Game{
				Players: []*player.Player{
					player.NewPlayer(newCard(8), newCard(9), newCard(5)),
					player.NewPlayer(newCard(9), newCard(9), newCard(10)),
				},
				Dealer: player.NewDealer(newCard(1), newCard(7)),
			},
			expected: expected{
				nbWPlayer: 0,
				wPlayer:   nil,
				wDealer:   player.NewDealer(newCard(1), newCard(7)),
			},
		},
		"draw with player 1 and dealer": {
			given: Game{
				Players: []*player.Player{
					player.NewPlayer(newCard(1), newCard(9)),
					player.NewPlayer(newCard(9), newCard(9), newCard(10)),
				},
				Dealer: player.NewDealer(newCard(10), newCard(10)),
			},
			expected: expected{
				nbWPlayer: 1,
				wPlayer:   player.NewPlayer(newCard(1), newCard(9)),
				wDealer:   player.NewDealer(newCard(10), newCard(10)),
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actualNbWPlayer, actualWPlayer, actualWDealer := tt.given.getWinner()
			if tt.expected.wDealer != nil && (actualWDealer == nil || !tt.expected.wDealer.Player.Equals(&actualWDealer.Player)) {
				t.Errorf("expected dealer as winner %v, actual %v", tt.expected.wDealer, actualWDealer)
			}
			if tt.expected.wPlayer != nil && (actualWPlayer == nil || !tt.expected.wPlayer.Equals(actualWPlayer) || tt.expected.nbWPlayer != actualNbWPlayer) {
				t.Errorf("expected player %d as winner %v, actual player %d as winner %v", tt.expected.nbWPlayer, tt.expected.wPlayer, actualNbWPlayer, actualWPlayer)
			}
		})
	}
}

func newCard(rank int) deck.Card {
	return deck.Card{Suit: deck.Spade, Rank: deck.Rank(rank)}
}
