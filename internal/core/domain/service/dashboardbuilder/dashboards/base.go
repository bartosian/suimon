package dashboards

import (
	"fmt"

	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/suimon/internal/core/domain/host"
)

const (
	dashboardName  = "💧 SUIMON: PRESS Q or ESC TO QUIT"
	emptyRowHeight = 1
)

var DashboardConfigDefault = []container.Option{
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

var CellConfigDefault = []container.Option{
	container.Border(linestyle.Light),
	container.AlignVertical(align.VerticalMiddle),
	container.AlignHorizontal(align.HorizontalCenter),
	container.TitleColor(cell.ColorRed),
	container.FocusedColor(cell.ColorWhite),
}

// GetColumnsConfig returns the columns configuration based on the specified dashboard type.
func GetColumnsConfig(dashboard enums.TableType) (ColumnsConfig, error) {
	configMap := map[enums.TableType]ColumnsConfig{
		enums.TableTypeNode:               ColumnsConfigNode,
		enums.TableTypeValidator:          ColumnsConfigValidator,
		enums.TableTypeRPC:                ColumnsConfigRPC,
		enums.TableTypeGasPriceAndSubsidy: ColumnsConfigSystemState,
	}

	config, ok := configMap[dashboard]
	if !ok {
		return nil, fmt.Errorf("unknown dashboard type: %v", dashboard)
	}

	return config, nil
}

// GetColumnsValues returns the columns values based on the specified dashboard type and host.
func GetColumnsValues(dashboard enums.TableType, host *domainhost.Host) (ColumnValues, error) {
	columnsValuesFuncMap := map[enums.TableType]func(*domainhost.Host) (ColumnValues, error){
		enums.TableTypeNode:               GetNodeColumnValues,
		enums.TableTypeValidator:          GetValidatorColumnValues,
		enums.TableTypeRPC:                GetRPCColumnValues,
		enums.TableTypeGasPriceAndSubsidy: GeSystemStateColumnValues,
	}

	if columnValuesFunc, ok := columnsValuesFuncMap[dashboard]; ok {
		return columnValuesFunc(host)
	}

	return nil, fmt.Errorf("unknown dashboard type: %v", dashboard)
}

// GetRowsConfig returns the rows configuration based on the specified dashboard type.
func GetRowsConfig(dashboard enums.TableType) (RowsConfig, error) {
	rowsConfigMap := map[enums.TableType]RowsConfig{
		enums.TableTypeNode:               RowsConfigNode,
		enums.TableTypeValidator:          RowsConfigValidator,
		enums.TableTypeRPC:                RowsConfigRPC,
		enums.TableTypeGasPriceAndSubsidy: RowsConfigSystemState,
	}

	config, ok := rowsConfigMap[dashboard]
	if !ok {
		return nil, fmt.Errorf("unknown dashboard type: %v", dashboard)
	}

	return config, nil
}

// GetCellsConfig returns the cells configuration based on the specified dashboard type.
func GetCellsConfig(dashboard enums.TableType) (CellsConfig, error) {
	cellsConfigMap := map[enums.TableType]CellsConfig{
		enums.TableTypeNode:               CellsConfigNode,
		enums.TableTypeValidator:          CellsConfigValidator,
		enums.TableTypeRPC:                CellsConfigRPC,
		enums.TableTypeGasPriceAndSubsidy: CellsConfigSystemState,
	}

	config, ok := cellsConfigMap[dashboard]
	if !ok {
		return nil, fmt.Errorf("unknown dashboard type: %v", dashboard)
	}

	return config, nil
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

// GetRows creates a new set of rows based on the configuration provided.
// It accepts a RowsConfig object that specifies the height and columns for each row,
// a Cells object that maps cell names to cell objects,
// and a Columns object that maps column names to column objects.
// It returns a Rows object and an error. The Rows object is a slice of Row objects,
// where each Row object represents a row in the terminal grid.
func GetRows(rowsConfig RowsConfig, columns Columns) (Rows, error) {
	rows := make(Rows, 0, len(rowsConfig))

	for _, rowConfig := range rowsConfig {
		rowColumns := make([]grid.Element, 0, len(rowConfig.Columns))

		for _, columnName := range rowConfig.Columns {
			column, ok := columns[columnName]
			if !ok {
				return nil, fmt.Errorf("failed to get column %s", columnName)
			}

			rowColumns = append(rowColumns, column)
		}

		rows = append(rows, NewRowPct(rowConfig.Height, rowColumns...))
	}

	// add empty row to limit last row height
	rows = append(rows, NewRowPct(emptyRowHeight))

	return rows, nil
}

// GetCells creates a new set of cells based on the configuration provided.
// It accepts a CellsConfig object that maps column names to cell names,
// and a terminal object that represents the terminal where the cells will be displayed.
// It returns a Cells object and an error. The Cells object is a map that maps
// column names to cell objects.
func GetCells(cellsConfig CellsConfig) (Cells, error) {
	cells := make(Cells, len(cellsConfig))

	for columnName, cellConfig := range cellsConfig {
		widget, err := newWidgetByColumnName(columnName, cellConfig.Color)
		if err != nil {
			return nil, err
		}

		dashCell, err := NewCell(cellConfig.Title, cellConfig.Color, widget)
		if err != nil {
			return nil, fmt.Errorf("failed to create new cell for %s: %w", columnName, err)
		}

		cells[columnName] = dashCell
	}

	return cells, nil
}
