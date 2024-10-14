package tablebuilder

import (
	"sort"
	"strconv"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainmetrics "github.com/bartosian/suimon/internal/core/domain/metrics"
	"github.com/bartosian/suimon/internal/core/domain/service/tablebuilder/tables"
)

const base = 10

// handleValidatorsAtRiskTable handles the configuration for the Validators At Risk table.
// It takes the system state, extracts the necessary data, and updates the table configuration.
func (tb *Builder) handleValidatorsAtRiskTable(systemState *domainmetrics.SuiSystemState) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeValidatorsAtRisk)

	validatorsAtRisk := systemState.ValidatorsAtRiskParsed

	// Optimized sorting logic
	sort.SliceStable(validatorsAtRisk, func(i, j int) bool {
		leftEpochs, leftErr := strconv.ParseInt(validatorsAtRisk[i].EpochsAtRisk, base, 64)
		rightEpochs, rightErr := strconv.ParseInt(validatorsAtRisk[j].EpochsAtRisk, base, 64)

		if leftErr != nil || rightErr != nil {
			return leftErr == nil
		}

		if leftEpochs != rightEpochs {
			return leftEpochs > rightEpochs
		}

		return validatorsAtRisk[i].Name < validatorsAtRisk[j].Name
	})

	for idx, validator := range validatorsAtRisk {
		columnValues := tables.GetValidatorAtRiskColumnValues(idx, validator)

		tableConfig.Columns.SetColumnValues(columnValues)

		tableConfig.RowsCount++
	}

	tb.config = tableConfig

	return nil
}
