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
func (addr *AddressInfo) GetURLRPC() (string, error) {
	endpoint := addr.Endpoint
	ports := addr.Ports
	protocol := getProtocol(endpoint.SSL)

	hostURL, err := url.Parse(fmt.Sprintf("%s://", protocol))
	if err != nil {
		return "", err
	}

	if endpoint.Host != nil {
		hostURL.Host = *endpoint.Host
	} else {
		hostURL.Host = *endpoint.IP
	}

	if rpcPort, ok := ports[enums.PortTypeRPC]; ok {
		hostURL.Host = net.JoinHostPort(hostURL.Host, rpcPort)
	} else if endpoint.IP != nil {
		hostURL.Host = net.JoinHostPort(*endpoint.IP, rpcPortDefault)
	}

	if endpoint.Path != nil {
		hostURL.Path = *endpoint.Path
	}

	return hostURL.String(), nil
}

// GetUrlPrometheus generates a URL for the Prometheus endpoint of the address.
// It constructs the URL using the protocol, host, port, and path
// components of the endpoint, as well as the default port and path values.
func (addr *AddressInfo) GetURLPrometheus() (string, error) {
	endpoint := addr.Endpoint
	ports := addr.Ports
	protocol := getProtocol(endpoint.SSL)

	hostURL, err := url.Parse(fmt.Sprintf("%s://", protocol))
	if err != nil {
		return "", err
	}

	if endpoint.Host != nil {
		hostURL.Host = *endpoint.Host
	} else {
		hostURL.Host = *endpoint.IP
	}

	if metricsPort, ok := ports[enums.PortTypeMetrics]; ok {
		hostURL.Host = net.JoinHostPort(hostURL.Host, metricsPort)
	} else if endpoint.IP != nil {
		hostURL.Host = net.JoinHostPort(*endpoint.IP, metricsPortDefault)
	}

	hostURL.Path = metricsPathDefault

	return hostURL.String(), nil
}

// getProtocol returns the protocol based on the secure flag.
func getProtocol(secure bool) string {
	protocol := protocolHTTP
	if secure {
		protocol = protocolHTTPS
	}

	return protocol
}
