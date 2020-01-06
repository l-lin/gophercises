package invoice

// Unit represents a line in the bill
type Unit struct {
	UnitName       string
	PricePerUnit   int // in cent
	UnitsPurchased int
}
