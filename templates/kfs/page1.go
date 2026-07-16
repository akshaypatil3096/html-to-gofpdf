package kfs

import (
	"html-to-gofpdf/models"
	"html-to-gofpdf/renderer"
	"html-to-gofpdf/utils"
)

// RenderPage1 draws the layout of KFS Part 1 onto the active page.
func RenderPage1(ctx *renderer.Context, req models.PDFRequest) error {
	// 1. Render Header (Logo and Titles)
	if err := renderPage1Header(ctx, req); err != nil {
		return err
	}

	// 2. Set up the table column widths as percentages matching the HTML <colgroup>
	colWidths := []float64{4.93, 7.69, 10.23, 10.23, 12.03, 12.42, 12.03, 15.02, 15.42}

	// 3. Create Rows
	rows := [][]renderer.TableCell{
		// Row 1: Loan Account No & Type of Loan
		{
			cCell("1", true, "C", 1, 1),
			cCell("Loan Account No.", true, "L", 2, 1),
			cCell(req.Data.LoadAccountNumber, false, "L", 2, 1),
			cCell("Type of Loan", true, "L", 1, 1),
			cCell("Overdraft Against Fixed Deposit", false, "L", 3, 1),
		},
		// Row 2: Sanctioned Loan amount (in Rupees)
		{
			cCell("2", true, "C", 1, 1),
			cCell("Sanctioned Loan amount (in Rupees)", true, "L", 4, 1),
			cCell(req.Data.SanctionedLoanAmount, false, "L", 4, 1),
		},
		// Row 3: Disbursal schedule
		{
			cCell("3", true, "C", 1, 1),
			cCell("Disbursal schedule\n(i) Disbursement in stages or 100% upfront.\n(ii) If it is stage wise, mention the clause of loan agreement having relevant details", false, "L", 4, 1),
			cCell("100% upfront", false, "C", 4, 1),
		},
		// Row 4: Loan term
		{
			cCell("4", true, "C", 1, 1),
			cCell("Loan term (year/months/days)", true, "L", 4, 1, 1.0, 1.0, 2.1, 2.1), // padding-top/bottom: 8px (2.11mm)
			cCell(req.Data.LoanTerm, false, "L", 4, 1),
		},
		// Row 5: Instalment details header
		{
			cCell("5", true, "C", 1, 1),
			cCell("Instalment details", true, "L", 8, 1),
		},
		// Row 5 Subrow 1
		{
			skipCell(),
			cCell("Type of instalments", false, "L", 2, 1),
			cCell("Number of EPIs", false, "L", 2, 1),
			cCell("EPI (Rs)", false, "L", 2, 1),
			cCell("Commencement of repayment, post sanction", false, "L", 2, 1),
		},
		// Row 5 Subrow 2
		{
			skipCell(),
			cCell("Not applicable", false, "C", 2, 1),
			cCell("Not applicable", false, "C", 2, 1),
			cCell("Not applicable", false, "C", 2, 1),
			cCell("Not applicable", false, "C", 2, 1),
		},
		// Row 6: Interest rate (%) and type
		{
			cCell("6", true, "C", 1, 1),
			cCell("Interest rate (%) and type (fixed or floating or hybrid)", true, "L", 4, 1, 1.0, 1.0, 2.1, 2.1),
			cCell("Floating", false, "L", 4, 1),
		},
		// Row 7: Floating interest details
		{
			cCell("7", true, "C", 1, 4),
			cCell("Additional Information in case of Floating rate of interest", true, "L", 8, 1),
		},
		// Row 7 Subheader 1
		{
			cCell("Reference\nBenchmark", false, "C", 1, 2),
			cCell("Benchmark\nrate\n(%) (B)", false, "C", 1, 2),
			cCell("Spread\n(s)\n(S)", false, "C", 1, 2),
			cCell("Final rate\n(%) R = (B)\n+ (S)", false, "C", 1, 2),
			cCell("Reset periodicity\n2 (Months)", false, "C", 2, 1),
			cCell("Impact of change in the reference benchmark (for 25 bps change in 'R', change in:3)", false, "C", 2, 1),
		},
		// Row 7 Subheader 2
		{
			cCell("B", false, "C", 1, 1),
			cCell("S", false, "C", 1, 1),
			cCell("EPI (Rs)", false, "C", 1, 1),
			cCell("No. of\nEPIs", false, "C", 1, 1),
		},
		// Row 7 Data
		{
			cCell("Fixed\nDeposit\n(FD)", false, "L", 1, 1),
			cCell(req.Data.BenchmarkRate, false, "C", 1, 1),
			cCell(req.Data.Spread, false, "C", 1, 1),
			cCell(req.Data.FinatRate, false, "C", 1, 1),
			cCell("On\nrenewal\nof FD", false, "C", 1, 1),
			cCell("Not\napplicable", false, "C", 1, 1),
			cCell("Not\napplicable", false, "C", 1, 1),
			cCell("Not\napplicable", false, "C", 1, 1),
		},
		// Row 8: Fee/ Charges
		{
			cCell("8", true, "C", 1, 7),
			cCell("Fee/ Charges 4", true, "L", 8, 1),
		},
		// Row 8 Subheader 1
		{
			cCell("", false, "C", 1, 2),
			cCell("", false, "C", 1, 2),
			cCell("Payable to the RE (A)", false, "C", 3, 1),
			cCell("Payable to a third party through RE (B)", false, "C", 3, 1),
		},
		// Row 8 Subheader 2
		{
			cCell("One-time/\nRecurring", false, "C", 1, 1),
			cCell("Amount (in Rs) or Percentage (%) as applicable5", false, "C", 2, 1),
			cCell("One-time/\nRecurring", false, "C", 1, 1),
			cCell("Amount (in Rs) or Percentage (%) as applicable5", false, "C", 2, 1),
		},
		// Row 8 Data (i) Processing fees
		{
			cCell("(i)", false, "C", 1, 1),
			cCell("Processing fees", false, "L", 1, 1),
			cCell("Nil", false, "C", 1, 1),
			cCell("Nil", false, "C", 2, 1),
			cCell("Nil", false, "C", 1, 1),
			cCell("Nil", false, "C", 2, 1),
		},
		// Row 8 Data (ii) Insurance charges
		{
			cCell("(ii)", false, "C", 1, 1),
			cCell("Insurance charges", false, "L", 1, 1),
			cCell("Nil", false, "C", 1, 1),
			cCell("Nil", false, "C", 2, 1),
			cCell("Nil", false, "C", 1, 1),
			cCell("Nil", false, "C", 2, 1),
		},
		// Row 8 Data (iii) Valuation fees
		{
			cCell("(iii)", false, "C", 1, 1),
			cCell("Valuation fees", false, "L", 1, 1),
			cCell("Nil", false, "C", 1, 1),
			cCell("Nil", false, "C", 2, 1),
			cCell("Nil", false, "C", 1, 1),
			cCell("Nil", false, "C", 2, 1),
		},
		// Row 8 Data (iv) Any other
		{
			cCell("(iv)", false, "C", 1, 1),
			cCell("Any other (please specify)", false, "L", 1, 1),
			cCell("Nil", false, "C", 1, 1),
			cCell("Nil", false, "C", 2, 1),
			cCell("Nil", false, "C", 1, 1),
			cCell("Nil", false, "C", 2, 1),
		},
		// Row 9: Annual Percentage Rate (APR)
		{
			cCell("9", true, "C", 1, 1),
			cCell("Annual Percentage Rate (APR) (%)6", true, "L", 4, 1),
			cCell(req.Data.APR, false, "L", 4, 1),
		},
		// Row 10: Details of Contingent Charges
		{
			cCell("10", true, "C", 1, 6),
			cCell("Details of Contingent Charges (in Rs or %, as applicable)", true, "L", 8, 1),
		},
		// Row 10 Data (i)
		{
			cCell("(i)", false, "C", 1, 1),
			cCell("Penal charges, if any, in case of delayed payment", false, "L", 5, 1),
			cCell("18% per annum on the\noverdrawn amount.", false, "C", 2, 1),
		},
		// Row 10 Data (ii)
		{
			cCell("(ii)", false, "C", 1, 1),
			cCell("Other penal charges, if any", false, "L", 5, 1),
			cCell("Nil", false, "C", 2, 1),
		},
		// Row 10 Data (iii)
		{
			cCell("(iii)", false, "C", 1, 1),
			cCell("Foreclosure charges, if applicable", false, "L", 5, 1),
			cCell("Nil", false, "C", 2, 1),
		},
		// Row 10 Data (iv)
		{
			cCell("(iv)", false, "C", 1, 1),
			cCell("Charges for switching of loans from floating to fixed rate and vice versa", false, "L", 5, 1),
			cCell("Nil", false, "C", 2, 1),
		},
		// Row 10 Data (v)
		{
			cCell("(v)", false, "C", 1, 1),
			cCell("Any other charges (please specify)", false, "L", 5, 1),
			cCell("Nil", false, "C", 2, 1),
		},
	}

	table := &renderer.Table{
		ColWidths: colWidths,
		Rows:      rows,
	}

	// 4. Render Table
	// Y starts at 0 since Page 1 top margin accounts for padding (34.84mm)
	sz, err := table.Measure(ctx, ctx.GetDrawableWidth())
	if err != nil {
		return err
	}
	return table.Draw(ctx, 0, 0, sz.Width, sz.Height)
}

func renderPage1Header(ctx *renderer.Context, req models.PDFRequest) error {
	// Logo
	if req.Logo.Path != "" || len(req.Logo.Bytes) > 0 {
		img := &renderer.Image{
			Path:  req.Logo.Path,
			Bytes: req.Logo.Bytes,
		}
		// X = 27.69 - ML(14.17) = 13.52
		// Y = 13.41 - MT(34.84) = -21.43
		if err := img.Draw(ctx, 13.52, -21.43, 31.77, 4.15); err != nil {
			return err
		}
	}

	// Title: Key Facts Statement
	if err := ctx.SetFont("B", 15); err != nil {
		return err
	}
	// Y = 22.4 - MT(34.84) = -12.44
	ctx.DrawLogicalText(0, -12.44, ctx.GetDrawableWidth(), utils.PtToMm(15), "Key Facts Statement", "C", "", false)

	// Subtitle: Part 1
	if err := ctx.SetFont("", 10); err != nil {
		return err
	}
	// Y = -12.44 + title_height(15pt) + 0.87 = -12.44 + 5.29 + 0.87 = -6.28
	ctx.DrawLogicalText(0, -6.28, ctx.GetDrawableWidth(), utils.PtToMm(10), "Part 1 (Interest rate and fees/charges)", "C", "", false)

	return nil
}

func defaultBorder() renderer.BorderSpec {
	return renderer.BorderSpec{
		Left:   true,
		Right:  true,
		Top:    true,
		Bottom: true,
		Color:  [3]int{0, 0, 0},
		Width:  0.2,
	}
}

func skipCell() renderer.TableCell {
	return renderer.TableCell{
		Component: &renderer.Cell{
			Border: renderer.BorderSpec{
				Left:   true,
				Right:  true,
				Top:    false, // Hidden top border
				Bottom: true,
				Color:  [3]int{0, 0, 0},
				Width:  0.2,
			},
		},
		ColSpan: 1,
		RowSpan: 1,
	}
}

func cCell(text string, bold bool, align string, colspan, rowspan int, padding ...float64) renderer.TableCell {
	style := ""
	if bold {
		style = "B"
	}
	pLeft, pRight, pTop, pBottom := 1.0, 1.0, 0.7, 0.7
	if len(padding) >= 4 {
		pLeft = padding[0]
		pRight = padding[1]
		pTop = padding[2]
		pBottom = padding[3]
	}
	return renderer.TableCell{
		Component: &renderer.Cell{
			Text:          text,
			FontSize:      10,
			FontStyle:     style,
			Align:         align,
			Border:        defaultBorder(),
			PaddingLeft:   pLeft,
			PaddingRight:  pRight,
			PaddingTop:    pTop,
			PaddingBottom: pBottom,
		},
		ColSpan: colspan,
		RowSpan: rowspan,
	}
}
