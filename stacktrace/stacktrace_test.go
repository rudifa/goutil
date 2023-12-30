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

	st := stacktrace.CapturedStacktrace()

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

const longerStackTrace = `printer.go:504: Stack trace:
goroutine 1 [running]:
cuelang.org/go/cue/format.LogStackTrace()
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/printer.go:503 +0x45
cuelang.org/go/cue/format.(*printer).append(0xc0003d1900, {0x17d9776, 0x6}, {0xc00055a788?, 0x1, 0x17ecf9b?})
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/printer.go:457 +0x478
cuelang.org/go/cue/format.(*printer).writeString(0xc0003d1900, {0xc0004b3840, 0xe}, 0x1)
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/printer.go:366 +0xd9
cuelang.org/go/cue/format.(*printer).Print(0xc0003d1900, {0x1760a80?, 0xc000479800?})
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/printer.go:252 +0x81a
cuelang.org/go/cue/format.(*formatter).printComment(0xc00055bd40, 0xc0004d8ea0)
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/format.go:339 +0x16c
cuelang.org/go/cue/format.(*formatter).visitComments(0xc00055bd40, 0x3)
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/format.go:317 +0x8c
cuelang.org/go/cue/format.(*formatter).after(0xc00055bd40, {0x18e9c28?, 0xc0004c99f0?})
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/format.go:305 +0xd1
cuelang.org/go/cue/format.(*formatter).expr1(0xc0003d1900?, {0x18e9c28?, 0xc0004c99f0}, 0x1062af5?, 0x0?)
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/node.go:524 +0x9c
cuelang.org/go/cue/format.(*formatter).expr(...)
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/node.go:513
cuelang.org/go/cue/format.(*formatter).selectorExpr(0xc00055bd40, 0xc0004d8e70, 0xc0001820c0?)
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/node.go:908 +0x316
cuelang.org/go/cue/format.(*formatter).exprRaw(0xc00055bd40, {0x18e9de0?, 0xc0004d8e70?}, 0x16f5060?, 0x1)
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/node.go:587 +0x526
cuelang.org/go/cue/format.(*formatter).walkListElems(0xc00055bd40, {0xc000485400, 0x3, 0xc0000061a0?})
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/node.go:208 +0x285
cuelang.org/go/cue/format.(*formatter).exprRaw(0xc00055bd40, {0x18e9d30?, 0xc0004c9ae0?}, 0x1013e45?, 0x1)
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/node.go:663 +0xa69
cuelang.org/go/cue/format.(*formatter).expr1(0x1cef110?, {0x18e9d30?, 0xc0004c9ae0}, 0xc000117340?, 0xc0004baaa8?)
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/node.go:522 +0x79
cuelang.org/go/cue/format.(*formatter).expr(0xc0003d1900?, {0x18e9d30?, 0xc0004c9ae0?})
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/node.go:513 +0x2a
cuelang.org/go/cue/format.(*formatter).decl(0xc00055bd40, {0x18e8900, 0xc0004baa80?})
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/node.go:316 +0x2135
cuelang.org/go/cue/format.(*formatter).walkDeclList(0xc00055bd40, {0xc000479880, 0x2, 0xc00035fc01?})
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/node.go:134 +0x32b
cuelang.org/go/cue/format.(*formatter).file(0xc00055bd40, 0xc0004f2660)
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/node.go:229 +0x4c
cuelang.org/go/cue/format.printNode({0x1796b40?, 0xc0004f2660?}, 0xc0003d1900)
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/node.go:39 +0x157
cuelang.org/go/cue/format.(*config).fprint(0xc0004b1740, {0x1796b40, 0xc0004f2660})
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/format.go:145 +0xcd
cuelang.org/go/cue/format.Node({0x18e6fe0?, 0xc0004f2660}, {0x0, 0x0, 0xc00035fee0?})
	/Users/rudifarkas/GitHub/golang/src/cue/cue/format/format.go:92 +0x98
main.runParseAndFormat({0x17ea4ce, 0x18}, 0x0)
	/Users/rudifarkas/GitHub/golang/src/cue-issues-fmt-comments/main.go:107 +0x245
main.runParseAndFormat_2274(...)
	/Users/rudifarkas/GitHub/golang/src/cue-issues-fmt-comments/main.go:41
main.main()
	/Users/rudifarkas/GitHub/golang/src/cue-issues-fmt-comments/main.go:28 +0x69`

func TestNewStacktrace(t *testing.T) {
	st := stacktrace.NewStacktraceFrom(longerStackTrace)
	oneline := st.OnelineString()
	// fmt.Printf("oneline: %s\n", oneline)
	wanted := `main.main => main.runParseAndFormat_2274 => main.runParseAndFormat => format.Node => format.(*config).fprint => format.printNode => format.(*formatter).file => format.(*formatter).walkDeclList => format.(*formatter).decl => format.(*formatter).expr => format.(*formatter).expr1 => format.(*formatter).exprRaw => format.(*formatter).walkListElems => format.(*formatter).exprRaw => format.(*formatter).selectorExpr => format.(*formatter).expr => format.(*formatter).expr1 => format.(*formatter).after => format.(*formatter).visitComments => format.(*formatter).printComment => format.(*printer).Print => format.(*printer).writeString => format.(*printer).append => format.LogStackTrace`
	if oneline != wanted {
		t.Errorf("wanted |%s|, got |%s|", wanted, oneline)
	}
}
