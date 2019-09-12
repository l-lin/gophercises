package game

import (
	"github.com/l-lin/gophercises/blackjack/player"
	"github.com/l-lin/gophercises/deck"
)

func initDealer(cards []deck.Card) ([]deck.Card, *player.Dealer) {
	dealerCards := make([]deck.Card, nbCardsOnStart)
	for i := 0; i < nbCardsOnStart; i++ {
		dealerCards[i] = cards[i]
	}
	dealerCards[0].Hidden = true
	dealer := player.NewDealer(dealerCards...)
	return cards[nbCardsOnStart:], dealer
}

func initPlayers(cards []deck.Card, nbPlayers int) ([]deck.Card, []*player.Player) {
	players := make([]*player.Player, nbPlayers)
	for i := 0; i < nbPlayers; i++ {
		playerCards := make([]deck.Card, nbCardsOnStart)
		for j := 0; j < nbCardsOnStart; j++ {
			playerCards[j] = cards[i+j]
		}
		players[i] = player.NewPlayer(playerCards...)
		cards = cards[nbCardsOnStart:]
	}
	return cards, players
}
