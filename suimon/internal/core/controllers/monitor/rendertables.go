package monitor

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
)

// RenderTables renders the selected tables. The function checks whether data has been provided for each table
// and enables or disables the table based on the availability of data. For each selected table, the function
// retrieves the corresponding table builder from the static table builders map and calls its Render method.
// The function returns nil if all selected tables have been rendered successfully.
func (c *Controller) RenderTables() error {
	selectedTables := c.selectedTables

	rpcProvided := len(c.hosts.rpc) > 0
	nodeProvided := len(c.hosts.node) > 0
	peersProvided := len(c.hosts.peers) > 0
	validatorProvided := len(c.hosts.validator) > 0

	tableTypeEnabled := map[enums.TableType]bool{
		enums.TableTypeRPC:              rpcProvided,
		enums.TableTypeNode:             nodeProvided,
		enums.TableTypeValidator:        validatorProvided,
		enums.TableTypePeers:            peersProvided,
		enums.TableTypeSystemState:      rpcProvided,
		enums.TableTypeValidatorsCounts: rpcProvided,
		enums.TableTypeValidatorsAtRisk: rpcProvided,
		enums.TableTypeValidatorReports: rpcProvided,
		enums.TableTypeActiveValidators: rpcProvided,
	}

	for _, tableType := range selectedTables {
		if !tableTypeEnabled[tableType] {
			continue
		}

		builder := c.builders.static[tableType]

		builder.Render()
	}

	return nil
}