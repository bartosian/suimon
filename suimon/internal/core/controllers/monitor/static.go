package monitor

import (
	"fmt"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
)

func (c *Controller) Static() error {
	return nil
}

// InitTables initializes the enabled tables based on the display configuration.
// It retrieves the corresponding hosts for each table and initializes the table builder.
// If an error occurs during table initialization, it returns an error.
func (c *Controller) InitTables() error {
	displayConfig := c.config.MonitorsConfig
	tableConfigMap := map[enums.TableType]bool{
		enums.TableTypeRPC:              displayConfig.RPCTable.Display,
		enums.TableTypeNode:             displayConfig.NodeTable.Display,
		enums.TableTypeValidator:        displayConfig.ValidatorTable.Display,
		enums.TableTypePeers:            displayConfig.PeersTable.Display,
		enums.TableTypeSystemState:      displayConfig.SystemStateTable.Display,
		enums.TableTypeValidatorsCounts: displayConfig.ValidatorsCountsTable.Display,
		enums.TableTypeValidatorsAtRisk: displayConfig.ValidatorsAtRiskTable.Display,
		enums.TableTypeValidatorReports: displayConfig.ValidatorReportsTable.Display,
		enums.TableTypeActiveValidators: displayConfig.ActiveValidatorsTable.Display,
	}

	for table, isEnabled := range tableConfigMap {
		if !isEnabled {
			continue
		}

		builder := c.builders.static[table]

		hosts, err := c.getHostsByTableType(table)
		if err != nil {
			return err
		}

		err = builder.Init(table, hosts)
		if err != nil {
			return fmt.Errorf("error initializing table %s: %w", table, err)
		}
	}

	return nil
}
