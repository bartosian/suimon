package tablebuilder

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainmetrics "github.com/bartosian/suimon/internal/core/domain/metrics"
	"github.com/bartosian/suimon/internal/core/domain/service/tablebuilder/tables"
)

// handleActiveValidatorsTable handles the configuration for the Active Validators table.
// It takes the system state, extracts the necessary data, and updates the table configuration.
func (tb *Builder) handleActiveValidatorsTable(metrics *domainmetrics.Metrics) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeActiveValidators)

	activeValidators := metrics.SystemState.ActiveValidators
	validatorsApy := metrics.ValidatorsApyParsed

	const base = 10

	sort.SliceStable(activeValidators, func(i, j int) bool {
		leftVotingPower, leftErr := strconv.ParseInt(activeValidators[i].VotingPower, base, 64)
		rightVotingPower, rightErr := strconv.ParseInt(activeValidators[j].VotingPower, base, 64)

		if leftErr != nil {
			return false
		}

		if rightErr != nil {
			return true
		}

		leftNextEpochStake, leftStakeErr := strconv.ParseInt(activeValidators[i].NextEpochStake, base, 64)
		rightNextEpochStake, rightStakeErr := strconv.ParseInt(activeValidators[j].NextEpochStake, base, 64)

		if leftStakeErr != nil {
			return false
		}

		if rightStakeErr != nil {
			return true
		}

		if leftVotingPower != rightVotingPower {
			return leftVotingPower > rightVotingPower
		}

		if leftNextEpochStake != rightNextEpochStake {
			return leftNextEpochStake > rightNextEpochStake
		}

		return activeValidators[i].Name < activeValidators[j].Name
	})

	for idx, validator := range activeValidators {
		validatorApy, ok := validatorsApy[validator.SuiAddress]
		if !ok {
			return fmt.Errorf("failed to lookup validator APY by address: %s", validator.SuiAddress)
		}

		validator.APY = strconv.FormatFloat(validatorApy*100, 'f', 3, 64)

		columnValues, err := tables.GetActiveValidatorColumnValues(idx, validator)
		if err != nil {
			return err
		}

		tableConfig.Columns.SetColumnValues(columnValues)

		tableConfig.RowsCount++
	}

	tb.config = tableConfig

	return nil
}
