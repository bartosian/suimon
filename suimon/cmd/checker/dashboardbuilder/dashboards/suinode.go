package dashboards

//go:generate go run github.com/dmarkham/enumer -type=CellName -json -transform=snake-upper -output=./cellname.gen.go
type CellName int

const (
	CellNameStatus CellName = iota
	CellNameAddress
	CellNameTransactionsPerSecond
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
)

var cellNameStringValues = [...]string{
	"STATUS",
	"ADDRESS",
	"TPS",
	"TOTAL TX",
	"LATEST CHECK",
	"SYNCED CHECK",
	"PEERS",
	"SYNC TX PROGRESS",
	"SYNC CHECK PROGRESS",
	"UPTIME",
	"VERSION",
	"COMMIT",
	"PROVIDER",
	"COUNTRY",
}

func (i CellName) CellNameString() string {
	return cellNameStringValues[i]
}
