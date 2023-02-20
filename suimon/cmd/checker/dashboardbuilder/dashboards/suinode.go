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
		2: {Height: 7, Columns: []CellName{CellNameTXSyncProgress, CellNameCheckSyncProgress}},
		3: {Height: 20, Columns: []CellName{CellNameEpoch}},
		4: {Height: 10, Columns: []CellName{CellNameEpochEnd}},
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
			Width: 28,
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
			Width: 23,
		},
		CellNameTXSyncProgress: {
			Name:  "SYNC TRANSACTIONS STATUS",
			Width: 113,
		},
		CellNameCheckSyncProgress: {
			Name:  "SYNC CHECKPOINTS STATUS",
			Width: 113,
		},
		CellNameUptime: {
			Name:  "UPTIME",
			Width: 32,
		},
		CellNameVersion: {
			Name:  "VERSION",
			Width: 49,
		},
		CellNameCommit: {
			Name:  "COMMIT",
			Width: 66,
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
			Width: 50,
		},
		CellNameEpochEnd: {
			Name:  "TIME TILL THE END OF EPOCH",
			Width: 50,
		},
	}
)
