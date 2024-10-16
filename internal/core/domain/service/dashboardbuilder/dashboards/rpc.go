package dashboards

import (
	"github.com/mum4k/termdash/cell"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/suimon/internal/core/domain/host"
)

var (
	ColumnsConfigRPC = ColumnsConfig{
		enums.ColumnNameCurrentEpoch:            ColumnWidth19,
		enums.ColumnNameSystemTimeTillNextEpoch: ColumnWidth19,
		enums.ColumnNameTotalTransactionBlocks:  ColumnWidth30,
		enums.ColumnNameLatestCheckpoint:        ColumnWidth30,
	}

	RowsConfigRPC = RowsConfig{
		0: {
			Height: RowHeight14,
			Columns: []enums.ColumnName{
				enums.ColumnNameCurrentEpoch,
				enums.ColumnNameSystemTimeTillNextEpoch,
				enums.ColumnNameTotalTransactionBlocks,
				enums.ColumnNameLatestCheckpoint,
			},
		},
	}

	CellsConfigRPC = CellsConfig{
		enums.ColumnNameCurrentEpoch:            {"CURRENT EPOCH", cell.ColorGreen},
		enums.ColumnNameSystemTimeTillNextEpoch: {"TIME TILL NEXT EPOCH", cell.ColorGreen},
		enums.ColumnNameTotalTransactionBlocks:  {"TOTAL TRANSACTION BLOCKS", cell.ColorYellow},
		enums.ColumnNameLatestCheckpoint:        {"LATEST CHECKPOINT", cell.ColorBlue},
	}
)

// GetRPCColumnValues returns a map of ColumnName values to corresponding values for a node at the specified index on the specified host.
// The function retrieves information about the node from the host's internal state and formats it into a map of NodeColumnName keys and corresponding values.
// The function also includes emoji values in the map if the specified flag is true.
func GetRPCColumnValues(host *domainhost.Host) (ColumnValues, error) {
	return ColumnValues{
		enums.ColumnNameTotalTransactionBlocks:  host.Metrics.TotalTransactionsBlocks,
		enums.ColumnNameLatestCheckpoint:        host.Metrics.LatestCheckpoint,
		enums.ColumnNameCurrentEpoch:            host.Metrics.SystemState.Epoch,
		enums.ColumnNameSystemTimeTillNextEpoch: host.Metrics.DurationTillEpochEndHHMM,
	}, nil
}
