//nolint:dupl // these files are different.
package cligw

import (
	"fmt"

	"github.com/fatih/color"
)

const warnIcon = "⚠️"

var (
	iconWarnColor    = color.New(color.FgYellow, color.Bold)
	messageWarnColor = color.New(color.FgWhite, color.Bold)
)

func (gateway *Gateway) Warn(msg string) {
	gateway.WarnWithOpts(msg, MsgOpts{})
}

func (gateway *Gateway) Warnf(msg string, vars ...interface{}) {
	gateway.WarnfWithOpts(msg, MsgOpts{}, vars)
}

func (gateway *Gateway) WarnfWithOpts(msg string, opts MsgOpts, vars ...interface{}) {
	msg = fmt.Sprintf(msg, vars)

	gateway.WarnWithOpts(msg, opts)
}

func (Gateway) WarnWithOpts(msg string, opts MsgOpts) {
	var newIcon string

	for icon, i := warnIcon, opts.Indent; i > 0; i-- {
		newIcon = fmt.Sprintf("  %s", icon)
	}

	formattedIcon := iconWarnColor.Sprint(newIcon)
	formattedMsg := messageWarnColor.Sprint(msg)

	result := fmt.Sprintf("%s  %s", formattedIcon, formattedMsg)

	fmt.Println(result)
}
