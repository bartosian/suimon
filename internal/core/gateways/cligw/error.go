package cligw

import (
	"fmt"

	"github.com/fatih/color"
)

const errorIcon = "❗"

func (gateway *Gateway) Error(msg string) {
	gateway.ErrorWithOpts(msg, MsgOpts{})
}

func (gateway *Gateway) Errorf(msg string, vars ...interface{}) {
	gateway.ErrorfWithOpts(msg, MsgOpts{}, vars)
}

func (gateway *Gateway) ErrorfWithOpts(msg string, opts MsgOpts, vars ...interface{}) {
	msg = fmt.Sprintf(msg, vars)
	gateway.ErrorWithOpts(msg, opts)
}

func (Gateway) ErrorWithOpts(msg string, opts MsgOpts) {
	icon := errorIcon
	for i := opts.Indent; i > 0; i-- {
		icon = fmt.Sprintf("  %s", icon)
	}

	formattedIcon := color.New(color.FgRed, color.Bold).Sprint(icon)
	formattedMsg := color.New(color.FgWhite, color.Bold).Sprint(msg)
	result := fmt.Sprintf("%s %s", formattedIcon, formattedMsg)

	fmt.Println(result)
}
