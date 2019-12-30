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
		},
	}
	g.Generate()
}
