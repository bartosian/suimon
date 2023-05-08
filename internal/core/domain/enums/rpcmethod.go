package enums

type RPCMethod string

const (
	RPCMethodGetTotalTransactionBlocks         RPCMethod = "sui_getTotalTransactionBlocks"
	RPCMethodGetSuiSystemState                 RPCMethod = "suix_getLatestSuiSystemState"
	RPCMethodGetLatestCheckpointSequenceNumber RPCMethod = "sui_getLatestCheckpointSequenceNumber"
	RPCMethodGetValidatorsApy                  RPCMethod = "suix_getValidatorsApy"
	RPCMethodGetEpochs                         RPCMethod = "suix_getEpochs"
)

func (e RPCMethod) String() string {
	return string(e)
}
