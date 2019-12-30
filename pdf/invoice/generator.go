package invoice

import (
	"log"

	"github.com/jung-kurt/gofpdf"
)

// Generator invoice in PDF
type Generator struct {
	CompanyDetails
	CompanyAddress
}

// Generate invoice in PDF
func (g *Generator) Generate() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeA4, "")
	pdf.SetTopMargin(70)
	pdf.SetHeaderFuncMode(g.headerFn(pdf), true)
	pdf.SetFooterFunc(g.footerFn(pdf))
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")
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
		pdf.MoveTo(230, 10)
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
