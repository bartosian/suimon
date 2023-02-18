package enums

import (
	"github.com/jedib0t/go-pretty/v6/text"
)

type Status string

const (
	StatusGreen  Status = "\U0001F7E9"
	StatusYellow Status = "\U0001F7E8"
	StatusRed    Status = "\U0001F7E5"
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

func (i Status) StatusToDashboard() string {
	switch i {
	case StatusRed:
		return "\U0001F7E5\U0001F7E5\U0001F7E5\n\U0001F7E5\U0001F7E5\U0001F7E5\n\U0001F7E5\U0001F7E5\U0001F7E5"
	case StatusYellow:
		return "\U0001F7E8\U0001F7E8\U0001F7E8\n\U0001F7E8\U0001F7E8\U0001F7E8\n\U0001F7E8\U0001F7E8\U0001F7E8"
	case StatusGreen:
		return "\U0001F7E9\U0001F7E9\U0001F7E9\n\U0001F7E9\U0001F7E9\U0001F7E9\n\U0001F7E9\U0001F7E9\U0001F7E9"
	}

	return ""
}
