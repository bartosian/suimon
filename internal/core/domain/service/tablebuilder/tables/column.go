package tables

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/suimon/internal/core/domain/enums"
)

const (
	TableNoData    = "no data"
	EmptyValue     = ""
	RPCPortDefault = "9000"
)

type (
	// ColumnsConfig represents a mapping between column names and their corresponding column configuration.
	ColumnsConfig map[enums.ColumnName]Column

	// ColumnValues represents a mapping between column names and their corresponding values.
	ColumnValues map[enums.ColumnName]any

	Column struct {
		Config *table.ColumnConfig
		Values []any
	}
)

// NewDefaultColumnConfig creates a new column configuration with the given alignment settings and hidden flag.
func NewDefaultColumnConfig(alignHeader, align text.Align, hidden bool) Column {
	return Column{
		Config: &table.ColumnConfig{
			Align:        align,
			AlignHeader:  alignHeader,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       hidden,
		},
	}
}

// SetValue sets the value of the column to the given value, appending it to the existing values slice.
func (col *Column) SetValue(value any) {
	if value == nil || value == "" {
		value = TableNoData
	}

	col.Values = append(col.Values, value)
}

// SetNoDataValue creates a new copy of the column with a nil value set.
func (col *Column) SetNoDataValue() *Column {
	newColumn := *col
	newColumn.SetValue(nil)

	return &newColumn
}

// SetNoDataValue creates a new slice of columns with a nil value set for each column.
func (cols ColumnsConfig) SetNoDataValue() ColumnsConfig {
	newCols := make(ColumnsConfig, len(cols))
	for idx, col := range cols {
		newCols[idx] = *col.SetNoDataValue()
	}

	return newCols
}

// SetColumnValues creates a new slice of columns with the given values set for each column at the corresponding index.
func (cols ColumnsConfig) SetColumnValues(values ColumnValues) {
	for idx, col := range cols {
		newCol := col

		if value, ok := values[idx]; ok {
			newCol.SetValue(value)
		}

		cols[idx] = newCol
	}
}
