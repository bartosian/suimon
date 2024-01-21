package log

import (
	"regexp"
)

// RemoveNonPrintableChars removes non-printable characters from the input string.
// It takes a string as input and returns a new string with non-printable characters removed.
// Non-printable characters are defined as any characters that are not visible when printed.
// The function uses a regular expression to replace non-printable characters with an empty string.
// It returns the modified string with only printable characters.
func RemoveNonPrintableChars(str string) string {
	reg := regexp.MustCompile("[^[:print:]\n]")
	return reg.ReplaceAllString(str, "")
}
