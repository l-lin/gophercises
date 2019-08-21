package deck

// FilterOut card by given card
func FilterOut(f func(Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		result := make([]Card, 0)
		for _, c := range cards {
			if !f(c) {
				result = append(result, c)
			}
		}

		return result
	}
}
