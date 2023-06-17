// Package: fexec_test
package fexec_test

import (
	"log"
	"testing"

	"github.com/rudifa/goutil/pkg/fexec"
)

func TestRunCommand(t *testing.T) {

	log.Println("TestRunCommand")
	stdout, stderr, err := fexec.RunCommand("ls", "-l")

	if err != nil {
		t.Errorf("RunCommand failed: %v", err)
	}

	t.Logf("1. stdout: %v", stdout)
	t.Logf("1. stderr: %v", stderr)

	stdout, stderr, err = fexec.RunCommand("ls", "boing")

	if err == nil {
		t.Errorf("RunCommand should have failed")
	}

	t.Logf("2. stdout: %v", stdout)
	t.Logf("2. stderr: %v", stderr)
}
