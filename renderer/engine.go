package renderer

import (
	"html-to-gofpdf/fonts"
	"html-to-gofpdf/language"
	"html-to-gofpdf/utils"

	"github.com/jung-kurt/gofpdf"
)

// Context wraps gofpdf, holding layout rules, active styling, and coordinate bounds.
type Context struct {
	PDF          *gofpdf.Fpdf
	LangConfig   language.LanguageConfig
	Fonts        *fonts.FontManager
	PageWidth    float64
	PageHeight   float64
	MarginLeft   float64
	MarginRight  float64
	MarginTop    float64
	MarginBottom float64

	CurrentFontFamily string
	CurrentFontSize   float64
	CurrentFontStyle  string
}

// NewContext creates a Context initialized with standard gofpdf attributes.
func NewContext(pdf *gofpdf.Fpdf, lang language.Language, fm *fonts.FontManager) *Context {
	langConfig := language.GetDefaultConfig(lang)
	if customCfg := fm.GetFontConfig(lang); customCfg.Family != "" {
		langConfig.Font = customCfg
	}

	w, h := pdf.GetPageSize()
	ml, mt, mr, mb := pdf.GetMargins()
	if langConfig.Direction == language.DirectionRTL {
		ml, mr = mr, ml
	}

	return &Context{
		PDF:          pdf,
		LangConfig:   langConfig,
		Fonts:        fm,
		PageWidth:    w,
		PageHeight:   h,
		MarginLeft:   ml,
		MarginRight:  mr,
		MarginTop:    mt,
		MarginBottom: mb,
	}
}

// GetDrawableWidth returns the printable horizontal span inside the page margins.
func (c *Context) GetDrawableWidth() float64 {
	return c.PageWidth - c.MarginLeft - c.MarginRight
}

// SetFont configures the font size, family, and styling for the current language config.
func (c *Context) SetFont(style string, size float64) error {
	family, err := c.Fonts.LoadFontForStyle(c.PDF, c.LangConfig.Name, style)
	if err != nil {
		return err
	}
	c.CurrentFontFamily = family
	c.CurrentFontStyle = style
	c.CurrentFontSize = size
	c.PDF.SetFont(family, style, size)
	return nil
}

// DrawLogicalText draws a text string within a logical coordinate block.
// Logical (x, y) values are converted to RTL bounds automatically if direction is RTL.
func (c *Context) DrawLogicalText(x, y, w, h float64, text string, align string, border string, fill bool) {
	physX := utils.GetXPhysical(x, w, c.PageWidth, c.MarginLeft, c.MarginRight, c.LangConfig.Direction)
	physY := c.MarginTop + y
	physAlign := utils.FlipAlignment(align, c.LangConfig.Direction)

	c.PDF.SetXY(physX, physY)
	c.PDF.CellFormat(w, h, text, border, 0, physAlign, fill, 0, "")
}

// DrawLogicalRect draws a border rectangle utilizing direction-aware coordinates.
func (c *Context) DrawLogicalRect(x, y, w, h float64, style string) {
	physX := utils.GetXPhysical(x, w, c.PageWidth, c.MarginLeft, c.MarginRight, c.LangConfig.Direction)
	physY := c.MarginTop + y
	c.PDF.Rect(physX, physY, w, h, style)
}

// DrawLogicalLine draws a line between logical endpoints.
func (c *Context) DrawLogicalLine(x1, y1, x2, y2 float64) {
	physX1 := utils.GetXPhysical(x1, 0, c.PageWidth, c.MarginLeft, c.MarginRight, c.LangConfig.Direction)
	physX2 := utils.GetXPhysical(x2, 0, c.PageWidth, c.MarginLeft, c.MarginRight, c.LangConfig.Direction)
	physY1 := c.MarginTop + y1
	physY2 := c.MarginTop + y2
	c.PDF.Line(physX1, physY1, physX2, physY2)
}

// IsCoreFont returns true if the current active font is a standard core font.
func (c *Context) IsCoreFont() bool {
	return fonts.IsCoreFont(c.CurrentFontFamily)
}

