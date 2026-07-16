package kfs

import (
	"html-to-gofpdf/models"
	"html-to-gofpdf/renderer"
	"html-to-gofpdf/utils"
)

// RenderPage3 draws the layout of Annex B on Page 3.
func RenderPage3(ctx *renderer.Context, req models.PDFRequest) error {
	// 1. Render Header (Logo, Annex label, and Title)
	if err := renderPage3Header(ctx, req); err != nil {
		return err
	}

	// 2. Set up the table column widths as percentages matching the HTML
	colWidths := []float64{8.0, 70.0, 22.0}

	pLeft15 := 15.0 * utils.PxToMmFactor // padding-left: 15px

	rows := [][]renderer.TableCell{
		// Row 1: Header
		{
			cCell("SR\nNo.", true, "C", 1, 1),
			cCell("Parameter", true, "L", 1, 1, pLeft15, 1.0, 0.7, 0.7),
			cCell("Details", true, "L", 1, 1, pLeft15, 1.0, 0.7, 0.7),
		},
		// Row 2: 1
		{
			cCell("1", true, "C", 1, 1),
			cCell("Sanctioned Loan amount (in Rupees) (Sl no. 2 of the KFS template – Part 1)", false, "L", 1, 1),
			cCell(req.Data.SanctionedLoanAmount, false, "C", 1, 1),
		},
		// Row 3: 2
		{
			cCell("2", true, "C", 1, 1),
			cCell("Loan Term (in years/ months/ days) (Sl No.4 of the KFS template – Part 1)", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell(req.Data.LoanTerm, false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
		// Row 4: a)
		{
			cCell("a)", false, "C", 1, 1),
			cCell("No. of instalments for payment of principal, in case of nonequated periodic loans", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell("Not applicable", false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
		// Row 5: b)
		{
			cCell("b)", false, "C", 1, 1),
			cCell("Type of EPI Amount of each EPI (in Rupees) and nos. of EPIs (e.g., no. of EMIs in case of monthly instalments) (Sl No. 5 of the KFS template – Part 1)", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell("Not applicable", false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
		// Row 6: c)
		{
			cCell("c)", false, "C", 1, 1),
			cCell("No. of instalments for payment of capitalised interest, if any", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell("Not applicable", false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
		// Row 7: d)
		{
			cCell("d)", false, "C", 1, 1),
			cCell("Commencement of repayments, post sanction (Sl No. 5 of the KFS template – Part 1)", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell("Not applicable", false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
		// Row 8: 3
		{
			cCell("3", true, "C", 1, 1),
			cCell("Interest rate type (fixed or floating or hybrid) (Sl No. 6 of the KFS template – Part 1)", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell("Floating", false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
		// Row 9: 4
		{
			cCell("4", true, "C", 1, 1),
			cCell("Rate of Interest (Sl No. 7 of the KFS template – Part 1)", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell(req.Data.RateOfInterest, false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
		// Row 10: 5
		{
			cCell("5", true, "C", 1, 1),
			cCell("Total Interest Amount to be charged during the entire tenor of the loan as per the rate prevailing on sanction date (in Rupees)", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell("Not applicable", false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
		// Row 11: 6
		{
			cCell("6", true, "C", 1, 1),
			cCell("Fee/ Charges payable8 (in Rupees)", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell("NIL", false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
		// Row 12: A
		{
			cCell("A", false, "C", 1, 1),
			cCell("Payable to the RE (Sl No.8A of the KFS template-Part 1)", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell("NIL", false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
		// Row 13: B
		{
			cCell("B", false, "C", 1, 1),
			cCell("Payable to third-party routed through RE (Sl No.8B of the KFS template – Part 1)", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell("NIL", false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
		// Row 14: 7
		{
			cCell("7", true, "C", 1, 1),
			cCell("Net disbursed amount (1-6) (in Rupees)", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell("NIL", false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
		// Row 15: 8
		{
			cCell("8", true, "C", 1, 1),
			cCell("Total amount to be paid by the borrower (sum of 1 and 5) (in Rupees)", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell("NIL", false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
		// Row 16: 9
		{
			cCell("9", true, "C", 1, 1),
			cCell("Annual Percentage rate- Effective annualized interest rate (in percentage)10 (Sl No.9 of the KFS template-Part 1)", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell(req.Data.APR, false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
		// Row 17: 10
		{
			cCell("10", true, "C", 1, 1),
			cCell("Schedule of disbursement as per terms and conditions", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell("Not applicable", false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
		// Row 18: 11
		{
			cCell("11", true, "C", 1, 1),
			cCell("Due date of payment of instalment and interest", false, "L", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
			cCell("Not applicable", false, "C", 1, 1, 1.0, 1.0, 8.0*utils.PxToMmFactor, 8.0*utils.PxToMmFactor),
		},
	}

	table := &renderer.Table{
		ColWidths: colWidths,
		Rows:      rows,
	}

	sz, err := table.Measure(ctx, ctx.GetDrawableWidth())
	if err != nil {
		return err
	}
	return table.Draw(ctx, 0, 0, sz.Width, sz.Height)
}

func renderPage3Header(ctx *renderer.Context, req models.PDFRequest) error {
	// Logo
	if req.Logo.Path != "" || len(req.Logo.Bytes) > 0 {
		img := &renderer.Image{
			Path:  req.Logo.Path,
			Bytes: req.Logo.Bytes,
		}
		// X = 27.69 - ML(14.17) = 13.52
		// Y = 13.41 - MT(53.49) = -40.08
		if err := img.Draw(ctx, 13.52, -40.08, 31.77, 4.15); err != nil {
			return err
		}
	}

	// Annex Header (PART 2 / Annex B) at absolute top: 25.12mm.
	// Since MT is 53.49mm:
	// "PART 2" is at Y = 25.12 - 53.49 = -28.37
	// line-height is 2.2 (approx. 7.76mm separation)
	// "Annex B" is at Y = -28.37 + 7.76 = -20.61
	// They align right to right: 31.82mm (which is 210.0 - 31.82 = 178.18mm physical)
	// Since MR is 8.82: logical X from start is 178.18 - ML(14.17) = 164.01
	// Or we can draw it as a block of width GetDrawableWidth() right aligned!
	// Yes! That automatically maps it right aligned inside margins!
	// Let's check:
	// LTR: right edge of drawable width is W - MR.
	// In HTML: right: 31.82mm.
	// So we can draw a block of width 210.0 - 14.17 - 31.82 = 164.01mm.
	// To keep it simple: draw right-aligned inside a block of width ctx.GetDrawableWidth() - (31.82 - MR).
	// Let's do that!
	// 31.82 - MR(8.82) = 23.0mm from the drawable right margin.
	// So width is GetDrawableWidth() - 23.0.
	w := ctx.GetDrawableWidth() - (31.82 - ctx.MarginRight)

	if err := ctx.SetFont("", 10); err != nil {
		return err
	}
	ctx.DrawLogicalText(0, -28.37, w, utils.PtToMm(10), "PART 2", "R", "", false)

	if err := ctx.SetFont("B", 10); err != nil {
		return err
	}
	ctx.DrawLogicalText(0, -20.61, w, utils.PtToMm(10), " Annex B", "R", "", false)

	// Page Title: Illustration for Computation of APR
	if err := ctx.SetFont("", 10); err != nil {
		return err
	}
	// Y = 47.61 - MT(53.49) = -5.88
	ctx.DrawLogicalText(0, -5.88, ctx.GetDrawableWidth(), utils.PtToMm(10), "Illustration for Computation of APR", "C", "", false)

	return nil
}
