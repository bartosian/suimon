package checker

import (
	"fmt"
	"net"

	"github.com/oschwald/geoip2-golang"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/peer_checker/domain/enums"
)

type Peer struct {
	Address     string
	AddressType enums.AddressType
	Port        string
	Location    Location

	rpcClient   jsonrpc.RPCClient
	geoDbClient *geoip2.Reader
}

func newPeer(geoDB *geoip2.Reader, address, port string) *Peer {
	return &Peer{
		Address:     address,
		Port:        port,
		geoDbClient: geoDB,
	}
}

func (peer *Peer) Parse() error {
	if ip := net.ParseIP(peer.Address); ip != nil {
		peer.AddressType = enums.AddressTypeIP

		record, err := peer.geoDbClient.Country(ip)
		if err == nil {
			peer.Location = newLocation(record)
		}
	} else if isValidDomain(peer.Address) {
		peer.AddressType = enums.AddressTypeDomain
	} else {
		return fmt.Errorf("invalid ip/host value provided %s", peer.Address)
	}

	if !isValidPort(peer.Port) {
		return fmt.Errorf("invalid port value provided %s", peer.Port)
	}

	return nil
}
