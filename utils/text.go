package utils

import (
	"strings"
)

// SanitizeText cleans up common HTML entities and replaces Unicode punctuation.
// If isCoreFont is true, it replaces any rune > 255 with '?' to prevent index out of bounds in gofpdf's character width table.
func SanitizeText(text string, isCoreFont bool) string {
	// Normalize line breaks and spaces
	text = strings.ReplaceAll(text, "\r\n", "\n")

	// Replace HTML entities
	text = strings.ReplaceAll(text, "&amp;", "&")
	text = strings.ReplaceAll(text, "&lsquo;", "'")
	text = strings.ReplaceAll(text, "&rsquo;", "'")
	text = strings.ReplaceAll(text, "&ldquo;", "\"")
	text = strings.ReplaceAll(text, "&rdquo;", "\"")
	text = strings.ReplaceAll(text, "&ndash;", "-")
	text = strings.ReplaceAll(text, "&mdash;", "-")

	// Replace common Unicode punctuation
	text = strings.ReplaceAll(text, "\u2013", "-") // en dash
	text = strings.ReplaceAll(text, "\u2014", "-") // em dash
	text = strings.ReplaceAll(text, "\u2018", "'") // curly single quote left
	text = strings.ReplaceAll(text, "\u2019", "'") // curly single quote right
	text = strings.ReplaceAll(text, "\u201c", "\"") // curly double quote left
	text = strings.ReplaceAll(text, "\u201d", "\"") // curly double quote right

	if isCoreFont {
		runes := []rune(text)
		for i, r := range runes {
			if r > 255 {
				runes[i] = '?'
			}
		}
		text = string(runes)
	}

	return text
}
