package enums

type ColumnName string

const (
	ColumnNameIndex   ColumnName = "IDX"
	ColumnNameHealth  ColumnName = "HEALTH"
	ColumnNameAddress ColumnName = "ADDRESS"
	ColumnNamePortRPC ColumnName = "RPC"
	ColumnNameUptime  ColumnName = "UPTIME"
	ColumnNameVersion ColumnName = "VERSION"
	ColumnNameCommit  ColumnName = "COMMIT"
	ColumnNameCountry ColumnName = "COUNTRY"

	ColumnNameTotalTransactionBlocks       ColumnName = "TOTAL TX\nBLOCKS"
	ColumnNameTotalTransactionCertificates ColumnName = "TOTAL TX\nCERTIFICATES"
	ColumnNameTotalTransactionEffects      ColumnName = "TOTAL TX\nEFFECTS"
	ColumnNameTXSyncPercentage             ColumnName = "TX SYNC\nPERCENTAGE"
	ColumnNameLatestCheckpoint             ColumnName = "LATEST\nCHECKPOINT"
	ColumnNameHighestKnownCheckpoint       ColumnName = "HIGHEST KNOWN\nCHECKPOINT"
	ColumnNameLastExecutedCheckpoint       ColumnName = "LAST EXECUTED\nCHECKPOINT"
	ColumnNameHighestSyncedCheckpoint      ColumnName = "HIGHEST SYNCED\nCHECKPOINT"
	ColumnNameCheckpointExecBacklog        ColumnName = "CHECKPOINT\nEXEC BACKLOG"
	ColumnNameCheckpointSyncBacklog        ColumnName = "CHECKPOINT\nSYNC BACKLOG"
	ColumnNameCheckSyncPercentage          ColumnName = "CHECKPOINT\nSYNC PERCENTAGE"
	ColumnNameCurrentEpoch                 ColumnName = "CURRENT\nEPOCH"
	ColumnNameNetworkPeers                 ColumnName = "NETWORK\nPEERS"
	ColumnNamePrimaryNetworkPeers          ColumnName = "PRIMARY\nNETWORK PEERS"
	ColumnNameWorkerNetworkPeers           ColumnName = "WORKER\nNETWORK PEERS"
	ColumnNameCurrentRound                 ColumnName = "CURRENT\nROUND"
	ColumnNameHighestProcessedRound        ColumnName = "HIGHEST\nPROCESSED ROUND"
	ColumnNameLastCommittedRound           ColumnName = "LAST COMMITTED\nROUND"
	ColumnNameSkippedConsensusTransactions ColumnName = "SKIPPED\nCONSENSUS TX"
	ColumnNameTotalSignatureErrors         ColumnName = "TOTAL\nSIGNATURE ERRORS"

	ColumnNameValidatorName                     ColumnName = "NAME"
	ColumnNameValidatorNetAddress               ColumnName = "NET ADDRESS"
	ColumnNameValidatorVotingPower              ColumnName = "VOTING POWER"
	ColumnNameValidatorGasPrice                 ColumnName = "GAS PRICE"
	ColumnNameValidatorCommissionRate           ColumnName = "COMMISSION RATE"
	ColumnNameValidatorNextEpochStake           ColumnName = "NEXT EPOCH STAKE"
	ColumnNameValidatorNextEpochGasPrice        ColumnName = "NEXT EPOCH GAS PRICE"
	ColumnNameValidatorNextEpochCommissionRate  ColumnName = "NEXT EPOCH COMMISSION RATE"
	ColumnNameValidatorStakingPoolSuiBalance    ColumnName = "STAKING POOL SUI BALANCE"
	ColumnNameValidatorRewardsPool              ColumnName = "REWARDS POOL"
	ColumnNameValidatorPoolTokenBalance         ColumnName = "POOL TOKEN BALANCE"
	ColumnNameValidatorPendingStake             ColumnName = "PENDING STAKE"
	ColumnNameValidatorPendingTotalSuiWithdraw  ColumnName = "PENDING TOTAL SUI WITHDRAW"
	ColumnNameValidatorPendingPoolTokenWithdraw ColumnName = "PENDING POOL TOKEN WITHDRAW"

	ColumnNameSystemEpoch                                 ColumnName = "EPOCH"
	ColumnNameSystemEpochStartTimestampMs                 ColumnName = "EPOCH START TIME MS"
	ColumnNameSystemEpochDurationMs                       ColumnName = "EPOCH DURATION MS"
	ColumnNameSystemTotalStake                            ColumnName = "TOTAL STAKE"
	ColumnNameSystemStorageFundTotalObjectStorageRebates  ColumnName = "STORAGE FUND TOTAL OBJECT STORAGE REBATES"
	ColumnNameSystemStorageFundNonRefundableBalance       ColumnName = "STORAGE FUND REFUNDABLE BALANCE"
	ColumnNameSystemReferenceGasPrice                     ColumnName = "REFERENCE GAS PRICE"
	ColumnNameSystemStakeSubsidyStartEpoch                ColumnName = "STAKE SUBSIDY START EPOCH"
	ColumnNameSystemMaxValidatorCount                     ColumnName = "MAX VALIDATOR COUNT"
	ColumnNameSystemMinValidatorJoiningStake              ColumnName = "MIN VALIDATOR JOINING STAKE"
	ColumnNameSystemValidatorLowStakeThreshold            ColumnName = "VALIDATOR LOW STAKE THRESHOLD"
	ColumnNameSystemValidatorVeryLowStakeThreshold        ColumnName = "VALIDATOR VERY LOW STAKE THRESHOLD"
	ColumnNameSystemValidatorLowStakeGracePeriod          ColumnName = "VALIDATOR LOW STAKE GRACE PERIOD"
	ColumnNameSystemStakeSubsidyBalance                   ColumnName = "STAKE SUBSIDY BALANCE"
	ColumnNameSystemStakeSubsidyDistributionCounter       ColumnName = "STAKE SUBSIDY DISTRIBUTION COUNTER"
	ColumnNameSystemStakeSubsidyCurrentDistributionAmount ColumnName = "STAKE SUBSIDY CURRENT DISTRIBUTION AMOUNT"
	ColumnNameSystemStakeSubsidyPeriodLength              ColumnName = "STAKE SUBSIDY PERIOD LENGTH"
	ColumnNameSystemStakeSubsidyDecreaseRate              ColumnName = "STAKE SUBSIDY DECREASE RATE"
	ColumnNameSystemActiveValidatorCount                  ColumnName = "ACTIVE VALIDATOR COUNT"
	ColumnNameSystemPendingActiveValidatorCount           ColumnName = "PENDING ACTIVE VALIDATORS COUNT"
	ColumnNameSystemPendingRemovalsCount                  ColumnName = "PENDING REMOVALS COUNT"
	ColumnNameSystemValidatorCandidateCount               ColumnName = "VALIDATOR CANDIDATE COUNT"
	ColumnNameSystemAtRiskValidatorCount                  ColumnName = "VALIDATOR AT RISK COUNT"
	ColumnNameSystemAtRiskValidatorName                   ColumnName = "VALIDATOR AT RISK NAME"
	ColumnNameSystemAtRiskValidatorSinceEpoch             ColumnName = "VALIDATOR AT RISK SINCE EPOCH"
	ColumnNameSystemValidatorReportReporter               ColumnName = "VALIDATOR REPORTER"
	ColumnNameSystemValidatorReportReported               ColumnName = "VALIDATOR REPORTED"
)
