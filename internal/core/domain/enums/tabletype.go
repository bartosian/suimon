package enums

type TableType string

const (
	TableTypeRPC                TableType = "📡 REFERENCE RPC"
	TableTypeNode               TableType = "💻 FULL NODES"
	TableTypeValidator          TableType = "🤖 VALIDATORS"
	TableTypeGasPriceAndSubsidy TableType = "💾 SYSTEM STATE"
	TableTypeProtocol           TableType = "🌐 PROTOCOL"
	TableTypeValidatorParams    TableType = "📊 VALIDATOR PARAMS"
	TableTypeValidatorsAtRisk   TableType = "🚨 VALIDATORS AT RISK"
	TableTypeValidatorReports   TableType = "📢 VALIDATOR REPORTS"
	TableTypeActiveValidators   TableType = "✅ ACTIVE VALIDATORS"
	TableTypeReleases           TableType = "📈 RELEASE HISTORY"
)

func (e TableType) ToString() string {
	return string(e)
}
