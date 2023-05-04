package address

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	externalIP "github.com/glendc/go-external-ip"

	"github.com/bartosian/sui_helpers/suimon/pkg/validation"
)

type HostPort struct {
	Address string
	IP      *string
	Host    *string
	Path    *string
	Port    *string
}

func (hp *HostPort) GetHostWithPath() *string {
	if hp.Host == nil {
		return nil
	}

	hostPath := *hp.Host

	if hp.Path != nil {
		hostPath = hostPath + *hp.Path
	}

	return &hostPath
}

func ParseIpPort(address string) (*HostPort, error) {
	ip, port, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}

	if validation.IsInvalidPort(port) {
		return nil, fmt.Errorf("invalid port provided: %s", address)
	}

	if parsedIP := net.ParseIP(ip); parsedIP.IsLoopback() || parsedIP.IsUnspecified() {
		ip = GetPublicIP().String()
		address = fmt.Sprintf("%s:%s", ip, port)
	}

	return &HostPort{
		Address: address,
		IP:      &ip,
		Port:    &port,
	}, nil
}

func ParsePeer(address string) (*HostPort, error) {
	components := strings.Split(address, "/")

	if len(components) != 5 || components[3] != "udp" ||
		components[0] != "" || (components[1] != "ip4" && components[1] != "dns") {
		return nil, fmt.Errorf("invalid peer provided: %s", address)
	}

	host, port := components[2], components[4]

	if validation.IsInvalidPort(port) {
		return nil, fmt.Errorf("invalid port provided: %s", address)
	}

	hostPort := &HostPort{
		Address: address,
		Port:    &port,
	}

	if ip, err := ParseIP(host); err != nil {
		hostPort.Host = &host
	} else {
		hostPort.IP = ip
	}

	return hostPort, nil
}

func ParseURL(address string) (*HostPort, error) {
	u, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	hostPort := &HostPort{
		Address: address,
	}

	hostName, port, path := u.Hostname(), u.Port(), u.Path
	if hostName == "" {
		return nil, fmt.Errorf("invalid url provided: %s", address)
	}

	hostPort.Host = &hostName

	if port != "" {
		if validation.IsInvalidPort(port) {
			return nil, fmt.Errorf("invalid port provided: %s", address)
		}

		hostPort.Port = &port
	}

	if path != "" {
		hostPort.Path = &path
	}

	ip, err := GetIPByDomain(address)
	if err == nil {
		hostPort.IP = ip
	}

	return hostPort, nil
}

func GetIPByDomain(address string) (*string, error) {
	ips, err := net.LookupIP(address)
	if err != nil {
		return nil, err
	}

	ip := ips[0].String()

	return &ip, nil
}

func ParseIP(address string) (*string, error) {
	if ip := net.ParseIP(address); ip != nil {
		if ip.IsLoopback() {
			ip = GetPublicIP()
		}

		ipResult := ip.String()

		return &ipResult, nil
	}

	return nil, fmt.Errorf("invalid ip provided: %s", address)
}

func GetPublicIP() net.IP {
	consensus := externalIP.DefaultConsensus(nil, nil)
	consensus.UseIPProtocol(4)

	ip, err := consensus.ExternalIP()
	if err != nil {
		return nil
	}

	return ip
}
