package monitor

import (
	"errors"
	"sort"
	"sync"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/pkg/progress"
)

// ParseConfigData retrieves data from hosts and sets their health based on the selected tables.
// The function first retrieves data from the RPC table, sorts the hosts, and sets their health. It then
// retrieves data from the selected tables in parallel using goroutines. For each selected table, the function
// retrieves data from the hosts and sets their health. Any errors that occur during this process are sent to
// a channel. If an error is received from the channel, it is returned immediately. If no errors are received,
// the function returns nil.
func (c *Controller) ParseConfigData(monitorType enums.MonitorType) error {
	if err := c.ParseConfigRPC(); err != nil {
		return err
	}

	var (
		rpcTables = map[enums.TableType]bool{
			enums.TableTypeActiveValidators:   true,
			enums.TableTypeValidatorReports:   true,
			enums.TableTypeValidatorsAtRisk:   true,
			enums.TableTypeGasPriceAndSubsidy: true,
			enums.TableTypeValidatorsParams:   true,
			enums.TableTypeRPC:                true,
		}

		tablesToParse []enums.TableType
	)

	switch monitorType {
	case enums.MonitorTypeStatic:
		tablesToParse = make([]enums.TableType, 0, len(c.selectedTables))

		for _, table := range c.selectedTables {
			if _, ok := rpcTables[table]; ok {
				continue
			}

			tablesToParse = append(tablesToParse, table)
		}
	case enums.MonitorTypeDynamic:
		tablesToParse = []enums.TableType{c.selectedDashboard}
	}

	errChan := make(chan error, len(tablesToParse))

	var wg sync.WaitGroup

	for _, tableType := range tablesToParse {
		if tableType == enums.TableTypeGasPriceAndSubsidy {
			continue
		}

		wg.Add(1)

		go func(table enums.TableType) {
			defer wg.Done()

			if err := c.getHostsData(table); err != nil {
				errChan <- err

				return
			}

			if err := c.setHostsHealth(table); err != nil {
				errChan <- err
			}
		}(tableType)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

// ParseConfigRPC fetches hosts data for the RPC table, sorts the hosts in
// alphabetical order, and sets their health status.
func (c *Controller) ParseConfigRPC() error {
	if err := c.getHostsData(enums.TableTypeRPC); err != nil {
		return err
	}

	if len(c.hosts.rpc) == 0 {
		return errors.New("no public RPC provided")
	}

	if err := c.sortHosts(enums.TableTypeRPC); err != nil {
		return err
	}

	if err := c.setHostsHealth(enums.TableTypeRPC); err != nil {
		return err
	}

	return nil
}

// getHostsData retrieves the latest data for the specified table type from all active hosts and updates the MonitorController's internal state with the new data.
// The function retrieves data for each host in parallel and displays a progress bar indicating the progress of the data retrieval process.
// Returns an error if the data cannot be retrieved from any of the active hosts or if there is an issue updating the CheckerController's internal state.
func (c *Controller) getHostsData(table enums.TableType) error {
	progressChan := progress.NewProgressBar("PARSING DATA FOR "+string(table), progress.ColorBlue)
	defer func() { progressChan <- struct{}{} }()

	addresses, err := c.getAddressInfoByTableType(table)
	if err != nil {
		return err
	}

	hosts, err := c.createHosts(table, addresses)
	if err != nil {
		return err
	}

	return c.setHostsByTableType(table, hosts)
}

// sortHosts sorts the active hosts for the specified table type based on their corresponding metric values.
// The function retrieves the relevant metric for each host, sorts the hosts by their metric values, and updates the CheckerController's internal state accordingly.
// Returns an error if the specified table type is invalid or if there is an issue sorting the hosts based on their corresponding metric values.
func (c *Controller) sortHosts(tableType enums.TableType) error {

	if tableType == enums.TableTypeGasPriceAndSubsidy {
		return nil
	}

	hosts, err := c.getHostsByTableType(tableType)
	if err != nil {
		return err
	}

	if len(hosts) > 1 {
		sort.Slice(hosts, func(left, right int) bool {
			return hosts[left].Status > hosts[right].Status
		})

		sort.SliceStable(hosts, func(left, right int) bool {
			return hosts[left].Metrics.TotalTransactionsBlocks > hosts[right].Metrics.TotalTransactionsBlocks
		})
	}

	return nil
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
