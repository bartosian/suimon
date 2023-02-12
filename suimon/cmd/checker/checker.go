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
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/tablebuilder/tables"
)

const (
	freeIpInfoToken = "55f30ce0213aa7"
	ipInfoCacheExp  = 5 * time.Minute
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
	}
)

func NewChecker(suimonConfig config.SuimonConfig, nodeConfig config.NodeConfig, networkConfig enums.NetworkType) (*Checker, error) {
	suimonConfig.SetNetworkConfig(networkConfig)

	rpcClient := jsonrpc.NewClient(suimonConfig.NetworkType.ToRPC())
	httpClient := &http.Client{Timeout: httpClientTimeout}
	ipClient := ipinfo.NewClient(httpClient, ipinfo.NewCache(cache.NewInMemory().WithExpiration(ipInfoCacheExp)), freeIpInfoToken)

	return &Checker{
		rpcClient:    rpcClient,
		httpClient:   httpClient,
		ipClient:     ipClient,
		suimonConfig: suimonConfig,
		nodeConfig:   nodeConfig,
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

func (checker *Checker) InitTables() {
	displayConfig := checker.suimonConfig.MonitorsConfig

	if displayConfig.RPCTable.Display {
		checker.InitTable(enums.TableTypeRPC)
	}

	if displayConfig.NodeTable.Display {
		checker.InitTable(enums.TableTypeNode)
	}

	if displayConfig.PeersTable.Display {
		checker.InitTable(enums.TableTypePeers)
	}
}

func (checker *Checker) InitTable(tableType enums.TableType) {
	hosts := checker.getHostsByTableType(tableType)
	suimonConfig := checker.suimonConfig

	tableConfig := tablebuilder.TableConfig{
		Name:         tables.GetTableTitleSUI(suimonConfig.NetworkType, tableType, suimonConfig.MonitorsVisual.EnableEmojis),
		Colors:       tablebuilder.GetTableColorsFromString(suimonConfig.MonitorsVisual.ColorScheme),
		Tag:          tables.TableTagSUINode,
		Style:        tables.TableStyleSUINode,
		RowsCount:    0,
		ColumnsCount: len(tables.ColumnConfigSUINode),
		SortConfig:   tables.TableSortConfigSUINode,
	}

	columns := make(tablebuilder.Columns, len(tables.ColumnConfigSUINode))
	emojisEnabled := checker.suimonConfig.MonitorsVisual.EnableEmojis

	for idx, config := range tables.ColumnConfigSUINode {
		columns[idx].Config = config
	}

	for _, host := range hosts {
		if !host.Metrics.Updated {
			continue
		}

		tableConfig.RowsCount++

		var status any = host.Status
		if !emojisEnabled {
			status = host.Status.StatusToPlaceholder()
		}

		port := host.Ports[enums.PortTypeRPC]
		if tableType == enums.TableTypePeers {
			port = host.Ports[enums.PortTypePeer]
		}

		columns[tables.ColumnNameSUINodeStatus].SetValue(status)
		columns[tables.ColumnNameSUINodeAddress].SetValue(host.HostPort.Address)
		columns[tables.ColumnNameSUINodePortRPC].SetValue(port)
		columns[tables.ColumnNameSUINodeTotalTransactions].SetValue(host.Metrics.TotalTransactionNumber)
		columns[tables.ColumnNameSUINodeHighestCheckpoints].SetValue(host.Metrics.HighestSyncedCheckpoint)
		columns[tables.ColumnNameSUINodeConnectedPeers].SetValue(host.Metrics.SuiNetworkPeers)
		columns[tables.ColumnNameSUINodeUptime].SetValue(host.Metrics.Uptime)
		columns[tables.ColumnNameSUINodeVersion].SetValue(host.Metrics.Version)
		columns[tables.ColumnNameSUINodeCommit].SetValue(host.Metrics.Commit)

		if host.Location == nil {
			columns[tables.ColumnNameSUINodeCompany].SetValue(nil)
			columns[tables.ColumnNameSUINodeCountry].SetValue(nil)

			continue
		}

		columns[tables.ColumnNameSUINodeCompany].SetValue(host.Location.Provider)

		var country any = host.Location.String()
		if !emojisEnabled {
			country = host.Location.CountryName
		}

		columns[tables.ColumnNameSUINodeCountry].SetValue(country)
	}

	tableConfig.Columns = columns

	checker.setTableBuilderTableType(tableType, tableConfig)
}

func (checker *Checker) DrawTable() {
	if checker.suimonConfig.MonitorsConfig.NodeTable.Display {
		checker.tableBuilderNode.Build()
	}
	if checker.suimonConfig.MonitorsConfig.RPCTable.Display {
		checker.tableBuilderRPC.Build()
	}
	if checker.suimonConfig.MonitorsConfig.PeersTable.Display {
		checker.tableBuilderPeer.Build()
	}
}
