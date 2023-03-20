package enums

//go:generate go run github.com/dmarkham/enumer -type=NodeColumnName -json -transform=snake-upper -output=./nodecolumnname.gen.go
type NodeColumnName int

const (
	ColumnNameHealth NodeColumnName = iota
	ColumnNameAddress
	ColumnNamePortRPC
	ColumnNameTotalTransactions
	ColumnNameTotalTransactionCertificates
	ColumnNameTotalTransactionEffects
	ColumnNameLatestCheckpoint
	ColumnNameHighestKnownCheckpoint
	ColumnNameHighestSyncedCheckpoint
	ColumnNameCurrentEpoch
	ColumnNameTXSyncProgress
	ColumnNameCheckSyncProgress
	ColumnNameNetworkPeers
	ColumnNameUptime
	ColumnNameVersion
	ColumnNameCommit
	ColumnNameCompany
	ColumnNameCountry
)
