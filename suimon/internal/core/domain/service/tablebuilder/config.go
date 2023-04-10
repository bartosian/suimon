package tablebuilder

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/tablebuilder/tables"
)

const suiEmoji = "💧"

var (
	tableStyleDefault = table.Style{
		Name: "DEFAULT",
		Box: table.BoxStyle{
			BottomLeft:       "└",
			BottomRight:      "┘",
			BottomSeparator:  "┴",
			EmptySeparator:   text.RepeatAndTrim(" ", text.RuneWidthWithoutEscSequences("┼")),
			Left:             "│",
			LeftSeparator:    "├",
			MiddleHorizontal: "─",
			MiddleSeparator:  "┼",
			MiddleVertical:   "│",
			PaddingLeft:      " ",
			PaddingRight:     " ",
			PageSeparator:    "\n",
			Right:            "│",
			RightSeparator:   "┤",
			TopLeft:          "┌",
			TopRight:         "┐",
			TopSeparator:     "┬",
			UnfinishedRow:    " ≈",
		},
		Color: table.ColorOptions{
			Header: text.Colors{text.FgBlack, text.BgWhite},
			Row:    text.Colors{text.BgWhite},
			Footer: text.Colors{text.BgHiBlue, text.FgBlack},
		},
		Options: table.Options{
			DoNotColorBordersAndSeparators: true,
			DrawBorder:                     true,
			SeparateColumns:                true,
			SeparateFooter:                 true,
			SeparateHeader:                 true,
			SeparateRows:                   true,
		},
		Title: table.TitleOptions{
			Align:  text.AlignLeft,
			Colors: text.Colors{text.BgHiBlue, text.FgBlack},
		},
	}
)

// RowsConfig represents the configuration of rows in a table, each row being a slice of column names.
type RowsConfig [][]enums.ColumnName

// SortConfig represents a slice of sorting configurations for a table.
type SortConfig []table.SortBy

// TableConfig represents the overall configuration for a table.
type TableConfig struct {
	Name         string        // The name of the table
	Style        table.Style   // The style of the table
	Sort         SortConfig    // The sorting configuration for the table
	AutoIndex    bool          // Whether to automatically generate indices for the table
	Columns      ColumnsConfig // Configuration for the table's columns
	Rows         RowsConfig    // Configuration for the table's rows
	ColumnsCount int           // The total number of columns in the table
	RowsCount    int           // The total number of rows in the table
}

// NewDefaultTableConfig returns a new default table configuration based on the specified table type.
// It sets the table name, style, sort, rows, columns, column count, and auto-index.
func NewDefaultTableConfig(table enums.TableType) *TableConfig {
	tableName := fmt.Sprintf("%s [ %s ]", suiEmoji, table)
	columnsConfig := getColumnsConfig(table)
	sortConfig := getTableSortConfig(table)
	rowsConfig := getRowsConfig(table)
	autoIndex := getAutoIndex(table)

	return &TableConfig{
		Name:         tableName,
		Style:        tableStyleDefault,
		Sort:         sortConfig,
		Rows:         rowsConfig,
		Columns:      columnsConfig,
		ColumnsCount: len(columnsConfig),
		AutoIndex:    autoIndex,
	}
}

// getColumnsConfig returns the columns configuration based on the specified table type.
func getColumnsConfig(table enums.TableType) ColumnsConfig {
	switch table {
	case enums.TableTypeRPC:
		return tables.ColumnsConfigRPC
	case enums.TableTypePeers:
		return tables.ColumnsConfigPeer
	case enums.TableTypeValidator:
		return tables.ColumnsConfigValidator
	case enums.TableTypeNode:
		return tables.ColumnsConfigNode
	case enums.TableTypeSystemState,
		enums.TableTypeValidatorsCounts,
		enums.TableTypeValidatorsAtRisk,
		enums.TableTypeValidatorReports:
		return tables.ColumnsConfigSystem
	case enums.TableTypeActiveValidators:
		return tables.ColumnsConfigActiveValidator
	default:
		return nil
	}
}

// getTableSortConfig returns the sort configuration based on the specified table type.
func getTableSortConfig(table enums.TableType) SortConfig {
	switch table {
	case enums.TableTypeRPC:
		return tables.SortConfigRPC
	case enums.TableTypePeers:
		return tables.SortConfigPeer
	case enums.TableTypeValidator:
		return tables.SortConfigValidator
	case enums.TableTypeSystemState,
		enums.TableTypeValidatorsCounts,
		enums.TableTypeValidatorsAtRisk,
		enums.TableTypeValidatorReports:
		return tables.SortConfigSystem
	case enums.TableTypeNode:
		return tables.SortConfigNode
	case enums.TableTypeActiveValidators:
		return tables.SortConfigActiveValidator
	default:
		return nil
	}
}

// getRowsConfig returns the rows configuration based on the specified table type.
func getRowsConfig(table enums.TableType) RowsConfig {
	switch table {
	case enums.TableTypeRPC:
		return tables.RowsConfigRPC
	case enums.TableTypePeers:
		return tables.RowsConfigPeer
	case enums.TableTypeValidator:
		return tables.RowsConfigValidator
	case enums.TableTypeSystemState:
		return tables.RowsConfigSystemState
	case enums.TableTypeValidatorsCounts:
		return tables.RowsConfigValidatorCounts
	case enums.TableTypeValidatorsAtRisk:
		return tables.RowsConfigValidatorsAtRisk
	case enums.TableTypeValidatorReports:
		return tables.RowsConfigValidatorReports
	case enums.TableTypeNode:
		return tables.RowsConfigNode
	case enums.TableTypeActiveValidators:
		return tables.RowsActiveValidator
	default:
		return nil
	}
}

// getAutoIndex returns true if the specified table type needs to be auto-indexed; otherwise, false.
func getAutoIndex(table enums.TableType) bool {
	tableToAutoIndexMap := map[enums.TableType]bool{
		enums.TableTypeSystemState:      false,
		enums.TableTypeValidatorsCounts: false,
		enums.TableTypeValidatorReports: false,
		enums.TableTypePeers:            false,
		enums.TableTypeNode:             false,
		enums.TableTypeValidatorsAtRisk: false,
		enums.TableTypeValidator:        false,
		enums.TableTypeRPC:              true,
		enums.TableTypeActiveValidators: true,
	}

	return tableToAutoIndexMap[table]
}
