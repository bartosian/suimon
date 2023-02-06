package checker

import (
	"bufio"
	"io"
	"net/http"
	"strings"

	"github.com/bartosian/sui_helpers/peer_checker/domain/enums"
)

const (
	metricSeparator = " "
	metricComment   = "#"
)

func (peer *Peer) GetMetrics() {
	result, err := http.Get(peer.getUrl(requestTypeMetrics, false))
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
		parts := strings.SplitN(line, metricSeparator, 2)
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		parsedResp[key] = value
	}

	return
}
