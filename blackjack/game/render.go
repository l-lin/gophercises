package game

import (
	"fmt"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/l-lin/gophercises/blackjack/player"
)

type display struct {
	scoreWidget   *widgets.Paragraph
	infoWidget    *widgets.Paragraph
	dealerWidget  *widgets.Paragraph
	playersWidget []*widgets.Paragraph
}

func newDisplay(nbPlayers int) *display {
	s := newScore(nbPlayers)
	scoreWidget := widgets.NewParagraph()
	scoreWidget.Title = "Round 1"
	scoreWidget.Text = s.print()

	infoWidget := widgets.NewParagraph()
	infoWidget.Text = ""

	dealerWidget := widgets.NewParagraph()
	dealerWidget.Title = "Dealer"
	dealerWidget.Border = false

	playersWidget := make([]*widgets.Paragraph, nbPlayers)

	for i := 0; i < nbPlayers; i++ {
		p := widgets.NewParagraph()
		p.Title = fmt.Sprintf("Player %d", i+1)
		p.Border = false
		playersWidget[i] = p
	}

	return &display{
		scoreWidget:   scoreWidget,
		infoWidget:    infoWidget,
		dealerWidget:  dealerWidget,
		playersWidget: playersWidget,
	}
}

func (d *display) renderDealer(dealer *player.Dealer) {
	if dealer.HandCard.Cards[0].Hidden {
		d.dealerWidget.Title = "Dealer"
	} else {
		dealerScore, _ := dealer.HandCard.Compute()
		d.dealerWidget.Title = fmt.Sprintf("Dealer (%d points)", dealerScore)
	}
	d.dealerWidget.Text = dealer.HandCard.ToASCII()
}

func (d *display) renderPlayers(players []*player.Player) {
	for i, p := range players {
		score, _ := p.HandCard.Compute()
		d.playersWidget[i].Title = fmt.Sprintf("Player %d (%d points)", i+1, score)
		d.playersWidget[i].Text = p.HandCard.ToASCII()
	}
}

func (d *display) renderActive(dealerTurn bool, playerNb int) {
	d.renderAllInactive()
	if dealerTurn {
		d.dealerWidget.TitleStyle.Fg = ui.ColorGreen
		d.dealerWidget.TextStyle.Fg = ui.ColorGreen
	} else {
		d.playersWidget[playerNb-1].TitleStyle.Fg = ui.ColorGreen
		d.playersWidget[playerNb-1].TextStyle.Fg = ui.ColorGreen
	}
}

func (d *display) renderAllInactive() {
	d.dealerWidget.TitleStyle.Fg = ui.ColorWhite
	d.dealerWidget.TextStyle.Fg = ui.ColorWhite
	for _, playerWidget := range d.playersWidget {
		playerWidget.TitleStyle.Fg = ui.ColorWhite
		playerWidget.TextStyle.Fg = ui.ColorWhite
	}
}

func (d *display) renderScore(s *Score, round int) {
	d.scoreWidget.Title = fmt.Sprintf("Round %d", round)
	d.scoreWidget.Text = s.print()
}

func (d *display) renderInfoPlayerTurn(playerNb int) {
	d.infoWidget.Text = fmt.Sprintf(`[Player %d turn](fg:blue,mod:bold)
h: HIT
s: STAND
q: QUIT`, playerNb)
}

func (d *display) renderInfoDealerTurn() {
	d.infoWidget.Text = "[Dealer turn](fg:yellow,mod:bold)"
}

func (d *display) renderInfoWinner(wPlayerNb int, wPlayer *player.Player, wDealer *player.Dealer) {
	var b strings.Builder
	if wDealer != nil {
		b.WriteString("[Dealer wins](fg:yellow,mod:bold)")
	} else if wPlayer != nil {
		b.WriteString(fmt.Sprintf("[Player %d wins](fg:blue,mod:bold)", wPlayerNb))
	} else {
		b.WriteString("[DRAW](fg:white,mod:bold)")
	}
	b.WriteString("\nPress any key to start a new round")
	d.infoWidget.Text = b.String()
}

func (d *display) renderInfoGameOver(s *Score) {
	var b strings.Builder
	b.WriteString("[GAME OVER](fg:red,mod:bold)\n")
	b.WriteString("----------------------------\n")
	isDealerWinner := true
	wPlayerNb := 1
	for i, playerScore := range s.Players {
		if s.Players[wPlayerNb] < playerScore {
			wPlayerNb = i
		}
	}
	if s.Dealer < s.Players[wPlayerNb] {
		isDealerWinner = false
	}
	if isDealerWinner {
		if s.Dealer == s.Players[wPlayerNb] {
			b.WriteString(fmt.Sprintf("[DRAW BETWEEN DEALER AND PLAYER %d](fg:white,mod:bold)", wPlayerNb))
		} else {
			b.WriteString("[DEALER WINS](fg:yellow,mod:bold)")
		}
	} else {
		b.WriteString(fmt.Sprintf("[PLAYER %d WINS](fg:blue,mod:bold)", wPlayerNb))
	}
	d.infoWidget.Text = b.String()
}
