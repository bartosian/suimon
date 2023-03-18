package enums

//go:generate go run github.com/dmarkham/enumer -type=ValidatorColumnName -json -transform=snake-upper -output=./validatorcolumnname.gen.go
type ValidatorColumnName int

const (
	ValidatorColumnNameName ValidatorColumnName = iota
	ValidatorColumnNameAddress
	ValidatorColumnNameVotingPower
	ValidatorColumnNameGasPrice
	ValidatorColumnNameStakingPoolSuiBalance
	ValidatorColumnNameCommissionRate
	ValidatorColumnNameNextEpochGasPrice
	ValidatorColumnNameNextEpochStake
	ValidatorColumnNameNextEpochCommissionRate
)
