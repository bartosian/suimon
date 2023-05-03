package cligw

import (
	"fmt"

	"github.com/fatih/color"
)

const warnIcon = "⚠️"

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
	icon := warnIcon
	for i := opts.Indent; i > 0; i-- {
		icon = fmt.Sprintf("  %s", icon)
	}

	formattedIcon := color.New(color.FgYellow, color.Bold).Sprint(icon)
	formattedMsg := color.New(color.FgWhite, color.Bold).Sprint(msg)
	result := fmt.Sprintf("%s  %s", formattedIcon, formattedMsg)

	fmt.Println(result)
}
