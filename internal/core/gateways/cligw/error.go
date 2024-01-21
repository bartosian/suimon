//nolint:dupl // temporary disabled
package cligw

import (
	"fmt"

	"github.com/fatih/color"
)

const errorIcon = "❗"

var (
	iconErrColor    = color.New(color.FgRed, color.Bold)
	messageErrColor = color.New(color.FgWhite, color.Bold)
)

func (gateway *Gateway) Error(msg string) {
	gateway.ErrorWithOpts(msg, MsgOpts{})
}

func (gateway *Gateway) Errorf(msg string, vars ...interface{}) {
	gateway.ErrorfWithOpts(msg, MsgOpts{}, vars)
}

func (gateway *Gateway) ErrorfWithOpts(msg string, opts MsgOpts, vars ...interface{}) {
	msg = fmt.Sprintf(msg, vars...)

	gateway.ErrorWithOpts(msg, opts)
}

func (Gateway) ErrorWithOpts(msg string, opts MsgOpts) {
	var icon string

	for icon, i := errorIcon, opts.Indent; i > 0; i-- {
		icon = fmt.Sprintf("  %s", icon)
	}

	formattedIcon := iconErrColor.Sprint(icon)
	formattedMsg := messageErrColor.Sprint(msg)

	result := fmt.Sprintf("%s %s", formattedIcon, formattedMsg)

	fmt.Println(result)
}
