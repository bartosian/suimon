package tablebuilder

import (
	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainmetrics "github.com/bartosian/suimon/internal/core/domain/metrics"
	"github.com/bartosian/suimon/internal/core/domain/service/tablebuilder/tables"
)

// handleSystemStateTable handles the configuration for the System State table.
func (tb *Builder) handleSystemStateTable(metrics *domainmetrics.Metrics) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeGasPriceAndSubsidy)

	columnValues, err := tables.GetSystemStateColumnValues(metrics)
	if err != nil {
		return err
	}

	tableConfig.Columns.SetColumnValues(columnValues)

	tableConfig.RowsCount++

	tb.config = tableConfig

	return nil
}
