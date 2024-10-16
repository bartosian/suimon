package metrics

const (
	TransactionsPerSecondWindow     = 5
	CheckpointsPerSecondWindow      = 5
	RoundsPerSecondWindow           = 5
	CertificatesPerSecondWindow     = 5
	TransactionsPerSecondLag        = 5
	CheckpointsPerSecondLag         = 10
	LatestCheckpointLag             = 30
	HighestSyncedCheckpointLag      = 30
	TotalTransactionsSyncPercentage = 99
	TotalCheckpointsSyncPercentage  = 99
)

type (
	// Transactions represents information about transactions on the Sui blockchain network.
	Transactions struct {
		TransactionsHistory                 []int
		CertificatesHistory                 []int
		TotalTransactionsBlocks             int
		TotalTransactionCertificates        int
		TotalTransactionCertificatesCreated int
		CertificatesPerSecond               int
		NonConsensusLatency                 int
		TotalTransactionEffects             int
		TransactionsPerSecond               int
		TxSyncPercentage                    int
	}

	// Checkpoints represents information about checkpoints on the Sui blockchain network.
	Checkpoints struct {
		CheckpointsHistory      []int
		LatestCheckpoint        int
		HighestKnownCheckpoint  int
		HighestSyncedCheckpoint int
		LastExecutedCheckpoint  int
		CheckpointsPerSecond    int
		CheckpointExecBacklog   int
		CheckpointSyncBacklog   int
		CheckSyncPercentage     int
	}

	// Rounds represents information about rounds on the Sui blockchain network.
	Rounds struct {
		RoundsHistory                        []int
		LastCommittedLeaderRound             int
		HighestAcceptedRound                 int
		RoundsPerSecond                      int
		ConsensusRoundProberCurrentRoundGaps int
	}

	// Peers represents information about peers on the Sui blockchain network.
	Peers struct {
		NetworkPeers int
	}

	// Epoch represents information about the current epoch on the Sui blockchain network.
	Epoch struct {
		EpochStartTimeUTC        string
		EpochDurationHHMM        string
		DurationTillEpochEndHHMM string
		CurrentEpoch             int
		EpochTotalDuration       int
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

	// Object represents information about objects on the Sui blockchain network.
	Objects struct {
		NumberSharedObjectTransactions int
	}

	// Metrics represents various metrics about the Sui blockchain network.
	Metrics struct {
		ValidatorsApyParsed ValidatorsApyParsed

		Uptime  string
		Version string
		Commit  string

		SystemState SuiSystemState

		CurrentVotingRight float64

		Epoch
		Protocol
		Rounds
		Objects
		Transactions
		Checkpoints
		GasPrice
		Peers
		Errors
		Updated bool
	}
)
