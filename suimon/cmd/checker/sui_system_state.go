package checker

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
		NextEpochStake               int64       `json:"nextEpochStake"`
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
		Epoch                          int           `json:"epoch"`
		ProtocolVersion                int           `json:"protocolVersion"`
		SystemStateVersion             int           `json:"systemStateVersion"`
		StorageFund                    int           `json:"storageFund"`
		ReferenceGasPrice              int           `json:"referenceGasPrice"`
		SafeMode                       bool          `json:"safeMode"`
		EpochStartTimestampMs          int           `json:"epochStartTimestampMs"`
		GovernanceStartEpoch           int           `json:"governanceStartEpoch"`
		EpochDurationMs                int           `json:"epochDurationMs"`
		StakeSubsidyEpochCounter       int           `json:"stakeSubsidyEpochCounter"`
		StakeSubsidyBalance            int64         `json:"stakeSubsidyBalance"`
		StakeSubsidyCurrentEpochAmount int64         `json:"stakeSubsidyCurrentEpochAmount"`
		TotalStake                     int64         `json:"totalStake"`
		ActiveValidators               []Validator   `json:"activeValidators"`
		PendingActiveValidatorsID      string        `json:"pendingActiveValidatorsId"`
		PendingActiveValidatorsSize    int           `json:"pendingActiveValidatorsSize"`
		PendingRemovals                []interface{} `json:"pendingRemovals"`
		StakingPoolMappingsID          string        `json:"stakingPoolMappingsId"`
		StakingPoolMappingsSize        int           `json:"stakingPoolMappingsSize"`
		InactivePoolsID                string        `json:"inactivePoolsId"`
		InactivePoolsSize              int           `json:"inactivePoolsSize"`
		ValidatorCandidatesID          string        `json:"validatorCandidatesId"`
		ValidatorCandidatesSize        int           `json:"validatorCandidatesSize"`
		AtRiskValidators               []interface{} `json:"atRiskValidators"`
		ValidatorReportRecords         []interface{} `json:"validatorReportRecords"`
	}
)
