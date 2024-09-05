// Package: files_test
package files_test

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/rudifa/goutil/files"
)

const testDir = ".test"

func TestEnsureDirectoryExists(t *testing.T) {

	// test EnsureDirectoryExists
	log.Println("TestEnsureDirectoryExists")
	err := files.EnsureDirectoryExists(testDir)
	if err != nil {
		t.Errorf("EnsureDirectoryExists failed: %v", err)
	}

	err = files.EnsureDirectoryExists(testDir + "/A/B/C")
	if err != nil {
		t.Errorf("EnsureDirectoryExists failed: %v", err)
	}

	// remove test directory
	err = files.RemoveDirectory(testDir)
	if err != nil {
		t.Errorf("RemoveDirectory failed: %v", err)
	}
}

func TestRemoveFilesMatching(t *testing.T) {
	log.Println("TestRemoveFilesMatching")

	testDir := "test" // Define the test directory

	// EnsureDirectoryExists
	err := files.EnsureDirectoryExists(testDir)
	if err != nil {
		t.Errorf("EnsureDirectoryExists failed: %v", err)
	}

	// Write test files
	err = files.WriteToFile(filepath.Join(testDir, "test1.txt"), "Hello World!")
	if err != nil {
		t.Errorf("WriteToFile failed: %v", err)
	}
	err = files.WriteToFile(filepath.Join(testDir, "test2.txt"), "Hello World!")
	if err != nil {
		t.Errorf("WriteToFile failed: %v", err)
	}
	err = files.WriteToFile(filepath.Join(testDir, "test3.txt"), "Hello World!")
	if err != nil {
		t.Errorf("WriteToFile failed: %v", err)
	}

	// check if files exist
	fileNames, err := files.ListFileNames(testDir)
	if err != nil {
		t.Errorf("ListFileNames failed: %v", err)
	}

	if len(fileNames) != 3 {
		t.Errorf("Expected 3 files in testDir, but found %d files", len(fileNames))
	}

	err = files.RemoveFilesMatching(testDir, "txt")
	if err != nil {
		t.Errorf("RemoveFilesMatching failed: %v", err)
	}

	// check if files exist
	fileNames, err = files.ListFileNames(testDir)
	if err != nil {
		t.Errorf("ListFileNames failed: %v", err)
	}

	if len(fileNames) != 0 {
		t.Errorf("Expected no files in testDir, but found %d files", len(fileNames))
	}

	// remove test directory
	err = files.RemoveDirectory(testDir)
	if err != nil {
		t.Errorf("RemoveDirectory failed: %v", err)
	}
}

func TestListFileNames(t *testing.T) {

	// test ListFileNames
	log.Println("TestListFileNames")
	goFiles, err := files.ListFileNames(".", "go")
	if err != nil {
		t.Errorf("ListFileNames failed: %v", err)
	}
	if len(goFiles) == 0 {
		t.Errorf("ListFileNames failed: no files found")
	}

	anyFiles, err := files.ListFileNames(".")
	if err != nil {
		t.Errorf("ListFileNames failed: %v", err)
	}
	if len(anyFiles) == 0 {
		t.Errorf("ListFileNames failed: no files found")
	}

	if len(goFiles) > len(anyFiles) {
		t.Errorf("ListFileNames failed: goFiles >= anyFiles")
	}
}
