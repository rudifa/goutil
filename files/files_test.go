// Package: files_test
package files_test

import (
	"testing"

	"github.com/rudifa/goutil/files"
)

// TestWriteToFileAndReadFromFile tests WriteTo and ReadFromFile
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
