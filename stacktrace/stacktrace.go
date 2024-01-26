/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

package stacktrace

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"

	"github.com/rudifa/goutil/util"
)

/* Usage example:

func LogStackOneline() {
	// create a stacktrace instance (capturing the stack at this point)
	st := stacktrace.CapturedStacktrace()

	// get the stacktrace as a one-line string with function `Callpoints`, similar to this:
	// "main main.go:26 => runParseAndFormat_2567 main.go:37 => runParseAndFormat main.go:104 => writeString printer.go:365 => append printer.go:454 => LogStackTrace printer.go:506"

	callpoints := st.StacklineCallpoints()
	log.Printf("StacklineCallpoints: %s\n", callpoints)

	// and/or get the stacktrace as a one-line string with `Callersigs`, similar to this:
	// "main.main => main.runParseAndFormat_2567 => main.runParseAndFormat => format.(*printer).writeString => format.(*printer).append => format.LogStackTrace"

	callersigs := st.StacklineCallersigs()
	log.Printf("StacklineCallersigs: %s", callersigs)
}
*/

// stacktrace api ----------------------------------------------------------------------------------

// Stacktrace represents a stacktrace and has transform methods
type Stacktrace struct {
	// each pair of raw lines from runtime.Stack() describing a call is joined into a single line
	// the pairs are in the original order, i.e. the tip is at index 0
	PairedRawLines []string
}

// CapturedStacktrace returns a new instance of Stacktrace
// initialized from runtime.Stack() at the point where it was called
func CapturedStacktrace() *Stacktrace {
	buf := make([]byte, 65536)
	n := runtime.Stack(buf, false)
	if n >= len(buf) {
		// The buffer is full, and some data might have been lost.
		panic("stacktrace.CapturedStacktrace: stacktrace buffer is full")
	}
	return NewStacktraceFrom(string(buf[:n])) // Only convert the part of the buffer that was written to.
}

// NewStacktraceFrom returns a new instance of Stacktrace
// initialized from a string obtained from runtime.Stack()
func NewStacktraceFrom(rawString string) *Stacktrace {
	pairedLines := strings.Split(JoinLinePairs(rawString), "\n")
	return &Stacktrace{pairedLines}
}

// Trim removes atRoot lines from the root of the stacktrace and
// atTip lines from the tip of the stacktrace
func (st *Stacktrace) Trim(atRoot, atTip int) error {
	// Protect against invalid indices
	if atRoot < 0 || atTip < 0 || atRoot+atTip > len(st.PairedRawLines) {
		return fmt.Errorf("stacktrace.Trim: invalid indices: atRoot=%d, atTip=%d, len(st.PairedRawLines)=%d", atRoot, atTip, len(st.PairedRawLines))
	}

	// Remove atTip lines from the tip of the stacktrace
	st.PairedRawLines = st.PairedRawLines[atTip:]

	// Remove atRoot lines from the root of the stacktrace
	st.PairedRawLines = st.PairedRawLines[:len(st.PairedRawLines)-atRoot]

	return nil
}

// Len returns the number of calls in the stacktrace
func (st *Stacktrace) Len() int {
	return len(st.PairedRawLines)
}

// StacklineCallersigs returns a one-line string representation of the stacktrace
// with the function signatures (path (receiver) funcname) for each call
// captured at the point of creation of the Stacktrace instance, e.g.
// "main.main => main.runParseAndFormat_2274 => main.runParseAndFormat => format.Node => format.(*config).fprint =>..."
func (st Stacktrace) StacklineCallersigs() string {
	pairedRawLines := st.PairedRawLines
	return StacklineCallersigs(pairedRawLines)
}

// StacklineCallersigs returns an one-line string representation of the stacktrace, e.g.
// "main.main => main.runParseAndFormat_2567 => main.runParseAndFormat => format.(*printer).writeString => format.(*printer).append => format.LogStackTrace"
func StacklineCallersigs(rawLines []string) string {
	funcSigs := []string{}
	for _, line := range rawLines {
		funcSig := ExtractFuncSignature(line)
		if funcSig != "" {
			funcSigs = append(funcSigs, funcSig)
		}
	}
	util.ReverseSlice(funcSigs)
	return strings.Join(funcSigs, " => ")
}

// StacklineCallpoints returns an one-line string representation of the stacktrace
// with the calling function name and file:line for each call
// captured at the point of creation of the Stacktrace instance, e.g.
// "main main.go:26 => runParseAndFormat_2567 main.go:37 => runParseAndFormat main.go:104 => writeString printer.go:365 => append printer.go:454 => LogStackTrace printer.go:506"
func (st Stacktrace) StacklineCallpoints(short ...bool) string {
	return StacklineCallpoints(st.PairedRawLines, short...)
}

// StacklineCallpoints returns an one-line string representation of the stacktrace
// with the calling function name and file:line for each call
func StacklineCallpoints(pairedLines []string, short ...bool) string {
	funcCallSigs := []string{}
	for _, line := range pairedLines {
		funcSig := ExtractCallpoint(line)
		if funcSig != "" {
			funcCallSigs = append(funcCallSigs, funcSig)
		}
	}
	util.ReverseSlice(funcCallSigs)
	if len(short) > 0 && short[0] {
		funcCallSigs = shorten(funcCallSigs)
	}
	return strings.Join(funcCallSigs, " => ")
}

// stacktrace utilities ----------------------------------------------------------------------------

// ExtractFuncSignature extracts the function signature from a stack line
// that was generated by runtime.Stack(), e.g.
// "format.(*printer).writeString"
func ExtractFuncSignature(line string) string {
	r := regexp.MustCompile(`((\w+)\.((\(\*\w+\))\.)?(\w+))\((.*?)\)`)
	matches := r.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// ExtractCallsite returns a `Callsite` from a stack line, e.g. "printer.go:506"
func ExtractCallsite(line string) string {
	r := regexp.MustCompile(`(\w+\.go:\d+)`)
	matches := r.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// ExtractCallpoint returns a `Callpoint` e.g. "LogStackTrace printer.go:506"
func ExtractCallpoint(line string) (callpoint string) {
	// TODO combine the regex into one
	re1 := regexp.MustCompile(`((\w+)\.((\(\*\w+\))\.)?(\w+))\((.*?)\)`)
	re2 := regexp.MustCompile(`.+(\b\w+\.go:\d+)`)

	matches1 := re1.FindStringSubmatch(line)
	// log.Printf("len(matches)=%d\n", len(matches1))
	// for i, match := range matches1 {
	// 	log.Printf("matches[%d]=%s\n", i, match)
	// }
	matches2 := re2.FindStringSubmatch(line)
	// log.Printf("len(matches)=%d\n", len(matches2))
	// for i, match := range matches2 {
	// 	log.Printf("matches[%d]=%s\n", i, match)
	// }

	if len(matches1) > 5 && len(matches2) > 1 {
		return matches1[5] + ` ` + matches2[1]
	}
	return ""
}

// JoinLinePairs returns a string with each pair of lines describing a call joined into a single line, e.g.
// "cuelang.org/go/cue/format.LogStackTrace() /Users/rudifarkas/GitHub/golang/src/cue/cue/format/printer.go:506 +0x3a"
func JoinLinePairs(rawString string) string {
	const (
		pattern1 = `(\w+.\w+.+)`
		pattern2 = `\n\s+`
		pattern3 = `(.+)`
		pattern  = pattern1 + pattern2 + pattern3
	)
	// re should match pairs of lines in rawString from runtime.Stack(buf, false)
	re := regexp.MustCompile(pattern)

	// log.Printf("rawString=\n%s\n", rawString)

	// in every occurrence of re in rawString, replace pattern2 by ` `
	result := re.ReplaceAllStringFunc(rawString, func(s string) string {
		matches := re.FindStringSubmatch(s)
		// log.Printf("len(matches)=%d\n", len(matches))

		// for i, match := range matches {
		// 	log.Printf("matches[%d]=%s\n", i, match)
		// }
		if len(matches) > 0 {
			// log.Printf("matches[0]=%s\n", matches[0])
			return matches[1] + " " + matches[len(matches)-1]
		}
		return s
	})

	return result
}

// shorten returns a slice of strings with the filenames removed from the callpoints
// when repeated from previous callpoint
func shorten(sigs []string) []string {

	outsigs := []string{}
	prevFilename := ""

	for _, sig := range sigs {
		// Use a regex to extract the filename and the rest into separate strings from sig
		re := regexp.MustCompile(`(.+)\s(.+):(\d+)`)
		matches := re.FindStringSubmatch(sig)

		if len(matches) < 4 {
			// log.Printf("sig does not match the expected format: %s\n", sig)
			continue
		}

		funcName, filename, lineNumber := matches[1], matches[2], matches[3]

		// If the filename is the same as prevFilename, append the rest to shortsigs
		// Else append the full sig to shortsigs and update prevFilename
		if filename == prevFilename {
			outsigs = append(outsigs, fmt.Sprintf("%s :%s", funcName, lineNumber))
		} else {
			outsigs = append(outsigs, sig)
			prevFilename = filename
		}
	}
	return outsigs
}
