package monitor

import (
	"fmt"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/core/domain/service/tablebuilder"
)

// Static is a method of the Controller struct, responsible for initializing and rendering tables
// based on the configuration data.
func (c *Controller) Static() error {
	// Parse the configuration data.
	if err := c.ParseConfigData(enums.MonitorTypeStatic); err != nil {
		return err
	}

	// Initialize tables based on the configuration data.
	if err := c.InitTables(); err != nil {
		return err
	}

	// Render the tables and return error if any
	return c.RenderTables()
}

// InitTables initializes the enabled tables based on the display configuration.
// It retrieves the corresponding hosts for each table and initializes the table builder.
// If an error occurs during table initialization, it returns an error.
func (c *Controller) InitTables() error {
	selectedTables := c.selectedTables

	for _, tableType := range selectedTables {
		hosts, err := c.getHostsByTableType(tableType)
		if err != nil {
			return err
		}

		if len(hosts) == 0 {
			continue
		}

		builder := tablebuilder.NewBuilder(tableType, hosts, c.gateways.cli)
		c.builders.static[tableType] = builder

		if err = builder.Init(); err != nil {
			return fmt.Errorf("error initializing table %s: %w", tableType, err)
		}
	}

	return nil
}
