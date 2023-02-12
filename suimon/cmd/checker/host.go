package checker

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
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

func (host *Host) SetLocation() {
	var parseLocation = func(ip string) {
		record, err := host.ipClient.GetIPInfo(net.IP(ip))
		if err != nil {
			return
		}

		countryISOCode := record.Country
		countryName := record.CountryName
		flag := record.CountryFlag.Emoji
		company := record.Company.Name

		host.Location = newLocation(countryISOCode, countryName, flag, company)
	}

	if host.HostPort.IP != nil {
		parseLocation(*host.HostPort.IP)
	}
}

func (host *Host) getUrl(request requestType, secure bool) string {
	var (
		address  string
		port     string
		protocol = "http"
	)

	if secure {
		protocol = protocol + "s"
	}

	hostPort := host.HostPort

	if hostPort.IP != nil {
		address = *hostPort.IP
	} else if hostPort.Host != nil {
		address = *hostPort.GetHostWithPath()

		if hostPort.Path != nil {
			return fmt.Sprintf("%s://%s", protocol, address)
		}
	}

	switch request {
	case requestTypeRPC:
		port = rpcPortDefault
		if portRPC, ok := host.Ports[enums.PortTypeRPC]; ok {
			port = portRPC
		}

		return fmt.Sprintf("%s://%s:%s", protocol, address, port)
	case requestTypeMetrics:
		fallthrough
	default:
		port = metricsPortDefault
		if portMetrics, ok := host.Ports[enums.PortTypeMetrics]; ok {
			port = portMetrics
		}

		return fmt.Sprintf("%s://%s:%s/metrics", protocol, address, port)
	}
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

func (host *Host) GetTotalTransactionNumber() {
	if result := getFromRPC(host.rpcHttpClient, enums.RPCMethodGetTotalTransactionNumber); result != nil {
		totalTransactionNumber := strconv.Itoa(*result)

		host.Metrics.SetValue(enums.MetricTypeTotalTransactionsNumber, totalTransactionNumber)

		return
	}

	if result := getFromRPC(host.rpcHttpsClient, enums.RPCMethodGetTotalTransactionNumber); result != nil {
		totalTransactionNumber := strconv.Itoa(*result)

		host.Metrics.SetValue(enums.MetricTypeTotalTransactionsNumber, totalTransactionNumber)
	}
}

func (host *Host) GetLatestCheckpoint() {
	if result := getFromRPC(host.rpcHttpClient, enums.RPCMethodGetLatestCheckpointSequenceNumber); result != nil {
		latestCheckpoint := strconv.Itoa(*result)

		host.Metrics.SetValue(enums.MetricTypeLatestCheckpoint, latestCheckpoint)

		return
	}

	if result := getFromRPC(host.rpcHttpsClient, enums.RPCMethodGetLatestCheckpointSequenceNumber); result != nil {
		latestCheckpoint := strconv.Itoa(*result)

		host.Metrics.SetValue(enums.MetricTypeLatestCheckpoint, latestCheckpoint)
	}
}

func getFromRPC(rpcClient jsonrpc.RPCClient, method enums.RPCMethod) *int {
	respChan := make(chan *int)
	timeout := time.After(rpcClientTimeout)

	go func() {
		var response *int

		if err := rpcClient.CallFor(context.Background(), &response, method.String()); err != nil {
			return
		}

		respChan <- response
	}()

	select {
	case response := <-respChan:
		return response
	case <-timeout:
		return nil
	}
}

func (host *Host) GetMetrics(httpClient *http.Client) {
	metricsURL := host.getUrl(requestTypeMetrics, false)

	result, err := httpClient.Get(metricsURL)
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

		if strings.HasPrefix(line, "#") {
			continue
		}

		metric := strings.Split(line, " ")
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
			version := strings.Split(versionMetric[1], "=")

			uptimeSeconds, err := strconv.Atoi(value)
			if err != nil {
				continue
			}

			value = fmt.Sprintf("%.1f days", float64(uptimeSeconds)/(60*60*24))

			versionInfo := strings.Split(version[1], "-")
			if len(versionInfo) != 2 {
				continue
			}

			host.Metrics.SetValue(enums.MetricTypeVersion, versionInfo[0])
			host.Metrics.SetValue(enums.MetricTypeCommit, versionInfo[1])
		}

		host.Metrics.SetValue(metricName, value)
	}
}
