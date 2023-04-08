package tables

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/tablebuilder"
)

var (
	SortConfigRPC = tablebuilder.SortConfig{
		{Name: enums.ColumnNameHealth.ToString(), Mode: table.Dsc},
		{Name: enums.ColumnNameTotalTransactionBlocks.ToString(), Mode: table.Dsc},
		{Name: enums.ColumnNameLatestCheckpoint.ToString(), Mode: table.Dsc},
	}
	ColumnsConfigRPC = tablebuilder.ColumnsConfig{
		enums.ColumnNameIndex:                  tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHealth:                 tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameAddress:                tablebuilder.NewDefaultColumnConfig(text.AlignLeft, text.AlignCenter, false),
		enums.ColumnNamePortRPC:                tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionBlocks: tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameLatestCheckpoint:       tablebuilder.NewDefaultColumnConfig(text.AlignLeft, text.AlignLeft, false),
	}
	RowsConfigRPC = tablebuilder.RowsConfig{
		0: {
			enums.ColumnNameHealth,
			enums.ColumnNameAddress,
			enums.ColumnNamePortRPC,
			enums.ColumnNameTotalTransactionBlocks,
			enums.ColumnNameLatestCheckpoint,
		},
	}
)

// GetRPCColumnValues returns a map of NodeColumnName values to corresponding values for the RPC service on the specified host.
// The function retrieves information about the RPC service from the host's internal state and formats it into a map of NodeColumnName keys and corresponding values.
// Returns a map of NodeColumnName keys to corresponding values.
func GetRPCColumnValues(idx int, host host.Host) tablebuilder.ColumnValues {
	status := host.Status.StatusToPlaceholder()
	port := host.Ports[enums.PortTypeRPC]
	if port == "" {
		port = tablebuilder.RpcPortDefault
	}
	address := host.HostPort.Address

	return tablebuilder.ColumnValues{
		enums.ColumnNameIndex:                  idx + 1,
		enums.ColumnNameHealth:                 status,
		enums.ColumnNameAddress:                address,
		enums.ColumnNamePortRPC:                port,
		enums.ColumnNameTotalTransactionBlocks: host.Metrics.TotalTransactionsBlocks,
		enums.ColumnNameLatestCheckpoint:       host.Metrics.LatestCheckpoint,
	}
}
