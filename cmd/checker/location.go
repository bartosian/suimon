package checker

import (
	"fmt"
	"net"
)

type Location struct {
	Provider    string
	CountryCode string
	CountryName string
	Flag        string
}

func newLocation(countryCode, countryName, flag, company string) *Location {
	return &Location{
		Provider:    company,
		CountryCode: countryCode,
		CountryName: countryName,
		Flag:        flag,
	}
}

func (loc *Location) String() string {
	if loc == nil {
		return ""
	}

	return fmt.Sprintf("%s  %s", loc.Flag, loc.CountryName)
}

func (host *Host) SetLocation() {
	var parseLocation = func(ip net.IP) {
		record, err := host.ipClient.GetIPInfo(ip)
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

		host.Location = newLocation(countryISOCode, countryName, flag, company)
	}

	if host.HostPort.IP == nil {
		return
	}

	if ip := net.ParseIP(*host.HostPort.IP); ip != nil {
		parseLocation(ip)
	}
}
