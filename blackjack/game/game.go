package game

import (
	"fmt"

	"github.com/l-lin/gophercises/deck"
)

// Run blackjack game
func Run() {
	d := deck.NewDeck(deck.Shuffle)
	fmt.Println(d)

}
