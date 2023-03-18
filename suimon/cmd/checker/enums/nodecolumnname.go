package enums

//go:generate go run github.com/dmarkham/enumer -type=NodeColumnName -json -transform=snake-upper -output=./nodecolumnname.gen.go
type NodeColumnName int

const (
	ColumnNameStatus NodeColumnName = iota
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
