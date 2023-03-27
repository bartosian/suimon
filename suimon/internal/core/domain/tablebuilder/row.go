package tablebuilder

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

type Row struct {
	Values   table.Row
	Config   table.RowConfig
	IsHeader bool
}

func NewRow(isHeader bool, length int) Row {
	return Row{
		Values:   make(table.Row, 0, length),
		Config:   table.RowConfig{AutoMerge: true},
		IsHeader: isHeader,
	}
}

func (row *Row) SetValue(value any) {
	switch t := value.(type) {
	case string:
		if t == emptyValue {
			row.Values = append(table.Row{emptyValue}, row.Values...)

			return
		}
	}

	row.Values = append(row.Values, value)
}
