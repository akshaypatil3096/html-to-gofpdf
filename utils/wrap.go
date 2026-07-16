package utils

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// SplitText wraps text by splitting it at newlines first, and then using
// the gofpdf SplitText helper to wrap lines that exceed the target width.
// isCoreFont ensures Unicode characters are replaced appropriately for core PDF fonts.
func SplitText(pdf *gofpdf.Fpdf, text string, width float64, isCoreFont bool) []string {
	if text == "" {
		return []string{""}
	}
	// Normalize line endings and placeholders
	text = strings.ReplaceAll(text, "\r\n", "\n")
	rawLines := strings.Split(text, "\n")

	var wrappedLines []string
	for _, line := range rawLines {
		cleanedLine := strings.ReplaceAll(line, "<br>", "\n")
		cleanedLine = strings.ReplaceAll(cleanedLine, "<br/>", "\n")
		cleanedLine = strings.ReplaceAll(cleanedLine, "<br />", "\n")

		subRawLines := strings.Split(cleanedLine, "\n")
		for _, subRawLine := range subRawLines {
			sanitized := SanitizeText(subRawLine, isCoreFont)
			if sanitized == "" {
				wrappedLines = append(wrappedLines, "")
				continue
			}
			// Use gofpdf's native UTF-8 SplitText
			subLines := pdf.SplitText(sanitized, width)
			wrappedLines = append(wrappedLines, subLines...)
		}
	}

	return wrappedLines
}
