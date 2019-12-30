package invoice

import "fmt"

// CompanyAddress contains the address of the company
type CompanyAddress struct {
	Number  int
	Street  string
	ZipCode string
}

func (c *CompanyAddress) String() string {
	return fmt.Sprintf("%d %s\n%s", c.Number, c.Street, c.ZipCode)
}

// CompanyDetails contains the details of the company
type CompanyDetails struct {
	Email  string
	Phone  string
	Domain string
}

func (c *CompanyDetails) String() string {
	return fmt.Sprintf("%s\n%s\n%s", c.Phone, c.Email, c.Domain)
}
