package enums

type Status string

const (
	StatusGreen  Status = "\U0001F7E2"
	StatusYellow Status = "\U0001F7E1"
	StatusRed    Status = "🔴"
)

var statusValues = map[Status]string{
	StatusGreen:  "+",
	StatusYellow: "+/-",
	StatusRed:    "-",
}

func (i Status) StatusToPlaceholder() string {
	return statusValues[i]
}
