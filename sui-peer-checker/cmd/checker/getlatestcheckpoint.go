package checker

import (
	"context"
	"strconv"
	"time"

	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/enums"
)

func (host *RPCHost) GetLatestCheckpoint() {
	if result := getFromRPC(host.rpcClient, enums.RPCMethodGetLatestCheckpointSequenceNumber); result != nil {
		latestCheckpoint := strconv.Itoa(*result)

		host.Metrics.SetValue(enums.MetricTypeLatestCheckpoint, latestCheckpoint)
	}

	return
}

func getFromRPC(rpcClient jsonrpc.RPCClient, method enums.RPCMethod) *int {
	respChan := make(chan *int)
	timeout := time.After(time.Second * 2)

	go func() {
		var response *int

		if err := rpcClient.CallFor(context.Background(), &response, method.String()); err != nil {
			return
		}

		respChan <- response
	}()

	select {
	case response := <-respChan:
		return response
	case <-timeout:
		return nil
	}
}
