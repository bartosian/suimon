package validation

import (
	"regexp"
	"strconv"
	"strings"
)

const domainRegexp = `^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`

func IsValidDomain(domain string) bool {
	match, _ := regexp.MatchString(domainRegexp, domain)

	return match
}

func IsValidPort(port string) bool {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return false
	}

	return portInt >= 1 || portInt <= 65535
}

func IsInvalidPort(port string) bool {
	return !IsValidPort(port)
}

func IsValidCharCount(str string, char string, count int) bool {
	charCount := strings.Count(str, char)
	if charCount != count {
		return false
	}
	return true
}
