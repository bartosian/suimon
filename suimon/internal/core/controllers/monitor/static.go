package monitor

import (
	"fmt"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/tablebuilder"
)

// Static is a method of the Controller struct, responsible for initializing and rendering tables
// based on the configuration data.
func (c *Controller) Static() error {

	// Parse the configuration data.
	if err := c.ParseConfigData(); err != nil {
		return err
	}

	// Initialize tables based on the configuration data.
	if err := c.InitTables(); err != nil {
		return err
	}

	// Render the tables.
	if err := c.RenderTables(); err != nil {
		return err
	}

	// If everything ran successfully, return nil.
	return nil
}

// InitTables initializes the enabled tables based on the display configuration.
// It retrieves the corresponding hosts for each table and initializes the table builder.
// If an error occurs during table initialization, it returns an error.
func (c *Controller) InitTables() error {
	enabledTables := map[enums.TableType]bool{
		enums.TableTypeRPC:              c.config.MonitorsConfig.RPCTable.Display,
		enums.TableTypeNode:             c.config.MonitorsConfig.NodeTable.Display,
		enums.TableTypeValidator:        c.config.MonitorsConfig.ValidatorTable.Display,
		enums.TableTypePeers:            c.config.MonitorsConfig.PeersTable.Display,
		enums.TableTypeSystemState:      c.config.MonitorsConfig.SystemStateTable.Display,
		enums.TableTypeValidatorsCounts: c.config.MonitorsConfig.ValidatorsCountsTable.Display,
		enums.TableTypeValidatorsAtRisk: c.config.MonitorsConfig.ValidatorsAtRiskTable.Display,
		enums.TableTypeValidatorReports: c.config.MonitorsConfig.ValidatorReportsTable.Display,
		enums.TableTypeActiveValidators: c.config.MonitorsConfig.ActiveValidatorsTable.Display,
	}

	for table, isEnabled := range enabledTables {
		if !isEnabled {
			continue
		}

		hosts, err := c.getHostsByTableType(table)
		if err != nil {
			return err
		}

		builder := tablebuilder.NewBuilder(table, c.gateways.cli)
		c.builders.static[table] = builder

		err = builder.Init(hosts)
		if err != nil {
			return fmt.Errorf("error initializing table %s: %w", table, err)
		}
	}

	return nil
}
