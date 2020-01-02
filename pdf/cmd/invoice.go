package cmd

import (
	"github.com/l-lin/gophercises/pdf/invoice"
	"github.com/spf13/cobra"
)

var invoiceCmd = &cobra.Command{
	Use:   "invoice",
	Short: "Generate invoice PDF",
	Run:   runInvoice,
}

func init() {
	rootCmd.AddCommand(invoiceCmd)
}

func runInvoice(cmd *cobra.Command, args []string) {
	g := invoice.Generator{
		CompanyDetails: invoice.CompanyDetails{
			Phone:  "0123456789",
			Email:  "l-lin@foobar.com",
			Domain: "foobar.com",
		},
		CompanyAddress: invoice.CompanyAddress{
			Number:  123,
			Street:  "Foobar street",
			ZipCode: "12345",
			Country: "France",
		},
		Bill: invoice.Bill{
			ClientName: "Someone Giving Money",
			ClientAddress: invoice.CompanyAddress{
				Number:  321,
				Street:  "Client street",
				ZipCode: "54321",
				Country: "US",
			},
			InvoiceNumber: "0000000123",
			DateOfIssue:   "05/09/2018",
		},
	}
	g.Generate()
}
