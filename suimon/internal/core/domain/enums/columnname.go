package enums

type ColumnName string

const (
	ColumnNameIndex   ColumnName = "IDX"
	ColumnNameHealth  ColumnName = "HEALTH"
	ColumnNameAddress ColumnName = "ADDRESS"
	ColumnNamePortRPC ColumnName = "RPC"
	ColumnNameUptime  ColumnName = "UPTIME DAYS"
	ColumnNameVersion ColumnName = "VERSION"
	ColumnNameCommit  ColumnName = "COMMIT"
	ColumnNameCountry ColumnName = "COUNTRY"

	ColumnNameTotalTransactionBlocks       ColumnName = "TOTAL TX\nBLOCKS"
	ColumnNameTotalTransactionCertificates ColumnName = "TOTAL TX\nCERTIFICATES"
	ColumnNameTotalTransactionEffects      ColumnName = "TOTAL TX\nEFFECTS"
	ColumnNameTXSyncPercentage             ColumnName = "TX SYNC PCT"
	ColumnNameLatestCheckpoint             ColumnName = "LATEST\nCHECKPOINT"
	ColumnNameHighestKnownCheckpoint       ColumnName = "HIGHEST KNOWN\nCHECKPOINT"
	ColumnNameLastExecutedCheckpoint       ColumnName = "LAST EXECUTED\nCHECKPOINT"
	ColumnNameHighestSyncedCheckpoint      ColumnName = "HIGHEST SYNCED\nCHECKPOINT"
	ColumnNameCheckpointExecBacklog        ColumnName = "CHECKPOINT\nEXEC BACKLOG"
	ColumnNameCheckpointSyncBacklog        ColumnName = "CHECKPOINT\nSYNC BACKLOG"
	ColumnNameCheckSyncPercentage          ColumnName = "CHECKPOINT\nSYNC PCT"
	ColumnNameCurrentEpoch                 ColumnName = "CURRENT\nEPOCH"
	ColumnNameNetworkPeers                 ColumnName = "NETWORK\nPEERS"
	ColumnNamePrimaryNetworkPeers          ColumnName = "PRIMARY\nNETWORK PEERS"
	ColumnNameWorkerNetworkPeers           ColumnName = "WORKER\nNETWORK PEERS"
	ColumnNameCurrentRound                 ColumnName = "CURRENT\nROUND"
	ColumnNameHighestProcessedRound        ColumnName = "HIGHEST\nPROCESSED ROUND"
	ColumnNameLastCommittedRound           ColumnName = "LAST COMMITTED\nROUND"
	ColumnNameSkippedConsensusTransactions ColumnName = "SKIPPED\nCONSENSUS TX"
	ColumnNameTotalSignatureErrors         ColumnName = "SIGNATURE\nERRORS"

	ColumnNameValidatorName                     ColumnName = "NAME"
	ColumnNameValidatorNetAddress               ColumnName = "NET ADDRESS"
	ColumnNameValidatorVotingPower              ColumnName = "VOTING\nPOWER"
	ColumnNameValidatorGasPrice                 ColumnName = "GAS\nPRICE"
	ColumnNameValidatorCommissionRate           ColumnName = "COMMISSION\nRATE"
	ColumnNameValidatorNextEpochStake           ColumnName = "NEXT EPOCH STAKE"
	ColumnNameValidatorNextEpochGasPrice        ColumnName = "NEXT EPOCH GAS\nPRICE"
	ColumnNameValidatorNextEpochCommissionRate  ColumnName = "NEXT EPOCH\nCOMMISSION RATE"
	ColumnNameValidatorStakingPoolSuiBalance    ColumnName = "STAKING POOL SUI\nBALANCE"
	ColumnNameValidatorRewardsPool              ColumnName = "REWARDS POOL"
	ColumnNameValidatorPoolTokenBalance         ColumnName = "POOL TOKEN\nBALANCE"
	ColumnNameValidatorPendingStake             ColumnName = "PENDING STAKE"
	ColumnNameValidatorPendingTotalSuiWithdraw  ColumnName = "PENDING TOTAL\nSUI WITHDRAW"
	ColumnNameValidatorPendingPoolTokenWithdraw ColumnName = "PENDING POOL\nTOKEN WITHDRAW"

	ColumnNameSystemEpoch                                ColumnName = "EPOCH"
	ColumnNameSystemEpochStartTimestamp                  ColumnName = "EPOCH START TIME UTC"
	ColumnNameSystemEpochDuration                        ColumnName = "EPOCH DURATION"
	ColumnNameSystemTotalStake                           ColumnName = "TOTAL STAKE"
	ColumnNameSystemStorageFundTotalObjectStorageRebates ColumnName = "STORAGE FUND TOTAL\nOBJECT REBATES"
	ColumnNameSystemStorageFundNonRefundableBalance      ColumnName = "STORAGE FUND\nREFUNDABLE BALANCE"
	ColumnNameSystemReferenceGasPrice                    ColumnName = "REFERENCE\nGAS PRICE"

	ColumnNameSystemMaxValidatorCount              ColumnName = "MAX VALIDATOR\nCOUNT"
	ColumnNameSystemActiveValidatorCount           ColumnName = "ACTIVE VALIDATOR\nCOUNT"
	ColumnNameSystemPendingActiveValidatorCount    ColumnName = "PENDING ACTIVE\nVALIDATORS COUNT"
	ColumnNameSystemValidatorCandidateCount        ColumnName = "VALIDATOR\nCANDIDATE COUNT"
	ColumnNameSystemPendingRemovalsCount           ColumnName = "PENDING VALIDATOR\nREMOVALS COUNT"
	ColumnNameSystemAtRiskValidatorCount           ColumnName = "VALIDATOR AT RISK\nCOUNT"
	ColumnNameSystemMinValidatorJoiningStake       ColumnName = "MIN VALIDATOR\nJOINING STAKE"
	ColumnNameSystemValidatorLowStakeThreshold     ColumnName = "VALIDATOR LOW\nSTAKE THRESHOLD"
	ColumnNameSystemValidatorVeryLowStakeThreshold ColumnName = "VALIDATOR VERY LOW\nSTAKE THRESHOLD"
	ColumnNameSystemValidatorLowStakeGracePeriod   ColumnName = "VALIDATOR LOW STAKE\nGRACE PERIOD"

	ColumnNameSystemAtRiskValidatorName           ColumnName = "VALIDATOR NAME"
	ColumnNameSystemAtRiskValidatorAddress        ColumnName = "VALIDATOR ADDRESS"
	ColumnNameSystemAtRiskValidatorNumberOfEpochs ColumnName = "NUMBER OF EPOCHS\nAT RISK"

	ColumnNameSystemValidatorReporterName    ColumnName = "VALIDATOR REPORTER NAME"
	ColumnNameSystemValidatorReporterAddress ColumnName = "VALIDATOR REPORTER ADDRESS"
	ColumnNameSystemValidatorReportedName    ColumnName = "VALIDATOR REPORTED NAME"
	ColumnNameSystemValidatorReportedAddress ColumnName = "VALIDATOR REPORTED ADDRESS"

	ColumnNameSystemStakeSubsidyStartEpoch                ColumnName = "STAKE SUBSIDY\nSTART EPOCH"
	ColumnNameSystemStakeSubsidyBalance                   ColumnName = "STAKE SUBSIDY\nBALANCE"
	ColumnNameSystemStakeSubsidyDistributionCounter       ColumnName = "STAKE SUBSIDY\nDISTRIBUTION COUNTER"
	ColumnNameSystemStakeSubsidyCurrentDistributionAmount ColumnName = "STAKE SUBSIDY\nDISTRIBUTION AMOUNT"
	ColumnNameSystemStakeSubsidyPeriodLength              ColumnName = "STAKE SUBSIDY\nPERIOD LENGTH"
	ColumnNameSystemStakeSubsidyDecreaseRate              ColumnName = "STAKE SUBSIDY\nDECREASE RATE"
)

func (e ColumnName) ToString() string {
	return string(e)
}
