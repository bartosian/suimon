package log

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

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

func (logger *Logger) StreamFromProcess(processName string, stream chan string) error {
	var (
		cmdService = "sudo journalctl -f -u %s -o cat"
		cmdPID     = "sudo journalctl -f _PID=%s -o cat"
		stdout     io.ReadCloser
		cmd        *exec.Cmd
		err        error
	)

	if serviceExists(processName) {
		cmd = exec.Command("bash", "-c", fmt.Sprintf(cmdService, processName))
	} else {
		var pid string

		if pid, err = getPID(processName); err != nil {
			return err
		}

		cmd = exec.Command("bash", "-c", fmt.Sprintf(cmdPID, pid))
	}

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

func (logger *Logger) StreamFromContainer(imageName string, stream chan string) error {
	var (
		cli        *client.Client
		containers []types.Container
		err        error
	)

	if cli, err = client.NewClientWithOpts(client.FromEnv); err != nil {
		return err
	}

	if containers, err = cli.ContainerList(context.Background(), types.ContainerListOptions{}); err != nil {
		return err
	}

	for _, container := range containers {
		imageID := container.ImageID
		imageInspect, _, err := cli.ImageInspectWithRaw(context.Background(), imageID)
		if err != nil {
			return err
		}

		if strings.Contains(imageInspect.RepoTags[0], imageName) && container.State == "running" {
			var logs io.ReadCloser

			if logs, err = cli.ContainerLogs(context.Background(), container.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Follow: true}); err != nil {
				return err
			}

			defer logs.Close()

			scanner := bufio.NewScanner(logs)
			for scanner.Scan() {
				select {
				case stream <- scanner.Text():
				case <-stream:
					break
				}
			}

			if err != nil && err.Error() != "EOF" {
				return err
			}

			break
		}
	}

	return nil
}

func RemoveNonPrintableChars(str string) string {
	reg := regexp.MustCompile("[^[:print:]\n]")
	return reg.ReplaceAllString(str, "")
}

func serviceExists(name string) bool {
	var (
		cmd    = exec.Command("systemctl", "status", name)
		output bytes.Buffer
		err    error
	)

	cmd.Stdout = &output

	if err = cmd.Run(); err != nil {
		return false
	}

	return true
}

func getPID(command string) (string, error) {
	var (
		cmd    = exec.Command("pgrep", command)
		output bytes.Buffer
		err    error
	)

	cmd.Stdout = &output

	if err = cmd.Run(); err != nil {
		return "", err
	}

	return output.String(), nil
}

func PrintLogo(text string, fontName string, color string) {
	logo := figure.NewColorFigure(text, fontName, color, true)
	fmt.Print("\n\n")
	logo.Print()
	fmt.Print("\n\n")
}
