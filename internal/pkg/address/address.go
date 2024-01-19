package address

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	externalIP "github.com/glendc/go-external-ip"

	"github.com/bartosian/suimon/internal/pkg/validation"
)

type Endpoint struct {
	Address string
	IP      *string
	Host    *string
	Path    *string
	Port    *string
	SSL     bool
}

const (
	errInvalidPeerFormatProvided = "invalid peer format provided: %s"
	errInvalidPortProvided       = "invalid port provided: %s"
	errInvalidIpProvided         = "invalid ip provided: %s"
	errInvalidUrlProvided        = "invalid url provided: %s"
)

// GetHostWithPath returns the host with the path, if available.
// If the host is nil, it returns nil. Otherwise, it concatenates the host and path and returns the result.
// Returns a pointer to the concatenated host with the path.
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

// ParseIpPort parses the given address and returns an Endpoint and an error.
// If the address is in the format "ip:port", it returns the IP and port as an Endpoint.
// If the address is in the format "protocol/ip4/host/udp/port", it returns the host and port as an Endpoint.
// If the IP provided is a loopback or unspecified IP, it replaces it with the public IP.
// Returns the parsed Endpoint and nil error if successful, otherwise returns nil and an error.
func ParseIpPort(address string) (*Endpoint, error) {
	ip, port, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}

	if validation.IsInvalidPort(port) {
		return nil, fmt.Errorf(errInvalidPortProvided, address)
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

// ParsePeer parses the given address and returns an Endpoint and an error.
// If the address is in the format "/ip4/host/udp/port", it returns the host and port as an Endpoint.
// If the IP provided is a loopback or unspecified IP, it replaces it with the public IP.
// Returns the parsed Endpoint and nil error if successful, otherwise returns nil and an error.
func ParsePeer(address string) (*Endpoint, error) {
	components := strings.Split(address, "/")

	if len(components) != 5 {
		return nil, fmt.Errorf(errInvalidPeerFormatProvided, address)
	}

	validProtocol := components[3] == "udp"
	validFirstComponent := components[0] == ""
	validSecondComponent := components[1] == "ip4" || components[1] == "dns"

	if !validProtocol || !validFirstComponent || !validSecondComponent {
		return nil, fmt.Errorf(errInvalidPeerFormatProvided, address)
	}

	host, port := components[2], components[4]

	if validation.IsInvalidPort(port) {
		return nil, fmt.Errorf(errInvalidPortProvided, address)
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

// ParseURL parses the given address and returns an Endpoint and an error.
// If the address does not start with "http", it adds "http://" to the beginning of the address.
// It then parses the address using the url.Parse function and constructs an Endpoint with the parsed information.
// Returns the parsed Endpoint and nil error if successful, otherwise returns nil and an error.
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
		return nil, fmt.Errorf(errInvalidUrlProvided, address)
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
			return nil, fmt.Errorf(errInvalidPortProvided, address)
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

// GetIPByDomain performs a DNS lookup to retrieve the IP address associated with the provided domain.
// It takes the domain name as input and returns the IP address and nil error if successful, otherwise returns nil and an error.
func GetIPByDomain(address string) (*string, error) {
	ips, err := net.LookupIP(address)
	if err != nil {
		return nil, err
	}

	ip := ips[0].String()

	return &ip, nil
}

// ParseIP parses the given IP address and returns the parsed IP and nil error if successful, otherwise returns nil and an error.
// If the provided IP is a loopback or unspecified IP, it replaces it with the public IP.
func ParseIP(address string) (*string, error) {
	if ip := net.ParseIP(address); ip != nil {
		if ip.IsLoopback() {
			ip = GetPublicIP()
		}

		ipResult := ip.String()

		return &ipResult, nil
	}

	return nil, fmt.Errorf(errInvalidIpProvided, address)
}

// GetPublicIP retrieves the public IP address using the default consensus and returns it.
// It returns the public IP address if successful, otherwise returns nil.
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
