package game

import (
	"testing"

	"github.com/l-lin/gophercises/deck"
)

func TestInitDealer(t *testing.T) {
	cards := deck.NewDeck(deck.Shuffle)
	nbInitCards := len(cards)
	result := initDealer(cards)

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
	result := initPlayers(cards, nbPlayers)

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
