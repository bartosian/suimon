package checker

import (
	"fmt"
	"net"

	"github.com/bartosian/sui_helpers/peer_checker/domain/enums"
)

type Peer struct {
	Address     string
	AddressType enums.AddressType
	Port        string
}

func newPeer(address, port string) (*Peer, error) {
	var addressType enums.AddressType

	if net.ParseIP(address) != nil {
		addressType = enums.AddressTypeIP
	} else if isValidDomain(address) {
		addressType = enums.AddressTypeDomain
	} else {
		return nil, fmt.Errorf("invalid ip/host value provided %s", address)
	}

	if !isValidPort(port) {
		return nil, fmt.Errorf("invalid port value provided %s", port)
	}

	return &Peer{
		Address:     address,
		AddressType: addressType,
		Port:        port,
	}, nil
}
