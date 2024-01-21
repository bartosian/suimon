package enums

import (
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
)

const statusWidgetLength = 100

type Status string

const (
	StatusGreen  Status = "\U0001F7E9"
	StatusYellow Status = "\U0001F7E8"
	StatusRed    Status = "\U0001F7E5"
	StatusGrey   Status = "\U0001F7E4"
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
	case StatusGrey:
	}

	return colors.Sprint("|    |")
}

func (i Status) DashboardStatus() string {
	statusWidget := make([]string, statusWidgetLength)

	repeatedPattern := strings.Repeat("    ", statusWidgetLength)

	for idx := range statusWidget {
		statusWidget[idx] = repeatedPattern
	}

	return strings.Join(statusWidget, "\n")
}
