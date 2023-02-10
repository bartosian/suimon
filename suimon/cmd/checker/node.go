package checker

import (
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/ipinfo/go/v2/ipinfo"
	"net/http"
	"time"

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
	ipClient *ipinfo.Client,
	httpClient *http.Client,
	address, rpcPort, metricsPort string,
) *Node {
	node := &Node{
		Peer: Peer{
			Address:  address,
			Port:     rpcPort,
			ipClient: ipClient,
			logger:   log.NewLogger(),
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
	} else if metrics.TotalTransactionNumber == "" || metrics.LatestCheckpoint == "" {
		time.Sleep(3 * time.Second)

		node.Status = enums.StatusYellow
	} else {
		node.Status = enums.StatusGreen
	}
}
