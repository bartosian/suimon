package tables

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/suimon/cmd/checker/enums"
)

const suiEmoji = "💧 "

var nameTransformer = text.Transformer(func(val interface{}) string {
	return text.Bold.Sprint(val)
})

func GetTableTitleSUI(network enums.NetworkType, table enums.TableType, emojisEnabled bool) string {
	var emoji string
	if emojisEnabled {
		emoji = suiEmoji
	}

	switch network {
	case enums.NetworkTypeTestnet:
		return fmt.Sprintf("%sSUIMON %sv0.1.0%s %s[ %s %s ]%s", emoji, enums.ColorGreen, enums.ColorReset, enums.ColorRed, table, network, enums.ColorReset)
	case enums.NetworkTypeDevnet:
		fallthrough
	default:
		return fmt.Sprintf("%sSUIMON %sv0.1.0%s %s[ %s %s ]%s", emoji, enums.ColorGreen, enums.ColorReset, enums.ColorRed, table, network, enums.ColorReset)
	}
}
