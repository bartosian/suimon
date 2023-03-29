package enums

import (
	"strings"
)

type ColorTable string

const (
	ColorTableWhite ColorTable = "WHITE"
	ColorTableDark  ColorTable = "DARK"
	ColorTableColor ColorTable = "COLOR"
)

func (e ColorTable) String() string {
	return string(e)
}

func ColorTableFromString(value string) ColorTable {
	value = strings.ToUpper(strings.TrimSpace(value))

	result, ok := map[string]ColorTable{
		"WHITE": ColorTableWhite,
		"DARK":  ColorTableDark,
		"COLOR": ColorTableColor,
	}[value]

	if ok {
		return result
	}

	return ColorTableWhite
}
