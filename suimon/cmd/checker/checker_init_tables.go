package checker

import (
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/tablebuilder"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/tablebuilder/tables"
)

func (checker *Checker) InitTables() {
	displayConfig := checker.suimonConfig.MonitorsConfig

	if displayConfig.RPCTable.Display {
		checker.InitTable(enums.TableTypeRPC)
	}

	if displayConfig.NodeTable.Display {
		checker.InitTable(enums.TableTypeNode)
	}

	if displayConfig.PeersTable.Display {
		checker.InitTable(enums.TableTypePeers)
	}
}

func (checker *Checker) InitTable(tableType enums.TableType) {
	hosts := checker.getHostsByTableType(tableType)
	suimonConfig := checker.suimonConfig

	tableConfig := tablebuilder.TableConfig{
		Name:       tables.GetTableTitleSUI(suimonConfig.NetworkType, tableType, suimonConfig.MonitorsVisual.EnableEmojis),
		Colors:     tablebuilder.GetTableColorsFromString(suimonConfig.MonitorsVisual.ColorScheme),
		Tag:        tables.TableTagSUINode,
		Style:      tables.TableStyleSUINode,
		RowsCount:  0,
		SortConfig: tables.TableSortConfigSUINode,
	}

	lastColumn := tables.ColumnNameSUINodeCountry
	if tableType == enums.TableTypeRPC {
		lastColumn = tables.ColumnNameSUINodeLatestCheckpoint
	}

	columnsCount := int(lastColumn) + 1
	columns := make(tablebuilder.Columns, columnsCount)
	emojisEnabled := checker.suimonConfig.MonitorsVisual.EnableEmojis

	for i := 0; i < columnsCount; i++ {
		columns[i].Config = tables.ColumnConfigSUINode[i]
	}

	for _, host := range hosts {
		if !host.Metrics.Updated && tableType == enums.TableTypePeers {
			continue
		}

		tableConfig.RowsCount++

		var status any = host.Status
		if !emojisEnabled {
			status = host.Status.StatusToPlaceholder()
		}

		port := host.Ports[enums.PortTypeRPC]
		if port == "" {
			port = rpcPortDefault
		}

		address := host.HostPort.Address
		if tableType == enums.TableTypeNode {
			address = *host.HostPort.IP
		}

		columns[tables.ColumnNameSUINodeStatus].SetValue(status)
		columns[tables.ColumnNameSUINodeAddress].SetValue(address)
		columns[tables.ColumnNameSUINodePortRPC].SetValue(port)
		columns[tables.ColumnNameSUINodeTotalTransactions].SetValue(host.Metrics.TotalTransactionNumber)
		columns[tables.ColumnNameSUINodeLatestCheckpoint].SetValue(host.Metrics.LatestCheckpoint)

		if tableType != enums.TableTypeRPC {
			columns[tables.ColumnNameSUINodeHighestCheckpoint].SetValue(host.Metrics.HighestSyncedCheckpoint)
			columns[tables.ColumnNameSUINodeConnectedPeers].SetValue(host.Metrics.SuiNetworkPeers)
			columns[tables.ColumnNameSUINodeUptime].SetValue(host.Metrics.Uptime)
			columns[tables.ColumnNameSUINodeVersion].SetValue(host.Metrics.Version)
			columns[tables.ColumnNameSUINodeCommit].SetValue(host.Metrics.Commit)

			if host.Location == nil {
				columns[tables.ColumnNameSUINodeCompany].SetValue(nil)
				columns[tables.ColumnNameSUINodeCountry].SetValue(nil)

				continue
			}

			columns[tables.ColumnNameSUINodeCompany].SetValue(host.Location.Provider)

			var country any = host.Location.String()
			if !emojisEnabled {
				country = host.Location.CountryName
			}

			columns[tables.ColumnNameSUINodeCountry].SetValue(country)
		}
	}

	tableConfig.Columns = columns

	checker.setTableBuilderTableType(tableType, tableConfig)
}
