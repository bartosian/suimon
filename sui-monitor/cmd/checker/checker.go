package checker

import (
	"net/http"
	"os"

	"github.com/oschwald/geoip2-golang"
	"github.com/ybbus/jsonrpc/v3"
	"gopkg.in/yaml.v3"

	"github.com/bartosian/sui_helpers/sui-monitor/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/sui-monitor/cmd/checker/tablebuilder"
)

const (
	peerSeparator    = "/"
	addressSeparator = ":"
	peerCount        = 4
)

type (
	PeerData struct {
		Address string `yaml:"address"`
	}

	Genesis struct {
		GenesisFileLocation string `yaml:"genesis-file-location"`
	}

	Config struct {
		SeedPeers []PeerData `yaml:"seed-peers"`
	}

	NodeYaml struct {
		DbPath                string  `yaml:"db-path"`
		MetricsAddress        string  `yaml:"metrics-address"`
		JsonRPCAddress        string  `yaml:"json-rpc-address"`
		WebsocketAddress      string  `yaml:"websocket-address"`
		EnableEventProcessing bool    `yaml:"enable-event-processing"`
		Config                Config  `yaml:"p2p-config"`
		Genesis               Genesis `yaml:"genesis"`
	}

	Checker struct {
		peers            []Peer
		rpcList          []RPCHost
		node             Node
		rpcClient        jsonrpc.RPCClient
		httpClient       *http.Client
		geoDbClient      *geoip2.Reader
		tableBuilderPeer *tablebuilder.TableBuilder
		tableBuilderNode *tablebuilder.TableBuilder
		tableBuilderRPC  *tablebuilder.TableBuilder
		tableConfig      tablebuilder.TableConfig
		network          enums.NetworkType
		nodeYaml         NodeYaml
	}
)

func NewChecker(path string, network enums.NetworkType) (*Checker, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var result NodeYaml
	err = yaml.Unmarshal(file, &result)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Timeout: httpClientTimeout,
	}

	return &Checker{
		rpcClient:  jsonrpc.NewClient(network.ToRPC()),
		httpClient: httpClient,
		network:    network,
		nodeYaml:   result,
	}, nil
}
