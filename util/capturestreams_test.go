// Fakestdio tests.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
package util

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
	"unicode/utf8"
)

func TestExecuteFnAndCaptureStdoutAndStderr(t *testing.T) {

	cmd := "some command"
	args := []string{"arg1", "arg2"}

	stdout, stderr, err := ExecuteFnAndCaptureStdoutAndStderr(writeToStdoutAndStdErr, cmd, args...)

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

func firstNChars(s string, n int) string {
	var i int
	for j := 0; j < n && i < len(s); j++ {
		_, size := utf8.DecodeRuneInString(s[i:])
		i += size
	}
	return s[:i]
}
