package enums

type TableType string

const (
	TableTypeRPC              TableType = "PUBLIC RPC"
	TableTypeNode             TableType = "YOUR NODE"
	TableTypeValidator        TableType = "YOUR VALIDATOR"
	TableTypePeers            TableType = "PEERS"
	TableTypeSystemState      TableType = "SYSTEM STATE"
	TableTypeValidatorsCounts TableType = "VALIDATORS COUNTS"
	TableTypeValidatorsAtRisk TableType = "VALIDATORS AT RISK"
	TableTypeValidatorReports TableType = "VALIDATOR REPORTS"
	TableTypeActiveValidators TableType = "ACTIVE VALIDATORS"
)
