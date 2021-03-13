package goutils

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"syscall"
)

// SpawnCommand runs a command and shows its output. It's going to start the command and give
//      control to that process, so if we, say, do Ctrl+C, we're basically existing from the whole
//      application and we won't return to our original process
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

// RunCommand runs a shell command as a subprocess and returns its output
func RunCommand(cmdString string) (string, error) {
	cmd := exec.Command("bash", "-c", cmdString)

	combined, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(string(combined))
	}

	return string(combined), nil
}

// StartCommand uses StartProcess to start a command. It's going to start the command, but the
//      "control" is still in the calling process, so when the command finishes, we return back
//      to our original process
func StartCommand(workingDir, cmdStr string) error {
	bash, err := exec.LookPath("bash")
	if err != nil {
		return err
	}

	var attr os.ProcAttr
	attr.Files = []*os.File{
		os.Stdin,
		os.Stderr,
		os.Stdout,
	}
	attr.Dir = workingDir

	proc, err := os.StartProcess(bash, []string{"bash", "-c", cmdStr}, &attr)
	proc.Wait()

	return err
}
