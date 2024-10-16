package tablebuilder

import (
	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainmetrics "github.com/bartosian/suimon/internal/core/domain/metrics"
	"github.com/bartosian/suimon/internal/core/domain/service/tablebuilder/tables"
)

// handleValidatorParamsTable handles the configuration for the Validator Counts table.
func (tb *Builder) handleValidatorParamsTable(systemState *domainmetrics.SuiSystemState) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeValidatorParams)

	columnValues, err := tables.GetValidatorParamsColumnValues(systemState)
	if err != nil {
		return err
	}

	tableConfig.Columns.SetColumnValues(columnValues)

	tableConfig.RowsCount++

	tb.config = tableConfig

	return nil
}
