package dashboards

//go:generate go run github.com/dmarkham/enumer -type=CellName -json -transform=snake-upper -output=./cellname.gen.go
type CellName int

const (
	CellNameNodeStatus CellName = iota
	CellNameNetworkStatus
	CellNameAddress
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
	CellNameCompany
	CellNameCountry
	CellNameEpoch
	CellNameEpochEnd
	CellNameDiskUsage
	CellNameDatabaseSize
)

var cellNameStringValues = [...]string{
	"NODE STATUS",
	"NETWORK STATUS",
	"ADDRESS",
	"TPS",
	"CPS",
	"TOTAL TRANSACTIONS",
	"LATEST CHECKPOINT",
	"HIGHEST SYNCED CHECKPOINT",
	"CONNECTED PEERS",
	"SYNC TRANSACTIONS STATUS",
	"SYNC CHECKPOINTS STATUS",
	"UPTIME DAYS",
	"VERSION",
	"COMMIT",
	"PROVIDER",
	"COUNTRY",
	"EPOCH",
	"TIME TILL THE END OF EPOCH",
	"DISK USAGE",
	"DATABASE SIZE",
}

func (i CellName) CellNameString() string {
	return cellNameStringValues[i]
}
