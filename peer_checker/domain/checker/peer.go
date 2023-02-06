package checker

import (
	"fmt"
	"net"

	emoji "github.com/jayco/go-emoji-flag"
	"github.com/oschwald/geoip2-golang"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/peer_checker/domain/enums"
)

type Location struct {
	CountryCode string
	CountryName string
	Flag        string
}

type Peer struct {
	Address     string
	AddressType enums.AddressType
	Port        string
	Location    Location

	rpcClient jsonrpc.RPCClient
}

func newPeer(geoDB *geoip2.Reader, address, port string) (*Peer, error) {
	peer := &Peer{
		Address: address,
		Port:    port,
	}

	if ip := net.ParseIP(address); ip != nil {
		peer.AddressType = enums.AddressTypeIP

		record, err := geoDB.Country(ip)
		if err == nil {
			peer.Location = Location{
				CountryCode: record.Country.IsoCode,
				CountryName: record.Country.Names["en"],
				Flag:        emoji.GetFlag(record.Country.IsoCode),
			}
		}
	} else if isValidDomain(address) {
		peer.AddressType = enums.AddressTypeDomain
	} else {
		return nil, fmt.Errorf("invalid ip/host value provided %s", address)
	}

	if !isValidPort(port) {
		return nil, fmt.Errorf("invalid port value provided %s", port)
	}

	return peer, nil
}
