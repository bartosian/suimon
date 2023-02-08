package checker

import (
	"strconv"

	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/enums"
)

func (peer *Peer) GetTotalTransactionNumber() {
	if result := getFromRPC(peer.rpcClient, enums.RPCMethodGetTotalTransactionNumber); result != nil {
		totalTransactionNumber := strconv.Itoa(*result)

		peer.Metrics.SetValue(enums.MetricTypeTotalTransactionsNumber, totalTransactionNumber)
	}

	return
}

func (host *RPCHost) GetTotalTransactionNumber() {
	if result := getFromRPC(host.rpcClient, enums.RPCMethodGetTotalTransactionNumber); result != nil {
		totalTransactionNumber := strconv.Itoa(*result)

		host.Metrics.SetValue(enums.MetricTypeTotalTransactionsNumber, totalTransactionNumber)
	}

	return
}
