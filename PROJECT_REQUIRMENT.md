# PROJECT REQUIREMENTS

## Project

Create a production-grade PDF generation framework in Go using **gofpdf**.

The framework will generate PDF documents that are visually identical to the supplied HTML templates. The supplied HTML files are **design specifications only** and must NOT be rendered at runtime.

Every layout element must be recreated programmatically using gofpdf.

The framework must be reusable, maintainable, extensible, and capable of supporting multiple document templates and multiple languages.

---

# Primary Goal

The generated PDF must look visually identical to the supplied HTML.

No layout simplification is allowed.

This includes:

- page size
- margins
- spacing
- padding
- borders
- merged cells
- row heights
- column widths
- fonts
- font sizes
- line spacing
- page breaks
- images
- wrapping
- alignments

The HTML is the design reference.

The generated PDF should be visually indistinguishable from the HTML.

---

# Important

DO NOT parse HTML.

DO NOT use HTMLBasicNew().

DO NOT convert HTML into PDF.

Instead,

Analyze the HTML once and recreate the document using gofpdf drawing APIs.

---

# Future Requirement

The system must support multiple document templates.

For example:

- English KFS
- Urdu KFS
- Hindi KFS
- Arabic KFS
- Any future template

without rewriting the rendering engine.

---

# Multi Language Support

The framework must fully support multilingual PDF generation.

Examples:

- English
- Urdu
- Arabic
- Hindi
- Marathi
- Unicode languages

---

# Left-To-Right and Right-To-Left Support

This is one of the most important requirements.

The rendering engine must support both

- Left-To-Right (LTR)
- Right-To-Left (RTL)

layouts.

Examples

English

```
Loan Account Number
123456789
```

Urdu

```
اکاؤنٹ نمبر
123456789
```

The rendering engine should automatically handle

- text direction
- alignment
- wrapped text
- positioning
- cell rendering

based on document configuration.

The rendering logic should NOT be duplicated for RTL documents.

Instead,

there should be a configurable rendering mode.

Example

```
DocumentDirectionLTR

DocumentDirectionRTL
```

---

# Font Management

The framework must support

- TrueType fonts
- UTF-8 fonts
- Unicode fonts

The font system must allow different fonts per language.

Example

English

Arial

Urdu

Noto Nastaliq Urdu

Arabic

Amiri

Hindi

Noto Sans Devanagari

The rendering engine must automatically select the configured font.

---

# Image Support

Support

PNG

JPEG

transparent PNG

Maintain aspect ratio.

Allow images to be supplied using

- file path
- byte array
- io.Reader

The renderer should place images exactly at the configured coordinates.

---

# Dynamic Data

Every dynamic field in the HTML should become Go fields.

Example

Loan Account Number

APR

Loan Amount

Loan Term

Benchmark Rate

Spread

Interest Rate

etc.

No hardcoded business values.

---

# Rendering Architecture

Separate

Layout

Rendering

Data

Fonts

Images

Pages

Configuration

Each page should have its own renderer.

Example

Page1Renderer

Page2Renderer

Page3Renderer

---

# Layout Engine

Build reusable primitives.

Do NOT hardcode thousands of Cell() calls.

Reusable components should include

Rectangle

Border

Cell

Merged Cell

Wrapped Cell

Paragraph

Image

Horizontal Line

Vertical Line

Table

Header

Footer

Spacer

Container

Absolute Position

Relative Position

These primitives should be reusable across all templates.

---

# Table Engine

The supplied HTML contains many complex tables.

The framework must support

merged cells

rowspan

colspan

nested tables

dynamic row height

dynamic wrapped text

automatic border alignment

automatic height calculation

table splitting across pages (future support)

---

# Text Rendering

Support

Left aligned

Center aligned

Right aligned

RTL aligned

Automatic wrapping

Automatic height calculation

Unicode rendering

Multi-line cells

Vertical alignment

---

# Page Management

Support

multiple pages

custom headers

custom footers

page numbers

automatic page breaks

manual page breaks

future templates

---

# Configuration

Everything should be configurable.

Example

Page Size

Margins

Fonts

Direction

Language

Image

Colors

Border Width

Cell Padding

---

# Public API

The final API should be simple.

Example

```go
type PDFRequest struct {
    Language Language

    Direction TextDirection

    Logo ImageSource

    Data KFSData
}

func GeneratePDF(
    req PDFRequest,
    outputFile string,
) error
```

---

# Language Configuration

Avoid hardcoding language-specific logic.

Instead

```go
type LanguageConfig struct {
    Name string

    Direction TextDirection

    Font FontConfig

    Locale string
}
```

The rendering engine should work entirely based on this configuration.

---

# Project Structure

```
pdf/

    generator.go

    renderer/

        engine.go

        page.go

        table.go

        cell.go

        border.go

        image.go

        paragraph.go

        layout.go

    fonts/

        manager.go

    language/

        config.go

        rtl.go

        ltr.go

    templates/

        kfs/

            page1.go

            page2.go

            page3.go

    models/

        request.go

        data.go

    utils/

        measurement.go

        text.go

        wrap.go
```

---

# Extensibility

Adding a new document should require only

1. New template layout

NOT

modifying the rendering engine.

---

# Performance

The framework should

avoid duplicated calculations

reuse measurements

cache loaded fonts

cache images

minimize allocations

avoid unnecessary object creation

---

# Error Handling

Return descriptive errors for

missing fonts

missing images

invalid coordinates

text overflow

page overflow

layout inconsistencies

---

# Testing

The code should be easily unit testable.

The rendering engine should be separated from business logic.

---

# Coding Standards

Follow

Effective Go

SOLID

DRY

KISS

Clean Architecture

Idiomatic Go

Small reusable functions

Proper documentation

Meaningful package structure

No duplicated rendering logic.

---

# Final Deliverable

Generate a production-ready PDF rendering framework capable of producing pixel-perfect PDFs that match the supplied HTML layouts exactly.

The framework must support:

- exact layout reproduction
- reusable rendering engine
- multiple document templates
- multilingual documents
- Unicode
- Left-To-Right documents
- Right-To-Left documents (Urdu, Arabic)
- configurable fonts
- configurable images
- dynamic data
- future extensibility

The generated code should be suitable for enterprise production use and should not require architectural changes when new languages or templates are introduced.