package host

import (
	"net"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/location"
)

// SetLocation sets the location data for the Host struct by querying an IP address database.
// This method does not accept any parameters and does not return anything.
func (host *Host) SetLocation() {
	var parseLocation = func(ip net.IP) {
		record, err := host.clients.ipClient.GetIPInfo(ip)
		if err != nil {
			return
		}

		countryISOCode := record.Country
		countryName := record.CountryName
		flag := record.CountryFlag.Emoji

		var company string
		if record.Company != nil {
			company = record.Company.Name
		}

		host.Location = location.NewLocation(countryISOCode, countryName, flag, company)
	}

	if host.HostPort.IP == nil {
		return
	}

	if ip := net.ParseIP(*host.HostPort.IP); ip != nil {
		parseLocation(ip)
	}
}
