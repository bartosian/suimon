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
		0: {Height: 7, Columns: []CellName{CellNameStatus, CellNameTransactionsPerSecond, CellNameUptime, CellNameConnectedPeers, CellNameVersion, CellNameCommit}},
		1: {Height: 7, Columns: []CellName{CellNameTotalTransactions, CellNameLatestCheckpoint, CellNameHighestCheckpoint}},
		2: {Height: 7, Columns: []CellName{CellNameTXSyncProgress, CellNameCheckSyncProgress}},
		3: {Height: 20, Columns: []CellName{CellNameEpoch}},
	}

	ColumnConfigSUI = [...]ColumnConfig{
		CellNameStatus: {
			Name:  "STATUS",
			Width: 8,
		},
		CellNameAddress: {
			Name:  "ADDRESS",
			Width: 8,
		},
		CellNameTransactionsPerSecond: {
			Name:  "TPS",
			Width: 30,
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
			Width: 105,
		},
		CellNameCheckSyncProgress: {
			Name:  "SYNC CHECKPOINTS STATUS",
			Width: 105,
		},
		CellNameUptime: {
			Name:  "UPTIME",
			Width: 30,
		},
		CellNameVersion: {
			Name:  "VERSION",
			Width: 45,
		},
		CellNameCommit: {
			Name:  "COMMIT",
			Width: 72,
		},
		CellNameCompany: {
			Name:  "PROVIDER",
			Width: 100,
		},
		CellNameCountry: {
			Name:  "PROVIDER",
			Width: 100,
		},
		CellNameEpoch: {
			Name:  "EPOCH",
			Width: 45,
		},
	}
)
