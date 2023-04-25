package enums

type TableType string

const (
	TableTypeRPC                TableType = "📡 PUBLIC RPC"
	TableTypeNode               TableType = "💻 FULL NODES"
	TableTypeValidator          TableType = "🤖 VALIDATORS"
	TableTypePeers              TableType = "🤝 PEERS"
	TableTypeGasPriceAndSubsidy TableType = "💰 EPOCH, GAS PRICE AND SUBSIDY"
	TableTypeValidatorsCounts   TableType = "📊 VALIDATORS COUNTS AND THRESHOLDS"
	TableTypeValidatorsAtRisk   TableType = "🚨 VALIDATORS AT RISK"
	TableTypeValidatorReports   TableType = "📢 VALIDATORS REPORTS"
	TableTypeActiveValidators   TableType = "✅ ACTIVE VALIDATORS"
)

func (e TableType) ToString() string {
	return string(e)
}
