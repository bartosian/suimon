package monitor

import (
	"fmt"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
)

func (c *Controller) Static() error {
	return nil
}

func (c *Controller) initTables() error {
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

	for tableType, shouldDisplay := range tableConfigMap {
		if shouldDisplay {
			if err := c.initTable(tableType); err != nil {
				return fmt.Errorf("error initializing table %s: %w", tableType, err)
			}
		}
	}

	return nil
}
