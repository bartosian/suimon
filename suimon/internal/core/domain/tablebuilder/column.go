package tablebuilder

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/jedib0t/go-pretty/v6/table"
)

const (
	TableNoData = "no data"
	EmptyValue  = " "
)

type (
	Columns map[enums.ColumnName]Column

	Column struct {
		Values []any
		Config table.ColumnConfig
	}
)

func (col *Column) SetValue(value any) {
	if value == nil || value == "" {
		value = TableNoData
	}

	col.Values = append(col.Values, value)
}

func (col *Column) SetNoDataValue() {
	col.SetValue(nil)
}

func (col Columns) SetNoDataValue() {
	for idx, column := range col {
		column.SetNoDataValue()

		col[idx] = column
	}
}
