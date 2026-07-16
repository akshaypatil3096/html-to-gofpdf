package renderer

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html-to-gofpdf/utils"

	"github.com/jung-kurt/gofpdf"
)

// Image renders PNG or JPEG images from a file path or a byte buffer.
type Image struct {
	Path  string
	Bytes []byte
}

// Measure provides layout sizing. If height is set to 0, gofpdf will calculate it to preserve aspect ratio.
func (img *Image) Measure(ctx *Context, maxWidth float64) (Size, error) {
	return Size{Width: maxWidth, Height: 0}, nil
}

func detectImageType(data []byte) string {
	if len(data) >= 4 && bytes.Equal(data[:4], []byte{0x89, 0x50, 0x4E, 0x47}) {
		return "png"
	}
	if len(data) >= 2 && bytes.Equal(data[:2], []byte{0xFF, 0xD8}) {
		return "jpg"
	}
	return "png" // default fallback
}

// Draw registers and renders the image onto the page at logical boundaries.
func (img *Image) Draw(ctx *Context, x, y, width, height float64) error {
	var name string

	if len(img.Bytes) > 0 {
		// Generate a unique cache key based on content hash
		hasher := md5.New()
		hasher.Write(img.Bytes)
		name = fmt.Sprintf("bytes_%x", hasher.Sum(nil))

		// Register the reader with gofpdf if not registered yet
		if ctx.PDF.GetImageInfo(name) == nil {
			imgType := detectImageType(img.Bytes)
			ctx.PDF.RegisterImageOptionsReader(name, gofpdf.ImageOptions{ImageType: imgType}, bytes.NewReader(img.Bytes))
		}
	} else if img.Path != "" {
		name = img.Path
	} else {
		return nil
	}

	physX := utils.GetXPhysical(x, width, ctx.PageWidth, ctx.MarginLeft, ctx.MarginRight, ctx.LangConfig.Direction)
	physY := ctx.MarginTop + y

	ctx.PDF.ImageOptions(name, physX, physY, width, height, false, gofpdf.ImageOptions{}, 0, "")
	return nil
}
