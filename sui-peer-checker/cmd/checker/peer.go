package checker

import (
	"fmt"
	emoji "github.com/jayco/go-emoji-flag"
	"net"
	"net/http"

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
	Location    *Location

	Metrics Metrics

	rpcClient   jsonrpc.RPCClient
	httpClient  *http.Client
	geoDbClient *geoip2.Reader

	logger log.Logger
}

func newPeer(geoDB *geoip2.Reader, httpClient *http.Client, address, port string) *Peer {
	peer := &Peer{
		Address:     address,
		Port:        port,
		geoDbClient: geoDB,
		logger:      log.NewLogger(),
	}

	peer.rpcClient = jsonrpc.NewClient(peer.getUrl(requestTypeRPC, false))
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

func (peer *Peer) getUrl(request requestType, secure bool) string {
	protocol := "http"

	if secure {
		protocol = protocol + "s"
	}

	switch request {
	case requestTypeRPC:
		return fmt.Sprintf("%s://%s:%s", protocol, peer.Address, rpcPortDefault)
	case requestTypeMetrics:
		fallthrough
	default:
		return fmt.Sprintf("%s://%s:%s/metrics", protocol, peer.Address, metricsPortDefault)
	}
}
