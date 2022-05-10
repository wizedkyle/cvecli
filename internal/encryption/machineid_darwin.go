package encryption

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/wizedkyle/cvecli/internal/logging"
)

func extractId(lines string) string {
	for _, line := range strings.Split(lines, "\n") {
		if strings.Contains(line, "IOPlatformUUID") {
			parts := strings.SplitAfter(line, `" = "`)
			if len(parts) == 2 {
				return strings.TrimRight(parts[1], `""`)
			}
		}
	}
	return ""
}

func getMachineId() string {
	buffer := &bytes.Buffer{}
	err := runCommand(buffer, os.Stderr, "ioreg", "-rd1", "-c", "IOPlatformExpertDevice")
	if err != nil {
		logging.Console().Error().Err(err).Msg("failed to run command to get machine id")
	}
	machineId := extractId(buffer.String())
	if machineId == "" {
		logging.Console().Error().Msg("failed to retrieve machine id")
	}
	return trim(machineId)
}

func runCommand(stdout, stderr io.Writer, cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	command.Stdin = os.Stdin
	command.Stdout = stdout
	command.Stderr = stderr
	return command.Run()
}

func trim(s string) string {
	return strings.TrimSpace(strings.Trim(s, "\n"))
}
