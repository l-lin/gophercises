package invoice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

const fileName = "invoice.pdf"

// Generator invoice in PDF
type Generator struct {
	CompanyDetails
	CompanyAddress
	Bill
}

// Generate invoice in PDF
func (g *Generator) Generate() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeA4, "")
	pdf.SetTopMargin(100)
	pdf.SetHeaderFuncMode(g.headerFn(pdf), true)
	pdf.SetFooterFunc(g.footerFn(pdf))
	pdf.AddPage()

	g.invoiceHeader(pdf)

	// Write PDF file
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Generator) headerFn(pdf *gofpdf.Fpdf) func() {
	return func() {
		w, _ := pdf.GetPageSize()
		pdf.SetXY(0, 0)
		// Fill color
		pdf.SetFont("Arial", "B", 30)
		pdf.SetFillColor(64, 64, 64)
		pdf.CellFormat(w, 60, "", "", 0, "LM", true, 0, "")
		// Image
		pdf.ImageOptions("images/header.jpg", 250, 3, 0, 55, false, gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: "",
		}, 0, "")
		// Invoice text
		pdf.SetTextColor(255, 255, 255)
		pdf.Text(50, 42, "INVOICE")
		// Details
		pdf.SetFont("Arial", "", 10)
		pdf.SetTextColor(255, 255, 255)
		pdf.MoveTo(130, 0)
		pdf.MultiCell(330, 20, g.CompanyDetails.String(), gofpdf.BorderNone, gofpdf.AlignRight, false)
		// Address
		pdf.SetFont("Arial", "", 10)
		pdf.SetTextColor(255, 255, 255)
		pdf.MoveTo(230, 0)
		pdf.MultiCell(330, 20, g.CompanyAddress.String(), gofpdf.BorderNone, gofpdf.AlignRight, false)
	}
}

func (g *Generator) footerFn(pdf *gofpdf.Fpdf) func() {
	return func() {
		w, h := pdf.GetPageSize()
		pdf.SetXY(0, h-60)
		// Fill color
		pdf.SetFont("Arial", "B", 30)
		pdf.SetFillColor(64, 64, 64)
		pdf.CellFormat(w, 60, "", "", 0, "LM", true, 0, "")
	}
}

func (g *Generator) invoiceHeader(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 16)
	pdf.SetTextColor(180, 180, 180)
	// First line
	pdf.Cell(200, 10, "Billed To")
	pdf.Cell(250, 10, "Invoice Number")
	pdf.Cell(200, 10, "Invoice Total")
	// Client address
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 0)
	pdf.MoveTo(30, 120)
	pdf.MultiCell(200, 20, g.Bill.FullClientAddress(), gofpdf.BorderNone, gofpdf.AlignLeft, false)
	// Invoice number
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 0)
	pdf.MoveTo(230, 123)
	pdf.Cell(200, 10, g.Bill.InvoiceNumber)
	// Date of issue
	pdf.SetFont("Arial", "B", 16)
	pdf.SetTextColor(180, 180, 180)
	pdf.MoveTo(230, 165)
	pdf.Cell(200, 0, "Date of Issue")
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 0)
	pdf.MoveTo(230, 184)
	pdf.Cell(200, 10, g.Bill.DateOfIssue)
	// Invoice total
	pdf.SetFont("Arial", "B", 50)
	pdf.SetTextColor(64, 64, 64)
	pdf.MoveTo(368, 155)
	pdf.Cell(400, 0, fmt.Sprintf("%s%s", g.Bill.Currency, strconv.FormatFloat(g.Bill.InvoiceTotal(), 'f', 2, 64)))
	// Line break
	w, _ := pdf.GetPageSize()
	pdf.SetFont("Arial", "B", 30)
	pdf.SetFillColor(64, 64, 64)
	pdf.MoveTo(20, 230)
	pdf.CellFormat(w-40, 5, "", "", 0, "LM", true, 0, "")
}
