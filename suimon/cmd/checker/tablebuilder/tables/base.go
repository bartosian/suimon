package tables

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
)

const suiEmoji = "💧 "

var nameTransformer = text.Transformer(func(val interface{}) string {
	return text.Bold.Sprint(val)
})

func GetTableTitle(network enums.NetworkType, table enums.TableType, emojisEnabled bool) string {
	var emoji string
	if emojisEnabled {
		emoji = suiEmoji
	}

	switch network {
	case enums.NetworkTypeTestnet:
		return fmt.Sprintf("%s%sSUIMON%s %s[ %s %s ]%s", emoji, enums.ColorGreen, enums.ColorReset, enums.ColorRed, table, network, enums.ColorReset)
	case enums.NetworkTypeDevnet:
		fallthrough
	default:
		return fmt.Sprintf("%s%sSUIMON%s %s[ %s %s ]%s", emoji, enums.ColorGreen, enums.ColorReset, enums.ColorRed, table, network, enums.ColorReset)
	}
}
