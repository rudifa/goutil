package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rudifa/goutil/pkg/files"
)

func main() {
	fmt.Println("Here we go")

	OsDemo()
}

func OsDemo() {
	err := files.EnsureDirectoryExists(".subdir")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(".subdir", 0755)
	if err == nil {
		log.Fatal(err) // should fail
	}

	err = os.MkdirAll(".subdir/a/b/c", 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = files.CreateFileIfNotExists(".subdir/a/b/c/d.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = files.FileExists(".subdir/a/b/c/d.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = files.RemoveDirectoryIfExists(".subdir")
	if err != nil {
		log.Fatal(err)
	}
}
