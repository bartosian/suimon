package checker

import (
	"net/http"
	"os"

	"github.com/ybbus/jsonrpc/v3"
	"gopkg.in/yaml.v3"

	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/tablebuilder"
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
		peers            []Peer
		rpcList          []RPCHost
		rpcClient        jsonrpc.RPCClient
		httpClient       *http.Client
		tableBuilderPeer *tablebuilder.TableBuilder
		tableBuilderRPC  *tablebuilder.TableBuilder
		tableConfig      tablebuilder.TableConfig
		network          enums.NetworkType
	}
)

var rpcList = map[enums.NetworkType][]string{
	enums.NetworkTypeDevnet: {
		"https://fullnode.devnet.sui.io",
	},
	enums.NetworkTypeTestnet: {
		"https://rpc-office.cosmostation.io/sui-testnet-wave-2",
		"https://rpc.ankr.com/sui_testnet",
		"https://sui-testnet.public.blastapi.io",
		"https://sui-api.rpcpool.com",
		"https://fullnode.testnet.sui.io",
	},
}

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

	peers, err := parsePeers(nodeYaml.Config.SeedPeers)
	if err != nil {
		return nil, err
	}

	hosts, err := parseRPCHosts(rpcList[network])
	if err != nil {
		return nil, err
	}

	return &Checker{
		peers:     peers,
		rpcClient: jsonrpc.NewClient(network.ToRPC()),
		rpcList:   hosts,
		network:   network,
	}, nil
}
