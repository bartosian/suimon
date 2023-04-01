package tables

import (
	"fmt"

	styles "github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums/columnnames"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/tablebuilder"
)

const suiEmoji = "💧"

type (
	tableRows         [][]int
	tableSortConfig   []styles.SortBy
	tableColumnConfig []styles.ColumnConfig

	ColumnName interface {
		columnnames.NodeColumnName | columnnames.ValidatorColumnName | columnnames.ActiveValidatorColumnName | columnnames.SystemColumnName
	}
)

var nameTransformer = text.Transformer(func(val interface{}) string {
	return text.Bold.Sprint(val)
})

func SetColumnValues[E ColumnName](columns tablebuilder.Columns, values map[E]any) {
	for name, value := range values {
		columns[name].SetValue(value)
	}
}

func GetTableTitle(table enums.TableType) string {
	return fmt.Sprintf("%s [ %s ]", suiEmoji, table)
}

func GetTableTag(table enums.TableType) string {
	switch table {
	case enums.TableTypeRPC:
		return TableTagRPC
	case enums.TableTypePeers:
		return TableTagPeer
	case enums.TableTypeValidator:
		return TableTagValidator
	case enums.TableTypeSystemState:
		return TableTagSystem
	case enums.TableTypeNode:
		return TableTagNode
	case enums.TableTypeActiveValidators:
		return TableTagActiveValidator
	default:
		return ""
	}
}

func GetTableStyle(table enums.TableType) styles.Style {
	switch table {
	case enums.TableTypeRPC:
		return TableStyleRPC
	case enums.TableTypePeers:
		return TableStylePeer
	case enums.TableTypeValidator:
		return TableStyleValidator
	case enums.TableTypeSystemState:
		return TableStyleSystem
	case enums.TableTypeNode:
		return TableStyleNode
	case enums.TableTypeActiveValidators:
		return TableStyleActiveValidator
	default:
		return styles.StyleLight
	}
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
	case enums.TableTypeSystemState:
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
	case enums.TableTypeSystemState:
		return ColumnsConfigSystem
	case enums.TableTypeNode:
		return ColumnsConfigNode
	case enums.TableTypeActiveValidators:
		return ColumnConfigActiveValidator
	default:
		return nil
	}
}
