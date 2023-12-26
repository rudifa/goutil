/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package stacktrace_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/rudifa/goutil/stacktrace"
)

func TestStacktrace(t *testing.T) {

	st := stacktrace.NewStacktrace()

	fmt.Printf("stacktrace:\n")
	for _, line := range st.RawLines {
		fmt.Printf("  %s\n", line)
	}
}

const (
	sampleStackTrace2 = `cuelang.org/go/cue/format.LogStackTrace()
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/printer.go:506 +0x3a`

	sampleStackTrace3 = `main.runParseAndFormat({0x1bd7719, 0x13}, 0x0)
	/Users/rudifarkas/GitHub/golang/src/cue-issues-fmt-comments/main.go:104 +0x4f2`

	sampleStackTrace4 = `cuelang.org/go/cue/format.(*formatter).print(...)
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/format.go:257`
)

func TestExtractFuncSignature(t *testing.T) {

	const sampleStackTrace = `cuelang.org/go/cue/format.LogStackTrace()
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/printer.go:506 +0x3a
cuelang.org/go/cue/format.(*printer).append(0xc00031b2c0, {0x17c374c, 0x6}, {0xc0002dee40?, 0x1, 0x11a2d25?})
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/printer.go:454 +0x4aa
cuelang.org/go/cue/format.(*printer).writeString(0xc00031b2c0, {0x18ce528, 0x1}, 0x0)
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/printer.go:365 +0xc7
main.runParseAndFormat({0x1bd7719, 0x13}, 0x0)
	/Users/rudifarkas/GitHub/golang/src/cue-issues-fmt-comments/main.go:104 +0x4f2
main.runParseAndFormat_2567()
	/Users/rudifarkas/GitHub/golang/src/cue-issues-fmt-comments/main.go:37 +0x25
main.main()
	/Users/rudifarkas/GitHub/golang/src/cue-issues-fmt-comments/main.go:26 +0x58`

	extractedFromSampleStackTrace := `format.LogStackTrace

format.(*printer).append

format.(*printer).writeString

main.runParseAndFormat

main.runParseAndFormat_2567

main.main
`

	lines := strings.Split(sampleStackTrace, "\n")

	wantedSigs := strings.Split(extractedFromSampleStackTrace, "\n")

	for i, line := range lines {
		got := stacktrace.ExtractFuncSignature(line)
		wanted := wantedSigs[i]

		if got != wanted {
			t.Errorf("%d wanted |%s|, got |%s|", i, wanted, got)
		}
	}

	oneline := stacktrace.OnelineString(lines)

	wantedSigsString := strings.ReplaceAll(extractedFromSampleStackTrace, "\n\n", " => ")
	wantedSigsString = strings.ReplaceAll(wantedSigsString, "\n", "") // remove trailing newline

	if oneline != wantedSigsString {
		t.Errorf("wanted |%s|, got |%s|", wantedSigsString, oneline)
	}
}
