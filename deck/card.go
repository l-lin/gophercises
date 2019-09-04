package deck

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/logrusorgru/aurora"
)

var coeff int

// Card from a deck
type Card struct {
	Suit   Suit
	Rank   Rank
	Hidden bool
}

func (c Card) String() string {
	for _, s := range suits {
		if c.Suit.Equals(s) {
			return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
		}
	}
	return fmt.Sprintf("%s", c.Suit)
}

// Equals checks if the given card is the same or not
func (c Card) Equals(to Card) bool {
	return c.Suit.Equals(to.Suit) && c.Rank == to.Rank
}

// ToASCII renders the card in ASCII art
func (c Card) ToASCII() []string {
	ascii := []string{}
	if c.Hidden {
		ascii = cardTemplate
	} else if c.Suit.HasRank() {
		top := c.Rank.Single()
		bottom := c.Rank.Single()
		if c.Rank != Ten {
			top = fmt.Sprintf("%s ", top)
			bottom = fmt.Sprintf(" %s", bottom)
		}
		for i, t := range c.Suit.ASCIITemplate() {
			switch i {
			case 1:
				ascii = append(ascii, fmt.Sprintf(t, top))
			case 4:
				ascii = append(ascii, fmt.Sprintf(t, bottom))
			default:
				ascii = append(ascii, t)
			}
		}
	} else {
		ascii = c.Suit.ASCIITemplate()
	}
	return ascii
}

// Print card in ASCII with color
func (c Card) Print() string {
	return Print([]Card{c})
}

func (c Card) absRank() int {
	return c.Suit.Value()*coeff + int(c.Rank)
}

// NewDeck given an operation function to perform on them (e.g. sort, shuffle...)
func NewDeck(opts ...func([]Card) []Card) []Card {
	cards := []Card{}
	for _, s := range suits {
		for i := minRank; i <= maxRank; i++ {
			cards = append(cards, Card{Suit: s, Rank: i})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

// FromDecks constructs a single deck composed of multiple decks
func FromDecks(decks ...[]Card) []Card {
	cards := make([]Card, 0)
	for _, d := range decks {
		for _, c := range d {
			cards = append(cards, c)
		}
	}
	return cards
}

// ToASCII renders the cards in ASCII art
func ToASCII(cards []Card) string {
	if len(cards) == 0 {
		return ""
	}
	var b strings.Builder
	templates := [][]string{}
	for _, c := range cards {
		templates = append(templates, c.ToASCII())
	}
	for i := 0; i < len(templates[0]); i++ {
		for _, t := range templates {
			b.WriteString(t[i])
		}
		b.WriteString("\n")
	}
	return strings.Trim(b.String(), "\n")
}

// Print cards in color & ASCII art
func Print(cards []Card) string {
	if len(cards) == 0 {
		return ""
	}
	var b strings.Builder
	templates := [][]string{}
	for _, c := range cards {
		templates = append(templates, c.ToASCII())
	}
	for i := 0; i < len(templates[0]); i++ {
		for j, t := range templates {
			if !cards[j].Hidden {
				b.WriteString(aurora.Sprintf(cards[j].Suit.Color(t[i])))
			} else {
				b.WriteString(t[i])
			}
		}
		b.WriteString("\n")
	}
	return strings.Trim(b.String(), "\n")
}

func init() {
	coeff = computeCoeff(int(maxRank))
}

func computeCoeff(base int) int {
	str := strconv.Itoa(base * 10)
	var buff strings.Builder
	buff.WriteString("1")
	for i := 1; i < len(str); i++ {
		buff.WriteString("0")
	}
	result, err := strconv.Atoi(buff.String())
	if err != nil {
		log.Fatalln(err)
	}
	return result
}
