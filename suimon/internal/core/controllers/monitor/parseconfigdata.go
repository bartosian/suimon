package monitor

import (
	"errors"
	"sort"
	"sync"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/progress"
)

// ParseConfigData retrieves the latest data from all active hosts and updates the MonitorController's internal state with the new data.
// The function parses the data for each table type and sets the corresponding rpcgw and dashboard options accordingly.
// Returns an error if the data cannot be retrieved from any of the active hosts or if there is an issue parsing the data for any table type.
func (c *Controller) ParseConfigData() error {
	monitorsConfig := c.config.MonitorsConfig

	enabledSystemState := monitorsConfig.SystemStateTable.Display || monitorsConfig.ActiveValidatorsTable.Display || monitorsConfig.ValidatorReportsTable.Display ||
		monitorsConfig.ValidatorsAtRiskTable.Display || monitorsConfig.ValidatorsCountsTable.Display

	tableMap := map[enums.TableType]bool{
		enums.TableTypeNode:        monitorsConfig.NodeTable.Display,
		enums.TableTypeValidator:   monitorsConfig.ValidatorTable.Display,
		enums.TableTypePeers:       monitorsConfig.PeersTable.Display,
		enums.TableTypeSystemState: enabledSystemState,
	}

	if err := c.getHostsData(enums.TableTypeRPC, progress.ColorBlue); err != nil {
		return err
	}

	if err := c.sortHosts(enums.TableTypeRPC); err != nil {
		return err
	}

	if err := c.setHostsHealth(enums.TableTypeRPC); err != nil {
		return err
	}

	disabledTables := 0

	errChan := make(chan error, len(tableMap))
	defer close(errChan)

	var wg sync.WaitGroup

	for tableType, isEnabled := range tableMap {
		if !isEnabled {
			disabledTables++

			continue
		}

		wg.Add(1)

		go func(tt enums.TableType) {
			defer wg.Done()

			if err := c.getHostsData(tt, progress.ColorBlue); err != nil {
				errChan <- err

				return
			}

			if err := c.setHostsHealth(tt); err != nil {
				errChan <- err
			}
		}(tableType)
	}

	wg.Wait()

	select {
	case err := <-errChan:
		return err
	default:
		if disabledTables == len(tableMap) {
			return errors.New("all tables disabled in suimon.yaml")
		}
		return nil
	}
}

// getHostsData retrieves the latest data for the specified table type from all active hosts and updates the MonitorController's internal state with the new data.
// The function retrieves data for each host in parallel and displays a progress bar indicating the progress of the data retrieval process.
// Returns an error if the data cannot be retrieved from any of the active hosts or if there is an issue updating the CheckerController's internal state.
func (c *Controller) getHostsData(tableType enums.TableType, progressColor progress.Color) error {
	progressChan := progress.NewProgressBar("PARSING DATA FOR "+string(tableType), progressColor)
	defer func() { progressChan <- struct{}{} }()

	var (
		addresses []host.AddressInfo
		hosts     []host.Host
		err       error
	)

	if tableType == enums.TableTypeSystemState {
		rpcHost := c.hosts.rpc[0]

		if err := rpcHost.GetDataByMetric(enums.RPCMethodGetSuiSystemState); err != nil {
			return err
		}
	}

	if addresses, err = c.getAddressInfoByTableType(tableType); err != nil {
		return err
	}

	if hosts, err = c.createHosts(tableType, addresses); err != nil {
		return err
	}

	return c.setHostsByTableType(tableType, hosts)
}

// sortHosts sorts the active hosts for the specified table type based on their corresponding metric values.
// The function retrieves the relevant metric for each host, sorts the hosts by their metric values, and updates the CheckerController's internal state accordingly.
// Returns an error if the specified table type is invalid or if there is an issue sorting the hosts based on their corresponding metric values.
func (c *Controller) sortHosts(tableType enums.TableType) error {
	hosts, err := c.getHostsByTableType(tableType)
	if err != nil {
		return err
	}

	if len(hosts) > 1 {
		sort.Slice(hosts, func(left, right int) bool {
			return hosts[left].Metrics.TotalTransactionsBlocks > hosts[right].Metrics.TotalTransactionsBlocks
		})

		sort.SliceStable(hosts, func(left, right int) bool {
			return hosts[left].Metrics.LatestCheckpoint > hosts[right].Metrics.LatestCheckpoint
		})
	}

	return c.setHostsByTableType(tableType, hosts)
}

// setHostsHealth retrieves the latest health information for all active hosts and updates the CheckerController's internal state with the new information.
// The function retrieves health information for each host in parallel and sets the corresponding health status in the internal state.
// Returns an error if the health information cannot be retrieved from any of the active hosts or if there is an issue updating the CheckerController's internal state.
func (c *Controller) setHostsHealth(tableType enums.TableType) error {
	hosts, err := c.getHostsByTableType(tableType)
	if err != nil {
		return err
	}

	rpcHost := c.hosts.rpc[0]

	for idx := range hosts {
		metrics := hosts[idx].Metrics

		checkpointExecBacklog := metrics.HighestKnownCheckpoint - metrics.LastExecutedCheckpoint
		checkpointSyncBacklog := metrics.HighestKnownCheckpoint - metrics.HighestSyncedCheckpoint

		if err = hosts[idx].SetPctProgress(enums.MetricTypeTxSyncPercentage, rpcHost); err != nil {
			return err
		}

		if err = hosts[idx].SetPctProgress(enums.MetricTypeCheckSyncPercentage, rpcHost); err != nil {
			return err
		}

		if err = hosts[idx].Metrics.SetValue(enums.MetricTypeCheckpointExecBacklog, checkpointExecBacklog); err != nil {
			return err
		}

		if err = hosts[idx].Metrics.SetValue(enums.MetricTypeCheckpointSyncBacklog, checkpointSyncBacklog); err != nil {
			return err
		}

		hosts[idx].SetStatus(rpcHost)
	}

	return nil
}
