package monitor

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/cligw"
)

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

	selectedConfigName, err := c.gateways.cli.SelectOne("Which config would you like to use?", configsChoiceList)
	if err != nil {
		c.gateways.cli.Error("failed to parse user selection")

		return err
	}

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

	if selectedMonitorType.Value == string(enums.MonitorTypeStatic) {
		// Select the tables to render.
		tableTypeChoiceList := cligw.NewSimpleSelectChoiceList(
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

			return err
		}

		if len(selectedTableTypes) == 0 {
			c.gateways.cli.Error("no tables selected to render")

			return nil
		}

		tablesToRender := make([]enums.TableType, 0, len(selectedTableTypes))

		for _, selectedTable := range selectedTableTypes {
			tablesToRender = append(tablesToRender, enums.TableType(selectedTable.Label))
		}

		c.selectedConfig = c.configs[selectedConfigName.Value]
		c.selectedTables = tablesToRender

		return c.Static()
	}

	return nil
}
