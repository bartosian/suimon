package tables

import (
	"fmt"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/jedib0t/go-pretty/v6/text"
)

var nameTransformer = text.Transformer(func(val interface{}) string {
	return text.Bold.Sprint(val)
})

func GetTableTitleSUI(network enums.NetworkType, table enums.TableType) string {
	switch network {
	case enums.NetworkTypeTestnet:
		return fmt.Sprintf("💧 SUIMON %sv0.1.0%s %s[ %s %s ]%s", enums.ColorGreen, enums.ColorReset, enums.ColorRed, table, network, enums.ColorReset)
	case enums.NetworkTypeDevnet:
		fallthrough
	default:
		return fmt.Sprintf("💧 SUIMON %sv0.1.0%s %s[ %s %s ]%s", enums.ColorGreen, enums.ColorReset, enums.ColorRed, table, network, enums.ColorReset)
	}
}
