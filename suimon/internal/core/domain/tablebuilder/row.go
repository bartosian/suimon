package tablebuilder

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type Row struct {
	Values   table.Row
	Config   table.RowConfig
	IsHeader bool
}

func NewRow(isHeader bool, length int) Row {
	return Row{
		Values:   make(table.Row, 0, length),
		Config:   table.RowConfig{AutoMerge: true, AutoMergeAlign: text.AlignCenter},
		IsHeader: isHeader,
	}
}

func (row *Row) AppendValue(value any) {
	row.Values = append(row.Values, value)
}

func (row *Row) PrependValue(value any) {
	row.Values = append(table.Row{value}, row.Values...)
}
