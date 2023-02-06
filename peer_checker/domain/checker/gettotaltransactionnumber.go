package checker

import (
	"context"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/peer_checker/domain/enums"
)

func (checker *Checker) GetTotalTransactionNumber() (*uint64, error) {
	return getTotalTransactionNumber(checker.rpcClient)
}

func (peer *Peer) GetTotalTransactionNumber() {
	result, err := getTotalTransactionNumber(peer.rpcClient)
	if err == nil {
		peer.TotalTransactionNumber = result
	}
}

func getTotalTransactionNumber(rpcClient jsonrpc.RPCClient) (*uint64, error) {
	var response *uint64

	err := rpcClient.CallFor(context.Background(), &response, enums.RPCMethodGetTotalTransactionNumber.String())
	if err != nil {
		return nil, err
	}

	return response, nil
}
