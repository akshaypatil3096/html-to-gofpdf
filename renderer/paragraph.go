package renderer

import (
	"html-to-gofpdf/utils"
)

// Paragraph represents a block of multi-line text with custom spacing and alignments.
type Paragraph struct {
	Text         string
	FontSize     float64
	FontStyle    string
	TextColor    [3]int
	HasTextColor bool
	Align        string
	LineSpacing  float64
}

// Measure calculates the vertical height needed to display the paragraph.
func (p *Paragraph) Measure(ctx *Context, maxWidth float64) (Size, error) {
	if p.FontSize > 0 {
		if err := ctx.SetFont(p.FontStyle, p.FontSize); err != nil {
			return Size{}, err
		}
	}

	lines := utils.SplitText(ctx.PDF, p.Text, maxWidth, ctx.IsCoreFont())
	lineHeight := p.FontSize
	if lineHeight == 0 {
		lineHeight = ctx.CurrentFontSize
	}
	if p.LineSpacing == 0 {
		p.LineSpacing = 1.2
	}

	h := float64(len(lines)) * utils.PtToMm(lineHeight) * p.LineSpacing
	return Size{Width: maxWidth, Height: h}, nil
}

// Draw renders the paragraph's lines line by line on the page.
func (p *Paragraph) Draw(ctx *Context, x, y, width, height float64) error {
	if p.FontSize > 0 {
		if err := ctx.SetFont(p.FontStyle, p.FontSize); err != nil {
			return err
		}
	}
	if p.HasTextColor {
		ctx.PDF.SetTextColor(p.TextColor[0], p.TextColor[1], p.TextColor[2])
	} else {
		ctx.PDF.SetTextColor(0, 0, 0)
	}

	lines := utils.SplitText(ctx.PDF, p.Text, width, ctx.IsCoreFont())
	lineHeightVal := p.FontSize
	if lineHeightVal == 0 {
		lineHeightVal = ctx.CurrentFontSize
	}
	if p.LineSpacing == 0 {
		p.LineSpacing = 1.2
	}
	singleLineHeight := utils.PtToMm(lineHeightVal) * p.LineSpacing

	align := p.Align
	if align == "" {
		align = "L"
	}

	for i, line := range lines {
		ly := y + float64(i)*singleLineHeight
		ctx.DrawLogicalText(x, ly, width, singleLineHeight, line, align, "", false)
	}
	return nil
}
