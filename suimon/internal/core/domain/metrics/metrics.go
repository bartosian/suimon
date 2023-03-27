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

type Metrics struct {
	Updated bool

	SystemState SuiSystemState

	TotalTransactions            int
	TotalTransactionCertificates int
	TotalTransactionEffects      int
	TransactionsPerSecond        int
	TransactionsHistory          []int
	LatestCheckpoint             int
	HighestKnownCheckpoint       int
	HighestSyncedCheckpoint      int
	LastExecutedCheckpoint       int
	CheckpointsPerSecond         int
	CheckpointExecBacklog        int
	CheckpointSyncBacklog        int
	CheckpointsHistory           []int
	CurrentEpoch                 int
	EpochTotalDuration           int
	EpochPercentage              int
	TimeTillNextEpochMs          int
	TxSyncPercentage             int
	CheckSyncPercentage          int
	NetworkPeers                 int
	Uptime                       string
	Version                      string
	Commit                       string
	CurrentRound                 int
	HighestProcessedRound        int
	LastCommittedRound           int
	PrimaryNetworkPeers          int
	WorkerNetworkPeers           int
	SkippedConsensusTransactions int
	TotalSignatureErrors         int
}

// NewMetrics creates and returns a new instance of the Metrics struct.
// Returns: a new instance of the Metrics struct.
func NewMetrics() Metrics {
	return Metrics{
		TransactionsHistory: make([]int, 0, TransactionsPerSecondWindow),
		CheckpointsHistory:  make([]int, 0, CheckpointsPerSecondWindow),
	}
}
