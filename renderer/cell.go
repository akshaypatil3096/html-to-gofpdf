package renderer

import (
	"html-to-gofpdf/utils"
)

// BorderSpec configures borders for individual sides of a cell.
type BorderSpec struct {
	Left   bool
	Right  bool
	Top    bool
	Bottom bool
	Color  [3]int  // RGB color
	Width  float64 // Width in mm
}

// Cell is a standard rectangular area with text alignment, padding, border, and background attributes.
type Cell struct {
	Text            string
	Child           Component
	Width           float64
	Height          float64
	PaddingLeft     float64
	PaddingRight    float64
	PaddingTop      float64
	PaddingBottom   float64
	Border          BorderSpec
	BackgroundFill  bool
	BackgroundColor [3]int
	TextColor       [3]int
	HasTextColor    bool
	Align           string // "L", "C", "R"
	FontStyle       string // "", "B", "I", "BI"
	FontSize        float64
	LineSpacing     float64
}

// Measure calculates cell width and height considering children or lines of text.
func (c *Cell) Measure(ctx *Context, maxWidth float64) (Size, error) {
	if c.FontSize > 0 {
		if err := ctx.SetFont(c.FontStyle, c.FontSize); err != nil {
			return Size{}, err
		}
	}

	contentWidth := maxWidth - c.PaddingLeft - c.PaddingRight
	if c.Width > 0 && c.Width < maxWidth {
		contentWidth = c.Width - c.PaddingLeft - c.PaddingRight
	}

	var h float64
	if c.Child != nil {
		sz, err := c.Child.Measure(ctx, contentWidth)
		if err != nil {
			return Size{}, err
		}
		h = sz.Height
	} else {
		lines := utils.SplitText(ctx.PDF, c.Text, contentWidth, ctx.IsCoreFont())
		lineHeight := c.FontSize
		if lineHeight == 0 {
			lineHeight = ctx.CurrentFontSize
		}
		if c.LineSpacing == 0 {
			c.LineSpacing = 1.2
		}
		h = float64(len(lines)) * utils.PtToMm(lineHeight) * c.LineSpacing
	}

	totalHeight := h + c.PaddingTop + c.PaddingBottom
	if c.Height > totalHeight {
		totalHeight = c.Height
	}

	w := c.Width
	if w <= 0 {
		w = maxWidth
	}

	return Size{Width: w, Height: totalHeight}, nil
}

// Draw renders the cell backgrounds, lines of text, custom borders, and child components.
func (c *Cell) Draw(ctx *Context, x, y, width, height float64) error {
	if c.FontSize > 0 {
		if err := ctx.SetFont(c.FontStyle, c.FontSize); err != nil {
			return err
		}
	}

	// Apply colors
	if c.BackgroundFill {
		ctx.PDF.SetFillColor(c.BackgroundColor[0], c.BackgroundColor[1], c.BackgroundColor[2])
		ctx.DrawLogicalRect(x, y, width, height, "F")
	}
	if c.HasTextColor {
		ctx.PDF.SetTextColor(c.TextColor[0], c.TextColor[1], c.TextColor[2])
	} else {
		ctx.PDF.SetTextColor(0, 0, 0)
	}

	// Draw custom borders
	borderWidth := c.Border.Width
	if borderWidth == 0 {
		borderWidth = 0.2 // default border width in mm
	}
	ctx.PDF.SetLineWidth(borderWidth)
	ctx.PDF.SetDrawColor(c.Border.Color[0], c.Border.Color[1], c.Border.Color[2])

	if c.Border.Top {
		ctx.DrawLogicalLine(x, y, x+width, y)
	}
	if c.Border.Bottom {
		ctx.DrawLogicalLine(x, y+height, x+width, y+height)
	}
	if c.Border.Left {
		ctx.DrawLogicalLine(x, y, x, y+height)
	}
	if c.Border.Right {
		ctx.DrawLogicalLine(x+width, y, x+width, y+height)
	}

	// Calculate inner bounds
	cx := x + c.PaddingLeft
	cy := y + c.PaddingTop
	cw := width - c.PaddingLeft - c.PaddingRight
	ch := height - c.PaddingTop - c.PaddingBottom

	if c.Child != nil {
		return c.Child.Draw(ctx, cx, cy, cw, ch)
	}

	// Draw wrapping text lines
	lines := utils.SplitText(ctx.PDF, c.Text, cw, ctx.IsCoreFont())
	lineHeightVal := c.FontSize
	if lineHeightVal == 0 {
		lineHeightVal = ctx.CurrentFontSize
	}
	if c.LineSpacing == 0 {
		c.LineSpacing = 1.2
	}
	singleLineHeight := utils.PtToMm(lineHeightVal) * c.LineSpacing

	align := c.Align
	if align == "" {
		align = "L"
	}

	textHeight := float64(len(lines)) * singleLineHeight
	offsetY := 0.0
	if textHeight < ch {
		// Vertically align in the middle
		offsetY = (ch - textHeight) / 2.0
	}

	for i, line := range lines {
		ly := cy + offsetY + float64(i)*singleLineHeight
		ctx.DrawLogicalText(cx, ly, cw, singleLineHeight, line, align, "", false)
	}

	return nil
}
