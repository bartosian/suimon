package checker

import (
	"fmt"
	"net"

	"github.com/oschwald/geoip2-golang"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/peer_checker/domain/enums"
	"github.com/bartosian/sui_helpers/peer_checker/pkg/validation"
)

const (
	rpcPortDefault     = "9000"
	metricsPortDefault = "9184"
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

	rpcClient   jsonrpc.RPCClient
	geoDbClient *geoip2.Reader

	TotalTransactionNumber  *uint64
	HighestSyncedCheckpoint *string
	SuiNetworkPeers         *string
	Uptime                  *string
}

func newPeer(geoDB *geoip2.Reader, address, port string) *Peer {
	peer := &Peer{
		Address:     address,
		Port:        port,
		geoDbClient: geoDB,
	}

	peer.rpcClient = jsonrpc.NewClient(peer.getUrl(requestTypeRPC, false))

	return peer
}

func (peer *Peer) Parse() error {
	if ip := net.ParseIP(peer.Address); ip != nil {
		peer.AddressType = enums.AddressTypeIP

		record, err := peer.geoDbClient.Country(ip)
		if err == nil {
			peer.Location = newLocation(record)
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
