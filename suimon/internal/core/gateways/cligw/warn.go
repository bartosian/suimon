package cligw

import (
	"fmt"

	"github.com/fatih/color"
)

func (ucg *Gateway) Warn(msg string) {
	ucg.WarnWithOpts(msg, MsgOpts{})
}

func (ucg *Gateway) Warnf(msg string, vars ...interface{}) {
	ucg.WarnfWithOpts(msg, MsgOpts{}, vars)
}

func (ucg *Gateway) WarnfWithOpts(msg string, opts MsgOpts, vars ...interface{}) {
	msg = fmt.Sprintf(msg, vars)
	ucg.WarnWithOpts(msg, opts)
}

func (Gateway) WarnWithOpts(msg string, opts MsgOpts) {
	icon := "⚠️"
	for i := opts.Indent; i > 0; i-- {
		icon = fmt.Sprintf("  %s", icon)
	}

	formattedIcon := color.New(color.FgYellow, color.Bold).Sprint(icon)
	formattedMsg := color.New(color.FgWhite, color.Bold).Sprint(msg)
	result := fmt.Sprintf("%s  %s", formattedIcon, formattedMsg)

	fmt.Println(result)
}
