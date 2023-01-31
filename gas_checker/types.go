package main

type SuiSystemState struct {
	Epoch                  uint64           `json:"epoch"`
	EpochStartTimestampMs  uint64           `json:"epoch_start_timestamp_ms"`
	Info                   string           `json:"uid"`
	Parameters             SystemParameters `json:"parameters"`
	ReferenceGasPrice      uint64           `json:"reference_gas_price"`
	SafeMode               bool             `json:"safe_mode"`
	StakeSubsidy           StakeSubsidy     `json:"stake_subsidy"`
	StorageFundBalance     Balance          `json:"storage_fund"`
	TreasuryCap            Supply           `json:"treasury_cap"`
	ValidatorReportRecords interface{}      `json:"validator_report_records"`
	Validators             ValidatorSet     `json:"validators"`
}

type SystemParameters struct {
	MinValidatorStake          uint64 `json:"min_validator_stake"`
	MaxValidatorCandidateCount uint64 `json:"max_validator_candidate_count"`
	StorageGasPrice            uint64 `json:"storage_gas_price"`
}

type StakeSubsidy struct {
	EpochCounter       uint64 `json:"epoch_counter"`
	Balance            Balance
	CurrentEpochAmount uint64 `json:"current_epoch_amount"`
}

type Balance struct {
	Value uint64
}

type Supply struct {
	Value uint64
}

type ValidatorSet struct {
	ValidatorStake            uint64      `json:"validator_stake"`
	DelegationStake           uint64      `json:"delegation_stake"`
	ActiveValidators          []Validator `json:"active_validators"`
	PendingDelegationSwitches interface{} `json:"pending_delegation_switches"`
}

type Validator struct {
	Metadata              ValidatorMetadata `json:"metadata"`
	VotingPower           uint64            `json:"voting_power"`
	StakeAmount           uint64            `json:"stake_amount"`
	PendingStake          uint64            `json:"pending_stake"`
	PendingWithdraw       uint64            `json:"pending_withdraw"`
	GasPrice              uint64            `json:"gas_price"`
	DelegationStakingPool StakePool         `json:"delegation_staking_pool"`
	CommissionRate        uint64            `json:"commission_rate"`
}

type ValidatorMetadata struct {
	SuiAddress              string `json:"sui_address"`
	PubkeyBytes             []byte `json:"pubkey_bytes"`
	NetworkPubkeyBytes      []byte `json:"network_pubkey_bytes"`
	WorkerPubkeyBytes       []byte `json:"worker_pubkey_bytes"`
	ProofOfPossessionBytes  []byte `json:"proof_of_possession_bytes"`
	Name                    []byte
	Description             []byte
	ImageUrl                []byte `json:"image_url"`
	ProjectUrl              []byte `json:"project_url"`
	NetAddress              []byte `json:"net_address"`
	ConsensusAddress        []byte `json:"consensus_address"`
	WorkerAddress           []byte `json:"worker_address"`
	NextEpochStake          uint64 `json:"next_epoch_stake"`
	NextEpochDelegation     uint64 `json:"next_epoch_delegation"`
	NextEpochGasPrice       uint64 `json:"next_epoch_gas_price"`
	NextEpochCommissionRate uint64 `json:"next_epoch_comission_rate"`
}

type StakePool struct {
	ValidatorAddress      string      `json:"validator_address"`
	StartingEpoch         uint64      `json:"starting_epoch"`
	SuiBalance            uint64      `json:"sui_balance"`
	RewardsPool           Balance     `json:"rewards_pool"`
	DelegationTokenSupply Supply      `json:"delegation_token_supply"`
	PendingWithdraws      Withdrawals `json:"pending_withdraws"`
	PendingDelegations    Delegations `json:"pending_delegations"`
}

type Delegations struct {
	Id   string
	Size uint64
	Head interface{}
	Tail interface{}
}

type Withdrawals struct {
	Contents WithdrawContent
}

type WithdrawContent struct {
	Id   string
	Size uint64
}
