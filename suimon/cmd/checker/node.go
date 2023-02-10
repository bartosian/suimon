package checker

import (
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"net/http"

	"github.com/oschwald/geoip2-golang"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/pkg/log"
)

type Node struct {
	Peer

	Status      enums.Status
	RpcPort     string
	MetricsPort string
}

func newNode(
	geoDB *geoip2.Reader,
	httpClient *http.Client,
	address, rpcPort, metricsPort string,
) *Node {
	node := &Node{
		Peer: Peer{
			Address:     address,
			Port:        rpcPort,
			geoDbClient: geoDB,
			logger:      log.NewLogger(),
		},
		RpcPort:     rpcPort,
		MetricsPort: metricsPort,
	}

	rpcURL := node.getUrl(requestTypeRPC, false, &rpcPort)

	node.rpcClient = jsonrpc.NewClient(rpcURL)
	node.httpClient = httpClient

	return node
}

func (node *Node) SetStatus() {
	metrics := node.Metrics

	if !metrics.Updated {
		node.Status = enums.StatusRed
	} else if metrics.TotalTransactionNumber != "" && metrics.LatestCheckpoint != "" {
		node.Status = enums.StatusGreen
	} else {
		node.Status = enums.StatusYellow
	}
}
