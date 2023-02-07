package checker

import (
	"context"
	"strconv"

	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/enums"
)

func (checker *Checker) GetTotalTransactionNumber() (*int, error) {
	return getTotalTransactionNumber(checker.rpcClient)
}

func (peer *Peer) GetTotalTransactionNumber() {
	result, err := getTotalTransactionNumber(peer.rpcClient)
	if err != nil {
		return
	}

	totalTransactionNumber := strconv.Itoa(*result)

	peer.Metrics = Metrics{
		TotalTransactionNumber: totalTransactionNumber,
	}
}

func getTotalTransactionNumber(rpcClient jsonrpc.RPCClient) (*int, error) {
	var response *int

	err := rpcClient.CallFor(context.Background(), &response, enums.RPCMethodGetTotalTransactionNumber.String())
	if err != nil {
		return nil, err
	}

	return response, nil
}
