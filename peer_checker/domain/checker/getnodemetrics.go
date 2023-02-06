package checker

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/bartosian/sui_helpers/peer_checker/domain/enums"
)

const (
	metricSeparator     = " "
	metricComment       = "#"
	metricKeySeparator  = "="
	metricVersionRegexp = `\{(.*?)\}`
)

var versionRegex = regexp.MustCompile(metricVersionRegexp)

func (peer *Peer) GetMetrics() {
	result, err := peer.httpClient.Get(peer.getUrl(requestTypeMetrics, false))
	if err != nil {
		return
	}

	defer result.Body.Close()

	reader := bufio.NewReader(result.Body)
	parsedResp := make(map[enums.MetricName]string)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if strings.HasPrefix(line, metricComment) {
			continue
		}

		// Split the line into a key and value
		parts := strings.Split(line, metricSeparator)
		if len(parts) != 2 {
			continue
		}

		key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

		metricName, err := enums.MetricNameFromString(key)
		if err != nil {
			continue
		}

		if metricName == enums.MetricNameUptime {
			version := versionRegex.FindStringSubmatch(key)
			versionInfo := strings.Split(version[1], metricKeySeparator)

			uptimeSeconds, err := strconv.Atoi(value)
			if err != nil {
				continue
			}

			value = fmt.Sprintf("%.1f days", float64(uptimeSeconds)/(60*60*24))
			parsedResp[enums.MetricNameVersion] = versionInfo[1]
		}

		parsedResp[metricName] = value
	}

	if len(parsedResp) == 0 {
		return
	}

	peer.Metrics = NewMetrics(parsedResp)

	return
}
