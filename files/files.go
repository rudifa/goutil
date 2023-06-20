// Package files provides functions for working with files and directories.
package files

import (
	"log"
	"os"
)

// FileExists returns true if the file exists, false if not,
// or an error if it cannot be determined.
func FileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return !info.IsDir(), nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

// DirectoryExists returns true if the directory exists, false if not,
// or an error if it cannot be determined.
func DirectoryExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir(), nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

// CreateFileIfNotExists creates a file if it does not exist.
func CreateFileIfNotExists(filePath string) error {
	// Check if the file already exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Create the file
		_, err := os.Create(filePath)
		if err != nil {
			log.Fatal("failed to create file:", err)
		}
		log.Printf("File created: %s\n", filePath)
	} else if err != nil {
		log.Fatal("failed to check file existence:", err)
	}

	return nil
}

// EnsureDirectoryExists creates a directory if it does not exist.
func EnsureDirectoryExists(directoryPath string) error {
	// Check if the directory already exists
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		// Create the directory
		err := os.MkdirAll(directoryPath, 0755)
		if err != nil {
			log.Fatal("failed to create directory:", err)
		}
		log.Printf("Directory created: %s\n", directoryPath)
	} else if err != nil {
		log.Fatal("failed to check directory existence:", err)
	}

	return nil
}

// RemoveFileIfExists removes a file if it exists.
func RemoveFileIfExists(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// File does not exist, return nil
		return nil
	}
	return os.Remove(filePath)
}

// RemoveDirectoryIfExists removes a directory if it exists.
func RemoveDirectoryIfExists(directoryPath string) error {
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		// Directory does not exist, return nil
		return nil
	}
	return os.RemoveAll(directoryPath)
}

// WriteToFile writes content to a file.
func WriteToFile(filename string, content string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}