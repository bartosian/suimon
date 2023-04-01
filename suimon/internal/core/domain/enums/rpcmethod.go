package enums

type RPCMethod string

const (
	RPCMethodGetTotalTransactionBlocks         RPCMethod = "sui_getTotalTransactionBlocks"
	RPCMethodGetSuiSystemState                 RPCMethod = "suix_getLatestSuiSystemState"
	RPCMethodGetLatestCheckpointSequenceNumber RPCMethod = "sui_getLatestCheckpointSequenceNumber"
)

func (e RPCMethod) String() string {
	return string(e)
}
