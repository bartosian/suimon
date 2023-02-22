package dashboards

import (
	"fmt"

	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
)

const suiEmoji = "💧"

type ColumnConfig struct {
	Name  string
	Width int
}

type RowConfig struct {
	Height  int
	Columns []CellName
}

var (
	DashboardConfigSUI = []container.Option{
		container.Border(linestyle.Light),
		container.BorderColor(cell.ColorGreen),
		container.FocusedColor(cell.ColorGreen),
		container.AlignHorizontal(align.HorizontalCenter),
		container.AlignVertical(align.VerticalMiddle),
		container.BorderTitle(fmt.Sprintf("%s SUIMON: PRESS Q or ESC TO QUIT", suiEmoji)),
		container.Focused(),
	}

	RowConfigSUI = [...]RowConfig{
		0: {Height: 7, Columns: []CellName{CellNameNodeStatus, CellNameNetworkStatus, CellNameTotalTransactions, CellNameLatestCheckpoint, CellNameHighestCheckpoint}},
		1: {Height: 7, Columns: []CellName{CellNameUptime, CellNameTransactionsPerSecond, CellNameCheckpointsPerSecond, CellNameConnectedPeers, CellNameVersion, CellNameCommit}},
		2: {Height: 7, Columns: []CellName{CellNameEpochEnd, CellNameDatabaseSize, CellNameTXSyncProgress, CellNameCheckSyncProgress}},
		3: {Height: 18, Columns: []CellName{CellNameEpoch, CellNameDiskUsage}},
		4: {Height: 18, Columns: []CellName{CellNameCpuUsage, CellNameMemoryUsage}},
		5: {Height: 7, Columns: []CellName{CellNameBytesSent, CellNameBytesReceived}},
	}

	ColumnConfigSUI = [...]ColumnConfig{
		CellNameNodeStatus: {
			Name:  "NODE",
			Width: 8,
		},
		CellNameNetworkStatus: {
			Name:  "NET",
			Width: 8,
		},
		CellNameAddress: {
			Name:  "ADDRESS",
			Width: 8,
		},
		CellNameTransactionsPerSecond: {
			Name:  "TRANSACTIONS PER SECOND",
			Width: 28,
		},
		CellNameCheckpointsPerSecond: {
			Name:  "CHECKPOINTS PER SECOND",
			Width: 26,
		},
		CellNameTotalTransactions: {
			Name:  "TOTAL TRANSACTIONS",
			Width: 70,
		},
		CellNameLatestCheckpoint: {
			Name:  "LATEST CHECKPOINT",
			Width: 70,
		},
		CellNameHighestCheckpoint: {
			Name:  "HIGHEST SYNCED CHECKPOINT",
			Width: 70,
		},
		CellNameConnectedPeers: {
			Name:  "CONNECTED PEERS",
			Width: 25,
		},
		CellNameTXSyncProgress: {
			Name:  "SYNC TRANSACTIONS STATUS",
			Width: 70,
		},
		CellNameCheckSyncProgress: {
			Name:  "SYNC CHECKPOINTS STATUS",
			Width: 70,
		},
		CellNameUptime: {
			Name:  "UPTIME",
			Width: 32,
		},
		CellNameVersion: {
			Name:  "VERSION",
			Width: 45,
		},
		CellNameCommit: {
			Name:  "COMMIT",
			Width: 70,
		},
		CellNameCompany: {
			Name:  "PROVIDER",
			Width: 100,
		},
		CellNameCountry: {
			Name:  "COUNTRY",
			Width: 100,
		},
		CellNameEpoch: {
			Name:  "EPOCH",
			Width: 43,
		},
		CellNameEpochEnd: {
			Name:  "TIME TILL THE END OF EPOCH",
			Width: 43,
		},
		CellNameDiskUsage: {
			Name:  "DISK USAGE",
			Width: 43,
		},
		CellNameDatabaseSize: {
			Name:  "DATABASE SIZE",
			Width: 43,
		},
		CellNameBytesSent: {
			Name:  "NETWORK BYTES SENT",
			Width: 43,
		},
		CellNameBytesReceived: {
			Name:  "NETWORK BYTES RECEIVED",
			Width: 43,
		},
		CellNameMemoryUsage: {
			Name:  "MEMORY USAGE",
			Width: 43,
		},
		CellNameCpuUsage: {
			Name:  "CPU USAGE",
			Width: 43,
		},
	}
)
