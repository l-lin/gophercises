//go:generate stringer -type=Rank
package deck

import "strconv"

// Rank of a card
type Rank int

const (
	_ Rank = iota
	// Ace of card
	Ace
	// Two of card
	Two
	// Three of card
	Three
	// Four of card
	Four
	// Five of card
	Five
	// Six of card
	Six
	// Seven of card
	Seven
	// Eight of card
	Eight
	// Nine of card
	Nine
	// Ten of card
	Ten
	// Jack of card
	Jack
	// Queen of card
	Queen
	// King of card
	King
)

const (
	minRank = Ace
	maxRank = King
)

// Single character of the rank
func (r Rank) Single() string {
	if r > Ace && r < Jack {
		return strconv.Itoa(int(r))
	}
	if r == Ace {
		return "A"
	}
	return string(r.String()[0])
}
