package invoice

import "fmt"

// CompanyAddress contains the address of the company
type CompanyAddress struct {
	Number  int
	Street  string
	ZipCode string
	Country string
}

func (c CompanyAddress) String() string {
	return fmt.Sprintf("%d %s\n%s\n%s", c.Number, c.Street, c.Country, c.ZipCode)
}

// CompanyDetails contains the details of the company
type CompanyDetails struct {
	Email  string
	Phone  string
	Domain string
}

func (c CompanyDetails) String() string {
	return fmt.Sprintf("%s\n%s\n%s", c.Phone, c.Email, c.Domain)
}

// Bill represent the overview representation of a bill
type Bill struct {
	ClientName    string
	ClientAddress CompanyAddress
	InvoiceNumber string
	DateOfIssue   string
	// nothing fancy here and complicated, we will just prepend the currency with the price
	Currency string
	Units    []Unit
}

// FullClientAddress returns the complete information of the client in string
func (b Bill) FullClientAddress() string {
	return fmt.Sprintf("%s\n%s", b.ClientName, b.ClientAddress.String())
}

// InvoiceTotal computes the total of the invoice due by the client
func (b Bill) InvoiceTotal() (total float64) {
	for _, unit := range b.Units {
		total = total + float64(unit.PricePerUnit*unit.UnitsPurchased)
	}
	return total / 100
}
