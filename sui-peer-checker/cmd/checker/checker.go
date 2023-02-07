package checker

import (
	"os"

	"github.com/ybbus/jsonrpc/v3"
	"gopkg.in/yaml.v3"

	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/enums"
)

const (
	peerSeparator = "/"
	peerCount     = 4
)

type (
	PeerData struct {
		Address string `yaml:"address"`
	}

	NodeYaml struct {
		Config Config `yaml:"p2p-config"`
	}

	Checker struct {
		peers        []Peer
		rpcClient    jsonrpc.RPCClient
		tableBuilder *TableBuilder
	}
)

func NewChecker(path string, network enums.NetworkType) (*Checker, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var nodeYaml NodeYaml
	err = yaml.Unmarshal(file, &nodeYaml)
	if err != nil {
		return nil, err
	}

	peers, err := nodeYaml.Config.parsePeers()
	if err != nil {
		return nil, err
	}

	return &Checker{
		peers:        peers,
		rpcClient:    jsonrpc.NewClient(network.ToRPC()),
		tableBuilder: NewTableBuilder(),
	}, nil
}
