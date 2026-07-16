package pdf

import (
	"html-to-gofpdf/fonts"
	"html-to-gofpdf/models"
	"html-to-gofpdf/renderer"
	"html-to-gofpdf/templates/kfs"

	"github.com/jung-kurt/gofpdf"
)

// GeneratePDF creates a new PDF document matching the request details and saves it to outputFile.
func GeneratePDF(req models.PDFRequest, outputFile string) error {
	// Initialize a standard Portrait, millimeters, A4 PDF instance.
	// Empty font directory since we use full filesystem paths in AddUTF8Font.
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Disable automatic page breaks so each page is rendered strictly according to the HTML template sizes.
	pdf.SetAutoPageBreak(false, 0)

	// Create the Font Manager.
	fm := fonts.NewFontManager()

	// 1. Page 1 (Part 1 Table)
	ml1, mt1, mr1, mb1 := 14.17, 34.84, 8.82, 19.78
	pdf.SetMargins(ml1, mt1, mr1)
	pdf.AddPage()
	ctx1 := renderer.NewContext(pdf, req.Language, fm)
	ctx1.MarginBottom = mb1

	if err := kfs.RenderPage1(ctx1, req); err != nil {
		return err
	}

	// 2. Page 2 (Qualitative Information Table)
	ml2, mt2, mr2, mb2 := 14.17, 50.94, 8.82, 61.94
	pdf.SetMargins(ml2, mt2, mr2)
	pdf.AddPage()
	ctx2 := renderer.NewContext(pdf, req.Language, fm)
	ctx2.MarginBottom = mb2

	if err := kfs.RenderPage2(ctx2, req); err != nil {
		return err
	}

	// 3. Page 3 (Annex B Computations Table)
	ml3, mt3, mr3, mb3 := 14.17, 53.49, 8.82, 65.07
	pdf.SetMargins(ml3, mt3, mr3)
	pdf.AddPage()
	ctx3 := renderer.NewContext(pdf, req.Language, fm)
	ctx3.MarginBottom = mb3

	if err := kfs.RenderPage3(ctx3, req); err != nil {
		return err
	}

	// Save output file.
	return pdf.OutputFileAndClose(outputFile)
}
