package game

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gosuri/uilive"
	"github.com/l-lin/gophercises/blackjack/player"
	"github.com/l-lin/gophercises/deck"
	"github.com/tcnksm/go-input"
)

// Game represents a blackjack game
type Game struct {
	Dealer  *player.Dealer
	Players []*player.Player
	Cards   []deck.Card
}

const nbCardsOnStart = 2

// New game
func New(nbPlayers int) Game {
	g := Game{}
	g.init(nbPlayers)
	return g
}

// Run blackjack game
func (g Game) Run() {
	writer := uilive.New()
	writer.Start()
	for i := 0; i < 10; i++ {
		g.displayCards(writer)
		g.playersTurn(writer)
	}
	writer.Stop()
}

func (g *Game) init(nbPlayers int) {
	cards := deck.NewDeck(deck.Shuffle)
	cards, g.Dealer = initDealer(cards)
	cards, g.Players = initPlayers(cards, nbPlayers)
	g.Cards = cards
}

func (g Game) displayCards(w io.Writer) {
	fmt.Fprintf(w, "Dealer:\n%s\n", g.Dealer.HandCard.ToASCII())
	for j, p := range g.Players {
		fmt.Fprintf(w, "Player %d:\n%s\n", j, p.HandCard.ToASCII())
	}
}

func (g *Game) playersTurn(w io.Writer) {
	for _, p := range g.Players {
		ui := &input.UI{
			Writer: w,
			Reader: os.Stdin,
		}
		choice, err := ui.Select("What's your choice?", []string{"Hit", "Stand"}, &input.Options{
			Default:  "Hit",
			Required: true,
			Loop:     true,
		})
		if err != nil {
			log.Fatal(err)
		}

		if choice == "Hit" {
			c, err := g.hit()
			if err != nil {
				log.Fatal(err)
			}
			p.HandCard.Add(*c)
		}
	}
}

func (g *Game) hit() (*deck.Card, error) {
	if len(g.Cards) == 0 {
		return nil, fmt.Errorf("no cards left")
	}
	c := g.Cards[0]
	g.Cards = g.Cards[1:]
	return &c, nil
}

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
