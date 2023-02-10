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
		Name:         tables.GetTableTitleSUI(checker.suimonConfig.NetworkType, enums.TableTypePeers, checker.suimonConfig.MonitorsVisual.EnableEmojis),
		Tag:          tables.TableTagSUINode,
		Colors:       tablebuilder.GetTBColorsFromString(checker.suimonConfig.MonitorsVisual.ColorScheme),
		Style:        tables.TableStyleSUINode,
		RowsCount:    0,
		ColumnsCount: len(tables.ColumnConfigSUINode),
		SortConfig:   tables.TableSortConfigSUINode,
	}

	columns := make(tablebuilder.Columns, len(tables.ColumnConfigSUINode))

	for idx, config := range tables.ColumnConfigSUINode {
		columns[idx].Config = config
	}

	tableConfig.RowsCount++

	node := checker.node
	emojisEnabled := checker.suimonConfig.MonitorsVisual.EnableEmojis

	if emojisEnabled {
		columns[tables.ColumnNameSUINodeStatus].SetValue(node.Status)
	} else {
		columns[tables.ColumnNameSUINodeStatus].SetValue(node.Status.StatusToPlaceholder())
	}

	columns[tables.ColumnNameSUINodeAddress].SetValue(node.Address)
	columns[tables.ColumnNameSUINodePortRPC].SetValue(node.RpcPort)
	columns[tables.ColumnNameSUINodeTotalTransactions].SetValue(node.Metrics.TotalTransactionNumber)
	columns[tables.ColumnNameSUINodeHighestCheckpoints].SetValue(node.Metrics.HighestSyncedCheckpoint)
	columns[tables.ColumnNameSUINodeConnectedPeers].SetValue(node.Metrics.SuiNetworkPeers)
	columns[tables.ColumnNameSUINodeUptime].SetValue(node.Metrics.Uptime)
	columns[tables.ColumnNameSUINodeVersion].SetValue(node.Metrics.Version)
	columns[tables.ColumnNameSUINodeCommit].SetValue(node.Metrics.Commit)

	if node.Location == nil {
		columns[tables.ColumnNameSUINodeCompany].SetValue(nil)
		columns[tables.ColumnNameSUINodeCountry].SetValue(nil)
	} else {
		columns[tables.ColumnNameSUINodeCompany].SetValue(node.Location.Provider)

		if emojisEnabled {
			columns[tables.ColumnNameSUINodeCountry].SetValue(node.Location.String())
		} else {
			columns[tables.ColumnNameSUINodeCountry].SetValue(node.Location.CountryName)
		}
	}

	tableConfig.Columns = columns

	checker.tableBuilderNode = tablebuilder.NewTableBuilder(tableConfig)
}
