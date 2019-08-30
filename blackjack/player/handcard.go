package player

import (
	"math"
	"strings"

	"github.com/l-lin/gophercises/deck"
)

const topScore = 21

// HandCard of a player
type HandCard struct {
	Cards []deck.Card
}

// NewHandCard instanciates a new handcard
func NewHandCard(cards ...deck.Card) HandCard {
	c := []deck.Card{}
	c = append(c, cards...)
	return HandCard{
		Cards: c,
	}
}

// Add a new card to the handcard
func (h HandCard) Add(c deck.Card) {
	h.Cards = append(h.Cards, c)
}

// IsBlackJack checks if the current cards are a blackjack
// combinaison which 2 cards & Ace + face card (J/Q/K)
func (h HandCard) IsBlackJack() bool {
	return len(h.Cards) == 2 && h.compute() == topScore
}

// IsOver returns true if the score is over 21
func (h HandCard) IsOver() bool {
	return h.compute() > topScore
}

// CompareTo to h2 in term of being closest to the topScore
// without being over the topScore
// Returns:
//   1 if h1 is closest to topScore
//   -1 if h2 is the one closest
//   0 if both h1 & h2 have the same number of cards and are
//     both close to topScore or both are over the topScore
func (h HandCard) CompareTo(to HandCard) int {
	if h.IsOver() {
		if to.IsOver() {
			return 0
		}
		return -1
	}
	if to.IsOver() {
		return 1
	}
	val1 := h.compute()
	val2 := to.compute()
	final1 := math.Abs(float64(topScore - val1))
	final2 := math.Abs(float64(topScore - val2))

	if final1 == final2 {
		if len(h.Cards) == len(to.Cards) {
			return 0
		} else if len(h.Cards) < len(to.Cards) {
			return 1
		}
		return -1
	}
	if final2-final1 > 0 {
		return 1
	}
	return -1
}

// Print the cards in ASCII art with colors
func (h HandCard) Print() string {
	var b strings.Builder
	for i := 0; i < len(h.Cards); i++ {
		b.WriteString(h.Cards[i].Print())
		b.WriteString(" ")
	}
	return b.String()
}

func (h HandCard) compute() int {
	result := 0
	nbAces := 0
	for _, c := range h.Cards {
		var val int
		if c.Rank == deck.Jack || c.Rank == deck.Queen || c.Rank == deck.King {
			val = 10
		} else if c.Rank == deck.Ace {
			nbAces++
			val = 0
		} else {
			val = int(c.Rank)
		}
		result += val
	}
	if nbAces > 0 {
		for i := 0; i < nbAces; i++ {
			tmp := 11
			if result+tmp > topScore {
				tmp = 1
			}
			result += tmp
		}
	}

	return result
}
