package checker

import (
	"fmt"
	"net"
	"net/http"

	emoji "github.com/jayco/go-emoji-flag"
	"github.com/oschwald/geoip2-golang"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/sui-peer-checker/pkg/log"
	"github.com/bartosian/sui_helpers/sui-peer-checker/pkg/validation"
)

const (
	rpcPortDefault      = "9000"
	metricsPortDefault  = "9184"
	countryNameLanguage = "en"
)

type requestType int

const (
	requestTypeRPC requestType = iota
	requestTypeMetrics
)

type Peer struct {
	Address     string
	AddressType enums.AddressType
	Port        string

	Location *Location

	Metrics Metrics

	rpcClient   jsonrpc.RPCClient
	httpClient  *http.Client
	geoDbClient *geoip2.Reader

	logger log.Logger
}

func newPeer(
	geoDB *geoip2.Reader,
	httpClient *http.Client,
	address, port string,
) *Peer {
	peer := &Peer{
		Address:     address,
		Port:        port,
		geoDbClient: geoDB,
		logger:      log.NewLogger(),
	}

	rpcURL := peer.getUrl(requestTypeRPC, false, nil)

	peer.rpcClient = jsonrpc.NewClient(rpcURL)
	peer.httpClient = httpClient

	return peer
}

func (peer *Peer) Parse() error {
	if ip := net.ParseIP(peer.Address); ip != nil {
		peer.AddressType = enums.AddressTypeIP

		record, err := peer.geoDbClient.Country(ip)
		if err == nil {
			countryISOCode := record.Country.IsoCode
			countryName := record.Country.Names[countryNameLanguage]
			flag := emoji.GetFlag(record.Country.IsoCode)
			peer.Location = newLocation(countryISOCode, countryName, flag)
		}
	} else if validation.IsValidDomain(peer.Address) {
		peer.AddressType = enums.AddressTypeDomain
	} else {
		return fmt.Errorf("invalid ip/host value provided %s", peer.Address)
	}

	if !validation.IsValidPort(peer.Port) {
		return fmt.Errorf("invalid port value provided %s", peer.Port)
	}

	return nil
}

func (peer *Peer) getUrl(
	request requestType,
	secure bool,
	port *string,
) string {
	protocol := "http"
	if secure {
		protocol = protocol + "s"
	}

	portRPC := rpcPortDefault
	portMetrics := metricsPortDefault

	switch request {
	case requestTypeRPC:
		if port != nil {
			portRPC = *port
		}

		return fmt.Sprintf("%s://%s:%s", protocol, peer.Address, portRPC)
	case requestTypeMetrics:
		fallthrough
	default:
		if port != nil {
			portMetrics = *port
		}

		return fmt.Sprintf("%s://%s:%s/metrics", protocol, peer.Address, portMetrics)
	}
}
