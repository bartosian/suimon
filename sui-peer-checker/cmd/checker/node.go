package checker

import (
	"net/http"

	"github.com/oschwald/geoip2-golang"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/sui-peer-checker/pkg/log"
)

type Node struct {
	Peer

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
