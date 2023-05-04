package checker

import (
	"net/http"
	"sync"
	"time"

	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ipinfo/go/v2/ipinfo/cache"
	"github.com/mum4k/termdash"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/suimon/cmd/checker/config"
	"github.com/bartosian/suimon/cmd/checker/dashboardbuilder"
	"github.com/bartosian/suimon/cmd/checker/dashboardbuilder/dashboards"
	"github.com/bartosian/suimon/cmd/checker/enums"
	"github.com/bartosian/suimon/cmd/checker/tablebuilder"
	"github.com/bartosian/suimon/pkg/log"
)

const ipInfoCacheExp = 5 * time.Minute

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
		tableConfig      tablebuilder.TableConfig

		DashboardBuilder *dashboardbuilder.DashboardBuilder

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

func (checker *Checker) DrawTables() {
	if checker.suimonConfig.MonitorsConfig.RPCTable.Display && len(checker.rpc) > 0 {
		checker.tableBuilderRPC.Build()
	}
	if checker.suimonConfig.MonitorsConfig.NodeTable.Display && len(checker.node) > 0 {
		checker.tableBuilderNode.Build()
	}
	if checker.suimonConfig.MonitorsConfig.PeersTable.Display && len(checker.peers) > 0 {
		checker.tableBuilderPeer.Build()
	}
}

func (checker *Checker) DrawDashboards() {
	var (
		monitorsConfig   = checker.suimonConfig.MonitorsConfig
		dashboardBuilder = checker.DashboardBuilder
		dashCells        = dashboardBuilder.Cells
		ticker           = time.NewTicker(watchHostsTimeout)
		logsCH           = make(chan string)
		wg               sync.WaitGroup
	)

	defer ticker.Stop()

	var stream = func() {
		defer wg.Done()

		if err := checker.logger.StreamFromService("suid", logsCH); err != nil {
			if err := checker.logger.StreamFromContainer("sui-node", logsCH); err != nil {
				if err := checker.logger.StreamFromScreen("sui", logsCH); err != nil {
					return
				}
			}
		}
	}

	var draw = func(hosts []Host) {
		defer wg.Done()

		doneCH := make(chan struct{}, len(hosts))

		for {
			select {
			case <-ticker.C:
				for _, host := range hosts {
					go func(host Host) {
						for idx, dashCell := range dashCells {
							cellName := dashboards.CellName(idx)

							metric := checker.getMetricForDashboardCell(cellName)
							options := checker.getOptionsForDashboardCell(cellName)

							dashCell.Write(metric, options)
						}

						doneCH <- struct{}{}
					}(host)
				}

				for i := 0; i < len(hosts); i++ {
					<-doneCH
				}
			case log := <-logsCH:
				dashCell := dashCells[dashboards.CellNameNodeLogs]
				options := checker.getOptionsForDashboardCell(dashboards.CellNameNodeLogs)

				dashCell.Write(log+"\n", options)
			case <-dashboardBuilder.Ctx.Done():
				close(doneCH)
				close(logsCH)

				return
			}
		}
	}

	wg.Add(1)

	go func() {
		defer wg.Done()

		checker.WatchHosts()
	}()

	if monitorsConfig.NodeTable.Display && len(checker.node) > 0 {
		wg.Add(2)

		go stream()
		go draw(checker.node)
	}

	if err := termdash.Run(dashboardBuilder.Ctx, dashboardBuilder.Terminal, dashboardBuilder.Dashboard, termdash.KeyboardSubscriber(dashboardBuilder.Quitter)); err != nil {
		panic(err)
	}

	wg.Wait()
}
