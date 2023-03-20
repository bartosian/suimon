package enums

//go:generate go run github.com/dmarkham/enumer -type=ColumnName -json -transform=snake-upper -output=./columnname.gen.go
type ColumnName int

const (
	ColumnNameHealth ColumnName = iota
	ColumnNameAddress
	ColumnNamePortRPC
	ColumnNameTotalTransactions
	ColumnNameLatestCheckpoint
	ColumnNameTotalTransactionCertificates
	ColumnNameTotalTransactionEffects
	ColumnNameHighestKnownCheckpoint
	ColumnNameHighestSyncedCheckpoint
	ColumnNameLastExecutedCheckpoint
	ColumnNameCheckpointExecBacklog
	ColumnNameCheckpointSyncBacklog
	ColumnNameCurrentEpoch
	ColumnNameTXSyncPercentage
	ColumnNameCheckSyncPercentage
	ColumnNameNetworkPeers
	ColumnNameUptime
	ColumnNameVersion
	ColumnNameCommit
	ColumnNameCountry
)
