package game

import (
	"fmt"

	"github.com/l-lin/gophercises/blackjack/player"
	"github.com/l-lin/gophercises/deck"
)

var (
	dealer  player.Dealer
	players []player.Player
	cards   []deck.Card
)

const nbCardsOnStart = 2

// Run blackjack game
func Run(nbPlayers int) {
	initGame(nbPlayers)
	fmt.Println(dealer.HandCard.Print())
}

func initGame(nbPlayers int) {
	cards = deck.NewDeck(deck.Shuffle)
	cards = initDealer(cards)
	cards = initPlayers(cards, nbPlayers)
	fmt.Println(len(cards))
}

func initDealer(cards []deck.Card) []deck.Card {
	dealerCards := make([]deck.Card, nbCardsOnStart)
	for i := 0; i < nbCardsOnStart; i++ {
		dealerCards[i] = cards[i]
	}
	dealer = player.NewDealer(dealerCards...)
	return cards[nbCardsOnStart:]
}

func initPlayers(cards []deck.Card, nbPlayers int) []deck.Card {
	players = make([]player.Player, nbPlayers)
	for i := 0; i < nbPlayers; i++ {
		playerCards := make([]deck.Card, nbCardsOnStart)
		for j := 0; j < nbCardsOnStart; j++ {
			playerCards[j] = cards[i+j]
		}
		players[i] = player.NewPlayer(playerCards...)
		cards = cards[nbCardsOnStart:]
	}
	return cards
}
