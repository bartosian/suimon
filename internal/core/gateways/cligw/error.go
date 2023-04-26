package cligw

import (
	"fmt"

	"github.com/fatih/color"
)

func (ucg *Gateway) Error(msg string) {
	ucg.ErrorWithOpts(msg, MsgOpts{})
}

func (ucg *Gateway) Errorf(msg string, vars ...interface{}) {
	ucg.ErrorfWithOpts(msg, MsgOpts{}, vars)
}

func (ucg *Gateway) ErrorfWithOpts(msg string, opts MsgOpts, vars ...interface{}) {
	msg = fmt.Sprintf(msg, vars)
	ucg.ErrorWithOpts(msg, opts)
}

func (Gateway) ErrorWithOpts(msg string, opts MsgOpts) {
	icon := "❗"
	for i := opts.Indent; i > 0; i-- {
		icon = fmt.Sprintf("  %s", icon)
	}

	formattedIcon := color.New(color.FgRed, color.Bold).Sprint(icon)
	formattedMsg := color.New(color.FgWhite, color.Bold).Sprint(msg)
	result := fmt.Sprintf("%s %s", formattedIcon, formattedMsg)

	fmt.Println(result)
}
