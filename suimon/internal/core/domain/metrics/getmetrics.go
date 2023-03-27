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

// GetValue returns the current value of the specified metric.
// Parameters:
// - metric: an enums.MetricType representing the metric whose value to retrieve.
// - rpc: a boolean indicating whether to retrieve the value for RPC host.
// Returns: the current value of the specified metric, which can be of any type.
func (metrics *Metrics) GetValue(metric enums.MetricType, rpc bool) any {
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
	case enums.MetricTypeTotalTransactions:
		return metrics.TotalTransactions
	case enums.MetricTypeTransactionsPerSecond:
		return metrics.TransactionsPerSecond
	case enums.MetricTypeCheckpointsPerSecond:
		return metrics.CheckpointsPerSecond
	case enums.MetricTypeTxSyncPercentage:
		return metrics.TotalTransactions
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

// GetTimeTillNextEpoch returns the time remaining until the next epoch in milliseconds.
// Returns: an integer representing the time remaining until the next epoch in seconds.
func (metrics *Metrics) GetTimeTillNextEpoch() int {
	nextEpochStartMs := metrics.SystemState.EpochStartTimestampMs + metrics.SystemState.EpochDurationMs
	currentTimeMs := int(time.Now().UnixNano() / 1000000)

	return nextEpochStartMs - currentTimeMs
}

// GetEpochTimer returns an array of strings representing the epoch timer data.
// Returns: an array of strings representing the epoch timer data.
func (metrics *Metrics) GetEpochTimer() []string {
	duration := time.Duration(metrics.TimeTillNextEpochMs) * time.Millisecond
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) - (hours * 60)
	second := time.Now().Second()

	if hours < 0 {
		return []string{""}
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

// GetDonutUsageMetric retrieves the usage data for a specific unit and returns it as a formatted string and a percentage value for display in a donut chart.
// Parameters:
// - unit: a string representing the unit for which to retrieve usage data.
// - option: a function that returns a pointer to a utility.UsageData object and an error, used to retrieve the actual usage data.
// Returns:
// - a string representing the formatted usage data for display in a donut chart.
// - an integer representing the percentage value of the usage data for display in a donut chart.
func GetDonutUsageMetric(unit string, option func() (*utility.UsageData, error)) (string, int) {
	var (
		usageLabel      = "LOADING..."
		usagePercentage = 1
		usageData       *utility.UsageData
		err             error
	)

	if usageData, err = option(); err == nil {
		usageLabel = fmt.Sprintf("TOTAL/USED: %d/%d%s", usageData.Total, usageData.Used, unit)
		usagePercentage = usageData.PercentageUsed

		if usagePercentage == 0 {
			usagePercentage = 1
		}
	}

	return usageLabel, usagePercentage
}

// GetNetworkUsageMetric retrieves the usage data for a specific network metric and returns it as an array of formatted strings for display in a table.
// Parameters: networkMetric: a dashboards.CellName representing the name of the network metric for which to retrieve usage data.
// Returns: an array of strings representing the formatted usage data for display in a table.
func GetNetworkUsageMetric(networkMetric enums.CellName) []string {
	var (
		usageData    = ""
		networkUsage *utility.NetworkUsage
		unit         string
		err          error
	)

	if networkUsage, err = utility.GetNetworkUsage(); err == nil {
		metric := networkUsage.Sent
		formatString := "%.02f"
		unit = "GB"

		if networkMetric == enums.CellNameBytesReceived {
			metric = networkUsage.Recv
		}

		if metric >= 100 {
			metric = metric / 1024
			unit = "TB"

			if metric >= 100 {
				formatString = "%.01f"
			}
		}

		usageData = fmt.Sprintf(formatString, metric)
	}

	return []string{usageData, unit}
}

// GetDirectorySize retrieves the sizes of all files and directories within a given directory and returns them as an array of formatted strings for display in a table.
// Parameters:
// - dirPath: a string representing the path of the directory for which to retrieve sizes.
// Returns:
// - an array of strings representing the formatted size data for all files and directories within the given directory.
func GetDirectorySize(dirPath string) []string {
	var (
		usageData = ""
		dirSize   float64
		unit      string
		err       error
	)

	if dirPath == "suidb" {
		var homeDir string

		if homeDir, err = os.UserHomeDir(); err != nil {
			return nil
		}

		dirPath = filepath.Join(homeDir, "sui", dirPath)
	}

	var processSize = func() {
		formatString := "%.02f"
		unit = "GB"

		if dirSize >= 100 {
			dirSize = dirSize / 1024
			unit = "TB"

			if dirSize >= 100 {
				formatString = "%.01f"
			}
		}

		if usageData = fmt.Sprintf(formatString, dirSize); usageData == "0.00" {
			usageData = ""
		}
	}

	if dirSize, err = utility.GetDirSize(dirPath); err == nil {
		processSize()

		return []string{usageData, unit}
	}

	if dirSize, err = utility.GetVolumeSize("suidb"); err == nil {
		processSize()
	}

	return []string{usageData, unit}
}
