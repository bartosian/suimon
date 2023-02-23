package log

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"

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

func (logger *Logger) StreamFrom(processName string, stream chan string) error {
	var (
		command = fmt.Sprintf("sudo journalctl -f -u %s -o cat", processName)
		cmd     = exec.Command("bash", "-c", command)
		stdout  io.ReadCloser
		err     error
	)

	if stdout, err = cmd.StdoutPipe(); err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		select {
		case stream <- scanner.Text():
		case <-stream:
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return err
	}

	return nil
}
