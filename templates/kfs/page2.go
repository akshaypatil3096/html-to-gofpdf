package kfs

import (
	"html-to-gofpdf/models"
	"html-to-gofpdf/renderer"
	"html-to-gofpdf/utils"
)

// RenderPage2 draws the layout of KFS Part 2.
func RenderPage2(ctx *renderer.Context, req models.PDFRequest) error {
	// 1. Render Header (Logo and Titles)
	if err := renderPage2Header(ctx, req); err != nil {
		return err
	}

	// 2. Build Table 1 (Other Qualitative info list)
	colWidths1 := []float64{6.5, 46.5, 47.0}
	rows1 := [][]renderer.TableCell{
		// Row 1
		{
			emptyCell(),
			cCell("Clause of Loan agreement relating to engagement of recovery agents", false, "L", 1, 1, 10.0*utils.PxToMmFactor, 30.0*utils.PxToMmFactor, 30.0*utils.PxToMmFactor, 30.0*utils.PxToMmFactor),
			cCell("Not applicable", false, "C", 1, 1),
		},
		// Row 2
		{
			emptyCell(),
			cCell("Clause of Loan agreement which details grievance redressal mechanism", false, "L", 1, 1, 10.0*utils.PxToMmFactor, 30.0*utils.PxToMmFactor, 30.0*utils.PxToMmFactor, 30.0*utils.PxToMmFactor),
			cCell("Designation: Grievance Redressal\nOfficer Name: Mr. Kannan Ramaseshan\nAddress: Grievance Redressal Cell,\nHDFC Bank Limited, 1st Floor, Empire\nPlaza- 1, Lal Bahadur Shastri Marg,\nChandan Nagar, Vikhroli West, Mumbai\n– 400083.", false, "L", 1, 1, 10.0*utils.PxToMmFactor, 10.0*utils.PxToMmFactor, 10.0*utils.PxToMmFactor, 10.0*utils.PxToMmFactor),
		},
		// Row 3
		{
			emptyCell(),
			cCell("Phone number and email id of the nodal grievance redressal officer", false, "L", 1, 1, 10.0*utils.PxToMmFactor, 30.0*utils.PxToMmFactor, 30.0*utils.PxToMmFactor, 30.0*utils.PxToMmFactor),
			cCell("Toll free number: 18002664060\nEmail: grievance.redressaldl@hdfcbank.com\nAvailability - Monday to Saturday on\n18002664060 between 9.30am to 5.30 pm\nPlease note this facility is not available on 2nd &\n4th Saturdays, all Sundays and Banks Holidays", false, "L", 1, 1, 10.0*utils.PxToMmFactor, 10.0*utils.PxToMmFactor, 10.0*utils.PxToMmFactor, 10.0*utils.PxToMmFactor),
		},
		// Row 4
		{
			emptyCell(),
			cCell("Whether the loan is, or in future maybe, subject to transfer to other REs or securitisation (Yes/ No)", false, "L", 1, 1, 10.0*utils.PxToMmFactor, 40.0*utils.PxToMmFactor, 20.0*utils.PxToMmFactor, 20.0*utils.PxToMmFactor),
			cCell("Not applicable", false, "C", 1, 1),
		},
		// Row 5
		{
			emptyCell(),
			cCell("In case of lending under collaborative lending arrangements (e.g., co-lending/ outsourcing), following additional details may be furnished:", false, "L", 2, 1, 10.0*utils.PxToMmFactor, 60.0*utils.PxToMmFactor, 20.0*utils.PxToMmFactor, 20.0*utils.PxToMmFactor),
		},
	}

	table1 := &renderer.Table{
		ColWidths: colWidths1,
		Rows:      rows1,
	}

	sz1, err := table1.Measure(ctx, ctx.GetDrawableWidth())
	if err != nil {
		return err
	}
	if err := table1.Draw(ctx, 0, 0, sz1.Width, sz1.Height); err != nil {
		return err
	}

	// 3. Build Table 2 (Collaborative lending arrangements)
	colWidths2 := []float64{6.5, 31.16, 31.16, 31.18}
	rows2 := [][]renderer.TableCell{
		// Row 1
		{
			emptyCell(),
			cCell("Name of the originating RE,\nalong with its funding\nproportion", false, "C", 1, 1, 1.0, 1.0, 10.0*utils.PxToMmFactor, 0.7),
			cCell("Name of the partner RE along\nwith its proportion of\nfunding", false, "C", 1, 1, 1.0, 1.0, 10.0*utils.PxToMmFactor, 0.7),
			cCell("Blended rate of\ninterest", false, "C", 1, 1, 1.0, 1.0, 10.0*utils.PxToMmFactor, 0.7),
		},
		// Row 2
		{
			emptyCell(),
			cCell("Not applicable", false, "C", 1, 1, 1.0, 1.0, 10.0*utils.PxToMmFactor, 10.0*utils.PxToMmFactor),
			cCell("Not applicable", false, "C", 1, 1, 1.0, 1.0, 10.0*utils.PxToMmFactor, 10.0*utils.PxToMmFactor),
			cCell("Not applicable", false, "C", 1, 1, 1.0, 1.0, 10.0*utils.PxToMmFactor, 10.0*utils.PxToMmFactor),
		},
	}

	table2 := &renderer.Table{
		ColWidths: colWidths2,
		Rows:      rows2,
	}

	sz2, err := table2.Measure(ctx, ctx.GetDrawableWidth())
	if err != nil {
		return err
	}
	// Shift by -0.26mm to collapse the top border with table1's bottom border
	y2 := sz1.Height - utils.PxToMm(1.0)
	if err := table2.Draw(ctx, 0, y2, sz2.Width, sz2.Height); err != nil {
		return err
	}

	// 4. Build Table 3 (Digital loans specific disclosures)
	colWidths3 := []float64{6.5, 46.5, 47.0}
	rows3 := [][]renderer.TableCell{
		// Row 1
		{
			emptyCell(),
			cCell("In case of digital loans, following specific disclosures may be furnished:", false, "L", 2, 1, 1.0, 1.0, 10.0*utils.PxToMmFactor, 10.0*utils.PxToMmFactor),
		},
		// Row 2
		{
			emptyCell(),
			cCell("Cooling off/look-up period, in terms of RE's board approved policy, during which borrower shall not be charged any penalty on prepayment of loan", false, "L", 1, 1, 1.0, 1.0, 10.0*utils.PxToMmFactor, 10.0*utils.PxToMmFactor),
			cCell("The overdraft can be closed any time without penalty during the tenure of the loan", false, "L", 1, 1, 1.0, 1.0, 10.0*utils.PxToMmFactor, 10.0*utils.PxToMmFactor),
		},
		// Row 3
		{
			emptyCell(),
			cCell("Details of LSP acting as recovery agent and authorized to approach the borrower", false, "L", 1, 1, 1.0, 1.0, 10.0*utils.PxToMmFactor, 10.0*utils.PxToMmFactor),
			cCell("Not applicable", false, "C", 1, 1, 1.0, 1.0, 10.0*utils.PxToMmFactor, 10.0*utils.PxToMmFactor),
		},
	}

	table3 := &renderer.Table{
		ColWidths: colWidths3,
		Rows:      rows3,
	}

	sz3, err := table3.Measure(ctx, ctx.GetDrawableWidth())
	if err != nil {
		return err
	}
	y3 := y2 + sz2.Height - utils.PxToMm(1.0)
	return table3.Draw(ctx, 0, y3, sz3.Width, sz3.Height)
}

func renderPage2Header(ctx *renderer.Context, req models.PDFRequest) error {
	// Logo
	if req.Logo.Path != "" || len(req.Logo.Bytes) > 0 {
		img := &renderer.Image{
			Path:  req.Logo.Path,
			Bytes: req.Logo.Bytes,
		}
		// X = 27.69 - ML(14.17) = 13.52
		// Y = 13.41 - MT(50.94) = -37.53
		if err := img.Draw(ctx, 13.52, -37.53, 31.77, 4.15); err != nil {
			return err
		}
	}

	// Title: Part 2 (Other qualitative information)
	if err := ctx.SetFont("", 10); err != nil {
		return err
	}
	// Y = 46.81 - MT(50.94) = -4.13
	ctx.DrawLogicalText(0, -4.13, ctx.GetDrawableWidth(), utils.PtToMm(10), "Part 2 (Other qualitative information)", "C", "", false)

	return nil
}

func emptyCell() renderer.TableCell {
	return renderer.TableCell{
		Component: &renderer.Cell{
			Border: renderer.BorderSpec{
				Left:   false,
				Right:  false,
				Top:    false,
				Bottom: false,
			},
		},
		ColSpan: 1,
		RowSpan: 1,
	}
}
