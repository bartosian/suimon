package checker

import "github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"

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
