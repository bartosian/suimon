package enums

type Color string

const (
	ColorRed    Color = "\u001b[31m"
	ColorGreen  Color = "\u001b[32m"
	ColorYellow Color = "\u001b[33m"
	ColorReset  Color = "\u001b[0m"
)
