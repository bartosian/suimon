package checker

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder/dashboards"
	"github.com/bartosian/sui_helpers/suimon/pkg/utility"
)

// getDonutUsageMetric retrieves the usage data for a specific unit and returns it as a formatted string and a percentage value for display in a donut chart.
// Parameters:
// - unit: a string representing the unit for which to retrieve usage data.
// - option: a function that returns a pointer to a utility.UsageData object and an error, used to retrieve the actual usage data.
// Returns:
// - a string representing the formatted usage data for display in a donut chart.
// - an integer representing the percentage value of the usage data for display in a donut chart.
func getDonutUsageMetric(unit string, option func() (*utility.UsageData, error)) (string, int) {
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

// getNetworkUsageMetric retrieves the usage data for a specific network metric and returns it as an array of formatted strings for display in a table.
// Parameters: networkMetric: a dashboards.CellName representing the name of the network metric for which to retrieve usage data.
// Returns: an array of strings representing the formatted usage data for display in a table.
func getNetworkUsageMetric(networkMetric dashboards.CellName) []string {
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

		if networkMetric == dashboards.CellNameBytesReceived {
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

// getDirectorySize retrieves the sizes of all files and directories within a given directory and returns them as an array of formatted strings for display in a table.
// Parameters:
// - dirPath: a string representing the path of the directory for which to retrieve sizes.
// Returns:
// - an array of strings representing the formatted size data for all files and directories within the given directory.
func getDirectorySize(dirPath string) []string {
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
