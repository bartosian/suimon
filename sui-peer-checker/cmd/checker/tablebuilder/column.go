package tablebuilder

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

const tableNoData = "no data"

type (
	Column struct {
		Values []any
		Config table.ColumnConfig
	}
)

func (col *Column) SetValue(value any) {
	if value == nil || value == "" {
		value = tableNoData
	}

	col.Values = append(col.Values, value)
}
