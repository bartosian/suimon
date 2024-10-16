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

type NewRowConfig struct {
	IsHeader       bool
	IsFooter       bool
	Length         int
	AutoMerge      bool
	AutoMergeAlign text.Align
}

// NewRow creates a new row based on the provided configuration.
func NewRow(config NewRowConfig) Row {
	return Row{
		Values:   make(table.Row, 0, config.Length),
		Config:   table.RowConfig{AutoMerge: config.AutoMerge, AutoMergeAlign: config.AutoMergeAlign},
		IsHeader: config.IsHeader,
		IsFooter: config.IsFooter,
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
