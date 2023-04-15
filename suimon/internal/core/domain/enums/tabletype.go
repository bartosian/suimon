package enums

type TableType string

const (
	TableTypeRPC              TableType = "PUBLIC RPC"
	TableTypeNode             TableType = "FULL NODES"
	TableTypeValidator        TableType = "VALIDATORS"
	TableTypePeers            TableType = "PEERS"
	TableTypeSystemState      TableType = "SYSTEM STATE"
	TableTypeValidatorsCounts TableType = "VALIDATORS STATISTICS"
	TableTypeValidatorsAtRisk TableType = "VALIDATORS AT RISK"
	TableTypeValidatorReports TableType = "VALIDATORS REPORTS"
	TableTypeActiveValidators TableType = "ACTIVE VALIDATORS"
)

func (e TableType) ToString() string {
	return string(e)
}
