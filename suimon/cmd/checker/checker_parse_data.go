package checker

import (
	"errors"
	"sort"
	"sync"
	"time"

	"github.com/hashicorp/go-multierror"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/pkg/address"
	"github.com/bartosian/sui_helpers/suimon/pkg/progress"
)

const (
	httpClientTimeout = 3 * time.Second
)

func (checker *Checker) getAddressInfoByTableType(tableType enums.TableType) ([]AddressInfo, error) {
	var addresses []AddressInfo

	switch tableType {
	case enums.TableTypeNode:
		addressRPC, addressMetrics := checker.nodeConfig.JSONRPCAddress, checker.nodeConfig.MetricsAddress

		var (
			hostPortRPC     *address.HostPort
			hostPortMetrics *address.HostPort
		)

		if addressRPC != "" {
			hostPort, err := address.ParseIpPort(addressRPC)
			if err != nil {
				return nil, errors.New("invalid json-rpc-address in fullnode.yaml")
			}

			hostPortRPC = hostPort
		}

		if addressMetrics != "" {
			hostPort, err := address.ParseIpPort(addressMetrics)
			if err != nil {
				return nil, errors.New("invalid metrics-address in fullnode.yaml")
			}

			hostPortMetrics = hostPort
		}

		addressInfo := AddressInfo{HostPort: *hostPortRPC, Ports: make(map[enums.PortType]string)}
		if hostPortRPC.Port != nil {
			addressInfo.Ports[enums.PortTypeRPC] = *hostPortRPC.Port
		}

		if hostPortMetrics.Port != nil {
			addressInfo.Ports[enums.PortTypeMetrics] = *hostPortMetrics.Port
		}

		addresses = append(addresses, addressInfo)
	case enums.TableTypePeers:
		peers := checker.nodeConfig.P2PConfig.SeedPeers
		if len(peers) == 0 {
			return nil, errors.New("peers not found in fullnode.yaml")
		}

		for _, peer := range peers {
			hostPort, err := address.ParsePeer(peer.Address)
			if err != nil {
				return nil, errors.New("invalid peer in fullnode.yaml")
			}

			addressInfo := AddressInfo{HostPort: *hostPort, Ports: make(map[enums.PortType]string)}
			if hostPort.Port != nil {
				addressInfo.Ports[enums.PortTypePeer] = *hostPort.Port
			}

			addresses = append(addresses, addressInfo)
		}
	case enums.TableTypeRPC:
		rpc := checker.suimonConfig.GetRPCByNetwork()
		if len(rpc) == 0 {
			return nil, errors.New("rpc not found in suimon.yaml")
		}

		for _, url := range rpc {
			hostPort, err := address.ParseURL(url)
			if err != nil {
				return nil, errors.New("invalid rpc url in suimon.yaml")
			}

			addressInfo := AddressInfo{HostPort: *hostPort, Ports: make(map[enums.PortType]string)}
			if hostPort.Port != nil {
				addressInfo.Ports[enums.PortTypeRPC] = *hostPort.Port
			}

			addresses = append(addresses, addressInfo)
		}
	}

	return addresses, nil
}

func (checker *Checker) Init() error {
	var (
		wg             sync.WaitGroup
		errChan        = make(chan error, 3)
		suimonConfig   = checker.suimonConfig
		monitorsConfig = suimonConfig.MonitorsConfig
	)

	var getData = func(tableType enums.TableType, progressColor progress.Color) {
		var (
			progressChan = progress.NewProgressBar("PARSING DATA FOR "+string(tableType), progressColor)
			addresses    []AddressInfo
			hosts        []Host
			err          error
		)

		defer wg.Done()

		if addresses, err = checker.getAddressInfoByTableType(tableType); err != nil {
			errChan <- err
		}

		if hosts, err = checker.createHosts(addresses); err != nil {
			errChan <- err
		}

		checker.setHostsByTableType(tableType, hosts)

		progressChan <- struct{}{}
	}

	// parse data for the RPC servers
	wg.Add(1)

	go getData(enums.TableTypeRPC, progress.ColorBlue)

	// parse data for the user servers
	if monitorsConfig.NodeTable.Display {
		wg.Add(1)

		go getData(enums.TableTypeNode, progress.ColorRed)
	}

	// parse data for the peers servers
	if monitorsConfig.PeersTable.Display {
		wg.Add(1)

		go getData(enums.TableTypePeers, progress.ColorGreen)
	}

	wg.Wait()
	close(errChan)

	if len(checker.rpc) == 0 || len(checker.peers) == 0 && len(checker.node) == 0 {
		var err error

		for parseErr := range errChan {
			err = multierror.Append(err, parseErr)
		}

		return err
	}

	rpc := checker.rpc

	if len(rpc) > 1 {
		sort.Slice(rpc, func(left, right int) bool {
			return rpc[left].Metrics.TotalTransactionNumber > rpc[right].Metrics.TotalTransactionNumber
		})

		sort.SliceStable(rpc, func(left, right int) bool {
			return rpc[left].Metrics.LatestCheckpoint > rpc[right].Metrics.LatestCheckpoint
		})
	}

	checker.rpc = rpc

	var setStatus = func(tableType enums.TableType) {
		hosts := checker.getHostsByTableType(tableType)
		comparatorRPC := rpc[0]

		for idx := range hosts {
			hosts[idx].SetPctProgress(enums.MetricTypeTxSyncProgress, comparatorRPC)
			hosts[idx].SetPctProgress(enums.MetricTypeCheckSyncProgress, comparatorRPC)
			hosts[idx].SetStatus(comparatorRPC)
		}
	}

	setStatus(enums.TableTypeRPC)

	if monitorsConfig.NodeTable.Display {
		setStatus(enums.TableTypeNode)
	}

	if monitorsConfig.PeersTable.Display {
		setStatus(enums.TableTypePeers)
	}

	return nil
}
