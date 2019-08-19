package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var coeff int

// Card from a deck
type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) String() string {
	for _, s := range suits {
		if c.Suit == s {
			return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
		}
	}
	return fmt.Sprintf("%s", c.Rank)
}

func (c Card) absRank() int {
	return int(c.Suit)*coeff + int(c.Rank)
}

// NewDeck given an operation function to perform on them (e.g. sort, shuffle...)
func NewDeck(opt func([]Card), additionalCards ...Card) []Card {
	cards := []Card{}
	for _, s := range suits {
		for i := minRank; i <= maxRank; i++ {
			cards = append(cards, Card{Suit: s, Rank: i})
		}
	}
	cards = append(cards, additionalCards...)

	opt(cards)
	return cards
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
