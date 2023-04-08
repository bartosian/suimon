package cligw

import (
	"fmt"

	"github.com/fatih/color"
)

type MsgOpts struct {
	Indent int
}

func (ucg *Gateway) Info(label string, value string) {
	ucg.InfoWithOpts(label, value, MsgOpts{})
}

func (ucg *Gateway) InfoWithOpts(label string, value string, opts MsgOpts) {
	// Build label line
	labelLine := ucg.infoLabel(label, opts.Indent)

	// Build value line
	valueLine := ""
	if value != "" {
		valueLine = fmt.Sprintf("%s %s", color.New(color.FgGreen, color.Bold).Sprint("->"), value)
	}

	// Combine
	result := fmt.Sprintf("%s %s", labelLine, valueLine)

	// Render
	fmt.Println(result)
}

func (Gateway) infoLabel(label string, indent int) string {
	icon := "ℹ️ "
	for i := indent; i > 0; i-- {
		icon = fmt.Sprintf("  %s", icon)
	}

	bang := color.New(color.FgRed, color.Bold).Sprint(icon)
	formattedlabel := color.New(color.FgWhite, color.Bold).Sprint(label)

	return fmt.Sprintf("%s %s", bang, formattedlabel)
}
