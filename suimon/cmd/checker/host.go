package checker

import (
	"regexp"
	"time"

	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/pkg/address"
	"github.com/bartosian/sui_helpers/suimon/pkg/log"
)

type requestType int

const (
	rpcPortDefault      = "9000"
	metricsPortDefault  = "9184"
	rpcClientTimeout    = 3 * time.Second
	metricVersionRegexp = `\{(.*?)\}`

	requestTypeRPC requestType = iota
	requestTypeMetrics
)

var versionRegex = regexp.MustCompile(metricVersionRegexp)

type AddressInfo struct {
	HostPort address.HostPort
	Ports    map[enums.PortType]string
}

type Host struct {
	AddressInfo

	Status   enums.Status
	Location *Location
	Metrics  Metrics

	rpcHttpClient  jsonrpc.RPCClient
	rpcHttpsClient jsonrpc.RPCClient
	ipClient       *ipinfo.Client

	logger log.Logger
}

func newHost(addressInfo AddressInfo, ipClient *ipinfo.Client) *Host {
	host := &Host{
		AddressInfo: addressInfo,
		ipClient:    ipClient,
		logger:      log.NewLogger(),
	}

	host.rpcHttpClient = jsonrpc.NewClient(host.getUrl(requestTypeRPC, false))
	host.rpcHttpsClient = jsonrpc.NewClient(host.getUrl(requestTypeRPC, true))

	return host
}

func (host *Host) SetStatus(tableType enums.TableType, rpc Host) {
	status := enums.StatusGreen
	metricsHost := host.Metrics
	metricsRPC := rpc.Metrics

	switch metricsHost.Updated {
	case false:
		status = enums.StatusRed

		if tableType == enums.TableTypePeers {
			status = enums.StatusYellow
		}
	case true:
		if metricsHost.TotalTransactionNumber == "" {
			status = enums.StatusRed

			if tableType == enums.TableTypePeers {
				status = enums.StatusYellow
			}
		} else if metricsHost.IsUnhealthy(enums.MetricTypeTotalTransactionsNumber, metricsRPC.TotalTransactionNumber) ||
			metricsHost.IsUnhealthy(enums.MetricTypeLatestCheckpoint, metricsRPC.LatestCheckpoint) {
			status = enums.StatusYellow
		}
	}

	host.Status = status
}
