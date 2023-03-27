package enums

type CellName int

const (
	CellNameNodeHealth CellName = iota
	CellNameNetworkHealth
	CellNameCurrentEpoch
	CellNameEpochTimeTillTheEnd
	CellNameUptime
	CellNameVersion
	CellNameCommit
	CellNameTotalTransactions
	CellNameTotalTransactionCertificates
	CellNameTotalTransactionEffects
	CellNameLatestCheckpoint
	CellNameHighestKnownCheckpoint
	CellNameLastExecutedCheckpoint
	CellNameHighestSyncedCheckpoint
	CellNameCheckpointSyncBacklog
	CellNameCheckpointExecBacklog
	CellNameTransactionsPerSecond
	CellNameCheckpointsPerSecond
	CellNameConnectedPeers
	CellNameTXSyncProgress
	CellNameCheckSyncProgress
	CellNameEpochProgress
	CellNameDiskUsage
	CellNameDatabaseSize
	CellNameBytesSent
	CellNameBytesReceived
	CellNameMemoryUsage
	CellNameCpuUsage
	CellNameNodeLogs
	CellNameButtonQuit
	CellNameTPSTracker
	CellNameCPSTracker
)

var cellNameStringValues = [...]string{
	"NODE",
	"NETWORK",
	"CURRENT EPOCH",
	"TIME TILL THE END OF EPOCH",
	"UPTIME DAYS",
	"VERSION",
	"COMMIT",
	"TOTAL TRANSACTIONS",
	"TOTAL TRANSACTION CERTIFICATES",
	"TOTAL TRANSACTION EFFECTS",
	"LATEST CHECKPOINT",
	"HIGHEST KNOWN CHECKPOINT",
	"LAST EXECUTED CHECKPOINT",
	"HIGHEST SYNCED CHECKPOINT",
	"CHECKPOINT SYNC BACKLOG",
	"CHECKPOINT EXEC BACKLOG",
	"TRANSACTIONS PER SECOND",
	"CHECKPOINTS PER SECOND",
	"CONNECTED PEERS",
	"SYNC TRANSACTIONS STATUS",
	"SYNC CHECKPOINTS STATUS",
	"EPOCH PROGRESS",
	"DISK USAGE",
	"DATABASE SIZE",
	"NETWORK BYTES SENT",
	"NETWORK BYTES RECEIVED",
	"MEMORY USAGE",
	"CPU USAGE",
	"NODE LOGS",
	"QUIT",
	"TRANSACTIONS PER SECOND TRACKER",
	"CHECKPOINTS PER SECOND TRACKER",
}

func (i CellName) CellNameString() string {
	return cellNameStringValues[i]
}
