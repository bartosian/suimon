package cligw

import (
	"fmt"
	"log/slog"

	"github.com/fatih/color"
)

const infoIcon = "ℹ️ "

var (
	iconInfoColor    = color.New(color.FgGreen, color.Bold)
	messageInfoColor = color.New(color.FgWhite, color.Bold)
)

func (gateway *Gateway) Info(label, value string) {
	gateway.InfoWithOpts(label, value, MsgOpts{})
}

func (gateway *Gateway) InfoWithOpts(label, value string, opts MsgOpts) {
	labelLine := gateway.infoLabel(label, opts.Indent)

	valueLine := ""
	if value != "" {
		valueLine = fmt.Sprintf("%s %s", color.New(color.FgGreen, color.Bold).Sprint("->"), value)
	}

	result := fmt.Sprintf("%s %s", labelLine, valueLine)

	slog.Info(result)
}

func (Gateway) infoLabel(label string, indent int) string {
	var newIcon string

	for icon, i := infoIcon, indent; i > 0; i-- {
		newIcon = fmt.Sprintf("  %s", icon)
	}

	bang := iconInfoColor.Sprint(newIcon)
	formattedLabel := messageInfoColor.Sprint(label)

	return fmt.Sprintf("%s %s", bang, formattedLabel)
}
