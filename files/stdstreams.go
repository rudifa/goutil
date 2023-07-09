// Package files provides functions for capturing stdout and stderr streams.
package files

import (
	"io"
	"os"
)

// RunFnAndCaptureStdoutAndStderr runs the provided function and captures its stdout and stderr streams.
func RunFnAndCaptureStdoutAndStderr(fn func(...string), args ...string) (string, string) {
	// Save original stdout and stderr
	origStdout := os.Stdout
	origStderr := os.Stderr

	// Create pipes for capturing stdout and stderr
	stdoutReader, stdoutWriter, _ := os.Pipe()
	stderrReader, stderrWriter, _ := os.Pipe()

	// Redirect stdout and stderr to the pipes
	os.Stdout = stdoutWriter
	os.Stderr = stderrWriter

	// Call the provided function
	fn(args...)

	// Restore original stdout and stderr
	os.Stdout = origStdout
	os.Stderr = origStderr

	// Close the writers
	stdoutWriter.Close()
	stderrWriter.Close()

	// Read captured stdout and stderr
	var stdout, stderr string

	// Read from the pipes
	stdoutBytes, _ := io.ReadAll(stdoutReader)
	stdout = string(stdoutBytes)
	stderrBytes, _ := io.ReadAll(stderrReader)
	stderr = string(stderrBytes)

	// Close the readers
	stdoutReader.Close()
	stderrReader.Close()

	return stdout, stderr
}
