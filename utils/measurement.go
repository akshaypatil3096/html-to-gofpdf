package utils

import "html-to-gofpdf/language"

const (
	// PtToMmFactor converts font points to millimeters.
	PtToMmFactor = 25.4 / 72.0

	// PxToMmFactor converts CSS pixels to millimeters.
	PxToMmFactor = 25.4 / 96.0
)

// PtToMm converts points to millimeters.
func PtToMm(pt float64) float64 {
	return pt * PtToMmFactor
}

// PxToMm converts pixels to millimeters.
func PxToMm(px float64) float64 {
	return px * PxToMmFactor
}

// GetXPhysical translates a logical X position (from the start margin)
// into a physical X position on the page based on text direction.
// - xLogical: logical coordinate measured from the active margin.
// - elementWidth: width of the element.
// - pageWidth: total width of the page.
// - marginLeft: left margin.
// - marginRight: right margin.
// - dir: Text direction (LTR or RTL).
func GetXPhysical(xLogical, elementWidth, pageWidth, marginLeft, marginRight float64, dir language.TextDirection) float64 {
	if dir == language.DirectionRTL {
		// In RTL, 0 starts at (pageWidth - marginRight) and grows to the left.
		return pageWidth - marginRight - xLogical - elementWidth
	}
	return marginLeft + xLogical
}

// FlipAlignment flips the text alignment string ("L" or "R") if direction is RTL.
func FlipAlignment(align string, dir language.TextDirection) string {
	if dir == language.DirectionRTL {
		switch align {
		case "L":
			return "R"
		case "R":
			return "L"
		}
	}
	return align
}
