package location

import (
	"fmt"
)

type Location struct {
	Provider    string
	CountryCode string
	CountryName string
	Flag        string
}

func NewLocation(countryCode, countryName, flag, company string) *Location {
	return &Location{
		Provider:    company,
		CountryCode: countryCode,
		CountryName: countryName,
		Flag:        flag,
	}
}

// String returns the string representation of the Location struct.
// Returns: a string representation of the Location struct.
func (loc *Location) String() string {
	if loc == nil {
		return ""
	}

	return fmt.Sprintf("%s  %s", loc.Flag, loc.CountryName)
}
