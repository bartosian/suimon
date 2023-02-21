package checker

import (
	"fmt"
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
		Tag:        tables.TableTagSUI,
		Style:      tables.TableStyleSUI,
		RowsCount:  0,
		SortConfig: tables.TableSortConfigSUI,
	}

	lastColumn := tables.ColumnNameCountry
	if tableType == enums.TableTypeRPC {
		lastColumn = tables.ColumnNameLatestCheckpoint
	}

	columnsCount := int(lastColumn) + 1
	columns := make(tablebuilder.Columns, columnsCount)
	emojisEnabled := checker.suimonConfig.MonitorsVisual.EnableEmojis

	for i := 0; i < columnsCount; i++ {
		columns[i].Config = tables.ColumnConfigSUI[i]
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

		columns[tables.ColumnNameStatus].SetValue(status)
		columns[tables.ColumnNameAddress].SetValue(address)
		columns[tables.ColumnNamePortRPC].SetValue(port)
		columns[tables.ColumnNameTotalTransactions].SetValue(host.Metrics.TotalTransactionNumber)
		columns[tables.ColumnNameLatestCheckpoint].SetValue(host.Metrics.LatestCheckpoint)

		if tableType != enums.TableTypeRPC {
			columns[tables.ColumnNameHighestCheckpoint].SetValue(host.Metrics.HighestSyncedCheckpoint)
			columns[tables.ColumnNameTXSyncProgress].SetValue(fmt.Sprintf("%v%%", host.Metrics.TxSyncPercentage))
			columns[tables.ColumnNameCheckSyncProgress].SetValue(fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage))
			columns[tables.ColumnNameConnectedPeers].SetValue(host.Metrics.SuiNetworkPeers)
			columns[tables.ColumnNameUptime].SetValue(host.Metrics.Uptime)
			columns[tables.ColumnNameVersion].SetValue(host.Metrics.Version)
			columns[tables.ColumnNameCommit].SetValue(host.Metrics.Commit)

			if host.Location == nil {
				columns[tables.ColumnNameCompany].SetValue(nil)
				columns[tables.ColumnNameCountry].SetValue(nil)

				continue
			}

			columns[tables.ColumnNameCompany].SetValue(host.Location.Provider)

			var country any = host.Location.String()
			if !emojisEnabled {
				country = host.Location.CountryName
			}

			columns[tables.ColumnNameCountry].SetValue(country)
		}
	}

	tableConfig.Columns = columns

	checker.setTableBuilderTableType(tableType, tableConfig)
}
