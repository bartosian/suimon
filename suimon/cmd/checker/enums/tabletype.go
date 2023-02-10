package enums

type TableType string

const (
	TableTypeRPC   TableType = "REMOTE RPC"
	TableTypeNode  TableType = "LOCAL NODE"
	TableTypePeers TableType = "PEERS"
)
