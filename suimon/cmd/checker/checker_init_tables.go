package checker

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums/columnnames"
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

	if displayConfig.ValidatorTable.Display {
		checker.InitTable(enums.TableTypeValidator)
	}

	if displayConfig.PeersTable.Display {
		checker.InitTable(enums.TableTypePeers)
	}

	if displayConfig.SystemTable.Display {
		checker.InitTable(enums.TableTypeSystemState)
	}

	if displayConfig.ActiveValidatorsTable.Display {
		checker.InitTable(enums.TableTypeActiveValidators)
	}
}

// InitTable initializes a specific table of the "Checker" struct instance passed as a pointer receiver,
// based on the provided "TableType".
// Parameters: tableType: an enums.TableType representing the type of table to initialize.
// Returns: None.
func (checker *Checker) InitTable(tableType enums.TableType) {
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
		columnConfig = tables.ColumnConfigNode[:columnnames.NodeColumnNameCurrentRound]
		columns = make(tablebuilder.Columns, len(columnConfig))

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		tableConfig = tablebuilder.TableConfig{
			Name:       tables.GetTableTitle(enums.TableTypeNode, enabledEmojis),
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

			columns[columnnames.NodeColumnNameHealth].SetValue(status)
			columns[columnnames.NodeColumnNameAddress].SetValue(address)
			columns[columnnames.NodeColumnNamePortRPC].SetValue(port)
			columns[columnnames.NodeColumnNameTotalTransactions].SetValue(host.Metrics.TotalTransactions)
			columns[columnnames.NodeColumnNameTotalTransactionCertificates].SetValue(host.Metrics.TotalTransactionCertificates)
			columns[columnnames.NodeColumnNameTotalTransactionEffects].SetValue(host.Metrics.TotalTransactionEffects)
			columns[columnnames.NodeColumnNameLatestCheckpoint].SetValue(host.Metrics.LatestCheckpoint)
			columns[columnnames.NodeColumnNameHighestKnownCheckpoint].SetValue(host.Metrics.HighestKnownCheckpoint)
			columns[columnnames.NodeColumnNameHighestSyncedCheckpoint].SetValue(host.Metrics.HighestSyncedCheckpoint)
			columns[columnnames.NodeColumnNameLastExecutedCheckpoint].SetValue(host.Metrics.LastExecutedCheckpoint)
			columns[columnnames.NodeColumnNameCheckpointExecBacklog].SetValue(host.Metrics.CheckpointExecBacklog)
			columns[columnnames.NodeColumnNameCheckpointSyncBacklog].SetValue(host.Metrics.CheckpointSyncBacklog)
			columns[columnnames.NodeColumnNameCurrentEpoch].SetValue(host.Metrics.CurrentEpoch)
			columns[columnnames.NodeColumnNameTXSyncPercentage].SetValue(fmt.Sprintf("%v%%", host.Metrics.TxSyncPercentage))
			columns[columnnames.NodeColumnNameCheckSyncPercentage].SetValue(fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage))
			columns[columnnames.NodeColumnNameNetworkPeers].SetValue(host.Metrics.NetworkPeers)
			columns[columnnames.NodeColumnNameUptime].SetValue(host.Metrics.Uptime)
			columns[columnnames.NodeColumnNameVersion].SetValue(host.Metrics.Version)
			columns[columnnames.NodeColumnNameCommit].SetValue(host.Metrics.Commit)

			if host.Location == nil {
				columns[columnnames.NodeColumnNameCountry].SetValue(nil)

				continue
			}

			if !enabledEmojis {
				country = host.Location.CountryName
			}

			columns[columnnames.NodeColumnNameCountry].SetValue(country)
		}
	case enums.TableTypeValidator:
		hosts = checker.getHostsByTableType(enums.TableTypeValidator)
		columnConfig = tables.ColumnConfigNode
		columns = make(tablebuilder.Columns, len(columnConfig))

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		tableConfig = tablebuilder.TableConfig{
			Name:       tables.GetTableTitle(enums.TableTypeValidator, enabledEmojis),
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

			columns[columnnames.NodeColumnNameHealth].SetValue(status)
			columns[columnnames.NodeColumnNameAddress].SetValue(address)
			columns[columnnames.NodeColumnNamePortRPC].SetValue(port)
			columns[columnnames.NodeColumnNameTotalTransactions].SetValue(host.Metrics.TotalTransactions)
			columns[columnnames.NodeColumnNameTotalTransactionCertificates].SetValue(host.Metrics.TotalTransactionCertificates)
			columns[columnnames.NodeColumnNameTotalTransactionEffects].SetValue(host.Metrics.TotalTransactionEffects)
			columns[columnnames.NodeColumnNameLatestCheckpoint].SetValue(host.Metrics.LatestCheckpoint)
			columns[columnnames.NodeColumnNameHighestKnownCheckpoint].SetValue(host.Metrics.HighestKnownCheckpoint)
			columns[columnnames.NodeColumnNameHighestSyncedCheckpoint].SetValue(host.Metrics.HighestSyncedCheckpoint)
			columns[columnnames.NodeColumnNameLastExecutedCheckpoint].SetValue(host.Metrics.LastExecutedCheckpoint)
			columns[columnnames.NodeColumnNameCheckpointExecBacklog].SetValue(host.Metrics.CheckpointExecBacklog)
			columns[columnnames.NodeColumnNameCheckpointSyncBacklog].SetValue(host.Metrics.CheckpointSyncBacklog)
			columns[columnnames.NodeColumnNameCurrentEpoch].SetValue(host.Metrics.CurrentEpoch)
			columns[columnnames.NodeColumnNameTXSyncPercentage].SetValue(fmt.Sprintf("%v%%", host.Metrics.TxSyncPercentage))
			columns[columnnames.NodeColumnNameCheckSyncPercentage].SetValue(fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage))
			columns[columnnames.NodeColumnNameNetworkPeers].SetValue(host.Metrics.NetworkPeers)
			columns[columnnames.NodeColumnNameUptime].SetValue(host.Metrics.Uptime)
			columns[columnnames.NodeColumnNameVersion].SetValue(host.Metrics.Version)
			columns[columnnames.NodeColumnNameCommit].SetValue(host.Metrics.Commit)
			columns[columnnames.NodeColumnNameCurrentRound].SetValue(host.Metrics.CurrentRound)
			columns[columnnames.NodeColumnNameHighestProcessedRound].SetValue(host.Metrics.HighestProcessedRound)
			columns[columnnames.NodeColumnNameLastCommittedRound].SetValue(host.Metrics.LastCommittedRound)
			columns[columnnames.NodeColumnNamePrimaryNetworkPeers].SetValue(host.Metrics.PrimaryNetworkPeers)
			columns[columnnames.NodeColumnNameWorkerNetworkPeers].SetValue(host.Metrics.WorkerNetworkPeers)
			columns[columnnames.NodeColumnNameSkippedConsensusTransactions].SetValue(host.Metrics.SkippedConsensusTransactions)
			columns[columnnames.NodeColumnNameTotalSignatureErrors].SetValue(host.Metrics.TotalSignatureErrors)

			if host.Location == nil {
				columns[columnnames.NodeColumnNameCountry].SetValue(nil)

				continue
			}

			if !enabledEmojis {
				country = host.Location.CountryName
			}

			columns[columnnames.NodeColumnNameCountry].SetValue(country)

		}
	case enums.TableTypePeers:
		hosts = checker.getHostsByTableType(enums.TableTypePeers)
		columnConfig = tables.ColumnConfigPeer
		columns = make(tablebuilder.Columns, len(columnConfig))

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		tableConfig = tablebuilder.TableConfig{
			Name:       tables.GetTableTitle(enums.TableTypePeers, enabledEmojis),
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

			columns[columnnames.NodeColumnNameHealth].SetValue(status)
			columns[columnnames.NodeColumnNameAddress].SetValue(address)
			columns[columnnames.NodeColumnNamePortRPC].SetValue(port)
			columns[columnnames.NodeColumnNameTotalTransactions].SetValue(host.Metrics.TotalTransactions)
			columns[columnnames.NodeColumnNameTotalTransactionCertificates].SetValue(host.Metrics.TotalTransactionCertificates)
			columns[columnnames.NodeColumnNameTotalTransactionEffects].SetValue(host.Metrics.TotalTransactionEffects)
			columns[columnnames.NodeColumnNameLatestCheckpoint].SetValue(host.Metrics.LatestCheckpoint)
			columns[columnnames.NodeColumnNameHighestKnownCheckpoint].SetValue(host.Metrics.HighestKnownCheckpoint)
			columns[columnnames.NodeColumnNameHighestSyncedCheckpoint].SetValue(host.Metrics.HighestSyncedCheckpoint)
			columns[columnnames.NodeColumnNameLastExecutedCheckpoint].SetValue(host.Metrics.LastExecutedCheckpoint)
			columns[columnnames.NodeColumnNameCheckpointExecBacklog].SetValue(host.Metrics.CheckpointExecBacklog)
			columns[columnnames.NodeColumnNameCheckpointSyncBacklog].SetValue(host.Metrics.CheckpointSyncBacklog)
			columns[columnnames.NodeColumnNameCurrentEpoch].SetValue(host.Metrics.CurrentEpoch)
			columns[columnnames.NodeColumnNameTXSyncPercentage].SetValue(fmt.Sprintf("%v%%", host.Metrics.TxSyncPercentage))
			columns[columnnames.NodeColumnNameCheckSyncPercentage].SetValue(fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage))
			columns[columnnames.NodeColumnNameNetworkPeers].SetValue(host.Metrics.NetworkPeers)
			columns[columnnames.NodeColumnNameUptime].SetValue(host.Metrics.Uptime)
			columns[columnnames.NodeColumnNameVersion].SetValue(host.Metrics.Version)
			columns[columnnames.NodeColumnNameCommit].SetValue(host.Metrics.Commit)

			if host.Location == nil {
				columns[columnnames.NodeColumnNameCountry].SetValue(nil)

				continue
			}

			if !enabledEmojis {
				country = host.Location.CountryName
			}

			columns[columnnames.NodeColumnNameCountry].SetValue(country)
		}
	case enums.TableTypeRPC:
		hosts = checker.getHostsByTableType(enums.TableTypeRPC)
		columnConfig = tables.ColumnConfigRPC
		columns = make(tablebuilder.Columns, len(columnConfig))

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		tableConfig = tablebuilder.TableConfig{
			Name:       tables.GetTableTitle(enums.TableTypeRPC, enabledEmojis),
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

			columns[columnnames.NodeColumnNameHealth].SetValue(status)
			columns[columnnames.NodeColumnNameAddress].SetValue(address)
			columns[columnnames.NodeColumnNamePortRPC].SetValue(port)
			columns[columnnames.NodeColumnNameTotalTransactions].SetValue(host.Metrics.TotalTransactions)
			columns[columnnames.NodeColumnNameLatestCheckpoint].SetValue(host.Metrics.LatestCheckpoint)
		}
	case enums.TableTypeSystemState:
		hosts = checker.getHostsByTableType(enums.TableTypeSystemState)
		columnConfig = tables.ColumnConfigSystem
		columns = make(tablebuilder.Columns, len(columnConfig))

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		tableConfig = tablebuilder.TableConfig{
			Name:       tables.GetTableTitle(enums.TableTypeSystemState, enabledEmojis),
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

			columns[columnnames.SystemColumnNameStorageFund].SetValue(systemState.StorageFund)
			columns[columnnames.SystemColumnNameReferenceGasPrice].SetValue(systemState.ReferenceGasPrice)
			columns[columnnames.SystemColumnNameEpochDurationMs].SetValue(systemState.EpochDurationMs)
			columns[columnnames.SystemColumnNameStakeSubsidyCounter].SetValue(systemState.StakeSubsidyEpochCounter)
			columns[columnnames.SystemColumnNameStakeSubsidyBalance].SetValue(systemState.StakeSubsidyBalance)
			columns[columnnames.SystemColumnNameStakeSubsidyCurrentEpochAmount].SetValue(systemState.StakeSubsidyCurrentEpochAmount)
			columns[columnnames.SystemColumnNameTotalStake].SetValue(systemState.TotalStake)
			columns[columnnames.SystemColumnNameValidatorsCount].SetValue(len(systemState.ActiveValidators))
			columns[columnnames.SystemColumnNameValidatorsAtRiskCount].SetValue(len(systemState.AtRiskValidators))
		}
	case enums.TableTypeActiveValidators:
		hosts = checker.getHostsByTableType(enums.TableTypeActiveValidators)
		columnConfig = tables.ColumnConfigActiveValidator
		columns = make(tablebuilder.Columns, len(columnConfig))

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		tableConfig = tablebuilder.TableConfig{
			Name:       tables.GetTableTitle(enums.TableTypeActiveValidators, enabledEmojis),
			Colors:     tablebuilder.GetTableColorsFromString(colorScheme),
			Tag:        tables.TableTagActiveValidator,
			Style:      tables.TableStyleActiveValidator,
			RowsCount:  0,
			Columns:    columns,
			SortConfig: tables.TableSortConfigActiveValidator,
		}

		activeValidators := hosts[0].Metrics.SystemState.ActiveValidators

		for _, validator := range activeValidators {
			tableConfig.RowsCount++

			columns[columnnames.ActiveValidatorColumnNameName].SetValue(validator.Name)
			columns[columnnames.ActiveValidatorColumnNameNetAddress].SetValue(validator.NetAddress)
			columns[columnnames.ActiveValidatorColumnNameVotingPower].SetValue(validator.VotingPower)
			columns[columnnames.ActiveValidatorColumnNameGasPrice].SetValue(validator.GasPrice)
			columns[columnnames.ActiveValidatorColumnNameCommissionRate].SetValue(validator.CommissionRate)
			columns[columnnames.ActiveValidatorColumnNameNextEpochStake].SetValue(validator.NextEpochStake)
			columns[columnnames.ActiveValidatorColumnNameNextEpochGasPrice].SetValue(validator.NextEpochGasPrice)
			columns[columnnames.ActiveValidatorColumnNameNextEpochCommissionRate].SetValue(validator.NextEpochCommissionRate)
			columns[columnnames.ActiveValidatorColumnNameStakingPoolSuiBalance].SetValue(validator.StakingPoolSuiBalance)
			columns[columnnames.ActiveValidatorColumnNameRewardsPool].SetValue(validator.RewardsPool)
			columns[columnnames.ActiveValidatorColumnNamePoolTokenBalance].SetValue(validator.PoolTokenBalance)
			columns[columnnames.ActiveValidatorColumnNamePendingStake].SetValue(validator.PendingStake)
		}
	}

	checker.setBuilderTableType(tableType, tableConfig)
}
