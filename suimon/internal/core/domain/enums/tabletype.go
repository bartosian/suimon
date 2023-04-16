package enums

type TableType string

const (
	TableTypeRPC                TableType = "PUBLIC RPC"
	TableTypeNode               TableType = "FULL NODES"
	TableTypeValidator          TableType = "VALIDATORS"
	TableTypePeers              TableType = "PEERS"
	TableTypeGasPriceAndSubsidy TableType = "GAS PRICE AND SUBSIDY"
	TableTypeValidatorsCounts   TableType = "VALIDATORS STATISTICS"
	TableTypeValidatorsAtRisk   TableType = "VALIDATORS AT RISK"
	TableTypeValidatorReports   TableType = "VALIDATORS REPORTS"
	TableTypeActiveValidators   TableType = "ACTIVE VALIDATORS"
)

func (e TableType) ToString() string {
	return string(e)
}
