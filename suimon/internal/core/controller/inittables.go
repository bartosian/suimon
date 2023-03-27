package controller

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums/columnnames"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/tablebuilder"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/tablebuilder/tables"
)

const (
	rpcPortDefault = "9000"
)

// InitTables initializes the tables of the "Checker" struct instance passed as a pointer receiver.
// Parameters: None.
// Returns: None.
func (checker *CheckerController) InitTables() error {
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

	return nil
}

// InitTable initializes a specific table of the "Checker" struct instance passed as a pointer receiver,
// based on the provided "TableType".
// Parameters: tableType: an enums.TableType representing the type of table to initialize.
// Returns: None.
func (checker *CheckerController) InitTable(tableType enums.TableType) {
	enabledEmojis := checker.suimonConfig.MonitorsVisual.EnableEmojis

	var (
		hosts        []host.Host
		tableConfig  tablebuilder.TableConfig
		columnConfig []table.ColumnConfig
		columns      tablebuilder.Columns
	)

	switch tableType {
	case enums.TableTypeNode:
		hosts = checker.getHostsByTableType(enums.TableTypeNode)
		columnConfig = tables.ColumnsConfigNode
		columns = make(tablebuilder.Columns, len(columnConfig))

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		tableConfig = tablebuilder.TableConfig{
			Name:         tables.GetTableTitle(enums.TableTypeNode),
			Tag:          tables.TableTagNode,
			Style:        tables.TableStyleNode,
			Columns:      columns,
			ColumnsCount: len(columns),
			Rows:         tables.RowsNode,
			RowsCount:    0,
			SortConfig:   tables.TableSortConfigNode,
		}

		for idx, host := range hosts {
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

			columns[columnnames.NodeColumnNameIndex].SetValue(idx + 1)
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
		columnConfig = tables.ColumnsConfigValidator
		columns = make(tablebuilder.Columns, len(columnConfig))

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		tableConfig = tablebuilder.TableConfig{
			Name:         tables.GetTableTitle(enums.TableTypeValidator),
			Tag:          tables.TableTagValidator,
			Style:        tables.TableStyleValidator,
			Columns:      columns,
			ColumnsCount: len(columns),
			Rows:         tables.RowsValidator,
			RowsCount:    0,
			SortConfig:   tables.TableSortConfigValidator,
		}

		for idx, host := range hosts {
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

			columns[columnnames.ValidatorColumnNameIndex].SetValue(idx + 1)
			columns[columnnames.ValidatorColumnNameHealth].SetValue(status)
			columns[columnnames.ValidatorColumnNameAddress].SetValue(address)
			columns[columnnames.ValidatorColumnNameTotalTransactionCertificates].SetValue(host.Metrics.TotalTransactionCertificates)
			columns[columnnames.ValidatorColumnNameTotalTransactionEffects].SetValue(host.Metrics.TotalTransactionEffects)
			columns[columnnames.ValidatorColumnNameHighestKnownCheckpoint].SetValue(host.Metrics.HighestKnownCheckpoint)
			columns[columnnames.ValidatorColumnNameHighestSyncedCheckpoint].SetValue(host.Metrics.HighestSyncedCheckpoint)
			columns[columnnames.ValidatorColumnNameLastExecutedCheckpoint].SetValue(host.Metrics.LastExecutedCheckpoint)
			columns[columnnames.ValidatorColumnNameCheckpointExecBacklog].SetValue(host.Metrics.CheckpointExecBacklog)
			columns[columnnames.ValidatorColumnNameCheckpointSyncBacklog].SetValue(host.Metrics.CheckpointSyncBacklog)
			columns[columnnames.ValidatorColumnNameCurrentEpoch].SetValue(host.Metrics.CurrentEpoch)
			columns[columnnames.ValidatorColumnNameCheckSyncPercentage].SetValue(fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage))
			columns[columnnames.ValidatorColumnNameNetworkPeers].SetValue(host.Metrics.NetworkPeers)
			columns[columnnames.ValidatorColumnNameUptime].SetValue(host.Metrics.Uptime)
			columns[columnnames.ValidatorColumnNameVersion].SetValue(host.Metrics.Version)
			columns[columnnames.ValidatorColumnNameCommit].SetValue(host.Metrics.Commit)
			columns[columnnames.ValidatorColumnNameCurrentRound].SetValue(host.Metrics.CurrentRound)
			columns[columnnames.ValidatorColumnNameHighestProcessedRound].SetValue(host.Metrics.HighestProcessedRound)
			columns[columnnames.ValidatorColumnNameLastCommittedRound].SetValue(host.Metrics.LastCommittedRound)
			columns[columnnames.ValidatorColumnNamePrimaryNetworkPeers].SetValue(host.Metrics.PrimaryNetworkPeers)
			columns[columnnames.ValidatorColumnNameWorkerNetworkPeers].SetValue(host.Metrics.WorkerNetworkPeers)
			columns[columnnames.ValidatorColumnNameSkippedConsensusTransactions].SetValue(host.Metrics.SkippedConsensusTransactions)
			columns[columnnames.ValidatorColumnNameTotalSignatureErrors].SetValue(host.Metrics.TotalSignatureErrors)

			if host.Location == nil {
				columns[columnnames.ValidatorColumnNameCountry].SetValue(nil)

				continue
			}

			if !enabledEmojis {
				country = host.Location.CountryName
			}

			columns[columnnames.ValidatorColumnNameCountry].SetValue(country)

		}
	case enums.TableTypePeers:
		hosts = checker.getHostsByTableType(enums.TableTypePeers)
		columnConfig = tables.ColumnsConfigPeer
		columns = make(tablebuilder.Columns, len(columnConfig))

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		tableConfig = tablebuilder.TableConfig{
			Name:         tables.GetTableTitle(enums.TableTypePeers),
			Tag:          tables.TableTagPeer,
			Style:        tables.TableStylePeer,
			Columns:      columns,
			ColumnsCount: len(columns),
			Rows:         tables.RowsPeer,
			RowsCount:    0,
			SortConfig:   tables.TableSortConfigPeer,
		}

		for idx, host := range hosts {
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

			columns[columnnames.NodeColumnNameIndex].SetValue(idx + 1)
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
		columnConfig = tables.ColumnsConfigRPC
		columns = make(tablebuilder.Columns, len(columnConfig))

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		tableConfig = tablebuilder.TableConfig{
			Name:         tables.GetTableTitle(enums.TableTypeRPC),
			Tag:          tables.TableTagRPC,
			Style:        tables.TableStyleRPC,
			Columns:      columns,
			ColumnsCount: len(columns),
			Rows:         tables.RowsRPC,
			RowsCount:    0,
			SortConfig:   tables.TableSortConfigRPC,
		}

		for idx, host := range hosts {
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

			columns[columnnames.NodeColumnNameIndex].SetValue(idx + 1)
			columns[columnnames.NodeColumnNameHealth].SetValue(status)
			columns[columnnames.NodeColumnNameAddress].SetValue(address)
			columns[columnnames.NodeColumnNamePortRPC].SetValue(port)
			columns[columnnames.NodeColumnNameTotalTransactions].SetValue(host.Metrics.TotalTransactions)
			columns[columnnames.NodeColumnNameLatestCheckpoint].SetValue(host.Metrics.LatestCheckpoint)
		}
	case enums.TableTypeSystemState:
		hosts = checker.getHostsByTableType(enums.TableTypeSystemState)
		columnConfig = tables.ColumnsConfigSystem
		columns = make(tablebuilder.Columns, len(columnConfig))

		for i := 0; i < len(columnConfig); i++ {
			columns[i].Config = columnConfig[i]
		}

		tableConfig = tablebuilder.TableConfig{
			Name:         tables.GetTableTitle(enums.TableTypeSystemState),
			Tag:          tables.TableTagSystem,
			Style:        tables.TableStyleSystem,
			Columns:      columns,
			ColumnsCount: len(columns),
			Rows:         tables.RowsSystemState,
			RowsCount:    0,
			SortConfig:   tables.TableSortConfigSystem,
		}

		for idx, host := range hosts {
			if !host.Metrics.Updated {
				continue
			}

			tableConfig.RowsCount++

			systemState := host.Metrics.SystemState

			columns[columnnames.ValidatorColumnNameIndex].SetValue(idx + 1)
			columns[columnnames.SystemColumnNameEpoch].SetValue(systemState.Epoch)
			columns[columnnames.SystemColumnNameEpochDurationMs].SetValue(systemState.EpochDurationMs)
			columns[columnnames.SystemColumnNameStorageFund].SetValue(systemState.StorageFund)
			columns[columnnames.SystemColumnNameReferenceGasPrice].SetValue(systemState.ReferenceGasPrice)
			columns[columnnames.SystemColumnNameStakeSubsidyCounter].SetValue(systemState.StakeSubsidyEpochCounter)
			columns[columnnames.SystemColumnNameStakeSubsidyBalance].SetValue(systemState.StakeSubsidyBalance)
			columns[columnnames.SystemColumnNameStakeSubsidyCurrentEpochAmount].SetValue(systemState.StakeSubsidyCurrentEpochAmount)
			columns[columnnames.SystemColumnNameTotalStake].SetValue(systemState.TotalStake)
			columns[columnnames.SystemColumnNameValidatorsCount].SetValue(len(systemState.ActiveValidators))
			columns[columnnames.SystemColumnNamePendingActiveValidatorsSize].SetValue(systemState.PendingActiveValidatorsSize)
			columns[columnnames.SystemColumnNamePendingRemovals].SetValue(len(systemState.PendingRemovals))
			columns[columnnames.SystemColumnNameValidatorsCandidateSize].SetValue(systemState.ValidatorCandidatesSize)
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
			Name:         tables.GetTableTitle(enums.TableTypeActiveValidators),
			Tag:          tables.TableTagActiveValidator,
			Style:        tables.TableStyleActiveValidator,
			Columns:      columns,
			ColumnsCount: len(columns),
			Rows:         tables.RowsActiveValidator,
			RowsCount:    0,
			SortConfig:   tables.TableSortConfigActiveValidator,
		}

		activeValidators := hosts[0].Metrics.SystemState.ActiveValidators

		for idx, validator := range activeValidators {
			tableConfig.RowsCount++

			columns[columnnames.ActiveValidatorColumnNameIndex].SetValue(idx + 1)
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
