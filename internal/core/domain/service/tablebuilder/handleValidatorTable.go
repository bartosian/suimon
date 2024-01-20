package tablebuilder

import (
	"sort"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/suimon/internal/core/domain/host"
	"github.com/bartosian/suimon/internal/core/domain/service/tablebuilder/tables"
)

// handleValidatorTable handles the configuration for the Validator table.
func (tb *Builder) handleValidatorTable(hosts []domainhost.Host) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeValidator)

	sort.SliceStable(hosts, func(i, j int) bool {
		left, right := hosts[i], hosts[j]
		if left.Status != right.Status {
			return left.Status > right.Status
		}
		if left.Metrics.CurrentRound != right.Metrics.CurrentRound {
			return left.Metrics.CurrentRound > right.Metrics.CurrentRound
		}
		return left.Metrics.HighestSyncedCheckpoint > right.Metrics.HighestSyncedCheckpoint
	})

	for idx, host := range hosts {
		if !host.Metrics.Updated {
			continue
		}

		columnValues := tables.GetValidatorColumnValues(idx, host)

		tableConfig.Columns.SetColumnValues(columnValues)
		tableConfig.RowsCount++
	}

	tb.config = tableConfig

	return nil
}
