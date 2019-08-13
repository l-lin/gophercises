package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := NewDeck(func(cards []Card) {})
	expectedCardsNb := 52
	if len(cards) != expectedCardsNb {
		t.Errorf("not enough cards, expected %d, got %d", expectedCardsNb, len(cards))
	}
}
