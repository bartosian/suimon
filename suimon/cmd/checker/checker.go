package checker

import (
	"github.com/ipinfo/go/v2/ipinfo/cache"
	"net/http"
	"time"

	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/config"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/tablebuilder"
)

const (
	peerSeparator    = "/"
	addressSeparator = ":"
	peerCount        = 4
	freeIpInfoToken  = "55f30ce0213aa7"
	ipInfoCacheExp   = 5 * time.Minute
)

type (
	Checker struct {
		suimonConfig config.SuimonConfig
		nodeConfig   config.NodeConfig

		peers   []Peer
		rpcList []RPCHost
		node    Node

		rpcClient  jsonrpc.RPCClient
		httpClient *http.Client
		ipClient   *ipinfo.Client

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

	ipClient := ipinfo.NewClient(
		nil, ipinfo.NewCache(cache.NewInMemory().WithExpiration(ipInfoCacheExp)), freeIpInfoToken)

	return &Checker{
		rpcClient:    jsonrpc.NewClient(suimonConfig.NetworkType.ToRPC()),
		httpClient:   httpClient,
		ipClient:     ipClient,
		suimonConfig: suimonConfig,
		nodeConfig:   nodeConfig,
	}, nil
}
