// Fakestdio tests.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
package util_test

import (
	"fmt"

	"log"
	"os"
	"strconv"
	"testing"
	"unicode/utf8"

	"github.com/rudifa/goutil/util"
)

// TestExecuteFnAndCaptureStdoutAndStderr1 tests with short strings
func TestExecuteFnAndCaptureStdoutAndStderr1(t *testing.T) {

	args := []string{"some_command", "arg1", "arg2"}

	stdout, stderr, err := util.ExecuteFnAndCaptureStdoutAndStderr1(writeToStdoutAndStdErr1, args...)

	if err != nil {
		t.Errorf("got %v, want nil", err)
	}

	want := "[some_command arg1 arg2]"
	if stdout != want {
		t.Errorf("stdout got '%v', want '%v'", stdout, want)
	}
	if stderr != want {
		t.Errorf("stderr got '%v', want '%v'", stderr, want)
	}
}

// TestExecuteFnAndCaptureStdoutAndStderr tests with long strings (> 2.8 MB)
func TestExecuteFnAndCaptureStdoutAndStderr(t *testing.T) {

	cmd := "some command"
	args := []string{"arg1", "arg2"}

	stdout, stderr, err := util.ExecuteFnAndCaptureStdoutAndStderr2(writeToStdoutAndStdErr, cmd, args...)

	if err != nil {
		t.Errorf("got %v, want nil", err)
	}

	log.Println("---> len(stdout):", len(stdout))
	log.Println("---> len(stderr):", len(stderr))

	expected := "some command[arg1 arg2]0123456789"

	got := firstNChars(stdout, len(expected))
	if got != expected {
		t.Errorf("stdout got '%v', want '%v'", got, expected)
	}
	got = firstNChars(stderr, len(expected))
	if got != expected {
		t.Errorf("stderr got '%v', want '%v'", got, expected)
	}

	expectedLen := 2888913
	if len(stdout) != expectedLen {
		t.Errorf("got %v, want %v", len(stdout), expectedLen)
	}
	if len(stderr) != expectedLen {
		t.Errorf("got %v, want %v", len(stderr), expectedLen)
	}
}

func writeToStdoutAndStdErr(cmd string, args ...string) error {
	fmt.Fprint(os.Stdout, cmd, args)
	fmt.Fprint(os.Stderr, cmd, args)
	for i := 0; i < 500000; i++ {
		snippet := strconv.Itoa(i)
		fmt.Fprint(os.Stdout, snippet)
		fmt.Fprint(os.Stderr, snippet)
	}
	return nil
}

func writeToStdoutAndStdErr1(args ...string) error {
	fmt.Fprint(os.Stdout, args)
	fmt.Fprint(os.Stderr, args)
	return nil
}

func firstNChars(s string, n int) string {
	var i int
	for j := 0; j < n && i < len(s); j++ {
		_, size := utf8.DecodeRuneInString(s[i:])
		i += size
	}
	return s[:i]
}

// TestExecuteFnAndCaptureStdoutAndStderr3 tests also reading from stdin

func TestExecuteFnAndCaptureStdoutAndStderr3(t *testing.T) {

	input := "some data"
	args := []string{"arg1", "arg2"}

	stdout, stderr, err := util.ExecuteFnAndCaptureStdoutAndStderr3(readFromStdInAndWriteToStdoutAndStdErr, input, args...)

	if err != nil {
		t.Errorf("got %v, want nil", err)
	}

	want := "stdin:some"
	if stdout != want {
		t.Errorf("stdout got '%v', want '%v'", stdout, want)
	}

	want = "args:[arg1 arg2]"
	if stderr != want {
		t.Errorf("stderr got '%v', want '%v'", stderr, want)
	}
}

func readFromStdInAndWriteToStdoutAndStdErr(args ...string) error {
	// read from stdin
	var stdin string
	fmt.Scan(&stdin)
	// above reads only up to the first space
	// we should find a better way to read the whole line

	// write to stdout and stderr
	fmt.Fprint(os.Stdout, "stdin:", stdin)
	fmt.Fprint(os.Stderr, "args:", args)
	return nil
}
