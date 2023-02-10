package checker

import (
	"fmt"
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
