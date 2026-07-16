package models

import "html-to-gofpdf/language"

// ImageSource specifies an image either by its local file path or raw bytes.
type ImageSource struct {
	Path  string
	Bytes []byte
}

// PDFRequest holds all inputs required to generate the PDF document.
type PDFRequest struct {
	Language  language.Language
	Direction language.TextDirection
	Logo      ImageSource
	Data      KFSData
}
