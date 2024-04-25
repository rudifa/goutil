// package main

package main

import (
	"fmt"
	"os"

	"github.com/rudifa/goutil/runcue"
)

func main() {

	fmt.Println("Here we go.")

	// if the first argument is "cue", then run the cue command with the remaining arguments
	// e.g. `goutil cue eval runcue/testdata/sample.cue``
	// or `goutil cue version`
	if len(os.Args) > 1 && os.Args[1] == "cue" {
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "-h")
		}
		RunCueDemo(os.Args[2:]...)
		return
	}
}

func RunCueDemo(args ...string) {
	fmt.Println("RunCueDemo: ", args)
	runcue.RunCue(args...)
}
