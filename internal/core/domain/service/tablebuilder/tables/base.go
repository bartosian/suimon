package tables

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/suimon/internal/core/domain/enums"
)

const suiEmoji = "💧"

// tableStyleDefault defines the default style for the tables.
// It includes the box style, color options, and other table options.
// The box style defines the characters used for the table borders and separators.
// The color options define the colors used for the header, rows, and footer.
// The table options define various settings such as whether to draw borders and separate columns, rows, etc.
var tableStyleDefault = table.Style{
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

// Define the mapping of TableType enums to their corresponding ColumnsConfig.
var columnsConfigMap = map[enums.TableType]ColumnsConfig{
	enums.TableTypeRPC:                ColumnsConfigRPC,
	enums.TableTypeValidator:          ColumnsConfigValidator,
	enums.TableTypeNode:               ColumnsConfigNode,
	enums.TableTypeGasPriceAndSubsidy: ColumnsConfigSystem,
	enums.TableTypeValidatorsParams:   ColumnsConfigSystem,
	enums.TableTypeValidatorsAtRisk:   ColumnsConfigSystem,
	enums.TableTypeValidatorReports:   ColumnsConfigSystem,
	enums.TableTypeActiveValidators:   ColumnsConfigActiveValidator,
	enums.TableTypeReleases:           ColumnsConfigRelease,
	enums.TableTypeProtocol:           ColumnsConfigProtocol,
}

// Define the mapping of TableType enums to their corresponding RowsConfig.
var rowsConfigMap = map[enums.TableType]RowsConfig{
	enums.TableTypeRPC:                RowsConfigRPC,
	enums.TableTypeValidator:          RowsConfigValidator,
	enums.TableTypeGasPriceAndSubsidy: RowsConfigSystemState,
	enums.TableTypeValidatorsParams:   RowsConfigValidatorParams,
	enums.TableTypeValidatorsAtRisk:   RowsConfigValidatorsAtRisk,
	enums.TableTypeValidatorReports:   RowsConfigValidatorReports,
	enums.TableTypeNode:               RowsConfigNode,
	enums.TableTypeActiveValidators:   RowsActiveValidator,
	enums.TableTypeReleases:           RowsRelease,
	enums.TableTypeProtocol:           RowsConfigProtocol,
}

// Define the mapping of TableType enums to their corresponding text.Colors.
var tableColorMap = map[enums.TableType]text.Colors{
	enums.TableTypeRPC:              {text.BgHiBlue, text.FgBlack},
	enums.TableTypeValidator:        {text.BgHiBlue, text.FgBlack},
	enums.TableTypeValidatorsAtRisk: {text.BgHiBlue, text.FgBlack},
	enums.TableTypeActiveValidators: {text.BgHiBlue, text.FgBlack},
	enums.TableTypeProtocol:         {text.BgHiBlue, text.FgBlack},
}

// defaultTableColor defines the default color configuration.
var defaultTableColor = text.Colors{text.BgHiGreen, text.FgBlack}

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
func NewDefaultTableConfig(domainTable enums.TableType) *TableConfig {
	tableName := fmt.Sprintf("%s [ %s ]", suiEmoji, domainTable)
	columnsConfig := GetColumnsConfig(domainTable)
	rowsConfig := GetRowsConfig(domainTable)
	tableColor := GetTableColor(domainTable)

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

// GetColumnsConfig returns the ColumnsConfig based on the provided table type.
func GetColumnsConfig(domainTable enums.TableType) ColumnsConfig {
	config, ok := columnsConfigMap[domainTable]
	if !ok {
		return nil
	}

	return config
}

// GetRowsConfig returns the rows configuration based on the specified table type.
func GetRowsConfig(domainTable enums.TableType) RowsConfig {
	config, ok := rowsConfigMap[domainTable]
	if !ok {
		return nil
	}

	return config
}

// GetTableColor returns the color configuration based on the specified table type.
func GetTableColor(domainTable enums.TableType) text.Colors {
	color, ok := tableColorMap[domainTable]
	if !ok {
		return defaultTableColor
	}

	return color
}
