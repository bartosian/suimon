package enums

type RPCMethod string

const (
	RPCMethodGetTotalTransactionNumber         RPCMethod = "sui_getTotalTransactionBlocks"
	RPCMethodGetSuiSystemState                 RPCMethod = "suix_getLatestSuiSystemState"
	RPCMethodGetLatestCheckpointSequenceNumber RPCMethod = "sui_getLatestCheckpointSequenceNumber"
	RPCMethodGetCheckpointSummary              RPCMethod = "sui_getCheckpointSummary"
)

func (e RPCMethod) String() string {
	return string(e)
}
