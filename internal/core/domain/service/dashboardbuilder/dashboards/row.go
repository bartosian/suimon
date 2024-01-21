package dashboards

import (
	"github.com/mum4k/termdash/container/grid"

	"github.com/bartosian/suimon/internal/core/domain/enums"
)

// RowsConfig is a type that represents the configuration for a set of rows in a grid.
// Each `RowConfig` contains a height and a list of column names that should be included
// in the row.
type RowsConfig []RowConfig

// RowConfig is a type that represents the configuration for a single row in a grid.
// It contains a height and a list of column names that should be included in the row.
type RowConfig struct {
	Columns []enums.ColumnName
	Height  int
}

// Rows is a type that represents a set of grid rows, each of which is an element in the grid.
type Rows []grid.Element

// NewRowFixed creates a new row element with a fixed height and a list of sub-elements.
// The `height` parameter specifies the height of the row in pixels, and the `elements` parameter
// is a list of sub-elements that will be contained within the row.
func NewRowFixed(height int, elements ...grid.Element) grid.Element {
	return grid.RowHeightFixed(height, elements...)
}

// NewRowPct creates a new row element with a height proportional to the total grid height
// and a list of sub-elements. The `height` parameter specifies the percentage of the grid height
// that the row should occupy, and the `elements` parameter is a list of sub-elements that will
// be contained within the row.
func NewRowPct(height int, elements ...grid.Element) grid.Element {
	return grid.RowHeightPerc(height, elements...)
}
