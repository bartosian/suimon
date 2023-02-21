package tables

//go:generate go run github.com/dmarkham/enumer -type=ColumnName -json -transform=snake-upper -output=./columnname.gen.go
type ColumnName int

const (
	ColumnNameStatus ColumnName = iota
	ColumnNameAddress
	ColumnNamePortRPC
	ColumnNameTotalTransactions
	ColumnNameLatestCheckpoint
	ColumnNameHighestCheckpoint
	ColumnNameTXSyncProgress
	ColumnNameCheckSyncProgress
	ColumnNameConnectedPeers
	ColumnNameUptime
	ColumnNameVersion
	ColumnNameCommit
	ColumnNameCompany
	ColumnNameCountry
)
