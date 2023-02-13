package enums

import (
	"github.com/jedib0t/go-pretty/v6/text"
)

type Status string

const (
	StatusGreen  Status = "\U0001F7E2"
	StatusYellow Status = "\U0001F7E1"
	StatusRed    Status = "🔴"
)

func (i Status) StatusToPlaceholder() string {
	return i.ColorStatus()
}

func (i Status) ColorStatus() string {
	colors := text.Colors{text.Bold}

	switch i {
	case StatusRed:
		colors = append(colors, text.BgRed, text.FgRed)
	case StatusYellow:
		colors = append(colors, text.BgYellow, text.FgYellow)
	case StatusGreen:
		colors = append(colors, text.BgGreen, text.FgGreen)
	}

	return colors.Sprint("|    |")
}
