package tablebuilder

import (
	"fmt"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainmetrics "github.com/bartosian/suimon/internal/core/domain/metrics"
	"github.com/bartosian/suimon/internal/core/domain/service/tablebuilder/tables"
)

// handleValidatorReportsTable handles the configuration for the Validator Reports table.
// It takes the system state, extracts the necessary data, and updates the table configuration.
func (tb *Builder) handleValidatorReportsTable(systemState *domainmetrics.SuiSystemState) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeValidatorReports)

	validatorReports := systemState.ValidatorReportsParsed

	for _, report := range validatorReports {
		for j, reporter := range report.Reporters {
			reportedName := report.Name
			slashingPct := fmt.Sprintf("%.2f", report.SlashingPercentage)

			if j > 0 {
				reportedName = " "
				slashingPct = " "
			}

			columnValues := tables.GetValidatorReportColumnValues(reportedName, slashingPct, reporter)

			tableConfig.Columns.SetColumnValues(columnValues)

			tableConfig.RowsCount++
		}
	}

	tb.config = tableConfig

	return nil
}
