package enums

type TableType string

const (
	TableTypeRPC   TableType = "REMOTE RPC"
	TableTypeNode  TableType = "YOUR NODE"
	TableTypePeers TableType = "PEERS"
)
