package dashboards

//go:generate go run github.com/dmarkham/enumer -type=CellName -json -transform=snake-upper -output=./cellname.gen.go
type CellName int

const (
	CellNameNodeStatus CellName = iota
	CellNameNetworkStatus
	CellNameTransactionsPerSecond
	CellNameCheckpointsPerSecond
	CellNameTotalTransactions
	CellNameLatestCheckpoint
	CellNameHighestCheckpoint
	CellNameConnectedPeers
	CellNameTXSyncProgress
	CellNameCheckSyncProgress
	CellNameUptime
	CellNameVersion
	CellNameCommit
	CellNameCurrentEpoch
	CellNameEpochProgress
	CellNameEpochEnd
	CellNameDiskUsage
	CellNameDatabaseSize
	CellNameBytesSent
	CellNameBytesReceived
	CellNameMemoryUsage
	CellNameCpuUsage
	CellNameNodeLogs
	CellNameButtonQuit
)

var cellNameStringValues = [...]string{
	"NODE",
	"NETWORK",
	"TRANSACTIONS PER SECOND",
	"CHECKPOINTS PER SECOND",
	"TOTAL TRANSACTIONS",
	"LATEST CHECKPOINT",
	"HIGHEST SYNCED CHECKPOINT",
	"CONNECTED PEERS",
	"SYNC TRANSACTIONS STATUS",
	"SYNC CHECKPOINTS STATUS",
	"UPTIME",
	"VERSION",
	"COMMIT",
	"CURRENT EPOCH",
	"EPOCH PROGRESS",
	"TIME TILL THE END OF EPOCH",
	"DISK USAGE",
	"DATABASE SIZE",
	"NETWORK BYTES SENT",
	"NETWORK BYTES RECEIVED",
	"MEMORY USAGE",
	"CPU USAGE",
	"NODE LOGS",
	"QUIT",
}

func (i CellName) CellNameString() string {
	return cellNameStringValues[i]
}
