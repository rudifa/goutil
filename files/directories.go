// Package files provides functions for working with files and directories.
package files

import (
	"log"
	"os"
	"path/filepath"
)

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

// RemoveDirectory removes a directory if it exists.
func RemoveDirectory(directoryPath string) error {
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		// Directory does not exist, return nil
		return nil
	}
	return os.RemoveAll(directoryPath)
}

// RemoveDirectoryIfExists removes a directory if it exists.
// Deprecated: Use RemoveDirectory instead.
func RemoveDirectoryIfExists(directoryPath string) error {
	return RemoveDirectory(directoryPath)
}

// RemoveFilesMatching removes files in the specified directory that match the given extension.
func RemoveFilesMatching(dir, extension string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if extension == "" || file.Name()[len(file.Name())-len(extension):] == extension {
			filePath := filepath.Join(dir, file.Name())
			err := os.Remove(filePath)
			if err != nil {
				log.Printf("failed to remove file: %s", filePath)
				return err
			}
		}
	}
	return nil
}

// ListFileNames returns a slice of file names in the specified directory
// that match the given extensions.
func ListFileNames(dir string, extensions ...string) ([]string, error) {
	// Open the directory
	d, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer d.Close()

	// Read the directory contents
	files, err := d.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	// Return all files if no extensions are provided
	if len(extensions) == 0 {
		return files, nil
	}

	// Add a dot to the extensions if it is missing
	for i, ext := range extensions {
		if ext[0] != '.' {
			extensions[i] = "." + ext
		}
	}

	// Filter files based on extensions
	var filteredFiles []string
	for _, file := range files {
		for _, ext := range extensions {
			if filepath.Ext(file) == ext {
				filteredFiles = append(filteredFiles, file)
				break
			}
		}
	}

	return filteredFiles, nil
}
