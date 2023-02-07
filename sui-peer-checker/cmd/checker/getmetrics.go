package checker

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/enums"
)

const (
	metricSeparator     = " "
	metricComment       = "#"
	metricKeySeparator  = "="
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
	metrics := make(map[enums.MetricType]string)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
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
			metrics[enums.MetricTypeVersion] = version[1]
		}

		metrics[metricName] = value
	}

	if len(metrics) == 0 {
		return
	}

	metrics[enums.MetricTypeTotalTransactionsNumber] = peer.Metrics.TotalTransactionNumber
	peer.Metrics = NewMetrics(metrics)

	return
}
