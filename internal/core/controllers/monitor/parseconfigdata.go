package monitor

import (
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainmetrics "github.com/bartosian/suimon/internal/core/domain/metrics"
	"github.com/bartosian/suimon/internal/pkg/progress"
)

var rpcTables = map[enums.TableType]bool{
	enums.TableTypeActiveValidators:   true,
	enums.TableTypeValidatorReports:   true,
	enums.TableTypeValidatorsAtRisk:   true,
	enums.TableTypeGasPriceAndSubsidy: true,
	enums.TableTypeValidatorsParams:   true,
	enums.TableTypeRPC:                true,
	enums.TableTypeProtocol:           true,
}

// ParseConfigData retrieves data from hosts and sets their health based on the selected tables.
// The function first retrieves data from the RPC table, sorts the hosts, and sets their health. It then
// retrieves data from the selected tables in parallel using goroutines. For each selected table, the function
// retrieves data from the hosts and sets their health. Any errors that occur during this process are sent to
// a channel. If an error is received from the channel, it is returned immediately. If no errors are received,
// the function returns nil.
func (c *Controller) ParseConfigData(monitorType enums.MonitorType) error {
	if len(c.selectedTables) == 0 || (len(c.selectedTables) > 1 || c.selectedTables[0] != enums.TableTypeReleases) {
		if err := c.ParseConfigRPC(); err != nil {
			return err
		}
	}

	tablesToParse := c.determineTablesToParse(monitorType)

	errChan := make(chan error, len(tablesToParse))

	var wg sync.WaitGroup

	for _, tableType := range tablesToParse {
		wg.Add(1)

		go func(table enums.TableType) {
			defer wg.Done()

			if err := c.getTableData(table); err != nil {
				errChan <- fmt.Errorf("error processing table %s: %w", table, err)
			}
		}(tableType)
	}

	wg.Wait()
	close(errChan)

	return checkErrors(errChan)
}

// determineTablesToParse decides which tables to parse based on the monitor type.
// If the monitor type is static, it checks each selected table against a predefined list of tables to retrieve data.
// If the table is in the list, it is added to the tables to parse.
// If the monitor type is dynamic, only the selected dashboard is added to the tables to parse.
// The function returns a slice of tables to parse.
func (c *Controller) determineTablesToParse(monitorType enums.MonitorType) []enums.TableType {
	var tablesToParse []enums.TableType

	switch monitorType {
	case enums.MonitorTypeStatic:
		tablesToParse = make([]enums.TableType, 0, len(c.selectedTables))

		for _, table := range c.selectedTables {
			if _, ok := rpcTables[table]; !ok {
				tablesToParse = append(tablesToParse, table)
			}
		}
	case enums.MonitorTypeDynamic:
		tablesToParse = []enums.TableType{c.selectedDashboard}
	}

	return tablesToParse
}

// checkErrors collects errors from the channel and returns the first non-nil error.
func checkErrors(errChan <-chan error) error {
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
	if err := c.getTableData(enums.TableTypeRPC); err != nil {
		return err
	}

	if len(c.hosts.rpc) == 0 {
		return errors.New("no public RPC provided")
	}

	if err := c.sortHosts(enums.TableTypeRPC); err != nil {
		return err
	}

	return c.setHostsHealth(enums.TableTypeRPC)
}

// getTableData fetches the data for the specified table type.
// It uses a progress bar to indicate the progress of the data fetching process.
// If the table type is 'Releases', it processes the releases data.
// For other table types, it processes the data accordingly.
// The function returns an error if there is an issue fetching or processing the data.
func (c *Controller) getTableData(tableType enums.TableType) error {
	progressChan := progress.NewProgressBar("PARSING DATA FOR "+string(tableType), progress.ColorBlue)
	defer func() { progressChan <- struct{}{} }()

	if tableType == enums.TableTypeReleases {
		return c.processReleases()
	}

	return c.processStandardTableTypes(tableType)
}

// processReleases fetches the release data for the current network.
// It stores the fetched releases in the Controller's state and returns any error encountered during the process.
func (c *Controller) processReleases() error {
	releases, err := domainmetrics.GetReleases(c.network)
	if err != nil {
		return fmt.Errorf("error getting releases: %w", err)
	}

	c.releases = releases

	return nil
}

// processStandardTableTypes fetches the data for the specified table type other than 'Releases'.
// It retrieves the address information based on the table type, creates hosts, sets the hosts by table type, and sets their health status.
// The function returns an error if there is an issue fetching the address information, creating hosts, setting hosts by table type, or setting their health status.
func (c *Controller) processStandardTableTypes(tableType enums.TableType) error {
	addresses, err := c.getAddressInfoByTableType(tableType)
	if err != nil {
		return fmt.Errorf("error getting address info: %w", err)
	}

	hosts, err := c.createHosts(tableType, addresses)
	if err != nil {
		return fmt.Errorf("error creating hosts: %w", err)
	}

	if err = c.setHostsByTableType(tableType, hosts); err != nil {
		return fmt.Errorf("error setting hosts by table type: %w", err)
	}

	return c.setHostsHealth(tableType)
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
		return fmt.Errorf("error fetching hosts for table %s: %w", tableType, err)
	}

	rpcHost := c.hosts.rpc[0]

	for idx := range hosts {
		metrics := hosts[idx].Metrics

		checkpointExecBacklog := metrics.HighestKnownCheckpoint - metrics.LastExecutedCheckpoint
		checkpointSyncBacklog := metrics.HighestKnownCheckpoint - metrics.HighestSyncedCheckpoint

		// Set transaction sync percentage.
		if setPctProgressErr := hosts[idx].SetPctProgress(enums.MetricTypeTxSyncPercentage, &rpcHost); setPctProgressErr != nil {
			return fmt.Errorf("error setting transaction sync percentage for host: %w", setPctProgressErr)
		}

		// Set checkpoint sync percentage.
		if setPctProgressErr := hosts[idx].SetPctProgress(enums.MetricTypeCheckSyncPercentage, &rpcHost); setPctProgressErr != nil {
			return fmt.Errorf("error setting checkpoint sync percentage for host: %w", setPctProgressErr)
		}

		// Set checkpoint execution backlog.
		if setCheckpointExecBacklogErr := hosts[idx].Metrics.SetValue(enums.MetricTypeCheckpointExecBacklog, checkpointExecBacklog); setCheckpointExecBacklogErr != nil {
			return fmt.Errorf("error setting checkpoint execution backlog for host: %w", setCheckpointExecBacklogErr)
		}

		// Set checkpoint sync backlog.
		if setCheckpointSyncBacklogErr := hosts[idx].Metrics.SetValue(enums.MetricTypeCheckpointSyncBacklog, checkpointSyncBacklog); setCheckpointSyncBacklogErr != nil {
			return fmt.Errorf("error setting checkpoint sync backlog for host: %w", setCheckpointSyncBacklogErr)
		}

		// Set host status.
		hosts[idx].SetStatus(&rpcHost)
	}

	return nil
}
