package files

import (
	"log"
	"os"
)

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

func EnsureDirectoryExists(directoryPath string) error {
	// Check if the directory already exists
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		// Create the directory
		err := os.Mkdir(directoryPath, 0755)
		if err != nil {
			log.Fatal("failed to create directory:", err)
		}
		log.Printf("Directory created: %s\n", directoryPath)
	} else if err != nil {
		log.Fatal("failed to check directory existence:", err)
	}

	return nil
}

func RemoveFileIfExists(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// File does not exist, return nil
		return nil
	}
	return os.Remove(filePath)
}

func RemoveDirectoryIfExists(directoryPath string) error {
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		// Directory does not exist, return nil
		return nil
	}
	return os.RemoveAll(directoryPath)
}

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
