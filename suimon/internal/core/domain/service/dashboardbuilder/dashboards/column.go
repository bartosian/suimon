package dashboards

import (
	"github.com/mum4k/termdash/container/grid"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
)

const (
	TableNoData    = "no data"
	EmptyValue     = ""
	RpcPortDefault = "9000"
)

// ColumnsConfig is a type that maps column names to their respective widths.
type ColumnsConfig map[enums.ColumnName]int

// ColumnValues represents a mapping between column names and their corresponding values.
type ColumnValues map[enums.ColumnName]any

// ColumnOptions represents a mapping between column names and their corresponding configuration options.
type ColumnOptions map[enums.ColumnName]any

// Columns is a type that maps column names to their respective grid elements.
type Columns map[enums.ColumnName]grid.Element

// NewColumnFixed creates a new column element with a fixed width and a list of sub-elements.
// The `width` parameter specifies the width of the column in pixels, and the `elements` parameter
// is a list of sub-elements that will be contained within the column.
func NewColumnFixed(width int, elements ...grid.Element) grid.Element {
	return grid.ColWidthFixed(width, elements...)
}

// NewColumnPct creates a new column element with a width proportional to the total grid width
// and a list of sub-elements. The `width` parameter specifies the percentage of the grid width
// that the column should occupy, and the `elements` parameter is a list of sub-elements that will
// be contained within the column.
func NewColumnPct(width int, elements ...grid.Element) grid.Element {
	return grid.ColWidthPerc(width, elements...)
}
