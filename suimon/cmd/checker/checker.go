package checker

import (
	"net/http"
	"time"

	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ipinfo/go/v2/ipinfo/cache"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/config"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/tablebuilder"
	"github.com/bartosian/sui_helpers/suimon/pkg/log"
)

const (
	ipInfoCacheExp = 5 * time.Minute
)

type (
	Checker struct {
		suimonConfig config.SuimonConfig
		nodeConfig   config.NodeConfig

		rpc   []Host
		node  []Host
		peers []Host

		rpcClient  jsonrpc.RPCClient
		httpClient *http.Client
		ipClient   *ipinfo.Client

		tableBuilderPeer *tablebuilder.TableBuilder
		tableBuilderNode *tablebuilder.TableBuilder
		tableBuilderRPC  *tablebuilder.TableBuilder

		tableConfig tablebuilder.TableConfig

		logger log.Logger
	}
)

func NewChecker(suimonConfig config.SuimonConfig, nodeConfig config.NodeConfig, networkConfig enums.NetworkType) (*Checker, error) {
	suimonConfig.SetNetworkConfig(networkConfig)

	rpcClient := jsonrpc.NewClient(suimonConfig.NetworkType.ToRPC())
	httpClient := &http.Client{Timeout: httpClientTimeout}
	ipClient := ipinfo.NewClient(httpClient, ipinfo.NewCache(cache.NewInMemory().WithExpiration(ipInfoCacheExp)), suimonConfig.IPLookup.AccessToken)

	return &Checker{
		rpcClient:    rpcClient,
		httpClient:   httpClient,
		ipClient:     ipClient,
		suimonConfig: suimonConfig,
		nodeConfig:   nodeConfig,
		logger:       log.NewLogger(),
	}, nil
}

func (checker *Checker) getHostsByTableType(tableType enums.TableType) []Host {
	var hosts []Host

	switch tableType {
	case enums.TableTypeNode:
		hosts = checker.node
	case enums.TableTypePeers:
		hosts = checker.peers
	case enums.TableTypeRPC:
		hosts = checker.rpc
	}

	return hosts
}

func (checker *Checker) setHostsByTableType(tableType enums.TableType, hosts []Host) {
	switch tableType {
	case enums.TableTypeNode:
		checker.node = hosts
	case enums.TableTypePeers:
		checker.peers = hosts
	case enums.TableTypeRPC:
		checker.rpc = hosts
	}
}

func (checker *Checker) setTableBuilderTableType(tableType enums.TableType, tc tablebuilder.TableConfig) {
	tableBuilder := tablebuilder.NewTableBuilder(tc)

	switch tableType {
	case enums.TableTypeNode:
		checker.tableBuilderNode = tableBuilder
	case enums.TableTypePeers:
		checker.tableBuilderPeer = tableBuilder
	case enums.TableTypeRPC:
		checker.tableBuilderRPC = tableBuilder
	}
}

func (checker *Checker) DrawTable() {
	if checker.suimonConfig.MonitorsConfig.RPCTable.Display {
		checker.tableBuilderRPC.Build()
	}
	if checker.suimonConfig.MonitorsConfig.NodeTable.Display {
		checker.tableBuilderNode.Build()
	}
	if checker.suimonConfig.MonitorsConfig.PeersTable.Display {
		checker.tableBuilderPeer.Build()
	}
}
