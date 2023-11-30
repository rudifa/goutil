// Package files provides functions for working with files and directories.
package files

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

// RemoveFileIfExists removes a file if it exists.
func RemoveFileIfExists(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// File does not exist, return nil
		return nil
	}
	return os.Remove(filePath)
}

// WriteBytes writes bytes to a file.
func WriteBytes(filepath string, bytes []byte) error {

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

// WriteString writes the string content to a file.
func WriteString(filepath, content string) error {
	return WriteToFile(filepath, content)
}

// WriteToFile writes the string content to a file.
func WriteToFile(filepath, content string) error {
	return WriteBytes(filepath, []byte(content))
}

// ReadBytes reads content from a file and returns it as a byte slice.
func ReadBytes(filepath string) ([]byte, error) {
	// Read the file contents using os.ReadFile
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// Return the content
	return content, nil
}

// ReadString reads content from a file and returns it as a string.
func ReadString(filepath string) (string, error) {
	bytes, err := ReadBytes(filepath)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Deprecated: ReadFromFile is deprecated. Use ReadString instead.
func ReadFromFile(filepath string) (string, error) {
	return ReadString(filepath)
}

func AppendTo(filename, text string) error {

	// Open the file in append mode. If the file doesn't exist, create it.
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening or creating file:", err)
		return err
	}
	defer file.Close()

	// Write the text and a newline to the file.
	_, err = fmt.Fprintln(file, text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil
}

// RemoveFrom removes text from the file
func RemoveFrom(filename, text string) error {
	// Read the file
	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Convert to string and replace the text
	content := string(contentBytes)
	content = strings.Replace(content, text, "", -1)

	// Write the modified content back to the file
	err = os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}

// Touch creates a file if it doesn't exist, or updates the modified time if it does.
func Touch(filename string) error {

	// Open the file in append mode. If the file doesn't exist, create it.
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening or creating file:", err)
		return err
	}
	file.Close()
	return nil
}

// CatFile prints the content of the file to stdout
func CatFile(copyFilename string) {
	content, err := os.ReadFile(copyFilename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}

// CopyFile copies a file from src to dst
func CopyFile(src, dst string) (err error) {
	// Open source file for reading
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// Sync to ensure all changes have been written to disk
	err = dstFile.Sync()
	if err != nil {
		return err
	}

	return nil
}
