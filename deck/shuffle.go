package deck

import (
	"math/rand"
	"time"
)

// Shuffle cards
func Shuffle(cards []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	result := make([]Card, len(cards))
	perm := r.Perm(len(cards))
	for i, randIndex := range perm {
		result[i] = cards[randIndex]
	}
	copy(cards, result)
	return result
}
