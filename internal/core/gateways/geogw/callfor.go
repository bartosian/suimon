package geogw

import (
	"fmt"
	"net"

	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
)

func (gateway *Gateway) CallFor(ip net.IP) (result *ports.IPResult, err error) {
	if ip == nil {
		return nil, fmt.Errorf("no IP provided")
	}

	data, err := gateway.client.GetIPInfo(ip)
	if err != nil {
		return nil, fmt.Errorf("failed to get IP data for %s: %w", ip, err)
	}

	company := new(ports.Company)

	if data.Company != nil {
		company = &ports.Company{
			Name:   data.Company.Name,
			Domain: data.Company.Domain,
			Type:   data.Company.Type,
		}
	}

	return &ports.IPResult{
		IP:           data.IP,
		Hostname:     data.Hostname,
		City:         data.City,
		Region:       data.Region,
		Country:      data.Country,
		CountryName:  data.CountryName,
		CountryEmoji: data.CountryFlag.Emoji,
		Location:     data.Location,
		Company:      company,
	}, nil
}
