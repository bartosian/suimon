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

	if displayConfig.ValidatorsTable.Display {
		checker.InitTable(enums.TableTypeValidators)
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

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		tableConfig = tablebuilder.TableConfig{
			Name:       tables.GetTableTitle(networkType, enums.TableTypeNode, enabledEmojis),
			Colors:     tablebuilder.GetTableColorsFromString(colorScheme),
			Tag:        tables.TableTagNode,
			Style:      tables.TableStyleNode,
			RowsCount:  0,
			Columns:    columns,
			SortConfig: tables.TableSortConfigNode,
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

			columns[enums.NodeColumnNameHealth].SetValue(status)
			columns[enums.NodeColumnNameAddress].SetValue(address)
			columns[enums.NodeColumnNamePortRPC].SetValue(port)
			columns[enums.NodeColumnNameTotalTransactions].SetValue(host.Metrics.TotalTransactions)
			columns[enums.NodeColumnNameTotalTransactionCertificates].SetValue(host.Metrics.TotalTransactionCertificates)
			columns[enums.NodeColumnNameTotalTransactionEffects].SetValue(host.Metrics.TotalTransactionEffects)
			columns[enums.NodeColumnNameLatestCheckpoint].SetValue(host.Metrics.LatestCheckpoint)
			columns[enums.NodeColumnNameHighestKnownCheckpoint].SetValue(host.Metrics.HighestKnownCheckpoint)
			columns[enums.NodeColumnNameHighestSyncedCheckpoint].SetValue(host.Metrics.HighestSyncedCheckpoint)
			columns[enums.NodeColumnNameLastExecutedCheckpoint].SetValue(host.Metrics.LastExecutedCheckpoint)
			columns[enums.NodeColumnNameCheckpointExecBacklog].SetValue(host.Metrics.CheckpointExecBacklog)
			columns[enums.NodeColumnNameCheckpointSyncBacklog].SetValue(host.Metrics.CheckpointSyncBacklog)
			columns[enums.NodeColumnNameCurrentEpoch].SetValue(host.Metrics.CurrentEpoch)
			columns[enums.NodeColumnNameTXSyncPercentage].SetValue(fmt.Sprintf("%v%%", host.Metrics.TxSyncPercentage))
			columns[enums.NodeColumnNameCheckSyncPercentage].SetValue(fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage))
			columns[enums.NodeColumnNameNetworkPeers].SetValue(host.Metrics.NetworkPeers)
			columns[enums.NodeColumnNameUptime].SetValue(host.Metrics.Uptime)
			columns[enums.NodeColumnNameVersion].SetValue(host.Metrics.Version)
			columns[enums.NodeColumnNameCommit].SetValue(host.Metrics.Commit)

			if host.Location == nil {
				columns[enums.NodeColumnNameCountry].SetValue(nil)

				continue
			}

			if !enabledEmojis {
				country = host.Location.CountryName
			}

			columns[enums.NodeColumnNameCountry].SetValue(country)
		}
	case enums.TableTypePeers:
		hosts = checker.getHostsByTableType(enums.TableTypePeers)
		columnConfig = tables.ColumnConfigPeer
		columns = make(tablebuilder.Columns, len(columnConfig))

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		tableConfig = tablebuilder.TableConfig{
			Name:       tables.GetTableTitle(networkType, enums.TableTypePeers, enabledEmojis),
			Colors:     tablebuilder.GetTableColorsFromString(colorScheme),
			Tag:        tables.TableTagPeer,
			Style:      tables.TableStylePeer,
			RowsCount:  0,
			Columns:    columns,
			SortConfig: tables.TableSortConfigPeer,
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

			columns[enums.NodeColumnNameHealth].SetValue(status)
			columns[enums.NodeColumnNameAddress].SetValue(address)
			columns[enums.NodeColumnNamePortRPC].SetValue(port)
			columns[enums.NodeColumnNameTotalTransactions].SetValue(host.Metrics.TotalTransactions)
			columns[enums.NodeColumnNameTotalTransactionCertificates].SetValue(host.Metrics.TotalTransactionCertificates)
			columns[enums.NodeColumnNameTotalTransactionEffects].SetValue(host.Metrics.TotalTransactionEffects)
			columns[enums.NodeColumnNameLatestCheckpoint].SetValue(host.Metrics.LatestCheckpoint)
			columns[enums.NodeColumnNameHighestKnownCheckpoint].SetValue(host.Metrics.HighestKnownCheckpoint)
			columns[enums.NodeColumnNameHighestSyncedCheckpoint].SetValue(host.Metrics.HighestSyncedCheckpoint)
			columns[enums.NodeColumnNameLastExecutedCheckpoint].SetValue(host.Metrics.LastExecutedCheckpoint)
			columns[enums.NodeColumnNameCheckpointExecBacklog].SetValue(host.Metrics.CheckpointExecBacklog)
			columns[enums.NodeColumnNameCheckpointSyncBacklog].SetValue(host.Metrics.CheckpointSyncBacklog)
			columns[enums.NodeColumnNameCurrentEpoch].SetValue(host.Metrics.CurrentEpoch)
			columns[enums.NodeColumnNameTXSyncPercentage].SetValue(fmt.Sprintf("%v%%", host.Metrics.TxSyncPercentage))
			columns[enums.NodeColumnNameCheckSyncPercentage].SetValue(fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage))
			columns[enums.NodeColumnNameNetworkPeers].SetValue(host.Metrics.NetworkPeers)
			columns[enums.NodeColumnNameUptime].SetValue(host.Metrics.Uptime)
			columns[enums.NodeColumnNameVersion].SetValue(host.Metrics.Version)
			columns[enums.NodeColumnNameCommit].SetValue(host.Metrics.Commit)

			if host.Location == nil {
				columns[enums.NodeColumnNameCountry].SetValue(nil)

				continue
			}

			if !enabledEmojis {
				country = host.Location.CountryName
			}

			columns[enums.NodeColumnNameCountry].SetValue(country)
		}
	case enums.TableTypeRPC:
		hosts = checker.getHostsByTableType(enums.TableTypeRPC)
		columnConfig = tables.ColumnConfigRPC
		columns = make(tablebuilder.Columns, len(columnConfig))

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		tableConfig = tablebuilder.TableConfig{
			Name:       tables.GetTableTitle(networkType, enums.TableTypeRPC, enabledEmojis),
			Colors:     tablebuilder.GetTableColorsFromString(colorScheme),
			Tag:        tables.TableTagRPC,
			Style:      tables.TableStyleRPC,
			RowsCount:  0,
			Columns:    columns,
			SortConfig: tables.TableSortConfigRPC,
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

			columns[enums.NodeColumnNameHealth].SetValue(status)
			columns[enums.NodeColumnNameAddress].SetValue(address)
			columns[enums.NodeColumnNamePortRPC].SetValue(port)
			columns[enums.NodeColumnNameTotalTransactions].SetValue(host.Metrics.TotalTransactions)
			columns[enums.NodeColumnNameLatestCheckpoint].SetValue(host.Metrics.LatestCheckpoint)
		}
	case enums.TableTypeValidators:
		hosts = checker.getHostsByTableType(enums.TableTypeValidators)
		columnConfig = tables.ColumnConfigSystem
		columns = make(tablebuilder.Columns, len(columnConfig))

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		tableConfig = tablebuilder.TableConfig{
			Name:       tables.GetTableTitle(networkType, enums.TableTypeSystemState, enabledEmojis),
			Colors:     tablebuilder.GetTableColorsFromString(colorScheme),
			Tag:        tables.TableTagSystem,
			Style:      tables.TableStyleSystem,
			RowsCount:  0,
			Columns:    columns,
			SortConfig: tables.TableSortConfigSystem,
		}

		for _, host := range hosts {
			if !host.Metrics.Updated {
				continue
			}

			tableConfig.RowsCount++

			systemState := host.Metrics.SystemState

			columns[enums.SystemColumnNameStorageFund].SetValue(systemState.StorageFund)
			columns[enums.SystemColumnNameReferenceGasPrice].SetValue(systemState.ReferenceGasPrice)
			columns[enums.SystemColumnNameEpochDurationMs].SetValue(systemState.EpochDurationMs)
			columns[enums.SystemColumnNameStakeSubsidyCounter].SetValue(systemState.StakeSubsidyEpochCounter)
			columns[enums.SystemColumnNameStakeSubsidyBalance].SetValue(systemState.StakeSubsidyBalance)
			columns[enums.SystemColumnNameStakeSubsidyCurrentEpochAmount].SetValue(systemState.StakeSubsidyCurrentEpochAmount)
			columns[enums.SystemColumnNameTotalStake].SetValue(systemState.TotalStake)
			columns[enums.SystemColumnNameValidatorsCount].SetValue(len(systemState.ActiveValidators))
			columns[enums.SystemColumnNameValidatorsAtRiskCount].SetValue(len(systemState.AtRiskValidators))
		}
	}

	checker.setBuilderTableType(tableType, tableConfig)
}
