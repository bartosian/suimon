package checker

import (
	"fmt"

	emoji "github.com/jayco/go-emoji-flag"
	"github.com/oschwald/geoip2-golang"
)

const countryNameLanguage = "en"

type Location struct {
	CountryCode string
	CountryName string
	Flag        string
}

func newLocation(country *geoip2.Country) *Location {
	return &Location{
		CountryCode: country.Country.IsoCode,
		CountryName: country.Country.Names[countryNameLanguage],
		Flag:        emoji.GetFlag(country.Country.IsoCode),
	}
}

func (loc *Location) String() string {
	return fmt.Sprintf("%s  %s", loc.Flag, loc.CountryName)
}
