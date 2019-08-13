package main

import (
	"sort"
)

// DefaultSort cards
func DefaultSort(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].absRank() < cards[j].absRank()
	})
}
