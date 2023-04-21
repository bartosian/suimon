package dashboards

import (
	"fmt"
	"strings"
	"time"

	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
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

type Element interface {
	isElement()
}

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
	rows := make(Rows, 0, len(rowsConfig))

	for rowIdx, rowConfig := range rowsConfig {
		rowColumns := make([]grid.Element, len(rowConfig.Columns))

		for columnIdx, columnName := range rowConfig.Columns {
			column, ok := columns[columnName]
			if !ok {
				return nil, fmt.Errorf("failed to get column %s", columnName)
			}

			rowColumns[columnIdx] = column
		}

		rows[rowIdx] = NewRowPct(rowConfig.Height, rowColumns...)
	}

	return rows, nil
}

// GetCells returns a map of Cells based on the given CellsConfig.
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

func DashboardLoadingBlinkValue() string {
	inProgress := strings.Repeat("-", 50)
	second := time.Now().Second() % 10

	if second%2 == 0 {
		inProgress = strings.Repeat("\u0020", 50)
	}

	return inProgress
}
