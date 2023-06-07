package fexec

import (
	"bytes"
	"log"
	"os/exec"
)

// RunCommand executes a command and returns stdout, stderr, and error
func RunCommand(name string, args ...string) (string, string, error) {

	cmd := exec.Command(name, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("Command execution failed: %v", err)
	}
	return stdout.String(), stderr.String(), err
}
