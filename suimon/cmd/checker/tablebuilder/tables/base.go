package tables

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
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

//func TableStyleFromString(value string) (table.ColorOptions, error) {
//	value = strings.ToUpper(strings.TrimSpace(value))
//
//	result, ok := map[string]table.ColorOptions{
//		"WHITE":                     {
//			Header: text.Colors{text.FgHiWhite},
//			Row:    text.Colors{text.FgHiWhite},
//			Footer: text.Colors{text.FgHiWhite},
//		},
//		"DARK":                      text.FgHiBlack,
//		"COLOR": text.Fg
//		"COMMIT":                    MetricTypeCommit,
//		"HIGHEST_SYNCED_CHECKPOINT": MetricTypeHighestSyncedCheckpoint,
//		"SUI_NETWORK_PEERS":         MetricTypeSuiNetworkPeers,
//		"TOTAL_TRANSACTIONS_NUMBER": MetricTypeTotalTransactionsNumber,
//		"LATEST_CHECKPOINT":         MetricTypeLatestCheckpoint,
//	}[value]
//
//	if ok {
//		return result, nil
//	}
//
//	return MetricTypeUndefined, fmt.Errorf("unsupported metric type enum string: %s", value)
//}
