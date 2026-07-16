package pdf

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	"html-to-gofpdf/language"
	"html-to-gofpdf/models"
)

// createDummyPNG creates a small blue PNG in memory for testing image rendering.
func createDummyPNG() []byte {
	img := image.NewRGBA(image.Rect(0, 0, 120, 20))
	for x := 0; x < 120; x++ {
		for y := 0; y < 20; y++ {
			img.Set(x, y, color.RGBA{0, 102, 204, 255}) // blue
		}
	}
	var buf bytes.Buffer
	_ = png.Encode(&buf, img)
	return buf.Bytes()
}

func TestGeneratePDF_English(t *testing.T) {
	logoBytes := createDummyPNG()

	req := models.PDFRequest{
		Language:  language.LangEnglish,
		Direction: language.DirectionLTR,
		Logo: models.ImageSource{
			Bytes: logoBytes,
		},
		Data: models.KFSData{
			LoadAccountNumber:    "123456789",
			SanctionedLoanAmount: "5,00,000",
			LoanTerm:             "1 Year",
			BenchmarkRate:        "7.25",
			Spread:               "1.50",
			FinatRate:            "8.75",
			APR:                  "8.75",
			RateOfInterest:       "8.75",
		},
	}

	tempFile := "test_english_output.pdf"
	defer os.Remove(tempFile)

	err := GeneratePDF(req, tempFile)
	if err != nil {
		t.Fatalf("Failed to generate PDF: %v", err)
	}

	fi, err := os.Stat(tempFile)
	if err != nil {
		t.Fatalf("Output file does not exist: %v", err)
	}

	if fi.Size() == 0 {
		t.Error("Generated PDF file is empty")
	}
}
