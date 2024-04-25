package runcue_test

import (
	"log"
	"strings"
	"testing"

	"github.com/rudifa/goutil/runcue"
	"github.com/rudifa/goutil/util"
)

func TestRunCue(t *testing.T) {

	args := []string{"cue", "version"}
	runcue.RunCue(args...)

}

func TestExecuteFnAndCaptureStdoutAndStderr(t *testing.T) {

	wrapperFunc := func(args ...string) error {
		runcue.RunCue(args...)
		return nil
	}

	arg1 := "version"
	stdout, stderr, err := util.ExecuteFnAndCaptureStdoutAndStderr1(wrapperFunc, arg1)

	if err != nil {
		t.Errorf("got %v, want nil", err)
	}

	log.Println("stdout:", stdout)
	log.Println("stderr:", stderr)

	if !strings.Contains(stdout, "cue version") {
    	t.Errorf("Expected 'cue version' to be a substring of stdout, but it was not: %v", stdout)
	}

	if !strings.Contains(stdout, "go version") {
		t.Errorf("Expected 'go version' to be a substring of stdout, but it was not: %v", stdout)
	}

}
