package monitor

import (
	"fmt"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/cligw"
)

const allTablesSelection = "🌐 ALL TABLES"

// Monitor prompts the user to select the type of monitor to render, and then renders the monitor.
// For static monitors, the user is prompted to select the tables to render. Only tables that are enabled
// in the configuration file are displayed in the list of choices. If no tables are enabled, an error is
// displayed and the function returns without rendering any tables.
func (c *Controller) Monitor() error {
	// List available configurations.
	configNames := make([]string, 0, len(c.configs))

	for configName := range c.configs {
		configNames = append(configNames, configName)
	}

	configsChoiceList := cligw.NewSimpleSelectChoiceList(configNames...)

	selectedConfigName, err := c.gateways.cli.SelectOne("Which dashboards would you like to use?", configsChoiceList)
	if err != nil {
		c.gateways.cli.Error("failed to parse user selection")

		return err
	}

	c.selectedConfig = c.configs[selectedConfigName.Value]

	// Select the monitor type.
	monitorTypeChoiceList := cligw.NewSimpleSelectChoiceList(
		string(enums.MonitorTypeStatic),
		string(enums.MonitorTypeDynamic),
	)

	selectedMonitorType, err := c.gateways.cli.SelectOne("Which monitors would you like to render?", monitorTypeChoiceList)
	if err != nil {
		c.gateways.cli.Error("failed to parse user selection")

		return err
	}

	switch selectedMonitorType.Value {
	case string(enums.MonitorTypeStatic):
		tablesToRender, err := c.selectStaticTables()
		if err != nil {
			return err
		}

		if tablesToRender == nil {
			return nil
		}

		c.selectedTables = tablesToRender

		return c.Static()

	case string(enums.MonitorTypeDynamic):
		dashboardToRender, err := c.selectDynamicDashboard()
		if err != nil {
			return err
		}

		if dashboardToRender == nil {
			return nil
		}

		c.selectedDashboard = *dashboardToRender

		return c.Dynamic()
	default:
		return fmt.Errorf("not supported monitoring type provided %s", selectedMonitorType.Value)
	}
}

// selectStaticTables prompts the user to select the static tables to render.
// It returns a slice of enums.TableType representing the selected tables,
// or an error if the user's selection cannot be parsed or no tables are selected.
func (c *Controller) selectStaticTables() ([]enums.TableType, error) {
	// Select the tables to render.
	tableTypeChoiceList := cligw.NewSimpleSelectChoiceList(
		allTablesSelection,
		string(enums.TableTypeRPC),
		string(enums.TableTypeNode),
		string(enums.TableTypeValidator),
		string(enums.TableTypePeers),
		string(enums.TableTypeGasPriceAndSubsidy),
		string(enums.TableTypeValidatorsCounts),
		string(enums.TableTypeValidatorsAtRisk),
		string(enums.TableTypeValidatorReports),
		string(enums.TableTypeActiveValidators),
	)

	selectedTableTypes, err := c.gateways.cli.SelectMany("Which tables do you want to render?", tableTypeChoiceList)
	if err != nil {
		c.gateways.cli.Error("failed to parse user selection")

		return nil, err
	}

	if len(selectedTableTypes) == 0 {
		c.gateways.cli.Error("no tables selected to render")

		return nil, nil
	}

	tablesToRender := make([]enums.TableType, 0, len(selectedTableTypes))

	for _, selectedTable := range selectedTableTypes {
		if selectedTable.Value == allTablesSelection {
			tablesToRender = append(tablesToRender,
				enums.TableTypeRPC,
				enums.TableTypeNode,
				enums.TableTypeValidator,
				enums.TableTypePeers,
				enums.TableTypeGasPriceAndSubsidy,
				enums.TableTypeValidatorsCounts,
				enums.TableTypeValidatorsAtRisk,
				enums.TableTypeValidatorReports,
				enums.TableTypeActiveValidators,
			)

			break
		}

		tablesToRender = append(tablesToRender, enums.TableType(selectedTable.Value))
	}

	return tablesToRender, nil
}

// selectDynamicDashboard prompts the user to select the dynamic dashboard to render.
// It returns a slice of enums.TableType representing the selected dashboard,
// or an error if the user's selection cannot be parsed or no dashboard is selected.
func (c *Controller) selectDynamicDashboard() (*enums.TableType, error) {
	// Select the dashboard to render.
	dashboardTypeChoiceList := cligw.NewSimpleSelectChoiceList(
		string(enums.TableTypeNode),
		string(enums.TableTypeValidator),
		string(enums.TableTypeRPC),
		string(enums.TableTypeGasPriceAndSubsidy),
	)

	selectedDashboardType, err := c.gateways.cli.SelectOne("Which dashboard do you want to render?", dashboardTypeChoiceList)
	if err != nil {
		c.gateways.cli.Error("failed to parse user selection")

		return nil, err
	}

	if selectedDashboardType == nil {
		c.gateways.cli.Error("no dashboard selected to render")

		return nil, nil
	}

	dashboardType := enums.TableType(selectedDashboardType.Value)

	return &dashboardType, nil
}

// selectHostForDashboard selects a host to render a dashboard for, based on the selected dashboard.
// It prompts the user to select a host from the list of hosts that support the selected dashboard,
// and returns the selected host, or an error if the user's selection cannot be parsed or no host is selected.
func (c *Controller) selectHostForDashboard() (*domainhost.Host, error) {
	selectedDashboard := c.selectedDashboard

	hosts, err := c.getHostsByTableType(selectedDashboard)
	if err != nil {
		return nil, err
	}

	if len(hosts) == 0 {
		return nil, nil
	}

	if len(hosts) == 1 {
		return &hosts[0], nil
	}

	// Create a list of host addresses for the user to select from.
	hostAddresses := make([]string, len(hosts))
	for i, host := range hosts {
		hostAddresses[i] = host.Endpoint.Address
	}

	// Select the host to render.
	hostChoiceList := cligw.NewSimpleSelectChoiceList(hostAddresses...)

	selectedHostAddress, err := c.gateways.cli.SelectOne("Which host do you want to render dashboard for?", hostChoiceList)
	if err != nil {
		c.gateways.cli.Error("failed to parse user selection")

		return nil, err
	}

	if selectedHostAddress == nil {
		c.gateways.cli.Error("no host selected to render for")

		return nil, nil
	}

	hostAddress := selectedHostAddress.Value

	// Get the selected host from the slice of pointers.
	for _, host := range hosts {
		if host.Endpoint.Address == hostAddress {
			return &host, nil
		}
	}

	return nil, fmt.Errorf("selected host not found")
}
