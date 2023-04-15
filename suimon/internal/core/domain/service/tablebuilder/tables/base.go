package tables

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
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
	Columns      ColumnsConfig // Configuration for the table's columns
	Rows         RowsConfig    // Configuration for the table's rows
	ColumnsCount int           // The total number of columns in the table
	RowsCount    int           // The total number of rows in the table
}

// NewDefaultTableConfig returns a new default table configuration based on the specified table type.
// It sets the table name, style, sort, rows, columns, column count, and auto-index.
func NewDefaultTableConfig(table enums.TableType) *TableConfig {
	tableName := fmt.Sprintf("%s [ %s ]", suiEmoji, table)
	columnsConfig := GetColumnsConfig(table)
	rowsConfig := GetRowsConfig(table)
	tableColor := GetTableColor(table)

	tableStyle := tableStyleDefault
	tableStyle.Title.Colors = tableColor
	tableStyle.Color.Footer = tableColor

	return &TableConfig{
		Name:         tableName,
		Style:        tableStyle,
		Rows:         rowsConfig,
		Columns:      columnsConfig,
		ColumnsCount: len(columnsConfig),
	}
}

// GetColumnsConfig returns the columns configuration based on the specified table type.
func GetColumnsConfig(table enums.TableType) ColumnsConfig {
	switch table {
	case enums.TableTypeRPC:
		return ColumnsConfigRPC
	case enums.TableTypePeers:
		return ColumnsConfigPeer
	case enums.TableTypeValidator:
		return ColumnsConfigValidator
	case enums.TableTypeNode:
		return ColumnsConfigNode
	case enums.TableTypeSystemState,
		enums.TableTypeValidatorsCounts,
		enums.TableTypeValidatorsAtRisk,
		enums.TableTypeValidatorReports:
		return ColumnsConfigSystem
	case enums.TableTypeActiveValidators:
		return ColumnsConfigActiveValidator
	default:
		return nil
	}
}

// GetRowsConfig returns the rows configuration based on the specified table type.
func GetRowsConfig(table enums.TableType) RowsConfig {
	switch table {
	case enums.TableTypeRPC:
		return RowsConfigRPC
	case enums.TableTypePeers:
		return RowsConfigPeer
	case enums.TableTypeValidator:
		return RowsConfigValidator
	case enums.TableTypeSystemState:
		return RowsConfigSystemState
	case enums.TableTypeValidatorsCounts:
		return RowsConfigValidatorCounts
	case enums.TableTypeValidatorsAtRisk:
		return RowsConfigValidatorsAtRisk
	case enums.TableTypeValidatorReports:
		return RowsConfigValidatorReports
	case enums.TableTypeNode:
		return RowsConfigNode
	case enums.TableTypeActiveValidators:
		return RowsActiveValidator
	default:
		return nil
	}
}

// GetTableColor returns the color configuration based on the specified table type.
func GetTableColor(table enums.TableType) text.Colors {
	switch table {
	case enums.TableTypeRPC, enums.TableTypeValidator, enums.TableTypeSystemState, enums.TableTypeValidatorsAtRisk, enums.TableTypeActiveValidators:
		return text.Colors{text.BgHiBlue, text.FgBlack}
	default:
		return text.Colors{text.BgHiGreen, text.FgBlack}
	}
}
