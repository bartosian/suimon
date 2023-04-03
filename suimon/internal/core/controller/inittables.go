package controller

import (
	"fmt"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/metrics"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/tablebuilder"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/tablebuilder/tables"
)

const (
	rpcPortDefault = "9000"
)

// InitTables initializes the CheckerController's internal state for the tables by creating and setting the appropriate TableConfig objects for each table type.
// The function populates the CheckerController's internal state with information about the available hosts and their corresponding tables.
// Returns an error if the initialization process fails for any reason.
func (checker *CheckerController) InitTables() error {
	displayConfig := checker.suimonConfig.MonitorsConfig
	tableConfigMap := map[enums.TableType]bool{
		enums.TableTypeRPC:                   displayConfig.RPCTable.Display,
		enums.TableTypeNode:                  displayConfig.NodeTable.Display,
		enums.TableTypeValidator:             displayConfig.ValidatorTable.Display,
		enums.TableTypePeers:                 displayConfig.PeersTable.Display,
		enums.TableTypeSystemState:           displayConfig.SystemTable.Display,
		enums.TableTypeSystemStateValidators: displayConfig.SystemTable.Display,
		enums.TableTypeActiveValidators:      displayConfig.ActiveValidatorsTable.Display,
	}

	for tableType, shouldDisplay := range tableConfigMap {
		if shouldDisplay {
			checker.InitTable(tableType)
		}
	}

	return nil
}

// InitTable initializes the CheckerController's internal state for the specified table type by creating and setting the appropriate Host objects for each host that supports that table type.
// The function populates the CheckerController's internal state with information about the available hosts and their corresponding tables of the specified type.
// Returns nothing.
func (checker *CheckerController) InitTable(tableType enums.TableType) {
	enabledEmojis := checker.suimonConfig.MonitorsVisual.EnableEmojis
	columnConfig := tables.GetColumnConfig(tableType)
	columns := make(tablebuilder.Columns, len(columnConfig))
	hosts := checker.getHostsByTableType(tableType)

	for columnName, config := range columnConfig {
		column := tablebuilder.Column{
			Config: config,
		}

		columns[columnName] = &column
	}

	tableConfig := tablebuilder.TableConfig{
		Name:         tables.GetTableTitle(tableType),
		Tag:          tables.GetTableTag(tableType),
		Style:        tables.GetTableStyle(tableType),
		Colors:       tables.GetTableColors(tableType),
		Rows:         tables.GetTableRows(tableType),
		SortConfig:   tables.GetTableSortConfig(tableType),
		Columns:      columns,
		ColumnsCount: len(columns),
		RowsCount:    0,
	}

	for idx, hostToRender := range hosts {
		if !hostToRender.Metrics.Updated {
			continue
		}

		tableConfig.RowsCount++

		switch tableType {
		case enums.TableTypeNode:
			tables.SetColumnValues(columns, getNodeColumnValues(idx, &hostToRender, enabledEmojis))
		case enums.TableTypeValidator:
			tables.SetColumnValues(columns, getValidatorColumnValues(idx, &hostToRender, enabledEmojis))
		case enums.TableTypePeers:
			tables.SetColumnValues(columns, getNodeColumnValues(idx, &hostToRender, enabledEmojis))
		case enums.TableTypeRPC:
			tables.SetColumnValues(columns, getRPCColumnValues(idx, &hostToRender, enabledEmojis))
		case enums.TableTypeSystemState:
			tables.SetColumnValues(columns, getSystemStateColumnValues(idx, &hostToRender))
		case enums.TableTypeSystemStateValidators:
			tables.SetColumnValues(columns, getSystemStateValidatorsColumnValues(idx, &hostToRender))
		case enums.TableTypeActiveValidators:
			activeValidators := hosts[0].Metrics.SystemState.ActiveValidators

			for idx, validator := range activeValidators {
				columnValues := getActiveValidatorColumnValues(idx, &validator)

				tables.SetColumnValues(columns, columnValues)

				tableConfig.RowsCount++
			}

		}

	}

	checker.setBuilderTableType(tableType, tableConfig)
}

// getNodeColumnValues returns a map of NodeColumnName values to corresponding values for a node at the specified index on the specified host.
// The function retrieves information about the node from the host's internal state and formats it into a map of NodeColumnName keys and corresponding values.
// The function also includes emoji values in the map if the specified flag is true.
// Returns a map of NodeColumnName keys to corresponding values.
func getNodeColumnValues(idx int, host *host.Host, enabledEmojis bool) map[enums.ColumnName]any {
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

	columnValues := map[enums.ColumnName]any{
		enums.ColumnNameIndex:                        idx + 1,
		enums.ColumnNameHealth:                       status,
		enums.ColumnNameAddress:                      address,
		enums.ColumnNamePortRPC:                      port,
		enums.ColumnNameTotalTransactionBlocks:       host.Metrics.TotalTransactionsBlocks,
		enums.ColumnNameTotalTransactionCertificates: host.Metrics.TotalTransactionCertificates,
		enums.ColumnNameTotalTransactionEffects:      host.Metrics.TotalTransactionEffects,
		enums.ColumnNameLatestCheckpoint:             host.Metrics.LatestCheckpoint,
		enums.ColumnNameHighestKnownCheckpoint:       host.Metrics.HighestKnownCheckpoint,
		enums.ColumnNameHighestSyncedCheckpoint:      host.Metrics.HighestSyncedCheckpoint,
		enums.ColumnNameLastExecutedCheckpoint:       host.Metrics.LastExecutedCheckpoint,
		enums.ColumnNameCheckpointExecBacklog:        host.Metrics.CheckpointExecBacklog,
		enums.ColumnNameCheckpointSyncBacklog:        host.Metrics.CheckpointSyncBacklog,
		enums.ColumnNameCurrentEpoch:                 host.Metrics.CurrentEpoch,
		enums.ColumnNameTXSyncPercentage:             fmt.Sprintf("%v%%", host.Metrics.TxSyncPercentage),
		enums.ColumnNameCheckSyncPercentage:          fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage),
		enums.ColumnNameNetworkPeers:                 host.Metrics.NetworkPeers,
		enums.ColumnNameUptime:                       host.Metrics.Uptime,
		enums.ColumnNameVersion:                      host.Metrics.Version,
		enums.ColumnNameCommit:                       host.Metrics.Commit,
		enums.ColumnNameCountry:                      nil,
	}

	if host.Location != nil {
		if !enabledEmojis {
			country = host.Location.CountryName
		}

		columnValues[enums.ColumnNameCountry] = country
	}

	return columnValues
}

// getValidatorColumnValues returns a map of ValidatorColumnName values to corresponding values for a validator at the specified index on the specified host.
// The function retrieves information about the validator from the host's internal state and formats it into a map of ValidatorColumnName keys and corresponding values.
// The function also includes emoji values in the map if the specified flag is true.
// Returns a map of ValidatorColumnName keys to corresponding values.
func getValidatorColumnValues(idx int, host *host.Host, enabledEmojis bool) map[enums.ColumnName]any {
	var (
		status  any = host.Status
		country any = host.Location.String()
		address     = host.HostPort.Address
	)

	if !enabledEmojis {
		status = host.Status.StatusToPlaceholder()
	}

	columnValues := map[enums.ColumnName]any{
		enums.ColumnNameIndex:                        idx + 1,
		enums.ColumnNameHealth:                       status,
		enums.ColumnNameAddress:                      address,
		enums.ColumnNameTotalTransactionCertificates: host.Metrics.TotalTransactionCertificates,
		enums.ColumnNameTotalTransactionEffects:      host.Metrics.TotalTransactionEffects,
		enums.ColumnNameHighestKnownCheckpoint:       host.Metrics.HighestKnownCheckpoint,
		enums.ColumnNameHighestSyncedCheckpoint:      host.Metrics.HighestSyncedCheckpoint,
		enums.ColumnNameLastExecutedCheckpoint:       host.Metrics.LastExecutedCheckpoint,
		enums.ColumnNameCheckpointExecBacklog:        host.Metrics.CheckpointExecBacklog,
		enums.ColumnNameCheckpointSyncBacklog:        host.Metrics.CheckpointSyncBacklog,
		enums.ColumnNameCurrentEpoch:                 host.Metrics.CurrentEpoch,
		enums.ColumnNameCheckSyncPercentage:          fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage),
		enums.ColumnNameNetworkPeers:                 host.Metrics.NetworkPeers,
		enums.ColumnNameUptime:                       host.Metrics.Uptime,
		enums.ColumnNameVersion:                      host.Metrics.Version,
		enums.ColumnNameCommit:                       host.Metrics.Commit,
		enums.ColumnNameCurrentRound:                 host.Metrics.CurrentRound,
		enums.ColumnNameHighestProcessedRound:        host.Metrics.HighestProcessedRound,
		enums.ColumnNameLastCommittedRound:           host.Metrics.LastCommittedRound,
		enums.ColumnNamePrimaryNetworkPeers:          host.Metrics.PrimaryNetworkPeers,
		enums.ColumnNameWorkerNetworkPeers:           host.Metrics.WorkerNetworkPeers,
		enums.ColumnNameSkippedConsensusTransactions: host.Metrics.SkippedConsensusTransactions,
		enums.ColumnNameTotalSignatureErrors:         host.Metrics.TotalSignatureErrors,
		enums.ColumnNameCountry:                      nil,
	}

	if host.Location != nil {
		if !enabledEmojis {
			country = host.Location.CountryName
		}

		columnValues[enums.ColumnNameCountry] = country
	}

	return columnValues
}

// getRPCColumnValues returns a map of NodeColumnName values to corresponding values for the RPC service on the specified host.
// The function retrieves information about the RPC service from the host's internal state and formats it into a map of NodeColumnName keys and corresponding values.
// Returns a map of NodeColumnName keys to corresponding values.
func getRPCColumnValues(idx int, host *host.Host, enabledEmojis bool) map[enums.ColumnName]any {
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

	return map[enums.ColumnName]any{
		enums.ColumnNameIndex:                  idx + 1,
		enums.ColumnNameHealth:                 status,
		enums.ColumnNameAddress:                address,
		enums.ColumnNamePortRPC:                port,
		enums.ColumnNameTotalTransactionBlocks: host.Metrics.TotalTransactionsBlocks,
		enums.ColumnNameLatestCheckpoint:       host.Metrics.LatestCheckpoint,
	}
}

// getSystemStateColumnValues returns a map of SystemColumnName values to corresponding values for the system state on the specified host.
// The function retrieves information about the system state from the host's internal state and formats it into a map of SystemColumnName keys and corresponding values.
// Returns a map of SystemColumnName keys to corresponding values.
func getSystemStateColumnValues(idx int, host *host.Host) map[enums.ColumnName]any {
	systemState := host.Metrics.SystemState

	return map[enums.ColumnName]any{
		enums.ColumnNameIndex:                                       idx + 1,
		enums.ColumnNameSystemEpoch:                                 systemState.Epoch,
		enums.ColumnNameSystemEpochStartTimestampMs:                 systemState.EpochStartTimestampMs,
		enums.ColumnNameSystemEpochDurationMs:                       systemState.EpochDurationMs,
		enums.ColumnNameSystemTotalStake:                            systemState.TotalStake,
		enums.ColumnNameSystemStorageFundTotalObjectStorageRebates:  systemState.StorageFundTotalObjectStorageRebates,
		enums.ColumnNameSystemStorageFundNonRefundableBalance:       systemState.StorageFundNonRefundableBalance,
		enums.ColumnNameSystemReferenceGasPrice:                     systemState.ReferenceGasPrice,
		enums.ColumnNameSystemStakeSubsidyStartEpoch:                systemState.StakeSubsidyStartEpoch,
		enums.ColumnNameSystemStakeSubsidyBalance:                   systemState.StakeSubsidyBalance,
		enums.ColumnNameSystemStakeSubsidyDistributionCounter:       systemState.StakeSubsidyDistributionCounter,
		enums.ColumnNameSystemStakeSubsidyCurrentDistributionAmount: systemState.StakeSubsidyCurrentDistributionAmount,
		enums.ColumnNameSystemStakeSubsidyPeriodLength:              systemState.StakeSubsidyPeriodLength,
		enums.ColumnNameSystemStakeSubsidyDecreaseRate:              systemState.StakeSubsidyDecreaseRate,
	}
}

// getSystemStateValidatorsColumnValues returns a map of ColumnName values to corresponding values for the system state validators metrics.
// The function retrieves information about the system state from the host's internal state and formats it into a map of ColumnName keys and corresponding values.
// Returns a map of ColumnName keys to corresponding values.
func getSystemStateValidatorsColumnValues(idx int, host *host.Host) map[enums.ColumnName]any {
	systemState := host.Metrics.SystemState

	return map[enums.ColumnName]any{
		enums.ColumnNameIndex:                                idx + 1,
		enums.ColumnNameSystemMaxValidatorCount:              systemState.MaxValidatorCount,
		enums.ColumnNameSystemActiveValidatorCount:           len(systemState.ActiveValidators),
		enums.ColumnNameSystemPendingActiveValidatorCount:    systemState.PendingActiveValidatorsSize,
		enums.ColumnNameSystemValidatorCandidateCount:        systemState.ValidatorCandidatesSize,
		enums.ColumnNameSystemPendingRemovalsCount:           len(systemState.PendingRemovals),
		enums.ColumnNameSystemAtRiskValidatorCount:           len(systemState.AtRiskValidators),
		enums.ColumnNameSystemMinValidatorJoiningStake:       systemState.MinValidatorJoiningStake,
		enums.ColumnNameSystemValidatorLowStakeThreshold:     systemState.ValidatorLowStakeThreshold,
		enums.ColumnNameSystemValidatorVeryLowStakeThreshold: systemState.ValidatorVeryLowStakeThreshold,
		enums.ColumnNameSystemValidatorLowStakeGracePeriod:   systemState.ValidatorLowStakeGracePeriod,
	}
}

// getActiveValidatorColumnValues returns a map of ActiveValidatorColumnName values to corresponding values for the specified active validator.
// The function retrieves information about the active validator from the provided metrics.Validator object and formats it into a map of ActiveValidatorColumnName keys and corresponding values.
// Returns a map of ActiveValidatorColumnName keys to corresponding values.
func getActiveValidatorColumnValues(idx int, validator *metrics.Validator) map[enums.ColumnName]any {
	return map[enums.ColumnName]any{
		enums.ColumnNameIndex:                             idx + 1,
		enums.ColumnNameValidatorName:                     validator.Name,
		enums.ColumnNameValidatorNetAddress:               validator.NetAddress,
		enums.ColumnNameValidatorVotingPower:              validator.VotingPower,
		enums.ColumnNameValidatorGasPrice:                 validator.GasPrice,
		enums.ColumnNameValidatorCommissionRate:           validator.CommissionRate,
		enums.ColumnNameValidatorNextEpochStake:           validator.NextEpochStake,
		enums.ColumnNameValidatorNextEpochGasPrice:        validator.NextEpochGasPrice,
		enums.ColumnNameValidatorNextEpochCommissionRate:  validator.NextEpochCommissionRate,
		enums.ColumnNameValidatorStakingPoolSuiBalance:    validator.StakingPoolSuiBalance,
		enums.ColumnNameValidatorRewardsPool:              validator.RewardsPool,
		enums.ColumnNameValidatorPoolTokenBalance:         validator.PoolTokenBalance,
		enums.ColumnNameValidatorPendingStake:             validator.PendingStake,
		enums.ColumnNameValidatorPendingTotalSuiWithdraw:  validator.PendingTotalSuiWithdraw,
		enums.ColumnNameValidatorPendingPoolTokenWithdraw: validator.PendingPoolTokenWithdraw,
	}
}
