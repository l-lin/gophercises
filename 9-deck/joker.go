package main

// AddJokers adds jokers to the deck
func AddJokers(cards []Card) []Card {
	cards = append(cards, Card{Suit: BlackJoker})
	cards = append(cards, Card{Suit: RedJoker})
	return cards
}
