package checker

import (
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder/dashboards"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ipinfo/go/v2/ipinfo/cache"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/config"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/tablebuilder"
	"github.com/bartosian/sui_helpers/suimon/pkg/log"
)

const (
	ipInfoCacheExp  = 5 * time.Minute
	dashboardNoData = "⌛ loading"
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

func (checker *Checker) DrawDashboards() {
	var (
		monitorsConfig   = checker.suimonConfig.MonitorsConfig
		dashboardBuilder = checker.DashboardBuilder
		ticker           = time.NewTicker(watchHostsTimeout)
		wg               sync.WaitGroup
	)

	defer ticker.Stop()

	var draw = func(hosts []Host) {
		defer wg.Done()

		doneCH := make(chan struct{}, len(hosts))

		for {
			select {
			case <-ticker.C:
				for _, host := range hosts {
					go func(host Host) {
						dashCells := dashboardBuilder.Cells

						for idx, dashCell := range dashCells {
							metric := host.getMetricByDashboardCell(dashboards.CellName(idx))

							if metric == "" {
								metric = dashboardNoData
							}

							switch v := dashCell.Widget.(type) {
							case *text.Text:
								v.Reset()

								v.Write(metric, text.WriteCellOpts(cell.Bold()))
							case *gauge.Gauge:
								percentage := 0

								if metric != dashboardNoData {
									percentage, _ = strconv.Atoi(metric)
								}

								v.Percent(percentage)
							case *segmentdisplay.SegmentDisplay:
								v.Write([]*segmentdisplay.TextChunk{
									segmentdisplay.NewChunk(metric),
								})
							}
						}

						doneCH <- struct{}{}
					}(host)
				}

				for i := 0; i < len(hosts); i++ {
					<-doneCH
				}
			case <-dashboardBuilder.Ctx.Done():
				return
			}
		}
	}

	wg.Add(1)

	go func() {
		defer wg.Done()

		checker.WatchHosts()
	}()

	//if monitorsConfig.RPCTable.Display && len(checker.rpc) > 0 {
	//	wg.Add(1)
	//
	//	go draw(checker.rpc)
	//}

	if monitorsConfig.NodeTable.Display && len(checker.node) > 0 {
		wg.Add(1)

		go draw(checker.node)
	}

	//if monitorsConfig.PeersTable.Display && len(checker.peers) > 0 {
	//	wg.Add(1)
	//
	//	go draw(checker.peers)
	//}

	if err := termdash.Run(dashboardBuilder.Ctx, dashboardBuilder.Terminal, dashboardBuilder.Dashboard, termdash.KeyboardSubscriber(dashboardBuilder.Quitter)); err != nil {
		panic(err)
	}

	wg.Wait()
}

func (checker *Checker) InitDashboard() error {
	var (
		dashboard *dashboardbuilder.DashboardBuilder
		err       error
	)

	if dashboard, err = dashboardbuilder.NewDashboardBuilder(); err != nil {
		return err
	}

	checker.DashboardBuilder = dashboard

	return nil
}
