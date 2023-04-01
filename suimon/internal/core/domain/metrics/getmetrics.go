package metrics

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/dariubs/percent"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/utility"
)

const (
	bytesInTB   = 1024 * 1024 * 1024 * 1024
	bytesInGB   = 1024 * 1024 * 1024
	bytesInByte = 1024

	suiDBDirName = "suidb"
)

type (
	MetricValue interface{}
)

// GetValue returns the current value of the specified metric.
// Parameters:
// - metric: an enums.MetricType representing the metric whose value to retrieve.
// - rpc: a boolean indicating whether to retrieve the value for RPC host.
// Returns: the current value of the specified metric, which can be of any type.
func (metrics *Metrics) GetValue(metric enums.MetricType, rpc bool) MetricValue {
	switch metric {
	case enums.MetricTypeUptime:
		return metrics.Uptime
	case enums.MetricTypeVersion:
		return metrics.Version
	case enums.MetricTypeCommit:
		return metrics.Commit
	case enums.MetricTypeHighestSyncedCheckpoint:
		return metrics.HighestSyncedCheckpoint
	case enums.MetricTypeSuiNetworkPeers:
		return metrics.NetworkPeers
	case enums.MetricTypeTotalTransactionBlocks:
		return metrics.TotalTransactionsBlocks
	case enums.MetricTypeTransactionsPerSecond:
		return metrics.TransactionsPerSecond
	case enums.MetricTypeCheckpointsPerSecond:
		return metrics.CheckpointsPerSecond
	case enums.MetricTypeTxSyncPercentage:
		return metrics.TotalTransactionsBlocks
	case enums.MetricTypeCheckSyncPercentage:
		if rpc {
			return metrics.LatestCheckpoint
		}

		return metrics.HighestSyncedCheckpoint
	case enums.MetricTypeLatestCheckpoint:
		return metrics.LatestCheckpoint
	default:
		return nil
	}
}

// GetMillisecondsTillNextEpoch returns the time remaining until the next epoch in milliseconds.
// Returns: an integer representing the time remaining until the next epoch in seconds.
func (metrics *Metrics) GetMillisecondsTillNextEpoch() int {
	nextEpochStartMs := metrics.SystemState.EpochStartTimestampMs + metrics.SystemState.EpochDurationMs
	currentTimeMs := int(time.Now().UnixNano() / 1000000)

	return nextEpochStartMs - currentTimeMs
}

// GetTimeUntilNextEpochDisplay returns an array of strings representing the epoch timer data.
// Returns: an array of strings representing the epoch timer data.
func (metrics *Metrics) GetTimeUntilNextEpochDisplay() []string {
	duration := time.Duration(metrics.TimeTillNextEpochMs) * time.Millisecond
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) - (hours * 60)
	second := time.Now().Second()

	if hours < 0 {
		return nil
	}

	spacer := " "
	if second%2 == 0 {
		spacer = ":"
	}

	return []string{fmt.Sprintf("%02d%s%02d", hours, spacer, minutes), "H"}
}

// GetEpochLabel returns a string representing the epoch label.
// Returns: a string representing the epoch label.
func (metrics *Metrics) GetEpochLabel() string {
	return fmt.Sprintf("EPOCH %d", metrics.SystemState.Epoch)
}

// GetEpochProgress returns an integer representing the current epoch progress.
// Returns: an integer representing the current epoch progress.
func (metrics *Metrics) GetEpochProgress() int {
	epochCurrentLength := metrics.SystemState.EpochDurationMs - metrics.TimeTillNextEpochMs

	return int(percent.PercentOf(epochCurrentLength, metrics.SystemState.EpochDurationMs))
}

// GetUsageDataForDonutChart retrieves the usage data for a specific unit and returns it as a formatted string and a percentage value for display in a donut chart.
// Parameters:
// - unit: a string representing the unit for which to retrieve usage data.
// - option: a function that returns a pointer to a utility.UsageData object and an error, used to retrieve the actual usage data.
func GetUsageDataForDonutChart(unit enums.MetricUnit, option func() (*utility.UsageData, error)) (string, int) {
	var (
		usageLabel      = "LOADING..."
		usagePercentage = 1
	)

	if usageDataResult, err := option(); err == nil {
		usageLabel = fmt.Sprintf("TOTAL/USED: %d/%d%s", usageDataResult.Total, usageDataResult.Used, unit)
		usagePercentage = usageDataResult.PercentageUsed

		if usagePercentage == 0 {
			usagePercentage = 1
		}
	}

	return usageLabel, usagePercentage
}

type NetworkUsageResult struct {
	UsageData string
	Unit      enums.MetricUnit
}

// GetFormattedNetworkUsage retrieves the usage data for a specific network metric and returns it as an array of formatted strings for display in a table.
// Parameters: networkMetric: a dashboards.CellName representing the name of the network metric for which to retrieve usage data.
// Returns: an array of strings representing the formatted usage data for display in a table.
func GetFormattedNetworkUsage(networkMetric enums.CellName) NetworkUsageResult {
	var (
		usageData    = ""
		networkUsage *utility.NetworkUsage
		unit         enums.MetricUnit
		err          error
	)

	if networkUsage, err = utility.GetNetworkUsage(); err == nil {
		var metric float64
		if networkMetric == enums.CellNameBytesReceived {
			metric = networkUsage.Recv
		} else {
			metric = networkUsage.Sent
		}

		if metric >= bytesInTB {
			metric = metric / bytesInTB
			unit = enums.MetricUnitTB
		} else {
			metric = metric / bytesInGB
			unit = enums.MetricUnitGB
		}

		formatString := "%.02f"
		if metric >= 100 {
			formatString = "%.01f"
		}

		usageData = fmt.Sprintf(formatString, metric)
	}

	return NetworkUsageResult{UsageData: usageData, Unit: unit}
}

type FileSizeResult struct {
	Size float64
	Unit enums.MetricUnit
}

// GetFileSize retrieves the sizes of all files and directories within a given directory and returns them as an array of formatted strings for display in a table.
// Parameters:
// - dirPath: a string representing the path of the directory for which to retrieve sizes.
// Returns:
// - an array of strings representing the formatted size data for all files and directories within the given directory.
func GetFileSize(filePath string) FileSizeResult {
	var (
		usageData string
		fileSize  float64
		unit      enums.MetricUnit
		err       error
	)

	if filePath == suiDBDirName {
		var homeDir string

		if homeDir, err = os.UserHomeDir(); err != nil {
			return FileSizeResult{}
		}

		filePath = filepath.Join(homeDir, "sui", suiDBDirName)
	}

	if fileSize, err = utility.GetDirSize(filePath); err == nil {
		unit = enums.MetricUnitB
	} else if fileSize, err = utility.GetVolumeSize(filePath); err == nil {
		unit = enums.MetricUnitGB
		fileSize = fileSize / bytesInGB
	} else {
		return FileSizeResult{}
	}

	if fileSize >= 100 {
		fileSize = fileSize / bytesInByte
		unit = enums.MetricUnitTB
	}

	formatString := "%.02f"
	if fileSize >= 100 {
		formatString = "%.01f"
	}

	usageData = fmt.Sprintf(formatString, fileSize)

	if usageData == "0.00" {
		usageData = ""
	}

	return FileSizeResult{
		Size: fileSize,
		Unit: unit,
	}
}
