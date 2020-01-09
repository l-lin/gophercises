package invoice

// Unit represents a line in the bill
type Unit struct {
	UnitName       string
	PricePerUnit   int // in cent
	UnitsPurchased int
}

// Amount computes the total cost of the unit
func (u Unit) Amount() float64 {
	return float64(u.PricePerUnit*u.UnitsPurchased) / 100
}
