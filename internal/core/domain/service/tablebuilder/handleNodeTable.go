package tablebuilder

import (
	"sort"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/suimon/internal/core/domain/host"
	"github.com/bartosian/suimon/internal/core/domain/service/tablebuilder/tables"
)

// handleNodeTable handles the configuration for the Node table.
func (tb *Builder) handleNodeTable(hosts []domainhost.Host) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeNode)

	sort.SliceStable(hosts, func(i, j int) bool {
		left, right := hosts[i], hosts[j]
		if left.Status != right.Status {
			return left.Status > right.Status
		}
		if left.Metrics.TotalTransactionsBlocks != right.Metrics.TotalTransactionsBlocks {
			return left.Metrics.TotalTransactionsBlocks > right.Metrics.TotalTransactionsBlocks
		}
		return left.Metrics.HighestSyncedCheckpoint > right.Metrics.HighestSyncedCheckpoint
	})

	for idx, host := range hosts {
		if !host.Metrics.Updated {
			continue
		}

		columnValues := tables.GetNodeColumnValues(idx, host)

		tableConfig.Columns.SetColumnValues(columnValues)
		tableConfig.RowsCount++
	}

	tb.config = tableConfig

	return nil
}
