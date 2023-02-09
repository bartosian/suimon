package log

import (
	"fmt"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
)

type Logger struct{}

func NewLogger() Logger {
	return Logger{}
}

func (logger *Logger) Info(messages ...any) {
	colorPrint(enums.ColorGreen, messages)
}

func (logger *Logger) Warn(messages ...any) {
	colorPrint(enums.ColorYellow, messages)
}

func (logger *Logger) Error(messages ...any) {
	colorPrint(enums.ColorRed, messages)
}

func colorPrint(color enums.Color, messages ...any) {
	fmt.Println(color, messages, enums.ColorReset)
}
