package enums

type TableType string

const (
	TableTypeRPC                TableType = "📡 PUBLIC RPC"
	TableTypeNode               TableType = "💻 FULL NODES"
	TableTypeValidator          TableType = "🤖 VALIDATORS"
	TableTypeGasPriceAndSubsidy TableType = "💾 SYSTEM STATE"
	TableTypeProtocol           TableType = "🌐 PROTOCOL"
	TableTypeValidatorsParams   TableType = "📊 VALIDATORS PARAMS"
	TableTypeValidatorsAtRisk   TableType = "🚨 VALIDATORS AT RISK"
	TableTypeValidatorReports   TableType = "📢 VALIDATORS REPORTS"
	TableTypeActiveValidators   TableType = "✅ ACTIVE VALIDATORS"
	TableTypeReleases           TableType = "📈 RELEASE HISTORY"
)

func (e TableType) ToString() string {
	return string(e)
}
