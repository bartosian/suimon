package enums

type ColumnName string

// Overview section.
const (
	ColumnNameIndex   ColumnName = "IDX"
	ColumnNameHealth  ColumnName = "HEALTH"
	ColumnNameAddress ColumnName = "ADDRESS"
	ColumnNamePortRPC ColumnName = "RPC"
	ColumnNameUptime  ColumnName = "UPTIME DAYS"
	ColumnNameVersion ColumnName = "VERSION"
	ColumnNameCommit  ColumnName = "COMMIT"
	ColumnNameCountry ColumnName = "COUNTRY"
)

// Transactions section.
const (
	ColumnNameTotalTransactionBlocks                  ColumnName = "TOTAL TX\nBLOCKS"
	ColumnNameTotalTransactionCertificates            ColumnName = "TOTAL TX\nCERTIFICATES"
	ColumnNameTotalTransactionEffects                 ColumnName = "TOTAL TX\nEFFECTS"
	ColumnNameTXSyncPercentage                        ColumnName = "TX SYNC PCT"
	ColumnNameSkippedConsensusTransactions            ColumnName = "SKIPPED\nCONSENSUS TX"
	ColumnNameCertificatesCreated                     ColumnName = "CERTIFICATES\nCREATED"
	ColumnNameHandleCertificateNonConsensusLatencySum ColumnName = "CERTIFICATE\nNON CONSENSUS LATENCY"
	ColumnNameTotalSignatureErrors                    ColumnName = "SIGNATURE\nERRORS"
	ColumnNameTransactionsPerSecond                   ColumnName = "TRANSACTIONS PER SECOND"
	ColumnNameCertificatesPerSecond                   ColumnName = "CERTIFICATES PER SECOND"
)

// Checkpoints section.
const (
	ColumnNameLatestCheckpoint        ColumnName = "LATEST\nCHECKPOINT"
	ColumnNameHighestKnownCheckpoint  ColumnName = "HIGHEST KNOWN\nCHECKPOINT"
	ColumnNameLastExecutedCheckpoint  ColumnName = "LAST EXECUTED\nCHECKPOINT"
	ColumnNameHighestSyncedCheckpoint ColumnName = "HIGHEST SYNCED\nCHECKPOINT"
	ColumnNameCheckpointExecBacklog   ColumnName = "CHECKPOINT\nEXEC BACKLOG"
	ColumnNameCheckpointSyncBacklog   ColumnName = "CHECKPOINT\nSYNC BACKLOG"
	ColumnNameCheckSyncPercentage     ColumnName = "CHECKPOINT\nSYNC PCT"
	ColumnNameCheckpointsPerSecond    ColumnName = "CHECKPOINTS PER SECOND"
)

// Rounds section.
const (
	ColumnNameCurrentRound          ColumnName = "CURRENT\nROUND"
	ColumnNameHighestProcessedRound ColumnName = "HIGHEST\nPROCESSED ROUND"
	ColumnNameLastCommittedRound    ColumnName = "LAST COMMITTED\nROUND"
	ColumnNameRoundsPerSecond       ColumnName = "ROUNDS PER SECOND"
)

// Peers section.
const (
	ColumnNameNetworkPeers        ColumnName = "NETWORK\nPEERS"
	ColumnNamePrimaryNetworkPeers ColumnName = "PRIMARY\nNETWORK PEERS"
	ColumnNameWorkerNetworkPeers  ColumnName = "WORKER\nNETWORK PEERS"
)

// Validator section.
const (
	ColumnNameValidatorName                     ColumnName = "NAME"
	ColumnNameValidatorNetAddress               ColumnName = "NET ADDRESS"
	ColumnNameValidatorVotingPower              ColumnName = "VOTING\nPOWER"
	ColumnNameValidatorGasPrice                 ColumnName = "GAS\nPRICE"
	ColumnNameValidatorCommissionRate           ColumnName = "COMMISSION\nRATE"
	ColumnNameValidatorNextEpochStake           ColumnName = "NEXT EPOCH\nSTAKE, SUI"
	ColumnNameValidatorNextEpochGasPrice        ColumnName = "NEXT EPOCH\nGAS PRICE"
	ColumnNameValidatorNextEpochCommissionRate  ColumnName = "NEXT EPOCH\nCOMMISSION RATE"
	ColumnNameValidatorStakingPoolSuiBalance    ColumnName = "STAKING POOL SUI\nBALANCE, SUI"
	ColumnNameValidatorRewardsPool              ColumnName = "REWARDS POOL, SUI"
	ColumnNameValidatorPoolTokenBalance         ColumnName = "POOL TOKEN\nBALANCE, SUI"
	ColumnNameValidatorPendingStake             ColumnName = "PENDING STAKE, SUI"
	ColumnNameValidatorPendingTotalSuiWithdraw  ColumnName = "PENDING TOTAL\nSUI WITHDRAW"
	ColumnNameValidatorPendingPoolTokenWithdraw ColumnName = "PENDING POOL\nTOKEN WITHDRAW"
	ColumnNameValidatorApy                      ColumnName = "APY, %"
)

// Epoch section.
const (
	ColumnNameEpoch                             ColumnName = "EPOCH"
	ColumnNameEpochTotalTransactions            ColumnName = "TOTAL\nTRANSACTIONS"
	ColumnNameEpochStartTimestamp               ColumnName = "EPOCH START TIME UTC"
	ColumnNameEpochEndTimestamp                 ColumnName = "EPOCH END TIME UTC"
	ColumnNameEpochFirstCheckpointID            ColumnName = "FIRST\nCHECKPOINT ID"
	ColumnNameEpochLastCheckpointID             ColumnName = "LAST\nCHECKPOINT ID"
	ColumnNameEpochProtocolVersion              ColumnName = "PROTOCOL\nVERSION"
	ColumnNameEpochReferenceGasPrice            ColumnName = "REFERENCE\nGAS PRICE"
	ColumnNameEpochTotalStake                   ColumnName = "TOTAL STAKE,\nSUI"
	ColumnNameEpochStorageFundReinvestment      ColumnName = "STORAGE FUND\nREINVESTMENT, SUI"
	ColumnNameEpochStorageCharge                ColumnName = "STORAGE\nCHARGE, SUI"
	ColumnNameEpochStorageRebate                ColumnName = "STORAGE\nREBATE, SUI"
	ColumnNameEpochStorageFundBalance           ColumnName = "STORAGE FUND\nBALANCE, SUI"
	ColumnNameEpochStakeSubsidyAmount           ColumnName = "STORAGE SUBSIDY\nAMOUNT, SUI"
	ColumnNameEpochTotalGasFees                 ColumnName = "TOTAL GAS\nFEES, SUI"
	ColumnNameEpochTotalStakeRewardsDistributed ColumnName = "TOTAL STAKE REWARDS\nDISTRIBUTED, SUI"
	ColumnNameEpochLeftoverStorageFundInflow    ColumnName = "LEFTOVER STORAGE FUND\nINFLOW, SUI"
)

// nolint: dupl,nolintlint,gocritic // they are not duplicate
const (
	ColumnNameCurrentEpoch                                ColumnName = "CURRENT\nEPOCH"
	ColumnNameSystemEpoch                                 ColumnName = "EPOCH"
	ColumnNameSystemEpochStartTimestamp                   ColumnName = "EPOCH START TIME UTC"
	ColumnNameSystemEpochDuration                         ColumnName = "EPOCH\nDURATION"
	ColumnNameSystemTimeTillNextEpoch                     ColumnName = "TIME TILL\nNEXT EPOCH"
	ColumnNameSystemTotalStake                            ColumnName = "TOTAL STAKE, SUI"
	ColumnNameSystemStorageFundTotalObjectStorageRebates  ColumnName = "STORAGE FUND TOTAL\nOBJECT REBATES, SUI"
	ColumnNameSystemStorageFundNonRefundableBalance       ColumnName = "STORAGE FUND\nREFUNDABLE BALANCE, SUI"
	ColumnNameSystemReferenceGasPrice                     ColumnName = "REFERENCE\nGAS PRICE"
	ColumnNameSystemMinReferenceGasPrice                  ColumnName = "MIN REFERENCE\nGAS PRICE"
	ColumnNameSystemMaxReferenceGasPrice                  ColumnName = "MAX REFERENCE\nGAS PRICE"
	ColumnNameSystemMeanReferenceGasPrice                 ColumnName = "MEAN REFERENCE\nGAS PRICE"
	ColumnNameSystemStakeWeightedMeanReferenceGasPrice    ColumnName = "STAKE WEIGHTED MEAN\nREFERENCE GAS PRICE"
	ColumnNameSystemMedianReferenceGasPrice               ColumnName = "MEDIAN REFERENCE\nGAS PRICE"
	ColumnNameSystemEstimatedReferenceGasPrice            ColumnName = "ESTIMATED REFERENCE\nGAS PRICE"
	ColumnNameSystemMaxValidatorCount                     ColumnName = "MAX VALIDATOR\nCOUNT"
	ColumnNameSystemActiveValidatorCount                  ColumnName = "ACTIVE VALIDATOR\nCOUNT"
	ColumnNameSystemPendingActiveValidatorCount           ColumnName = "PENDING ACTIVE\nVALIDATORS COUNT"
	ColumnNameSystemValidatorCandidateCount               ColumnName = "VALIDATOR\nCANDIDATE COUNT"
	ColumnNameSystemPendingRemovalsCount                  ColumnName = "PENDING VALIDATOR\nREMOVALS COUNT"
	ColumnNameSystemAtRiskValidatorCount                  ColumnName = "VALIDATOR AT RISK\nCOUNT"
	ColumnNameSystemMinValidatorJoiningStake              ColumnName = "MIN VALIDATOR\nJOINING STAKE"
	ColumnNameSystemValidatorLowStakeThreshold            ColumnName = "VALIDATOR LOW\nSTAKE THRESHOLD"
	ColumnNameSystemValidatorVeryLowStakeThreshold        ColumnName = "VALIDATOR VERY LOW\nSTAKE THRESHOLD"
	ColumnNameSystemValidatorLowStakeGracePeriod          ColumnName = "VALIDATOR LOW STAKE\nGRACE PERIOD"
	ColumnNameSystemAtRiskValidatorName                   ColumnName = "VALIDATOR NAME"
	ColumnNameSystemAtRiskValidatorAddress                ColumnName = "VALIDATOR ADDRESS"
	ColumnNameSystemAtRiskValidatorNumberOfEpochs         ColumnName = "NUMBER OF EPOCHS\nAT RISK"
	ColumnNameSystemValidatorReporterName                 ColumnName = "REPORTER VALIDATOR"
	ColumnNameSystemValidatorReporterAddress              ColumnName = "REPORTER VALIDATOR ADDRESS"
	ColumnNameSystemValidatorReportedName                 ColumnName = "REPORTED VALIDATOR"
	ColumnNameSystemValidatorSlashingPercentage           ColumnName = "SLASHING PCT"
	ColumnNameSystemStakeSubsidyStartEpoch                ColumnName = "STAKE SUBSIDY\nSTART EPOCH"
	ColumnNameSystemStakeSubsidyBalance                   ColumnName = "STAKE SUBSIDY\nBALANCE, SUI"
	ColumnNameSystemStakeSubsidyDistributionCounter       ColumnName = "STAKE SUBSIDY\nDISTRIBUTION COUNTER"
	ColumnNameSystemStakeSubsidyCurrentDistributionAmount ColumnName = "STAKE SUBSIDY\nDISTRIBUTION AMOUNT, SUI"
	ColumnNameSystemStakeSubsidyPeriodLength              ColumnName = "STAKE SUBSIDY\nPERIOD LENGTH"
	ColumnNameSystemStakeSubsidyDecreaseRate              ColumnName = "STAKE SUBSIDY\nDECREASE RATE"
)

// Protocol section.
const (
	ColumnNameMinSupportedProtocolVersion ColumnName = "MIN SUPPORTED\nPROTOCOL VERSION"
	ColumnNameMaxSupportedProtocolVersion ColumnName = "MAX SUPPORTED\nPROTOCOL VERSION"
	ColumnNameProtocolVersion             ColumnName = "PROTOCOL VERSION"
)

// nolint: dupl,nolintlint,gocritic // they are not duplicate
const (
	ColumnNameFeatureFlagAcceptZkloginInMultisig                  ColumnName = "ACCEPT ZKLOGIN\nIN MULTISIG"
	ColumnNameFeatureFlagAdvanceEpochStartTimeInSafeMode          ColumnName = "ADVANCE EPOCH\nSTART TIME\nIN SAFE MODE"
	ColumnNameFeatureFlagAdvanceToHighestSupportedProtocolVersion ColumnName = "ADVANCE TO\nHIGHEST SUPPORTED\nPROTOCOL VERSION"
	ColumnNameFeatureFlagAllowReceivingObjectID                   ColumnName = "ALLOW RECEIVING\nOBJECT ID"
	ColumnNameFeatureFlagBanEntryInit                             ColumnName = "BAN ENTRY INIT"
	ColumnNameFeatureFlagCommitRootStateDigest                    ColumnName = "COMMIT ROOT\nSTATE DIGEST"
	ColumnNameFeatureFlagConsensusOrderEndOfEpochLast             ColumnName = "CONSENSUS ORDER\nEND OF EPOCH LAST"
	ColumnNameFeatureFlagDisableInvariantViolationCheckInSwapLoc  ColumnName = "DISABLE INVARIANT\nVIOLATION CHECK\nIN SWAP LOC"
	ColumnNameFeatureFlagDisallowAddingAbilitiesOnUpgrade         ColumnName = "DISALLOW ADDING\nABILITIES\nON UPGRADE"
	ColumnNameFeatureFlagDisallowChangeStructTypeParamsOnUpgrade  ColumnName = "DISALLOW CHANGE\nSTRUCT TYPE PARAMS\nON UPGRADE"
	ColumnNameFeatureFlagEnableEffectsV2                          ColumnName = "ENABLE\nEFFECTS V2"
	ColumnNameFeatureFlagEnableJwkConsensusUpdates                ColumnName = "ENABLE JWK\nCONSENSUS\nUPDATES"
	ColumnNameFeatureFlagEndOfEpochTransactionSupported           ColumnName = "END OF EPOCH\nTRANSACTION\nSUPPORTED"
	ColumnNameFeatureFlagHardenedOtwCheck                         ColumnName = "HARDENED\nOTW CHECK"
	ColumnNameFeatureFlagIncludeConsensusDigestInPrologue         ColumnName = "INCLUDE\nCONSENSUS DIGEST\nIN PROLOGUE"
	ColumnNameFeatureFlagLoadedChildObjectFormat                  ColumnName = "LOADED CHILD\nOBJECT FORMAT"
	ColumnNameFeatureFlagLoadedChildObjectFormatType              ColumnName = "LOADED CHILD\nOBJECT FORMAT\nTYPE"
	ColumnNameFeatureFlagLoadedChildObjectsFixed                  ColumnName = "LOADED CHILD\nOBJECTS FIXED"
	ColumnNameFeatureFlagMissingTypeIsCompatibilityError          ColumnName = "MISSING TYPE\nIS COMPATIBILITY\nERROR"
	ColumnNameFeatureFlagNarwhalCertificateV2                     ColumnName = "NARWHAL\nCERTIFICATE V2"
	ColumnNameFeatureFlagNarwhalHeaderV2                          ColumnName = "NARWHAL\nHEADER V2"
	ColumnNameFeatureFlagNarwhalNewLeaderElectionSchedule         ColumnName = "NARWHAL\nNEW LEADER\nELECTION SCHEDULE"
	ColumnNameFeatureFlagNarwhalVersionedMetadata                 ColumnName = "NARWHAL\nVERSIONED METADATA"
	ColumnNameFeatureFlagNoExtraneousModuleBytes                  ColumnName = "NO EXTRANEOUS\nMODULE BYTES"
	ColumnNameFeatureFlagPackageDigestHashModule                  ColumnName = "PACKAGE\nDIGEST HASH\nMODULE"
	ColumnNameFeatureFlagPackageUpgrades                          ColumnName = "PACKAGE\nUPGRADES"
	ColumnNameFeatureFlagRandomBeacon                             ColumnName = "RANDOM\nBEACON"
	ColumnNameFeatureFlagReceiveObjects                           ColumnName = "RECEIVE\nOBJECTS"
	ColumnNameFeatureFlagRecomputeHasPublicTransferInExecution    ColumnName = "RECOMPUTE HAS\nPUBLIC TRANSFER\nIN EXECUTION"
	ColumnNameFeatureFlagScoringDecisionWithValidityCutoff        ColumnName = "SCORING DECISION\nWITH VALIDITY\nCUTOFF"
	ColumnNameFeatureFlagSharedObjectDeletion                     ColumnName = "SHARED OBJECT\nDELETION"
	ColumnNameFeatureFlagSimpleConservationChecks                 ColumnName = "SIMPLE\nCONSERVATION\nCHECKS"
	ColumnNameFeatureFlagSimplifiedUnwrapThenDelete               ColumnName = "SIMPLIFIED\nUNWRAP\nTHEN DELETE"
	ColumnNameFeatureFlagThroughputAwareConsensusSubmission       ColumnName = "THROUGHPUT AWARE\nCONSENSUS\nSUBMISSION"
	ColumnNameFeatureFlagTxnBaseCostAsMultiplier                  ColumnName = "TXN BASE COST\nAS MULTIPLIER"
	ColumnNameFeatureFlagUpgradedMultisigSupported                ColumnName = "UPGRADED MULTISIG\nSUPPORTED"
	ColumnNameFeatureFlagVerifyLegacyZkloginAddress               ColumnName = "VERIFY LEGACY ZK\nLOGIN ADDRESS"
	ColumnNameFeatureFlagZkLoginAuth                              ColumnName = "ZK LOGIN AUTH"
)

// Release section.
const (
	ColumnNameTagName     ColumnName = "TAG"
	ColumnNameCommitHash  ColumnName = "TARGET"
	ColumnNameReleaseName ColumnName = "RELEASE"
	ColumnNameDraft       ColumnName = "DRAFT"
	ColumnNamePreRelease  ColumnName = "PRERELEASE"
	ColumnNameCreatedAt   ColumnName = "CREATED AT"
	ColumnNamePublishedAt ColumnName = "PUBLISHED AT"
	ColumnNameAuthor      ColumnName = "AUTHOR"
	ColumnNameURL         ColumnName = "RELEASE URL"
)

func (e ColumnName) ToString() string {
	return string(e)
}
