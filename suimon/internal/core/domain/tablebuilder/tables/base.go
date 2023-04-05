package tables

import (
	"fmt"

	styles "github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/tablebuilder"
)

const suiEmoji = "💧"

type (
	tableRows         [][]enums.ColumnName
	tableSortConfig   []styles.SortBy
	tableColumnConfig map[enums.ColumnName]styles.ColumnConfig
)

var (
	TableStyleDefault = styles.Style{
		Name: "DEFAULT",
		Box: styles.BoxStyle{
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
		Color: styles.ColorOptions{
			Header: text.Colors{text.FgBlack, text.BgWhite},
			Row:    text.Colors{text.BgWhite},
			Footer: text.Colors{text.BgHiBlue, text.FgBlack},
		},
		Options: styles.Options{
			DoNotColorBordersAndSeparators: true,
			DrawBorder:                     true,
			SeparateColumns:                true,
			SeparateFooter:                 true,
			SeparateHeader:                 true,
			SeparateRows:                   true,
		},
		Title: styles.TitleOptions{
			Align:  text.AlignLeft,
			Colors: text.Colors{text.BgHiBlue, text.FgBlack},
		},
	}
)

func SetColumnValues(columns tablebuilder.Columns, values map[enums.ColumnName]any) {
	for name, value := range values {
		columns[name].SetValue(value)
	}
}

func GetTableTitle(table enums.TableType) string {
	return fmt.Sprintf("%s [ %s ]", suiEmoji, table)
}

func GetTableStyle(table enums.TableType) styles.Style {
	return TableStyleDefault
}

func GetTableRows(table enums.TableType) tableRows {
	switch table {
	case enums.TableTypeRPC:
		return RowsRPC
	case enums.TableTypePeers:
		return RowsPeer
	case enums.TableTypeValidator:
		return RowsValidator
	case enums.TableTypeSystemState:
		return RowsSystemState
	case enums.TableTypeValidatorsCounts:
		return RowsValidatorCounts
	case enums.TableTypeValidatorsAtRisk:
		return RowsValidatorsAtRisk
	case enums.TableTypeValidatorReports:
		return RowsValidatorReports
	case enums.TableTypeNode:
		return RowsNode
	case enums.TableTypeActiveValidators:
		return RowsActiveValidator
	default:
		return nil
	}
}

func GetTableSortConfig(table enums.TableType) tableSortConfig {
	switch table {
	case enums.TableTypeRPC:
		return TableSortConfigRPC
	case enums.TableTypePeers:
		return TableSortConfigPeer
	case enums.TableTypeValidator:
		return TableSortConfigValidator
	case enums.TableTypeSystemState,
		enums.TableTypeValidatorsCounts,
		enums.TableTypeValidatorsAtRisk,
		enums.TableTypeValidatorReports:
		return TableSortConfigSystem
	case enums.TableTypeNode:
		return TableSortConfigNode
	case enums.TableTypeActiveValidators:
		return TableSortConfigActiveValidator
	default:
		return nil
	}
}

func GetColumnConfig(table enums.TableType) tableColumnConfig {
	switch table {
	case enums.TableTypeRPC:
		return ColumnsConfigRPC
	case enums.TableTypePeers:
		return ColumnsConfigPeer
	case enums.TableTypeValidator:
		return ColumnsConfigValidator
	case enums.TableTypeSystemState,
		enums.TableTypeValidatorsCounts,
		enums.TableTypeValidatorsAtRisk,
		enums.TableTypeValidatorReports:
		return ColumnsConfigSystem
	case enums.TableTypeNode:
		return ColumnsConfigNode
	case enums.TableTypeActiveValidators:
		return ColumnConfigActiveValidator
	default:
		return nil
	}
}

func GetAutoIndexConfig(table enums.TableType) bool {
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
