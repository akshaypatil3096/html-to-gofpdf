package renderer

// TableCell represents a layout block inside a Table row, supporting rowspan and colspan.
type TableCell struct {
	Component Component
	ColSpan   int
	RowSpan   int
}

// Table displays a dynamic grid of TableCells with automatic column normalization and sizing.
type Table struct {
	ColWidths []float64 // Widths as absolute mm values or percentage ratios.
	Rows      [][]TableCell
}

// gridElement holds grid lookup mapping info for spans.
type gridElement struct {
	cell      *TableCell
	width     float64
	originRow int
	originCol int
	isSpanned bool
}

// buildGrid resolves table rowspans and colspans into a flat 2D grid of elements.
func (t *Table) buildGrid() ([][]gridElement, int) {
	numRows := len(t.Rows)
	numCols := len(t.ColWidths)

	grid := make([][]gridElement, numRows)
	for i := 0; i < numRows; i++ {
		grid[i] = make([]gridElement, numCols)
	}

	for r := 0; r < numRows; r++ {
		cIdx := 0
		for _, cell := range t.Rows[r] {
			// Skip past grid cells marked as spanned by rowspans/colspans
			for cIdx < numCols && grid[r][cIdx].isSpanned {
				cIdx++
			}
			if cIdx >= numCols {
				break
			}

			colSpan := cell.ColSpan
			if colSpan <= 0 {
				colSpan = 1
			}
			rowSpan := cell.RowSpan
			if rowSpan <= 0 {
				rowSpan = 1
			}

			// Sum up spanned column widths
			w := 0.0
			for i := 0; i < colSpan && (cIdx+i) < numCols; i++ {
				w += t.ColWidths[cIdx+i]
			}

			cellRef := cell
			for rs := 0; rs < rowSpan && (r+rs) < numRows; rs++ {
				for cs := 0; cs < colSpan && (cIdx+cs) < numCols; cs++ {
					grid[r+rs][cIdx+cs] = gridElement{
						cell:      &cellRef,
						width:     w,
						originRow: r,
						originCol: cIdx,
						isSpanned: rs > 0 || cs > 0,
					}
				}
			}
			cIdx += colSpan
		}
	}
	return grid, numCols
}

// normalizeColWidths scales column ratios to match the target layout width.
func (t *Table) normalizeColWidths(maxWidth float64) []float64 {
	sum := 0.0
	for _, w := range t.ColWidths {
		sum += w
	}
	if sum == 0 {
		widths := make([]float64, len(t.ColWidths))
		for i := range widths {
			widths[i] = maxWidth / float64(len(widths))
		}
		return widths
	}

	widths := make([]float64, len(t.ColWidths))
	for i, w := range t.ColWidths {
		widths[i] = (w / sum) * maxWidth
	}
	return widths
}

// Measure calculates grid row heights to compute the total table height.
func (t *Table) Measure(ctx *Context, maxWidth float64) (Size, error) {
	if len(t.Rows) == 0 {
		return Size{Width: maxWidth, Height: 0}, nil
	}

	colWidths := t.normalizeColWidths(maxWidth)
	grid, numCols := t.buildGrid()
	numRows := len(t.Rows)

	rowHeights := make([]float64, numRows)

	// Phase 1: Measure non-rowspanned cells (rowspan = 1)
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			elem := grid[r][c]
			if elem.cell != nil && !elem.isSpanned {
				if elem.cell.RowSpan <= 1 {
					colSpan := elem.cell.ColSpan
					if colSpan <= 0 {
						colSpan = 1
					}
					cellWidth := 0.0
					for i := 0; i < colSpan && (c+i) < numCols; i++ {
						cellWidth += colWidths[c+i]
					}

					sz, err := elem.cell.Component.Measure(ctx, cellWidth)
					if err != nil {
						return Size{}, err
					}
					if sz.Height > rowHeights[r] {
						rowHeights[r] = sz.Height
					}
				}
			}
		}
	}

	// Phase 2: Measure rowspanned cells and distribute heights proportionally
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			elem := grid[r][c]
			if elem.cell != nil && !elem.isSpanned {
				if elem.cell.RowSpan > 1 {
					colSpan := elem.cell.ColSpan
					if colSpan <= 0 {
						colSpan = 1
					}
					cellWidth := 0.0
					for i := 0; i < colSpan && (c+i) < numCols; i++ {
						cellWidth += colWidths[c+i]
					}

					sz, err := elem.cell.Component.Measure(ctx, cellWidth)
					if err != nil {
						return Size{}, err
					}

					rowSpan := elem.cell.RowSpan
					spanHeight := 0.0
					for rs := 0; rs < rowSpan && (r+rs) < numRows; rs++ {
						spanHeight += rowHeights[r+rs]
					}

					if sz.Height > spanHeight {
						diff := sz.Height - spanHeight
						added := diff / float64(rowSpan)
						for rs := 0; rs < rowSpan && (r+rs) < numRows; rs++ {
							rowHeights[r+rs] += added
						}
					}
				}
			}
		}
	}

	totalHeight := 0.0
	for _, rh := range rowHeights {
		totalHeight += rh
	}

	return Size{Width: maxWidth, Height: totalHeight}, nil
}

// Draw renders each grid cell inside resolved boundaries.
func (t *Table) Draw(ctx *Context, x, y, width, height float64) error {
	if len(t.Rows) == 0 {
		return nil
	}

	colWidths := t.normalizeColWidths(width)
	grid, numCols := t.buildGrid()
	numRows := len(t.Rows)

	// Precompute row heights
	rowHeights := make([]float64, numRows)
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			elem := grid[r][c]
			if elem.cell != nil && !elem.isSpanned {
				if elem.cell.RowSpan <= 1 {
					colSpan := elem.cell.ColSpan
					if colSpan <= 0 {
						colSpan = 1
					}
					cellWidth := 0.0
					for i := 0; i < colSpan && (c+i) < numCols; i++ {
						cellWidth += colWidths[c+i]
					}
					sz, _ := elem.cell.Component.Measure(ctx, cellWidth)
					if sz.Height > rowHeights[r] {
						rowHeights[r] = sz.Height
					}
				}
			}
		}
	}

	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			elem := grid[r][c]
			if elem.cell != nil && !elem.isSpanned {
				if elem.cell.RowSpan > 1 {
					colSpan := elem.cell.ColSpan
					if colSpan <= 0 {
						colSpan = 1
					}
					cellWidth := 0.0
					for i := 0; i < colSpan && (c+i) < numCols; i++ {
						cellWidth += colWidths[c+i]
					}
					sz, _ := elem.cell.Component.Measure(ctx, cellWidth)
					rowSpan := elem.cell.RowSpan
					spanHeight := 0.0
					for rs := 0; rs < rowSpan && (r+rs) < numRows; rs++ {
						spanHeight += rowHeights[r+rs]
					}
					if sz.Height > spanHeight {
						diff := sz.Height - spanHeight
						added := diff / float64(rowSpan)
						for rs := 0; rs < rowSpan && (r+rs) < numRows; rs++ {
							rowHeights[r+rs] += added
						}
					}
				}
			}
		}
	}

	// Compute cell top and left offsets
	rowY := make([]float64, numRows)
	currentY := 0.0
	for r := 0; r < numRows; r++ {
		rowY[r] = currentY
		currentY += rowHeights[r]
	}

	colX := make([]float64, numCols)
	currentX := 0.0
	for c := 0; c < numCols; c++ {
		colX[c] = currentX
		currentX += colWidths[c]
	}

	// Render table cell components at calculated coordinates
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			elem := grid[r][c]
			if elem.cell != nil && !elem.isSpanned {
				cellX := colX[c]
				cellY := rowY[r]

				colSpan := elem.cell.ColSpan
				if colSpan <= 0 {
					colSpan = 1
				}
				rowSpan := elem.cell.RowSpan
				if rowSpan <= 0 {
					rowSpan = 1
				}

				cellWidth := 0.0
				for i := 0; i < colSpan && (c+i) < numCols; i++ {
					cellWidth += colWidths[c+i]
				}

				cellHeight := 0.0
				for i := 0; i < rowSpan && (r+i) < numRows; i++ {
					cellHeight += rowHeights[r+i]
				}

				err := elem.cell.Component.Draw(ctx, x+cellX, y+cellY, cellWidth, cellHeight)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
