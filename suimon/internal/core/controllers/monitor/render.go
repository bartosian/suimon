package monitor

import (
	"fmt"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
)

// RenderTables renders the tables for the hosts data provided to the controller.
// It iterates over the table types and corresponding builders in the tableTypeToBuilder map
// and renders the tables if they are enabled in the configuration and the data is provided for the table type.
// If rendering fails for any table, an error is returned.
func (c *Controller) RenderTables() error {
	monitorsConfig := c.config.MonitorsConfig

	rpcProvided := len(c.hosts.rpc) > 0
	nodeProvided := len(c.hosts.node) > 0
	peersProvided := len(c.hosts.peers) > 0
	validatorProvided := len(c.hosts.validator) > 0

	tableTypeToBuilder := map[enums.TableType]struct {
		builder ports.Builder
		enabled bool
	}{
		enums.TableTypeRPC: {
			builder: c.builders.static[enums.TableTypeRPC],
			enabled: monitorsConfig.RPCTable.Display && rpcProvided,
		},
		enums.TableTypeNode: {
			builder: c.builders.static[enums.TableTypeNode],
			enabled: monitorsConfig.NodeTable.Display && nodeProvided,
		},
		enums.TableTypeValidator: {
			builder: c.builders.static[enums.TableTypeValidator],
			enabled: monitorsConfig.ValidatorTable.Display && validatorProvided,
		},
		enums.TableTypePeers: {
			builder: c.builders.static[enums.TableTypePeers],
			enabled: monitorsConfig.PeersTable.Display && peersProvided,
		},
		enums.TableTypeSystemState: {
			builder: c.builders.static[enums.TableTypeSystemState],
			enabled: monitorsConfig.SystemStateTable.Display && rpcProvided,
		},
		enums.TableTypeValidatorsCounts: {
			builder: c.builders.static[enums.TableTypeValidatorsCounts],
			enabled: monitorsConfig.ValidatorsCountsTable.Display && rpcProvided,
		},
		enums.TableTypeValidatorsAtRisk: {
			builder: c.builders.static[enums.TableTypeValidatorsAtRisk],
			enabled: monitorsConfig.ValidatorsAtRiskTable.Display && rpcProvided,
		},
		enums.TableTypeValidatorReports: {
			builder: c.builders.static[enums.TableTypeValidatorReports],
			enabled: monitorsConfig.ValidatorReportsTable.Display && rpcProvided,
		},
		enums.TableTypeActiveValidators: {
			builder: c.builders.static[enums.TableTypeActiveValidators],
			enabled: monitorsConfig.ActiveValidatorsTable.Display && rpcProvided,
		},
	}

	for tableType, builderConfig := range tableTypeToBuilder {
		if !builderConfig.enabled {
			continue
		}

		if err := builderConfig.builder.Render(); err != nil {
			return fmt.Errorf("error rendering table %s: %w", tableType, err)
		}
	}

	return nil
}
