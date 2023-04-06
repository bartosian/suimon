package metrics

const (
	TransactionsPerSecondWindow     = 10
	CheckpointsPerSecondWindow      = 10
	TransactionsPerSecondLag        = 5
	TotalTransactionsLag            = 100
	CheckpointsPerSecondLag         = 10
	LatestCheckpointLag             = 30
	HighestSyncedCheckpointLag      = 30
	TotalTransactionsSyncPercentage = 99
	TotalCheckpointsSyncPercentage  = 99
)

type (
	Validator struct {
		SuiAddress                   string      `json:"suiAddress"`
		ProtocolPubkeyBytes          string      `json:"protocolPubkeyBytes"`
		NetworkPubkeyBytes           string      `json:"networkPubkeyBytes"`
		WorkerPubkeyBytes            string      `json:"workerPubkeyBytes"`
		ProofOfPossessionBytes       string      `json:"proofOfPossessionBytes"`
		Name                         string      `json:"name"`
		Description                  string      `json:"description"`
		ImageURL                     string      `json:"imageUrl"`
		ProjectURL                   string      `json:"projectUrl"`
		NetAddress                   string      `json:"netAddress"`
		P2PAddress                   string      `json:"p2pAddress"`
		PrimaryAddress               string      `json:"primaryAddress"`
		WorkerAddress                string      `json:"workerAddress"`
		NextEpochProtocolPubkeyBytes interface{} `json:"nextEpochProtocolPubkeyBytes"`
		NextEpochProofOfPossession   interface{} `json:"nextEpochProofOfPossession"`
		NextEpochNetworkPubkeyBytes  interface{} `json:"nextEpochNetworkPubkeyBytes"`
		NextEpochWorkerPubkeyBytes   interface{} `json:"nextEpochWorkerPubkeyBytes"`
		NextEpochNetAddress          interface{} `json:"nextEpochNetAddress"`
		NextEpochP2PAddress          interface{} `json:"nextEpochP2pAddress"`
		NextEpochPrimaryAddress      interface{} `json:"nextEpochPrimaryAddress"`
		NextEpochWorkerAddress       interface{} `json:"nextEpochWorkerAddress"`
		VotingPower                  int         `json:"votingPower"`
		OperationCapID               string      `json:"operationCapId"`
		GasPrice                     int         `json:"gasPrice"`
		CommissionRate               int         `json:"commissionRate"`
		NextEpochStake               int         `json:"nextEpochStake"`
		NextEpochGasPrice            int         `json:"nextEpochGasPrice"`
		NextEpochCommissionRate      int         `json:"nextEpochCommissionRate"`
		StakingPoolID                string      `json:"stakingPoolId"`
		StakingPoolActivationEpoch   int         `json:"stakingPoolActivationEpoch"`
		StakingPoolDeactivationEpoch interface{} `json:"stakingPoolDeactivationEpoch"`
		StakingPoolSuiBalance        int         `json:"stakingPoolSuiBalance"`
		RewardsPool                  int         `json:"rewardsPool"`
		PoolTokenBalance             int         `json:"poolTokenBalance"`
		PendingStake                 int         `json:"pendingStake"`
		PendingTotalSuiWithdraw      int         `json:"pendingTotalSuiWithdraw"`
		PendingPoolTokenWithdraw     int         `json:"pendingPoolTokenWithdraw"`
		ExchangeRatesID              string      `json:"exchangeRatesId"`
		ExchangeRatesSize            int         `json:"exchangeRatesSize"`
	}

	SuiSystemState struct {
		Epoch                                 int             `json:"epoch"`
		ProtocolVersion                       int             `json:"protocolVersion"`
		SystemStateVersion                    int             `json:"systemStateVersion"`
		StorageFundTotalObjectStorageRebates  int             `json:"storageFundTotalObjectStorageRebates"`
		StorageFundNonRefundableBalance       int             `json:"storageFundNonRefundableBalance"`
		ReferenceGasPrice                     int             `json:"referenceGasPrice"`
		SafeMode                              bool            `json:"safeMode"`
		SafeModeStorageRewards                int             `json:"safeModeStorageRewards"`
		SafeModeComputationRewards            int             `json:"safeModeComputationRewards"`
		SafeModeStorageRebates                int             `json:"safeModeStorageRebates"`
		SafeModeNonRefundableStorageFee       int             `json:"safeModeNonRefundableStorageFee"`
		EpochStartTimestampMs                 int64           `json:"epochStartTimestampMs"`
		EpochDurationMs                       int64           `json:"epochDurationMs"`
		StakeSubsidyStartEpoch                int             `json:"stakeSubsidyStartEpoch"`
		MaxValidatorCount                     int             `json:"maxValidatorCount"`
		MinValidatorJoiningStake              int             `json:"minValidatorJoiningStake"`
		ValidatorLowStakeThreshold            int             `json:"validatorLowStakeThreshold"`
		ValidatorVeryLowStakeThreshold        int             `json:"validatorVeryLowStakeThreshold"`
		ValidatorLowStakeGracePeriod          int             `json:"validatorLowStakeGracePeriod"`
		StakeSubsidyBalance                   int             `json:"stakeSubsidyBalance"`
		StakeSubsidyDistributionCounter       int             `json:"stakeSubsidyDistributionCounter"`
		StakeSubsidyCurrentDistributionAmount int             `json:"stakeSubsidyCurrentDistributionAmount"`
		StakeSubsidyPeriodLength              int             `json:"stakeSubsidyPeriodLength"`
		StakeSubsidyDecreaseRate              int             `json:"stakeSubsidyDecreaseRate"`
		TotalStake                            int             `json:"totalStake"`
		ActiveValidators                      []Validator     `json:"activeValidators"`
		PendingActiveValidatorsID             string          `json:"pendingActiveValidatorsId"`
		PendingActiveValidatorsSize           int             `json:"pendingActiveValidatorsSize"`
		PendingRemovals                       []interface{}   `json:"pendingRemovals"`
		StakingPoolMappingsID                 string          `json:"stakingPoolMappingsId"`
		StakingPoolMappingsSize               int             `json:"stakingPoolMappingsSize"`
		InactivePoolsID                       string          `json:"inactivePoolsId"`
		InactivePoolsSize                     int             `json:"inactivePoolsSize"`
		ValidatorCandidatesID                 string          `json:"validatorCandidatesId"`
		ValidatorCandidatesSize               int             `json:"validatorCandidatesSize"`
		AtRiskValidators                      [][]interface{} `json:"atRiskValidators"`
		ValidatorReportRecords                [][]interface{} `json:"validatorReportRecords"`
		AddressToValidatorName                map[string]string
	}

	Transactions struct {
		TotalTransactionsBlocks      int
		TotalTransactionCertificates int
		TotalTransactionEffects      int
		TransactionsPerSecond        int
		TxSyncPercentage             int
		TransactionsHistory          []int
	}

	Checkpoints struct {
		LatestCheckpoint        int
		HighestKnownCheckpoint  int
		HighestSyncedCheckpoint int
		LastExecutedCheckpoint  int
		CheckpointsPerSecond    int
		CheckpointExecBacklog   int
		CheckpointSyncBacklog   int
		CheckSyncPercentage     int
		CheckpointsHistory      []int
	}

	Rounds struct {
		CurrentRound          int
		HighestProcessedRound int
		LastCommittedRound    int
	}

	Peers struct {
		NetworkPeers        int
		PrimaryNetworkPeers int
		WorkerNetworkPeers  int
	}

	Epoch struct {
		CurrentEpoch       int
		EpochTotalDuration int
		EpochPercentage    int
		TimeTillNextEpoch  int64
	}

	Errors struct {
		SkippedConsensusTransactions int
		TotalSignatureErrors         int
	}

	Metrics struct {
		Updated bool

		SystemState SuiSystemState

		Uptime  string
		Version string
		Commit  string

		Transactions
		Checkpoints
		Rounds
		Peers
		Epoch
		Errors
	}
)
