package checker

type Location struct {
	CountryCode string
	CountryName string
	Flag        string
}

func newLocation(countryCode, countryName, flag string) *Location {
	return &Location{
		CountryCode: countryCode,
		CountryName: countryName,
		Flag:        flag,
	}
}

func (loc *Location) String() string {
	if loc == nil {
		return ""
	}

	return loc.CountryName
}
