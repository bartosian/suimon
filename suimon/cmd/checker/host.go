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

func (host *Host) SetStatus() {
	metrics := host.Metrics

	if !metrics.Updated {
		host.Status = enums.StatusRed
	} else if metrics.TotalTransactionNumber == "" || metrics.HighestSyncedCheckpoint == "" {
		host.Status = enums.StatusYellow
	} else {
		host.Status = enums.StatusGreen
	}
}

type RPCList struct {
	Testnet []string `yaml:"testnet"`
	Devnet  []string `yaml:"devnet"`
}

func (rpc RPCList) GetByNetwork(network enums.NetworkType) []string {
	switch network {
	case enums.NetworkTypeTestnet:
		return rpc.Testnet
	case enums.NetworkTypeDevnet:
		fallthrough
	default:
		return rpc.Devnet
	}
}
