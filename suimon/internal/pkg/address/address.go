package address

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	externalIP "github.com/glendc/go-external-ip"

	"github.com/bartosian/sui_helpers/suimon/internal/pkg/validation"
)

type Endpoint struct {
	Address string
	IP      *string
	Host    *string
	Path    *string
	Port    *string
	SSL     bool
}

func (hp *Endpoint) GetHostWithPath() *string {
	if hp.Host == nil {
		return nil
	}

	hostPath := *hp.Host

	if hp.Path != nil {
		hostPath = hostPath + *hp.Path
	}

	return &hostPath
}

func ParseIpPort(address string) (*Endpoint, error) {
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

	return &Endpoint{
		Address: address,
		IP:      &ip,
		Port:    &port,
	}, nil
}

func ParsePeer(address string) (*Endpoint, error) {
	components := strings.Split(address, "/")

	validLength := len(components) == 5
	validProtocol := components[3] == "udp"

	validFirstComponent := components[0] == ""
	validSecondComponent := components[1] == "ip4" || components[1] == "dns"

	host, port := components[2], components[4]

	if !validLength || !validProtocol || !validFirstComponent || !validSecondComponent {
		return nil, fmt.Errorf("invalid peer provided: %s", address)
	}

	if validation.IsInvalidPort(port) {
		return nil, fmt.Errorf("invalid port provided: %s", address)
	}

	endpoint := &Endpoint{
		Address: address,
		Port:    &port,
	}

	if ip, err := ParseIP(host); err != nil {
		endpoint.Host = &host
	} else {
		endpoint.IP = ip
	}

	return endpoint, nil
}

func ParseURL(address string) (*Endpoint, error) {
	if !strings.HasPrefix(address, "http") {
		address = "http://" + address
	}

	u, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	scheme, hostName, port, path := u.Scheme, u.Hostname(), u.Port(), u.Path
	if hostName == "" {
		return nil, fmt.Errorf("invalid url provided: %s", address)
	}

	endpoint := &Endpoint{
		Address: fmt.Sprintf("%s://%s%s", scheme, hostName, path),
		Host:    &hostName,
		SSL:     scheme == "https",
	}

	if ip, err := ParseIP(hostName); err == nil {
		endpoint.IP = ip
	}

	if port != "" {
		if validation.IsInvalidPort(port) {
			return nil, fmt.Errorf("invalid port provided: %s", address)
		}

		endpoint.Port = &port
	}

	if path != "" {
		endpoint.Path = &path
	}

	ip, err := GetIPByDomain(address)
	if err == nil {
		endpoint.IP = ip
	}

	return endpoint, nil
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
	if err := consensus.UseIPProtocol(4); err != nil {
		return nil
	}

	ip, err := consensus.ExternalIP()
	if err != nil {
		return nil
	}

	return ip
}
