package game

import (
	"github.com/l-lin/gophercises/blackjack/player"
)

var (
	dealer  player.Dealer
	players []player.Player
)

// Run blackjack game
func Run(nbPlayers int) {
	initGame(nbPlayers)
}

func initGame(nbPlayers int) {
	//cards := deck.NewDeck(deck.Shuffle)
	//pc1 := cards[0]
	//dc1 := cards[1]
	//pc2 := cards[2]
	//dc2 := cards[3]

	//cards = append(cards[:4], cards[4:]...)

	//p1 := player.NewPlayer(pc1, pc2)

}
