//go:generate stringer -type=Suit
package deck

// Suit represents the type of card
type Suit int

const (
	// Spade card type
	Spade Suit = iota
	// Diamond card type
	Diamond
	// Club card type
	Club
	// Hearth card type
	Hearth
	// RedJoker card type
	RedJoker
	// BlackJoker card type
	BlackJoker
)

var suits = [...]Suit{Spade, Diamond, Club, Hearth}
