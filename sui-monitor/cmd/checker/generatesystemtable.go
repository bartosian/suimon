package checker

import (
	"github.com/bartosian/sui_helpers/sui-monitor/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/sui-monitor/cmd/checker/tablebuilder"
	"github.com/bartosian/sui_helpers/sui-monitor/cmd/checker/tablebuilder/tables"
)

func (checker *Checker) GenerateSystemTable() {
	tableConfig := tablebuilder.TableConfig{
		Name:         tables.GetTableTitleSUI(checker.network, enums.TableTypeRPC),
		Tag:          tables.TableTagSystemSUI,
		Style:        tables.TableStyleSystemSUI,
		RowsCount:    0,
		ColumnsCount: len(tables.ColumnConfigSystemSUI),
		SortConfig:   tables.TableSortConfigSystemSUI,
	}

	columns := make(tablebuilder.Columns, len(tables.ColumnConfigSystemSUI))

	for idx, config := range tables.ColumnConfigSystemSUI {
		columns[idx].Config = config
	}

	for _, rpc := range checker.rpcList {
		tableConfig.RowsCount++

		columns[tables.ColumnNameSUISystemStatus].SetValue(rpc.Status)
		columns[tables.ColumnNameSUISystemRPC].SetValue(rpc.Address)
		columns[tables.ColumnNameSUISystemTotalTransactions].SetValue(rpc.Metrics.TotalTransactionNumber)
		columns[tables.ColumnNameSUISystemLatestCheckpoint].SetValue(rpc.Metrics.LatestCheckpoint)
	}

	if tableConfig.RowsCount == 0 {
		columns.SetNoDataValue()

		tableConfig.RowsCount++
	}

	tableConfig.Columns = columns

	checker.tableBuilderRPC = tablebuilder.NewTableBuilder(tableConfig)
}
