package log

import (
	"fmt"
	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/enums"
)

type Logger struct{}

func NewLogger() Logger {
	return Logger{}
}

func (logger *Logger) Info(msg string, data interface{}) {
	colorPrint(enums.ColorGreen, msg, data)
}

func (logger *Logger) Warn(msg string, data interface{}) {
	colorPrint(enums.ColorYellow, msg, data)
}

func (logger *Logger) Error(msg string, data interface{}) {
	colorPrint(enums.ColorRed, msg, data)
}

func colorPrint(color enums.Color, messages ...any) {
	fmt.Println(color, messages, enums.ColorReset)
}
