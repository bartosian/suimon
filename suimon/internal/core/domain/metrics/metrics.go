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
	// Validators represents a validator nodes on the Sui blockchain network.
	Validators []Validator
	Validator  struct {
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
		VotingPower                  string      `json:"votingPower"`
		OperationCapID               string      `json:"operationCapId"`
		GasPrice                     string      `json:"gasPrice"`
		CommissionRate               string      `json:"commissionRate"`
		NextEpochStake               string      `json:"nextEpochStake"`
		NextEpochGasPrice            string      `json:"nextEpochGasPrice"`
		NextEpochCommissionRate      string      `json:"nextEpochCommissionRate"`
		StakingPoolID                string      `json:"stakingPoolId"`
		StakingPoolActivationEpoch   string      `json:"stakingPoolActivationEpoch"`
		StakingPoolDeactivationEpoch interface{} `json:"stakingPoolDeactivationEpoch"`
		StakingPoolSuiBalance        string      `json:"stakingPoolSuiBalance"`
		RewardsPool                  string      `json:"rewardsPool"`
		PoolTokenBalance             string      `json:"poolTokenBalance"`
		PendingStake                 string      `json:"pendingStake"`
		PendingTotalSuiWithdraw      string      `json:"pendingTotalSuiWithdraw"`
		PendingPoolTokenWithdraw     string      `json:"pendingPoolTokenWithdraw"`
		ExchangeRatesID              string      `json:"exchangeRatesId"`
		ExchangeRatesSize            string      `json:"exchangeRatesSize"`
	}

	// SuiSystemState represents the current state of the Sui blockchain system.
	SuiSystemState struct {
		Epoch                                 string          `json:"epoch"`
		ProtocolVersion                       string          `json:"protocolVersion"`
		SystemStateVersion                    string          `json:"systemStateVersion"`
		StorageFundTotalObjectStorageRebates  string          `json:"storageFundTotalObjectStorageRebates"`
		StorageFundNonRefundableBalance       string          `json:"storageFundNonRefundableBalance"`
		ReferenceGasPrice                     string          `json:"referenceGasPrice"`
		SafeMode                              bool            `json:"safeMode"`
		SafeModeStorageRewards                string          `json:"safeModeStorageRewards"`
		SafeModeComputationRewards            string          `json:"safeModeComputationRewards"`
		SafeModeStorageRebates                string          `json:"safeModeStorageRebates"`
		SafeModeNonRefundableStorageFee       string          `json:"safeModeNonRefundableStorageFee"`
		EpochStartTimestampMs                 string          `json:"epochStartTimestampMs"`
		EpochDurationMs                       string          `json:"epochDurationMs"`
		StakeSubsidyStartEpoch                string          `json:"stakeSubsidyStartEpoch"`
		MaxValidatorCount                     string          `json:"maxValidatorCount"`
		MinValidatorJoiningStake              string          `json:"minValidatorJoiningStake"`
		ValidatorLowStakeThreshold            string          `json:"validatorLowStakeThreshold"`
		ValidatorVeryLowStakeThreshold        string          `json:"validatorVeryLowStakeThreshold"`
		ValidatorLowStakeGracePeriod          string          `json:"validatorLowStakeGracePeriod"`
		StakeSubsidyBalance                   string          `json:"stakeSubsidyBalance"`
		StakeSubsidyDistributionCounter       string          `json:"stakeSubsidyDistributionCounter"`
		StakeSubsidyCurrentDistributionAmount string          `json:"stakeSubsidyCurrentDistributionAmount"`
		StakeSubsidyPeriodLength              string          `json:"stakeSubsidyPeriodLength"`
		StakeSubsidyDecreaseRate              int             `json:"stakeSubsidyDecreaseRate"`
		TotalStake                            string          `json:"totalStake"`
		ActiveValidators                      Validators      `json:"activeValidators"`
		PendingActiveValidatorsID             string          `json:"pendingActiveValidatorsId"`
		PendingActiveValidatorsSize           string          `json:"pendingActiveValidatorsSize"`
		PendingRemovals                       []interface{}   `json:"pendingRemovals"`
		StakingPoolMappingsID                 string          `json:"stakingPoolMappingsId"`
		StakingPoolMappingsSize               string          `json:"stakingPoolMappingsSize"`
		InactivePoolsID                       string          `json:"inactivePoolsId"`
		InactivePoolsSize                     string          `json:"inactivePoolsSize"`
		ValidatorCandidatesID                 string          `json:"validatorCandidatesId"`
		ValidatorCandidatesSize               string          `json:"validatorCandidatesSize"`
		AtRiskValidators                      [][]interface{} `json:"atRiskValidators"`
		ValidatorReportRecords                [][]interface{} `json:"validatorReportRecords"`
		AddressToValidatorName                map[string]string
		ValidatorsAtRiskParsed                []ValidatorAtRisk
		ValidatorReportsParsed                []ValidatorReport
	}

	// Transactions represents information about transactions on the Sui blockchain network.
	Transactions struct {
		TotalTransactionsBlocks      int
		TotalTransactionCertificates int
		CertificatesCreated          int
		NonConsensusLatency          int
		TotalTransactionEffects      int
		TransactionsPerSecond        int
		TxSyncPercentage             int
		TransactionsHistory          []int
	}

	// Checkpoints represents information about checkpoints on the Sui blockchain network.
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

	// Rounds represents information about rounds on the Sui blockchain network.
	Rounds struct {
		CurrentRound          int
		HighestProcessedRound int
		LastCommittedRound    int
	}

	// Peers represents information about peers on the Sui blockchain network.
	Peers struct {
		NetworkPeers        int
		PrimaryNetworkPeers int
		WorkerNetworkPeers  int
	}

	// Epoch represents information about the current epoch on the Sui blockchain network.
	Epoch struct {
		CurrentEpoch             int
		EpochStartTimeUTC        string
		EpochTotalDuration       int
		EpochDurationHHMM        string
		DurationTillEpochEndHHMM string
		EpochPercentage          int
		TimeTillNextEpoch        int64
	}

	// Errors represents information about errors on the Sui blockchain network.
	Errors struct {
		SkippedConsensusTransactions int
		TotalSignatureErrors         int
	}

	// ValidatorReport represents a report about a validator node on the Sui blockchain network.
	ValidatorReport struct {
		ReportedName    string
		ReportedAddress string
		ReporterName    string
		ReporterAddress string
	}

	// ValidatorAtRisk represents a validator node at risk on the Sui blockchain network.
	ValidatorAtRisk struct {
		Name         string
		Address      string
		EpochsAtRisk string
	}

	// GasPrice represents the different reference gas prices used on the network.
	GasPrice struct {
		MinReferenceGasPrice               int // The minimum gas price (in wei) that transactions should pay in order to be included in the next block.
		MaxReferenceGasPrice               int // The maximum gas price (in wei) that transactions should pay in order to avoid overpaying and wasting funds.
		MeanReferenceGasPrice              int // The average gas price (in wei) of transactions that were included in the last few blocks.
		StakeWeightedMeanReferenceGasPrice int // The average gas price (in wei) weighted by the amount of stake that each validator has on the network.
		MedianReferenceGasPrice            int // The middle value of the sorted list of gas prices (in wei) that were included in the last few blocks.
		EstimatedNextReferenceGasPrice     int // The gas price (in wei) that is estimated to be included in the next block based on recent network activity and congestion.
	}

	// Metrics represents various metrics about the Sui blockchain network.
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
		GasPrice
		Errors
	}
)

// NewMetrics initializes a new instance of Metrics with default values.
func NewMetrics() *Metrics {
	return &Metrics{
		Updated:     false,
		SystemState: SuiSystemState{},
		Uptime:      "",
		Version:     "",
		Commit:      "",
		Transactions: Transactions{
			TotalTransactionsBlocks:      0,
			TotalTransactionCertificates: 0,
			TotalTransactionEffects:      0,
			TransactionsPerSecond:        0.0,
			TxSyncPercentage:             0,
		},
		Checkpoints: Checkpoints{
			LatestCheckpoint:        0,
			HighestKnownCheckpoint:  0,
			HighestSyncedCheckpoint: 0,
			LastExecutedCheckpoint:  0,
			CheckpointExecBacklog:   0,
			CheckpointSyncBacklog:   0,
			CheckpointsPerSecond:    0.0,
			CheckSyncPercentage:     0,
		},
		Rounds: Rounds{
			CurrentRound:          0,
			HighestProcessedRound: 0,
			LastCommittedRound:    0,
		},
		Peers: Peers{
			NetworkPeers:        0,
			PrimaryNetworkPeers: 0,
			WorkerNetworkPeers:  0,
		},
		Epoch: Epoch{
			CurrentEpoch:       0,
			EpochTotalDuration: 0,
			TimeTillNextEpoch:  0,
		},
		Errors: Errors{
			TotalSignatureErrors:         0,
			SkippedConsensusTransactions: 0,
		},
	}
}

// NewValidatorReport returns a new ValidatorReport instance.
func NewValidatorReport(reportedName, reportedAddress, reporterName, reporterAddress string) ValidatorReport {
	return ValidatorReport{
		ReportedName:    reportedName,
		ReportedAddress: reportedAddress,
		ReporterName:    reporterName,
		ReporterAddress: reporterAddress,
	}
}

// NewValidatorAtRisk returns a new ValidatorAtRisk instance.
func NewValidatorAtRisk(name, address, epochsAtRisk string) ValidatorAtRisk {
	return ValidatorAtRisk{
		Name:         name,
		Address:      address,
		EpochsAtRisk: epochsAtRisk,
	}
}
