package tablebuilder

import (
	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainmetrics "github.com/bartosian/suimon/internal/core/domain/metrics"
	"github.com/bartosian/suimon/internal/core/domain/service/tablebuilder/tables"
)

// handleProtocolTable handles the configuration for the Protocol table.
func (tb *Builder) handleProtocolTable(metrics *domainmetrics.Metrics) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeProtocol)

	columnValues, err := tables.GetProtocolColumnValues(metrics)
	if err != nil {
		return err
	}

	tableConfig.Columns.SetColumnValues(columnValues)

	tableConfig.RowsCount++

	tb.config = tableConfig

	return nil
}
