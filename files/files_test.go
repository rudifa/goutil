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

// TestWriteToandreadFromFile tests WriteTo and ReadFromFile
func TestWriteToFileAndReadFromFile(t *testing.T) {

	err := files.WriteToFile("test.txt", "Hello World!")
	if err != nil {
		t.Errorf("WriteTo failed: %v", err)
	}

	content, err := files.ReadFromFile("test.txt")
	if err != nil {
		t.Errorf("ReadFromFile failed: %v", err)
	}
	if content != "Hello World!" {
		t.Errorf("ReadFromFile failed: %v", err)
	}

	// remove test.txt
	err = files.RemoveFileIfExists("test.txt")
	if err != nil {
		t.Errorf("RemoveFileIfExists failed: %v", err)
	}
}
