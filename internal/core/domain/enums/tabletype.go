package enums

type TableType string

const (
	TableTypeRPC                TableType = "📡 PUBLIC RPC"
	TableTypeNode               TableType = "💻 FULL NODES"
	TableTypeValidator          TableType = "🤖 VALIDATORS"
	TableTypePeers              TableType = "🤝 PEERS"
	TableTypeGasPriceAndSubsidy TableType = "💰 EPOCH, GAS PRICE AND SUBSIDY"
	TableTypeValidatorsParams   TableType = "📊 VALIDATORS PARAMS"
	TableTypeValidatorsAtRisk   TableType = "🚨 VALIDATORS AT RISK"
	TableTypeValidatorReports   TableType = "📢 VALIDATORS REPORTS"
	TableTypeActiveValidators   TableType = "✅ ACTIVE VALIDATORS"
)

func (e TableType) ToString() string {
	return string(e)
}
