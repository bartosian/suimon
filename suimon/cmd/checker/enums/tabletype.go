package enums

type TableType string

const (
	TableTypeRPC         TableType = "PUBLIC RPC"
	TableTypeNode        TableType = "YOUR NODE"
	TableTypePeers       TableType = "PEERS"
	TableTypeSystemState TableType = "SYSTEM STATE"
	TableTypeValidators  TableType = "VALIDATORS"
)
