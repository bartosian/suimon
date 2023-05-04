package dashboards

import (
	"github.com/mum4k/termdash/cell"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/core/domain/host"
)

var (
	ColumnsConfigRPC = ColumnsConfig{
		// Overview section
		enums.ColumnNameCurrentEpoch:            19,
		enums.ColumnNameSystemTimeTillNextEpoch: 19,
		enums.ColumnNameTotalTransactionBlocks:  30,
		enums.ColumnNameLatestCheckpoint:        30,
	}

	RowsConfigRPC = RowsConfig{
		0: {
			Height: 14,
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
func GetRPCColumnValues(host host.Host) ColumnValues {
	return ColumnValues{
		enums.ColumnNameTotalTransactionBlocks:  host.Metrics.TotalTransactionsBlocks,
		enums.ColumnNameLatestCheckpoint:        host.Metrics.LatestCheckpoint,
		enums.ColumnNameCurrentEpoch:            host.Metrics.SystemState.Epoch,
		enums.ColumnNameSystemTimeTillNextEpoch: host.Metrics.DurationTillEpochEndHHMM,
	}
}
