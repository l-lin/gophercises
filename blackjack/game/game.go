package game

import (
	"fmt"
	"log"
	"time"

	"github.com/l-lin/gophercises/blackjack/player"
	"github.com/l-lin/gophercises/deck"
	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
)

// Game represents a blackjack game
type Game struct {
	Dealer  *player.Dealer
	Players []*player.Player
	Cards   []deck.Card
	Score   *Score
}

// Score of the players and dealer
type Score struct {
	Players map[int]int
	Dealer  int
}

const (
	nbCardsOnStart = 2
	maxNbCards     = 10
)

// New game
func New(nbPlayers int) Game {
	g := Game{Score: newScore(nbPlayers)}
	g.Cards = deck.NewDeck(deck.Shuffle)
	return g
}

// Run blackjack game
func (g *Game) Run(nbPlayers int) {
	round := 1
	for len(g.Cards) > maxNbCards {
		time.Sleep(time.Second * 1)
		g.init(nbPlayers)
		g.runRound(round)
		round++
		g.Score.diplay()
	}
}

func (g *Game) runRound(round int) {
	fmt.Println(aurora.Sprintf(aurora.BrightBlue("ROUND %d").Bold(), round))
	fmt.Println(aurora.BrightBlack("PLAYERS TURN").BgBrightBlue().Bold())
	for !g.haveAllPlayerFinished() {
		g.displayCards()
		g.playersTurn()
	}
	fmt.Println(aurora.BrightBlack("DEALER TURN").BgBrightYellow().Bold())
	g.dealerSetUp()
	for !g.hasDealerFinished() {
		time.Sleep(time.Second * 1)
		g.displayDealerCards()
		g.dealerTurn()
	}
	time.Sleep(time.Second * 1)
	g.displayCards()
	nbWPlayer, wPlayer, wDealer := g.getWinner()
	if wPlayer != nil {
		fmt.Println(aurora.BrightBlack(fmt.Sprintf("PLAYER %d WINS!", nbWPlayer)).BgBrightBlue().Bold())
		g.Score.Players[nbWPlayer]++
	} else if wDealer != nil {
		fmt.Println(aurora.BrightBlack("DEALER WINS").BgBrightYellow().Bold())
		g.Score.Dealer++
	} else {
		fmt.Println(aurora.BrightBlack("DRAW").BgBrightWhite().Bold())
	}
}

func (g *Game) init(nbPlayers int) {
	cards := g.Cards
	cards, g.Dealer = initDealer(cards)
	cards, g.Players = initPlayers(cards, nbPlayers)
	g.Cards = cards
}

func (g *Game) displayCards() {
	g.displayDealerCards()
	g.displayPlayerCards()
}

func (g *Game) displayDealerCards() {
	if g.Dealer.HandCard.Cards[0].Hidden {
		fmt.Printf("Dealer:\n%s\n", g.Dealer.HandCard.Print())
	} else {
		fmt.Printf("Dealer (%d points):\n%s\n", g.Dealer.HandCard.Compute(), g.Dealer.HandCard.Print())
	}
}

func (g *Game) displayPlayerCards() {
	for j, p := range g.Players {
		fmt.Printf("Player %d (%d points):\n%s\n", j+1, p.HandCard.Compute(), p.HandCard.Print())
	}
}

func (g *Game) playersTurn() {
	for i, p := range g.Players {
		playerNb := i + 1
		if p.Finished {
			continue
		}
		prompt := promptui.Select{
			Label: fmt.Sprintf("Player %d: What's your choice?", playerNb),
			Items: []string{"Hit", "Stand"},
		}
		_, choice, err := prompt.Run()
		if err != nil {
			log.Fatal(err)
		}

		if choice == "Hit" {
			c, err := g.hit()
			if err != nil {
				log.Fatal(err)
			}
			p.HandCard.Add(*c)
			fmt.Printf("Player %d, you picked the following which gives you %d points:\n%s\n", playerNb, p.HandCard.Compute(), c.Print())
			if p.HandCard.IsOver() {
				fmt.Println(aurora.Sprintf(aurora.BrightRed("Player %d, you have exceeded max score! You lose!"), playerNb))
				p.Finished = true
			}
		} else {
			p.Finished = true
		}
	}
}

func (g *Game) dealerSetUp() {
	g.Dealer.HandCard.Cards[0].Hidden = false
	if g.Dealer.Player.HandCard.Compute() > 16 {
		g.Dealer.Player.Finished = true
	}
}

func (g *Game) dealerTurn() {
	c, err := g.hit()
	if err != nil {
		log.Fatal(err)
	}
	g.Dealer.Player.HandCard.Add(*c)
	fmt.Printf("Dealer has picked the following which gives him %d points:\n%s\n", g.Dealer.Player.HandCard.Compute(), c.Print())
	if g.Dealer.Player.HandCard.Compute() > 16 {
		g.Dealer.Player.Finished = true
	}
	time.Sleep(time.Second * 1)
}

func (g *Game) hit() (*deck.Card, error) {
	if len(g.Cards) == 0 {
		return nil, fmt.Errorf("no cards left")
	}
	c := g.Cards[0]
	g.Cards = g.Cards[1:]
	return &c, nil
}

func (g *Game) haveAllPlayerFinished() bool {
	for _, p := range g.Players {
		if !p.Finished {
			return false
		}
	}
	return true
}

func (g *Game) hasDealerFinished() bool {
	return g.Dealer.Player.Finished
}

func (g *Game) getWinner() (int, *player.Player, *player.Dealer) {
	// get player winner
	nbWPlayer := 0
	var wPlayer *player.Player
	for i := 0; i < len(g.Players); i++ {
		if g.Players[i].HandCard.IsOver() {
			continue
		}
		if wPlayer == nil {
			wPlayer = g.Players[i]
			nbWPlayer = i
			continue
		}
		if wPlayer.CompareTo(g.Players[i]) < 0 {
			wPlayer = g.Players[i]
			nbWPlayer = i
		}
	}

	// dealer is over
	if g.Dealer.HandCard.IsOver() {
		// player is also over => draw, nobody win
		if wPlayer == nil {
			return 0, nil, nil
		}
		// player wins
		return nbWPlayer + 1, wPlayer, nil
	}
	// player is over => dealer wins
	if wPlayer == nil {
		return 0, nil, g.Dealer
	}

	result := wPlayer.CompareTo(&g.Dealer.Player)
	if result < 1 { // dealer wins
		return 0, nil, g.Dealer
	} else if result == 0 { // player and dealer are not over and it's a draw
		return nbWPlayer + 1, wPlayer, g.Dealer
	}
	// player wins
	return nbWPlayer + 1, wPlayer, nil
}

func newScore(nbPlayers int) *Score {
	scores := make(map[int]int, nbPlayers)
	for i := 1; i <= nbPlayers; i++ {
		scores[i] = 0
	}
	return &Score{Players: scores, Dealer: 0}
}
func (s *Score) diplay() {
	fmt.Println("")
	for i, v := range s.Players {
		fmt.Println(aurora.BrightBlack(fmt.Sprintf("Player %d: %d wins", i, v)).BgBrightBlue().Bold())
	}
	fmt.Println(aurora.BrightBlack(fmt.Sprintf("Dealer: %d wins", s.Dealer)).BgBrightYellow().Bold())
	fmt.Println("")
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
