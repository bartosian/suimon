package tablebuilder

import (
	"sort"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/suimon/internal/core/domain/host"
	"github.com/bartosian/suimon/internal/core/domain/service/tablebuilder/tables"
)

// handleRPCTable handles the configuration for the RPC table.
func (tb *Builder) handleRPCTable(hosts []domainhost.Host) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeRPC)

	sort.SliceStable(hosts, func(i, j int) bool {
		left, right := hosts[i], hosts[j]
		if left.Status != right.Status {
			return left.Status > right.Status
		}

		return left.Metrics.TotalTransactionsBlocks > right.Metrics.TotalTransactionsBlocks
	})

	for idx := range hosts {
		host := hosts[idx]

		if !host.Metrics.Updated {
			continue
		}

		columnValues := tables.GetRPCColumnValues(idx, &host)

		tableConfig.Columns.SetColumnValues(columnValues)

		tableConfig.RowsCount++
	}

	tb.config = tableConfig

	return nil
}
