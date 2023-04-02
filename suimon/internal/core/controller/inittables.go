package controller

import (
	"fmt"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums/columnnames"
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
		enums.TableTypeRPC:              displayConfig.RPCTable.Display,
		enums.TableTypeNode:             displayConfig.NodeTable.Display,
		enums.TableTypeValidator:        displayConfig.ValidatorTable.Display,
		enums.TableTypePeers:            displayConfig.PeersTable.Display,
		enums.TableTypeSystemState:      displayConfig.SystemTable.Display,
		enums.TableTypeActiveValidators: displayConfig.ActiveValidatorsTable.Display,
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
		columns[columnName].Config = config
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

	for idx, host := range hosts {
		if !host.Metrics.Updated {
			continue
		}

		tableConfig.RowsCount++

		switch tableType {
		case enums.TableTypeNode:
			tables.SetColumnValues(columns, getNodeColumnValues(idx, &host, enabledEmojis))
		case enums.TableTypeValidator:
			tables.SetColumnValues(columns, getValidatorColumnValues(idx, &host, enabledEmojis))
		case enums.TableTypePeers:
			tables.SetColumnValues(columns, getNodeColumnValues(idx, &host, enabledEmojis))
		case enums.TableTypeRPC:
			tables.SetColumnValues(columns, getRPCColumnValues(idx, &host, enabledEmojis))
		case enums.TableTypeSystemState:
			tables.SetColumnValues(columns, getSystemStateColumnValues(idx, &host))
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
func getNodeColumnValues(idx int, host *host.Host, enabledEmojis bool) map[columnnames.NodeColumnName]any {
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

	columnValues := map[columnnames.NodeColumnName]any{
		columnnames.NodeColumnNameIndex:                        idx + 1,
		columnnames.NodeColumnNameHealth:                       status,
		columnnames.NodeColumnNameAddress:                      address,
		columnnames.NodeColumnNamePortRPC:                      port,
		columnnames.NodeColumnNameTotalTransactionBlocks:       host.Metrics.TotalTransactionsBlocks,
		columnnames.NodeColumnNameTotalTransactionCertificates: host.Metrics.TotalTransactionCertificates,
		columnnames.NodeColumnNameTotalTransactionEffects:      host.Metrics.TotalTransactionEffects,
		columnnames.NodeColumnNameLatestCheckpoint:             host.Metrics.LatestCheckpoint,
		columnnames.NodeColumnNameHighestKnownCheckpoint:       host.Metrics.HighestKnownCheckpoint,
		columnnames.NodeColumnNameHighestSyncedCheckpoint:      host.Metrics.HighestSyncedCheckpoint,
		columnnames.NodeColumnNameLastExecutedCheckpoint:       host.Metrics.LastExecutedCheckpoint,
		columnnames.NodeColumnNameCheckpointExecBacklog:        host.Metrics.CheckpointExecBacklog,
		columnnames.NodeColumnNameCheckpointSyncBacklog:        host.Metrics.CheckpointSyncBacklog,
		columnnames.NodeColumnNameCurrentEpoch:                 host.Metrics.CurrentEpoch,
		columnnames.NodeColumnNameTXSyncPercentage:             fmt.Sprintf("%v%%", host.Metrics.TxSyncPercentage),
		columnnames.NodeColumnNameCheckSyncPercentage:          fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage),
		columnnames.NodeColumnNameNetworkPeers:                 host.Metrics.NetworkPeers,
		columnnames.NodeColumnNameUptime:                       host.Metrics.Uptime,
		columnnames.NodeColumnNameVersion:                      host.Metrics.Version,
		columnnames.NodeColumnNameCommit:                       host.Metrics.Commit,
		columnnames.NodeColumnNameCountry:                      nil,
	}

	if host.Location != nil {
		if !enabledEmojis {
			country = host.Location.CountryName
		}

		columnValues[columnnames.NodeColumnNameCountry] = country
	}

	return columnValues
}

// getValidatorColumnValues returns a map of ValidatorColumnName values to corresponding values for a validator at the specified index on the specified host.
// The function retrieves information about the validator from the host's internal state and formats it into a map of ValidatorColumnName keys and corresponding values.
// The function also includes emoji values in the map if the specified flag is true.
// Returns a map of ValidatorColumnName keys to corresponding values.
func getValidatorColumnValues(idx int, host *host.Host, enabledEmojis bool) map[columnnames.ValidatorColumnName]any {
	var (
		status  any = host.Status
		country any = host.Location.String()
		address     = host.HostPort.Address
	)

	if !enabledEmojis {
		status = host.Status.StatusToPlaceholder()
	}

	columnValues := map[columnnames.ValidatorColumnName]any{
		columnnames.ValidatorColumnNameIndex:                        idx + 1,
		columnnames.ValidatorColumnNameHealth:                       status,
		columnnames.ValidatorColumnNameAddress:                      address,
		columnnames.ValidatorColumnNameTotalTransactionCertificates: host.Metrics.TotalTransactionCertificates,
		columnnames.ValidatorColumnNameTotalTransactionEffects:      host.Metrics.TotalTransactionEffects,
		columnnames.ValidatorColumnNameHighestKnownCheckpoint:       host.Metrics.HighestKnownCheckpoint,
		columnnames.ValidatorColumnNameHighestSyncedCheckpoint:      host.Metrics.HighestSyncedCheckpoint,
		columnnames.ValidatorColumnNameLastExecutedCheckpoint:       host.Metrics.LastExecutedCheckpoint,
		columnnames.ValidatorColumnNameCheckpointExecBacklog:        host.Metrics.CheckpointExecBacklog,
		columnnames.ValidatorColumnNameCheckpointSyncBacklog:        host.Metrics.CheckpointSyncBacklog,
		columnnames.ValidatorColumnNameCurrentEpoch:                 host.Metrics.CurrentEpoch,
		columnnames.ValidatorColumnNameCheckSyncPercentage:          fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage),
		columnnames.ValidatorColumnNameNetworkPeers:                 host.Metrics.NetworkPeers,
		columnnames.ValidatorColumnNameUptime:                       host.Metrics.Uptime,
		columnnames.ValidatorColumnNameVersion:                      host.Metrics.Version,
		columnnames.ValidatorColumnNameCommit:                       host.Metrics.Commit,
		columnnames.ValidatorColumnNameCurrentRound:                 host.Metrics.CurrentRound,
		columnnames.ValidatorColumnNameHighestProcessedRound:        host.Metrics.HighestProcessedRound,
		columnnames.ValidatorColumnNameLastCommittedRound:           host.Metrics.LastCommittedRound,
		columnnames.ValidatorColumnNamePrimaryNetworkPeers:          host.Metrics.PrimaryNetworkPeers,
		columnnames.ValidatorColumnNameWorkerNetworkPeers:           host.Metrics.WorkerNetworkPeers,
		columnnames.ValidatorColumnNameSkippedConsensusTransactions: host.Metrics.SkippedConsensusTransactions,
		columnnames.ValidatorColumnNameTotalSignatureErrors:         host.Metrics.TotalSignatureErrors,
	}

	if host.Location != nil {
		if !enabledEmojis {
			country = host.Location.CountryName
		}

		columnValues[columnnames.ValidatorColumnNameCountry] = country
	}

	return columnValues
}

// getRPCColumnValues returns a map of NodeColumnName values to corresponding values for the RPC service on the specified host.
// The function retrieves information about the RPC service from the host's internal state and formats it into a map of NodeColumnName keys and corresponding values.
// Returns a map of NodeColumnName keys to corresponding values.
func getRPCColumnValues(idx int, host *host.Host, enabledEmojis bool) map[columnnames.NodeColumnName]any {
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

	return map[columnnames.NodeColumnName]any{
		columnnames.NodeColumnNameIndex:                  idx + 1,
		columnnames.NodeColumnNameHealth:                 status,
		columnnames.NodeColumnNameAddress:                address,
		columnnames.NodeColumnNamePortRPC:                port,
		columnnames.NodeColumnNameTotalTransactionBlocks: host.Metrics.TotalTransactionsBlocks,
		columnnames.NodeColumnNameLatestCheckpoint:       host.Metrics.LatestCheckpoint,
	}
}

// getSystemStateColumnValues returns a map of SystemColumnName values to corresponding values for the system state on the specified host.
// The function retrieves information about the system state from the host's internal state and formats it into a map of SystemColumnName keys and corresponding values.
// Returns a map of SystemColumnName keys to corresponding values.
func getSystemStateColumnValues(idx int, host *host.Host) map[columnnames.SystemColumnName]any {
	systemState := host.Metrics.SystemState

	return map[columnnames.SystemColumnName]any{
		columnnames.SystemColumnNameIndex:           idx + 1,
		columnnames.SystemColumnNameEpoch:           systemState.Epoch,
		columnnames.SystemColumnNameEpochDurationMs: systemState.EpochDurationMs,
		//columnnames.SystemColumnNameStorageFund:                    systemState.StorageFund,
		columnnames.SystemColumnNameReferenceGasPrice: systemState.ReferenceGasPrice,
		//columnnames.SystemColumnNameStakeSubsidyCounter:            systemState.StakeSubsidyEpochCounter,
		columnnames.SystemColumnNameStakeSubsidyBalance: systemState.StakeSubsidyBalance,
		//columnnames.SystemColumnNameStakeSubsidyCurrentEpochAmount: systemState.StakeSubsidyCurrentEpochAmount,
		columnnames.SystemColumnNameTotalStake:                  systemState.TotalStake,
		columnnames.SystemColumnNameValidatorsCount:             len(systemState.ActiveValidators),
		columnnames.SystemColumnNamePendingActiveValidatorsSize: systemState.PendingActiveValidatorsSize,
		columnnames.SystemColumnNamePendingRemovals:             len(systemState.PendingRemovals),
		columnnames.SystemColumnNameValidatorsCandidateSize:     systemState.ValidatorCandidatesSize,
		columnnames.SystemColumnNameValidatorsAtRiskCount:       len(systemState.AtRiskValidators),
	}
}

// getActiveValidatorColumnValues returns a map of ActiveValidatorColumnName values to corresponding values for the specified active validator.
// The function retrieves information about the active validator from the provided metrics.Validator object and formats it into a map of ActiveValidatorColumnName keys and corresponding values.
// Returns a map of ActiveValidatorColumnName keys to corresponding values.
func getActiveValidatorColumnValues(idx int, validator *metrics.Validator) map[columnnames.ActiveValidatorColumnName]any {
	return map[columnnames.ActiveValidatorColumnName]any{
		columnnames.ActiveValidatorColumnNameIndex:                   idx + 1,
		columnnames.ActiveValidatorColumnNameName:                    validator.Name,
		columnnames.ActiveValidatorColumnNameNetAddress:              validator.NetAddress,
		columnnames.ActiveValidatorColumnNameVotingPower:             validator.VotingPower,
		columnnames.ActiveValidatorColumnNameGasPrice:                validator.GasPrice,
		columnnames.ActiveValidatorColumnNameCommissionRate:          validator.CommissionRate,
		columnnames.ActiveValidatorColumnNameNextEpochStake:          validator.NextEpochStake,
		columnnames.ActiveValidatorColumnNameNextEpochGasPrice:       validator.NextEpochGasPrice,
		columnnames.ActiveValidatorColumnNameNextEpochCommissionRate: validator.NextEpochCommissionRate,
		columnnames.ActiveValidatorColumnNameStakingPoolSuiBalance:   validator.StakingPoolSuiBalance,
		columnnames.ActiveValidatorColumnNameRewardsPool:             validator.RewardsPool,
		columnnames.ActiveValidatorColumnNamePoolTokenBalance:        validator.PoolTokenBalance,
		columnnames.ActiveValidatorColumnNamePendingStake:            validator.PendingStake,
	}
}
