package services

import "github.com/jung-kurt/gofpdf"

func GeneratePDF(filePath string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Invoice")
	pdf.Ln(20)

	// TODO: Add job and user data

	return pdf.OutputFileAndClose(filePath)
}
