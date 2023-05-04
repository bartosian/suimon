package host

import (
	"fmt"
	"net"
	"net/url"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/pkg/address"
)

const (
	protocolHTTP  = "http"
	protocolHTTPS = "https"

	rpcPortDefault     = "9000"
	metricsPortDefault = "9184"

	metricsPathDefault = "/metrics"
)

type AddressInfo struct {
	Endpoint address.Endpoint
	Ports    map[enums.PortType]string
}

// GetUrlRPC generates a URL for the RPC endpoint of the address.
// It constructs the URL using the protocol, host, port, and path
// components of the endpoint, as well as the default port value.
func (addr *AddressInfo) GetUrlRPC() (string, error) {
	endpoint := addr.Endpoint
	ports := addr.Ports
	protocol := getProtocol(endpoint.SSL)

	hostUrl, err := url.Parse(fmt.Sprintf("%s://", protocol))
	if err != nil {
		return "", err
	}

	if endpoint.Host != nil {
		hostUrl.Host = *endpoint.Host
	} else {
		hostUrl.Host = *endpoint.IP
	}

	if rpcPort, ok := ports[enums.PortTypeRPC]; ok {
		hostUrl.Host = net.JoinHostPort(hostUrl.Host, rpcPort)
	} else if endpoint.IP != nil {
		hostUrl.Host = net.JoinHostPort(*endpoint.IP, rpcPortDefault)
	}

	if endpoint.Path != nil {
		hostUrl.Path = *endpoint.Path
	}

	return hostUrl.String(), nil
}

// GetUrlPrometheus generates a URL for the Prometheus endpoint of the address.
// It constructs the URL using the protocol, host, port, and path
// components of the endpoint, as well as the default port and path values.
func (addr *AddressInfo) GetUrlPrometheus() (string, error) {
	endpoint := addr.Endpoint
	ports := addr.Ports
	protocol := getProtocol(endpoint.SSL)

	hostUrl, err := url.Parse(fmt.Sprintf("%s://", protocol))
	if err != nil {
		return "", err
	}

	if endpoint.Host != nil {
		hostUrl.Host = *endpoint.Host
	} else {
		hostUrl.Host = *endpoint.IP
	}

	if metricsPort, ok := ports[enums.PortTypeMetrics]; ok {
		hostUrl.Host = net.JoinHostPort(hostUrl.Host, metricsPort)
	} else if endpoint.IP != nil {
		hostUrl.Host = net.JoinHostPort(*endpoint.IP, metricsPortDefault)
	}

	hostUrl.Path = metricsPathDefault

	return hostUrl.String(), nil
}

func getProtocol(secure bool) string {
	protocol := protocolHTTP
	if secure {
		protocol = protocolHTTPS
	}

	return protocol
}
