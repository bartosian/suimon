package dashboards

import (
	"fmt"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"
)

const dashboardName = "💧 SUIMON: PRESS Q or ESC TO QUIT"

var (
	DashboardConfigDefault = []container.Option{
		container.Border(linestyle.Light),
		container.BorderColor(cell.ColorGreen),
		container.BorderTitle(dashboardName),
		container.FocusedColor(cell.ColorGreen),
		container.AlignHorizontal(align.HorizontalCenter),
		container.AlignVertical(align.VerticalMiddle),
		container.TitleColor(cell.ColorRed),
		container.TitleFocusedColor(cell.ColorRed),
		container.Focused(),
	}

	CellConfigDefault = []container.Option{
		container.Border(linestyle.Light),
		container.BorderColor(cell.ColorGreen),
		container.AlignVertical(align.VerticalMiddle),
		container.AlignHorizontal(align.HorizontalCenter),
		container.TitleColor(cell.ColorRed),
		container.FocusedColor(cell.ColorWhite),
		container.FocusedColor(cell.ColorBlue),
	}
)

// GetColumnsConfig returns the columns configuration based on the specified dashboard type.
func GetColumnsConfig(dashboard enums.TableType) ColumnsConfig {
	switch dashboard {
	case enums.TableTypeNode:
		return ColumnsConfigNode
	default:
		return nil
	}
}

// GetRowsConfig returns the rows configuration based on the specified dashboard type.
func GetRowsConfig(dashboard enums.TableType) RowsConfig {
	switch dashboard {
	case enums.TableTypeNode:
		return RowsConfigNode
	default:
		return nil
	}
}

// GetCellsConfig returns the cells configuration based on the specified dashboard type.
func GetCellsConfig(dashboard enums.TableType) CellsConfig {
	switch dashboard {
	case enums.TableTypeNode:
		return CellsConfigNode
	default:
		return nil
	}
}

// GetColumns returns a slice of Columns based on the given ColumnsConfig and Cells.
func GetColumns(columnsConfig ColumnsConfig, cells Cells) (Columns, error) {
	columns := make(Columns, len(columnsConfig))

	for columnName, width := range columnsConfig {
		dashCell, ok := cells[columnName]
		if !ok {
			return nil, fmt.Errorf("failed to get cell for column %s", columnName)
		}

		columnPct := NewColumnPct(width, dashCell.GetWidget())

		columns[columnName] = columnPct
	}

	return columns, nil
}

// GetRows returns a slice of Rows based on the given RowsConfig and Columns.
func GetRows(rowsConfig RowsConfig, columns Columns) (Rows, error) {
	rows := make(Rows, len(rowsConfig)+1)

	for rowIdx, rowConfig := range rowsConfig {
		rowColumns := make([]grid.Element, len(rowConfig.Columns)+1)

		for columnIdx, columnName := range rowConfig.Columns {
			column, ok := columns[columnName]
			if !ok {
				return nil, fmt.Errorf("failed to get column %s", columnName)
			}

			rowColumns[columnIdx] = column
		}

		rows[rowIdx] = NewRowPct(rowConfig.Height, rowColumns...)
	}

	// add empty row to limit last row height
	rows = append(rows, NewRowPct(1))

	return rows, nil
}

// GetCells creates a new set of cells based on the configuration provided.
// It accepts a CellsConfig object that maps column names to cell names,
// and a terminal object that represents the terminal where the cells will be displayed.
// It returns a Cells object and an error. The Cells object is a map that maps
// column names to cell objects.
func GetCells(cellsConfig CellsConfig) (Cells, error) {
	cells := make(Cells, len(cellsConfig))

	for columnName, cellName := range cellsConfig {
		widget, err := newWidgetByColumnName(columnName)
		if err != nil {
			return nil, err
		}

		dashCell, err := NewCell(cellName, widget)
		if err != nil {
			return nil, fmt.Errorf("failed to create new cell for %s: %w", columnName, err)
		}

		cells[columnName] = dashCell
	}

	return cells, nil
}
