// Package: files_test
package files_test

import (
	"log"
	"testing"

	"github.com/rudifa/goutil/files"
)

func TestEnsureDirectoryExists(t *testing.T) {

	// test EnsureDirectoryExists
	log.Println("TestEnsureDirectoryExists")
	err := files.EnsureDirectoryExists("test")
	if err != nil {
		t.Errorf("EnsureDirectoryExists failed: %v", err)
	}

	err = files.EnsureDirectoryExists("test/A/B/C")
	if err != nil {
		t.Errorf("EnsureDirectoryExists failed: %v", err)
	}

	// remove test directory
	err = files.RemoveDirectoryIfExists("test")
	if err != nil {
		t.Errorf("RemoveDirectoryIfExists failed: %v", err)
	}
}
