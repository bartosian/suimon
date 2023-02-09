package enums

type RPCMethod string

const (
	RPCMethodGetTotalTransactionNumber         RPCMethod = "sui_getTotalTransactionNumber"
	RPCMethodGetSuiSystemState                 RPCMethod = "sui_getSuiSystemState"
	RPCMethodGetLatestCheckpointSequenceNumber RPCMethod = "sui_getLatestCheckpointSequenceNumber"
	RPCMethodGetCheckpointSummary              RPCMethod = "sui_getCheckpointSummary"
)

func (e RPCMethod) String() string {
	return string(e)
}
