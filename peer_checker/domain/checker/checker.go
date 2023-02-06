package checker

import (
	"os"

	"github.com/ybbus/jsonrpc/v3"
	"gopkg.in/yaml.v3"

	"github.com/bartosian/sui_helpers/peer_checker/domain/enums"
)

const (
	peerSeparator = "/"
	peerCount     = 4
)

type (
	PeerData struct {
		Address string `yaml:"address"`
	}

	P2PConfig struct {
		SeedPeers []PeerData `yaml:"seed-peers"`
	}

	NodeConfigYaml struct {
		P2PConfig P2PConfig `yaml:"p2p-config"`
	}

	Checker struct {
		Peers     []Peer
		rpcClient jsonrpc.RPCClient
	}
)

func NewChecker(path string, network enums.NetworkType) (*Checker, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var configData NodeConfigYaml

	err = yaml.Unmarshal(file, &configData)
	if err != nil {
		return nil, err
	}

	peers, err := configData.P2PConfig.parsePeers()
	if err != nil {
		return nil, err
	}

	return &Checker{
		Peers:     peers,
		rpcClient: jsonrpc.NewClient(network.ToRPC()),
	}, nil
}
