package tables

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// Row represents a row in a table.
type Row struct {
	Values   []any
	Config   table.RowConfig
	IsHeader bool
	IsFooter bool
}

// NewRow creates a new row with the given header flag and length.
func NewRow(isHeader, isFooter bool, length int, autoMerge bool, autoMergeAlign text.Align) Row {
	return Row{
		Values:   make(table.Row, 0, length),
		Config:   table.RowConfig{AutoMerge: autoMerge, AutoMergeAlign: autoMergeAlign},
		IsHeader: isHeader,
		IsFooter: isFooter,
	}
}

// AppendValue appends a value to the end of the row.
func (row *Row) AppendValue(value any) {
	row.Values = append(row.Values, value)
}

// PrependValue prepends a value to the beginning of the row.
func (row *Row) PrependValue(value any) {
	row.Values = append(table.Row{value}, row.Values...)
}
