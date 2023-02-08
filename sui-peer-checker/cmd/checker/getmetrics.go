package checker

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/enums"
)

const (
	metricSeparator     = " "
	metricComment       = "#"
	metricKeySeparator  = "="
	metricValSeparator  = "-"
	metricVersionRegexp = `\{(.*?)\}`
)

var versionRegex = regexp.MustCompile(metricVersionRegexp)

func (peer *Peer) GetMetrics() {
	metricsURL := peer.getUrl(requestTypeMetrics, false)

	result, err := peer.httpClient.Get(metricsURL)
	if err != nil {
		return
	}

	defer result.Body.Close()

	reader := bufio.NewReader(result.Body)
	for {
		line, err := reader.ReadString('\n')
		if len(line) == 0 && err != nil {
			break
		}

		if strings.HasPrefix(line, metricComment) {
			continue
		}

		metric := strings.Split(line, metricSeparator)
		if len(metric) != 2 {
			continue
		}

		key, value := strings.TrimSpace(metric[0]), strings.TrimSpace(metric[1])

		metricName, err := enums.MetricTypeFromString(key)
		if err != nil {
			continue
		}

		if metricName == enums.MetricTypeUptime {
			versionMetric := versionRegex.FindStringSubmatch(key)
			version := strings.Split(versionMetric[1], metricKeySeparator)

			uptimeSeconds, err := strconv.Atoi(value)
			if err != nil {
				continue
			}

			value = fmt.Sprintf("%.1f days", float64(uptimeSeconds)/(60*60*24))

			versionInfo := strings.Split(version[1], metricValSeparator)
			if len(versionInfo) != 2 {
				continue
			}

			peer.Metrics.SetValue(enums.MetricTypeVersion, versionInfo[0])
			peer.Metrics.SetValue(enums.MetricTypeCommit, versionInfo[1])
		}

		peer.Metrics.SetValue(metricName, value)
	}
}
