package checker

import (
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/pkg/log"
)

type RPCList struct {
	Testnet []string `yaml:"testnet"`
	Devnet  []string `yaml:"devnet"`
}

func (rpc RPCList) GetByNetwork(network enums.NetworkType) []string {
	switch network {
	case enums.NetworkTypeTestnet:
		return rpc.Testnet
	case enums.NetworkTypeDevnet:
		fallthrough
	default:
		return rpc.Devnet
	}
}

type RPCHost struct {
	Address     string
	AddressType enums.AddressType

	Metrics Metrics
	Status  enums.Status

	rpcClient jsonrpc.RPCClient
	logger    log.Logger
}

func newRPCHost(address string) *RPCHost {
	host := &RPCHost{
		Address: address,
		logger:  log.NewLogger(),
	}

	host.rpcClient = jsonrpc.NewClient(address)

	return host
}

func (host *RPCHost) SetStatus() {
	metrics := host.Metrics

	if !metrics.Updated {
		host.Status = enums.StatusRed
	} else if metrics.TotalTransactionNumber == "" && metrics.LatestCheckpoint == "" {
		host.Status = enums.StatusYellow
	} else {
		host.Status = enums.StatusGreen
	}
}
