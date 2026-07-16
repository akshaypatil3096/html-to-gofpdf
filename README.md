# Gofpdf PDF Generation Framework

A production-grade, highly extensible PDF generation framework in Go built on top of `github.com/jung-kurt/gofpdf`. It is designed to programmatically replicate design references (like A4 HTML templates) pixel-for-pixel, supporting dynamic data binding, multi-language Unicode typefaces, and Left-to-Right (LTR) or Right-to-Left (RTL) rendering directions.

## Key Features

1. **Pixel-Perfect HTML Replication**: Programmatically recreates margins, paddings, lines, borders, and column layouts from A4 page templates.
2. **Unified LTR & RTL Engine**: Supports horizontal coordinate mirroring, alignment flips, and column order reversal using logical LTR-defined coordinates.
3. **Advanced Table Layout**: Supports nested components, column widths scaled as percentage widths, multi-colspan, multi-rowspan, dynamic row heights, and merged borders.
4. **Typographic Crash Prevention**: Automatically cleans up HTML entities, converts curly punctuation marks to ASCII, and sanitizes character indexes above 255 for standard core PDF fonts to prevent array index-out-of-bounds crashes in `gofpdf`.
5. **Flexible Asset Registry**: Supports image rendering from local file paths or raw byte buffers with MD5-based register deduplication.

---

## Directory Structure

```
html-to-gofpdf/
├── main.go               # Command line tool to generate test PDFs
├── go.mod                # Module definitions
├── go.sum                # Package checksums
├── eng_index.html        # HTML layout design specification
├── pdf/
│   ├── generator.go      # Orchestrates A4 page margin configs and draws templates
│   └── generator_test.go # Automated unit test suite
├── models/
│   ├── data.go           # KFSData struct matching template placeholders
│   └── request.go        # Public PDFRequest input model
├── language/
│   ├── config.go         # Font configurations, locales, and LTR/RTL structures
│   ├── ltr.go            # LTR check helper
│   └── rtl.go            # RTL check helper
├── fonts/
│   └── manager.go        # Dynamic Unicode TTF loader and core font overrides
├── renderer/
│   ├── engine.go         # Logical Context containing drawing utilities
│   ├── layout.go         # Primitive Component interfaces
│   ├── cell.go           # Sized rectangular cell layout blocks
│   ├── paragraph.go      # Wrapped multiline text blocks
│   ├── image.go          # Image registration and drawing
│   └── table.go          # Matrix rowspan/colspan table renderer
├── templates/
│   └── kfs/
│       ├── page1.go      # Page 1 (Part 1 Statement Table)
│       ├── page2.go      # Page 2 (Qualitative Stacked Tables)
│       └── page3.go      # Page 3 (Annex B Computation Table)
└── utils/
    ├── measurement.go    # Coordinate scaling (px/pt to mm) and LTR/RTL X-transformations
    ├── text.go           # Punctuation cleaning and core font bounds safety
    └── wrap.go           # Text wrapping under width bounds
```

---

## Technical Highlights

### 1. Left-to-Right & Right-to-Left Mathematics
Symmetric mapping allows template layouts to be written once in a standard LTR logical format. During rendering, horizontal bounds are converted physically:
- **LTR**: $X_{physical} = Margin_{left} + X_{logical}$
- **RTL**: $X_{physical} = PageWidth - Margin_{right} - X_{logical} - ElementWidth$

Additionally, the Left and Right margins are automatically swapped internally by the context in RTL mode, ensuring that asymmetric print alignments (such as A4 bindings) align correctly.

### 2. Multi-Rowspan and Colspan Grid Engine
The table layout engine:
1. Translates tabular cell blocks into a unified 2D grid matrix of resolved positions.
2. Iterates over rows and measures the height required by cells with a `rowspan` of 1.
3. Iterates over rowspanned cells and distributes their excess height proportionally across all the rows they span.
4. Scale column widths based on target percentage widths so the table matches the width of the printable margins.

### 3. Core Font Safety Sanitization
Standard core PDF fonts (like Helvetica, Times, Courier) only support 256-index character maps. If text containing Unicode punctuation (like en-dashes `–` or curly single quotes `‘`) is passed to `SplitText`, `gofpdf` panics with `index out of range`.
The framework detects if a core font is active and automatically cleans the text runs:
- Normalizes HTML symbols like `&amp;` and `&rsquo;`.
- Translates curly punctuation and en-dashes into safe ASCII equivalents (`-`, `'`, `"`).
- Safely sanitizes any remaining runes $> 255$ into `?` to prevent panics, while allowing full Unicode rendering for registered custom TTF fonts.

---

## API Usage Example

```go
package main

import (
	"html-to-gofpdf/language"
	"html-to-gofpdf/models"
	"html-to-gofpdf/pdf"
)

func main() {
	req := models.PDFRequest{
		Language:  language.LangEnglish,
		Direction: language.DirectionLTR,
		Logo: models.ImageSource{
			Path: "path/to/logo.png",
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

	err := pdf.GeneratePDF(req, "output.pdf")
	if err != nil {
		panic(err)
	}
}
```

---

## Building and Verification

### Run Automated Unit Tests
To run the automated tests offline:
```bash
GOPROXY=off GOSUMDB=off go test -v ./...
```

### Run the Command Line Utility
To run the manual PDF compiler generating output files for multiple languages:
```bash
GOPROXY=off GOSUMDB=off go run main.go
```

The CLI tool compiles three test cases in the workspace:
1. `output_english.pdf`
2. `output_urdu.pdf`
3. `output_hindi.pdf`