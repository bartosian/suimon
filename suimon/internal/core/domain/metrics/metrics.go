package metrics

const (
	TransactionsPerSecondWindow     = 5
	CheckpointsPerSecondWindow      = 5
	RoundsPerSecondWindow           = 5
	CertificatesPerSecondWindow     = 5
	TransactionsPerSecondLag        = 5
	TotalTransactionsLag            = 100
	CheckpointsPerSecondLag         = 10
	LatestCheckpointLag             = 30
	HighestSyncedCheckpointLag      = 30
	TotalTransactionsSyncPercentage = 99
	TotalCheckpointsSyncPercentage  = 99
)

type (
	// Transactions represents information about transactions on the Sui blockchain network.
	Transactions struct {
		TotalTransactionsBlocks      int
		TotalTransactionCertificates int
		CertificatesCreated          int
		CertificatesPerSecond        int
		NonConsensusLatency          int
		TotalTransactionEffects      int
		TransactionsPerSecond        int
		TxSyncPercentage             int
		TransactionsHistory          []int
		CertificatesHistory          []int
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
		RoundsPerSecond       int
		LastCommittedRound    int
		RoundsHistory         []int
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
