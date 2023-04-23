package log

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Logger struct{}

func (logger *Logger) StreamFromService(serviceName string, stream chan string) error {
	var (
		cmdUnitLogs = "sudo journalctl -f -u %s -o cat"
		stdout      io.ReadCloser
		cmd         *exec.Cmd
		err         error
	)

	if !serviceExists(serviceName) {
		return fmt.Errorf("service %s not found", serviceName)
	}

	cmd = exec.Command("bash", "-c", fmt.Sprintf(cmdUnitLogs, serviceName))

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

	return fmt.Errorf("container with the image %s not found", imageName)
}

func (logger *Logger) StreamFromScreen(sessionName string, stream chan string) error {
	var (
		stdout    io.ReadCloser
		cmdAttach *exec.Cmd
		cmdDetach *exec.Cmd
		err       error
	)

	// attach screen for the logs piping
	cmdAttach = exec.Command("script", "-q", "-c", "screen -r "+sessionName, "/dev/null")

	if stdout, err = cmdAttach.StdoutPipe(); err != nil {
		return err
	}

	if err = cmdAttach.Start(); err != nil {
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

	// detach screen back
	cmdDetach = exec.Command("script", "-q", "-c", "screen -d "+sessionName, "/dev/null")
	if err = cmdDetach.Run(); err != nil {
		return err
	}

	if err = cmdAttach.Wait(); err != nil {
		return err
	}

	return nil
}

func RemoveNonPrintableChars(str string) string {
	reg := regexp.MustCompile("[^[:print:]\n]")
	return reg.ReplaceAllString(str, "")
}

func serviceExists(name string) bool {
	out, err := exec.Command("systemctl", "is-active", name).Output()
	if err != nil {
		return false
	}

	status := strings.TrimSpace(string(out))

	return status == "active"
}
