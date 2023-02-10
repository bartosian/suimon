package checker

import (
	"net/http"

	"github.com/oschwald/geoip2-golang"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/config"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/tablebuilder"
)

const (
	peerSeparator    = "/"
	addressSeparator = ":"
	peerCount        = 4
)

type (
	Checker struct {
		suimonConfig config.SuimonConfig
		nodeConfig   config.NodeConfig

		peers   []Peer
		rpcList []RPCHost
		node    Node

		rpcClient   jsonrpc.RPCClient
		httpClient  *http.Client
		geoDbClient *geoip2.Reader

		tableBuilderPeer *tablebuilder.TableBuilder
		tableBuilderNode *tablebuilder.TableBuilder
		tableBuilderRPC  *tablebuilder.TableBuilder

		tableConfig tablebuilder.TableConfig
	}
)

func NewChecker(suimonConfig config.SuimonConfig, nodeConfig config.NodeConfig) (*Checker, error) {
	httpClient := &http.Client{
		Timeout: httpClientTimeout,
	}

	return &Checker{
		rpcClient:    jsonrpc.NewClient(suimonConfig.NetworkType.ToRPC()),
		httpClient:   httpClient,
		suimonConfig: suimonConfig,
		nodeConfig:   nodeConfig,
	}, nil
}
