// Code generated by "stringer -type=Suit"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Spade-0]
	_ = x[Diamond-1]
	_ = x[Club-2]
	_ = x[Hearth-3]
	_ = x[RedJoker-4]
	_ = x[BlackJoker-5]
}

const _Suit_name = "SpadeDiamondClubHearthRedJokerBlackJoker"

var _Suit_index = [...]uint8{0, 5, 12, 16, 22, 30, 40}

func (i Suit) String() string {
	if i < 0 || i >= Suit(len(_Suit_index)-1) {
		return "Suit(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Suit_name[_Suit_index[i]:_Suit_index[i+1]]
}