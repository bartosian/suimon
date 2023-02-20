package dashboards

//go:generate go run github.com/dmarkham/enumer -type=CellName -json -transform=snake-upper -output=./cellname.gen.go
type CellName int

const (
	CellNameStatus CellName = iota
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
)

var cellNameStringValues = [...]string{
	"STATUS",
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
}

func (i CellName) CellNameString() string {
	return cellNameStringValues[i]
}
