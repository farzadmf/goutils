package goutils

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"syscall"
)

// SpawnCommand runs a command and shows its output
func SpawnCommand(workingDir, cmdStr string) error {
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Dir = workingDir
	var stdout, stderr bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdout)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)

	return cmd.Run()
}

// ExecCommand runs a command and shows its output
func ExecCommand(workingDir, cmd string) error {
	bash, err := exec.LookPath("bash")
	if err != nil {
		return err
	}

	if err := syscall.Chdir(workingDir); err != nil {
		return err
	}

	return syscall.Exec(bash, []string{"bash", "-c", cmd}, os.Environ())
}

// RunCommand runs a shell command
func RunCommand(cmdString string) (string, error) {
	cmd := exec.Command("bash", "-c", cmdString)

	combined, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(string(combined))
	}

	return string(combined), nil
}
