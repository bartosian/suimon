package columnnames

//go:generate go run github.com/dmarkham/enumer -type=SystemColumnName -json -transform=snake-upper -output=./systemcolumnname.gen.go
type SystemColumnName int

const (
	SystemColumnNameStorageFund SystemColumnName = iota
	SystemColumnNameReferenceGasPrice
	SystemColumnNameEpochDurationMs
	SystemColumnNameStakeSubsidyCounter
	SystemColumnNameStakeSubsidyBalance
	SystemColumnNameStakeSubsidyCurrentEpochAmount
	SystemColumnNameTotalStake
	SystemColumnNameValidatorsCount
	SystemColumnNameValidatorsAtRiskCount
)
