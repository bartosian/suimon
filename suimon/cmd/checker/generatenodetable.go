package checker

import (
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/tablebuilder"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/tablebuilder/tables"
)

func (checker *Checker) GenerateNodeTable() {
	if !checker.suimonConfig.MonitorsConfig.NodeTable.Display {
		return
	}

	tableConfig := tablebuilder.TableConfig{
		Name:         tables.GetTableTitleSUI(checker.suimonConfig.NetworkType, enums.TableTypeNode),
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

	columns[tables.ColumnNameSUINodeStatus].SetValue(node.Status)
	columns[tables.ColumnNameSUINodeAddress].SetValue(node.Address)
	columns[tables.ColumnNameSUINodePortRPC].SetValue(node.RpcPort)
	columns[tables.ColumnNameSUINodeTotalTransactions].SetValue(node.Metrics.TotalTransactionNumber)
	columns[tables.ColumnNameSUINodeHighestCheckpoints].SetValue(node.Metrics.HighestSyncedCheckpoint)
	columns[tables.ColumnNameSUINodeConnectedPeers].SetValue(node.Metrics.SuiNetworkPeers)
	columns[tables.ColumnNameSUINodeUptime].SetValue(node.Metrics.Uptime)
	columns[tables.ColumnNameSUINodeVersion].SetValue(node.Metrics.Version)
	columns[tables.ColumnNameSUINodeCommit].SetValue(node.Metrics.Commit)

	if checker.suimonConfig.HostLookupConfig.EnableLookup {
		columns[tables.ColumnNameSUINodeCountry].SetValue(node.Location.CountryName)
	} else {
		columns[tables.ColumnNameSUINodeCountry].SetValue(nil)
	}

	tableConfig.Columns = columns

	checker.tableBuilderNode = tablebuilder.NewTableBuilder(tableConfig)
}
