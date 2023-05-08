package metrics

type (
	// EpochsHistory represents a list of epoch data returned by an API.
	EpochsHistory struct {
		Data []EpochInfo `json:"data"`
	}

	// EpochInfo represents information about the epoch.
	EpochInfo struct {
		Epoch                  string          `json:"epoch"`
		Validators             []interface{}   `json:"validators"`
		EpochTotalTransactions string          `json:"epochTotalTransactions"`
		FirstCheckpointID      string          `json:"firstCheckpointId"`
		EpochStartTimestamp    string          `json:"epochStartTimestamp"`
		EndOfEpochInfo         *EndOfEpochInfo `json:"endOfEpochInfo"`
	}

	// EndOfEpochInfo represents information about the end of epoch.
	EndOfEpochInfo struct {
		LastCheckpointID             string `json:"lastCheckpointId"`
		EpochEndTimestamp            string `json:"epochEndTimestamp"`
		ProtocolVersion              string `json:"protocolVersion"`
		ReferenceGasPrice            string `json:"referenceGasPrice"`
		TotalStake                   string `json:"totalStake"`
		StorageFundReinvestment      string `json:"storageFundReinvestment"`
		StorageCharge                string `json:"storageCharge"`
		StorageRebate                string `json:"storageRebate"`
		StorageFundBalance           string `json:"storageFundBalance"`
		StakeSubsidyAmount           string `json:"stakeSubsidyAmount"`
		TotalGasFees                 string `json:"totalGasFees"`
		TotalStakeRewardsDistributed string `json:"totalStakeRewardsDistributed"`
		LeftoverStorageFundInflow    string `json:"leftoverStorageFundInflow"`
	}
)
