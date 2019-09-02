package deck

import (
	"github.com/logrusorgru/aurora"
)

// Suit represents the type of card
// Setting all properties to private to make it immutable
type Suit struct {
	value         int
	name          string
	color         func(interface{}) aurora.Value
	hasRank       bool
	asciiTemplate []string
}

func (s Suit) String() string {
	return s.name
}

// Value of the suit
func (s Suit) Value() int {
	return s.value
}

// Color of the suit
func (s Suit) Color(arg interface{}) aurora.Value {
	return s.color(arg)
}

// ASCIITemplate of the card suit in ASCII
func (s Suit) ASCIITemplate() []string {
	return s.asciiTemplate
}

// Equals checks if the given suit is the same or not
func (s Suit) Equals(to Suit) bool {
	return s.value == to.value && s.name == to.name
}

// HasRank specifies if the suit use card rank
func (s Suit) HasRank() bool {
	return s.hasRank
}

var (
	// Spade card type
	Spade = Suit{
		value:   1,
		name:    "Spade",
		color:   aurora.White,
		hasRank: true,
		asciiTemplate: []string{
			`┌────────┐`,
			`│%s .    │`,
			`│  / \   │`,
			`│ (_,_)  │`,
			`│   I  %s│`,
			`└────────┘`,
		},
	}
	// Diamond card type
	Diamond = Suit{
		value:   2,
		name:    "Diamond",
		color:   aurora.BrightRed,
		hasRank: true,
		asciiTemplate: []string{
			`┌────────┐`,
			`│%s /\   │`,
			`│  /  \  │`,
			`│  \  /  │`,
			`│   \/ %s│`,
			`└────────┘`,
		},
	}
	// Club card type
	Club = Suit{
		value:   3,
		name:    "Club",
		color:   aurora.White,
		hasRank: true,
		asciiTemplate: []string{
			`┌────────┐`,
			`│%s _    │`,
			`│  ( )   │`,
			`│ (_x_)  │`,
			`│   Y  %s│`,
			`└────────┘`,
		},
	}
	// Hearth card type
	Hearth = Suit{
		value:   4,
		name:    "Hearth",
		color:   aurora.BrightRed,
		hasRank: true,
		asciiTemplate: []string{
			`┌────────┐`,
			`│%s_  _  │`,
			`│ ( \/ ) │`,
			`│  \  /  │`,
			`│   \/ %s│`,
			`└────────┘`,
		},
	}
	// BlackJoker card type
	BlackJoker = Suit{
		value:   5,
		name:    "BlackJoker",
		color:   aurora.White,
		hasRank: false,
		asciiTemplate: []string{
			`┌────────┐`,
			`│* \||/ K│`,
			`│J /~~\ O│`,
			`│O( o o)J│`,
			`│K \ v/ *│`,
			`└────────┘`,
		},
	}
	// RedJoker card type
	RedJoker = Suit{
		value:   6,
		name:    "RedJoker",
		color:   aurora.BrightRed,
		hasRank: false,
		asciiTemplate: []string{
			`┌────────┐`,
			`│+ \||/ K│`,
			`│J /~~\ O│`,
			`│O( o o)J│`,
			`│K \ v/ +│`,
			`└────────┘`,
		},
	}
	suits        = [...]Suit{Spade, Diamond, Club, Hearth}
	cardTemplate = []string{
		`┌────────┐`,
		`│████████│`,
		`│████████│`,
		`│████████│`,
		`│████████│`,
		`└────────┘`,
	}
)
