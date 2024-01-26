/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package stacktrace_test

import (
	"log"
	"strings"
	"testing"

	"github.com/rudifa/goutil/stacktrace"
)

const verbose = false

func TestStacktrace(t *testing.T) {

	log.SetFlags(log.Lshortfile)

	// capture stacktrace from runtime.Stack()
	st := stacktrace.CapturedStacktrace()

	if verbose {
		log.Printf("stacktrace PairedRawLines:\n")
		for _, line := range st.PairedRawLines {
			log.Printf("  %s\n", line)
		}
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

const sampleStackTracePairsJoined = `cuelang.org/go/cue/format.LogStackTrace() /Users/rudifarkas/GitHub/golang/src/cue/cue/format/printer.go:506 +0x3a
cuelang.org/go/cue/format.(*printer).append(0xc00031b2c0, {0x17c374c, 0x6}, {0xc0002dee40?, 0x1, 0x11a2d25?}) /Users/rudifarkas/GitHub/golang/src/cue/cue/format/printer.go:454 +0x4aa
cuelang.org/go/cue/format.(*printer).writeString(0xc00031b2c0, {0x18ce528, 0x1}, 0x0) /Users/rudifarkas/GitHub/golang/src/cue/cue/format/printer.go:365 +0xc7
main.runParseAndFormat({0x1bd7719, 0x13}, 0x0) /Users/rudifarkas/GitHub/golang/src/cue-issues-fmt-comments/main.go:104 +0x4f2
main.runParseAndFormat_2567() /Users/rudifarkas/GitHub/golang/src/cue-issues-fmt-comments/main.go:37 +0x25
main.main() /Users/rudifarkas/GitHub/golang/src/cue-issues-fmt-comments/main.go:26 +0x58

`
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

// stacktrace api tests ----------------------------------------------------------------------------

func TestNewStacktraceFrom(t *testing.T) {
	st := stacktrace.NewStacktraceFrom(longerStackTrace)
	got := st.StacklineCallersigs()
	want := `main.main => main.runParseAndFormat_2274 => main.runParseAndFormat => format.Node => format.(*config).fprint => format.printNode => format.(*formatter).file => format.(*formatter).walkDeclList => format.(*formatter).decl => format.(*formatter).expr => format.(*formatter).expr1 => format.(*formatter).exprRaw => format.(*formatter).walkListElems => format.(*formatter).exprRaw => format.(*formatter).selectorExpr => format.(*formatter).expr => format.(*formatter).expr1 => format.(*formatter).after => format.(*formatter).visitComments => format.(*formatter).printComment => format.(*printer).Print => format.(*printer).writeString => format.(*printer).append => format.LogStackTrace`
	if want != got {
		t.Errorf("want |%s|\ngot |%s|\n", want, got)
	}
	if verbose {
		log.Printf("StacklineCallersigs returns %s\n", got)
	}

	// also test Len()
	wantlen := 26
	gotlen := st.Len()
	if wantlen != gotlen {
		t.Errorf("want |%d|, got |%d|", wantlen, gotlen)
	}
	if verbose {
		log.Printf("Len returns %d\n", gotlen)
	}
}

func TestTrim(t *testing.T) {
	st := stacktrace.NewStacktraceFrom(sampleStackTrace)
	{
		got := st.StacklineCallersigs()
		want := `main.main => main.runParseAndFormat_2567 => main.runParseAndFormat => format.(*printer).writeString => format.(*printer).append => format.LogStackTrace`
		if want != got {
			t.Errorf("want |%s|, got |%s|", want, got)
		}
		if verbose {
			log.Printf("StacklineCallersigs returns %s\n", got)
		}
	}
	{
		err := st.Trim(2, 2)
		if err != nil {
			t.Errorf("Trim returned error: %v", err)
		}
		got := st.StacklineCallersigs()
		want := `main.runParseAndFormat => format.(*printer).writeString`
		if want != got {
			t.Errorf("want |%s|, got |%s|", want, got)
		}
		if verbose {
			log.Printf("StacklineCallersigs returns %s\n", got)
		}
	}
	{
		err := st.Trim(3, 1)
		if err != nil {
			const expectedError = "stacktrace.Trim: invalid indices: atRoot=3, atTip=1, len(st.PairedRawLines)=2"
			if err.Error() != expectedError {
				t.Errorf("expected error message '%s', got '%s'", expectedError, err.Error())
			}
		} else {
			t.Error("expected an error, got nil")
		}
		// expect unchanged from previous test
		got := st.StacklineCallersigs()
		want := `main.runParseAndFormat => format.(*printer).writeString`
		if want != got {
			t.Errorf("want |%s|, got |%s|", want, got)
		}
		if verbose {
			log.Printf("StacklineCallersigs returns %s\n", got)
		}
	}
	{
		err := st.Trim(1, 1)
		if err != nil {
			t.Errorf("Trim returned error: %v", err)
		}
		got := st.StacklineCallersigs()
		want := ``
		if want != got {
			t.Errorf("want |%s|, got |%s|", want, got)
		}
		if verbose {
			log.Printf("StacklineCallersigs returns %s\n", got)
		}
	}
}

func TestStacklineCallersigs(t *testing.T) {

	lines := strings.Split(sampleStackTrace, "\n")
	got := stacktrace.StacklineCallersigs(lines)
	want := `main.main => main.runParseAndFormat_2567 => main.runParseAndFormat => format.(*printer).writeString => format.(*printer).append => format.LogStackTrace`

	if want != got {
		t.Errorf("want |%s|, got |%s|", want, got)
	}
	if verbose {
		log.Printf("StacklineCallersigs returns %s\n", got)
	}
}

func TestStacklineCallpoints(t *testing.T) {

	wantCallsFromSampleStackTrace := `main main.go:26 => runParseAndFormat_2567 main.go:37 => runParseAndFormat main.go:104 => writeString printer.go:365 => append printer.go:454 => LogStackTrace printer.go:506`

	lines := strings.Split(sampleStackTracePairsJoined, "\n")
	got := stacktrace.StacklineCallpoints(lines)

	if wantCallsFromSampleStackTrace != got {
		t.Errorf("want |%s|, got |%s|", wantCallsFromSampleStackTrace, got)
	}
	if verbose {
		log.Printf("StacklineCallpoints returns %s\n", got)

	}
}

func TestStacklineCallpointsShort(t *testing.T) {

	wantCallsFromSampleStackTrace := `main main.go:26 => runParseAndFormat_2567 :37 => runParseAndFormat :104 => writeString printer.go:365 => append :454 => LogStackTrace :506`

	lines := strings.Split(sampleStackTracePairsJoined, "\n")
	got := stacktrace.StacklineCallpoints(lines, true)

	if wantCallsFromSampleStackTrace != got {
		t.Errorf("want |%s|, got |%s|", wantCallsFromSampleStackTrace, got)
	}
	if verbose {
		log.Printf("StacklineCallpoints returns %s\n", got)

	}
}
// stacktrace utilities tests -----------------------------------------------------------------------

func TestJoinLinePairs(t *testing.T) {
	gotJoinedPairs := stacktrace.JoinLinePairs(sampleStackTrace)

	for _, line := range strings.Split(gotJoinedPairs, "\n") {
		log.Printf("line:|%v|\n", line)
	}

	for i, got := range strings.Split(gotJoinedPairs, "\n") {
		want := strings.Split(sampleStackTracePairsJoined, "\n")[i]
		if want != got {
			t.Errorf("%d want |%s|, got |%s|", i, want, got)
		}
		if verbose {
			log.Printf("JoinLinePairs returns %s\n", got)
		}
	}
}

func TestExtractFuncSignature(t *testing.T) {

	wantSignaturesFromSampleStackTrace := `format.LogStackTrace

format.(*printer).append

format.(*printer).writeString

main.runParseAndFormat

main.runParseAndFormat_2567

main.main
`

	lines := strings.Split(sampleStackTrace, "\n")

	wantSignatures := strings.Split(wantSignaturesFromSampleStackTrace, "\n")

	for i, line := range lines {
		got := stacktrace.ExtractFuncSignature(line)
		want := wantSignatures[i]

		if want != got {
			t.Errorf("%d want |%s|, got |%s|", i, want, got)
		}
		if verbose {
			log.Printf("ExtractFuncSignature returns %s\n", got)
		}
	}
}

func TestExtractCallsite(t *testing.T) {

	wantCallsitesFromSampleStackTrace := `
printer.go:506

printer.go:454

printer.go:365

main.go:104

main.go:37

main.go:26`

	lines := strings.Split(sampleStackTrace, "\n")

	log.Printf("wantCallsitesFromSampleStackTrace:\n%s\n", wantCallsitesFromSampleStackTrace)

	wantCallsites := strings.Split(wantCallsitesFromSampleStackTrace, "\n")

	for i, line := range lines {
		got := stacktrace.ExtractCallsite(line)
		want := wantCallsites[i]

		if want != got {
			t.Errorf("%d want |%s|, got |%s|", i, want, got)
		}
		if verbose {
			log.Printf("ExtractCallsite returns %s\n", got)
		}
	}
}

func TestExtractCallpoint(t *testing.T) {
	got := stacktrace.ExtractCallpoint(sampleStackTrace2)
	want := "LogStackTrace printer.go:506"
	if want != got {
		t.Errorf("want |%s|, got |%s|", want, got)
	}

	log.Printf("ExtractCallpoint returns %s\n", got)
}
