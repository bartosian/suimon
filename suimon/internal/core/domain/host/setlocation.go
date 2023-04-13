package host

import (
	"fmt"
	"net"
)

const (
	ErrInvalidIPAddressProvided = "invalid IP address: %v"
)

// SetIPInfo sets the IPInfo property of a Host struct by calling an external geolocation API with the host's IP address.
// It returns an error if the IP address is invalid or if the API call fails.
func (host *Host) SetIPInfo() error {
	if host.Endpoint.IP == nil {
		return nil
	}

	ip := net.ParseIP(*host.Endpoint.IP)
	if ip == nil {
		return fmt.Errorf(ErrInvalidIPAddressProvided, host.Endpoint.IP)
	}

	ipInfo, err := host.gateways.geo.CallFor(ip)
	if err != nil {
		return err
	}

	host.IPInfo = ipInfo

	return nil
}
