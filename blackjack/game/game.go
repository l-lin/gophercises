package game

import (
	"fmt"
	"log"
	"os"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/l-lin/gophercises/blackjack/player"
	"github.com/l-lin/gophercises/deck"
)

const (
	nbCardsOnStart = 2
	maxDealerScore = 16
)

// Game represents a blackjack game
type Game struct {
	Dealer  *player.Dealer
	Players []*player.Player
	Cards   []deck.Card
	Score   *Score
	display *display
	turn    *turn
}

// turn represents whether the dealer is playing or which player is playing
type turn struct {
	dealerTurn bool
	playerNb   int
}

func playerTurn(playerNb int) *turn {
	return &turn{dealerTurn: false, playerNb: playerNb}
}

func dealerTurn() *turn {
	return &turn{dealerTurn: true}
}

// New game
func New(nbPlayers int) Game {
	g := Game{
		Score:   newScore(nbPlayers),
		Players: make([]*player.Player, nbPlayers),
		display: newDisplay(nbPlayers),
	}
	g.Cards = deck.NewDeck(deck.Shuffle)
	return g
}

// Run blackjack game
func (g *Game) Run() {
	if err := ui.Init(); err != nil {
		log.Fatalln(err)
	}

	rowPlayers := make([]interface{}, len(g.display.playersWidget))
	for i, pPlayer := range g.display.playersWidget {
		rowPlayers[i] = ui.NewRow(1.0/float64(len(g.display.playersWidget)), pPlayer)
	}

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/4,
			ui.NewCol(1.0/2, g.display.scoreWidget),
			ui.NewCol(1.0/2, g.display.infoWidget),
		),
		ui.NewRow(1.0/4,
			ui.NewCol(1.0, g.display.dealerWidget),
		),
		ui.NewRow(2.0/4,
			rowPlayers...,
		),
	)

	ui.Render(grid)
	g.runRounds(grid)
}

func (g *Game) runRounds(grid *ui.Grid) {
	round := 1
	for len(g.Cards) > nbCardsOnStart*(len(g.Players)+1) {
		g.initRound(len(g.Players))
		g.runRound(round, grid)
		round++
	}
	g.display.renderInfoGameOver(g.Score)
	ui.Render(grid)
	waitKeyPress()
	ui.Close()
	os.Exit(0)
}

func (g *Game) initRound(nbPlayers int) {
	cards := g.Cards
	cards, g.Dealer = initDealer(cards)
	cards, g.Players = initPlayers(cards, nbPlayers)
	g.Cards = cards
}

func (g *Game) runRound(round int, grid *ui.Grid) {
	g.display.renderDealer(g.Dealer)
	// players turn
	g.playersTurn(grid)
	// dealer turn
	g.dealerSetUp()
	if !g.isEveryPlayersOver() {
		for !g.hasDealerFinished() {
			g.display.renderInfoDealerTurn()
			g.display.renderDealer(g.Dealer)
			g.display.renderActive(true, 0)
			ui.Render(grid)
			g.dealerTurn()
			time.Sleep(time.Second * 1)
		}
	}
	// display winner
	g.display.renderDealer(g.Dealer)
	g.display.renderAllInactive()
	ui.Render(grid)
	wPlayerNb, wPlayer, wDealer := g.getWinner()
	if wPlayer != nil {
		g.Score.Players[wPlayerNb]++
	} else if wDealer != nil {
		g.Score.Dealer++
	}
	g.display.renderScore(g.Score, round)
	g.display.renderInfoWinner(wPlayerNb, wPlayer, wDealer)
	ui.Render(grid)

	// wait for player input
	waitKeyPress()
}

func (g *Game) playersTurn(grid *ui.Grid) {
	for i, p := range g.Players {
		playerNb := i + 1
		// no cards left, do not continue
		if len(g.Cards) == 0 {
			p.Finished = true
			continue
		}
		if p.Finished {
			continue
		}
		for !p.Finished {
			g.display.renderInfoPlayerTurn(playerNb)
			g.display.renderPlayers(g.Players)
			g.display.renderActive(false, playerNb)
			ui.Render(grid)
			g.awaitEvent(grid, playerNb, p)
		}
		g.display.renderInfoPlayerTurn(playerNb)
		g.display.renderPlayers(g.Players)
		g.display.renderActive(false, playerNb)
		ui.Render(grid)
	}
}

func (g *Game) awaitEvent(grid *ui.Grid, playerNb int, p *player.Player) {
	uiEvents := ui.PollEvents()
	select {
	case e := <-uiEvents:
		switch e.ID {
		case "q", "<C-c>":
			ui.Close()
			os.Exit(0)
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			grid.SetRect(0, 0, payload.Width, payload.Height)
			ui.Clear()
			ui.Render(grid)
		case "h":
			g.hit(playerNb)
		case "s":
			p.Finished = true
		}
	}
}

func (g *Game) hit(playerNb int) {
	p := g.Players[playerNb-1]
	c, err := g.pickCard()
	if err != nil {
		p.Finished = true
		return
	}
	p.HandCard.Add(*c)
	if p.HandCard.IsOver() {
		p.Finished = true
	}
}

func (g *Game) dealerSetUp() {
	g.Dealer.HandCard.Cards[0].Hidden = false
	score, _ := g.Dealer.Player.HandCard.Compute()
	if score > maxDealerScore {
		g.Dealer.Player.Finished = true
	}
}

func (g *Game) dealerTurn() {
	// no cards left, do not play
	if len(g.Cards) == 0 {
		g.Dealer.Player.Finished = true
		return
	}
	c, err := g.pickCard()
	if err != nil {
		log.Fatal(err)
	}
	g.Dealer.Player.HandCard.Add(*c)
	score, isSoft := g.Dealer.Player.HandCard.Compute()
	if score > maxDealerScore && !isSoft {
		g.Dealer.Player.Finished = true
	}
}

func (g *Game) pickCard() (*deck.Card, error) {
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
	wPlayerNb := 0
	var wPlayer *player.Player
	for i := 0; i < len(g.Players); i++ {
		if g.Players[i].HandCard.IsOver() {
			continue
		}
		if wPlayer == nil {
			wPlayer = g.Players[i]
			wPlayerNb = i
			continue
		}
		if wPlayer.CompareTo(g.Players[i]) < 0 {
			wPlayer = g.Players[i]
			wPlayerNb = i
		}
	}

	// dealer is over
	if g.Dealer.HandCard.IsOver() {
		// player is also over => draw, nobody win
		if wPlayer == nil {
			return -1, nil, nil
		}
		// player wins
		return wPlayerNb + 1, wPlayer, nil
	}
	// player is over => dealer wins
	if wPlayer == nil {
		return -1, nil, g.Dealer
	}

	result := wPlayer.CompareTo(&g.Dealer.Player)
	if result < 0 { // dealer wins
		return -1, nil, g.Dealer
	} else if result == 0 { // player and dealer are not over and it's a draw
		return wPlayerNb + 1, wPlayer, g.Dealer
	}
	// player wins
	return wPlayerNb + 1, wPlayer, nil
}

func (g *Game) isEveryPlayersOver() bool {
	for _, p := range g.Players {
		if !p.HandCard.IsOver() {
			return false
		}
	}
	return true
}

func waitKeyPress() {
	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
