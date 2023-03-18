package checker

import (
	"fmt"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/tablebuilder"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/tablebuilder/tables"
)

// InitTables initializes the tables of the "Checker" struct instance passed as a pointer receiver.
// Parameters: None.
// Returns: None.
func (checker *Checker) InitTables() {
	displayConfig := checker.suimonConfig.MonitorsConfig

	if displayConfig.RPCTable.Display {
		checker.InitHostsTable(enums.TableTypeRPC)
	}

	if displayConfig.NodeTable.Display {
		checker.InitHostsTable(enums.TableTypeNode)
	}

	if displayConfig.PeersTable.Display {
		checker.InitHostsTable(enums.TableTypePeers)
	}

	if displayConfig.ValidatorsTable.Display {
		checker.InitValidatorsTable(enums.TableTypeValidators)
	}
}

// InitHostsTable initializes a specific table of the "Checker" struct instance passed as a pointer receiver,
// based on the provided "TableType".
// Parameters: tableType: an enums.TableType representing the type of table to initialize.
// Returns: None.
func (checker *Checker) InitHostsTable(tableType enums.TableType) {
	hosts := checker.getHostsByTableType(tableType)
	suimonConfig := checker.suimonConfig

	tableConfig := tablebuilder.TableConfig{
		Name:       tables.GetTableTitleSUI(suimonConfig.Network.NetworkType, tableType, suimonConfig.MonitorsVisual.EnableEmojis),
		Colors:     tablebuilder.GetTableColorsFromString(suimonConfig.MonitorsVisual.ColorScheme),
		Tag:        tables.TableTagSUI,
		Style:      tables.TableStyleSUI,
		RowsCount:  0,
		SortConfig: tables.TableSortConfigSUI,
	}

	lastColumn := enums.ColumnNameCountry
	if tableType == enums.TableTypeRPC {
		lastColumn = enums.ColumnNameLatestCheckpoint
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

		columns[enums.ColumnNameStatus].SetValue(status)
		columns[enums.ColumnNameAddress].SetValue(address)
		columns[enums.ColumnNamePortRPC].SetValue(port)
		columns[enums.ColumnNameTotalTransactions].SetValue(host.Metrics.TotalTransactionNumber)
		columns[enums.ColumnNameLatestCheckpoint].SetValue(host.Metrics.LatestCheckpoint)

		if tableType != enums.TableTypeRPC {
			columns[enums.ColumnNameHighestCheckpoint].SetValue(host.Metrics.HighestSyncedCheckpoint)
			columns[enums.ColumnNameTXSyncProgress].SetValue(fmt.Sprintf("%v%%", host.Metrics.TxSyncPercentage))
			columns[enums.ColumnNameCheckSyncProgress].SetValue(fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage))
			columns[enums.ColumnNameConnectedPeers].SetValue(host.Metrics.SuiNetworkPeers)
			columns[enums.ColumnNameUptime].SetValue(host.Metrics.Uptime)
			columns[enums.ColumnNameVersion].SetValue(host.Metrics.Version)
			columns[enums.ColumnNameCommit].SetValue(host.Metrics.Commit)

			if host.Location == nil {
				columns[enums.ColumnNameCompany].SetValue(nil)
				columns[enums.ColumnNameCountry].SetValue(nil)

				continue
			}

			columns[enums.ColumnNameCompany].SetValue(host.Location.Provider)

			var country any = host.Location.String()
			if !emojisEnabled {
				country = host.Location.CountryName
			}

			columns[enums.ColumnNameCountry].SetValue(country)
		}
	}

	tableConfig.Columns = columns

	checker.setBuilderTableType(tableType, tableConfig)
}

// InitValidatorsTable initializes a validators table of the "Checker" struct instance passed as a pointer receiver.
// Parameters: None.
// Returns: None.
func (checker *Checker) InitValidatorsTable() {
	suimonConfig := checker.suimonConfig

	tableConfig := tablebuilder.TableConfig{
		Name:       tables.GetTableTitleSUI(suimonConfig.Network.NetworkType, tableType, suimonConfig.MonitorsVisual.EnableEmojis),
		Colors:     tablebuilder.GetTableColorsFromString(suimonConfig.MonitorsVisual.ColorScheme),
		Tag:        tables.TableTagSUI,
		Style:      tables.TableStyleSUI,
		RowsCount:  0,
		SortConfig: tables.TableSortConfigSUI,
	}

	lastColumn := enums.ColumnNameCountry
	if tableType == enums.TableTypeRPC {
		lastColumn = enums.ColumnNameLatestCheckpoint
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

		columns[enums.ColumnNameStatus].SetValue(status)
		columns[enums.ColumnNameAddress].SetValue(address)
		columns[enums.ColumnNamePortRPC].SetValue(port)
		columns[enums.ColumnNameTotalTransactions].SetValue(host.Metrics.TotalTransactionNumber)
		columns[enums.ColumnNameLatestCheckpoint].SetValue(host.Metrics.LatestCheckpoint)

		if tableType != enums.TableTypeRPC {
			columns[enums.ColumnNameHighestCheckpoint].SetValue(host.Metrics.HighestSyncedCheckpoint)
			columns[enums.ColumnNameTXSyncProgress].SetValue(fmt.Sprintf("%v%%", host.Metrics.TxSyncPercentage))
			columns[enums.ColumnNameCheckSyncProgress].SetValue(fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage))
			columns[enums.ColumnNameConnectedPeers].SetValue(host.Metrics.SuiNetworkPeers)
			columns[enums.ColumnNameUptime].SetValue(host.Metrics.Uptime)
			columns[enums.ColumnNameVersion].SetValue(host.Metrics.Version)
			columns[enums.ColumnNameCommit].SetValue(host.Metrics.Commit)

			if host.Location == nil {
				columns[enums.ColumnNameCompany].SetValue(nil)
				columns[enums.ColumnNameCountry].SetValue(nil)

				continue
			}

			columns[enums.ColumnNameCompany].SetValue(host.Location.Provider)

			var country any = host.Location.String()
			if !emojisEnabled {
				country = host.Location.CountryName
			}

			columns[enums.ColumnNameCountry].SetValue(country)
		}
	}

	tableConfig.Columns = columns

	checker.setBuilderTableType(tableType, tableConfig)
}
