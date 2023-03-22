package enums

type TableType string

const (
	TableTypeRPC              TableType = "PUBLIC RPC"
	TableTypeNode             TableType = "YOUR NODE"
	TableTypeValidator        TableType = "YOUR VALIDATOR"
	TableTypePeers            TableType = "PEERS"
	TableTypeSystemState      TableType = "SYSTEM STATE"
	TableTypeActiveValidators TableType = "ACTIVE VALIDATORS"
)
