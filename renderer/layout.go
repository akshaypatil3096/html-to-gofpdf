package renderer

// Size represents width and height dimensions in millimeters.
type Size struct {
	Width  float64
	Height float64
}

// Component is the base interface for all layout elements (Cell, Table, Image, Paragraph, etc.)
// that are measured and rendered programmatically.
type Component interface {
	// Measure computes the dimensions of the component under a maximum width constraint.
	Measure(ctx *Context, maxWidth float64) (Size, error)

	// Draw renders the component on the PDF page at (x, y) with the specified width and height.
	// Coordinates (x, y) are logical coordinates from the active margin.
	Draw(ctx *Context, x, y, width, height float64) error
}
