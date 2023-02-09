package checker

import (
	"github.com/bartosian/sui_helpers/sui-monitor/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/sui-monitor/cmd/checker/tablebuilder"
	"github.com/bartosian/sui_helpers/sui-monitor/cmd/checker/tablebuilder/tables"
)

func (checker *Checker) GenerateNodeTable() {
	tableConfig := tablebuilder.TableConfig{
		Name:         tables.GetTableTitleSUI(checker.network, enums.TableTypeNode),
		Tag:          tables.TableTagSUINode,
		Style:        tables.TableStyleSUINode,
		RowsCount:    0,
		ColumnsCount: len(tables.ColumnConfigSUINode),
		SortConfig:   tables.TableSortConfigSUINode,
	}

	columns := make(tablebuilder.Columns, len(tables.ColumnConfigSUINode))

	for idx, config := range tables.ColumnConfigSUINode {
		columns[idx].Config = config
	}

	node := checker.node

	tableConfig.RowsCount++

	columns[tables.ColumnNameSUINodeAddress].SetValue(node.Address)
	columns[tables.ColumnNameSUINodePortRPC].SetValue(node.RpcPort)
	columns[tables.ColumnNameSUINodePortMetrics].SetValue(node.MetricsPort)
	columns[tables.ColumnNameSUINodeTotalTransactions].SetValue(node.Metrics.TotalTransactionNumber)
	columns[tables.ColumnNameSUINodeHighestCheckpoints].SetValue(node.Metrics.HighestSyncedCheckpoint)
	columns[tables.ColumnNameSUINodeConnectedPeers].SetValue(node.Metrics.SuiNetworkPeers)
	columns[tables.ColumnNameSUINodeUptime].SetValue(node.Metrics.Uptime)
	columns[tables.ColumnNameSUINodeVersion].SetValue(node.Metrics.Version)
	columns[tables.ColumnNameSUINodeCommit].SetValue(node.Metrics.Commit)
	columns[tables.ColumnNameSUINodeCountry].SetValue(node.Location.String())

	tableConfig.Columns = columns

	checker.tableBuilderNode = tablebuilder.NewTableBuilder(tableConfig)
}