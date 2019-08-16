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

func TestComputeCoeff(t *testing.T) {
	var tests = map[string]struct {
		given    int
		expected int
	}{
		"13":    {13, 100},
		"99":    {99, 100},
		"1":     {1, 10},
		"81234": {81234, 100000},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := computeCoeff(tt.given)
			if actual != tt.expected {
				t.Errorf("(%d): expected %d, actual %d", tt.given, tt.expected, actual)
			}
		})
	}
}
