package checker

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"

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
		checker.InitTable(enums.TableTypeRPC)
	}

	if displayConfig.NodeTable.Display {
		checker.InitTable(enums.TableTypeNode)
	}

	if displayConfig.PeersTable.Display {
		checker.InitTable(enums.TableTypePeers)
	}
}

// InitTable initializes a specific table of the "Checker" struct instance passed as a pointer receiver,
// based on the provided "TableType".
// Parameters: tableType: an enums.TableType representing the type of table to initialize.
// Returns: None.
func (checker *Checker) InitTable(tableType enums.TableType) {
	networkType := checker.suimonConfig.Network.NetworkType
	enabledEmojis := checker.suimonConfig.MonitorsVisual.EnableEmojis
	colorScheme := checker.suimonConfig.MonitorsVisual.ColorScheme

	var (
		hosts        []Host
		tableConfig  tablebuilder.TableConfig
		columnConfig []table.ColumnConfig
		columns      tablebuilder.Columns
	)

	switch tableType {
	case enums.TableTypeNode:
		hosts = checker.getHostsByTableType(enums.TableTypeNode)
		columnConfig = tables.ColumnConfigNode
		columns = make(tablebuilder.Columns, len(columnConfig))

		tableConfig = tablebuilder.TableConfig{
			Name:       tables.GetTableTitleSUI(networkType, enums.TableTypeNode, enabledEmojis),
			Colors:     tablebuilder.GetTableColorsFromString(colorScheme),
			Tag:        tables.TableTagNode,
			Style:      tables.TableStyleNode,
			RowsCount:  0,
			Columns:    columns,
			SortConfig: tables.TableSortConfigNode,
		}

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		for _, host := range hosts {
			if !host.Metrics.Updated {
				continue
			}

			tableConfig.RowsCount++

			var (
				status  any = host.Status
				country any = host.Location.String()
				port        = host.Ports[enums.PortTypeRPC]
				address     = *host.HostPort.IP
			)

			if !enabledEmojis {
				status = host.Status.StatusToPlaceholder()
			}

			if port == "" {
				port = rpcPortDefault
			}

			columns[enums.ColumnNameHealth].SetValue(status)
			columns[enums.ColumnNameAddress].SetValue(address)
			columns[enums.ColumnNamePortRPC].SetValue(port)
			columns[enums.ColumnNameTotalTransactions].SetValue(host.Metrics.TotalTransactions)
			columns[enums.ColumnNameTotalTransactionCertificates].SetValue(host.Metrics.TotalTransactionCertificates)
			columns[enums.ColumnNameTotalTransactionEffects].SetValue(host.Metrics.TotalTransactionEffects)
			columns[enums.ColumnNameLatestCheckpoint].SetValue(host.Metrics.LatestCheckpoint)
			columns[enums.ColumnNameHighestKnownCheckpoint].SetValue(host.Metrics.HighestKnownCheckpoint)
			columns[enums.ColumnNameHighestSyncedCheckpoint].SetValue(host.Metrics.HighestSyncedCheckpoint)
			columns[enums.ColumnNameLastExecutedCheckpoint].SetValue(host.Metrics.LastExecutedCheckpoint)
			columns[enums.ColumnNameCheckpointExecBacklog].SetValue(host.Metrics.CheckpointExecBacklog)
			columns[enums.ColumnNameCheckpointSyncBacklog].SetValue(host.Metrics.CheckpointSyncBacklog)
			columns[enums.ColumnNameCurrentEpoch].SetValue(host.Metrics.CurrentEpoch)
			columns[enums.ColumnNameTXSyncPercentage].SetValue(fmt.Sprintf("%v%%", host.Metrics.TxSyncPercentage))
			columns[enums.ColumnNameCheckSyncPercentage].SetValue(fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage))
			columns[enums.ColumnNameNetworkPeers].SetValue(host.Metrics.NetworkPeers)
			columns[enums.ColumnNameUptime].SetValue(host.Metrics.Uptime)
			columns[enums.ColumnNameVersion].SetValue(host.Metrics.Version)
			columns[enums.ColumnNameCommit].SetValue(host.Metrics.Commit)

			if host.Location == nil {
				columns[enums.ColumnNameCountry].SetValue(nil)

				continue
			}

			if !enabledEmojis {
				country = host.Location.CountryName
			}

			columns[enums.ColumnNameCountry].SetValue(country)
		}
	case enums.TableTypePeers:
		hosts = checker.getHostsByTableType(enums.TableTypePeers)
		columnConfig = tables.ColumnConfigPeer
		columns = make(tablebuilder.Columns, len(columnConfig))

		tableConfig = tablebuilder.TableConfig{
			Name:       tables.GetTableTitleSUI(networkType, enums.TableTypePeers, enabledEmojis),
			Colors:     tablebuilder.GetTableColorsFromString(colorScheme),
			Tag:        tables.TableTagPeer,
			Style:      tables.TableStylePeer,
			RowsCount:  0,
			Columns:    columns,
			SortConfig: tables.TableSortConfigPeer,
		}

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		for _, host := range hosts {
			if !host.Metrics.Updated {
				continue
			}

			tableConfig.RowsCount++

			var (
				status  any = host.Status
				country any = host.Location.String()
				port        = host.Ports[enums.PortTypeRPC]
				address     = host.HostPort.Address
			)

			if !enabledEmojis {
				status = host.Status.StatusToPlaceholder()
			}

			if port == "" {
				port = rpcPortDefault
			}

			columns[enums.ColumnNameHealth].SetValue(status)
			columns[enums.ColumnNameAddress].SetValue(address)
			columns[enums.ColumnNamePortRPC].SetValue(port)
			columns[enums.ColumnNameTotalTransactions].SetValue(host.Metrics.TotalTransactions)
			columns[enums.ColumnNameTotalTransactionCertificates].SetValue(host.Metrics.TotalTransactionCertificates)
			columns[enums.ColumnNameTotalTransactionEffects].SetValue(host.Metrics.TotalTransactionEffects)
			columns[enums.ColumnNameLatestCheckpoint].SetValue(host.Metrics.LatestCheckpoint)
			columns[enums.ColumnNameHighestKnownCheckpoint].SetValue(host.Metrics.HighestKnownCheckpoint)
			columns[enums.ColumnNameHighestSyncedCheckpoint].SetValue(host.Metrics.HighestSyncedCheckpoint)
			columns[enums.ColumnNameLastExecutedCheckpoint].SetValue(host.Metrics.LastExecutedCheckpoint)
			columns[enums.ColumnNameCheckpointExecBacklog].SetValue(host.Metrics.CheckpointExecBacklog)
			columns[enums.ColumnNameCheckpointSyncBacklog].SetValue(host.Metrics.CheckpointSyncBacklog)
			columns[enums.ColumnNameCurrentEpoch].SetValue(host.Metrics.CurrentEpoch)
			columns[enums.ColumnNameTXSyncPercentage].SetValue(fmt.Sprintf("%v%%", host.Metrics.TxSyncPercentage))
			columns[enums.ColumnNameCheckSyncPercentage].SetValue(fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage))
			columns[enums.ColumnNameNetworkPeers].SetValue(host.Metrics.NetworkPeers)
			columns[enums.ColumnNameUptime].SetValue(host.Metrics.Uptime)
			columns[enums.ColumnNameVersion].SetValue(host.Metrics.Version)
			columns[enums.ColumnNameCommit].SetValue(host.Metrics.Commit)

			if host.Location == nil {
				columns[enums.ColumnNameCountry].SetValue(nil)

				continue
			}

			if !enabledEmojis {
				country = host.Location.CountryName
			}

			columns[enums.ColumnNameCountry].SetValue(country)
		}
	case enums.TableTypeRPC:
		hosts = checker.getHostsByTableType(enums.TableTypeRPC)
		columnConfig = tables.ColumnConfigRPC
		columns = make(tablebuilder.Columns, len(columnConfig))

		tableConfig = tablebuilder.TableConfig{
			Name:       tables.GetTableTitleSUI(networkType, enums.TableTypeRPC, enabledEmojis),
			Colors:     tablebuilder.GetTableColorsFromString(colorScheme),
			Tag:        tables.TableTagRPC,
			Style:      tables.TableStyleRPC,
			RowsCount:  0,
			Columns:    columns,
			SortConfig: tables.TableSortConfigRPC,
		}

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		for _, host := range hosts {
			if !host.Metrics.Updated {
				continue
			}

			tableConfig.RowsCount++

			var (
				status  any = host.Status
				port        = host.Ports[enums.PortTypeRPC]
				address     = host.HostPort.Address
			)

			if !enabledEmojis {
				status = host.Status.StatusToPlaceholder()
			}

			if port == "" {
				port = rpcPortDefault
			}

			columns[enums.ColumnNameHealth].SetValue(status)
			columns[enums.ColumnNameAddress].SetValue(address)
			columns[enums.ColumnNamePortRPC].SetValue(port)
			columns[enums.ColumnNameTotalTransactions].SetValue(host.Metrics.TotalTransactions)
			columns[enums.ColumnNameLatestCheckpoint].SetValue(host.Metrics.LatestCheckpoint)
		}
	case enums.TableTypeValidators:
	}

	checker.setBuilderTableType(tableType, tableConfig)
}
