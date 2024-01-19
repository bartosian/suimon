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

// StreamFromService streams logs from the specified service to the provided channel.
// It takes the service name and a channel for streaming logs as input and returns an error if any.
// If the service does not exist, it returns an error indicating that the service was not found.
// If an error occurs while streaming logs, it returns the corresponding error.
// If the streaming is successful, it returns nil.
// The logs are streamed using the 'journalctl' command with the specified service name.
// The logs are read line by line and sent to the provided channel for further processing.
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

// StreamFromContainer streams logs from the running containers with the specified image name to the provided channel.
// It takes the image name and a channel for streaming logs as input and returns an error if any.
// If the container with the specified image name is not found or no running containers are found with the specified image name, it returns an error indicating that the container was not found.
// If an error occurs while streaming logs, it returns the corresponding error.
// If the streaming is successful, it returns nil.
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

// StreamFromScreen streams the output from a screen session to the provided channel.
// It takes the sessionName as input and the stream channel to send the output to.
// It returns an error if there was an issue streaming the output.
// The function uses the 'script' command to attach to and detach from the screen session.
// It reads the output from the attached session and sends it to the provided channel.
// If there is an error during the process, it returns the error.
func (logger *Logger) StreamFromScreen(sessionName string, stream chan string) error {
	var (
		stdout    io.ReadCloser
		cmdAttach *exec.Cmd
		cmdDetach *exec.Cmd
		err       error
	)

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

	cmdDetach = exec.Command("script", "-q", "-c", "screen -d "+sessionName, "/dev/null")
	if err = cmdDetach.Run(); err != nil {
		return err
	}

	if err = cmdAttach.Wait(); err != nil {
		return err
	}

	return nil
}

// RemoveNonPrintableChars removes non-printable characters from the input string.
// It takes a string as input and returns a new string with non-printable characters removed.
// Non-printable characters are defined as any characters that are not visible when printed.
// The function uses a regular expression to replace non-printable characters with an empty string.
// It returns the modified string with only printable characters.
func RemoveNonPrintableChars(str string) string {
	reg := regexp.MustCompile("[^[:print:]\n]")
	return reg.ReplaceAllString(str, "")
}

// serviceExists checks if the specified service exists and is active.
// It takes the service name as input and returns true if the service exists and is active, otherwise returns false.
// If an error occurs while checking the service status, it returns false.
// It uses the 'systemctl is-active' command to check the status of the service.
// Returns true if the service is active, false otherwise.
func serviceExists(name string) bool {
	out, err := exec.Command("systemctl", "is-active", name).Output()
	if err != nil {
		return false
	}

	status := strings.TrimSpace(string(out))

	return status == "active"
}
