package tablebuilder

import (
	"github.com/bartosian/suimon/cmd/checker/enums"
	"github.com/jedib0t/go-pretty/v6/table"
)

const tableNoData = "no data"

type (
	Columns []Column

	Column struct {
		Values []any
		Config table.ColumnConfig
	}
)

func (col *Column) SetValue(value any) {
	if value == nil || value == "" {
		value = enums.ColorYellow + tableNoData + enums.ColorReset
	}

	col.Values = append(col.Values, value)
}

func (col *Column) SetNoDataValue() {
	col.SetValue(nil)
}

func (col Columns) SetNoDataValue() {
	for idx := range col {
		col[idx].SetNoDataValue()
	}
}
