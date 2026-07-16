package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"html-to-gofpdf/language"
	"html-to-gofpdf/models"
	"html-to-gofpdf/pdf"
)

func createDummyLogo() []byte {
	img := image.NewRGBA(image.Rect(0, 0, 150, 20))
	// Draw a blue background block simulating a company logo
	for x := 0; x < 150; x++ {
		for y := 0; y < 20; y++ {
			img.Set(x, y, color.RGBA{0, 102, 204, 255})
		}
	}
	var buf bytes.Buffer
	_ = png.Encode(&buf, img)
	return buf.Bytes()
}

func main() {
	fmt.Println("Starting PDF Generation Framework...")

	logoBytes := createDummyLogo()

	// 1. Generate English LTR PDF
	engReq := models.PDFRequest{
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

	engOutput := "output_english.pdf"
	fmt.Printf("Generating English PDF -> %s...\n", engOutput)
	if err := pdf.GeneratePDF(engReq, engOutput); err != nil {
		log.Fatalf("Error generating English PDF: %v", err)
	}
	fmt.Printf("Successfully generated English PDF: %s\n\n", engOutput)

	// 2. Generate Urdu RTL PDF (if system font is present)
	urConfig := language.GetDefaultConfig(language.LangUrdu)
	if _, err := os.Stat(urConfig.Font.FilePath); err == nil {
		urReq := models.PDFRequest{
			Language:  language.LangUrdu,
			Direction: language.DirectionRTL,
			Logo: models.ImageSource{
				Bytes: logoBytes,
			},
			Data: models.KFSData{
				LoadAccountNumber:    "987654321",
				SanctionedLoanAmount: "10,00,000",
				LoanTerm:             "2 Years",
				BenchmarkRate:        "6.50",
				Spread:               "2.00",
				FinatRate:            "8.50",
				APR:                  "8.50",
				RateOfInterest:       "8.50",
			},
		}

		urOutput := "output_urdu.pdf"
		fmt.Printf("Generating Urdu PDF -> %s...\n", urOutput)
		if err := pdf.GeneratePDF(urReq, urOutput); err != nil {
			fmt.Printf("Warning: Failed to generate Urdu PDF: %v\n", err)
		} else {
			fmt.Printf("Successfully generated Urdu PDF: %s\n\n", urOutput)
		}
	} else {
		fmt.Printf("Urdu font not found at default location (%s). Skipping Urdu PDF rendering.\n", urConfig.Font.FilePath)
		fmt.Println("To generate Urdu PDFs, configure a valid Noto Nastaliq Urdu TTF file path in PDFRequest.")
	}

	// 3. Generate Hindi LTR PDF (if system font is present)
	hiConfig := language.GetDefaultConfig(language.LangHindi)
	if _, err := os.Stat(hiConfig.Font.FilePath); err == nil {
		hiReq := models.PDFRequest{
			Language:  language.LangHindi,
			Direction: language.DirectionLTR,
			Logo: models.ImageSource{
				Bytes: logoBytes,
			},
			Data: models.KFSData{
				LoadAccountNumber:    "555555555",
				SanctionedLoanAmount: "3,50,000",
				LoanTerm:             "6 Months",
				BenchmarkRate:        "8.00",
				Spread:               "1.25",
				FinatRate:            "9.25",
				APR:                  "9.25",
				RateOfInterest:       "9.25",
			},
		}

		hiOutput := "output_hindi.pdf"
		fmt.Printf("Generating Hindi PDF -> %s...\n", hiOutput)
		if err := pdf.GeneratePDF(hiReq, hiOutput); err != nil {
			fmt.Printf("Warning: Failed to generate Hindi PDF: %v\n", err)
		} else {
			fmt.Printf("Successfully generated Hindi PDF: %s\n\n", hiOutput)
		}
	} else {
		fmt.Printf("Hindi font not found at default location (%s). Skipping Hindi PDF rendering.\n", hiConfig.Font.FilePath)
		fmt.Println("To generate Hindi PDFs, configure a valid Noto Sans Devanagari TTF file path in PDFRequest.")
	}

	fmt.Println("PDF generation completed successfully.")
}
