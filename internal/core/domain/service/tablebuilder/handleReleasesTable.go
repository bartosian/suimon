package tablebuilder

import (
	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/core/domain/metrics"
	"github.com/bartosian/suimon/internal/core/domain/service/tablebuilder/tables"
)

// handleReleasesTable processes the provided releases and builds a table configuration for them.
// It iterates over each release, retrieves the column values for it, and sets these values in the table configuration.
// The function also increments the row count for each processed release.
// If an error occurs while getting the column values, the function returns the error.
// At the end, the built table configuration is set as the builder's configuration.
// The function returns nil if it completes successfully.
func (tb *Builder) handleReleasesTable(releases []metrics.Release) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeReleases)

	for idx := range releases {
		release := releases[idx]

		columnValues := tables.GetReleaseColumnValues(idx, &release)

		tableConfig.Columns.SetColumnValues(columnValues)
		tableConfig.RowsCount++
	}

	tb.config = tableConfig

	return nil
}
