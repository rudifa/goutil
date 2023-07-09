// Package: files_test
package files_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/rudifa/goutil/files"
	//"github.com/stretchr/testify "
	"github.com/stretchr/testify/assert"
)

func TestCaptureStdoutAndStderr(t *testing.T) {

	args := []string{"Hello", "World", "of", "Go"}

	stdout, stderr := files.RunFnAndCaptureStdoutAndStderr(writeToStdoutAndStderr, args...)

	assert.Equal(t, "This is stdout [Hello World of Go]\n", stdout)
	assert.Equal(t, "This is stderr [Hello World of Go]\n", stderr)
}

// writeToStdoutAndStderr writes to stdout and stderr
func writeToStdoutAndStderr(args ...string) {

	fmt.Fprintln(os.Stdout, "This is stdout", args)
	fmt.Fprintln(os.Stderr, "This is stderr", args)
}
