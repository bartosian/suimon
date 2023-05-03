package cligw

import (
	"fmt"

	"github.com/fatih/color"
)

const infoIcon = "ℹ️ "

type MsgOpts struct {
	Indent int
}

func (gateway *Gateway) Info(label string, value string) {
	gateway.InfoWithOpts(label, value, MsgOpts{})
}

func (gateway *Gateway) InfoWithOpts(label string, value string, opts MsgOpts) {
	labelLine := gateway.infoLabel(label, opts.Indent)

	valueLine := ""
	if value != "" {
		valueLine = fmt.Sprintf("%s %s", color.New(color.FgGreen, color.Bold).Sprint("->"), value)
	}

	result := fmt.Sprintf("%s %s", labelLine, valueLine)

	fmt.Println(result)
}

func (Gateway) infoLabel(label string, indent int) string {
	icon := infoIcon
	for i := indent; i > 0; i-- {
		icon = fmt.Sprintf("  %s", icon)
	}

	bang := color.New(color.FgRed, color.Bold).Sprint(icon)
	formattedlabel := color.New(color.FgWhite, color.Bold).Sprint(label)

	return fmt.Sprintf("%s %s", bang, formattedlabel)
}
